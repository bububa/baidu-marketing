package enum

// DeepConvertType 深度转化
type DeepConvertType Enum[int]

func (e DeepConvertType) Values() []DeepConvertType {
	return []DeepConvertType{
		{10, "商品支付成功"},
		{25, "注册"},
		{26, "付费"},
		{27, "客户自定义"},
		{28, "次日留存"},
		{42, "授信"},
		{45, "商品下单成功"},
		{53, "订单核对成功"},
		{54, "收货成功"},
	}
}

func (e *DeepConvertType) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = DeepConvertType(v)
	return nil
}
