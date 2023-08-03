package middleware

import (
	"common/response"
	"net/http"

	"github.com/5-say/go-tools/tools/ip"
	"github.com/5-say/go-tools/tools/t"
	"github.com/5-say/zero-auth/public/jwtx"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type AuthAdminMiddleware struct {
	jwtxClient jwtx.JwtxClient
	group      string
}

func NewAuthAdminMiddleware(jwtxClient jwtx.JwtxClient, group string) *AuthAdminMiddleware {
	return &AuthAdminMiddleware{
		jwtxClient: jwtxClient,
		group:      group,
	}
}

func (m *AuthAdminMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// 调用 jwtx-rpc 中间件
		rpcResponse, rpcError := m.jwtxClient.CheckToken(r.Context(), &jwtx.CheckToken_Request{
			Group:        m.group,
			RequestIP:    ip.GetRequestIP(r),
			RequestToken: r.Header.Get("Authorization"),
		})

		// token 校验失败
		if rpcError != nil {

			err := t.RPCErrorParse(rpcError)

			// 内部错误信息日志
			logx.Debug(err.PrivateMessage)

			// 加工错误对象
			appErr := response.Unauthorized().Message(err.Error())

			// 为响应对象写入错误信息
			httpx.ErrorCtx(r.Context(), w, appErr)

		} else {

			// 响应头填充刷新后的 token
			w.Header().Add("NEW-TOKEN", rpcResponse.NewToken)

			// 附加上下文
			newCtx := jwtx.WithValue(r.Context(), rpcResponse)

			// Passthrough to next handler if need
			next(w, r.WithContext(newCtx))
		}

	}
}
