package enums

import "strings"

type OrderBy int

const (
	NoOrderBy OrderBy = iota
	Desc
	Asc
)

var stringToOrderBy = map[string]OrderBy{
	"desc": Desc,
	"asc":  Asc,
}

func ParseOrderBy(s string) OrderBy {
	if val, ok := stringToOrderBy[toLowerSafe(s)]; ok {
		return val
	}
	return NoOrderBy
}

func toLowerSafe(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s)
}
