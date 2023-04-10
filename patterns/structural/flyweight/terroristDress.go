package flyweight

type TerroristDress struct {
	color string
}

func (t *TerroristDress) GetColor() string {
	return t.color
}

func newTerroristDress() *TerroristDress {
	return &TerroristDress{color: "red"}
}
