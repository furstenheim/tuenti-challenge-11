package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"os"
	"sort"
	"strings"
)


func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	marshallErr := challenge_encoding.Unmarshall(challenge, reader)
	if marshallErr != nil {
		log.Fatal(marshallErr)
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

type City string


func solveCase (c Case) string {
	graph := map[City][]City{}

	for _, p := range c.Connections {
		if pathFromStart, ok := graph[p.Start]; !ok {
			graph[p.Start] = []City{p.End}
		} else {
			graph[p.Start] = append(pathFromStart, p.End)
		}

		if pathFromEnd, ok := graph[p.End]; !ok {
			graph[p.End] = []City{p.Start}
		} else {
			graph[p.End] = append(pathFromEnd, p.Start)
		}
	}
	result := []string{}
	for city, _ := range(graph) {
		isNotCritical := canNavigateWithoutCity(graph, city)
		if !isNotCritical {
			result = append(result, string(city))
		}
	}
	if len(result) == 0 {
		return "-"
	}
	sort.Strings(result)
	return strings.Join(result, ",")
}


func canNavigateWithoutCity (graph map[City][]City, removeNode City) bool {
	if len(graph) == 1 {
		return true
	}
	var startingCity City

	for startingCity, _ = range graph {
		if startingCity != removeNode {
			break
		}
	}
	visited := map[City]bool{startingCity: true}
	toVisit := []City{startingCity}
	currentCity := startingCity

	for len(toVisit) != 0 {
		currentCity, toVisit = toVisit[0], toVisit[1:]
		for _, candidateCity := range(graph[currentCity]) {
			if candidateCity == removeNode {
				continue
			}
			if _, ok := visited[candidateCity]; !ok {
				visited[candidateCity] = true
				toVisit = append(toVisit, candidateCity)
			}
		}
	}

	return len(visited) == len(graph) - 1
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
	NConnections int `index:"0"`
	Connections []Connection `index:"1" indexed:"NConnections"`
}
type Connection struct {
	Start City `index:"0" delimiter:","`
	End   City `index:"1"`
}