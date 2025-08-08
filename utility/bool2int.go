package utility

import (
	_ "time/tzdata"
)

func Bool2Int(b bool) int {
	if b {
		return 1
	}
	return 0
}
