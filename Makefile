
run:api-plugin-ts

help:
	@echo ""
	@echo "------------------------------------------"
	@echo "make i   | make init   : 初始化工作区"
	@echo "------------------------------------------"
	@echo ""

init:
	sh shell/install-vscode-extension.sh
	cd dev && git clone git@github.com:5-say/web-client-demo.git client && code client
	cd dev && git clone git@github.com:5-say/web-manage-demo.git client && code manage

ts:api-ts
api-ts:
	go run . ts --api ../../define/api/api.api --dir ts

api-plugin-ts:
	@go build .
	@goctl api plugin -p ./goctl-ts -api ../../define/api/api.api -dir ts