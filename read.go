package postgresql

// GetAllRecords
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func GetAllRecords(dsn string, blankSchema interface{}) ([]map[string]interface{}, int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return nil, 0, err
	}

	result := []map[string]interface{}{}
	query := db.Model(blankSchema).Find(&result)
	return result, query.RowsAffected, query.Error
}

// GetRecords
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func GetRecords(dsn string, blankSchema interface{}, queryCondition QueryCondition, queryParameters ...QueryParameters) ([]map[string]interface{}, int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return nil, 0, err
	}

	orderBy := "id"
	orderSort := "asc"
	limit := -1
	if len(queryParameters) == 1 {
		if queryParameters[0].OrderBy != "" {
			orderBy = queryParameters[0].OrderBy
		}
		if queryParameters[0].OrderSort != "" {
			orderSort = queryParameters[0].OrderSort
		}
		if queryParameters[0].Limit != 0 {
			limit = queryParameters[0].Limit
		}
	}

	result := []map[string]interface{}{}
	queryString, args, err := createWhereClause(queryCondition)
	if err != nil {
		return nil, 0, err
	}
	query := db.Model(blankSchema).Where(queryString, args...).Order(orderBy + " " + orderSort).Limit(limit).Find(&result)
	return result, query.RowsAffected, query.Error
}

// GetRecordsStrcut
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func GetRecordsStruct(dsn string, blankSchema interface{}, filters interface{}, queryParameters ...QueryParameters) ([]map[string]interface{}, int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return nil, 0, err
	}

	orderBy := "id"
	orderSort := "asc"
	limit := -1
	if len(queryParameters) == 1 {
		if queryParameters[0].OrderBy != "" {
			orderBy = queryParameters[0].OrderBy
		}
		if queryParameters[0].OrderSort != "" {
			orderSort = queryParameters[0].OrderSort
		}
		if queryParameters[0].Limit != 0 {
			limit = queryParameters[0].Limit
		}
	}

	result := []map[string]interface{}{}
	query := db.Model(blankSchema).Where(filters).Order(orderBy + " " + orderSort).Limit(limit).Find(&result)
	return result, query.RowsAffected, query.Error
}
