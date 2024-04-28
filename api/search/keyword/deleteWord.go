package keyword

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/search/keyword"
)

// DeleteWord 删除关键词
func DeleteWord(clt *core.SDKClient, auth *model.RequestHeader, keywordIds ...uint64) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body:   &keyword.DeleteWordRequest{KeywordIds: keywordIds},
	}
	var resp keyword.DeleteWordResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
