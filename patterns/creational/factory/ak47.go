package factory

type Ak47 struct {
	Gun
}

func newAk47() GunInterface {
	return &Ak47{
		Gun: Gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}
