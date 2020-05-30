package params


import (
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var escapeRegex = regexp.MustCompile(`,|'|"|;`)

func GetStringArgs(c *gin.Context, key string) string {
	safeKey := escapeRegex.ReplaceAllString(c.Query(key), " ")
	return strings.TrimSpace(safeKey)
}

func GetIntArgs(c *gin.Context, key string) int {
	ret, _ := strconv.Atoi(c.Query(key))
	if 0 == ret {
		ret = 0
	}

	return ret
}

func GetInt64Args(c *gin.Context, key string) int64 {
	return int64(GetIntArgs(c, key))
}

func GetUint64Args(c *gin.Context, key string) uint64 {
	return uint64(GetIntArgs(c, key))
}
