package models

import (
	"database/sql"
	"log"

	"github.com/innv8/api-introduction/entities"
)

func GetUsers(db *sql.DB) (users []entities.User, err error) {
	// lets assume it gets a list of users from the database and returns it
	var query = "SELECT id, name, created FROM user"

	// to select we do this
	rows, err := db.Query(query)
	if err != nil {
		return
	}

	// here we have something from the db
	// we loop through the rows buffer

	/*
		 there are three types of linear data in Go


		 1. Arrays

		 var ages = []int {1,2,3}
		 for i, age := range ages {
			 fmt.Println(i, age)
		 }

		 for i := 0; i < len(ages); i++ {
			  fmt.Println(i, age)
		 }

		 2. Channels
		 These are used in goroutines. (jobs and results)
		 They are read and written by goroutines.
		 You loop through them this way

		 var jobs = make(chan string, 10)

		 // inside a goroutine
		 for j := range jobs {

		 }

		 3. Stream
		 This is kind of an array that we don't know the length.
		 Think of it as a linked list.
		 You read the next value in a while loop.
		 So in the first loop you will get the first item,
		 in the second loop you will get the second item and so on until we reach the end.

		 rows, err := db.Query(query)
		 for rows.Next() {
			 // in each loop we will get the next value until we reach the end
			 // and then the loop will end.
		 }

	*/

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.Id, &user.Name, &user.Created)
		if err != nil {
			log.Println("unable to read user : ", err)
			return
		}

		// here we've read the user details, so append to the array of users
		users = append(users, user)
	}

	return users, nil
}
