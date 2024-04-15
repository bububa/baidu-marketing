package keyword

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// UpdateWordRequest 更新关键词 API Request
type UpdateWordRequest struct {
	// KeywordTypes 新增关键词对象数组
	// 参见KeywordType说明 单次请求不超过10000 建议分批多次请求
	KeywordTypes []Keyword `json:"keywordTypes,omitempty"`
}

func (r UpdateWordRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_SMS, "KeywordService/updateWord")
}

// UpdateWordResponse 更新关键词 API Response
type UpdateWordResponse struct {
	Data []Keyword `json:"data,omitempty"`
}
