package trans

// AppConversion 包含应用转化的信息
type OcpcTrans struct {
	AppTransId     int64    `json:"appTransId"`               // 转化追踪ID
	TransFrom      int      `json:"transFrom"`                // 接入方式，枚举值详见oCPC数据对象transFrom字段
	TransName      string   `json:"transName"`                // 转化名称
	TransTypes     []int    `json:"transTypes"`               // 转化类型，枚举值详见oCPC数据对象transTypes字段
	MonitorUrl     string   `json:"monitorUrl,omitempty"`     // 点击监测地址，仅当transFrom=1，4，13时返回该字段
	AppType        int      `json:"appType,omitempty"`        // 应用类型，仅当transFrom=1，4，13时返回该字段，枚举值详见oCPC数据对象appType字段
	DownloadUrl    string   `json:"downloadUrl,omitempty"`    // 下载URL，仅当transFrom=1，4，13时返回该字段
	ExposureUrl    string   `json:"exposureUrl,omitempty"`    // 曝光监测地址，仅当transFrom=1，4，13时返回该字段
	LpUrl          string   `json:"lpUrl,omitempty"`          // 转化URL，仅当transFrom=2，5，7，8时返回该字段
	RelatedUrls    []string `json:"relatedUrls,omitempty"`    // 推广URL，仅当transFrom=5，7，8时返回该字段
	Mode           int      `json:"mode,omitempty"`           // 监测方式，仅当transFrom=1，4，13时返回该字段，枚举值详见oCPC数据对象mode字段
	TransStatus    int      `json:"transStatus"`              // 激活状态，枚举值详见oCPC数据对象transStatus字段
	DeepTransTypes []int    `json:"deepTransTypes,omitempty"` // 深度转化类型，枚举值详见oCPC数据对象deepTransTypes字段
	Docid          string   `json:"docid,omitempty"`          // Android渠道包ID，仅当transFrom=1，且appType=2时返回该字段
	AppName        string   `json:"appName,omitempty"`        // 应用名称，仅当transFrom=1，4，13时返回该字段
	ApkName        string   `json:"apkName,omitempty"`        // 应用包名，仅当transFrom=1，4，13时返回该字段；当transFrom=1或4时，只有appType=2结果才包含该字段
	ChannelId      int64    `json:"channelId,omitempty"`      // Android渠道包ID，仅当transFrom=1，且appType=2时返回该字段
}

type OcpcTransResponse struct {
	Data []OcpcTrans
}
