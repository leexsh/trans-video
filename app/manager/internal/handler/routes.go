package handler

import (
	"net/http"
	_ "net/http/pprof"

	"dist-encoder/app/manager/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)


func NilHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("nil function"))
	}
}
