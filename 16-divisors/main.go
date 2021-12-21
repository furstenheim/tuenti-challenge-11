package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)
const N_QUESTIONS = 1500
const SIZE = 100
func main () {
	conn, err := net.Dial("tcp", "codechallenge-daemons.0x14.net:7162")
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(conn)

	line1, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1)
	}
	log.Println(line1)

	size := SIZE
	divisors := make([][]PrimePower, SIZE)
	for i, _ := range divisors {
		divisors[i] = []PrimePower{}
	}
	mostlyNotPrimes := make([]bool, SIZE)

	nQuestions := 1
	currentQuestion: for currentIndex := 0; currentIndex < SIZE; currentIndex++ {
		if nQuestions == N_QUESTIONS {
			break
		}
		currentDivisors := divisors[currentIndex]
		if mostlyNotPrimes[currentIndex] {
			continue
		}

		for bIndex := currentIndex + 1; bIndex < size; bIndex++ {
			if nQuestions == N_QUESTIONS {
				break currentQuestion
			}/*
			if mostlyNotPrimes[bIndex] {
				continue
			}*/
			/*if bDivisors, aDivisors := divisors[bIndex], divisors[currentIndex]; len(bDivisors) == 1 && len(aDivisors) == 1 && aDivisors[0].base == bDivisors[0].base {
				continue
			}*/
			nQuestions++

			question := fmt.Sprintf("? %d %d\n", currentIndex + 1 , bIndex + 1)
			log.Println(question)
			_, questionErr := conn.Write([]byte(question))
			if questionErr != nil {
				log.Fatal(questionErr)
			}

			answerLine, readAnswerErr := reader.ReadString('\n')
			if readAnswerErr != nil {
				log.Fatal("read error", readAnswerErr)
			}
			log.Println(answerLine, nQuestions)
			parsed, answerErr := strconv.Atoi(strings.Trim(answerLine, "\n"))
			if answerErr != nil {
				log.Fatal("answer error", answerErr)
			}
			factorization := FactorizeNumber(parsed)
			if len (factorization) == 0 {
				continue
			}

			newBDivisors := mergeDivisors(divisors[bIndex], factorization)
			divisors[bIndex] = newBDivisors
			mostlyNotPrimes[bIndex] = isMostlyNotPrime(newBDivisors)

			currentDivisors = mergeDivisors(currentDivisors, factorization)
			divisors[currentIndex] = currentDivisors
			mostlyNotPrimes[currentIndex] = isMostlyNotPrime(currentDivisors)
			if mostlyNotPrimes[currentIndex] {
				continue currentQuestion
			}
		}
	}

	answer := []string{"!"}

	for i, v := range mostlyNotPrimes {
		if !v {
			answer = append(answer, strconv.Itoa(i + 1))
		}
		if len(answer) == 26 + 1 {
			break
		}
	}
	answerString := strings.Join(answer, " ")
	log.Println(answerString)
	_, answerErr := conn.Write([]byte(answerString + "\n"))
	if answerErr != nil {
		log.Fatal(answerErr)
	}

	answerLine, readAnswerErr := reader.ReadString('\n')
	if readAnswerErr != nil {
		log.Fatal("read error", readAnswerErr)
	}

	log.Println(answerLine)

}

func mergeDivisors (aDivisors, bDivisors []PrimePower) []PrimePower {
	if len(aDivisors) > 1 {
		return aDivisors
	}
	if len (bDivisors) > 1 {
		return bDivisors
	}
	if len (aDivisors) == 0 {
		return bDivisors
	}
	if len (bDivisors) == 0 {
		return aDivisors
	}

	if aDivisors[0].base != bDivisors[0].base {
		return []PrimePower{aDivisors[0], bDivisors[0]}
	}

	return []PrimePower{{
		base:  aDivisors[0].base,
		power: max(aDivisors[0].power, bDivisors[0].power),
	}}

}

func isMostlyNotPrime (divisors []PrimePower) bool {
	return len(divisors) > 1 || (len(divisors) == 1 && divisors[0].power > 1)
}

func max (i, j int) int {
	if i > j {
		return i
	}
	return j
}

type PrimePower struct {
	base int
	power int
}

func FactorizeNumber (input int) []PrimePower {
	result := []PrimePower{}
	for i := 2; i * i <= input; i++ {
		primePower := PrimePower{base: i}
		for input % i == 0 {
			primePower.power++
			input /= i
		}
		if primePower.power > 0 {
			result = append(result, primePower)
		}
	}
	if input > 1 {
		result = append(result, PrimePower{
			base:  input,
			power: 1,
		})
	}
	return result
}