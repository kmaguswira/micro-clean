package utils

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type FindAllRequest struct {
	Limit      int32
	Offset     int32
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
		tmp, _ := strconv.ParseInt(limit, 10, 32)
		findAllRequest.Limit = int32(tmp)
	}

	if offset, isExist := c.GetQuery(OFFSET_ID_FIND_ALL_REQUEST); isExist {
		tmp, _ := strconv.ParseInt(offset, 10, 32)
		findAllRequest.Offset = int32(tmp)
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
