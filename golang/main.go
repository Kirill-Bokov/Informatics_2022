package main

import (
	"fmt"
	"strings"
)

func main() {
	inps := "2"
	fmt.Println(len(inps))
	pa := []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż"}
	ea := []string{"a", "c", "e", "l", "n", "o", "s", "z", "z"}
	fmt.Println(len(ea), len(pa))
	psl := strings.Split(inps, "")[:]
	for i := 0; i <= len(inps)-1; i++ {
		fmt.Println("i =", i)
		for j := 0; j <= len(pa)-1; j++ {
			fmt.Println("j =", j)
			if psl[i] == pa[j] {
				psl[i] = ea[j]
			}
		}
	}
	engstr := ""
	for k := 0; k <= len(psl)-1; k++ {
		engstr = engstr + psl[k]
	}
	fmt.Println(engstr)
}
