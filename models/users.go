package models

import "github.com/innv8/api-introduction/entities"

func GetUsers() (users []entities.User, err error) {
	// lets assume it gets a list of users from the database and returns it
	users = []entities.User{
		{Name: "Habiba", Age: 24},
		{Name: "Sam", Age: 25},
		{Name: "Denno", Age: 27},
		{Name: "Maxine", Age: 24},
	}
	return users, nil
}
