package factory

type musket struct {
	Gun
}

func newMusket() GunInterface {
	return &musket{
		Gun: Gun{
			name:  "Musket gun",
			power: 1,
		},
	}
}
