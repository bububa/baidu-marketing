package enum

// Atype 表明用户的转化数据
type AType string

const (
	// AType_ACTIVATE 激活
	AType_ACTIVATE AType = "activate"
	// ATYPE_REGISTER 注册
	AType_REGISTER AType = "register"
	// ATYPE_ORDERS 付费（成单）
	AType_ORDERS AType = "orders"
	// ATYPE_RETAIN_1DAY 次日留存
	AType_RETAIN_1DAY AType = "retain_1day"
	// AType_USER_DEFINED 客户自定义
	AType_USER_DEFINED AType = "user_defined"
	// AType_ECBuy 商品下单成功
	AType_ECBuy AType = "ec_buy"
	// AType_DEEP_PAGE_ACCESS 深度页面访问
	AType_DEEP_PAGE_ACCESS AType = "deep_page_access"
	// ATYPE_CREDT_GRANTING 授信
	AType_CREDT_GRANTING AType = "credit_granting"
	// AType_DEEPLINK 应用调起
	AType_DEEPLINK AType = "deep_link"
	// AType_FEED_DEEPLINK 应用调起
	AType_FEED_DEEPLINK AType = "feed_deeplink"
	// AType_PAY_TO_READ 付费阅读
	AType_PAY_TO_READ AType = "pay_to_read"
	// AType_ENTER_BOOKSTORE_READ 进入书城阅读
	AType_ENTER_BOOKSTORE_READ AType = "enter_bookstore_read"
	// AType_ADD_TO_DESKTOP 添加到桌面
	AType_ADD_TO_DESKTOP AType = "add_to_desktop"
	// AType_LOGIN  登录
	AType_LOGIN AType = "log_in"
	// AType_ORDER_SUBMIT_SUCCESS 订单提交成功
	AType_ORDER_SUBMIT_SUCCESS AType = "order_submit_success"
	// AType_PAY_TO_WATCH 付费观剧
	AType_PAY_TO_WATCH AType = "pay_to_watch"
	// AType_KEY_ACTION 关键行为
	AType_KEY_ACTION AType = "key_action"
	// AType_DERIVED_EVENT 衍生事件
	AType_DERIVED_EVENT AType = "derived_event"
)
