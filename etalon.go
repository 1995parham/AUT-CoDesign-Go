/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     etalon.go
 * +===============================================
 */
package main

import (
	"fmt"
	"os"
)

var E [27][27]uint64

func LoadEtalon() {
	f, err := os.Open("etalon.txt")
	if err != nil {
		panic(err)
	}

	var ltable string
	for i := 0; i < 28; i++ {
		for j := 0; j < 28; j++ {
			if i == 0 {
				fmt.Fscanf(f, "%s", &ltable)
			} else if j == 0 {
				fmt.Fscanf(f, "%s", &ltable)
			} else {
				fmt.Fscanf(f, "%s %d", &ltable, &E[i-1][j-1])
			}
		}
		fmt.Fscanf(f, "\n")
	}
}
