package pkg

import (
	"errors"
	"slices"
)

type PackageService struct {
	entities []Package
}

var ErrInvalidIdx = errors.New("invalid idx")

const pageSize = 5

func (s *PackageService) isValidIdx(idx int) bool {
	if idx < 0 || idx >= len(s.entities) {
		return false
	}
	return true
}

func NewPackageService() *PackageService {
	return &PackageService{
		entities: []Package{
			{"item1"},
			{"item2"},
			{"item3"},
			{"item4"},
			{"item5"},
			{"item6"},
			{"item7"},
			{"item8"},
			{"item9"},
			{"item10"},
			{"item11"},
			{"item12"},
			{"item13"},
			{"item14"},
			{"item15"},
			{"item16"},
			{"item17"},
		},
	}
}

func (s *PackageService) List() []Package {
	return s.entities
}

func (s *PackageService) Get(idx int) (*Package, error) {
	if !s.isValidIdx(idx) {
		return nil, ErrInvalidIdx
	}
	return &s.entities[idx], nil
}

func (s *PackageService) Edit(idx int, value string) error {
	if !s.isValidIdx(idx) {
		return ErrInvalidIdx
	}
	s.entities[idx] = Package{value}

	return nil
}

func (s *PackageService) Delete(idx int) error {
	if !s.isValidIdx(idx) {
		return ErrInvalidIdx
	}
	s.entities = slices.Delete(s.entities, idx, idx+1)
	return nil
}

func (s *PackageService) New(value string) error {
	s.entities = append(s.entities, Package{value})
	return nil
}
