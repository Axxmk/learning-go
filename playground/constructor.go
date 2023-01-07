package main

import "fmt"

// Service is a component
type Service struct {
	CommentRepo CommentRepo
}

// CommentRepo These dependencies should be modelled as interfaces to help ensure the components are loosely coupled
type CommentRepo interface {
	GetComments() ([]Comment, error)
}

type Comment struct {
	Author string
	Body   string
	Slug   string
}

// NewService is a constructor function
func NewService(cmtRepo CommentRepo) (*Service, error) {
	svc := &Service{
		CommentRepo: cmtRepo,
	}

	// handles other potentially more complex setup logic

	return svc, nil
}

func (s *Service) GetAllComments() error {
	comments, err := s.CommentRepo.GetComments()
	if err != nil {
		return err
	}

	fmt.Println(comments)
	return nil
}
