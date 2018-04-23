package sentence

import (
	"os"
	"log"
	"io/ioutil"
	"strings"
)

func ReadContent(path string) sentenceList {
	contentBytes := loadTextFile(path)
	content := contentFromBytes(contentBytes)
	contentSpliced := splitAndClearContent(content)
	return createSentenceListFromContent(contentSpliced)
}

func createSentenceListFromContent(content []string) sentenceList {
	sentenceList := *NewSentenceList()
	for _, s := range content {
		ss := ParseRealClause(s)
		sentenceList.AddSentence(*ss)
	}
	return sentenceList
}

func loadTextFile(path string) []byte{
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	contentBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	return contentBytes
}

func splitAndClearContent(content string) []string{
	contentSpliced := strings.Split(content, ".")
	for i, sentence := range contentSpliced {
		if sentence == "" {
			contentSpliced = append(contentSpliced[:i-1], contentSpliced[i+1:]...)
		}
	}
	return contentSpliced
}

func contentFromBytes(contentB []byte) string {
	content := string(contentB)
	cases := []struct {
		old string
		new string
	}{
		{"' ", ". "},
		{" '", " "},
		{"..", "."},
		{";", "."},
		{"?", "."},
		{"!", "."},
		{":", "."},
		{"..", "."},
		{"...", "."},
		{"\n", " "},
		{"  ", " "},
		{"   ", " "},
		{"    ", " "},
	}

	for _, c := range cases {
		content = strings.Replace(content, c.old, c.new, -1)
	}

	return content
}
