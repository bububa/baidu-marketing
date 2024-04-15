package campaign

import "github.com/bububa/baidu-marketing/model"

// Campaign 推广计划
type Campaign struct {
	// CampaignId 计划ID
	CampaignId uint64 `json:"campaignId,omitempty"`
	// CampaignName 计划名称;长度限制：最大30个字节（1个中文按2个字节计算，英文、数字按1个字节计算）
	CampaignName string `json:"campaignName,omitempty"`
	// Budget 计划每日预算;取值范围：[50, Min(10000000, 账户预算)]
	Budget float64 `json:"budget,omitempty"`
	// RegionTarget 计划推广地域
	RegionTarget []int64 `json:"regionTarget,omitempty"`
	// NegativeKeywords 短语否定关键词列表
	// 空数组[]-不设置短语否定关键词；
	// 根据用户等级（账户对象userLevel字段）上限不等：未生效客户(4)：200；一星客户(3)：200； 二星客户(2): 400；三星客户(1): 500；
	// 注：商品计划尚未支持此字段
	NegativeKeywords []string `json:"negativeKeywords,omitempty"`
	// ExactNegativeKeywords 精确否定关键词列表
	//     空数组[]-不设置精确否定关键词；
	// 根据用户等级（账户对象userLevel字段）上限不等：未生效客户(4)：200；一星客户(3)：400；二星客户(2)：700；三星客户(1)：900；
	// 注：商品计划尚未支持此字段
	ExactNegativeKeywords []string `json:"exactNegativeKeywords,omitempty"`
	// Schedule 计划推广暂停时段
	// 默认不设置，表示无推广暂停时段限制；
	// 数组元素个数限制：每天可设置最多12个推广暂停时间，每周可设置最多84个推广暂停时间；
	// 在updateCampaign接口中，如果该字段设置为空数组,即"schedule":[]，表示清空原有暂停时段设置"
	Schedule []model.Schedule `json:"schedule,omitempty"`
	// BudgetOfflineTime 预算下线时间;数组元素个数限制：最近有过下线时段的7个自然日的下线和上线时段（这7个自然日中若某日期距当前已超过30天，则不返回）;
	BudgetOfflineTime []model.OfflineTime `json:"budgetOfflineTime,omitempty"`
	// Pause 暂停状态;true - 暂停;false - 启用
	Pause *bool `json:"pause,omitempty"`
	// Status 计划状态
	// 21 - 有效
	// 22 - 处于暂停时段
	// 23 - 暂停推广
	// 24 - 计划预算不足
	// 25 - 账户预算不足
	Status int `json:"status,omitempty"`
	// AdType 计划类型
	// 默认值：0
	// 取值范围：枚举值，列表如下
	// 0 - 普通计划
	// 14 - 商品计划
	// 6 - 网址定向计划
	AdType *int `json:"adType,omitempty"`
	// BusinessPointId 推广业务ID
	BusinessPointId uint64 `json:"businessPointId,omitempty"`
	// BusinessPointName 推广业务字面
	BusinessPointName string `json:"businessPointName,omitempty"`
	// SmartRegion 商品计划: 智能地域开关
	SmartRegion bool `json:"smartRegion,omitempty"`
	// PaDevice 商品计划: 计划的投放设备
	// 默认值：0
	// 取值范围：枚举值，列表如下
	// 0 - 全部
	// 1 - 移动
	// 2 - 计算机
	// 商品目录新建计划equipmentType与paDevice投放设备类型保持一致
	// 商品目录更新计划不支持修改equipmentType与paDevice
	PaDevice *int `json:"paDevice,omitempty"`
	// Os 商品计划: 计划的投放设备平台
	// 默认值：全选
	// 取值范围：枚举值，列表如下
	// IPHONE - 苹果手机
	// ANDROID - 安卓手机
	// OTHERS - 其他类型
	Os []string `json:"os,omitempty"`
	// RegionPriceFactor 分地域出价系数
	// 选填，默认为账户投放地域，出价系数为1
	RegionPriceFactor []model.RegionPriceFactor `json:"regionPriceFactor,omitempty"`
	// SchedulePriceFactor 分时段出价系数
	// 选填，默认为全时段投放，出价系数为1
	SchedulePriceFactor []model.SchedulePriceFactor `json:"schedulePriceFactor,omitempty"`
	// MarketingTargetId 营销目标类型
	// 取值范围：枚举值，列表如下
	// 0 - 网站链接
	// 1 - 应用推广
	// 2 - 本地推广
	// 4 - 电商店铺推广
	// 5 - 商品目录
	MarketingTargetId *int `json:"marketingTargetId,omitempty"`
	// ShopType 电商店铺类型
	// 取值范围：枚举值，列表如下
	// 1 - 度小店
	// 3 - 第三方店铺
	// 31 - 淘宝（含天猫）
	// 32 - 京东
	// 33 - 拼多多
	// 34 - 苏宁
	// 当营销目标为"电商店铺推广"时必填，其他营销目标不支持
	ShopType int `json:"shopType,omitempty"`
	// EquipmentType  推广设备
	// 默认值：3
	// 取值范围：枚举值，列表如下
	// 1 - 计算机
	// 2 - 移动
	// 3 - 不限
	// marketingTargetId=0，equipmentType可取1,2,3
	// marketingTargetId=1，equipmentType可取2,3
	// marketingTargetId=2，equipmentType可取2,3
	// marketingTargetId=4，shopType=1，equipmentType可取2,3
	// marketingTargetId=4，shopType=3,31,32,33,34，equipmentType可取1,2,3
	// marketingTargetId=5，equipmentType可取1,2,3
	// 商品目录新建计划equipmentType与paDevice投放设备类型保持一致
	// 商品目录更新计划不支持修改equipmentType与paDevice
	EquipmentType int `json:"equipmentType,omitempty"`
	// CampaignBidType 计划出价方式
	// 默认值：0
	// 取值范围：枚举值，列表如下
	// 0 - 点击
	// 1 - 转化
	// 该字段不支持修改
	CampaignBidType *int `json:"campaignBidType,omitempty"`
	// CampaignBid 计划点击出价
	// 默认值：0
	// 取值范围：[0.01, 999.99]
	// 单位：元/点击
	// 不填写会使用单元出价
	// campaignBidType=0时api选填
	// campaignBidType=1时不支持该字段
	CampaignBid *float64 `json:"campaignBid,omitempty"`
	// CampaignOcpcBidType 计划出价模式
	// 取值范围：枚举值，列表如下
	// 0 - cpc
	// 1 - 目标转化成本
	// 2 - 增强模式
	// 3 - 放量模式
	// campaignBidType=0时，默认值0: cpc
	// campaignBidType=1时，默认值1: 目标转化成本
	// 当campaignBidType=0，修改时仅支持0,2
	// 当campaignBidType=1，修改时仅支持1,3
	CampaignOcpcBidType *int `json:"campaignOcpcBidType,omitempty"`
	// CampaignOcpcBid 转化计划出价
	// 默认值：0
	// 取值范围：[0.1, 9999]
	// 单位：元/转化
	// campaignOcpcBidType=1时必填，否则禁止填写
	CampaignOcpcBid *float64 `json:"campaignOcpcBid,omitempty"`
	// CampaignTransType 计划目标转化
	// 取值范围：枚举值，列表如下
	// 1 - 咨询按钮点击
	// 2 - 电话按钮点击
	// 3 - 表单提交成功
	// 4 - APP激活
	// 5 - 表单按钮点击
	// 6 - 下载（预约）按钮点击
	// 7 - 购买按钮点击
	// 9 - 电商订单
	// 10 - 服务购买成功
	// 12 - 预约按钮点击
	// 13 - 表单有效请求
	// 14 - 订单提交成功
	// 15 - 加入购物车按钮点击
	// 16 - 表单调起按钮点击
	// 17 - 三句话咨询
	// 18 - 留线索
	// 19 - 一句话咨询
	// 20 - 关键页面浏览
	// 25 - APP注册
	// 26 - APP付费
	// 27 - 客户自定义
	// 28 - 次日留存
	// 30 - 电话拨通>
	// 35 - 微信复制按钮点击
	// 36 - 评价按钮点击
	// 40 - 关注按钮点击
	// 41 - 申请小额贷款
	// 42 - 授信
	// 43 - 小程序线索
	// 45 - 商品下单成功
	// 46 - 加入购物车
	// 49 - 登录
	// 50 - 预约
	// 52 - 深度使用
	// 56 - 到店
	// 57 - 调起点击
	// 67 - 微信调起按钮点击
	// 68 - 粉丝关注成功
	// 71 - 应用调起(小流量)
	// 72 - 聊到相关业务(小流量)
	// 73 - 回访-电话接通(小流量)
	// 74 - 回访-信息确认(小流量)
	// 75 - 回访-发现意向(小流量)
	// 76 - 回访-高潜成交(小流量)
	// 77 - 回访-成单客户(小流量)
	// 78 - 店铺停留(小流量)
	// 79 - 微信加粉成功
	// 89 - 放款
	// 90 - 商品支付成功
	// 92 - 有效咨询(小流量)
	// 99 - 其他
	// campaignBidType=1时必填，
	// campaignOcpcBidType = 2 时选填，其他情况禁止填写
	// 数组长度限制[1, 5]，且目标转化+深度转化数量之和要小于等于5
	// 只能选择已在转化追踪工具添加的目标转化
	CampaignTransTypes []int `json:"campaignTransTypes,omitempty"`
	// CampaignDeepTransTypes 计划深度转化
	// 取值范围：枚举值，列表如下
	// 9 - 电商订单
	// 10 - 服务购买成功
	// 14 - 订单提交成功
	// 18 - 留线索
	// 25 - APP注册
	// 26 - APP付费
	// 28 - 次日留存
	// 30 - 电话拨通
	// 42 - 授信
	// 49 - 登录
	// 50 - 预约
	// 52 - 深度使用
	// 56 - 到店
	// 72 - 聊到相关业务(小流量)
	// 73 - 回访-电话接通(小流量)
	// 74 - 回访-信息确认(小流量)
	// 75 - 回访-发现意向(小流量)
	// 76 - 回访-高潜成交(小流量)
	// 77 - 回访-成单客户(小流量)
	// 79 - 微信加粉成功
	// 89 - 放款
	// campaignOcpcBidType = 0 禁止填写
	// 需为转化追踪下同一类型目标转化，数组长度限制[1, 5]，且目标转化+深度转化数量之和要小于等于5
	CampaignDeepTransTypes []int `json:"campaignDeepTransTypes,omitempty"`
	// CampaignCvSources 数据来源
	// 取值范围：枚举值，列表如下
	// 1000 - 不限
	// 1 - 网页JS布码
	// 2 - 线索API
	// 3 - 咨询工具授权
	// 4 - 基木鱼/度小店
	// 5 - 应用API
	// 6 - 电话数据授权
	// 7 - 百度智能小程序SDK
	// 8 - 应用SDK
	// 9 - 爱番番
	// 10 - 百度APP
	// 23 - 百度统计网站导入
	// 24 - 百度统计小程序导入
	// 非必填选项，默认为不限
	// web端在增强模式，目标转化成本，放量模式下可选，会根据客户联调或可用数据做过滤，且在投放设备包括pc时有文案提示
	// 仅展示客户联调或可用数据接入方式
	CampaignCvSources []int `json:"campaignCvSources,omitempty"`
	// StorePageInfos 本地计划设置的门店落地页信息
	StorePageInfos []model.StorePageInfo `json:"storePageInfos,omitempty"`
	// TransAsset 转化资产类型 取值范围：枚举值，列表如下
	// 0 - 不限
	// 2 - 指定资产
	TransAsset *int `json:"transAsset,omitempty"`
	// TransAssetId 转化资产ID
	TransAssetId uint64 `json:"transAssetId,omitempty"`
	// GeolocationStatus 推广地域地理位置选项
	// 取值范围：枚举值，列表如下
	// 0 - 该地区内或搜索意图在该地区的所有用户（包含：正在该地区的用户、长时间内居住或者工作在该地区的用户、在搜索词中对该地区表现出明确兴趣的用户）
	// 1 - 该地区内的所有用户（包含：正在该地区的用户、长时间内居住或者工作在该地区的用户）
	GeolocationStatus *int `json:"geoLocationStatus,omitempty"`
}
