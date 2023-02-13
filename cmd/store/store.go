package store

import "errors"

type Store map[string]string

var (
	ErrKeyNotFound = errors.New("Key not found.")
)

func NewStore() Store {
	s := make(Store)
	return s
}

func (s Store) Get(key string) (string, error) {
	val, ok := s[key]

	if !ok {
		return "", ErrKeyNotFound
	}

	return val, nil
}

func (s Store) Set(key, val string) {
	s[key] = val
}

func (s Store) Del(key string) {
	delete(s, key)
}
