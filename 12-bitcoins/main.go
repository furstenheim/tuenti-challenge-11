package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
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

	graph := computeGraph(c)

	visitedOneCoin := map[CryptoId]PoorStatus{"BTC": {Length: 0}}
	visitedMoreCoins := map[CryptoId]RichStatus{}

	toVisit := []VisitStatus{{
		Length: 0,
		Amount: 1,
		Crypto: "BTC",
	}}

	for len(toVisit) > 0 {
		var current VisitStatus
		current, toVisit = toVisit[0], toVisit[1:]
		neighbours := graph[current.Crypto]
		nextLength := current.Length + 1

		for n, increase := range neighbours {
			toHave := increase * current.Amount

			if toHave == 1 {
				if _, ok := visitedOneCoin[n]; !ok {
					visitedOneCoin[n] = PoorStatus{Length: nextLength}
					toVisit = append(toVisit, VisitStatus{
						Length: nextLength,
						Amount: 1,
						Crypto: n,
					})
				}
			} else if toHave == 0 {
				log.Fatal("to have 0", current.Amount, )
			} else {
				if visitValue, ok := visitedMoreCoins[n]; !ok || (visitValue.Length == nextLength && visitValue.Amount < toHave) {
					visitedMoreCoins[n] = RichStatus{
						Length: nextLength,
						Amount: toHave,
					}
					toVisit = append(toVisit, VisitStatus{
						Length: nextLength,
						Amount: toHave,
						Crypto: n,
					})
				}
			}
		}
	}

	if value, ok := visitedMoreCoins["BTC"]; !ok {
		return "1"
	} else {
		return fmt.Sprintf("%d", value.Amount)
	}
}

type VisitStatus struct {
	Length int
	Amount int
	Crypto CryptoId
}

type PoorStatus struct {
	Length int
}
type RichStatus struct {
	Length int
	Amount int
}



func computeGraph (c Case) map[CryptoId]map[CryptoId]int {
	graph := map[CryptoId]map[CryptoId]int{}
	for _, m := range c.Markets {
		for _, e := range m.Exchanges {
			if e.Amount == 0 {
				continue
			}
			existingMap, firstMapOk := graph[e.Start]

			if !firstMapOk {
				existingMap = map[CryptoId]int{}
			}

			existingMap[e.End] = max(existingMap[e.End], e.Amount)
			graph[e.Start] = existingMap
		}
	}
	return graph
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
	NMarkets int `index:"0"`
	Markets []Market `index:"1" indexed:"NMarkets"`
}

type Market struct {
	Name string `index:"0" delimiter:" "`
	NExchanges int `index:"1"`
	Exchanges []Exchange `index:"2" indexed:"NExchanges"`
}

type Exchange struct {
	Start CryptoId `index:"0" delimiter:"-"`
	Amount int `index:"1" delimiter:"-"`
	End CryptoId `index:"2"`
}

type CryptoId string
