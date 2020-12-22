package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type cardDeck struct {
	cards     []int
	deckScore uint64
}

func (s *cardDeck) new() *cardDeck {
	s.cards = []int{}
	s.deckScore = 0
	return s
}

func (s *cardDeck) enqueue(t int) {
	s.cards = append(s.cards, t)
	for _, v := range s.cards {
		s.deckScore += uint64(v)
	}
}

func (s *cardDeck) dequeue() int {
	s.deckScore -= uint64(s.size()) * uint64(s.cards[0])
	cardTop := s.cards[0]
	s.cards = s.cards[1:len(s.cards)]
	return cardTop
}

func (s *cardDeck) front() int {
	tempCard := s.cards[0]
	return tempCard
}

func (s *cardDeck) isEmpty() bool {
	return len(s.cards) == 0
}

func (s *cardDeck) size() int {
	return len(s.cards)
}

func (s *cardDeck) copy(t int) *cardDeck {
	temp := new(cardDeck)
	for x := 0; x < t; x++ {
		temp.enqueue(s.cards[x])
	}
	return temp
}

func mapContainsUint64(dict map[uint64]int, s uint64) bool {
	if _, ok := dict[s]; ok {
		return ok
	}
	return false
}

func copyIntSlice(old []int) []int {
	temp := make([]int, len(old))
	copy(temp, old)
	return temp
}

func compIntSlice(x []int, y []int) bool {
	if len(x) != len(y) {
		return false
	} else {
		for i, v := range x {
			if v != y[i] {
				return false
			}
		}
	}
	return true
}

var game int = 1

func recursiveWar(player1 *cardDeck, player2 *cardDeck) (*cardDeck, int) {
	memo := [][][]int{}

	counter := 1
	for !player1.isEmpty() && !player2.isEmpty() {
		for _, oldDecks := range memo {
			old1, old2 := oldDecks[0], oldDecks[1]
			if compIntSlice(player1.cards, old1) || compIntSlice(player1.cards, old2) {
				return player1, 1
			}
		}

		memo = append(memo, [][]int{copyIntSlice(player1.cards), copyIntSlice(player2.cards)})

		card1 := player1.dequeue()
		card2 := player2.dequeue()
		if player1.size() >= int(card1) && player2.size() >= int(card2) {
			tempDeck1 := player1.copy(int(card1))
			tempDeck2 := player2.copy(int(card2))
			game++
			_, temp := recursiveWar(tempDeck1, tempDeck2)
			if temp == 1 {
				player1.enqueue(card1)
				player1.enqueue(card2)
			} else {
				player2.enqueue(card2)
				player2.enqueue(card1)
			}
		} else {
			if card1 > card2 {
				player1.enqueue(card1)
				player1.enqueue(card2)
			} else {
				player2.enqueue(card2)
				player2.enqueue(card1)
			}
		}
		counter++
	}
	if !player1.isEmpty() {
		return player1, 1
	} else {
		return player2, 2
	}
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
			player1deck.enqueue(temp)
		} else {
			temp, _ := strconv.Atoi(line)
			player2deck.enqueue(temp)
		}
	}

	acc := 0
	winningDeck, _ := recursiveWar(player1deck, player2deck)
	for x := winningDeck.size(); x > 0; x-- {
		temp := winningDeck.dequeue()
		acc += int(temp) * x
	}
	fmt.Println(acc)
}
