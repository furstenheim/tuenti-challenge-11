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
	log.Println(challenge.Cases[3])
	for i, g := range(challenge.Cases) {
		result := solveCase(g)
		logResult(i, result)
	}
}


func logResult(i int, result string) {
	fmt.Printf("Case #%d: %s\n", i + 1, result)
}



func solveCase (c Case) string {
	gcd := big.NewInt(0)
	_, b := gcd.SetString(string(c.Interferences[0]), 16)
	if !b {
		log.Fatal("Could not set value")
	}
	for i, v := range c.Interferences {
		if i == 0 {
			continue
		}
		newInt := big.NewInt(0)
		_, couldSet := newInt.SetString(string(v), 16)
		if !couldSet {
			log.Fatal("Could not set", v)
		}

		gcd = computeGCD(gcd, newInt)
	}
	two := big.NewInt(2)
	rem := big.NewInt(2)
	rem.Rem(gcd, two)
	for rem.Uint64() == 0{
		gcd.Div(gcd, two)
		rem.Rem(gcd, two)
	}
	return fmt.Sprintf("%x", gcd)
}

func computeGCD(a1, a2 *big.Int) *big.Int {
	r1 := new(big.Int)
	r1.Set(a1)
	r2 := new(big.Int)
	r2.Set(a2)

	if r1.Cmp(r2) > 0 {
		r1, r2 = r2, r1
	}
	zero := big.NewInt(0)
	for r1.Cmp(zero) > 0 {
		remove := big.NewInt(0)
		for r1.Cmp(r2) <= 0 {
			remove.Lsh(r1, uint(r2.BitLen() - r1.BitLen()))
			r2.Xor(r2, remove)
		}
		r1, r2 = r2, r1
	}
	return r2
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
	NInterferences int `index:"0"`
	Interferences  []Interference `index:"1" indexed:"NInterferences"`
}

type Interference string
