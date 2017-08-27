package ex2_4

var pc [256]byte

func init(){
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountWithoutTable(x uint64) int {
	bitCount := 0
	for n := uint(0); n < 64; n++ {
		if x&1 == 1 { bitCount++ }
		x >>= 1
	}
	return bitCount
}
