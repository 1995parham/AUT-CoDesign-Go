/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 18-05-2016
 * |
 * | File Name:     main.go
 * +===============================================
 */

package main

import (
	"github.com/tarm/serial"
	"log"
)

func main() {
	c := &serial.Config{Name: "/dev/tnt0", Baud: 9600}
	c.Parity = serial.ParityNone
	c.StopBits = serial.Stop1
	c.Size = serial.DefaultSize

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n := 0
	buf := make([]byte, 128)

	/* Initiate LFSRs */
	dlfsr81 = NewDummyLFSR8()
	dlfsr82 = NewDummyLFSR8()
	dlfsr83 = NewDummyLFSR8()

	/* Read seed number 1 */
	var s1 uint8
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("We get more than one byte :(")
	}
	s1 = uint8(buf[0])
	log.Printf("%d", s1)
	dlfsr81.Init(0, s1)

	/* Read seed number 2 */
	var s2 uint8
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	s2 = uint8(buf[0])
	log.Printf("%d", s2)
	dlfsr82.Init(0, s2)

	/* Read seed number 3 */
	var s3 uint8
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	s3 = uint8(buf[0])
	log.Printf("%d", s3)
	dlfsr83.Init(0, s3)

	LoadRefText()

	LoadEtalon()

	p := NewPopulationFromFile()

	p.Report()

	for i := 0; i < 30; i++ {
		p.Next()
		p.Report()
	}
}
