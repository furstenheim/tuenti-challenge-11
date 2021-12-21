package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	challenge_encoding.Unmarshall(challenge, reader)
	log.Println(challenge.Cases[0])
	for i, g := range(challenge.Cases) {
		result := solveCase(g)
		logResult(i, result)
	}
}

func logResult(i int, result string) {
	fmt.Printf("Case #%d: %s\n", i + 1, result)
}

func solveCase (c Case) string {
	startIndex := 0
	index := 0
	var left, right string
	var parsingType ParsingType
	for index < len(c.Input) {
		if c.Input[index] == '-' {
			left = c.Input[startIndex: index]
			startIndex = index + 1
		} else if c.Input[index] == '|' {
			right = c.Input[startIndex: index]
			startIndex = index
			if c.Input[index + 1] == '{' {
				parsingType = Object
			} else if c.Input[index + 1] == '[' {
				parsingType = Tuples
			} else {
				parsingType = List
			}
			break
		}
		index++
	}

	var valueMap map[byte]*big.Rat

	switch parsingType {
	case Tuples:
		valueMap = parseTuple(c.Input[index + 1:])
	case List:
		valueMap = parseList(c.Input[index + 1:])
	case Object:
		valueMap = parseObject(c.Input[index + 1:])
	}
	// log.Println(valueMap, left, right)
	leftValue := computeValue(left,valueMap)
	rightValue := computeValue(right,valueMap)

	compare := leftValue.Cmp(rightValue)
	if compare < 0 {
		return right
	}
	if compare > 0 {
		return left
	}
	return "-"
}

func computeValue (input string, byValue map[byte]*big.Rat) *big.Rat {
	result := big.NewRat(0, 1)
	for i := 0; i < len(input); i++ {
		value, ok := byValue[input[i]]
		if !ok {
			log.Fatal("Missing letter", string(input[i]))
		}
		result.Add(result, value)
	}
	return result
}

func parseList (input string) map[byte]*big.Rat {
	parsed := strings.Split(input, ",")
	byValue := map[byte]*big.Rat{}
	for _, p := range(parsed) {
		if p[1] != '=' {
			log.Fatal("Unexpected format for list", p)
		}
		rat := parseRat(p[2:])
		if _, ok := byValue[p[0]]; ok {
			log.Fatal("Existing value for list: ", string(p[0]))
		}
		byValue[p[0]] = rat
	}
	return byValue
}

func parseObject (input string) map[byte]*big.Rat {
	// remove start end end {}
	input = input[1: len(input) - 1]
	parsed := strings.Split(input, ", ")
	byValue := map[byte]*big.Rat{}
	for _, p := range(parsed) {
		if p[0] != '\'' || p[2] != '\'' || p[3] != ':' || p[4] != ' ' {
			log.Fatal("Unexpected format for object", p)
		}
		rat := parseRat(p[5:])
		if _, ok := byValue[p[1]]; ok {
			log.Fatal("Existing value for object: ", string(p[1]))
		}
		byValue[p[1]] = rat
	}
	return byValue
}

func parseTuple (input string) map[byte]*big.Rat {
	// remove start end end [( )]
	input = input[2: len(input) - 2]
	parsed := strings.Split(input, "), (")
	byValue := map[byte]*big.Rat{}
	for _, p := range(parsed) {
		if p[0] != '\'' || p[2] != '\'' || p[3] != ',' || p[4] != ' ' {
			log.Fatal("Unexpected format for tuple", p)
		}
		rat := parseRat(p[5:])
		if _, ok := byValue[p[1]]; ok {
			log.Fatal("Existing value for tuple", string(p[1]))
		}
		byValue[p[1]] = rat
	}
	return byValue
}

func parseRat (input string) *big.Rat {
	splitByDelimiter := strings.Split(input, "/")
	if len(splitByDelimiter) == 1 {
		atoi, parseErr := strconv.Atoi(splitByDelimiter[0])
		if parseErr != nil {
			log.Fatal(parseErr)
		}
		return big.NewRat(int64(atoi), 1)
	}

	if (len(splitByDelimiter) == 2) {
		atoiLeft, parseErrLeft := strconv.Atoi(splitByDelimiter[0])
		if parseErrLeft != nil {
			log.Fatal(parseErrLeft)
		}
		atoiRight, parseErrRight := strconv.Atoi(splitByDelimiter[1])
		if parseErrRight != nil {
			log.Fatal(parseErrRight)
		}
		return big.NewRat(int64(atoiLeft), int64(atoiRight))
	}
	log.Fatal("Failed", input)
	return nil
}

type ParsingType uint64

const (
	Tuples ParsingType = iota + 1
	List
	Object
)

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
	Cases  []Case `index:"1" indexed:"NCases"`
}
type Case struct {
	Input string `index:"0"`
}