/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     text.go
 * +===============================================
 */
package main

import (
	"io/ioutil"
)

var RefText string

func LoadRefText() {
	dat, err := ioutil.ReadFile("refText.txt")
	if err != nil {
		panic(err)
	}

	RefText = string(dat)
}
