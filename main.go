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
	"strconv"
)

func main() {
	c := &serial.Config{Name: "/dev/ttyUSB0", Baud: 9600}
	c.Parity = serial.ParityNone
	c.StopBits = serial.Stop1
	c.Size = serial.DefaultSize
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	n := 0
	buf := make([]byte, 128)

	/* Read seed number 1 */
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	s1, err := strconv.Atoi(string(buf[:n]))
	log.Printf("%d", s1)

	/* Read seed number 2 */
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	s2, err := strconv.Atoi(string(buf[:n]))
	log.Printf("%d", s2)

	/* Read seed number 3 */
	n, err = s.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	s3, err := strconv.Atoi(string(buf[:n]))
	log.Printf("%d", s3)
}
