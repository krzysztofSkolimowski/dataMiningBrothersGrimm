package main

import (
	"fmt"
	"dataMiningBrothersGrim/sentence"
)

func main() {

	fmt.Println("Hello Brothers Grimm")
	someText := *sentence.NewSentence("Hello", "Bro", "I", "am")
	anotherText := *sentence.NewSentence("Hello", "Bro", "I", "am", "who", "I", "am")
	realClauseTest := *sentence.ParseRealClause("Hmm, co by, tutaj, napisac.")

	sl := sentence.NewSentenceList(someText, anotherText, realClauseTest)
	fmt.Println(sl.Sentences())
}
