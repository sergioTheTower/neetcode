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
		// Find the hash tag.
		pos := strings.Index(encoded[i:], "#") + i
		// Use i position and the hashtag index to find the length of the word.
		length, _ := strconv.Atoi(encoded[i:pos])
		// Now using slicing to grab the word.
		res = append(res, encoded[pos+1:pos+1+length])
		// Make sure the loop starts at the end of the word.
		i = pos + 1 + length
	}
	return res
}
