package main

import (
	"math/rand"
	"time"
)

type CharGenerator interface{
	Next() byte
}

type RandCharGenerator struct {
	codes string
}

func NewRandCharGenerator(codes string) *RandCharGenerator {
	rand.Seed(time.Now().Unix())
	r := RandCharGenerator{
		codes: codes,
	}
	return &r
}

func (g *RandCharGenerator) Next() byte {
	index := rand.Intn(len(g.codes))
	return g.codes[index]
}

type SingleCharGenerator struct {
	Char byte
}

func (g *SingleCharGenerator) Next() byte {
	return g.Char
}