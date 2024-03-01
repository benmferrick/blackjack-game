package main

import (
    "fmt"
    "math/rand"
    "time"
    "bufio"
    "os"
    "strings"
)

type Suit string
const (
    Clubs    Suit = "Clubs"
    Diamonds Suit = "Diamonds"
    Hearts   Suit = "Hearts"
    Spades   Suit = "Spades"
)

type Result string
const (
    WinNatural Result = "WinNatural"
    WinBetterHand Result = "WinBetterHand"
    WinDealerBust Result = "WinDealerBust"
    LoseBust Result = "LostBust"
    LoseWorseHand Result = "LoseWorseHand"
    LoseDealerNatural Result = "LoseDealerNatural"
    Push Result = "Push"
    SplitBothWin Result = "SplitBothWin"
    SplitBothLose Result = "SplitBothLose"
    SplitPush Result = "SplitPush"
    WaitingOtherHand Result = "WaitingOtherHand"
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

// askForHit asks the user if they want to hit and returns true if they do.
func askForHit() bool {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Do you want to hit? (yes/no): ")
    input, _ := reader.ReadString('\n') // Read the input until the newline character
    input = strings.TrimSpace(input) // Trim whitespace and newline characters

    // Check the input and return true if the user wants to hit
    return strings.ToLower(input) == "yes"
}

func askForSplit() bool {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Do you want to split? (yes/no): ")
    input, _ := reader.ReadString('\n') // Read the input until the newline character
    input = strings.TrimSpace(input) // Trim whitespace and newline characters

    // Check the input and return true if the user wants to hit
    return strings.ToLower(input) == "yes"
}

//runs one hand
func doAHand(playerCards []Card, dealerCards []Card) Result {

    fmt.Println("Player's cards:", playerCards)
    fmt.Println("Player Total:", calcTotal(playerCards))
    fmt.Println("Dealer's cards:", dealerCards)
    fmt.Println("Dealer Total:", calcTotal(dealerCards))

    //handles dealer getting 21 with first two cards
    if calcTotal(dealerCards) == 21 {
        return LoseDealerNatural
    }

    //if dealer does not have 21 and user does, automatic win for user
    if calcTotal(playerCards) == 21 {
        return WinNatural
    }

    if (isSplit(playerCards)) {
        if askForSplit() {
            
        }
    }

    //handles user turn to hit cards
	for {
        if askForHit() {
            fmt.Println("User chose to hit.")
            hit(&playerCards)
            fmt.Println("Player's cards:", playerCards)
            fmt.Println("Player Total:", calcTotal(playerCards))

            if (calcTotal(playerCards) > 21) {
                return LoseBust
            }

        } else {
            fmt.Println("User chose not to hit.")
            break // Exit the loop if the user chooses not to hit
        }
    }

    //handles dealer hits if <=16 or soft 17
    for calcTotal(dealerCards) <= 16 || isSoft17(dealerCards) {
        hit(&dealerCards)
        fmt.Println("Dealer's cards:", dealerCards)
        fmt.Println("Dealer Total:", calcTotal(dealerCards))
    }

    //handles tie
    if calcTotal(dealerCards) == calcTotal(playerCards) {
        return Push
    }

    //returns win or loss if neither bust
    if calcTotal(playerCards) > calcTotal(dealerCards) {
        return WinBetterHand
    }
    return LoseWorseHand
}

// calcTotal calculates the total value of a slice of cards
// where Ace is worth 11, and Jack, Queen, King are worth 10 each.
func calcTotal(cards []Card) int {
    aceCount := 0
    total := 0
    for _, card := range cards {
        switch card.rank {
        case Ace:
            total += 11
            aceCount += 1
        case Jack, Queen, King:
            total += 10
        case Two:
            total += 2
        case Three:
            total += 3
        case Four:
            total += 4
        case Five:
            total += 5
        case Six:
            total += 6
        case Seven:
            total += 7
        case Eight:
            total += 8
        case Nine:
            total += 9
        case Ten:
            total += 10
        }
    }

    if (total > 21) {
        for aceCount > 0 {
            total -= 10
            aceCount -= 1
        }
    }

    return total
}

func isSoft17(cards []Card) bool {
    aceCount := 0
    total := 0
    for _, card := range cards {
        switch card.rank {
        case Ace:
            total += 11
            aceCount += 1
        case Jack, Queen, King:
            total += 10
        case Two:
            total += 2
        case Three:
            total += 3
        case Four:
            total += 4
        case Five:
            total += 5
        case Six:
            total += 6
        case Seven:
            total += 7
        case Eight:
            total += 8
        case Nine:
            total += 9
        case Ten:
            total += 10
        }
    }

    if (total == 17 && aceCount == 1) {
        return true
    }

    return false
}

func isSplit(playerCards []Card) bool {
    if playerCards[0].rank == playerCards[1].rank {
        return true
    }

    return false
}

func main() {

    //get a commmit in today
    
    playerCards, dealerCards := deal()
    println(doAHand(playerCards, dealerCards))

}