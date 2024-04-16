package creative

// Creative 创意
type Creative struct {
	// CreativeId 创意ID
	CreativeId uint64 `json:"creativeFeedId,omitempty"`
	// AdgroupFeedId 所属推广单元ID
	AdgroupId uint64 `json:"adgroupFeedId,omitempty"`
	// Materialstyle 创意样式ID
	Materialstyle int `json:"materialstyle,omitempty"`
	// CreativeName 创意名称
	CreativeName *string `json:"creativeFeedName,omitempty"`
	// Pause 是否暂停推广
	Pause *bool `json:"pause,omitempty"`
	// Status 创意状态
	// 取值范围：枚举值，列表如下
	// 0 - 有效
	// 1 - 暂停推广
	// 2 - 审核中
	// 3 - 审核未通过
	// 4 - 无效（因与所属推广单元所选的投放版位不匹配而无法投放，推广单元修改投放版位后将自动重置此状态。请参考信息流创意样式列表的说明。）
	// 6 - 未审核
	Status int `json:"status,omitempty"`
	// Material 物料内容
	// 物料内容为string形式的json对象，包含该创意所需的标题、图片、落地页URL等展示物料，以及监测URL等功能性物料。各种创意样式所需的物料字段有所不同，请参考信息流创意物料说明和信息流创意样式列表
	Material string `json:"material,omitempty"`
	// Refusereason 审核未通过的原因（审核拒绝理由）
	// 原因可能有多个，该字段为string形式的json数组，数组元素为string类型，每个元素是一个原因。
	// 示例：
	// "[\"原因A\",\"原因B\"]"
	Refusereason string `json:"refusereason,omitempty"`
	// Playnum 视频创意的视频播放量
	// 取值范围：枚举值，列表如下
	// 0 - 关闭
	// 1 - 开启

	// 仅支持如下样式ID：
	// 108：大图视频链接
	// 109：大图视频下载
	// 113：程序化创意下载
	// 114：程序化创意链接
	Playnum *int `json:"playnum,omitempty"`
	// IdeaType 创意类型
	// 取值范围：枚举值，列表如下
	// 0 - 自定义创意
	// 1 - 程序化创意
	// 小程序-微信小游戏仅支持程序化创意
	IdeaType *int `json:"ideaType,omitempty"`
	// Addtime 创意的创建时间，仅获取创意列表时传递有效
	Addtime string `json:"addtime,omitempty"`
	// Approvemsgnew JSON格式化后的审核未通过的原因（审核拒绝理由）
	// 原因可能有多个，该字段为string形式的json数组，数组元素为string类型，每个元素是一个原因。
	// 示例：
	// {"audit_detail":[{"desc":"拒绝字段名称","text":"拒绝理由"}],"image_detail":[{"img":"图片URL","text":"拒绝理由"}],"detail":[{"img":"图片URL","text":"拒绝理由"}],"global_reason":""}
	// audit_detail: 文本拒绝理由 1.非精标产品线拒绝理由：desc为物料内容，text为物料拒绝理由。 2.精标产品线拒绝理由：desc为拒绝字段（业务方提供），text为字段拒绝理由。
	// image_detail: 图片拒绝理由 1.精标产品线拒绝理由：desc为拒绝理由字段（业务方提供），text为图片拒绝理由 2.非精标产品线：空数组
	// detail: 截图拒绝理由 1.截图产品线：desc为拒绝理由的url，text为该拒绝理由截图话术 2.非截图产品线：空数组
	// global_reason: 历史整体拒绝理由
	// 特殊说明：创意为拒绝时该字段可能为空，该字段为空时请取refusereason字段
	Approvemsgnew string `json:"approvemsgnew,omitempty"`
	// AuditTimeModel 预估审核时间
	AuditTimeModel *AuditTimeModel `json:"auditTimeModel,omitempty"`
	// Commentnum 评论数
	// 取值范围：枚举值，列表如下
	// 0 - 关闭
	// 1 - 开启
	// 单选小说流量不支持，且仅支持营销目标为应用下载
	Commentnum *int `json:"commentnum,omitempty"`
	// Readnum 阅读量
	// 取值范围：枚举值，列表如下
	// 0 - 关闭
	// 1 - 开启

	// 单选小说流量不支持，仅支持如下样式ID：：
	// 101：单图链接样式
	// 102：三图链接样式
	// 107：大图链接样式
	// 110：百度APP橱窗URL
	// 171：百度APP-首页-单品单图-商品URL
	// 172：百度APP-首页-单品三图-商品URL
	// 178：百度APP-首页-单品大图-商品URL
	// 173：百度APP-首页-三品三图-商品URL
	// 187：百度APP-首页-橱窗-商品URL
	// 114：程序化创意链接
	Readnum *int `json:"readnum,omitempty"`
	// Template 商品创意物料
	// 仅商品目录创意使用, 商品创意的具体物料内容，请参照下方Template对象的具体定义。相当于非商品目录的material字段，可参考信息流创意样式列表填写
	Template *PaIdeaTemplate `json:"template,omitempty"`
}

