package image

import (
	"fmt"

	"github.com/bububa/baidu-marketing/model"
)

type GetImageRequest struct {
	Resolution  [][]int  `json:"resolution,omitempty"`  // 图片规格，选填，二维数组，支持多个分辨率，分别表示为：[[宽，高]，…,[宽，高]]
	Size        []int    `json:"size,omitempty"`        // 图片大小区间，选填，数组长度为2，表示获取大小在[size[0],size[1]]这个区间的图片
	Format      []string `json:"format,omitempty"`      // 图片格式，选填，可选范围为上传图片时限定的范围。["jpg","png"]
	Date        []string `json:"date,omitempty"`        // 时间区间, 选填，数组长度为2，表示获取上传日期在[date[0], date[1]]这个区间的图片，时间格式为yyyy-MM-ddTHH:mm:ss
	Page        int      `json:"page"`                  // 当前页
	PageSize    int      `json:"pageSize"`              // 当前页大小
	IsCollected int      `json:"isCollected,omitempty"` // 图片收藏
}

func (r GetImageRequest) Url() string {
	return fmt.Sprintf("%sImageService/getImage", model.BASE_URL_FEED)
}
