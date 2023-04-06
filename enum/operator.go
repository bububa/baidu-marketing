package enum

// Operator 支持：GT GTE LT LTE EQ NOT_EQ IN NOT_IN
type Operator string

const (
	GT     Operator = "GT"
	GTE             = "GTE"
	LT              = "LT"
	LTE             = "LTE"
	EQ              = "EQ"
	NOT_EQ          = "NOT_EQ"
	IN              = "IN"
	NOT_IN          = "NOT_IN"
)
