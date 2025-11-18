package day_6_7

import "errors"

type undoStack struct {
	data []string
}

func (s *undoStack) push(msg string) {
	s.data = append(s.data, msg)
}

func (s *undoStack) pop() (string, error) {
	if len(s.data) == 0 {
		return "", errors.New("nothing to undo")
	}
	last := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return last, nil
}
