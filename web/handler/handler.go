package handler

import (
	"encoding/json"
	"net/http"
	"time"
	"web/common"
	"web/context"
	"web/logger"

	"github.com/gin-gonic/gin"
	"github.com/pquerna/ffjson/ffjson"
)

const (
	MethodGET    = "GET"
	MethodPOST   = "POST"
	MethodPUT    = "PUT"
	MethodDELETE = "DELETE"
	MethodHEAD   = "HEAD"
)

type Response struct {
	Data      any    `json:"data"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	RequestID string `json:"request_id"`
}

func (r *Response) InitCode(code int) {
	r.Code = code
	r.Msg = common.GetMsg(code)
}

func (r *Response) ToStruct(obj any) error {
	b, _ := ffjson.Marshal(r.Data)
	return ffjson.Unmarshal(b, obj)
}

func render(c *gin.Context, v any, err error) {
	requestId := context.GetRequestID(c)
	resp := Response{Data: v, Code: common.SUCCESS, Msg: common.GetMsg(common.SUCCESS)}
	if err != nil {
		if tem, ok := err.(*common.Error); ok {
			resp = Response{Data: v, Code: tem.Code, Msg: tem.Error(), RequestID: requestId}
		} else {
			resp = Response{Data: v, Code: common.Unknown, Msg: common.GetMsg(common.Unknown), RequestID: requestId}
		}
	}

	start := time.UnixMilli(context.GetRequestTIme(c))
	res, _ := json.Marshal(resp)

	logger.Infof("\033[0;32m [SMART REQUEST OUT]\033[0m [Request ID: %s] [Processing time:%6d ms] [res: %s]", requestId, time.Since(start).Milliseconds(), string(res))

	c.JSON(http.StatusOK, resp)

}

type (
	TRPathParamHandlerFunc[T any, R any] func(ctx *gin.Context, t *T) (R, error)
)

func TRPathParamHandler[T any, R any](
	handler TRPathParamHandlerFunc[T, R],
) gin.HandlerFunc {
	return handlerWithContext[T, R](handler)
}

func handlerWithContext[T any, R any](
	handlerFunc any,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := new(T)
		if err := c.ShouldBind(&t); err != nil {
			render(c, nil, common.New(common.ParamsErr))
			return
		}
		switch handler := handlerFunc.(type) {
		case TRPathParamHandlerFunc[T, R]:
			v, e := handler(c, t)
			render(c, v, e)
		}
	}
}

// 原样返回，不包装
func TRPathParamNudeHandler[T any, R any](
	handler TRPathParamHandlerFunc[T, R],
) gin.HandlerFunc {
	return handlerWithNudeContext[T, R](handler)
}

func handlerWithNudeContext[T any, R any](
	handlerFunc any,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := new(T)
		if err := c.ShouldBind(&t); err != nil {
			renderNude(c, nil, common.New(common.ParamsErr))
			return
		}
		switch handler := handlerFunc.(type) {
		case TRPathParamHandlerFunc[T, R]:
			v, e := handler(c, t)
			renderNude(c, v, e)
		}

	}
}

func renderNude(c *gin.Context, v any, err error) {
	resp := v
	c.JSON(http.StatusOK, resp)
}
