package builder

type iglooBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) SetWindowType() {
	b.WindowType = "Snow Window"
}

func (b *iglooBuilder) SetDoorType() {
	b.DoorType = "Snow Door"
}

func (b *iglooBuilder) SetNumFloor() {
	b.Floor = 1
}

func (b *iglooBuilder) GetHouse() house {
	return house{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
