package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"math/big"
	"os"
	"sort"
)

func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	unmarshallErr := challenge_encoding.Unmarshall(challenge, reader)
	if unmarshallErr != nil {
		log.Fatal(unmarshallErr)
	}
	log.Println(challenge.Cases[0])
	world := prepareWorld(challenge)
	for i, g := range(challenge.Cases) {
		result := solveCase(g, world)
		logResult(i, result)
	}
}
const MODULO = 100_000_007
func prepareWorld (c *Challenge) map[uint64]uint64 {
	input := make([]uint64, c.NCases)

	for i, v := range c.Cases {
		input[i] = v.Value
	}

	result := map[uint64]uint64{}

	sort.Slice(input, func(i, j int) bool { return input[i] < input[j] })
	for _, v := range input {
		result[v] = 0
	}

	max := input[len(input) - 1]
	var it uint64

	current := big.NewInt(1)
	modulo := big.NewInt(MODULO)
	placeholder := big.NewInt(1)


	for it = 1; it < min(max, MODULO) + 1; it++ {
		if !isTuentiNumber(it) {
			placeholder.SetUint64(it)
			current.Mul(current, placeholder)
			current.Rem(current, modulo)
		}

		if _, ok := result[it]; ok {
			result[it] = current.Uint64()
		}
	}
	return result
}

func isTuentiNumber (i uint64) bool {
	return ((i / 10) % 10 == 2) || ((i / 10_000) % 10 == 2) || ((i / 10_000_000) % 10 == 2)
}

func logResult(i int, result string) {
	fmt.Printf("Case #%d: %s\n", i + 1, result)
}



func solveCase (c Case, world map[uint64]uint64) string {
	value, ok := world[c.Value]
	if !ok {
		log.Fatal("Unknown value", value)
	}
	return fmt.Sprintf("%d", value)
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
	Value uint64 `index:"0"`
}
