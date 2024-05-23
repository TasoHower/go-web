package models

type (
	PingReq struct{}

	PingResp struct {
		Ping string `json:"ping"`
	}
)
