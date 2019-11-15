package article

import (
	"FirstGoProject/model"
)

// Store
type Store interface {
	// CreateArticle(*model.Article) error
	Create(*model.Origin) error
	Update(*model.Origin) error
}
