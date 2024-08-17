package tool

import (
	"math/rand"
	"time"
)

const (
	KcRandKindNum   int = iota // 纯数字
	KcRandKindLower            // 小写字母
	KcRandKindUpper            // 大写字母
	KcRandKindAll              // 数字、大小写字母
)

// 随机字符串
func Krand(size, kind int) string {
	ikind, kinds, result := kind, [][]int{{10, 48}, {26, 97}, {26, 65}}, make([]byte, size)
	isAll := kind > 2 || kind < 0
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < size; i++ {
		if isAll { // random ikind
			ikind = r.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return string(result)
}
