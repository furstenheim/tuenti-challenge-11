package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main () {
	conn, err := net.Dial("tcp", "codechallenge-daemons.0x14.net:4321")
	if err != nil {
		log.Fatal(err)
	}
	globalMap := map[Position]PositionInfo{}
	globalPosition := Position{X: 0, Y: 0}

	globalMap[globalPosition]=PositionInfo{
		ComingFrom: Position{},
		Length:     0,
	}

	reader := bufio.NewReader(conn)

	line1, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println(line1)

	toVisit := []Position{}
	firstLine, firstLineErr := reader.ReadString('\n')
	if firstLineErr != nil {
		log.Fatal(firstLineErr)
	}
	if !strings.HasPrefix(firstLine,"north - south - east - west - look - where am I - go to x,y - is exit? - bye" ) {
		log.Fatal(firstLine, globalPosition)
	}

	// Last log before finding  Moving forward {31 83} 1000 101100 608
	counter := 0
	for true {
		counter++

		_, lookErr := conn.Write([]byte("look\n"))
		if lookErr != nil {
			log.Fatal(lookErr)
		}

		lookLine, readLookLineErr := reader.ReadString('\n')
		if readLookLineErr != nil {
			log.Fatal(readLookLineErr)
		}

		_, isExitErr := conn.Write([]byte("is exit?\n"))
		if isExitErr != nil {
			log.Fatal(isExitErr)
		}


		isExitLine, readIsExitLineErr := reader.ReadString('\n')
		if readIsExitLineErr != nil {
			log.Fatal(readIsExitLineErr)
		}

		log.Println("---", counter)
		isExit := parseIsExitLine(isExitLine, globalPosition)
		movements := parseLookLine(lookLine, globalPosition)

		if isExit {
			log.Println(getMovementsBackFromOriginPath(globalPosition, globalMap))
			log.Fatal(isExit, globalPosition)
		}

		// log.Println("movements", movements)
		for _, mov := range(movements) {
			newPos := Position{
				X: globalPosition.X + mov.dX,
				Y: globalPosition.Y + mov.dY,
			}

			if _, ok := globalMap[newPos]; !ok {
				globalMap[newPos] = PositionInfo{
					ComingFrom: globalPosition,
					MovementFrom: mov,
					Length:     globalMap[globalPosition].Length + 1,
				}
				toVisit = append(toVisit, newPos)
			}
		}

		if len(toVisit) == 0 {
			log.Fatal("to visit is empty", globalPosition)
		}
		var nextPosition Position
		// log.Println("Pre global position", globalPosition)
		nextPosition, toVisit = toVisit[0], toVisit[1:]
		_, movErr := conn.Write([]byte(fmt.Sprintf("go to %d,%d\n", nextPosition.X, nextPosition.Y)))
		if movErr != nil {
			log.Fatal(movErr)
		}

		movLine, readMovLineErr := reader.ReadString('\n')
		if readMovLineErr != nil {
			log.Fatal(readMovLineErr)
		}

		if !strings.HasPrefix(movLine, "Great movement. Here is your new position:") {
			log.Fatalf("Not moved correctly: '%s', '%o', '%s'", firstLine, globalPosition, nextPosition)
		}

		globalPosition = nextPosition
		log.Println(globalPosition, globalMap[globalPosition].Passes)
	}
}

func parseLookLine(line string, globalPosition Position) []Movement{
	lookPrefix := "Well, well, well, my friend. You could do these movements: "
	if !strings.HasPrefix(line, lookPrefix) {
		log.Fatal("wrong look line ", line, globalPosition)
	}
	result := make([]Movement, 0, 4)
	possibleMovements := strings.Split(line[len(lookPrefix):len(line) - 1], " ")
	for _, m := range(possibleMovements) {
		if mov, ok := movementsByName[m]; !ok {
			log.Fatalf("Unknown movement in line %s, '%s'", line, m)
		} else {
			result = append(result, mov)
		}
	}
	return result
}

func parseIsExitLine (line string, globalPosition Position) bool {
	if strings.HasPrefix(line, "No. Sorry, traveller...") {
		return false
	}
	// {-19 17}
	// log.Fatal("is Exit", line, globalPosition)
	return true
}
func reverseArray (a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
func getMovementsBackToOrigin (position Position, globalMap map[Position]PositionInfo) []string{
	movements := make([]string, 0, globalMap[position].Length)
	startPosition := Position{}
	for position != startPosition {
		fromMov := globalMap[position].MovementFrom
		movements = append(movements, fromMov.reverse)
		position = globalMap[position].ComingFrom
	}
	return movements
}

func getMovementsBackFromOriginPath (position Position, globalMap map[Position]PositionInfo) string{
	movements := make([]string, 0, globalMap[position].Length)
	startPosition := Position{}
	for position != startPosition {
		movements = append(movements, fmt.Sprintf("(%d, %d)", position.X, position.Y))
		position = globalMap[position].ComingFrom
	}

	movements = append(movements, "(0, 0)")

	reverseArray(movements)

	return strings.Join(movements, ", ")
}

func getMovementsFromOrigin (position Position, globalMap map[Position]PositionInfo) []string{
	movements := make([]string, 0, globalMap[position].Length)
	startPosition := Position{}
	for position != startPosition {
		fromMov := globalMap[position].MovementFrom
		movements = append(movements, fromMov.name)
		position = globalMap[position].ComingFrom
	}
	reverseArray(movements)
	return movements
}


var movementsByName = map[string]Movement{
	"east": { dX: -1, dY: 0, name: "east", reverse: "west", index: 0, reverseIndex: 1 },
	"west": { dX: 1, dY: 0, name: "west", reverse: "east", index: 1, reverseIndex: 0 },
	"south": { dX: 0, dY: -1, name: "south", reverse: "north", index: 2, reverseIndex: 3 },
	"north": { dX: 0, dY: 1, name: "north", reverse: "south", index: 3, reverseIndex: 2 },
}


type Movement struct {
	dX, dY int
	name string
	reverse string
	index int
	reverseIndex int
}

type VisitedPath int

const (
	NotVisited VisitedPath = iota
	Visited
	VisitedTwice
)

type PositionInfo struct {
	VisitedPaths [4]VisitedPath
	MovementFrom Movement
	ComingFrom Position
	Length int
	Passes int
}

type PositionType int

const (
	Empty PositionType = iota + 1
	Wall
	Exit
)

type Position struct {
	X, Y int
}
