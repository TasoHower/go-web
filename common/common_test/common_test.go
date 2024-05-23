package common_test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"
)

type Resp struct {
	Data WebDiffResp `json:"data"`
}
type WebDiffResp struct {
	LocalExist  bool     `json:"local_exist"`
	RemoteExist bool     `json:"remote_exist"`
	Local       []string `json:"local"`
	Remote      []string `json:"remote"`
}

type JSONData struct {
	InscriptionID string `json:"inscription_id"`
	Op            string `json:"op"`
	Source        string `json:"source"`
	Tick          string `json:"tick"`
	Valid         string `json:"valid"`
}

func TestWeb(t *testing.T) {

	block := 832034
	urlBase := "https://validator.odinbtc.io/api/test/ins/diff?block=%d"

	method := "GET"
	url := fmt.Sprintf(urlBase, block)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	fmt.Println(err)
	req.Header.Add("Accept", "application/json")

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var resp Resp

	err = json.Unmarshal(body, &resp)

	for _, s := range resp.Data.Local {
		var tmp JSONData
		_ = json.Unmarshal([]byte(s), &tmp)
		fmt.Printf("'%s',\n", tmp.InscriptionID)
	}
}
