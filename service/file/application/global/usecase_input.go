package global

import (
	"strconv"
)

type FindAllInput struct {
	QueryKey   string
	QueryValue []interface{}
	Limit      int
	Offset     int
	Sort       string
}

func (t *FindAllInput) ParseValue(str []string) {
	for i, _ := range str {
		if val, err := strconv.Atoi(str[i]); err == nil {
			t.QueryValue = append(t.QueryValue, val)
		}

		if val, err := strconv.ParseBool(str[i]); err == nil {
			t.QueryValue = append(t.QueryValue, val)
		} else {
			t.QueryValue = append(t.QueryValue, str[i])
		}
	}
}
