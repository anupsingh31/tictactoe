package player

type Player struct {
	name, mark string
}

func New(playerName, pMark string) *Player {
	return &Player{
		name: playerName,
		mark: pMark,
	}
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetPlayerMark() string {
	return p.mark
}
