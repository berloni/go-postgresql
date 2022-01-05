package postgresql

import "errors"

// QueryCondition
type QueryCondition struct {
	AND     []AndCondition
	LIKE    []LikeCondition
	BETWEEN []BetweenCondition
}

type AndCondition struct {
	Column   string
	Operator string
	Value    string
}

type LikeCondition struct {
	Column string
	Value  string
}

type BetweenCondition struct {
	Column string
	Value1 string
	Value2 string
}

// QueryParameters
type QueryParameters struct {
	OrderBy   string
	OrderSort string
	Limit     int
}

// createWhereClause takes the query conditions specified and returns a dinamic query that can be used by gorm
func createWhereClause(queryCondition QueryCondition) (string, []interface{}, error) {
	var query string
	var args []interface{}

	and := queryCondition.AND
	like := queryCondition.LIKE
	between := queryCondition.BETWEEN

	if len(and) > 0 {
		for i, v := range and {
			// checks if the fields are empty
			if v.Column == "" || v.Operator == "" || v.Value == "" {
				return "", nil, errors.New("query: all the query fields are required")
			}
			if i > 0 {
				query += "AND "
			}
			query += v.Column + " " + v.Operator + " ? "
			args = append(args, v.Value)
		}
	}
	if len(like) > 0 {
		for i, v := range like {
			// checks if the fields are empty
			if v.Column == "" || v.Value == "" {
				return "", nil, errors.New("query: all the query fields are required")
			}
			if i == 0 && query != "" {
				query += "AND "
			}
			if i > 0 {
				query += "AND "
			}
			query += v.Column + " LIKE ? "
			args = append(args, "%"+v.Value+"%")
		}
	}
	if len(between) > 0 {
		for i, v := range between {
			// checks if the fields are empty
			if v.Column == "" || v.Value1 == "" || v.Value2 == "" {
				return "", nil, errors.New("query: all the query fields are required")
			}
			if i == 0 && query != "" {
				query += "AND "
			}
			if i > 0 {
				query += "AND "
			}
			query += v.Column + " BETWEEN ? AND ? "
			args = append(args, v.Value1, v.Value2)
		}
	}

	if query == "" {
		return "", nil, errors.New("query: at least one query condition is required")
	}

	return query, args, nil
}
