package main

import "fmt"

func printPointer(namePointer *string) {
	fmt.Println(namePointer)
}

func main() {

	name := "bill"

	namePointer := &name

	printPointer(namePointer)

	//card := "newvar"
	//fmt.Println(card)
	//// var card string = "Ace of spades"
	//// card := "Ace of spades"
	//// card = "five"
	//
	//// printState()
	//var cards = []string{"card1", "card2"}
	////var deck = deck{newCard(), newCard()}
	////deck = append(deck, "six")
	////
	////deck.print()
	//var newCards = append(cards, "newCard")
	//fmt.Println(newCards)
	//
	//for i, s := range newCards {
	//	fmt.Println(i, s)
	//}

	//greeting := "Hello"
	//fmt.Println([]byte(greeting))

	//var cards = newDeckFromFile("deck")
	//cards.shuffle()
	//cards.print()

	//
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()

	//alex := person{firstname: "Alex", lastname: "Anderson"}
	//fmt.Println(alex)

	//var osama = person{
	//	firstname: "osama",
	//	lastname:  "mohamed",
	//	contactInfo: contactInfo{
	//		email:   "osama.mn@live.com",
	//		zipCode: 1234,
	//	},
	//}
	//
	//segment := "segment"
	//updateSegmentName(segment)
	//fmt.Println(segment)
	//
	////osamaPointer := &osama
	//osama.updateName("Ahmed")
	//osama.print()
}

func updateSegmentName(s string) {
	s = string(s[0]) + s[1:] + "2"
}
func (p *person) updateName(newFirstName string) {
	(*p).firstname = newFirstName
}

func (p person) print() {
	fmt.Println(p.firstname)
}

func getAceOfSpades() string {
	return "Ace of spades"
}

func newCard() string {
	return "Five"
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}
