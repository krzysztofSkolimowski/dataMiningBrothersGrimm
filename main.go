package main

import (
	"fmt"
	"dataMiningBrothersGrim/sentence"
)

func main() {

	path := "rawData/GRIMMSFairyTales.txt"
	sl := sentence.ReadContent(path)
	fmt.Println("--------------------------------------------------------------------------------------------------")
	fmt.Println(sl)
	fmt.Println(len(sl.Sentences()))
	fmt.Println("--------------------------------------------------------------------------------------------------")

}
