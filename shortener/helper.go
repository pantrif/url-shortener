package shortener

import "strings"

//ALPHABET the supported alphabet
const ALPHABET = "23456789bcdfghjkmnpqrstvwxyzBCDFGHJKLMNPQRSTVWXYZ-_"

//BASE base51
const BASE = 51

func encode(num int64) string {
	str := ""
	for num > 0 {
		str = string(ALPHABET[num%BASE]) + str
		num = num / BASE
	}
	return str
}

func decode(str string) int64 {
	num := 0
	l := len([]rune(str))
	for i := 0; i < l; i++ {
		num = num*BASE + strings.Index(ALPHABET, string(str[i]))
	}
	return int64(num)
}
