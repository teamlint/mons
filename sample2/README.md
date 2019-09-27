# mons 微服务框架生成工具
## 操作指南 
1. 在 services 目录创建服务目录, 如: user

2. 在服务目录下创建 api 子目录, 并创建服务 proto 文件

3. 生成基础 grpc 代码
```shell
cd user
protoc --proto_path=./api --proto_path=../../templates/api --go_out=plugins=grpc:./api api/*.proto
```

4. 生成微服务框架
```shell
protoc --proto_path=./api --proto_path=../../templates/api --gotemplate_out=template_dir=../../templates,debug=true,all=false:. api/*.proto
```

5. 编写业务服务代码
打开 `app/service/service.go`, 编写业务服务代码

6. 运行微服务服务端
```shell
go run cmd/server/main.go
```

7. 运行微服务客户端
```shell
go run cmd/client/main.go
```

8. 维护
框架代码生成后,已生成Taskfile文件, 后续可以使用 task 命令维护管理
- 重新生成 proto 代码
```shell
task pb
```
- 重新生成服务框架代码
```shell
task gen
```

## TODO
- protoc grpc 生成代码简化, 只要模型和接口
- 编写grpc客户端, 修订nats及http使用grpc编解码
