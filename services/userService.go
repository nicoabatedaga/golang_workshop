package services

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/nicoabatedaga/golang_workshop/models"
	"github.com/nicoabatedaga/golang_workshop/storage"
)

type UserService interface {
	GetUser(id string) (*models.User, error)
	PostUser(user models.User) (*models.User, error)
	DeleteUser(id string) (*models.User, error)
	PutUser(id string, user models.User) (*models.User, error)
}

type UserServiceImp struct {
	storage   storage.StorageInterface
	partition string
}

func NewUserService(storage storage.StorageInterface) UserService {
	return &UserServiceImp{
		storage,
		"users",
	}
}

func (u *UserServiceImp) GetUser(id string) (*models.User, error) {
	rsp, err := u.storage.Get(u.partition, id)
	if err != nil {
		return nil, err
	}
	var user models.User
	if err := json.Unmarshal(rsp, &user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServiceImp) PostUser(user models.User) (*models.User, error) {
	// generate an uuid
	user.ID = uuid.New().String()
	userByte, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}
	if err = u.storage.Save(u.partition, user.ID, userByte); err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserServiceImp) DeleteUser(id string) (*models.User, error) {
	return nil, nil
}

func (u *UserServiceImp) PutUser(id string, user models.User) (*models.User, error) {
	return nil, nil
}
