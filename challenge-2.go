package main

import (
	"fmt"
	"unicode"
)

func main() {
	randStr := "САШАРВО"
	for i := 0; i < 5; i++ {
    fmt.Printf("Nilai i = %d\n", i)
		for j := 0; j <= 10; j++ {
			if i == 4 {
				if j == 5 {
					for k, v := range randStr {
						if unicode.Is(unicode.Cyrillic, v) {
							fmt.Printf("character %U '%c' starts at byte position %d\n", v, v, k)
						}
					}
					continue
				} 
				fmt.Println("Nilai j =", j)
			}
		}
	}
}

