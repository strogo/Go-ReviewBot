package user

import (
	"FirstGoProject/model"
)

// type Store interface {
// 	GetByID(uint) (*model.User, error)
// 	GetByEmail(string) (*model.User, error)
// 	GetByUsername(string) (*model.User, error)
// 	Create(*model.User) error
// 	Update(*model.User) error
// 	AddFollower(user *model.User, followerID uint) error
// 	RemoveFollower(user *model.User, followerID uint) error
// 	IsFollower(userID, followerID uint) (bool, error)
// }

// Store interface for Champ
type Store interface {
	GetByName(string) (*model.Champ, error)
	Create(*model.Champ) error
	Update(*model.Champ) error
}
