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
			Crossover(p.Kromosoms[i], p.Kromosoms[i+1])
	}
}

func (p *Population) Mutate() {
	for i := 0; i < 32; i++ {
		p.Kromosoms[i].Mutate()
	}
}
