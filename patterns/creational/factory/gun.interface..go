package factory

type GunInterface interface {
	SetName(name string)
	SetPower(power int)
	GetName() string
	GetPower() int
}
