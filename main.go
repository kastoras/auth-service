package main

func main() {

	dbUrl := "postgres://postgres:postgres@db:5432/restgolang"
	db, err := getDB(dbUrl)
	if err != nil {
		panic("Failed to connect to the db")
	}

	address := ":3030"
	server := NewAPIServer(address, db)
	server.Run()
}
