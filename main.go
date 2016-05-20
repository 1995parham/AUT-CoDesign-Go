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
	/* Load refrence text from file :) */
	LoadRefText()

	/* Load etalon array */
	LoadEtalon()

	/* Build population with first generation from keys.txt file */
	p := NewPopulationFromFile()

	log.Printf("We are setup :D")

	c := &serial.Config{Name: "/dev/tnt0", Baud: 9600}
	c.Parity = serial.ParityNone
	c.StopBits = serial.Stop1
	c.Size = serial.DefaultSize

	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n := 0
	input := make([]byte, 128)

	/* Initiate LFSRs */
	dlfsr81 = NewDummyLFSR8()
	dlfsr82 = NewDummyLFSR8()
	dlfsr83 = NewDummyLFSR8()

	/* Read seed number 1 from serial */
	var s1 uint8
	n, err = s.Read(input)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("We get more than one byte :(")
	}
	s1 = uint8(input[0])
	log.Printf("%d", s1)
	dlfsr81.Init(0, s1)

	/* Read seed number 2 from serial */
	var s2 uint8
	n, err = s.Read(input)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("We get more than one byte :(")
	}
	s2 = uint8(input[0])
	log.Printf("%d", s2)
	dlfsr82.Init(0, s2)

	/* Read seed number 3 from serial */
	var s3 uint8
	n, err = s.Read(input)
	if err != nil {
		log.Fatal(err)
	}
	if n != 1 {
		log.Fatal("We get more than one byte :(")
	}
	s3 = uint8(input[0])
	log.Printf("%d", s3)
	dlfsr83.Init(0, s3)

	/* Report first generation that read from file */
	p.Report()

	for i := 0; i < 29; i++ {
		p.Next()
	}
	p.HalfNext()

	output := make([]byte, 16)
	for i := 0; i < 16; i++ {
		output[i] = uint8(p.Kromosoms[31].Gen[i])
	}
	_, err = s.Write(output)
	if err != nil {
		panic(err)
	}
	s.Close()
}
