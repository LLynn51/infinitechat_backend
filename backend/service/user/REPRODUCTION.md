# User 模块复现演进记录

本文档是 `backend/service/user` 的演进台账。每个阶段必须在同一个 Git 提交中记录目标、范围、验证结果、已知问题和下一步。

## 基线

- 目标仓库：`D:\LLynn_Projects\My_Go\infinitechat`
- 参考仓库：`D:\LLynn_Projects\infinitechatgo`
- 参考提交：`52c4009`
- 技术路线：Go、go-zero、HTTP API、gRPC
- 复现原则：保留主要领域结构，但不复制参考项目中的密钥、阻塞缺陷或不必要耦合。

## 状态

| 步骤 | 状态 | 内容 | 提交主题 |
|---|---|---|---|
| USR-001 | 已完成 | 建立可启动、可验证的 API/RPC 最小骨架 | `[USR-001] chore(user): scaffold api and rpc services` |
| USR-002 | 计划中 | 建立本地依赖底座与环境配置 | 待定 |

## USR-001：建立最小骨架

### 目标

- 建立 user-api HTTP 进程。
- 建立 user-rpc gRPC 进程。
- 两个进程均不依赖 MySQL、Redis、Kafka、etcd 或 MinIO 即可启动。
- 提供可自动验证的健康接口，为后续部署和监控留下稳定入口。

### 实现范围

#### HTTP API

- `api/user.api` 定义 `GET /healthz`。
- 使用 `goctl api go` 生成入口、配置、路由、Handler、Logic、ServiceContext 和类型。
- 健康响应读取当前服务配置，返回：

  ```json
  {"service":"user-api","status":"ok"}
  ```

#### gRPC

- `rpc/user.proto` 定义 `UserRpc.Health`。
- 使用 `goctl rpc protoc` 生成 protobuf、gRPC Server、Client 和 Logic 脚手架。
- 健康响应返回服务名 `user.rpc` 和状态 `ok`。
- 开发和测试模式启用 gRPC reflection。

#### 配置

- user-api 默认监听 `0.0.0.0:8104`。
- user-rpc 默认监听 `0.0.0.0:8105`。
- 当前配置只包含启动和控制台日志所需字段，不包含任何凭据。

#### 测试

- API 健康 Logic 单元测试。
- RPC 健康 Logic 单元测试。

### 生成命令

在对应目录执行：

```powershell
goctl api go --api user.api --dir . --style go_zero
goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=. --style go_zero --module backend
```

生成器版本：`goctl 1.10.2`。

### 验证结果

- `go test ./service/user/api/... ./service/user/rpc/...`：通过。
- `go test ./...`：通过。
- user-api 启动成功，访问 `GET http://127.0.0.1:8104/healthz` 返回 HTTP 200 和预期 JSON。
- user-rpc 启动成功并监听 `0.0.0.0:8105`。
- `git diff --check`：通过。

### 本步明确不包含

- 用户注册、登录、验证码、Token 或资料接口。
- MySQL、Redis、Kafka、etcd、MinIO 和 SMTP 初始化。
- 数据库迁移及运行时数据。
- Gateway 和 Realtime 集成。
- 生产部署配置。

### 已知问题

- 当前 gRPC 依赖来自 go-zero 的完整 zrpc 依赖图，因此 `go mod tidy` 会引入 etcd 和 Kubernetes 客户端等传递依赖；这不是本步的运行时外部依赖。
- 已存在的 Model 目录尚未接入 API/RPC ServiceContext，本阶段只验证其仍可编译。

### 下一步

`USR-002` 建立可重复启动的本地依赖底座和环境配置，先明确 MySQL、Redis、etcd、Kafka、MinIO 的职责、版本、端口、健康检查与 Secret 注入方式，再连接业务代码。
