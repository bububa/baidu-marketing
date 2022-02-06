package image

// GetImageResponse 获取图片素材 API Response
type GetImageResponse struct {
	Data []Image `json:"data,omitempty"`
}
