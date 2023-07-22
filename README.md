# goctl-ap

goctl api plugin

## 开发环境

```sh
docker dev create -o --name goctl-ap git@github.com:5-say/goctl-ap.git
```

## 使用方法

> 安装插件

```sh
go install github.com/5-say/goctl-ap@latest
```

> 生成前端 TypeScript 接口文件

```sh
goctl api plugin -p goctl-ap="ts" --api example/define/api/platform/client.api  --dir example/ts/api
```

> 生成 swagger json 文件

```sh
goctl api plugin -p goctl-ap="swagger -f swagger.json" --api example/define/api/platform/api.api --dir example/swagger
```
