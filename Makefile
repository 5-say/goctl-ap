
run:test

i:
	@goctl api ts -api example/define/api/api.api -dir example/src/ts

init:
	sh shell/install-vscode-extension.sh
	cd dev && git clone git@github.com:5-say/web-client-demo.git client && code client
	cd dev && git clone git@github.com:5-say/web-manage-demo.git client && code manage

test:
	@go build .
	@goctl api plugin -p ./goctl-ts -api example/define/api/api.api -dir example/src/api