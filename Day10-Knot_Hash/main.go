package main

import (
	"fmt"
	"log"
)

func main() {
	partOne()
	partTwo()
}

func partTwo() {
	var pos, skip = 0, 0
	var max = 256
	var input = "83,0,193,1,254,237,187,40,88,27,2,255,149,29,42,100"
	//var input = "AoC 2017"
	//var input = "1,2,3"
	var ints []int
	var fixedLenthVals = []int{17, 31, 73, 47, 23}
	for _, i := range input {
		ints = append(ints, int(i))
	}
	for _, i := range fixedLenthVals {
		ints = append(ints, int(i))
	}
	hash := getSeed(max)
	for i := 0; i < 64; i++ {
		iPos, iSkip := pos, skip
		for idx, n := range ints {
			if n == 0 {
				iPos += iSkip
				if iPos > max {
					iPos = iPos % max
				}
				iSkip++
				continue
			}
			chain := make([]int, max)
			copy(chain, hash)
			tmp := []int{}
			if iPos+n > max {
				end := chain[iPos:]
				endLen := len(end)
				start := chain[:(iPos+n)%max]
				startLen := len(start)
				tmp = append(end, start...)
				tmp = reverse(tmp)
				end = tmp[:endLen]
				start = tmp[endLen:]
				var dt []int
				if len(start)+len(end) == max {
					dt = append(start, end...)
				} else {
					dt = append(start, append(hash[startLen:len(hash)-endLen], end...)...)
				}
				hash = dt
			} else {
				chain = chain[iPos : iPos+n]
				if len(chain) != 1 {
					chain = reverse(chain)
					tmp = append(hash[:iPos], append(chain, hash[n+iPos:]...)...)
					hash = tmp
				}
			}
			iPos += n + iSkip
			if iPos > max {
				iPos = iPos % max
			}

			iSkip++

			if idx == len(ints)-1 {
				pos = iPos
				skip = iSkip
			}
		}
	}

	hiThere := toSparseHash(hash)
	fmt.Println(hiThere)
}

func toSparseHash(hash []int) string {
	var hashes []int
	for len(hash) > 0 {
		cur := hash[:16]
		hash = hash[16:]
		var block int
		for _, n := range cur {
			if block == 0 {
				block = n
			} else {
				block = block ^ n
			}
		}
		hashes = append(hashes, block)
	}
	var s string
	for _, h := range hashes {
		s += fmt.Sprintf("%02x", h)
	}

	return s
}

func partOne() {
	var pos, skip = 0, 0
	//var input = []int{83, 0, 193, 1, 254, 237, 187, 40, 88, 27, 2, 255, 149, 29, 42, 100}
	//var max = 256
	var input = []int{3, 4, 1, 5}
	var max = 5
	hash := getSeed(max)
	for _, n := range input {
		if n == 0 {
			pos += skip
			if pos > max {
				pos = pos % max
			}
			skip++
			continue
		}
		chain := make([]int, max)
		copy(chain, hash)
		tmp := []int{}
		if pos+n > max {
			end := chain[pos:]
			endLen := len(end)
			start := chain[:(pos+n)%max]
			startLen := len(start)
			tmp = append(end, start...)
			tmp = reverse(tmp)
			end = tmp[:endLen]
			start = tmp[endLen:]
			var dt []int
			if len(start)+len(end) == max {
				dt = append(start, end...)
			} else {
				dt = append(start, append(hash[startLen:len(hash)-endLen], end...)...)
			}
			hash = dt
		} else {
			chain = chain[pos : (pos+n)%max]
			if len(chain) != 1 {
				chain = reverse(chain)
				tmp = append(hash[:pos], append(chain, hash[n+pos:]...)...)
				hash = tmp
			}
		}
		pos += n + skip
		if pos > max {
			pos = pos % max
		}
		skip++
		if len(hash) != max {
			log.Fatal("FUCK", n)
		}
	}
	fmt.Println(hash)
}

func getSeed(length int) []int {
	var seed []int
	for i := 0; i < length; i++ {
		seed = append(seed, i)
	}
	return seed
}

func reverse(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
