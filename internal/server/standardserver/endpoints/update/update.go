package update

import (
	handler "github.com/Vackhan/metrics/internal/server/pkg/httphandlers/update"
	"github.com/Vackhan/metrics/internal/server/pkg/storage"
)

type Update struct {
	repo storage.UpdateRepo
}

func (u *Update) GetURL() string {
	return "/update/"
}
func (u *Update) GetFunctionality() any {
	return handler.New(u.repo)
}

func NewUpdateEndpoint(repo storage.UpdateRepo) *Update {
	return &Update{repo}
}
