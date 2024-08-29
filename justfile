# set shell
set windows-shell := ["powershell.exe", "-c"]

# set `&&` or `;` for different OS
and := if os_family() == "windows" {";"} else {"&&"}

#====================================== alias start ============================================#

#======================================= alias end =============================================#


#===================================== targets start ===========================================#

# default target - `just` 默认目标
default: fmt lint  test


[unix]
copy:
    @mkdir -p target/usercenter-rpc && cp app/usercenter/cmd/rpc/etc/usercenter.yaml target/usercenter-rpc
    @mkdir -p target/usercenter-api && cp app/usercenter/cmd/api/etc/usercenter.yaml target/usercenter-api
    @mkdir -p target/content-rpc && cp app/content/cmd/rpc/etc/content.yaml target/content-rpc
    @mkdir -p target/content-api && cp app/content/cmd/api/etc/content.yaml target/content-api
    @mkdir -p target/interaction-rpc && cp app/interaction/cmd/rpc/etc/interaction.yaml target/interaction-rpc
    @mkdir -p target/interaction-api && cp app/interaction/cmd/api/etc/interaction.yaml target/interaction-api
    @mkdir -p target/counter-rpc && cp app/counter/cmd/rpc/etc/counter.yaml target/counter-rpc
    @mkdir -p target/counter-consumer && cp app/counter/cmd/consumer/etc/consumer.yaml target/counter-consumer
    @mkdir -p target/counter-api && cp app/counter/cmd/api/etc/counter.yaml target/counter-api
    @mkdir -p target/search-rpc && cp app/search/cmd/rpc/etc/search.yaml target/search-rpc
    @mkdir -p target/search-consumer && cp app/search/cmd/consumer/etc/consumer.yaml target/search-consumer
    @mkdir -p target/search-api && cp app/search/cmd/api/etc/search.yaml target/search-api

[windows]
copy:
    New-Item -ItemType Directory -Force -Path target/usercenter-rpc
    Copy-Item -Force app/usercenter/cmd/rpc/etc/usercenter.yaml target/usercenter-rpc
    New-Item -ItemType Directory -Force -Path target/usercenter-api
    Copy-Item -Force app/usercenter/cmd/api/etc/usercenter.yaml target/usercenter-api
    New-Item -ItemType Directory -Force -Path target/content-rpc
    Copy-Item -Force app/content/cmd/rpc/etc/content.yaml target/content-rpc
    New-Item -ItemType Directory -Force -Path target/content-api
    Copy-Item -Force app/content/cmd/api/etc/content.yaml target/content-api
    New-Item -ItemType Directory -Force -Path target/interaction-rpc
    Copy-Item -Force app/interaction/cmd/rpc/etc/interaction.yaml target/interaction-rpc
    New-Item -ItemType Directory -Force -Path target/interaction-api
    Copy-Item -Force app/interaction/cmd/api/etc/interaction.yaml target/interaction-api
    New-Item -ItemType Directory -Force -Path target/counter-rpc
    Copy-Item -Force app/counter/cmd/rpc/etc/counter.yaml target/counter-rpc
    New-Item -ItemType Directory -Force -Path target/counter-consumer
    Copy-Item -Force app/counter/cmd/consumer/etc/consumer.yaml target/counter-consumer
    New-Item -ItemType Directory -Force -Path target/counter-api
    Copy-Item -Force app/counter/cmd/api/etc/counter.yaml target/counter-api
    New-Item -ItemType Directory -Force -Path target/search-rpc
    Copy-Item -Force app/search/cmd/rpc/etc/search.yaml target/search-rpc
    New-Item -ItemType Directory -Force -Path target/search-consumer
    Copy-Item -Force app/search/cmd/consumer/etc/consumer.yaml target/search-consumer
    New-Item -ItemType Directory -Force -Path target/search-api
    Copy-Item -Force app/search/cmd/api/etc/search.yaml target/search-api

[unix]
clean: ## 清理目标
    @go clean
    @rm -rf target

[windows]
clean:
    @Remove-Item -Recurse -Force target

