package main

import (
    "fmt"
    "math/rand"
    "time"
)

type Suit string
const (
    Clubs    Suit = "Clubs"
    Diamonds Suit = "Diamonds"
    Hearts   Suit = "Hearts"
    Spades   Suit = "Spades"
)


type Rank string
const (
    Ace   Rank = "Ace"
    Two   Rank = "2"
    Three Rank = "3"
    Four  Rank = "4"
    Five  Rank = "5"
    Six   Rank = "6"
    Seven Rank = "7"
    Eight Rank = "8"
    Nine  Rank = "9"
    Ten   Rank = "10"
    Jack  Rank = "Jack"
    Queen Rank = "Queen"
    King  Rank = "King"
)

type Card struct {
    suit Suit
    rank Rank
}

var suits = []Suit{Clubs, Diamonds, Hearts, Spades}
var ranks = []Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

// NewCard creates a new card with the given suit and rank
func NewCard(suit Suit, rank Rank) Card {
    return Card{suit: suit, rank: rank}
}

// deals starter cards, **constant shuffle**
func deal() (playerCards []Card, dealerCards []Card) {
    rand.Seed(time.Now().UnixNano())
    for i := 0; i < 2; i++ {
        playerCards = append(playerCards, NewCard(suits[rand.Intn(len(suits))], ranks[rand.Intn(len(ranks))]))
        dealerCards = append(dealerCards, NewCard(suits[rand.Intn(len(suits))], ranks[rand.Intn(len(ranks))]))
    }
    return playerCards, dealerCards
}

// hit function to add a random card to the player's hand
func hit(hand *[]Card) {
    rand.Seed(time.Now().UnixNano())
    *hand = append(*hand, NewCard(suits[rand.Intn(len(suits))], ranks[rand.Intn(len(ranks))]))
}

func main() {
    playerCards, dealerCards := deal()
    fmt.Println("Player's cards:", playerCards)
    fmt.Println("Dealer's cards:", dealerCards)
	hit(&playerCards)
	fmt.Println("Player's cards:", playerCards)
}