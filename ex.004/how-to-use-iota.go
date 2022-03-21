package main

import (
	"flag"
	"fmt"
)

type ByteSize float64

const (
	_           = iota
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func (b ByteSize) String() string {
	switch {
	case b >= YB:
		return fmt.Sprintf("%0.2fYB", b/YB)
	case b >= ZB:
		return fmt.Sprintf("%0.2fZB", b/ZB)
	case b >= EB:
		return fmt.Sprintf("%0.2fEB", b/EB)
	case b >= PB:
		return fmt.Sprintf("%0.2fPB", b/PB)
	case b >= TB:
		return fmt.Sprintf("%0.2fTB", b/TB)
	case b >= GB:
		return fmt.Sprintf("%0.2fGB", b/GB)
	case b >= MB:
		return fmt.Sprintf("%0.2fMB", b/MB)
	case b >= MB:
		return fmt.Sprintf("%0.2fMB", b/MB)
	case b >= KB:
		return fmt.Sprintf("%0.2fKB", b/KB)
	default:
		return fmt.Sprintf("%fB", b)
	}
}

func main() {
	var arg float64
	flag.Float64Var(&arg, "v", 0, "value to print")
	flag.Parse()

	fmt.Printf("Input: %f\n", arg)
	fmt.Printf("Output: %s\n", ByteSize(arg))
}