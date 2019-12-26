package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// any value of type deck
// can use print() method.
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// join by commma all strings from d which is deck type
	// slice of strings and make one big string we return.
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// like above, we are converting to byte slice return value
	// of toString method, which return one large string.
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// convert stuff back to slice of strings
	// return deck type after econversion
	s := strings.Split(string(bs), ",")

	return deck(s)
}

func (d deck) shuffle() {
	// seed is current time in int64, we use that to
	// generate NewSource object and this is base 
	// for rng 
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
