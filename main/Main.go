package main

import (
	"math/rand"
	"time"
)

func main() {
	p := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	c := []string{"@", "!", "#", "$", "%", "^", "&", "*", "(", ")", "-", "_", "=", "+", "[", "]", "{", "}", "|", ";", ":", "<", ">", "?", "/", ".", ","}
	n := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	rand.Seed(time.Now().UnixNano())
	var txt string
	for i := 0; i < 10; i++ {
		randomNumber := rand.Intn(3)
		if randomNumber == 0 {
			txt += p[rand.Intn(len(p))]
		} else if randomNumber == 1 {
			txt += c[rand.Intn(len(c))]
		} else if randomNumber == 2 {
			txt += n[rand.Intn(len(n))]
		}

	}
	print(txt)

}
