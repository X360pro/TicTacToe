package service

import (
	"components"
	"errors"
)

type GameService struct {
	*ResultService
	player [2]*components.Player
}

func NewGameService(pl [2]*components.Player, size int) *GameService {
	return &GameService{NewResultService(NewBoardService(size)), pl}
}

var turn = 0

func (g *GameService) Play(pos uint8) (error, string) {
	if pos < 0 || pos > g.Size*g.Size-1 {
		return errors.New("Positoin is out of bounds"), "nil"
	}
	var res string
	if turn%2 == 0 {
		err := g.PutMarkInPosition(pos, g.player[0].Mark)
		if err != nil {
			return err, "nil"
		}
		res = g.GiveResult(g.player[0])
	} else {
		err := g.PutMarkInPosition(pos, g.player[1].Mark)
		if err != nil {
			return err, "nil"
		}
		res = g.GiveResult(g.player[1])
	}
	turn++
	return nil, res

}
