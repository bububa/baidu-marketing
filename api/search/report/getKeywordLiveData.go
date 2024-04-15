package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/report"
)

// GetKeywordLiveData 关键词实时数据
// 通过此接口获取关键词实时数据报告信息。
func GetKeywordLiveData(clt *core.SDKClient, auth *model.RequestHeader, reqBody *report.GetKeywordLiveDataRequest) (*model.ResponseHeader, []report.KeywordLiveData, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp report.GetKeywordLiveDataResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
