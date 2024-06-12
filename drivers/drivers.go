package drivers

type HWDevice interface {
	Test()
}

type DriversList = []string

type DeviceDriver interface {
	Name() string
	Init() HWDevice
	Close() error
}
