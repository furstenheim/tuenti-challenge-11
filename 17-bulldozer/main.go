package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"math/big"
	"os"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	unmarshallErr := challenge_encoding.Unmarshall(challenge, reader)
	if unmarshallErr != nil {
		log.Fatal(unmarshallErr)
	}
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
	a := big.NewInt(0)
	count1 := 0
	count2 := 0
	three := big.NewInt(3)
	for _, p := range c.Piles {
		_, ok := a.SetString(string(p), 10)
		if !ok {
			log.Fatal("Could not set volume", p)
		}
		a.Mod(a, three)
		module := a.Uint64()
		if module == 2 {
			count2++
		}
		if module == 1 {
			count1++
		}
	}
	if count1 % 2 == 0 && count2 % 2 == 0 {
		return "Alberto"
	}
	return "Edu"
}




func min (i, j uint64) uint64 {
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
	NPiles int `index:"0"`
	Piles  []Pile `index:"1" elem_delimiter:"space" indexed:"NPiles"`
}

type Pile string
