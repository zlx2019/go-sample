// @Title unit.go
// @Description 单位常量
// @Author Zero - 2024/7/27 19:29:00

package constant

import "time"

const (
	BYTE = 1 << iota
	KB   = 1 << (10 * iota)
	MB
	GB
	TB
	PB

	Day   = time.Hour * 24
	Month = Day * 30
	Year  = Month * 12
)
