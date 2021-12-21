package main

import (
	"bufio"
	"fmt"
	"github.com/furstenheim/challenge_encoding"
	"log"
	"os"
	"strings"
)

var allNotes = []Note{
	{rootValue: "F", faceValue: "F#", intValue: 0},
	{rootValue: "G", faceValue: "Gb", intValue: 0},
	{rootValue: "G", faceValue: "G", intValue: 1},
	{rootValue: "G", faceValue: "G#", intValue: 2},
	{rootValue: "A", faceValue: "Ab", intValue: 2},
	{rootValue: "A", faceValue: "A", intValue: 3},
	{rootValue: "A", faceValue: "A#", intValue: 4},
	{rootValue: "B", faceValue: "Bb", intValue: 4},
	{rootValue: "B", faceValue: "B", intValue: 5},
	{rootValue: "C", faceValue: "Cb", intValue: 5},
	{rootValue: "B", faceValue: "B#", intValue: 6},
	{rootValue: "C", faceValue: "C", intValue: 6},
	{rootValue: "C", faceValue: "C#", intValue: 7},
	{rootValue: "D", faceValue: "Db", intValue: 7},
	{rootValue: "D", faceValue: "D", intValue: 8},
	{rootValue: "D", faceValue: "D#", intValue: 9},
	{rootValue: "E", faceValue: "Eb", intValue: 9},
	{rootValue: "E", faceValue: "E", intValue: 10},
	{rootValue: "F", faceValue: "Fb", intValue: 10},
	{rootValue: "E", faceValue: "E#", intValue: 11},
	{rootValue: "F", faceValue: "F", intValue: 11},
}

var notesIndex = map[string]int{}

func main () {
	reader := bufio.NewReader(os.Stdin)
	challenge := &Challenge{}
	challenge_encoding.Unmarshall(challenge, reader)
	log.Println(challenge.Cases[0])
	for i, v := range(allNotes) {
		notesIndex[v.faceValue] = i
	}
	for i, g := range(challenge.Cases) {
		result := solveCase(g)
		logResult(i, result)
	}
}

func logResult(i int, result string) {
	fmt.Printf("Case #%d: %s\n", i + 1, result)
}

type Note struct {
	rootValue string
	faceValue string
	intValue int
}


func solveCase (c Case) string {
	toneSeries := strings.Split(c.Series, "")
	currentIndex := notesIndex[c.Start]
	startNote := allNotes[currentIndex]
	currentNote := startNote
	result := startNote.faceValue

	if len(toneSeries) != 7 {
		log.Fatal("Unexpected length")
	}
	for _, tone := range toneSeries {
		var diff int
		if tone == "s" {
			diff = 1
		} else if tone == "T" {
			diff = 2
		} else {
			log.Fatal("Unexpected tone", tone)
		}

		found := false
		for i := 0; i < 5; i++ {
			currentIndex = (currentIndex + 1) % len(allNotes)
			possibleNote := allNotes[currentIndex]
			if possibleNote.intValue == (currentNote.intValue + diff) % 12 && possibleNote.rootValue != currentNote.rootValue {
				currentNote = possibleNote
				found = true
				result = result + currentNote.faceValue
				break
			}
		}
		if !found {
			log.Fatal("Not found next note: ", result, diff)
		}
	}
	if startNote.faceValue != currentNote.faceValue {
		log.Fatal("Start and end do not match", result)
	}
	return result
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
	Start string `index:"0"`
	Series string `index:"1"`
}