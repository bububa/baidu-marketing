package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

// GetKeywordLiveDataRequest 关键词实时数据 API Request
type GetKeywordLiveDataRequest struct {
	// DateType 账户/计划实时数据查询
	DataType int `json:"dataType,omitempty"`
	// KeywordIds 关键词id，该集合为请求关键词实时数据的集合; 数组内关键词个数上限为200
	KeywordIds []int64 `json:"keywordIds,omitempty"`
	// Device 投放设备
	Device int `json:"device,omitempty"`
	// StartTime 起始时间; 选填， 默认值：当前请求时间之前30分钟时间点; 说明1：请求时间即为当前起始时间; 说明2：起始时间仅支持当前时间之前的30分钟内数据; 说明3：时间格式：2021-04-22 14:30:00; 举例：当前时间为2020-08-01 14:00:00，则startTime不能早于2020-08-01 13:31:00， 不能晚于2020-08-01 14:00:00；默认值为当前时间30分钟之前的时间
	StartTime string `json:"startTime,omitempty"`
	// EndTime 结束时间; 同起始时间，且不能早于startTime
	EndTime string `json:"endTime,omitempty"`
	// RegionTarget 地域; 选填，默认全部区域，取值范围参考地域代码，地域代码个数上限为10
	RegionTarget []int `json:"regionTarget,omitempty"`
}

func (r GetKeywordLiveDataRequest) Url() string {
	return fmt.Sprintf("%sLiveReportService/getKeywordLiveData", model.BASE_URL_SMS)
}
