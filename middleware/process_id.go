package middleware

import (
	"context"
	"github.com/riyanathariq/jwt-with-rbac/common"
	"github.com/riyanathariq/jwt-with-rbac/consts"
	"net/http"
	"time"
)

func ProcessIdInjector(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		regid := r.Header.Get(`X-Request-Id`)
		if regid == "" {
			regid = common.GenerateReferenceID("PID-")
			r.Header.Set(`X-Request-Id`, regid)
		}
		w.Header().Set(`X-Request-Id`, regid)

		ctx := context.WithValue(r.Context(), consts.ContextRequestID, regid)
		ctx = context.WithValue(ctx, consts.ContextStartTime, time.Now())
		ctx = context.WithValue(ctx, consts.ContextRequestIp, common.IPFromRequest(r))
		ctx = context.WithValue(ctx, consts.ContextRequestPath, r.URL.Path)
		ctx = context.WithValue(ctx, consts.ContextRequestMethod, r.Method)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
