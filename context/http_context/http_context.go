package http_context

import (
	"context"
	contexti "github.com/hopeio/lemon/context"
	contexti2 "github.com/hopeio/lemon/utils/context"
	"go.opencensus.io/trace"
	"net/http"
)

type Context = contexti.Context[http.Request]

func ContextFromRequest(r *http.Request, tracing bool) (*Context, *trace.Span) {
	ctxi, span := contexti2.CtxWithRequest(r, tracing)
	return &Context{Authorization: &contexti.Authorization{}, RequestContext: ctxi}, span
}

func ContextFromContext(ctx context.Context) *Context {
	return contexti.CtxFromContext[http.Request](ctx)
}
