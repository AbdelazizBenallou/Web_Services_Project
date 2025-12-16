package domain

import "errors"

var ErrNotEnoughStock = errors.New("not enough stock")

func (s *Stock) CanReserve(qty int) error {
	if qty <= 0 {
		return ErrNotEnoughStock
	}
	if s.Quantity < qty {
		return ErrNotEnoughStock
	}
	return nil
}

func (s *Stock) Reserve(qty int) error {
	if err := s.CanReserve(qty); err != nil {
		return err
	}
	s.Quantity -= qty
	return nil
}

