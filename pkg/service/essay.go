package service

import (
	"github.com/Marityr/gopitman"
	"github.com/Marityr/gopitman/pkg/repository"
)

type EssayService struct {
	repo repository.Essay
}

func NewEssayService(repo repository.Essay) *EssayService {
	return &EssayService{repo: repo}
}

func (s *EssayService) Create(input gopitman.Essay) error {
	return s.repo.Create(input)
}
