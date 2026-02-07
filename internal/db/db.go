package db

type Database struct {
	Name string
}

func NewDB() Database {
	return Database{Name: "I'm your database!"}
}
