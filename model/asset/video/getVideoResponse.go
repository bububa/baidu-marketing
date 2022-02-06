package video

// GetVideoResponse 获取视频素材 API Response
type GetVideoResponse struct {
	Data []Video `json:"data,omitempty"`
}
