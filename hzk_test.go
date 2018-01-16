package main

import (
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func TestMatrix(t *testing.T) {
	rand.Seed(time.Now().Unix())
	codes := "$"
	l := len(codes)
	var str = "新"
	matrix, err := Matrix([]byte(str))
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(len(matrix))
	var key = []byte{0x80, 0x40, 0x20, 0x10, 0x08, 0x04, 0x02, 0x01}
	for k := 0; k < 48; k++ {
		for j := 0; j < 6; j++ {
			for i := 0; i < 8; i++ {
				if flag := (matrix[6*k+j] & key[i]) == 0; flag {
					fmt.Print(" ")
				} else {
					index := rand.Intn(l)
					fmt.Print(string(codes[index]))
				}
			}
		}
		fmt.Println()
	}
}

func TestPrintWords(t *testing.T) {
	s := "工要在地"
	g := NewRandCharGenerator("$&JFDS")
	PrintWords(s,10, g)
}
