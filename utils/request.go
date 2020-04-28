package utils

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FindAllRequest struct {
	Limit      int64
	Offset     int64
	Sort       string
	QueryKey   string
	QueryValue []string
}

func QueryBuilder(c *gin.Context) *FindAllRequest {
	findAllRequest := FindAllRequest{
		Sort:   DEFAULT_SORT_FIND_ALL_REQUEST,
		Limit:  30,
		Offset: 0,
	}

	if sort, isExist := c.GetQuery(SORT_ID_FIND_ALL_REQUEST); isExist {
		findAllRequest.Sort = sort
	}

	if limit, isExist := c.GetQuery(LIMIT_ID_FIND_ALL_REQUEST); isExist {
		findAllRequest.Limit, _ = strconv.ParseInt(limit, 10, 64)
	}

	if offset, isExist := c.GetQuery(OFFSET_ID_FIND_ALL_REQUEST); isExist {
		findAllRequest.Offset, _ = strconv.ParseInt(offset, 10, 64)
	}

	if where, isExist := c.GetQuery(QUERY_ID_FIND_ALL_REQUEST); isExist {
		var query interface{}
		var queryKey []string
		var queryValue []string

		json.Unmarshal([]byte(where), &query)
		if query != nil {
			for i, _ := range query.(map[string]interface{}) {
				val := query.(map[string]interface{})[i]

				if val, ok := val.(string); ok {
					queryValue = append(queryValue, val)
					queryKey = append(queryKey, i+" ?")
				}
			}

			findAllRequest.QueryKey = strings.Join(queryKey[:], " AND ")
			findAllRequest.QueryValue = queryValue
		}
	}

	return &findAllRequest
}
