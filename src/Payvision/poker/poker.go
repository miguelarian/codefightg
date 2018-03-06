import (
	"math"
	"sort"
	"strings"
)
// *********** CONSTANTS ***********
// All Hands and values
const (
	DEFAULT_HAND int = 1 + iota // 0
	HIGH_CARD
	ONE_PAIR
	TWO_PAIR
	THREE_KIND
	STRAIGHT
	FLUSH
	FULL_HOUSE
	FOUR_KIND
	STRAIGHT_FLUSH
	ROYAL_STRAIGHT
)

// *********** STRUCTS ***********

// Card struct
type Card struct {
	value int
	suit  string
}
type ByCardValue []Card

func (items ByCardValue) Len() int      { return len(items) }
func (items ByCardValue) Swap(i, j int) { items[i], items[j] = items[j], items[i] }
func (items ByCardValue) Less(i, j int) bool {
	if items[i].value < items[j].value {
		return true
	}
	return false
}

// Hand struct
type Hand struct {
	cards              []Card
	isFlush            bool
	isStraight         bool
	isRoyal            bool
	highestCard        int
	handValue          int
	handBestCard       int
	handSecondBestCard int
}

// *********** HELPER FUNCTIONS ***********

func parseHand(value string) Hand {

	var cards []Card
	inputCards := strings.Split(value, " ")
	highestCard := 0.0

	// Some hands can be detected while parsing cards: color, straights...
	cardsSuit := string(inputCards[0][1])
	isFlush := true
	for _, inputCard := range inputCards {
		cardValue := parseCardValue(string(inputCard[0]))
		card := Card{cardValue, string(inputCard[1])}
		cards = append(cards, card)

		// Get highest card on the fly
		highestCard = math.Max(highestCard, float64(cardValue))

		// Check for flush
		if card.suit != cardsSuit {
			isFlush = false
		}
	}

	// Sort cards to help with straights
	sort.Sort(ByCardValue(cards))

	// Check for straight
	isStraight := math.Abs(float64(cards[0].value)-float64(cards[4].value)) == 4
	isRoyal := isStraight && cards[0].value == 10
	return Hand{cards: cards,
		highestCard: int(highestCard),
		isFlush:     isFlush,
		isStraight:  isStraight,
		isRoyal:     isRoyal,
		handValue:   DEFAULT_HAND}
}

func parseCardValue(value string) int {
	switch value {
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	default:
		res, _ := strconv.Atoi(value)
		return res
	}
}

// Checks if a card is repeated 'expectedCount' times: pairs, threes, pokers...
// Returns (isDuplicate, dusplicateSetsCount, bestHandCard (highest set)
func checkDuplicates(hand Hand, expectedCount int) (bool, int, int) {

	setsCount := 0
	handBestCard := 0
	setValues := ""

	// Iterate counting card ocurrences
	for _, targetCard := range hand.cards {
		counter := 0
		for _, card := range hand.cards {
			if targetCard.value == card.value {
				counter++
			}
		}
		if counter == expectedCount && !strings.Contains(setValues, string(targetCard.value)) {
			setsCount++
			handBestCard = targetCard.value
			setValues += string(targetCard.value) + "."
		}
	}
	return setsCount > 0, setsCount, handBestCard
}

func compareCardValues(value1, value2 int) int {
	if value1 > value2 {
		// Hand1 wins
		return -1
	} else if value1 < value2 {
		// Hand2 wins
		return 1
	}
	// Tie
	return 0
}

// ***********  HAND DETECTION FUNCTIONS ***********
type HandAnalyser func(hand Hand, channel chan Hand)

var isRoyalStraight = func(hand Hand, channel chan Hand) {
	if hand.isRoyal {
		hand.handValue = ROYAL_STRAIGHT
		hand.handBestCard = hand.highestCard
	}
	channel <- hand
}

var isStraightFlush = func(hand Hand, channel chan Hand) {
	if hand.isStraight && hand.isFlush {
		hand.handValue = STRAIGHT_FLUSH
		hand.handBestCard = hand.highestCard
	}
	channel <- hand
}

var isFourKind = func(hand Hand, channel chan Hand) {
	isFourKind, _, handBestCard := checkDuplicates(hand, 4)
	if isFourKind {
		hand.handValue = FOUR_KIND
		hand.handBestCard = handBestCard
	}
	channel <- hand
}

