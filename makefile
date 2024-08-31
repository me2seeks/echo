# 定义变量
ifndef GOPATH
	GOPATH := $(shell go env GOPATH)
endif

GOBIN=$(GOPATH)/bin
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) mod tidy

GOCTL=$(GOBIN)/goctl ## goctl

# 安装goctl代码生成工具
$(shell if [ ! -d $(GOCTL) ]; then \
	$(GOCMD) install github.com/zeromicro/go-zero/tools/goctl@latest; \
fi; \
)


clean: ## 清理目标
	$(GOCLEAN)
	rm -rf target

deps: ## 安装依赖目标
	@export GOPROXY=https://goproxy.cn,direct
	$(GOGET) -v

copy_config:
	mkdir -p target/usercenter-rpc && cp app/usercenter/cmd/rpc/etc/usercenter.yaml target/usercenter-rpc
	mkdir -p target/usercenter-api && cp app/usercenter/cmd/api/etc/usercenter.yaml target/usercenter-api
	mkdir -p target/content-rpc && cp app/content/cmd/rpc/etc/content.yaml target/content-rpc
	mkdir -p target/content-api && cp app/content/cmd/api/etc/content.yaml target/content-api
	mkdir -p target/interaction-rpc && cp app/interaction/cmd/rpc/etc/interaction.yaml target/interaction-rpc
	mkdir -p target/interaction-api && cp app/interaction/cmd/api/etc/interaction.yaml target/interaction-api 
	mkdir -p target/counter-rpc && cp app/counter/cmd/rpc/etc/counter.yaml target/counter-rpc
	mkdir -p target/counter-consumer && cp app/counter/cmd/consumer/etc/consumer.yaml target/counter-consumer
	mkdir -p target/counter-api && cp app/counter/cmd/api/etc/counter.yaml target/counter-api
	mkdir -p target/search-rpc && cp app/search/cmd/rpc/etc/search.yaml target/search-rpc
	mkdir -p target/search-consumer && cp app/search/cmd/consumer/etc/consumer.yaml target/search-consumer
	mkdir -p target/search-api && cp app/search/cmd/api/etc/search.yaml target/search-api

build: copy_config ## 构建目标
	$(GOBUILD) -o target/usercenter-rpc/usercenter-rpc app/usercenter/cmd/rpc/usercenter.go
	$(GOBUILD) -o target/usercenter-api/usercenter-api app/usercenter/cmd/api/usercenter.go
	$(GOBUILD) -o target/content-rpc/content-rpc app/content/cmd/rpc/content.go
	$(GOBUILD) -o target/content-api/content-api app/content/cmd/api/content.go
	$(GOBUILD) -o target/interaction-rpc/interaction-rpc app/interaction/cmd/rpc/interaction.go
	$(GOBUILD) -o target/interaction-api/interaction-api app/interaction/cmd/api/interaction.go
	$(GOBUILD) -o target/counter-rpc/counter-rpc app/counter/cmd/rpc/counter.go
	$(GOBUILD) -o target/counter-consumer/counter-consumer app/counter/cmd/consumer/consumer.go
	$(GOBUILD) -o target/counter-api/counter-api app/counter/cmd/api/counter.go
	$(GOBUILD) -o target/search-rpc/search-rpc app/search/cmd/rpc/search.go
	$(GOBUILD) -o target/search-consumer/search-consumer app/search/cmd/consumer/consumer.go
	$(GOBUILD) -o target/search-api/search-api app/search/cmd/api/search.go



start: ## 运行目标
	nohup ./target/usercenter-rpc/usercenter-rpc -f ./target/usercenter-rpc/usercenter.yaml  > /dev/null 2>&1 &
	nohup ./target/usercenter-api/usercenter-api -f ./target/usercenter-api/usercenter.yaml  > /dev/null 2>&1 &
	nohup ./target/content-rpc/content-rpc -f ./target/content-rpc/content.yaml  > /dev/null 2>&1 &
	nohup ./target/content-api/content-api -f ./target/content-api/content.yaml  > /dev/null 2>&1 &
	nohup ./target/interaction-rpc/interaction-rpc -f ./target/interaction-rpc/interaction.yaml  > /dev/null 2>&1 &
	nohup ./target/interaction-api/interaction-api -f ./target/interaction-api/interaction.yaml  > /dev/null 2>&1 &
	nohup ./target/counter-rpc/counter-rpc -f ./target/counter-rpc/counter.yaml  > /dev/null 2>&1 &
	nohup ./target/counter-consumer/counter-consumer -f ./target/counter-consumer/consumer.yaml  > /dev/null 2>&1 &
	nohup ./target/counter-api/counter-api -f ./target/counter-api/counter.yaml  > /dev/null 2>&1 &
	nohup ./target/search-rpc/search-rpc -f ./target/search-rpc/search.yaml  > /dev/null 2>&1 &
	nohup ./target/search-consumer/search-consumer -f ./target/search-consumer/consumer.yaml  > /dev/null 2>&1 &
	nohup ./target/search-api/search-api -f ./target/search-api/search.yaml  > /dev/null 2>&1 &

stop: ## 停止目标
	-pkill -f usercenter-rpc                                                                                                -pkill -f usercenter-api
	-pkill -f content-rpc
	-pkill -f content-api
	-pkill -f interaction-rpc
	-pkill -f interaction-api
	-pkill -f counter-rpc
	-pkill -f counter-consumer
	-pkill -f counter-api
	-pkill -f search-rpc                                                                                                    -pkill -f search-consumer
	-pkill -f search-api
	@for i in 5 4 3 2 1; do \
	echo -n "stop $$i"; \
	sleep 1; \
	echo " "; \
done

restart: stop start ## 重启项目


.DEFAULT_GOAL := all ## 默认构建目标