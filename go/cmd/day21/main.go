package main

import "fmt"

func main() {
	dd := NewDeterministicDice()

	roll(dd)
}

func roll(dice Dice) {
	player1 := NewPlayer(4)
	player2 := NewPlayer(3)

	for {
		player1.RollAndMove(dice)
		if player1.Score() >= 1000 {
			break
		}

		player2.RollAndMove(dice)
		if player2.Score() >= 1000 {
			break
		}
	}

	fmt.Printf("%s\n", player1)
	fmt.Printf("%s\n", player2)

	fmt.Printf("%d\n", dice.Rolls()*player1.Score())
	fmt.Printf("%d\n", dice.Rolls()*player2.Score())
}

type Dice interface {
	Roll() int
	Rolls() int
}

type DeterministicDice struct {
	count    int
	lastRoll int
}

func NewDeterministicDice() *DeterministicDice {
	return &DeterministicDice{count: 0, lastRoll: 0}
}

func (d *DeterministicDice) Roll() int {
	if d.lastRoll == 100 {
		d.lastRoll = 0
	}

	d.count++
	d.lastRoll++
	return d.lastRoll
}

func (d *DeterministicDice) Rolls() int {
	return d.count
}

type Player struct {
	currentSpace int
	points       int
}

func NewPlayer(startingPosition int) *Player {
	if startingPosition < 1 || startingPosition > 10 {
		panic(fmt.Sprintf("position must be between 1 and 10, not %d", startingPosition))
	}
	return &Player{
		currentSpace: startingPosition,
		points:       0,
	}
}

func (p *Player) RollAndMove(dice Dice) {
	rolls := 0
	rolls += dice.Roll()
	rolls += dice.Roll()
	rolls += dice.Roll()

	for i := 1; i <= rolls; i++ {
		if p.currentSpace == 10 {
			p.currentSpace = 1
		} else {
			p.currentSpace++
		}
	}

	p.points += p.currentSpace
}

func (p *Player) Score() int {
	return p.points
}