var isFullHouse = func(hand Hand, channel chan Hand) {
	isThreeKind, _, handBestCard := checkDuplicates(hand, 3)
	isPair, _, handSecondBestCard := checkDuplicates(hand, 2)
	if isThreeKind && isPair {
		hand.handValue = FULL_HOUSE
		hand.handBestCard = handBestCard
		hand.handSecondBestCard = handSecondBestCard
	}
	channel <- hand
}

var isFlush = func(hand Hand, channel chan Hand) {
	if hand.isFlush {
		hand.handValue = FLUSH
		hand.handBestCard = hand.highestCard
	}
	channel <- hand
}

var isStraight = func(hand Hand, channel chan Hand) {
	if hand.isStraight {
		hand.handValue = STRAIGHT
		hand.handBestCard = hand.highestCard
	}
	channel <- hand
}

var isThreeKind = func(hand Hand, channel chan Hand) {
	isThreeKind, _, handBestCard := checkDuplicates(hand, 3)
	if isThreeKind {
		hand.handValue = THREE_KIND
		hand.handBestCard = handBestCard
	}
	channel <- hand
}

var isTwoPair = func(hand Hand, channel chan Hand) {
	isTwoPair, pairsCount, handBestCard := checkDuplicates(hand, 2)
	if isTwoPair && pairsCount == 2 {
		hand.handValue = TWO_PAIR
		hand.handBestCard = handBestCard
	}
	channel <- hand
}

var isOnePair = func(hand Hand, channel chan Hand) {
	isOnePair, _, handBestCard := checkDuplicates(hand, 2)
	if isOnePair {
		hand.handValue = ONE_PAIR
		hand.handBestCard = handBestCard
	}

	channel <- hand
}

// This goroutine spawns threads to check all possible hands for a given hand
// Returns best hand
func processHand(hand Hand, channel chan Hand, analysers []HandAnalyser) {

	// Create gouroutines communication channels
	var channels []chan Hand
	for _, analyser := range analysers {
		returnChannel := make(chan Hand)
		channels = append(channels, returnChannel)
		go analyser(hand, returnChannel)
	}

	// Get messages back from goroutines
	var allResults []Hand
	for _, subchannel := range channels {
		result := <-subchannel
		allResults = append(allResults, result)
	}

	// Return best hand to parent go routine
	channel <- selectBestHand(allResults)
}

// Gets all hands processed and returns the winner
func selectBestHand(hands []Hand) Hand {

	winner := hands[0]
	for _, hand := range hands {
		if hand.handValue > winner.handValue {
			winner = hand
		}
	}
	if winner.handValue == 0 {
		winner.handValue = HIGH_CARD
	}

	return winner
}

func compare(hand1, hand2 Hand) int {

	// ********* No ties ****************
	result := compareCardValues(hand1.handValue, hand2.handValue)
	if result != 0 {
		return result
	}

	// ********* Resolve ties ***********
	// Check hand best card (valid for all ties)
	result = compareCardValues(hand1.handBestCard, hand2.handBestCard)
	if result != 0 {
		return result
	}

	// FullHouse tie
	if hand1.handValue == FULL_HOUSE {
		result = compareCardValues(hand1.handSecondBestCard, hand2.handSecondBestCard)
		if result != 0 {
			return result
		}

		return 0
	}

	// One pair tie
	if hand1.handValue == ONE_PAIR {
		// Check rest of cards
		for i := 2; i >= 0; i-- {
			result = compareCardValues(hand1.cards[i].value, hand2.cards[i].value)
			if result != 0 {
				return result
			}
		}
	}

	// Highest Card
	for i := 4; i >= 0; i-- {
		result = compareCardValues(hand1.cards[i].value, hand2.cards[i].value)
		if result != 0 {
			return result
		}
	}

	// Tie
	return 0
}

// ***********  MAIN ***********
func ComparePokerHands(fistHand string, secondHand string) int {
	// Array with all anaylser functions (list of delegates)
	var handAnalysers = []HandAnalyser{
		isOnePair,
		isTwoPair,
		isThreeKind,
		isFourKind,
		isStraight,
		isFlush,
		isFullHouse,
		isStraightFlush,
		isRoyalStraight}

	hand1Channel := make(chan Hand)
	hand2Channel := make(chan Hand)

	// Spawn two goroutine threads to process hands simultaneously
	go processHand(parseHand(fistHand), hand1Channel, handAnalysers)
	go processHand(parseHand(secondHand), hand2Channel, handAnalysers)

	hand1 := <-hand1Channel
	hand2 := <-hand2Channel

	return compare(hand1, hand2)
}



func Sort(hands []string) []string {

}
