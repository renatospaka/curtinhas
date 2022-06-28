package builder

type normalBuilder struct {
	WindowType string
	DoorType   string
	Floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) SetWindowType() {
	b.WindowType = "Wooden Window"
}

func (b *normalBuilder) SetDoorType() {
	b.DoorType = "Wooden Door"
}

func (b *normalBuilder) SetNumFloor() {
	b.Floor = 2
}

func (b *normalBuilder) GetHouse() house {
	return house{
		DoorType:   b.DoorType,
		WindowType: b.WindowType,
		Floor:      b.Floor,
	}
}
