package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"os"
	"regexp"
	"sort"
)

var asciiPattern = "^[a-zA-Z]+$"

func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	unmarshallErr := challenge_encoding.Unmarshall(challenge, reader)
	if unmarshallErr != nil {
		log.Fatal(unmarshallErr)
	}
	log.Println(challenge.Cases[29])
	for i, g := range(challenge.Cases) {
		result := solveCase(g)
		logResult(i, result)
	}
}

func logResult(i int, result string) {
	fmt.Printf("Case #%d: %s\n", i + 1, result)
}



func solveCase (c Case) string {
	sort.Strings(c.Functions)
	sum := 0

	leftStrings := c.Functions

	for j := 0; j < c.NTotalFunctions  / c.NFunctionsPerFile; j++ {
		currentBestScore := -1
		currentBestIndex := -1
		for i := 0; i < len(leftStrings) - c.NFunctionsPerFile + 1; i++ {
			a := c.Functions[i]
			endIndex := i + c.NFunctionsPerFile - 1
			b := c.Functions[endIndex]
			possibleBest := findCommonRoot(a, b)
			if possibleBest > currentBestScore {
				currentBestIndex = i
				currentBestScore = possibleBest
			}
		}
		leftStrings = append(leftStrings[0: currentBestIndex], leftStrings[currentBestIndex + c.NFunctionsPerFile:]...)
		sum += currentBestScore
	}

	return fmt.Sprintf("%d", sum)
}

func findBestForPosition (bestResults []BestResult, index int) int {
	for i := len(bestResults) - 1; i >= 0; i-- {
		if bestResults[i].index < index {
			return bestResults[i].value
		}
	}
	log.Fatal("find best should not reach here")
	return 0
}

func findCommonRoot (a, b string) int {
	res1, e1 := regexp.MatchString(asciiPattern, a)
	res2, e2 := regexp.MatchString(asciiPattern, b)
	if e1 != nil {
		log.Fatal("Could not match s1", e1, a)
	}
	if e2 != nil {
		log.Fatal("Could not match s1", e2, b)
	}
	if !res1 || ! res2 {
		log.Fatal("Non ascii function", a, b, res1, res2)
	}

	common := 0

	for i := 0; i < min(len(a), len(b)); i++ {
		if a[i] == b[i] {
			common++
		} else {
			break
		}
	}
	return common
}

type BestResult struct {
	index int
	value int
}

func min (i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max (i, j int) int {
	if i > j {
		return i
	}
	return j
}


type Challenge struct {
	NCases int    `index:"0"`
	Cases []Case `index:"1" indexed:"NCases"`
}
type Case struct {
	NTotalFunctions int `index:"0" delimiter:" "`
	NFunctionsPerFile int `index:"1"`
	Functions []string `index:"2" indexed:"NTotalFunctions"`
}
