package userstore

import (
	"fmt"
	"math/rand"
	"sync"
)

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

//TODO разобрать
func (us *UserStore) GetAllUsers() []User {
	us.Lock()
	defer us.Unlock()

	allUsers := make([]User, 0, len(us.users))
	for _, user := range us.users {
		allUsers = append(allUsers, user)
	}
	return allUsers
}

func New() *UserStore {
	us := &UserStore{}
	us.users = make(map[int]User)
	us.nextKey = rand.Intn(1000000)
	return us
}

func (us *UserStore) CreateUser(username string, url string ) int {
	us.Lock()
	defer us.Unlock()
	user := User{
		Key:  us.nextKey,
		Username: username,
		Url: url}

	us.users[us.nextKey] = user
	us.nextKey = rand.Intn(1000000)
	return user.Key
}

func (us *UserStore) DeleteUser(id int)  error {
	us.Lock()
	defer us.Unlock()

	if _, ok := us.users[id]; !ok {
		return fmt.Errorf("task with id=%d not found", id)
	}
	delete(us.users, id)
	return nil
}

