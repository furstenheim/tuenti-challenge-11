package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/furstenheim/challenge_encoding"
	"unicode/utf8"
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
	result := make([]rune, 0, c.NColumns * c.NRows)
	for _, row := range(c.Rows) {
		total := 0
		for _, s := range(row) {
			if s != ' ' {
				result = append(result, s)
				total ++
			}
		}
		if total != c.NColumns {
			log.Fatal("Wrong number of columns")
		}
	}

	pokemonByStartingLetter := map[rune][]string{}

	for _, pok := range(c.Pokemon) {
		firstLetter, _ := utf8.DecodeRuneInString(pok)
		if _, ok := pokemonByStartingLetter[firstLetter]; !ok {
			pokemonByStartingLetter[firstLetter] = []string{pok}
		} else {
			pokemonByStartingLetter[firstLetter] = append(pokemonByStartingLetter[firstLetter], pok)
		}
	}

	for _, _ = range(c.Pokemon) {
		found := false
		SearchLoop:
			for resultIndex, v := range(result) {
			if _, ok := pokemonByStartingLetter[v]; !ok {
				continue
			}
			pokemonsForLetter := pokemonByStartingLetter[v]
			for pi, pok := range(pokemonsForLetter) {
				if string(result[resultIndex: min(len(result), resultIndex+ len(pok))]) == pok {
					pokemonByStartingLetter[v] = append(pokemonsForLetter[0: pi], pokemonsForLetter[pi + 1: len(pokemonsForLetter)]...)
					result = append(result[0: resultIndex], result[resultIndex + len(pok): len(result)]...)
					found = true
					break SearchLoop
				}
				if Reverse(string(result[max(0, resultIndex + 1 - len(pok)): resultIndex + 1])) == pok {
					pokemonByStartingLetter[v] = append(pokemonsForLetter[0: pi], pokemonsForLetter[pi + 1: len(pokemonsForLetter)]...)
					result = append(result[0: resultIndex + 1 - len(pok)], result[resultIndex + 1: len(result)]...)
					found = true
					break SearchLoop
				}

			}
		}
		if !found {
			log.Fatal("Missing pokemons", pokemonByStartingLetter, string(result))
		}
	}
	return string(result)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
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
	Cases  []Case `index:"1" indexed:"NCases"`
}
type Case struct {
	NPokemon int `index:"0"  delimiter:"space"`
	NRows int `index:"1"  delimiter:"space"`
	NColumns int `index:"2"`
	Pokemon []string `index:"3" indexed:"NPokemon"`
	Rows []string `index:"4" indexed:"NRows"`
}