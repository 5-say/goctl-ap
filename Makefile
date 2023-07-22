
run:test-swagger

test-ts:
	@go build .
	@goctl api plugin -p ./goctl-ts="ts" --api example/define/api/platform/client.api  --dir example/ts/api

test-swagger:
	@go build .
	@goctl api plugin -p ./goctl-ts="swagger" --api example/define/api/platform/api.api --dir example/swagger
