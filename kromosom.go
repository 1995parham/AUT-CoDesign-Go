/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     kromosom.go
 * +===============================================
 */
package main

type Kromosom struct {
	Gen     [16]int
	Fitness uint64
}

func (k *Kromosom) CalculateFitness(s string) {
	var T [27][27]uint64

	k.Fitness = 0

	for i := 0; i < len(s)-1; i++ {
		var a, b int

		if s[i] == ' ' {
			a = 26
		} else {
			a = int(s[i] - 'a')
		}

		if s[i+1] == ' ' {
			b = 26
		} else {
			b = int(s[i+1] - 'a')
		}

		T[a][b]++
	}

	for i := 0; i < 27; i++ {
		for j := 0; j < 27; j++ {
			tmp := (T[i][j] * 0x10000) - E[i][j]

			if tmp > 0 {
				k.Fitness += tmp
			} else {
				k.Fitness -= tmp
			}
		}
	}
}

func (k *Kromosom) Mutate() {
	/* Get p from LFSR2 */
	var p uint8

	if p < 64 {
		/* Get indicator from LFSR3 */
		var indicator uint8

		a := indicator & 0x0F
		b := indicator >> 4

		k.Gen[a], k.Gen[b] = k.Gen[b], k.Gen[a]
	}
}

func Crossover(k1 *Kromosom, k2 *Kromosom) (Kromosom, Kromosom) {
	var d1 Kromosom
	var d2 Kromosom

	var d1i [16]bool
	var d2i [16]bool

	var alpha uint8
	/* Get alpha from LFSR1 */
	alpha = alpha & 0x0F

	for i := 0; i < int(alpha); i++ {
		d1.Gen[i] = k1.Gen[i]
		d1i[k1.Gen[i]-1] = true

		d2.Gen[i] = k2.Gen[i]
		d2i[k2.Gen[i]-1] = true

	}

	var i1, i2 int
	for i := alpha; i < 16; i++ {
		flag1 := true
		for ; i1 < 16 && flag1; i1++ {
			index := k2.Gen[i1]
			if d1i[index-1] == false {
				d1.Gen[i] = index
				d1i[index-1] = true
				flag1 = false
			}
		}

		flag2 := true
		for ; i2 < 16 && flag2; i2++ {
			index := k1.Gen[i2]
			if d2i[index-1] == false {
				d2.Gen[i] = index
				d2i[index-1] = true
				flag2 = false
			}
		}
	}

	return d1, d2
}
