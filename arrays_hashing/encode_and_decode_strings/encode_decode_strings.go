// Package encodedecode will solve the encode and decode strings problem.
package encodedecode

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	result := ""
	for _, str := range strs {
		// Take the word and concat "length#word" as the encoding.
		result += strconv.Itoa(len(str)) + "#" + str
	}
	return result
}

func (s *Solution) Decode(encoded string) []string {
	if len(encoded) == 0 {
		return []string{}
	}
	var res []string
	for i := 0; i < len(encoded); {
		// Find the next '#' starting from position i
		pos := strings.Index(encoded[i:], "#") + i
		length, _ := strconv.Atoi(encoded[i:pos])
		// Grab the string and update i in one go
		res = append(res, encoded[pos+1:pos+1+length])
		i = pos + 1 + length
	}
	return res
}
