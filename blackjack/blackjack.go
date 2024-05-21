package blackjack

import (
	"math/rand"
)

type Card struct {
	Suit  string
	Name  string
	Value []uint8
}

type CardDeck struct {
	Cards []Card
}

var CardSuits = []string{"Clubs", "Diamonds", "Hearts", "Spades"}

var CardTypes = map[string][]uint8{
	"Two":   {2},
	"Three": {3},
	"Four":  {4},
	"Five":  {5},
	"Dix":   {6},
	"Seven": {7},
	"Eight": {8},
	"Nine":  {9},
	"Ten":   {10},
	"Jacl":  {10},
	"Queen": {10},
	"King":  {10},
	"Ace":   {10}, // {1, 10},
}

func (deck *CardDeck) generate() {
	for _, suit := range CardSuits {
		for key, value := range CardTypes {
			card := Card{
				Suit:  suit,
				Name:  key,
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
}

type Rules struct {
	MaxPoints   uint8
	TotalRounds uint8
}

type Round struct {
	Points   uint8
	CardHand []Card
	Deck     CardDeck
}

type Game struct {
	Rules  Rules
	Losses uint
	Wins   uint
	Rounds []Round
}

func NewRound() *Round {
	round := &Round{
		Points: 0,
		Deck:   CardDeck{},
	}
	round.Deck.Shuffle()
	return round
}

func (r *Round) DrawCard() Card {
	card := r.Deck.Cards[0]
	r.Deck.Cards = r.Deck.Cards[1:]
	return card
}

func NewGame() *Game {
	return &Game{
		Rules: Rules{
			MaxPoints:   21,
			TotalRounds: 6,
		},
	}
}
