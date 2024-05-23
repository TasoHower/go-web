package ping

import (
	"web/logger"
	"web/web/models"

	"github.com/gin-gonic/gin"
)

func GetPingInfo(c *gin.Context, req *models.PingReq) (any, error) {
	var ret models.PingResp
	ret.Ping = `pong`
	logger.Infof("ping running. [ret:%s]", ret)

	return ret, nil
}
