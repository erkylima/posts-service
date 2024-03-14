package services

import (
	"github.com/erkylima/posts-service/internal/core/domains"
	"github.com/erkylima/posts-service/internal/core/ports"
)

type PageService struct {
	repo ports.Repository
}

func NewPageService(repo ports.Repository) *PageService {
	return &PageService{repo: repo}
}

func (ps *PageService) CreatePage(page *domains.Page) (string, error) {
	result, err := ps.repo.Create(page)
	if err != nil {
		return "", err
	}
	return result, nil
}

func (ps *PageService) ReadPage(slug string) (*domains.Page, error) {
	service := &domains.Page{}
	page, err := ps.repo.Read(slug, service)
	if err != nil {
		return nil, err
	}

	return page.(*domains.Page), nil
}

func (ps *PageService) UpdatePage(page domains.Page) (*domains.Page, error) {
	// ps.repo.UpdatePage(page)
	return nil, nil
}

func (ps *PageService) DeletePage(slug string) error {
	// ps.repo.DeletePage(slug)
	return nil
}

func (ps *PageService) ListPages() ([]domains.Page, error) {
	// ps.repo.ListPages()
	return nil, nil
}

func (ps *PageService) ListPagesByCategory(category string) ([]domains.Page, error) {
	// ps.repo.ListPagesByCategory(category)
	return nil, nil
}

func (ps *PageService) ListPagesByTag(tag string) ([]domains.Page, error) {
	// ps.repo.ListPagesByTag(tag)
	return nil, nil
}

func (ps *PageService) ListPagesByAuthor(author string) ([]domains.Page, error) {
	// ps.repo.ListPagesByAuthor(author)
	return nil, nil
}
