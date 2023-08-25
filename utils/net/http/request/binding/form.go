// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

const defaultMemory = 32 << 20

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct{}

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	if err := Decode(obj, req.Form); err != nil {
		return err
	}
	return validate(obj)
}

func (formBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseMultipartForm(defaultMemory); err != nil {
		if err != http.ErrNotMultipart {
			return err
		}
	}
	args := Args{formSource(ctx.Request.Form), paramSource(ctx.Params)}
	if err := mapForm(obj, args); err != nil {
		return err
	}
	return validate(obj)
}

func (formBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	if err := mapForm(obj, (*argsSource)(req.PostArgs())); err != nil {
		return err
	}
	return validate(obj)
}

func (formBinding) FiberBind(ctx *fiber.Ctx, obj interface{}) error {
	if err := mapForm(obj, (*argsSource)(ctx.Request().PostArgs())); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}
	if err := Decode(obj, req.PostForm); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseForm(); err != nil {
		return err
	}

	args := Args{formSource(ctx.Request.Form), paramSource(ctx.Params)}
	if err := mapForm(obj, args); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	if err := mapForm(obj, (*argsSource)(req.PostArgs())); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) FiberBind(ctx *fiber.Ctx, obj interface{}) error {
	if err := mapForm(obj, (*argsSource)(ctx.Request().PostArgs())); err != nil {
		return err
	}
	return validate(obj)
}

func (formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (formMultipartBinding) Bind(req *http.Request, obj interface{}) error {
	if err := req.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := mappingByPtr(obj, (*multipartRequest)(req), tag); err != nil {
		return err
	}

	return validate(obj)
}

func (formMultipartBinding) GinBind(ctx *gin.Context, obj interface{}) error {
	if err := ctx.Request.ParseMultipartForm(defaultMemory); err != nil {
		return err
	}
	if err := mappingByPtr(obj, (*multipartRequest)(ctx.Request), tag); err != nil {
		return err
	}

	return validate(obj)
}

func (formMultipartBinding) FasthttpBind(req *fasthttp.Request, obj interface{}) error {
	if err := mappingByPtr(obj, (*multipartFasthttpRequest)(req), tag); err != nil {
		return err
	}

	return validate(obj)
}

func (formMultipartBinding) FiberBind(ctx *fiber.Ctx, obj interface{}) error {
	if err := mappingByPtr(obj, (*multipartFasthttpRequest)(ctx.Request()), tag); err != nil {
		return err
	}

	return validate(obj)
}
