package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var resultStr string

func getModdedMultiple(input string, endIdx, num int) []string {
	start := endIdx
	words := []string{}
	for start > 0 && len(words) < num {
		start--
		if input[start] == ' ' || start == 0 {
			word := strings.TrimSpace(input[start:endIdx])
			words = append([]string{word}, words...) // reverse order
			endIdx = start
		}
	}
	return words
}

func parseMods(str string) {
	mods := []string{"hex", "bin", "up", "low", "cap"}

	for r := 0; r < len(str); r++ {
		if str[r] != ')' {
			continue
		}
		l := r
		for l > 0 && str[l] != '(' {
			l--
		}

		tag := str[l+1 : r]  // cap
		full := str[l : r+1] // (cap)
		endIdx := l - 1
		num := 1

		for _, mod := range mods {
			if strings.HasPrefix(tag, mod) {
				if commaIdx := strings.Index(tag, ","); commaIdx != -1 {
					n, _ := strconv.Atoi(strings.TrimSpace(tag[commaIdx+1:]))
					num = n
				}
				resultStr = strings.Replace(resultStr, full+" ", "", 1)
				words := getModdedMultiple(str, endIdx, num)
				Modify(words, mod)
				break
			}
		}
	}
}

func Modify(words []string, mod string) {
	str := strings.Join(words, " ")

	// lastIndex
	fmt.Println(strings.LastIndex(resultStr, str))

	switch mod {
	case "hex":
		n, _ := strconv.ParseInt(str, 16, 0)
		resultStr = strings.Replace(resultStr, str, strconv.Itoa(int(n)), 1)
	case "bin":
		n, _ := strconv.ParseInt(str, 2, 0)
		resultStr = strings.Replace(resultStr, str, strconv.Itoa(int(n)), 1)
	case "up":
		resultStr = strings.Replace(resultStr, str, strings.ToUpper(str), 1)
	case "low":
		resultStr = strings.Replace(resultStr, str, strings.ToLower(str), 1)
	case "cap":
		for _, w := range words {
			if len(w) > 0 {
				cap := strings.ToUpper(string(w[0])) + w[1:]
				resultStr = strings.Replace(resultStr, w, cap, 1)
			}
		}
	}
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
	resultStr = input

	parseMods(input)
	fmt.Println(resultStr)

	os.WriteFile(args[1], []byte(resultStr), 0644)
}
