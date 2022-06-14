package authservice

import (
	"boiler-plate/internal/services"
	"boiler-plate/internal/storage"
	"boiler-plate/models"
	"boiler-plate/pkg/global"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	storage storage.IStorage
}

func NewAuthService(storage storage.IStorage) services.IAuthService {
	return &AuthService{storage: storage}
}

func (a *AuthService) Login(userLoginDTO models.UserLoginDTO) (token string, err error) {
	user, err := a.storage.Users().FindByUsername(userLoginDTO.Username)
	if err != nil {
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginDTO.Password)); err != nil {
		return
	}

	return a.CreateToken(*user)
}

func (a *AuthService) Register(userRegisterDTO models.UserRegisterDTO) error {
	var user models.User

	if err := copier.Copy(&user, userRegisterDTO); err != nil {
		return err
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashPassword)

	if err := a.storage.Users().Create(&user); err != nil {
		return err
	}

	return nil
}

func (a *AuthService) CreateToken(user models.User) (tokenString string, err error) {
	expiryTime := time.Now().Add(time.Minute * time.Duration(global.TokenExpiryTime))

	claims := jwt.MapClaims{}
	newUUID := uuid.NewV4().String()
	claims["uuid"] = newUUID
	claims["exp"] = expiryTime.Unix()
	claims["user_id"] = user.ID

	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = newToken.SignedString([]byte(global.TokenSecretKey))
	if err != nil {
		return
	}

	newTokenDetail := models.TokenDetail{
		UUID:       newUUID,
		UserID:     user.ID,
		Token:      tokenString,
		ExpiryTime: expiryTime,
	}

	var currentTokenDetail models.TokenDetail

	if currentTokenDetail, err = a.storage.TokenDetails().GetByUserID(user.ID); err != nil {

		if err = a.storage.TokenDetails().Create(&newTokenDetail); err != nil {
			return
		}
		return
	}

	currentTokenDetail.UUID = newTokenDetail.UUID
	currentTokenDetail.UserID = newTokenDetail.UserID
	currentTokenDetail.Token = newTokenDetail.Token
	currentTokenDetail.ExpiryTime = newTokenDetail.ExpiryTime

	if err = a.storage.TokenDetails().Update(&currentTokenDetail); err != nil {
		return
	}

	return
}
