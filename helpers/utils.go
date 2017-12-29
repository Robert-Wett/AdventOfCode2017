package helpers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func GetInput(fp string) string {
	content, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func KnotHash(input string) string {
	var pos, skip = 0, 0
	var max = 256
	var ints []int
	var fixedLenthVals = []int{17, 31, 73, 47, 23}
	for _, i := range input {
		ints = append(ints, int(i))
	}
	for _, i := range fixedLenthVals {
		ints = append(ints, int(i))
	}

	var hash []int
	for i := 0; i < max; i++ {
		hash = append(hash, i)
	}

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

	hi := toSparseHash(hash)
	return hi
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

func reverse(a []int) []int {
	for left, right := 0, len(a)-1; left < right; left, right = left+1, right-1 {
		a[left], a[right] = a[right], a[left]
	}
	return a
}
func IndexOf(needle int, haystack []int) int {
	for i, v := range haystack {
		if v == needle {
			return i
		}
	}
	return -1
}
