package main

import (
	"fmt"
	"time"
)

func PrintWords(s string,speed int, generator CharGenerator) {
	for _, v := range s {
		PrintCharacter(v, speed, generator)
	}
}

func PrintCharacter(c rune, speed int, generator CharGenerator) {
	word := string(c)
	matrix, err := Matrix([]byte(word))
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Println(len(matrix))
	var key = []byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	for k := 0; k < 48; k++ {
		buf := make([]byte, 0)
		for j := 0; j < 6; j++ {
			for i := 0; i < 8; i++ {
				if flag := (matrix[6*k+j] & key[i]) == 0; flag {
					buf = append(buf, ' ')
				} else {
					buf = append(buf, generator.Next())
				}
			}
		}
		fmt.Println(string(buf))
		time.Sleep(time.Duration(speed) * time.Millisecond)
	}
}


