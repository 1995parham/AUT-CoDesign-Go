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
