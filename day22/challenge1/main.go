package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type card int

type cardDeck struct {
	cards []card
}

func (s *cardDeck) new() *cardDeck {
	s.cards = []card{}
	return s
}

func (s *cardDeck) enqueue(t card) {
	s.cards = append(s.cards, t)
}

func (s *cardDeck) dequeue() card {
	cardTop := s.cards[0]
	s.cards = s.cards[1:len(s.cards)]
	return cardTop
}

func (s *cardDeck) front() card {
	tempCard := s.cards[0]
	return tempCard
}

func (s *cardDeck) isEmpty() bool {
	return len(s.cards) == 0
}

func (s *cardDeck) size() int {
	return len(s.cards)
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatalln("Failed to open file", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var player1deck *cardDeck = new(cardDeck)
	var player2deck *cardDeck = new(cardDeck)

	newline := false

	for scanner.Scan() {
		var line string = scanner.Text()
		if line == "Player 1:" || line == "Player 2:" {
			continue
		} else if len(line) == 0 {
			newline = true
			continue
		}
		if !newline {
			temp, _ := strconv.Atoi(line)
			player1deck.enqueue(card(temp))
		} else {
			temp, _ := strconv.Atoi(line)
			player2deck.enqueue(card(temp))
		}
	}

	// acc := 1
	for !player1deck.isEmpty() && !player2deck.isEmpty() {
		// fmt.Printf("-- Round %d --\n", acc)
		// fmt.Print("Player 1's deck: ")
		// fmt.Println(player1deck.cards)
		// fmt.Print("Player 2's deck: ")
		// fmt.Println(player2deck.cards)
		card1 := player1deck.dequeue()
		// fmt.Printf("Player 1 plays: %d\n", card1)
		card2 := player2deck.dequeue()
		// fmt.Printf("Player 2 plays: %d\n", card2)
		if card1 > card2 {
			// fmt.Println("Player 1 wins the round!")
			player1deck.enqueue(card1)
			player1deck.enqueue(card2)
		} else {
			// fmt.Println("Player 2 wins the round!")
			player2deck.enqueue(card2)
			player2deck.enqueue(card1)
		}
		//acc++
	}
	acc := 0
	if !player1deck.isEmpty() {
		// fmt.Println("Player 1 won!")
		// fmt.Println(player1deck.cards)
		for x := player1deck.size(); x > 0; x-- {
			temp := player1deck.dequeue()
			acc += int(temp) * x
		}
	} else {
		// fmt.Println("Player 2 won!")
		// fmt.Println(player2deck.cards)
		for x := player2deck.size(); x > 0; x-- {
			temp := player2deck.dequeue()
			acc += int(temp) * x
		}
	}
	fmt.Printf("The winning score is: %d\n", acc)
}
