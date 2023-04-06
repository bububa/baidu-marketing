package report

type GetAdgroupReportDataResponse struct {
	Data []GetAdgroupReportData `json:"data,omitempty"`
}

type GetAdgroupReportData struct {
	// Rows 数据行
	Rows []AdgroupReport `json:"rows,omitempty"`
	// Summary 汇总行
	Summary Summary `json:"summary,omitempty"`
	// RowCount 当前返回的数据行数。
	RowCount int64 `json:"rowCount,omitempty"`
	// TotalRowCount 	所有符合条件的数据总行数。
	TotalRowCount int64 `json:"totalRowCount,omitempty"`
}