// AuditTimeModel 预估审核时间
type AuditTimeModel struct {
	// IdeaId 创意ID
	IdeaId uint64 `json:"ideaId,omitempty"`
	// IdeaState 审核状态
	// 取值范围：枚举值，列表如下
	// 0 - 审核中
	// 1 - 审核超时
	// 2 - 由于物料特殊无法预估审核时间
	// 3 - 审核预估时间计算中，请稍后查看
	IdeaState []int `json:"ideaState,omitempty"`
	// StartTime 审核开始时间
	StartTime string `json:"startTime,omitempty"`
	// EstimationTime 预计审核结束时间
	EstimationTime string `json:"estimationTime,omitempty"`
	// EstimationMinute 预计审核时长
	EstimationMinute []int `json:"estimationMinute,omitempty"`
	// CompleteRatio 完成百分比
	CompleteRadio string `json:"completeRadio,omitempty"`
}

// PaIdeaTemplate 商品创意物料
type PaIdeaTemplate struct {
	// Title 标题
	// 可使用商品库字段
	// 示例:
	// 大多数人不知道，这里买${name}更便宜
	Title string `json:"title,omitempty"`
	// Subtitle 副标题
	// 可使用商品库字段
	Subtitle string `json:"subtitle,omitempty"`
	// TargetUrl 推广链接
	// 可使用商品库字段，如${targetUrl}
	TargetUrl string `json:"targetUrl,omitempty"`
	// MonitorUrl 监控链接
	// 可使用商品库字段，如${monitorUrl}
	MonitorUrl string `json:"monitorUrl,omitempty"`
	// UserPortrait 用户头像
	UserPortrait string `json:"userPortrait,omitempty"`
	// NaUrl 调起链接
	// 仅在选取网站链接&应用调起类型商品计划时可用。可使用商品库字段填写deeplink，如判断用户安装了app，可直接从浏览器调起app，否则跳转至点击跳转url（目前仅支持手机百度首页，如投放了其他流量，将直接跳转至点击跳转url）
	NaUrl string `json:"naUrl,omitempty"`
	// Brand 品牌名
	Brand string `json:"brand,omitempty"`
	// Pictures 图片列表
	// 主尺寸图片列表。对应于创意列表中尺寸列表内的第一个尺寸
	// 非单选百青藤流量时必填，单选百青藤流量时选填
	// 根据样式所需的图片个数填写相应个数的元素，多图片样式部分图片自提时，对应的图片填写为相应的图片url
	// 示例
	// （单品三图）：[{"url":"${image}"},{"url":"http://example.com/example.jpg"},{"url":"${image}"}]
	Pictures []PaPicture `json:"pictures,omitempty"`
	// ExtraPictures 多尺寸图片列表
	// 可填写当前样式支持的更多尺寸;
	// 仅在默认流量，或自定义流量包含百青藤流量时可用;
	// 百青藤流量创意图片所有尺寸自提应均使用extraPictures字段，pictures字段不传;
	// 示例（单品大图）：[{"pictures":[{"image":"http://www.sample.jpg"}],"width":960,"height":640}]，
	// 示例（单品三图）:[{"pictures":[{"image":"http://www.sample1.jpg"},{"image":"http://www.sample2.jpg"},{"image":"http://www.sample3.jpg"}],"width":600,"height":300}]
	ExtraPictures []PaExtraPicture `json:"extraPictures,omitempty"`
	// Poster 视频封面
	//     仅视频样式使用
	// 可使用商品库字段
	// 横版视频未填写时，将复用pictures内第一个元素（图片物料）的url
	// 竖版视频未填写时，将使用商品库字段verticalPoster
	Poster string `json:"poster,omitempty"`
	// VideoId 视频id
	// 仅视频样式使用
	// 使用商品库字段时不填写，视频自提时填写
	// 自提视频ID可以通过相关接口，或信息流推广Web端的"视频库"中获取
	VideoId uint64 `json:"videoId,omitempty"`
	// VideoUrl 视频url
	// 仅视频样式使用
	// 可使用商品库字段，视频自提时与videoId选填一个
	// videoId与videoUrl均不填写时，横版视频将使用商品库字段videoUrl
	// 竖版视频将使用商品库字段verticalVideoUrl
	// 自提视频地址可以通过相关接口获取
	VideoUrl string `json:"videoUrl,omitempty"`
	// MiniProgramSchema 小程序链接
	// 可使用商品库字段
	MiniProgramSchema string `json:"miniProgramSchema,omitempty"`
	// ExposureUrl 曝光监测链接
	// 支持选填.可使用商品库字段，如${exposureUrl}
	ExposureUrl string `json:"exposureUrl,omitempty"`
	// UlkScheme ulk调起链接协议头
	// ulk调起链接协议头，选填，可使用商品库字段，非安卓样式且单元设备未选安卓时可用
	UlkScheme string `json:"ulkScheme,omitempty"`
	// UlkUrl ulk调起链接
	// ulk调起链接，选填，可使用商品库字段，非安卓样式且单元设备未选安卓时可用
	UlkUrl string `json:"ulkUrl,omitempty"`
	// Elements 程序化创意元素
	// 程序化创意信息，仅商品目录程序化创意(528、529)使用
	Elements *ProgramElements `json:"elements,omitempty"`
	// StyleList 样式
	// 取值范围：枚举值，列表如下
	// singlePic - 单图
	// largePic - 大图
	// package3Pic - 三图
	// horizontalVideos - 横版视频
	// verticalVideos - 竖版视频
	// 仅商品目录程序化创意(528、529使用，填写程序化创意支持的样式类型，枚举值数组)
	StyleList []string `json:"styleList,omitempty"`
	// PaCommon 商品创意设置
	// 仅商品目录创意使用，非必填
	PaCommon *PaIdeaCommon `json:"paCommon,omitempty"`
	// Huitus 创意绑定的慧图id列表
	// 仅商品目录创意使用，非必填，通过查询慧图接口 获取userTemplateId 并进行绑定，使用时请注意以下条件：
	// 1. 一次调用中，同一个单元下相同materialstyle的多个，huitus需相同，若不同则取并集；多次调用时，huitus会覆盖
	// 例：添加单元A materialstyle为501的两个创意1和创意2，创意1传参huitu1，创意2传参慧图2
	// 则单元A materialstyle501对应的慧图为1和2
	// 添加单元A materialstyle为501的创意1传参huitu1，此时对应慧图为1
	// 再次调接口添加单元A materialstyle为501的创意2传参huitu2，此时对应慧图为2
	// 即：同一单元同一样式下的创意对应同一组慧图
	// 2. 若慧图id不存在或id对应的type和创意样式不符合，则新建或更新创意全部失败
	Huitus []uint64 `json:"huitus,omitempty"`
	// ProgFlag 创意工具
	// 仅支持程序化创意，取值范围：枚举值，列表如下
	// [] - 不支持创意工具
	// [32] - 支持创意工具
	ProgFlag []int `json:"progFlag,omitempty"`
}

