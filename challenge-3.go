package main

import "fmt"

func main() {
	fmt.Println(countLetters("selamat malam"))
}

func countLetters(str string) map[string]int {
    counts := make(map[string]int)
    for _, c := range str {
        if string(c) == " " {
            counts[""]++
        } else {
            counts[string(c)]++
        }
        fmt.Printf("%c\n", c)
    }
    return counts
}