package entities

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Created string `json:"created"`
}

// define the structure Member according to the members table
/*
+-------------+-------------+------+-----+---------------------+-------------------------------+
| Field       | Type        | Null | Key | Default             | Extra                         |
+-------------+-------------+------+-----+---------------------+-------------------------------+
| id          | int(11)     | NO   | PRI | NULL                | auto_increment                |
| name        | varchar(30) | YES  | UNI | NULL                |                               |
| parentsName | varchar(30) | NO   |     | NULL                |                               |
| phone       | varchar(12) | NO   | UNI | NULL                |                               |
| position    | int(11)     | YES  | MUL | NULL                |                               |
| created     | datetime    | YES  | MUL | current_timestamp() |                               |
| modified    | datetime    | YES  |     | NULL                | on update current_timestamp() |
| nextOfKin   | varchar(30) | NO   |     | NULL                |                               |
+-------------+-------------+------+-----+---------------------+-------------------------------+
*/

type Member struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	ParentsName string `json:"parentsName"`
	Phone       string `json:"phone"`
	Position    int    `json:"position"`
	Created     string `json:"created"`
	Modified    string `json:"modified,omitempty"` // if dealing with json, you can use omitempty to omit empty values
	NextOfKin   string `json:"nextOfKin"`
}
