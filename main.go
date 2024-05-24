package main

func main() {
	arsenal := newDeck()
	arsenal.shuffle()
	hand, remainingArsenal := deal(arsenal, 2)
	hand.print()
	remainingArsenal.print()
}
