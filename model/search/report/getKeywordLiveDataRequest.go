package report

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetKeywordLiveDataRequest struct {
	DataType     int     `json:"dataType,omitempty"`     // 账户/计划实时数据查询
	KeywordIds   []int64 `json:"keywordIds,omitempty"`   // 关键词id，该集合为请求关键词实时数据的集合; 数组内关键词个数上限为200
	Device       int     `json:"device,omitempty"`       // 投放设备
	StartTime    string  `json:"startTime,omitempty"`    // 起始时间; 选填， 默认值：当前请求时间之前30分钟时间点; 说明1：请求时间即为当前起始时间; 说明2：起始时间仅支持当前时间之前的30分钟内数据; 说明3：时间格式：2021-04-22 14:30:00; 举例：当前时间为2020-08-01 14:00:00，则startTime不能早于2020-08-01 13:31:00， 不能晚于2020-08-01 14:00:00；默认值为当前时间30分钟之前的时间
	EndTime      string  `json:"endTime,omitempty"`      // 结束时间; 同起始时间，且不能早于startTime
	RegionTarget []int   `json:"regionTarget,omitempty"` // 地域; 选填，默认全部区域，取值范围参考地域代码，地域代码个数上限为10
}

func (r GetKeywordLiveDataRequest) Url() string {
	return fmt.Sprintf("%sLiveReportService/getKeywordLiveData", model.BASE_URL_SMS)
}
