
run:
	@echo ""
	@echo "-----------------------------------------"
	@echo "make ts       | 生成前端 TypeScript 接口文件"
	@echo "make swagger  | 生成 swagger json 文件"
	@echo "make service  | 生成服务初始化文件"
	@echo "-----------------------------------------"
	@echo ""

ts:
	@go build .
	@goctl api plugin -p ./goctl-ap="ts" --api example/define/api/platform/client.api  --dir example/ts/api

swagger:
	@go build .
	@goctl api plugin -p ./goctl-ap="swagger -f swagger.json" --api example/define/api/platform/api.api --dir example/swagger

service:
	@go run . service -a example/service/app -d example/service/define
