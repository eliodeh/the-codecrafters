package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	small := map[string]bool{"a": true, "an": true, "the": true, "of": true, "in": true, "on": true, "and": true, "to": true, "but": true, "or": true, "for": true, "nor": true, "at": true, "by": true, "up": true, "as": true, "is": true, "it": true}

	fmt.Println("STRING TRANSFORMER")
	for {
		fmt.Println("Valid commands: *upper <text>  *lower <text>  *cap <text>  *title <text>  *snake <text>  *reverse <text>   *exit")

		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			main()
			return
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])

		if command == "exit" {
			fmt.Println("Goodbye Codedcrafters")
			return
		}

		if len(parts) < 2 {
			fmt.Println("no text and command provided")
			continue
		}

		text := strings.Join(parts[1:], " ")

		var output string

		switch command {

		case "upper":
			output = strings.ToUpper(text)

		case "lower":
			output = strings.ToLower(text)

		case "cap":
			words := strings.Fields(text)
			for i, w := range words {

				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
			output = strings.Join(words, " ")

		case "title":
			words := strings.Fields(text)
			for i, w := range words {
				lw := strings.ToLower(w)
				if i == 0 || !small[lw] {
					words[i] = strings.ToUpper(string(lw[0])) + lw[1:]
				} else {
					words[i] = lw
				}
			}
			output = strings.Join(words, " ")

		case "snake":
			var r []rune
			for _, c := range text {
				if unicode.IsLetter(c) || unicode.IsDigit(c) {
					r = append(r, unicode.ToLower(c))
				} else if unicode.IsSpace(c) {
					r = append(r, '_')
				}
			}
			output = strings.Trim(strings.ReplaceAll(string(r), "__", "_"), "_")

		case "reverse":
			words := strings.Fields(text)
			for i, w := range words {
				r := []rune(w)
				for a, b := 0, len(r)-1; a < b; a, b = a+1, b-1 {
					r[a], r[b] = r[b], r[a]
				}
				words[i] = string(r)
			}
			output = strings.Join(words, " ")

		default:
			fmt.Println(" Unknown command")
			continue

		}

		fmt.Println("→", output)
	}
}
