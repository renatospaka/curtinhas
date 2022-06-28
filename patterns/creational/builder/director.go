package builder

type Director struct {
	builder iBuilder
}

func NewDirector(b iBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) SetBuilder(b iBuilder) {
	d.builder = b
}

func (d *Director) BuildHouse() house {
	d.builder.SetDoorType()
	d.builder.SetWindowType()
	d.builder.SetNumFloor()
	return d.builder.GetHouse()
}
