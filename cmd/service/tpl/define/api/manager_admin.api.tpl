info(
	description: "管理端 - 管理员登录接口"
)

@server (
	group: manager/admin
	prefix: manager/admin
	auth: jwt
	middleware: AuthAdminMiddleware
)

service {{ .serviceName }} {
    @doc(
        summary: Demo
    )
    @handler demo
    post /demo(Manager_Admin_Demo_Request) returns (Manager_Admin_Demo_Response)
}

// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - Demo
// - 请求
type (
    Manager_Admin_Demo_Request {}
)
// - 响应
type (
    Manager_Admin_Demo_Response {}
)
