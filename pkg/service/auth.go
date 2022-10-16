package service

import (
	"crypto/sha1"
	"fmt"

	rest "github.com/serdfxe/rest-api-golang/pkg"
	"github.com/serdfxe/rest-api-golang/pkg/repository"
)

const salt = "oiygefg3p97rt237grwurbgpvuwyeoug"

type AuthService struct {
	repo repository.Authorization
}

func newAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user rest.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
