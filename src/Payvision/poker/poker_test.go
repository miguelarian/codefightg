package main

import "testing"

func Test_StraightFlush_vs_RoyalStraight(t *testing.T) {

	//
	expected := 1
	output := ComparePokerHands("2H 3H 4H 5H 6H", "KS AS TS QS JS")

	if expected != output {
		t.Errorf("Failed Test_StraightFlush_vs_RoyalStraight. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_StraightFlush_vs_AcePoker(t *testing.T) {

	expected := -1
	output := ComparePokerHands("2H 3H 4H 5H 6H", "AS AD AC AH JD")

	if expected != output {
		t.Errorf("Failed Test_StraightFlush_vs_AcePoker. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_AcePoker_vs_JokerPoker(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS AH 2H AD AC", "JS JD JC JH 3D")

	if expected != output {
		t.Errorf("Failed Test_AcePoker_vs_JokerPoker. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_TwoPairs_vs_ThreeKind(t *testing.T) {

	expected := 1
	output := ComparePokerHands("AS AH 2H 3D 3C", "JS JD JC KH 3D")

	if expected != output {
		t.Errorf("Failed Test_TwoPairs_vs_ThreeKind. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_HighCard(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS 6H 2H 4D 3C", "3S 4D 6C JH QD")

	if expected != output {
		t.Errorf("Failed Test_HighCard. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_OnePairDraw(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS AH 2H 4D 3C", "KS KD AC JH QD")

	if expected != output {
		t.Errorf("Failed Test_OnePairDraw. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_OnePairDraw_And_HighestCard_1(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS AH 2H 5D 3C", "KS KD 3C JH 5H")

	if expected != output {
		t.Errorf("Failed Test_OnePairDraw. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_OnePairDraw_And_HighestCard_2(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS AH 8H 2D 3C", "AC AD 4C 5H 6H")

	if expected != output {
		t.Errorf("Failed Test_OnePairDraw. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_HighestCardDraw(t *testing.T) {

	expected := 1
	output := ComparePokerHands("AS JH 2H 6D 3C", "AC KD 3S 5H QD")

	if expected != output {
		t.Errorf("Failed Test_HighestCardDraw. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Draw_SameHands(t *testing.T) {

	expected := 0
	output := ComparePokerHands("AS JH 2H 6D 3C", "AS JH 2H 6D 3C")

	if expected != output {
		t.Errorf("Failed Test_OnePairDraw. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Fullhouse_Wins(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS AH AH 6D 6C", "AS AH AH 3D 4C")

	if expected != output {
		t.Errorf("Failed Test_Fullhouse_Wins. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Draw_Fullhouse_1(t *testing.T) {

	expected := 0
	output := ComparePokerHands("AS AH AH 6D 6C", "AS AH AH 6D 6C")

	if expected != output {
		t.Errorf("Failed Test_Draw_Fullhouse_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Draw_Fullhouse_2(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AS AH AH 6D 6C", "AS AH AH 3D 3C")

	if expected != output {
		t.Errorf("Failed Test_Draw_Fullhouse_2. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Draw_Fullhouse_3(t *testing.T) {

	expected := 1
	output := ComparePokerHands("QS QH QH 6D 6C", "AS AH AH 3D 3C")

	if expected != output {
		t.Errorf("Failed Test_Draw_Fullhouse_3. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Poker_1(t *testing.T) {

	expected := 1
	output := ComparePokerHands("3S 3H 3H 3D 6C", "4S 4H 4H 4D 3C")

	if expected != output {
		t.Errorf("Failed Test_Poker_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Poker_Draw_1(t *testing.T) {

	expected := 0
	output := ComparePokerHands("4S 4H 4H 4D 3C", "4S 4H 4H 4D 3C")

	if expected != output {
		t.Errorf("Failed Test_Poker_Draw_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Poker_Draw_2(t *testing.T) {

	expected := -1
	output := ComparePokerHands("4S 4H 4H 4D 5C", "4S 4H 4H 4D 3C")

	if expected != output {
		t.Errorf("Failed Test_Poker_Draw_2. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_TwoPairs_Draw_1(t *testing.T) {

	expected := 1
	output := ComparePokerHands("4S 4H 3H 3D 5C", "4S 4H 3H 3D 6C")

	if expected != output {
		t.Errorf("Failed Test_TwoPairs_Draw_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_TwoPairs_Draw_2(t *testing.T) {

	expected := -1
	output := ComparePokerHands("5S 5H 3H 3D 5C", "4S 4H 3H 3D 6C")

	if expected != output {
		t.Errorf("Failed Test_TwoPairs_Draw_2. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_TwoPairs_Draw_3(t *testing.T) {

	expected := -1
	output := ComparePokerHands("5S 5H 3H 3D 5C", "5S 5H 2H 2D 6C")

	if expected != output {
		t.Errorf("Failed Test_TwoPairs_Draw_3. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Straight_Draw_1(t *testing.T) {

	expected := 1
	output := ComparePokerHands("2S 3H 4H 5D 6C", "4S 5H 6H 7D 8C")

	if expected != output {
		t.Errorf("Failed Test_Straight_Draw_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Straight_Draw_2(t *testing.T) {

	expected := 1
	output := ComparePokerHands("QS KH AH 2D 3C", "4S 5H 6H 7D 8C")

	if expected != output {
		t.Errorf("Failed Test_Straight_Draw_2. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_StraightFlush_Draw_1(t *testing.T) {

	expected := 1
	output := ComparePokerHands("2H 3H 4H 5H 6H", "4H 5H 6H 7H 8H")

	if expected != output {
		t.Errorf("Failed Test_StraightFlush_Draw_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_StraightFlush_Win_1(t *testing.T) {

	expected := -1
	output := ComparePokerHands("2H 3H 4H 5H 6H", "4H 5Q 6H 7H 8H")

	if expected != output {
		t.Errorf("Failed Test_StraightFlush_Draw_1. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Flush_Win(t *testing.T) {

	expected := -1
	output := ComparePokerHands("QH KH AH 2H 3H", "4S 5H 6H 7D 8C")

	if expected != output {
		t.Errorf("Failed Test_Flush_Win. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_Flush_Draw(t *testing.T) {

	expected := -1
	output := ComparePokerHands("QH KH AH 2H 3H", "QH 2H 4H 2H 3H")

	if expected != output {
		t.Errorf("Failed Test_Flush_Draw. \n Expected: %v \n Output: %v ", expected, output)
	}
}

func Test_HighCard_Draw(t *testing.T) {

	expected := -1
	output := ComparePokerHands("AH KH QH JH 8S", "AH KH QH JH 6S")

	if expected != output {
		t.Errorf("Failed Test_HighCard_Draw. \n Expected: %v \n Output: %v ", expected, output)
	}
}
