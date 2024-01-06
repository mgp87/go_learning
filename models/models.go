package models

import "time"

type User struct {
	Id        int
	Name      string
	createdAt time.Time
	status    bool
}

func (user *User) AddUser(id int, name string, createdAt time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.createdAt = createdAt
	user.status = status
}
