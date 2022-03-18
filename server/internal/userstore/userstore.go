package userstore

import "sync"



type User struct {
	Key      int `json:"key"`
	Username string `json:"username"`
	Url      string `json:"url"`
}

type UserStore struct {
	sync.Mutex

	users map[int]User
	nextKey int
}


func New() *UserStore {
	us := &UserStore{}
	us.users = make(map[int]User)
	us.nextKey = 0
	return us
}