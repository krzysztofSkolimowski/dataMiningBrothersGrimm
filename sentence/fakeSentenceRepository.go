package sentence

type sentenceList struct {
	sentences []sentence
}

func NewSentenceList(s ...sentence) *sentenceList{
	return &sentenceList{s}
}

func (s *sentenceList)Sentences()  []sentence{
	return s.sentences
}

func (s *sentenceList)AddSentence(sentence sentence) []sentence{
	s.sentences = append(s.sentences, sentence)
	return s.sentences
}

