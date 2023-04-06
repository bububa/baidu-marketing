package report

// Summary 汇总行
type Summary struct {
	Date       string  `json:"date"`
	Click      int     `json:"click"`
	Cost       float64 `json:"cost"`
	CPC        float64 `json:"cpc"`
	Impression int     `json:"impression"`
	Conversion int     `json:"conversion"`
}
