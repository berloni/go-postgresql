package psql

// UpdateWithStruct
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func UpdateWithStruct(dsn string, blankSchema interface{}, currentValues interface{}, updatedValues interface{}) (int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return -1, err
	}

	query := db.Model(blankSchema).Where(currentValues).Updates(updatedValues)
	return query.RowsAffected, query.Error
}

// UpdateWithClause
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func UpdateWithClause(dsn string, blankSchema interface{}, queryCondition QueryCondition, updatedValues interface{}) (int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return -1, err
	}

	queryString, args, err := createWhereClause(queryCondition)
	if err != nil {
		return -1, err
	}
	query := db.Model(blankSchema).Where(queryString, args...).Updates(updatedValues)
	return query.RowsAffected, query.Error
}
