package main

import (
	"errors"
	"sync"
)

// an in-memory storage for users
type Storage struct {
	mu sync.Mutex
	users map[string]User
}

// Initialize the storage
func NewStorage() *Storage {
	return &Storage{
		users: make(map[string]User),
	}
}

// create a new user
func (s *Storage) CreateUser(user User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[user.ID]; exists {
		return errors.New("User already exists")
	}
	s.users[user.ID] = user
	return nil
}

// retrieve a user by the userId
func (s *Storage) GetUser(id string) (User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, exists := s.users[id]
	if !exists {
		return User{}, errors.New("User not found")
	}
	return user, nil
}

// update an existing user instance
func (s *Storage) UpdateUser(id string, updated User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("Users not found")
	}
	s.users[id] = updated
	return nil
}

// remove a user according to the userId
func (s *Storage) DeleteUser(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.users[id]; !exists {
		return errors.New("User not found")
	}
	delete(s.users, id)
	return nil
}