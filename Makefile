
help:
	@echo ""
	@echo "------------------------------------------"
	@echo "make i   | make init   : 初始化工作区"
	@echo "------------------------------------------"
	@echo ""

run:
	make help

init:
	sh shell/install-vscode-extension.sh
	cd dev && git clone git@github.com:5-say/web-client-demo.git client && code client
	cd dev && git clone git@github.com:5-say/web-manage-demo.git client && code manage
