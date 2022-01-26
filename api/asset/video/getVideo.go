package video

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/asset/video"
)

// 获取信息流&搜索视频素材信息，括视频id、规格、大小、格式、上传日期、最后修改时间、长度、名称、视频URL等等
func GetVideo(clt *core.SDKClient, auth model.RequestHeader, reqBody *video.GetVideoRequest) ([]video.Video, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp video.GetVideoResponse
	err := clt.Do(req, &resp)
	if err != nil {
		return resp.Data, err
	}
	return resp.Data, nil
}
