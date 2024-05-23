package context

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"io"
	"time"
	"web/logger"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

const (
	RequestId   = "request_id"
	RequestTime = "request_time"
)

func GetGinContextWithRequestId() *gin.Context {
	ctx := gin.Context{}

	value := uuid.Must(uuid.NewV4(), nil).String()
	m := md5.New()
	m.Write([]byte(value))
	requestId := hex.EncodeToString(m.Sum(nil))
	SetRequestID(&ctx, requestId)
	return &ctx
}

func SetContextData(ctx *gin.Context, key string, value any) {
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]any)
	}
	//ctx.Keys[key] = value
	ctx.Set(key, value)
}

func GetContextData(ctx *gin.Context, key string) any {
	if ctx.Keys == nil {
		return nil
	}
	//value, ok := ctx.Keys[key]
	value, ok := ctx.Get(key)
	if !ok {
		return nil
	}
	return value
}

func SetRequestID(ctx *gin.Context, requestId string) {
	SetContextData(ctx, RequestId, requestId)
}

func GetRequestID(ctx *gin.Context) string {
	v := GetContextData(ctx, RequestId)
	if v == nil {
		return ""
	}
	return v.(string)
}

func SetRequestTIme(ctx *gin.Context) {
	SetContextData(ctx, RequestTime, time.Now().UnixMilli())
}

func GetRequestTIme(ctx *gin.Context) int64 {
	v := GetContextData(ctx, RequestTime)
	if v == nil {
		return -1
	}
	return v.(int64)
}

func GetRequestId() (requestId string) {
	value := uuid.Must(uuid.NewV4(), nil).String()
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}

func AddRequestId(ctx *gin.Context) {
	requestId := GetRequestId()
	SetRequestID(ctx, requestId)
	body, err := ctx.GetRawData()
	if err != nil {
		logger.Error(err.Error())
	}
	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	logger.Infof("\033[0;32m[SMART REQUEST IN]\033[0;0m [RequestURI: %s] [Request ID: %s] [RequestBody: %v]", requestId, ctx.Request.Header, string(body))
}
