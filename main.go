package main

import (
	"fmt"
	"time"

	"github.com/kimdre/go-blackjack/blackjack"
	"github.com/kimdre/go-blackjack/utils"
)

func main() {
	game := blackjack.NewGame()

	for {
		player := blackjack.NewRound()
		dealer := blackjack.NewRound()

		var roundEnd = false
		var playerRound = true

		for i := 0; i < 2; i++ {
			dealerCard := player.DrawCard()
			dealer.Points += dealerCard.Value[0]
			dealer.CardHand = append(dealer.CardHand, dealerCard)

			card := player.DrawCard()
			player.Points += card.Value[0]
			player.CardHand = append(player.CardHand, card)
		}

		for {
			utils.ClearTerminal()
			fmt.Printf("Round # %v of %v\nCards left in Deck: %v\nWins: %v, Losses: %v\n\n", len(game.Rounds)+1, game.Rules.TotalRounds, len(player.Deck.Cards), game.Wins, game.Losses)
			if playerRound {
				fmt.Println("Dealer cards:", dealer.CardHand[0], "???")
				fmt.Println("Dealer points:", dealer.CardHand[0].Value)
			} else {
				fmt.Println("Dealer cards:", dealer.CardHand)
				fmt.Println("Dealer points:", dealer.Points)
			}
			fmt.Println()
			fmt.Println("Your cards:", player.CardHand)
			fmt.Println("Your points:", player.Points)
			fmt.Println()

			if player.Points <= game.Rules.MaxPoints && dealer.Points > game.Rules.MaxPoints {
				fmt.Println("You win.")
				fmt.Println()
				game.Wins += 1
				roundEnd = true
			} else if player.Points >= game.Rules.MaxPoints && dealer.Points >= game.Rules.MaxPoints {
				fmt.Println("Draw.")
				fmt.Println()
				roundEnd = true
			} else if player.Points > game.Rules.MaxPoints && dealer.Points <= game.Rules.MaxPoints {
				fmt.Println("You loose.")
				fmt.Println()
				game.Losses += 1
				roundEnd = true
			} else if player.Points < dealer.Points && dealer.Points <= game.Rules.MaxPoints && !playerRound {
				fmt.Println("You loose.")
				fmt.Println()
				game.Losses += 1
				roundEnd = true
			}

			if !roundEnd && !playerRound {
				fmt.Println("Dealer draws new card...")
				time.Sleep(1 * time.Second)
				dealerCard := player.DrawCard()
				dealer.Points += dealerCard.Value[0]
				dealer.CardHand = append(dealer.CardHand, dealerCard)
			}

			if !roundEnd && playerRound {
				var choice string
				fmt.Print("Draw another card? [y/n] ")
				fmt.Scan(&choice)

				switch choice {
				case "y":
					card := player.DrawCard()
					player.Points += card.Value[0]
					player.CardHand = append(player.CardHand, card)
				case "n":
					playerRound = false
				default:
					return
				}
			}

			if roundEnd {
				if len(game.Rounds)+1 >= int(game.Rules.TotalRounds) {
					fmt.Println("You finished the game")
					fmt.Printf("Wins: %v, Losses: %v\nTotal Rounds: %v\n", game.Wins, game.Losses, len(game.Rounds)+1)
					return
				}

				var newRound string
				fmt.Print("New round? [y/n] ")
				fmt.Scan(&newRound)

				if newRound != "y" {
					return
				}
				break
			}
		}

		game.Rounds = append(game.Rounds, *player)
	}
}
