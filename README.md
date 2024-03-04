# h***-go

## 简介

#### 『**』 Server But Golang Ver.
## 编译和运行环境

* Go >= 1.18
* Protoc >= 3.21
* Protoc Gen Go >= 1.28
* Docker >= 20.10
* Docker Compose >= 1.29

## 快速启动

* 首次需要安装工具

```shell
make dev_tool
```

* 生成协议

```shell
make gen_natsrpc      # 生成natsrpc协议
make gen_proto        # 生成客户端协议
make gen_client_proto # 生成客户端协议代理(非必要 详见gate/client_proto/README.md)
```

* 构建

```shell
make build         # 构建服务器二进制文件
make docker_config # 复制配置模板等文件
make docker_build  # 构建镜像
```

* 启动

```shell
cd docker
# 启动前请先确保各服务器的配置文件正确(如docker/node/bin/application.toml)
docker-compose up -d # 启动服务器
```

#### 第三方组件

* mongodb
* nats-server
* redis

#### 服务器组件

* node 节点服务器 (仅单节点 有状态)
* dispatch 登录服务器 (可多节点 无状态)
* gate 网关服务器 (可多节点 有状态)
* multi 多功能服务器 (可多节点 有状态 尚不完善非必要启动)
* gs 游戏服务器 (可多节点 有状态)
* gm 游戏管理服务器 (仅单节点 无状态)

#### 其它

* 部分服务器组件以本地原生进程方式启动需要添加以下环境变量

```shell
GOLANG_PROTOBUF_REGISTRATION_CONFLICT=ignore
```

## 代码提交规范

#### 欢迎提交PR

* 提交前**必须**格式化你的代码，如运行`go fmt`
* 进行全局格式化时，请跳过`gdconf/game_data_config`目录，这是配置表数据，包含大量的`json`、`lua`、`txt`等文件
