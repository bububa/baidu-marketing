package campaign

// OcpcModel oCPC相关设置
type OcpcModel struct {
	// OcpcBid 目标转化出价
	// 取值范围：[0.1, 99999.99]
	// 单位为元
	OcpcBid float64 `json:"ocpcBid,omitempty"`
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
	// 56 - 到店（小流量）
	// 57 - 店铺调起
	// 67 - 微信调起按钮点击
	// 68 - 粉丝关注成功
	// 71 - 应用调起
	// 72 - 聊到相关业务（小流量）
	// 73 - 回访-电话接通（小流量）
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
	// 计划选择网站链接营销目标后，无法选择应用推广的转化目标
	TransType int `json:"transType,omitempty"`
	// TransFrom 接入方式
	// 取值范围：枚举值，列表如下
	// 1 - 应用API
	// 2 - 基木鱼营销页
	// 4 - API激活
	// 5 - 网页JS布码
	// 7 - 线索API
	// 8 - 咨询工具授权
	// 13 - 应用SDK
	// 25 - 百度健康商城
	// 30 - 百度小程序导入
	TransFrom int `json:"transFrom,omitempty"`
	// AppTransId 转化追踪ID
	// 营销目标为应用推广Android的单元，仅能绑定应用操作系统为Android的转化追踪；
	// 营销目标为应用推广iOS的单元，仅能绑定应用操作系统为iOS的转化追踪
	AppTransId uint64 `json:"appTransId,omitempty"`
	// DeepOcpcBid 深度转化出价
	// 取值范围：[0.4, 99999.99]
	// 单位为元，取值需要大于ocpcBid，optimizeDeepTrans=true时必填
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
	// 56 - 到店（小流量）
	// 72 - 聊到相关业务（小流量）
	// 73 - 回访-电话接通（小流量）
	// 74 - 回访-信息确认（小流量）
	// 75 - 回访-发现意向（小流量）
	// 76 - 回访-高潜成交（小流量）
	// 77 - 回访-成单客户（小流量）
	// 79 - 微信加粉成功（小流量）
	// optimizeDeepTrans=true时必填。对于小流量字段仅部分名单客户可用，如果要申请可通过客服和销售同学反馈
	// 更新推广单元时不允许修改该字段
	DeepTransType int `json:"deepTransType,omitempty"`
	// UseRoi 使用ROI优化
	// 默认值：false
	// 不能与optimizeDeepTrans同时使用。目前仅支持部分目标转化，列表如下
	// 14 - 订单提交成功
	// 26 - 付费
	UseRoi *bool `json:"useRoi,omitempty"`
	// RoiRatio ROI转化率
	// 取值范围：[0.00, 100.00]
	// useRoi=true时必填，保留两位小数
	RoiRatio float64 `json:"roiRatio,omitempty"`
	// TransTypeAttribute 付费次数优化
	// 默认值：1
	// 取值范围：枚举值，列表如下
	// 1 - 表示按照付费人数优化
	// 2 - 表示按照付费次数优化
	// 该字段必须搭配transType=26使用，否则会报错，更新单元时该字段不可修改
	// 该功能目前正在小流量测试，如需使用请向营销顾问申请
	TransTypeAttribute int `json:"transTypeAttribute,omitempty"`
	// OptimizeDeepTrans 优化深度转化
	// 默认值：false
	// 仅当appTransId对应的转化追踪具有可选的深度转化类型时有效
	OptimizeDeepTrans *bool `json:"optimizeDeepTrans,omitempty"`
	// MiniProgramType 小程序子类型
	// 取值范围：枚举值，列表如下4 - 微信小游戏（需要开通小流量名单）
	MiniProgramType int `json:"miniProgramType,omitempty"`
}
