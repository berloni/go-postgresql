package postgresql

import "errors"

type ConnectionParams struct {
	Host     string
	User     string
	Password string
	DBName   string
	Port     string
}

// GenerateDsn generates the gorm's connection string based on the ConnectionParams
// Example of connection string (dsn): "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
func GenerateDsn(connectionParams ConnectionParams) (string, error) {
	err := checkConnectionParams(connectionParams)
	if err != nil {
		return "", err
	}

	dsn := "host=" + connectionParams.Host +
		" user=" + connectionParams.User +
		" password=" + connectionParams.Password +
		" dbname=" + connectionParams.DBName +
		" port=" + connectionParams.Port +
		" sslmode=disable"
	return dsn, nil
}

func checkConnectionParams(connectionParams ConnectionParams) error {
	if connectionParams.Host == "" {
		return errors.New("host is required")
	}

	if connectionParams.User == "" {
		return errors.New("user is required")
	}

	if connectionParams.Password == "" {
		return errors.New("password is required")
	}

	if connectionParams.DBName == "" {
		return errors.New("dbName is required")
	}

	if connectionParams.Port == "" {
		return errors.New("port is required")
	}

	return nil
}
