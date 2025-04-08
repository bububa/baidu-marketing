package trans

// OcpcTransFeed 转化追踪数据对象
type OcpcTransFeed struct {
	// AppTransId 转化追踪ID
	AppTransId uint64 `json:"appTransId"`
	// TransFrom 接入方式
	// 取值范围：枚举值，列表如下
	// 0 - 包含全部接入方式
	// 1 - 应用API
	// 2 - 基木鱼营销页
	// 4 - API激活
	// 5 - 网页JS布码
	// 7 - 线索API
	// 8 - 咨询工具授权
	// 13 - 应用SDK
	TransFrom *int `json:"transFrom"`
	// TransType 转化类型
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
	// 20 - 深度页面访问
	// 30 - 电话拨通
	// 35 - 微信复制按钮点击（小流量）
	// 41 - 申请（小流量）
	// 42 - 授信（小流量）
	// 45 - 商品下单成功
	// 46 - 加入购物车
	// 47 - 商品收藏
	// 48 - 商品详情页到达
	// 57 - 店铺调起
	// 67 - 微信调起按钮点击
	// 68 - 粉丝关注成功
	// 71 - 应用调起
	// 对于小流量字段仅部分名单客户可用，如果要申请可通过客服和销售同学反馈
	TransTypes []int `json:"transTypes"`
	// MonitorUrl 点击监测地址，仅当transFrom=1，4，13时返回该字段
	MonitorUrl *string `json:"monitorUrl,omitempty"`
	// LpUrl 转化URL，仅当transFrom=2，5，7，8时返回该字段
	LpUrl *string `json:"lpUrl,omitempty"`
	// RelatedUrls 推广URL
	RelatedUrls []string `json:"relatedUrls,omitempty"`
	// TransName 转化名称
	TransName *string `json:"transName"`
	// AppType 应用类型
	// 取值范围：枚举值，列表如下
	// 1 - iOS
	// 2 - Android
	// 3 - 应用API-调起
	AppType int `json:"appType,omitempty"`
	// DownloadUrl 下载URL
	DownloadUrl *string `json:"downloadUrl,omitempty"`
	// ExposureUrl 曝光监测地址
	ExposureUrl *string `json:"exposureUrl,omitempty"`
	// Mode 监测方式
	// 取值范围：枚举值，列表如下
	// 0 - 点击+曝光
	// 1 - 点击
	Mode *int `json:"mode,omitempty"`
	// TransStatus 激活状态
	// 取值范围：枚举值，列表如下
	// 0 - 无效默认值
	// 1 - 已激活
	// 2 - 未激活
	TransStatus int `json:"transStatus"`
	// DeepTransTypes 深度转化类型
	// 取值范围：枚举值，列表如下
	// 10 - 购买成功
	// 25 - 注册（小流量）
	// 26 - 付费（小流量）
	// 27 - 客户自定义（小流量）
	// 28 - 次日留存（小流量）
	// 42 - 授信（小流量）
	// 45 - 商品下单成功
	// 53 - 订单核对成功
	// 54 - 收货成功
	// 对于小流量字段仅部分名单客户可用，如果要申请可通过客服和销售同学反馈
	DeepTransTypes []int `json:"deepTransTypes,omitempty"`
	// Docid Android渠道包ID
	// 旧平台Android渠道包对应ID，目前已不支持通过docid新建转化追踪；对于早先使用docid字段进行创建的转化追踪，查询结果仍然会返回docid字段
	Docid uint64 `json:"docid,omitempty"`
	// AppName 应用名称
	AppName string `json:"appName,omitempty"`
	// ApkName 应用包名
	ApkName string `json:"apkName,omitempty"`
	// AppId APP Store ID
	AppId uint64 `json:"appId,omitempty"`
	// SdkAppId SDK应用ID
	SdkAppId uint64 `json:"sdkAppId,omitempty"`
	// SdkSecretKey SDK应用密钥
	SdkSecretKey string `json:"sdkSecretKey,omitempty"`
	// ChannelId Android渠道包ID
	ChannelId uint64 `json:"channelId,omitempty"`
}
