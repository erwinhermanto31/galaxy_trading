package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Main struct{}

func NewMain() *Main {
	return &Main{}
}

var sentence string

var romNum = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
}

var typeCur = map[string]float64{
	"Silver": 17,
	"Gold":   14450,
	"Iron":   195.5,
}

func FloatToString(input_num float64) string {
	// to convert a float number to a string
	return strconv.FormatFloat(input_num, 'f', 6, 64)
}

func (m *Main) Find(slice []string, val string) bool {
	for _, item := range slice {
		if val == item {
			return true
		}
	}
	return false
}

func (m *Main) ConvertToNumber(number string) int {
	var result int
	ln := len(number)
	for i := 0; i < ln; i++ {
		c := string(number[i])
		vc := romNum[c]
		if i < ln-1 {
			cnext := string(number[i+1])
			vcnext := romNum[cnext]
			if vc < vcnext {
				result += vcnext - vc
				i++
			} else {
				result += vc
			}
		} else {
			result += vc
		}
	}
	return result
}

func (m *Main) ConvertToRomawi(sentence []string) string {
	var romArray []string
	var stc string
	var nm float64

	for _, val := range sentence {
		if val == "glob" {
			romArray = append(romArray, "I")
			continue
		} else if val == "prok" {
			romArray = append(romArray, "V")
			continue
		} else if val == "pish" {
			romArray = append(romArray, "X")
			continue
		} else if val == "tegj" {
			romArray = append(romArray, "L")
			continue
		}
	}

	stc = strings.Join(romArray, "")
	number := m.ConvertToNumber(stc)

	for _, vlu := range sentence {
		if vlu == "Silver" {
			nm = float64(number) * typeCur[vlu]
		} else if vlu == "Gold" {
			nm = float64(number) * typeCur[vlu]
		} else if vlu == "Iron" {
			nm = float64(number) * typeCur[vlu]
		}
	}

	stc = strconv.Itoa(int(nm))

	return stc
}

func (m *Main) SentenceParsing(sentence string) (string, string) {
	var arraySentence []string
	var stc string
	var rmw string

	rmc := []string{"glob", "prok", "pish", "tegj", "Silver", "Gold", "Iron"}

	s := strings.Split(sentence, " ")

	for i := 0; i < len(s); i++ {

		status := m.Find(rmc, s[i])
		if status == true {
			arraySentence = append(arraySentence, s[i])
		}
	}

	stc = strings.Join(arraySentence, " ")
	rmw = m.ConvertToRomawi(arraySentence)
	return stc, rmw
}

func main() {
	m := NewMain()
	consoleReader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Question : ")

	sentence, _ := consoleReader.ReadString('\n')
	sentence = strings.TrimSuffix(sentence, "\n")
	s, r := m.SentenceParsing(sentence)

	fmt.Println(s, "is", r, "Credits")
}
