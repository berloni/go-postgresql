package psql

// Create
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func Create(dsn string, blankSchema interface{}, values interface{}) error {
	db, err := connect(dsn)
	if err != nil {
		return err
	}

	err = db.AutoMigrate(blankSchema)
	if err != nil {
		return err
	}

	err = db.Create(values).Error
	if err != nil {
		return err
	}

	return nil
}