[unix]
start:
    @nohup ./target/usercenter-rpc/usercenter-rpc -f ./target/usercenter-rpc/usercenter.yaml > /dev/null 2>&1 &
    @nohup ./target/usercenter-api/usercenter-api -f ./target/usercenter-api/usercenter.yaml > /dev/null 2>&1 &
    @nohup ./target/content-rpc/content-rpc -f ./target/content-rpc/content.yaml > /dev/null 2>&1 &
    @nohup ./target/content-api/content-api -f ./target/content-api/content.yaml > /dev/null 2>&1 &
    @nohup ./target/interaction-rpc/interaction-rpc -f ./target/interaction-rpc/interaction.yaml > /dev/null 2>&1 &
    @nohup ./target/interaction-api/interaction-api -f ./target/interaction-api/interaction.yaml > /dev/null 2>&1 &
    @nohup ./target/counter-rpc/counter-rpc -f ./target/counter-rpc/counter.yaml > /dev/null 2>&1 &
    @nohup ./target/counter-consumer/counter-consumer -f ./target/counter-consumer/counter.yaml > /dev/null 2>&1 &
    @nohup ./target/counter-api/counter-api -f ./target/counter-api/counter.yaml > /dev/null 2>&1 &
    @nohup ./target/search-rpc/search-rpc -f ./target/search-rpc/search.yaml > /dev/null 2>&1 &
    @nohup ./target/search-consumer/search-consumer -f ./target/search-consumer/search.yaml > /dev/null 2>&1 &
    @nohup ./target/search-api/search-api -f ./target/search-api/search.yaml > /dev/null 2>&1 &

[unix]
stop: ## 停止目标
    -pkill -f usercenter-rpc
    -pkill -f usercenter-api
    -pkill -f content-rpc
    -pkill -f content-api
    -pkill -f interaction-rpc
    -pkill -f interaction-api
    -pkill -f counter-rpc
    -pkill -f counter-consumer
    -pkill -f counter-api
    -pkill -f search-rpc
    -pkill -f search-consumer
    -pkill -f search-api
    @for i in 5 4 3 2 1; do \
        echo -n "stop $$i"; \
        sleep 1; \
        echo " "; \
    done


build: copy
    @echo "Building..."
    @go build -ldflags "-s -w" -o target/usercenter-rpc/usercenter-rpc app/usercenter/cmd/rpc/usercenter.go
    @go build -ldflags "-s -w" -o target/usercenter-api/usercenter-api app/usercenter/cmd/api/usercenter.go
    @go build -ldflags "-s -w" -o target/content-rpc/content-rpc app/content/cmd/rpc/content.go
    @go build -ldflags "-s -w" -o target/content-api/content-api app/content/cmd/api/content.go
    @go build -ldflags "-s -w" -o target/interaction-rpc/interaction-rpc app/interaction/cmd/rpc/interaction.go
    @go build -ldflags "-s -w" -o target/interaction-api/interaction-api app/interaction/cmd/api/interaction.go
    @go build -ldflags "-s -w" -o target/counter-rpc/counter-rpc app/counter/cmd/rpc/counter.go
    @go build -ldflags "-s -w" -o target/counter-consumer/counter-consumer app/counter/cmd/consumer/consumer.go
    @go build -ldflags "-s -w" -o target/counter-api/counter-api app/counter/cmd/api/counter.go
    @go build -ldflags "-s -w" -o target/search-rpc/search-rpc app/search/cmd/rpc/search.go
    @go build -ldflags "-s -w" -o target/search-consumer/search-consumer app/search/cmd/consumer/consumer.go
    @go build -ldflags "-s -w" -o target/search-api/search-api app/search/cmd/api/search.go
    @echo "Build done."


# go test
test:
    @go test -v {{join(".", "...")}}

# generate swagger docs - 生成 swagger 文档
# swag: dep-swag
#     @cd {{server}} {{and}} swag init -g swagger.go

# format code - 格式化代码
fmt: dep-gofumpt
    @echo "Formatting..."
    @gofumpt -w -extra .

# lint - 代码检查
lint: dep-golangci-lint
    @echo "Linting..."
    @go mod tidy 
    @golangci-lint run

# install dependencies - 安装依赖工具
dependencies:  dep-swag dep-golangci-lint dep-gofumpt

# a tool to help you write API docs - 一个帮助你编写 API 文档的工具
dep-swag:
    @go install github.com/swaggo/swag/cmd/swag@latest

# a linter for Go - 一个 Go 语言的代码检查工具
dep-golangci-lint:
    @go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# a stricter gofmt - 一个更严格的 gofmt
dep-gofumpt:
    @go install mvdan.cc/gofumpt@latest

#===================================== targets end ===========================================#

#=================================== variables start =========================================#

# project name - 项目名称
project_name := "echo-hub"

# project root directory - 项目根目录
root := justfile_directory()

# binary path - go build 输出的二进制文件路径
bin := join(root, "target")


#=================================== variables end =========================================#