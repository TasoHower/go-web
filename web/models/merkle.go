package models

type (
	BuildMerkleRequest struct {
		Block uint     `form:"block"`
		Ins   []string `form:"ins"`
		Trx   []string `form:"trx"`
	}

	BuildMerkleResponse struct{}
)

type (
	GetMerkleFileRequest struct {
		Block  uint `form:"block"`
		Remote bool `form:"remote"`
	}

	GetMerkleFileResp struct {
		Path string `json:"path"`
	}
)

type (
	GetLastPushRequest struct {
	}

	GetLastPushResponse struct {
		LocalLastPush  uint `json:"local_last_push"`
		RemoteLastPush uint `json:"remote_last_push"`
	}
)
