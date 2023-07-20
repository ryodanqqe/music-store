package main

import "errors"

type MemoryStorage struct {
	albums []album
}

type Storage interface {
	Create(album) album
	Read() []album
	ReadOne(id string) (album, error)
	Update(id string, a album) (album, error)
	Delete(id string) error
}

func (s *MemoryStorage) Create(am album) {
	s.albums = append(s.albums, am)
}

func (s MemoryStorage) Read() []album {
	return s.albums
}

func (s MemoryStorage) ReadOne(id string) (album, error) {
	for _, a := range s.albums {
		if a.ID == id {
			return a, nil
		}
	}
	return album{}, errors.New("not found")
}

func (s MemoryStorage) Update(id string, newAlbum album) (album, error) {
	for i := range s.albums {
		if s.albums[i].ID == id {
			s.albums[i] = newAlbum
		}
	}
	return album{}, errors.New("not found")
}

func (s *MemoryStorage) Delete(id string) error {
	for i := range s.albums {
		if s.albums[i].ID == id {
			s.albums = append(s.albums[:i], s.albums[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func NewStorage() MemoryStorage {
	return MemoryStorage{}
}
