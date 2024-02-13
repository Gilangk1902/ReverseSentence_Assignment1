package ReverseSentence

import "strings"

func Reverse(input string) string {
	words := strings.Fields(input)

	i := 0
	j := len(words) - 1

	for i < j {
		temp := words[i]
		words[i] = words[j]
		words[j] = temp

		i++
		j--
	}

	new_string := strings.Join(words, " ")

	return new_string
}