type PaPicture struct {
	// Url 图片地址
	// 使用商品库字段时固定填写${image}
	// 使用自提图片时，创意图片必须先使用ImageService的相关接口上传，使用接口返回的URL进行投放。
	// 橱窗样式不支持自提图片
	Url string `json:"url,omitempty"`
	// TargetUrl 图片跳转地址
	TargetUrl string `json:"targetUrl,omitempty"`
	// NaUrl 调起链接
	NaUrl string `json:"naUrl,omitempty"`
	// Desc1 图片描述1
	Desc1 string `json:"desc1,omitempty"`
	// Desc2 图片描述2
	Desc2 string `json:"desc2,omitempty"`
}

type PaExtraPicture struct {
	// Pictures 多尺寸图片列表
	// 图片链接列表。创意图片必须先使用ImageService的相关接口上传，使用接口返回的URL进行投放。元素个数应为创意样式所需图片个数。
	// 不支持商品通配符。
	// 示例（单品大图）：
	// [{"image":"http://www.sample.jpg"}]
	// 示例（单品三图）：
	// [{"image":"http://www.sample1.jpg"},{"image":"http://www.sample2.jpg"},{"image":"http://www.sample3.jpg"}]
	Pictures []ExtraPictureImage `json:"pictures,omitempty"`
	// Width 宽度
	Width int `json:"width,omitempty"`
	// Height 高度
	Height int `json:"height,omitempty"`
}

type ExtraPictureImage struct {
	// Image 图片地址
	Image string `json:"image,omitempty"`
}

// ProgramElements 程序化创意元素
type ProgramElements struct {
	// Title 标题
	// 最多6个标题，每个标题不超过60个字符
	Title []string `json:"title,omitempty"`
	// SinglePic 单图图片地址
	// 最多6张图片
	SinglePic []string `json:"singlePic,omitempty"`
	// LargePic 大图图片地址
	// 最多6张图片
	LargePic []string `json:"largePic,omitempty"`
	// Package3Pic 三图
	// 最多6套图片，单套三张图片地址用空格分割
	Package3Pic []string `json:"package3Pic,omitempty"`
	// HorizontalVideos 横版视频
	// 视频地址 空格 视频id 空格 视频封面图片地址
	HorizontalVideos []string `json:"horizontalVideos,omitempty"`
	// VerticalVideos 竖版视频
	// 视频地址 空格 视频id 空格 视频封面图片地址
	VerticalVideos []string `json:"verticalVideos,omitempty"`
}

// PaIdeaCommon 商品创意设置
type PaIdeaCommon struct {
	// PictureType 程序化图片视频形式
	// 取值范围：枚举值，列表如下
	// 0 - 默认单图大图使用${image},三图使用${img index},横版视频使用${videoUrl},竖版视频使用${verticalVideoUrl}
	// 1 - 自行在elements对象内提交图片、视频物料
	// 仅商品目录程序化创意(528、529)使用，程序化物料是使用商品库信息还是自行填写
	PictureType int `json:"pictureType,omitempty"`
}
