package api

import (
	"context"
	"net/http"

	"app/{{ .serviceName }}/api/doc"
	"app/{{ .serviceName }}/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwaggerDocLogic struct {
	logx.Logger
	w      http.ResponseWriter
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSwaggerDocLogic(svcCtx *svc.ServiceContext, w http.ResponseWriter, r *http.Request) *SwaggerDocLogic {
	return &SwaggerDocLogic{
		Logger: logx.WithContext(r.Context()),
		w:      w,
		r:      r,
		ctx:    r.Context(),
		svcCtx: svcCtx,
	}
}

func (l *SwaggerDocLogic) SwaggerDoc() error {
	// 写入文档数据
	l.w.Write(doc.SwaggerJson)

	return nil
}
