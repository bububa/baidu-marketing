package image

import (
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/util"
)

// GetImageRequest 获取图片素材API Request
type GetImageRequest struct {
	// Resolution 图片规格，选填，二维数组，支持多个分辨率，分别表示为：[[宽，高]，…,[宽，高]]
	Resolution [][]int `json:"resolution,omitempty"`
	// Size 图片大小区间，选填，数组长度为2，表示获取大小在[size[0],size[1]]这个区间的图片
	Size []int `json:"size,omitempty"`
	// Format 图片格式，选填，可选范围为上传图片时限定的范围。["jpg","png"]
	Format []string `json:"format,omitempty"`
	// Date 时间区间, 选填，数组长度为2，表示获取上传日期在[date[0], date[1]]这个区间的图片，时间格式为yyyy-MM-ddTHH:mm:ss
	Date []string `json:"date,omitempty"`
	// Page 当前页
	Page int `json:"page"`
	// PageSize 当前页大小
	PageSize int `json:"pageSize"`
	// IsCollectec 图片收藏
	IsCollected int `json:"isCollected,omitempty"`
}

func (r GetImageRequest) Url() string {
	return util.StringsJoin(model.BASE_URL_FEED, "ImageService/getImage")
}

// GetImageResponse 获取图片素材 API Response
type GetImageResponse struct {
	Data []Image `json:"data,omitempty"`
}
