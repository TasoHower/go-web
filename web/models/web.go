package models

type (
	WebCheckedReq struct {
	}

	WebCheckedResp struct {
		List []WebCheckedResult `json:"list"`
	}

	WebCheckedResult struct {
		Number int    `json:"number"`
		Hash   string `json:"hash"`
		Status int    `json:"status"`
	}
)

type (
	WebListReq struct {
		Block uint `form:"block"`
	}

	WebListResp struct {
		List []WebListResult `json:"list"`
	}

	WebListResult struct {
		Number int    `json:"number"`
		Hash   string `json:"hash"`
		Status int    `json:"status"`
	}
)

type (
	WebDiffReq struct {
		Block uint `form:"block"`
	}

	WebDiffResp struct {
		LocalExist  bool `json:"local_exist"`
		RemoteExist bool `json:"remote_exist"`

		Local  []string `json:"local"`
		Remote []string `json:"remote"`
	}

	WebDiffResult struct {
		Number int    `json:"number"`
		Hash   string `json:"hash"`
		Status int    `json:"status"`
	}
)
