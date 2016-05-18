/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     lfsr.go
 * +===============================================
 */
package main

var dlfsr81, dlfsr82, dlfsr83 LFSR8

type LFSR8 interface {
	Init(poly uint8, seed uint8)
	Next() uint8
}

type dlfsr8 struct {
	data uint8
}

func NewDummyLFSR8() LFSR8 {
	return &dlfsr8{}
}

func (d *dlfsr8) Init(poly uint8, seed uint8) {
	d.data = seed
}

func (d *dlfsr8) Next() uint8 {
	d.data = ((((d.data >> 7) ^ (d.data >> 5) ^ (d.data >> 4) ^ (d.data >> 3)) & 0x01) | (d.data << 1))
	return d.data
}
