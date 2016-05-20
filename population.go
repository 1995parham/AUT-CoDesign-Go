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
	"github.com/1995parham/LFSR-Go"
	"log"
	"os"
	"sort"
	"sync"
)

type Population struct {
	/* Population kromosoms */
	Kromosoms [32]Kromosom
	/* Current generation index */
	Generation int
	/* Crossover alpha LFSR */
	lfsr81 lfsr.LFSR8
	/* Mutate portability LFSR */
	lfsr82 lfsr.LFSR8
	/* Mutate indicator LFSR */
	lfsr83 lfsr.LFSR8
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
	var alpha uint8
	for i := 0; i < 16; i += 2 {
		alpha = p.lfsr81.Next()
		p.Kromosoms[i+16], p.Kromosoms[i+16+1] =
			Crossover(&p.Kromosoms[i], &p.Kromosoms[i+1], alpha)
	}
}

func (p *Population) Mutate() {
	var mp uint8
	var indicator uint8

	for i := 0; i < 32; i++ {
		/* Get p from LFSR2 */
		mp = p.lfsr82.Next()
		if mp < 64 {
			/* Get indicator from LFSR3 */
			indicator = p.lfsr83.Next()
			p.Kromosoms[i].Mutate(indicator)
		}
	}
}

func (p *Population) Next() {
	var w sync.WaitGroup
	for i := 0; i < 32; i++ {
		w.Add(1)
		go func(i int) {
			str := p.Kromosoms[i].Permute()
			p.Kromosoms[i].CalculateFitness(str)
			w.Done()
		}(i)
	}
	w.Wait()

	sort.Sort(p)

	p.Report()

	p.Crossover()
	p.Mutate()

	p.Generation++
}

func (p *Population) HalfNext() {
	var w sync.WaitGroup
	for i := 0; i < 32; i++ {
		w.Add(1)
		go func(i int) {
			str := p.Kromosoms[i].Permute()
			p.Kromosoms[i].CalculateFitness(str)
			w.Done()
		}(i)
	}
	w.Wait()

	sort.Sort(p)

	p.Report()

	p.Generation++
}

func (p *Population) Report() {
	log.Printf("==== %d ====", p.Generation)
	for i := 0; i < 32; i++ {
		log.Printf("-- %d: %v ** %v", i, p.Kromosoms[i].Gen, p.Kromosoms[i].Fitness)
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
