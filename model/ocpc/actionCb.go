package ocpc

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"

	"github.com/bububa/baidu-marketing/enum"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// ActionCbRequest 转化追踪回调 请求
type ActionCbRequest struct {
	// AKey 双方加密密钥，业务端创建转化时自动生成，同一广告主 akey 唯一
	// 监测 URL 签名：替换通配符后的完整监测 URL（不包含&sign=）+akey 进行标准 32 位 md5，生成签名值，在监测 URL 后添加&sign=签名值；Callback_URL签名：替换通配符后的完整回调URL（不包含&sign=）+ akey 进行标准32位md5，生成签名值，在回调URL后添加&sign=签名值，在百度投放后台获取。
	Akey string `json:"akey,omitempty"`
	// AType 必填，标识用户的转化事件
	// 1. 如您采用【方式一：使用百度提供的转化回调地址】，回调地址已给出a_type={{ATYPE}}，替换{{ATYPE}}即可；
	// 2. 如您采用【方式二：自主拼接转化回调地址】，需要您自行拼接a_type参数
	// orders：付费
	// retain_1day：次日留存
	// user_defined：客户自定义
	// ec_buy：商品下单成功
	// deep_page_access：关键页面浏览
	// credit_granting: 授信
	// deeplink：应用调起
	// retain_2day：2日留存
	// retain_3day：3日留存
	// retain_4day：4日留存
	// retain_5day：5日留存
	// retain_6day：6日留存
	// retain_7day：7日留存
	// retain_14day：14日留存
	// pay_to_read：付费阅读
	// enter_bookstore_read：进入书城阅读
	// add_to_desktop：添加至桌面
	// log_in：登录
	// order_submit_success：订单提交成功
	// pay_to_watch：付费观剧
	// key_action：关键行为
	// derived_event：衍生事件
	AType enum.AType `json:"a_type,omitempty"`
	// AValue     必填，标识转化事件的金额信息
	// 1. 如您采用【方式一：使用百度提供的转化回调地址】，回调地址已给出a_type={{AVALUE}}，替换{{AVALUE}}即可；
	// 2. 如您采用【方式二：自主拼接转化回调地址】，需要您自行拼接a_value参数
	// 转化类型为“orders：付费”时，此字段定义为“付费金额-单位(分)”。例：转化金额为12.3元，a_value=1230
	// ※注意：①单位是（分）；②无需回传转化金额时，a_value数值填写为“0”即可，即a_value=0
	AValue  int64        `json:"a_value,omitempty"`
	ActType enum.ActType `json:"act_type,omitempty"`
	// ATime 用户转化事件发生的具体时间, unix时间戳（精确到秒）
	ATime int64 `json:"a_time,omitempty"`
	// ExtInfo 必填，广告编码信息
	// 1. 如您采用【方式一：使用百度提供的转化回调地址】，回调地址已给出ext_info=xxxx，直接回调即可；
	// 2. 如您采用【方式二：自主拼接转化回调地址】，需要您自行拼接ext_info参数
	// ext_info部分场景较长，建议至少预留3kb存储空间，回调需完整，使用时避免字符串截断
	// callType=v2 用户使用
	ExtInfo string `json:"ext_info,omitempty"`
	// CallbackUrl 效果数据回传URL
	// 由百度 server 端调用监测 URL 时直接传输给广告主，产生转化时广告主替换转化相关参数（a_type、a_value）后生成签名进行回调，callback_url 中的 s,o,ext_info 字段由 baidu 公司 server 负责处理。callType=v2 的用户可以不拼接 s,o参数
	CallbackUrl string `json:"callback_url,omitempty"`
	// CbJoinType 必填，转化数据归因类型，用于辅助百度优化投放效果
	// 例如，使用oaid匹配用户信息成功后，回传百度转化时，回调url增加&join_type=oaid标识该条转化数据使用oaid归因成功
	// · IMEI（通过IMEI_MD5归因成功）
	// · OAID（通过OAID或OAID_MD5归因成功）
	// · ANDROID_ID（通过ANDROID_ID或ANDROID_ID_MD5归因成功）
	// · IP（通过IP和系统信息（UA、OS_VERSION、MODEL）归因成功）
	// · IDFA（通过IDFA归因成功）
	// · CAID（通过CAID归因成功）
	// · MAC（通过MAC或MAC_MD5归因成功）
	// · PAID（通过PAID归因成功，灰度使用，可咨询您的百度SEM顾问如何申请）
	CbJoinType string `json:"cb_join_type,omitempty"`
	// CbIdfa IOS设备标识：原值
	CbIdfa string `json:"cb_idfa,omitempty"`
	// CbImei Android设备标识：原值32
	CbImei string `json:"cb_imei,omitempty"`
	// CbImeiMd5 Android 设备标识
	// Android 设备标识：编码方式：标准32 位 md5 当无 IMEI 传递时为空（此处 imei 使用从客户端获取的原始值进行 md5 加密，不需要做大小写转换）例：imei[10bc955ac2a675d3]imei_md5[f703b39228c8c5cf8069051d86a20747]
	CbImeiMd5 string `json:"cb_imei_md5,omitempty"`
	// CbOaid Android设备标识：原值
	CbOaid string `json:"cb_oaid,omitempty"`
	// CbAndroidId Android设备标识：原值32位小写MD5
	CbAndroidId string `json:"cb_android_id,omitempty"`
	// CbAndroidMd5 Android 设备标识; Android 设备标识：标准 32 位 md5当无 android id 时传递时为 null
	CbAndroidIdMd5 string `json:"cb_android_id_md5,omitempty"`
	// CbEventTime 发生转化时间（精确到毫秒）
	CbEventTime int64 `json:"cb_event_time,omitempty"`
	// OrderType 订单类型
	OrderType string `json:"order_type,omitempty"`
	// PayAmount 付费金额
	PayAmount int64 `json:"pay_amount,omitempty"`
	// OrderId 订单 id
	OrderId string `json:"order_id,omitempty"`
	// ActionType 行为类型
	// 可传多组，逗号分隔，和 action_value 一一对应
	// 1：广告观看（IPU）；
	// 2：ARPU；
	// 3：eCPM 均值；
	// 4：通过关卡数；
	// 5：提现次数；
	// 6：使用时长；
	// 7：eCPM 最值；
	// 8：24h ARPU 门槛；
	// 9：预估 LTV 7；
	// 10：预估 LTV all；
	// 0：其他；
	ActionType []int `json:"action_type,omitempty"`
	// ActionValue 行为数值
	// 可传多组，逗号分隔，和 action_type 一一对应
	// action_type 对应的指标门槛值，如：ipu=5则传5
	// ※注意：eCPM和ARPU单位为分
	ActionValue []int64 `json:"action_value,omitempty"`
	// ActionTs 事件实际发生的时间, 时间戳（秒级别）
	ActionTs int64 `json:"action_ts,omitempty"`
	// IsGaConvert 是否关键行为转化
	// 是否把当前的回传作为关键行为转化
	// 1：关键行为，0：衍生行为
	IsGaConvert *int `json:"is_ga_convert,omitempty"`
	// Depth 行为深度
	//     关键行为定义的深度；模型学习其排序结果，不学习其绝对值，辅助样本区分。
	// 可取值为 1-100 之间整数，如 30、50
	Depth int `json:"depth,omitempty"`
}

