package psql

// DeleteSingleRecord
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func DeleteSingleRecord(dsn string, blankSchema interface{}, recordToDelete interface{}) (int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return -1, err
	}

	query := db.Model(blankSchema).Where(recordToDelete).Delete(&recordToDelete)
	return query.RowsAffected, query.Error
}

// DeleteSingleRecordID
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func DeleteSingleRecordID(dsn string, blankSchema interface{}, id string) (int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return -1, err
	}

	query := db.Delete(blankSchema, id)
	return query.RowsAffected, query.Error
}

// DeleteManyRecords
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func DeleteManyRecords(dsn string, blankSchema interface{}, queryCondition QueryCondition) (int64, error) {
	db, err := connect(dsn)
	if err != nil {
		return -1, err
	}

	queryString, args, err := createWhereClause(queryCondition)
	if err != nil {
		return -1, err
	}
	query := db.Model(blankSchema).Where(queryString, args...).Delete(blankSchema)
	return query.RowsAffected, query.Error
}
