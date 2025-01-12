package tool

import (
	"database/sql"
	"strconv"
	"strings"
)

func BuildQuery(ids []int64) string {
	len := len(ids)

	var builder strings.Builder
	builder.WriteString("(")
	for i := 0; i < len; i++ {
		if i > 0 {
			builder.WriteString(", ")
		}
		builder.WriteString(strconv.FormatInt(ids[i], 10))
	}
	builder.WriteString(")")

	return builder.String()
}

func GenMediaURL(value sql.NullString, baseURL string) string {
	if !value.Valid {
		return ""
	}
	return baseURL + value.String
}
