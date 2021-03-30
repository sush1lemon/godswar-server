package utility

import (
	"bytes"
	"math"
	"strconv"
)

var key_shift = []int{0, 1, -1, -3, 2, 7, 1, 2, -2, 7, 2, 4, 1, 2, 7, 23, 2, 1, 2, 7, 1, 2, 2, 4, 2, 7, 15, -2, 2, 1, 11}
var key_shifi = []int{0, -1, -3, 3, -2, 3, -1, -2, 4, 3, -2, 2, -1, -2, 3, 3, -2, -1, -2, 3, -1, 0, 0, 2, -2, 3, -5, -6, -2, -1, 1}
var key_f = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var key_i = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func Parser(dt string) string {
	var fullData string
	tempData := bytes.Trim([]byte(dt), "\u0000")

	lng := len(tempData)
	for _, td := range tempData {
		dd := string(td)
		isNumber := false
		if _, err := strconv.Atoi(string(dd)); err == nil {
			isNumber = true
		}
		if !isNumber {
			pos := 0
			found := false
			for x := 0; x < len(key_f); x++ {
				if dd == key_f[x] {
					pos = x
					found = true
					break
				}
			}
			if found {
				key := key_shift[lng]
				bf := len(key_f) - pos
				var nwt string
				if (bf - key) <= 0 {
					find := bf - key
					if find < 0 {
						nff := math.Abs(float64(find))
						nwt = key_f[int(nff)]
					} else {
						nwt = key_f[find]
					}
				} else {
					nwt = key_f[pos+key]
				}
				fullData += nwt
			} else {
				fullData += dd
			}
		} else {
			pos := 0
			for x := 0; x < len(key_i); x++ {
				if dd == key_i[x] {
					pos = x
					break
				}
			}
			key := key_shifi[lng]
			bf := len(key_i) - pos
			var nwt string
			if (bf - key) <= 0 {
				find := bf - key
				if find < 0 {
					nff := math.Abs(float64(find))
					nwt = key_i[int(nff)]
				} else {
					nwt = key_i[find]
				}
			} else {
				nwt = key_i[pos+key]
			}
			fullData += nwt
		}
	}

	return fullData
}
