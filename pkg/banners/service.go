package banners

import (
	"context"
	"errors"
	"sync"
)

// Service представляют собой сервис по управлению баннерами.
type Service struct {
	mu    sync.RWMutex
	items []*Banner
}

// NewService создаёт сервис.
func NewService() *Service {
	return &Service{items: make([]*Banner, 0)}
}

// Banner представляет собой баннер.
type Banner struct {
	ID      int64
	Title   string
	Content string
	Button  string
	Link    string
}

// All возвращает все существующие баннеры.
func (s *Service) All(ctx context.Context) ([]*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.items, nil
}

// ByID возвращает баннер по идентификатору.
func (s *Service) ByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, banner := range s.items {
		if banner.ID == id {
			return banner, nil
		}
	}
	return nil, errors.New("item not found")
}

// Save сохраняет/обновляет баннер.
func (s *Service) Save(ctx context.Context, item *Banner) (*Banner, error) {

	s.items = append(s.items, item)
	// return item, errors.New("invalid Save")
	return item, nil

}

// RemoveByID удаляет баннер по идентификатору.
func (s *Service) RemoveByID(ctx context.Context, id int64) (*Banner, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for i, banner := range s.items {
		if banner.ID == id {
			s.items = append(s.items[:i], s.items[i+1:]...)
			return banner, nil
		}
	}
	return nil, errors.New("item not found")
}
