package types

import (
	"strconv"

	"github.com/taoqun8316/goblog/pkg/logger"
)

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToUint64(str string) uint64 {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.LogError(err)
	}
	return i
}
