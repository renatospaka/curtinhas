package builder

type iBuilder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() house
}

func GetBuilder(builderType string) iBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}

	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}
