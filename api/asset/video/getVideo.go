package video

import (
	"github.com/xiaoshouchen/baidu-marketing/core"
	"github.com/xiaoshouchen/baidu-marketing/model"
	"github.com/xiaoshouchen/baidu-marketing/model/asset/video"
)

// GetVideo 获取信息流&搜索视频素材信息，括视频id、规格、大小、格式、上传日期、最后修改时间、长度、名称、视频URL等等
func GetVideo(clt *core.SDKClient, auth *model.RequestHeader, reqBody *video.GetVideoRequest) (*model.ResponseHeader, []video.Video, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp video.GetVideoResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data, err
}
