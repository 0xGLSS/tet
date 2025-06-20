package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// big num 666666
// (up, 6 6)

//I was sitting over there    ,and then BAMM !!
//[I was sitting over there, and then BAMM, !]

/*
take: "it was the worst of times (pu, 6)"
return: [it was the worst of times "(up, 6)"]
*/
func customFields(str string) []string {
	str = strings.ReplaceAll(str, "(", " (")
	str = strings.ReplaceAll(str, ")", ") ")

	slc := strings.Fields(str)
	result := []string{}

	for i := 0; i < len(slc); i++ {
		// Ensure there is a next element
		if i+1 < len(slc) && slc[i][0] == '(' && slc[i+1][len(slc[i+1])-1] == ')' {
			combined := slc[i] + slc[i+1]
			result = append(result, combined)
			i++ // Skip next element
		} else {
			result = append(result, slc[i])
		}
	}

	return result
}

// return words indexes (start & end) and they're tag
func getWordsTag(slc []string) {
	validTags := []string{"hex", "bin", "up", "low", "cap"}

	n := 1

	for i, w := range slc {
		if w[0] == '(' && w[len(w)-1] == ')' {
			for _, tag := range validTags {
				full := w[1:]

				// get tag
				if strings.HasPrefix(full, tag) {

					// get num
					comma := strings.Index(full, ",")
					if comma == (-1) {
						n = 1
					} else {
						s := strings.TrimSpace(full[comma+1 : len(full)-1])
						n, _ = strconv.Atoi(string(s))
					}

					// get words indexes
					start := i - n
					end := i

					// modify
					Modify(slc, tag, start, end)

					// remove tags
					slc[i] = "#"
				}

			}
		} else if w[0] == '(' && w[len(w)-1] != ')' || w[0] != '(' && w[len(w)-1] == ')' {
			return
		}
	}
}

// take tag, start, end
// from start to end, apply mod
func Modify(slc []string, tag string, start, end int) {

	if start < 0 {
		start = 0
	}

	switch tag {
	case "up":
		for i := start; i < end; i++ {
			slc[i] = strings.ToUpper(slc[i])
		}
	case "cap":
		for i := start; i < end; i++ {
			slc[i] = strings.Title(slc[i])
		}
	case "bin": // should be numbers (10)
		for i := start; i < end; i++ {
			bin, _ := strconv.ParseInt(string(slc[i]), 2, 0)
			slc[i] = strconv.Itoa(int(bin))
		}
	case "hex":
		for i := start; i < end; i++ {
			hex, _ := strconv.ParseInt(string(slc[i]), 16, 0)

			slc[i] = strconv.Itoa(int(hex))
		}
	}
}

func startWithVowel(s string) bool {
	switch s[0] {
	case 'a', 'e', 'i', 'o', 'u', 'h',
		'A', 'E', 'I', 'O', 'U', 'H':
		return true
	}
	return false
}

// I was sitting over there    ,and then BAMM !!
// I was sitting over there, and then BAMM !!
func FixSymbols(str []string) {
	for i := 0; i < len(str)-1; i++ {
		if str[i+1][0] == ',' {
			str[i+1] = str[i+1][1:]
			str[i] += ","
		}
	}
	fmt.Println(str)
}

// final step: remove brackets and change a to an and return output string
func RemoveTagChangeA(slc []string) string {
	result := ""

	for i, w := range slc {
		if (w == "a" || w == "A") && startWithVowel(slc[i+1]) {
			result += w + "n "
		} else if w != "#" {
			result += w + " "
		} else {
			continue
		}
	}

	return result
}

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Need 2 args: input output")
		return
	}

	data, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	input := string(data)
	slc := customFields(input)
	getWordsTag(slc)
	RemoveTagChangeA(slc)

	fmt.Println(input)
	FixSymbols(slc)
}
