package adgroup

// Ocpc oCPC设置对象
type Ocpc struct {
	// AppTransId 转化ID; 营销目标为应用推广Android的单元，仅能绑定应用操作系统为Android的转化追踪；营销目标为应用推广iOS的单元，仅能绑定应用操作系统为iOS的转化追踪。
	AppTransId uint64 `json:"appTransId,omitempty"`
	// TransFrom 接入方式; 1 - 应用API 2 - 基木鱼营销页; 4 - API激活; 5 - 网页JS布码; 7 - 线索API; 8 - 咨询工具授权; 13 - 应用SDK
	TransFrom int `json:"transFrom,omitempty"`
	// OcpcBid 目标转化出价; 取值范围：[0.3，999.99] 单位为元
	OcpcBid float64 `json:"ocpcBid,omitempty"`
	// LpUrl 推广URL; 必须以“http”或“https”开头
	LpUrl string `json:"lpUrl,omitempty"`
	// TransType 目标转化
	// 取值范围：枚举值，列表如下
	// 1 - 咨询按钮点击
	// 2 - 电话按钮点击
	// 3 - 表单提交成功
	// 4 - 激活
	// 5 - 表单按钮点击
	// 6 - 下载（预约）按钮点击（小流量）
	// 10 - 购买成功
	// 14 - 订单提交成功
	// 17 - 三句话咨询
	// 18 - 留线索
	// 19 - 一句话咨询
	// 20 - 关键页面浏览
	// 25 - 注册（小流量）
	// 26 - 付费（小流量）
	// 30 - 电话拨通
	// 35 - 微信复制按钮点击（小流量）
	// 41 - 申请（小流量）
	// 42 - 授信（小流量）
	// 45 - 商品下单成功
	// 46 - 加入购物车
	// 47 - 商品收藏
	// 48 - 商品详情页到达
	// 49 - 登录（注册激活后登录）
	// 57 - 店铺调起
	// 67 - 微信调起按钮点击
	// 68 - 粉丝关注成功
	// 71 - 应用调起
	// 72 - 聊到相关业务（小流量）
	// 73 - 回访-电话接通（小流量）
	// 56 - 到店（小流量）
	// 74 - 回访-信息确认（小流量）
	// 75 - 回访-发现意向（小流量）
	// 76 - 回访-高潜成交（小流量）
	// 77 - 回访-成单客户（小流量）
	// 79 - 微信加粉成功（小流量）
	// 80 - 直播间成单（小流量）
	// 82 - 直播间观看（小流量）
	// 83 - 直播间商品按钮点击（小流量）
	// 84 - 直播间停留（小流量）
	// 85 - 直播间评论（小流量）
	// 86 - 直播间打赏（小流量）
	// 87 - 直播间购物袋点击（小流量）
	// 90 - 商品支付成功
	// 93 - 付费阅读(小流量)
	// 118 - 付费观剧（小流量）
	// 119 - 关键行为（小流量）
	// 小程序-微信小游戏仅支持登录和付费转化目标。
	// 对于小流量字段仅部分名单客户可用，如果要申请可通过客服和销售同学反馈
	// 更新推广单元时不允许修改该字段
	TransType int `json:"transType,omitempty"`
	// OptimizeDeepTrans 优化深度转化;
	OptimizeDeepTrans *bool `json:"optimizeDeepTrans,omitempty"`
	// DeepOcpcBid 深度转化出价
	DeepOcpcBid float64 `json:"deepOcpcBid,omitempty"`
	// DeepTransType 深度转化类型
	// 取值范围：枚举值，列表如下
	// 10 - 购买成功
	// 18 - 留线索（小流量）
	// 25 - 注册（小流量）
	// 26 - 付费（小流量）
	// 27 - 客户自定义（小流量）
	// 28 - 次日留存（小流量）
	// 42 - 授信（小流量）
	// 45 - 商品下单成功
	// 53 - 订单核对成功
	// 54 - 收货成功
	// 72 - 聊到相关业务（小流量）
	// 73 - 回访-电话接通（小流量）
	// 56 - 到店（小流量）
	// 74 - 回访-信息确认（小流量）
	// 75 - 回访-发现意向（小流量）
	// 76 - 回访-高潜成交（小流量）
	// 77 - 回访-成单客户（小流量）
	// 79 - 微信加粉成功（小流量）
	// optimizeDeepTrans=true时必填。对于小流量字段仅部分名单客户可用，如果要申请可通过客服和销售同学反馈
	// 更新推广单元时不允许修改该字段
	DeepTransType int `json:"deepTransType,omitempty"`
	// UrlType 落地页类型
	// 取值范围：枚举值，列表如下
	// 1 - 普通落地页
	// 2 - 百度小程序
	// 3 - 直播间
	UrlType int `json:"urlType,omitempty"`
	// UseRoi 使用ROI优化
	// 不能与optimizeDeepTrans同时使用。目前仅支持部分目标转化，列表如下
	// 14 - 订单提交成功
	// 26 - 付费
	UseRoi *bool `json:"useRoi,omitempty"`
	// RoiRatio ROI转化率
	// 取值范围：[0.00, 100.00]
	// useRoi=true时必填，保留两位小数
	RoiRatio *float64 `json:"roiRatio,omitempty"`
	// MiniProgramType 百度小程序类型
	// 取值范围：枚举值，列表如下
	// 1 - 小程序
	// 2 - 小游戏
	// 落地页类型为百度小程序时必填
	MiniProgramType int `json:"miniProgramType,omitempty"`
	// AppKey 百度小程序appkey
	// 取值范围：[1, 32]
	// 落地页类型为百度小程序时必填
	AppKey string `json:"appKey,omitempty"`
	// PagePath 百度小程序页面路径
	// 落地页类型为百度小程序时必填
	PagePath string `json:"pagePath,omitempty"`
	// BroadCastMode 直播间投放模式
	//     取值范围：枚举值，列表如下
	// 1 - 默认
	// 2 - 仅直播时投放
	// 3 - 连续投放
	// 落地页为直播间时必填
	BroadCastMode int `json:"broadCastMode,omitempty"`
	// AnchorId 主播连续投放时的主播ID
	// 直播间投放模式为连续投放时必填
	AnchorId int64 `json:"anchorId,omitempty"`
	// TransTypeAttribute 付费次数优化
	// 取值范围：枚举值，列表如下
	// 1 - 表示按照付费人数优化
	// 2 - 表示按照付费次数优化
	// 该字段必须搭配transType=26使用，否则会报错，更新单元时该字段不可修改
	// 该功能目前正在小流量测试，如需使用请向营销顾问申请
	TransTypeAttribute int `json:"transTypeAttribute,omitempty"`
}
