package video

import (
	"context"

	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/asset/video"
)

// GetVideo 获取信息流&搜索视频素材信息，括视频id、规格、大小、格式、上传日期、最后修改时间、长度、名称、视频URL等等
func GetVideo(ctx context.Context, clt *core.SDKClient, auth *model.RequestHeader, reqBody *video.GetVideoRequest) (*model.ResponseHeader, []video.Video, error) {
	req := &model.Request{
		Header: auth,
		Body:   reqBody,
	}
	var resp video.GetVideoResponse
	header, err := clt.Do(ctx, req, &resp)
	return header, resp.Data, err
}
