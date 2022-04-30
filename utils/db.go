package utils

import "database/sql"

func ConnectToDB() (connection *sql.DB, err error) {

	var dbURI = "habiba:123456789@tcp(localhost:3306)/chama"

	connection, err = sql.Open("mysql", dbURI)
	if err != nil {
		return
	}

	// it is advisable to ping
	err = connection.Ping()
	if err != nil {
		return
	}

	return

}
