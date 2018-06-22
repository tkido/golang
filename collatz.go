package main

import "log"

var n = 27

func collatz(i int) int {
	if n&1 == 1 {
		return 3*n + 1
	}
	return n / 2
}

func main() {
	for {
		log.Printf("%032b %02d (%d)", n, count(n), n)
		if n == 1 {
			break
		}
		n = collatz(n)
	}
}

func count(n int) int {
	n = n&0x5555 + n>>1&0x5555
	n = n&0x3333 + n>>2&0x3333
	n = n&0x0f0f + n>>4&0x0f0f
	n = n&0x00ff + n>>8&0x00ff
	return n
}
