package main

// importing mutiple package, packageName/subpackageName

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type 'deck' which is slice of strings
// type deck extends the behaviour of slice of strings
type deck []string

// we don't need a receier as we are not calling it from any obj, we are calling this
// func to create a new Deck obj
func newDeck() deck {
	// declare a cards slice
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	// return the cards slice
	return cards

}

// function to print
// here, d deck is a receiver

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// we take a deck and split it into two as handsize
// we take a starting deck and split it into two different slice using handsize
// we can return multiple return values from one function
// we decide to not include receiver here

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// we can definately tur a slice of deck to a slice of byte but we decided to cnv to string first
//func signature:- 'func' kw + receiver type + funcName(arg list) + return type { logic code here }

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to file func which allows us to write on files in HDD, if error occurs; it returns error
// perm : 0666 default permission, anyone read this file

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// read the file from HDD: reverse process of writing and conv
// we will read slice of byte and conv it to slice of deck

func newDeckFromFile(filename string) deck {
	// byte slice and error obj
	// if something is wrong in reading flie, error will have a value; else it will be nil(no value)

	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// option #1 - log the error and return a call to newDeck()
		// option #2 - log the error and entirely quit the program
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	s := strings.Split(string(bs),",") //string(bs)=> Ace of Spacdes,Two of Spades,Thee of Spades
	return deck(s) // to get slice of deck from slice of string


}

// func shuffle, we need receiver so we can call the func like: cards.shuffle()
// we need not always get a ref to element we are iterating over
// in this case we just care about the index

func (d deck) shuffle() {
	// lets create a custom seed for creating randomness
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d)-1)
		// one lne swap
		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
