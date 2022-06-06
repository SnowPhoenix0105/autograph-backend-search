package rpc

import (
	"autograph-backend-search/server/common"
	"autograph-backend-search/utils"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetFileInfoResp struct {
	URL  string `json:"url"`
	Name string `json:"name"`
	Type string `json:"type"`
}

func ControllerGetFileInfo(ctx context.Context, id uint) (*GetFileInfoResp, error) {
	url := fmt.Sprintf("%s/fileinfo?id=%d", globalConfig.ControllerBaseUrl, id)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, utils.WrapError(err, "create req fail")
	}

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, utils.WrapError(err, "rpc call fail")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, utils.WrapErrorf(ErrHttpStatus, "status[code=%d, msg=%#v]", resp.StatusCode, resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, utils.WrapError(err, "read body fail")
	}

	var ret GetFileInfoResp
	base := common.BaseResp{
		Data: &ret,
	}
	err = json.Unmarshal(respBody, &base)
	if err != nil {
		return nil, utils.WrapErrorf(err, "unmarshal(%#v) fail", string(respBody))
	}

	return &ret, nil
}
