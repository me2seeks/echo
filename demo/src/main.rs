use actix::prelude::*;
use actix::Actor;
use actix_web::{web, App, Error, HttpRequest, HttpResponse, HttpServer};
use actix_web_actors::ws;
use serde::{Deserialize, Serialize};
use serde_json;
use std::collections::HashMap;
use std::sync::Arc;
use webrtc::api::interceptor_registry::register_default_interceptors;
use webrtc::api::media_engine::MediaEngine;
use webrtc::api::APIBuilder;
use webrtc::ice_transport::ice_candidate::RTCIceCandidateInit;
use webrtc::ice_transport::ice_server::RTCIceServer;
use webrtc::interceptor::registry::Registry;
use webrtc::peer_connection::configuration::RTCConfiguration;
use webrtc::peer_connection::peer_connection_state::RTCPeerConnectionState;
use webrtc::peer_connection::sdp::session_description::RTCSessionDescription;
use webrtc::peer_connection::RTCPeerConnection;

#[derive(Debug, Serialize, Deserialize)]
#[serde(tag = "event", content = "payload")]
pub enum WSMessage {
    #[serde(rename = "webrtc")]
    WebRTC(RTCSessionDescription),

    #[serde(rename = "candidate")]
    Candidate(RTCIceCandidateInit),
}

struct Room {
    broadcaster: Option<Addr<WebRTCActor>>,
    viewers: HashMap<String, Addr<WebRTCActor>>,
    viewer_count: u32,
}

impl Room {
    fn new() -> Self {
        Room {
            broadcaster: None,
            viewers: HashMap::new(),
            viewer_count: 0,
        }
    }
}

impl Actor for Room {
    type Context = Context<Self>;
}

#[derive(Message)]
#[rtype(result = "()")]
struct JoinRoom {
    addr: Addr<WebRTCActor>,
    is_broadcaster: bool,
}

impl Handler<JoinRoom> for Room {
    type Result = ();

    fn handle(&mut self, msg: JoinRoom, _: &mut Context<Self>) {
        if msg.is_broadcaster {
            self.broadcaster = Some(msg.addr);
        } else {
            self.viewers
                .insert(format!("viewer-{}", self.viewer_count), msg.addr);
            self.viewer_count += 1;
        }
    }
}

#[derive(Message, Serialize)]
#[rtype(result = "()")]
struct WebRTCMessage {
    sdp: RTCSessionDescription,
}

impl Handler<WebRTCMessage> for WebRTCActor {
    type Result = ();

    fn handle(&mut self, msg: WebRTCMessage, ctx: &mut Self::Context) {
        ctx.text(serde_json::to_string(&WSMessage::WebRTC(msg.sdp)).unwrap());
    }
}

struct WebRTCActor {
    pc: Arc<RTCPeerConnection>,
    room: Addr<Room>,
    is_broadcaster: bool,
}

impl WebRTCActor {
    async fn new(room: Addr<Room>, is_broadcaster: bool) -> Result<Self, webrtc::Error> {
        // Prepare the configuration
        let config = RTCConfiguration {
            ice_servers: vec![RTCIceServer {
                urls: vec!["stun:95.216.190.5:3478".to_owned()],
                ..Default::default()
            }],
            ..Default::default()
        };

        // Create a MediaEngine object to configure the supported codec
        let mut m = MediaEngine::default();
        m.register_default_codecs()?;

        let mut registry = Registry::new();

        // Use the default set of Interceptors
        registry = register_default_interceptors(registry, &mut m).unwrap();

        // Create the API object with the MediaEngine
        let api = APIBuilder::new()
            .with_media_engine(m)
            .with_interceptor_registry(registry)
            .build();
        let pc = api.new_peer_connection(config).await.unwrap();
        pc.on_peer_connection_state_change(Box::new(move |state: RTCPeerConnectionState| {
            println!("Peer Connection State has changed: {:?}", state);
            Box::pin(async {})
        }));

        Ok(WebRTCActor {
            pc: Arc::new(pc),
            room,
            is_broadcaster,
        })
    }
}

impl Actor for WebRTCActor {
    type Context = ws::WebsocketContext<Self>;

    fn started(&mut self, _: &mut Self::Context) {
        println!("WebRTCActor started");
    }

    fn stopped(&mut self, ctx: &mut Self::Context) {
        if self.pc.connection_state()
            == webrtc::peer_connection::peer_connection_state::RTCPeerConnectionState::Connected
        {
            self.room.do_send(JoinRoom {
                addr: ctx.address(),
                is_broadcaster: self.is_broadcaster,
            });
        }
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for WebRTCActor {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(ws::Message::Text(text)) => {
                println!("Received text message: {}", text);
                let pc = self.pc.clone();
                let addr = ctx.address();
                ctx.spawn(
                    async move {
                        match serde_json::from_str::<WSMessage>(&text) {
                            Ok(ws_message) => match ws_message {
                                WSMessage::WebRTC(sdp) => match sdp.sdp_type {
                                    webrtc::peer_connection::sdp::sdp_type::RTCSdpType::Offer => {
                                        println!("Received WebRTC offer: {}", sdp.sdp);
                                        if let Err(err) = pc.set_remote_description(sdp).await {
                                            println!("Error setting remote description: {:?}", err);
                                            return;
                                        }
                                        let answer = match pc.create_answer(None).await {
                                            Ok(answer) => answer,
                                            Err(err) => {
                                                println!("Error creating answer: {:?}", err);
                                                return;
                                            }
                                        };
                                        if let Err(err) =
                                            pc.set_local_description(answer.clone()).await
                                        {
                                            println!("Error setting local description: {:?}", err);
                                            return;
                                        }
                                        let answer = match pc.local_description().await {
                                            Some(answer) => answer,
                                            None => {
                                                println!("Error getting local description");
                                                return;
                                            }
                                        };
                                        addr.send(WebRTCMessage { sdp: answer }).await.unwrap();
                                    }
                                    _ => {}
                                },
                                WSMessage::Candidate(candidate) => {
                                    println!("Received WebRTC candidate: {:?}", candidate);
                                    if let Err(err) = pc.add_ice_candidate(candidate).await {
                                        println!("Error adding ICE candidate: {:?}", err);
                                    }
                                }
                            },
                            Err(err) => {
                                println!("Error parsing message: {:?}", err);
                            }
                        }
                    }
                    .into_actor(self),
                );
            }
            Ok(ws::Message::Binary(bin)) => {
                println!("Received binary message: {:?}", bin);
            }
            Ok(ws::Message::Ping(msg)) => {
                ctx.pong(&msg);
            }
            Ok(ws::Message::Pong(_)) => {
                println!("Received pong");
            }
            Ok(ws::Message::Close(reason)) => {
                ctx.close(reason);
                ctx.stop();
            }
            _ => (),
        }
    }
}

async fn webrtc_handler(
    req: HttpRequest,
    stream: web::Payload,
    room: web::Data<Addr<Room>>,
    is_broadcaster: bool,
) -> Result<HttpResponse, Error> {
    ws::start(
        WebRTCActor::new(room.get_ref().clone(), is_broadcaster)
            .await
            .unwrap(),
        &req,
        stream,
    )
}

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    let room = Room::new().start();

    HttpServer::new(move || {
        App::new()
            .app_data(web::Data::new(room.clone()))
            .route(
                "/ws/broadcaster",
                web::get().to(|req, stream, room| webrtc_handler(req, stream, room, true)),
            )
            .route(
                "/ws/viewer",
                web::get().to(|req, stream, room| webrtc_handler(req, stream, room, false)),
            )
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}
