package models

import (
	"database/sql"
	"log"

	"github.com/innv8/api-introduction/entities"
)

/*
func GetUsers(db *sql.DB) (users []entities.User, err error) {
	// lets assume it gets a list of users from the database and returns it
	var query = "SELECT id, name, created FROM user"

	// to select we do this
	rows, err := db.Query(query)
	if err != nil {
		log.Printf("unable to fetch members from db because %v", err)
		return
	}

	// here we have something from the db
	// we loop through the rows buffer


		//  there are three types of linear data in Go


		//  1. Arrays

		//  var ages = []int {1,2,3}
		//  for i, age := range ages {
		// 	 fmt.Println(i, age)
		//  }

		//  for i := 0; i < len(ages); i++ {
		// 	  fmt.Println(i, age)
		//  }

		//  2. Channels
		//  These are used in goroutines. (jobs and results)
		//  They are read and written by goroutines.
		//  You loop through them this way

		//  var jobs = make(chan string, 10)

		//  // inside a goroutine
		//  for j := range jobs {

		//  }

		//  3. Stream
		//  This is kind of an array that we don't know the length.
		//  Think of it as a linked list.
		//  You read the next value in a while loop.
		//  So in the first loop you will get the first item,
		//  in the second loop you will get the second item and so on until we reach the end.

		//  rows, err := db.Query(query)
		//  for rows.Next() {
		// 	 // in each loop we will get the next value until we reach the end
		// 	 // and then the loop will end.
		//  }



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

*/

func FetchMembers(db *sql.DB) (members []entities.Member, err error) {
	var fetchMembersQuery = `SELECT id, name, parentsName, phone, 
	position, created, nextOfKin, modified
	FROM members`
	rows, err := db.Query(fetchMembersQuery)
	if err != nil {
		log.Printf("unable to fetch members because %v", err)
		return
	}

	for rows.Next() {
		var member entities.Member
		// an important factor,
		// if a field is nullable, treat it as if it will always be null
		// because you can't scan a null value into a pointer
		// for modified (which is null in the db), we create a new variable
		var modified sql.NullString

		err = rows.Scan(
			&member.Id,
			&member.Name,
			&member.ParentsName,
			&member.Phone,
			&member.Position,
			&member.Created,
			&member.NextOfKin,
			&modified, // we are reading into an sql.NullString data type because this field can be null
		)
		if err != nil {
			log.Printf("unable to scan member because %v", err)
			return // or continue if you just want to skip this one
		}

		// after you are sure there is no error, read the nullable values into the member struct
		member.Modified = modified.String
		members = append(members, member)
	}

	log.Printf("found %d members", len(members))
	return
}

func FetchMembersByPosition(position int, db *sql.DB) (members []entities.Member, err error) {
	var fetchMembersQuery = `SELECT id, name, parentsName, phone, 
	position, created, nextOfKin, modified
	FROM members
	WHERE position = ?`

	// in queries we put variable arguments as ? and then in db.Query, we put the values after the query
	// in the order of question marks in the query
	rows, err := db.Query(fetchMembersQuery, position)
	if err != nil {
		log.Printf("unable to fetch members because %v", err)
		return
	}

	for rows.Next() {
		var member entities.Member
		// an important factor,
		// if a field is nullable, treat it as if it will always be null
		// because you can't scan a null value into a pointer
		// for modified (which is null in the db), we create a new variable
		var modified sql.NullString

		err = rows.Scan(
			&member.Id,
			&member.Name,
			&member.ParentsName,
			&member.Phone,
			&member.Position,
			&member.Created,
			&member.NextOfKin,
			&modified, // we are reading into an sql.NullString data type because this field can be null
		)
		if err != nil {
			log.Printf("unable to scan member because %v", err)
			return // or continue if you just want to skip this one
		}

		// after you are sure there is no error, read the nullable values into the member struct
		member.Modified = modified.String
		members = append(members, member)
	}

	log.Printf("found %d members", len(members))
	return
}
