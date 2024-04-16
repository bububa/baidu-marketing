package keyword

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// AddWordRequest 添加关键词 API Request
type AddWordRequest struct {
	// KeywordTypes 新增关键词对象数组
	// 集合长度限制：[1, 10000]
	// 建议分批多次请求
	KeywordTypes []Keyword `json:"keywordTypes,omitempty"`
}

func (r AddWordRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "KeywordService/addWord")
}

// AddWordResponse 添加关键词 API Response
type AddWordResponse struct {
	Data []Keyword `json:"data,omitempty"`
}
