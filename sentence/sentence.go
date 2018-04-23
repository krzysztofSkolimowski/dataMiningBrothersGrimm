package sentence

import "strings"

type sentence struct {
	words []string
}

func NewSentence(words ...string) *sentence {
	return &sentence{words}
}

func (s sentence) Word(i int) string {
	return s.words[i]
}

func ParseRealClause(s string) *sentence {
	words := strings.Replace(s, ",", "", -1)
	words = strings.Replace(words, ".", "", -1)
	return &sentence{strings.Split(words, " ")}
}
