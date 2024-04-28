package enum

// Device 推广设备
type Device Enum[int]

func (e Device) Values() []Device {
	return []Device{
		{0, "计算机"},
		{1, "移动设备"},
	}
}

// UnmarshalJSON implement json Unmarshal interface
func (e *Device) UnmarshalJSON(b []byte) error {
	v, err := UnmarshalJSON[int](b)
	if err != nil {
		return err
	}
	*e = Device(v)
	return nil
}
