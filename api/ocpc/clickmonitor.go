package ocpc

import "github.com/bububa/baidu-marketing/util"

var defaultFields = map[string]string{
	"advertiser_id": "__USER_ID__",
	"cid":           "__IDEA_ID__",
	"campaign_id":   "__PLAN_ID__",
	"aid":           "__UNIT_ID__",
	"word_id":       "__WORD_ID__",
	"callback_url":  "__CALLBACK_URL__",
	"ext_info":      "__EXT_INFO__",
	"clickid":       "__CLICK_ID__",
	"idfa":          "__IDFA__",
	"imei":          "__IMEI__",
	"android_id":    "__ANDROIDID__",
	"oaid_md5":      "__OAID_MD5__",
	"oaid":          "__OAID__",
	"caid":          "__CAID__",
	"ts":            "__TS__",
	"os_type":       "__OS_TYPE__",
	"bd_vid":        "__BD_VID__",
}

func ClickMonitorUrl(baseUrl string, fields []string, version int) {
	values := util.GetUrlValues()
	defer util.PutUrlValues(values)
	for _, f := range fields {
		if version == 2 && f == "callback_ur" {
			continue
		}
		if v, ok := defaultFields[f]; ok {
			values.Set(f, v)
		}
	}
	if version == 2 {
		values.Set("callType", "v2")
	} else {
		values.Set("callType", "v1")
	}
	values.Set("sign", "__SIGN__")
}