func (r ActionCbRequest) callbackUrl() string {
	if r.CallbackUrl != "" {
		link := strings.ReplaceAll(r.CallbackUrl, "{{ATYPE}}", string(r.AType))
		link = strings.ReplaceAll(link, "{{AVALUE}}", strconv.FormatInt(r.AValue, 10))
		return link
	}
	values := util.GetUrlValues()
	defer util.PutUrlValues(values)
	values.Set("a_type", string(r.AType))
	values.Set("a_value", strconv.FormatInt(r.AValue, 10))
	values.Set("act_type", strconv.Itoa(int(r.ActType)))
	values.Set("ext_info", r.ExtInfo)
	if r.ATime > 0 {
		values.Set("a_time", strconv.FormatInt(r.ATime, 10))
	}
	if r.CbJoinType != "" {
		values.Set("cb_join_type", r.CbJoinType)
	}
	if r.CbIdfa != "" {
		values.Set("cb_idfa", r.CbIdfa)
	}
	if r.CbImei != "" {
		values.Set("cb_imei", r.CbImei)
	}
	if r.CbImeiMd5 != "" {
		values.Set("cb_imei_md5", r.CbImeiMd5)
	}
	if r.CbOaid != "" {
		values.Set("cb_oaid", r.CbOaid)
	}
	if r.CbAndroidId != "" {
		values.Set("cb_android_id", r.CbAndroidId)
	}
	if r.CbAndroidIdMd5 != "" {
		values.Set("cb_android_id_md5", r.CbAndroidIdMd5)
	}
	if r.CbEventTime > 0 {
		values.Set("cb_event_time", strconv.FormatInt(r.CbEventTime, 10))
	}
	if r.OrderType != "" {
		values.Set("order_type", r.OrderType)
	}
	if r.PayAmount > 0 {
		values.Set("pay_amount", strconv.FormatInt(r.PayAmount, 10))
	}
	if r.OrderId != "" {
		values.Set("order_id", r.OrderId)
	}
	if len(r.ActionType) > 0 {
		vals := make([]string, 0, len(r.ActionType))
		for _, v := range r.ActionType {
			vals = append(vals, strconv.Itoa(v))
		}
		values.Set("action_type", strings.Join(vals, ","))
	}
	if len(r.ActionValue) > 0 {
		vals := make([]string, 0, len(r.ActionValue))
		for _, v := range r.ActionValue {
			vals = append(vals, strconv.FormatInt(v, 10))
		}
		values.Set("action_value", strings.Join(vals, ","))
	}
	if r.ActionTs > 0 {
		values.Set("action_ts", strconv.FormatInt(r.ActionTs, 10))
	}
	if r.IsGaConvert != nil {
		values.Set("is_ga_convert", strconv.Itoa(*r.IsGaConvert))
	}
	if r.Depth > 0 {
		values.Set("depth", strconv.Itoa(r.Depth))
	}
	return util.StringsJoin(model.ACTIONCB_URL, "?", values.Encode())
}

func (r ActionCbRequest) Sign() string {
	signUrl := util.StringsJoin(r.callbackUrl(), r.Akey)
	h := md5.New()
	h.Write([]byte(signUrl))
	return hex.EncodeToString(h.Sum(nil))
}

func (r ActionCbRequest) Url() string {
	return util.StringsJoin(r.callbackUrl(), "&sign=", r.Sign())
}
