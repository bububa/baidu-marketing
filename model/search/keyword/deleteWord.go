package keyword

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// DeleteWordRequest 删除关键词 API Request
type DeleteWordRequest struct {
	// KeywordIds 关键词ID集合
	// 集合长度限制：[1, 10000]
	// 建议分批多次请求
	KeywordIds []uint64 `json:"keywordIds,omitempty"`
}

func (r DeleteWordRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "KeywordService/deleteWord")
}

// DeleteWordResponse 删除关键词 API Response
type DeleteWordResponse struct {
	Data []Keyword `json:"data,omitempty"`
}
