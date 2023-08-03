
run:test-service

test-ts:
	@go build .
	@goctl api plugin -p ./goctl-ap="ts" --api example/define/api/platform/client.api  --dir example/ts/api

test-swagger:
	@go build .
	@goctl api plugin -p ./goctl-ap="swagger -f swagger.json" --api example/define/api/platform/api.api --dir example/swagger

test-service:
	@go run . service -a example/service/app -d example/service/define
