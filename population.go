/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     population.go
 * +===============================================
 */
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"sync"
)

type Population struct {
	Kromosoms [32]Kromosom
}

func (p *Population) Len() int {
	return 32
}

func (p *Population) Swap(i, j int) {
	p.Kromosoms[i], p.Kromosoms[j] =
		p.Kromosoms[j], p.Kromosoms[i]
}

func (p *Population) Less(i, j int) bool {
	return p.Kromosoms[i].Fitness < p.Kromosoms[j].Fitness
}

func (p *Population) Crossover() {
	for i := 0; i < 16; i += 2 {
		p.Kromosoms[i+16], p.Kromosoms[i+16+1] =
			Crossover(&p.Kromosoms[i], &p.Kromosoms[i+1])
	}
}

func (p *Population) Mutate() {
	for i := 0; i < 32; i++ {
		p.Kromosoms[i].Mutate()
	}
}

func (p *Population) Next() {
	var w sync.WaitGroup
	for i := 0; i < 32; i++ {
		w.Add(1)
		go func() {
			str := p.Kromosoms[i].Permute()
			p.Kromosoms[i].CalculateFitness(str)
			w.Done()
		}()
	}
	w.Wait()

	sort.Sort(p)

	p.Crossover()
	p.Mutate()
}

func (p *Population) Report() {
	for i := 0; i < 32; i++ {
		log.Printf("-- %d: %v\n", i, p.Kromosoms[i].Gen)
	}
}

func NewPopulationFromFile() *Population {
	f, err := os.Open("keys.txt")
	if err != nil {
		panic(err)
	}

	p := &Population{}

	for i := 0; i < 32; i++ {
		for j := 0; j < 16; j++ {
			_, err := fmt.Fscanf(f, "%d", &p.Kromosoms[i].Gen[j])

			if err != nil {
				panic(err)
			}
		}
		fmt.Fscanf(f, "\n")
	}

	return p
}
