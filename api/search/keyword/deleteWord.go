package keyword

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/search/keyword"
)

// DeleteWord 删除关键词
func DeleteWord(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, keywordIds ...uint64) (*model.ResponseHeader, []keyword.Keyword, error) {
	req := &model.Request{
		Header: auth,
		Body:   &keyword.DeleteWordRequest{KeywordIds: keywordIds},
	}
	var resp keyword.DeleteWordResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
