use actix::prelude::*;
use actix::Actor;
use actix_web::{web, App, Error, HttpRequest, HttpResponse, HttpServer};
use actix_web_actors::ws;
use serde::{Deserialize, Serialize};
use serde_json;
use std::collections::HashMap;

use webrtc::ice_transport::ice_candidate::RTCIceCandidate;
use webrtc::peer_connection::sdp::session_description::RTCSessionDescription;

#[derive(Debug, Serialize, Deserialize)]
pub struct WSMessage {
    pub event: String,
    pub payload: Payload,
}

#[derive(Debug, Serialize, Deserialize)]
#[serde(tag = "event", content = "payload")]
pub enum Payload {
    #[serde(rename = "webrtc")]
    WebRTC(RTCSessionDescription),

    #[serde(rename = "candidate")]
    Candidate(RTCIceCandidate),
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

struct WebRTCActor {
    room: Addr<Room>,
    is_broadcaster: bool,
}

impl WebRTCActor {
    fn new(room: Addr<Room>, is_broadcaster: bool) -> Self {
        WebRTCActor {
            room,
            is_broadcaster,
        }
    }
}

impl Actor for WebRTCActor {
    type Context = ws::WebsocketContext<Self>;
    fn started(&mut self, ctx: &mut Self::Context) {
        self.room.do_send(JoinRoom {
            addr: ctx.address(),
            is_broadcaster: self.is_broadcaster,
        });
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for WebRTCActor {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(ws::Message::Text(text)) => {
                println!("Received text message: {}", text);
                match serde_json::from_str::<WSMessage>(&text) {
                    Ok(ws_message) => {
                        println!("{:?}", ws_message);
                        match ws_message.event.as_str() {
                            "webrtc" => match ws_message.payload {
                                Payload::WebRTC(sdp) => match sdp.sdp_type {
                                    webrtc::peer_connection::sdp::sdp_type::RTCSdpType::Offer => {
                                        println!("Received WebRTC offer: {}", sdp.sdp);
                                    }
                                    webrtc::peer_connection::sdp::sdp_type::RTCSdpType::Answer => {
                                        println!("Received WebRTC answer: {}", sdp.sdp);
                                    }
                                    _ => {}
                                },
                                _ => {}
                            },
                            "candidate" => match ws_message.payload {
                                Payload::Candidate(candidate) => {
                                    println!("Received WebRTC candidate: {}", candidate);
                                }
                                _ => {}
                            },
                            _ => {}
                        }
                    }
                    Err(e) => {
                        println!("{:?}", e);
                    }
                }
            }
            Ok(ws::Message::Binary(bin)) => {
                ctx.binary(bin);
            }
            Ok(ws::Message::Ping(msg)) => {
                ctx.pong(&msg);
            }
            Ok(ws::Message::Close(reason)) => {
                println!("Received close message: {:?}", reason);
                ctx.stop();
            }
            _ => {}
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
        WebRTCActor::new(room.get_ref().clone(), is_broadcaster),
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
