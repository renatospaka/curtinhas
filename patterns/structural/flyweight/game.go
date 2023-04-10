package flyweight

type game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *game {
	return &game{
			terrorists:        make([]*Player, 1),
			counterTerrorists: make([]*Player, 1),
	}
}

func (c *game) AddTerrorist(dressType string) {
	player := newPlayer("T", dressType)
	c.terrorists = append(c.terrorists, player)
	return
}

func (c *game) AddCounterTerrorist(dressType string) {
	player := newPlayer("CT", dressType)
	c.counterTerrorists = append(c.counterTerrorists, player)
	return
}