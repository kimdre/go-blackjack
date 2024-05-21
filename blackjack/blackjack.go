package blackjack

import (
	"fmt"
	"math/rand"
)

type Card struct {
	Color string
	Name  string
	Value []uint8
}

type CardDeck struct {
	Cards []Card
}

var CardColors = []string{"Kreuz", "Pik", "Herz", "Karo"}

var CardTypes = map[string][]uint8{
	"zwei":   {2},
	"drei":   {3},
	"vier":   {4},
	"fünf":   {5},
	"sechs":  {6},
	"sieben": {7},
	"acht":   {8},
	"neun":   {9},
	"zehn":   {10},
	"bube":   {10},
	"dame":   {10},
	"könig":  {10},
	"ass":    {1, 10},
}

func (deck *CardDeck) generate() {
	for _, color := range CardColors {
		for key, value := range CardTypes {
			card := Card{
				Color: color,
				Name: key,
				Value: value,
			}
			deck.Cards = append(deck.Cards, card)
		}
	}

}

func (deck *CardDeck) Shuffle() {
	deck.generate()
	rand.Shuffle(len(deck.Cards), func(i, j int) {
		deck.Cards[i], deck.Cards[j] = deck.Cards[j], deck.Cards[i]
	})

	fmt.Println(deck.Cards)
}

type Rules struct {
	MaxPoints uint8
	Rounds uint8
}

type Round struct {
	Points uint8
	History []Card
	Deck CardDeck
}

type Game struct {
	Rules Rules 
	Round uint
	Losses uint
	Wins uint
	Rounds []Round
}

func (r *Round) NewRound() {
	r.Points = 0
	r.Deck.Shuffle()
}

func (r *Round) DrawCard() Card {
	card := r.Deck.Cards[0]
	r.Deck.Cards = r.Deck.Cards[1:]
	return card
}

func (g *Game) NewGame() {
	g.Rules.MaxPoints = 21
	g.Rules.Rounds = 6
	g.Losses = 0
	g.Wins = 0
}
