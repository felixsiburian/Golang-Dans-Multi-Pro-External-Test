package tools

import (
	"Golang-Dans-Multi-Pro-External-Test/service"
	"Golang-Dans-Multi-Pro-External-Test/service/model"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

type tokenRepository struct {
	toolRepo service.IToolsRepository
}

func NewTokenRepository(
	toolRepo service.IToolsRepository,
) service.ITokenRepository {
	return tokenRepository{
		toolRepo: toolRepo,
	}
}

func (t tokenRepository) CreateUserToken(args model.TokenArgs) (res model.LoginResponses, err error) {
	res.AccessTokenExpires = time.Now().Add(time.Hour * 24).Unix()
	res.AccessTokenUUID = t.toolRepo.GUID()

	res.RefreshTokenExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	res.RefreshTokenUUID = t.toolRepo.GUID()

	// Create Access Token
	{
		jwtClaims := jwt.MapClaims{}
		jwtClaims["authorized"] = true
		jwtClaims["access_uuid"] = res.AccessTokenUUID
		jwtClaims["id"] = args.ID
		jwtClaims["fullname"] = args.Name
		jwtClaims["email"] = args.Email
		jwtClaims["exp"] = res.AccessTokenExpires

		accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
		res.AccessToken, err = accessToken.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
		if err != nil {
			log.Print(err)
			err = errors.New("[repository][token] while signed string token")
			return res, err
		}
	}

	// Create Refresh Token
	{
		rtClaims := jwt.MapClaims{}
		rtClaims["refresh_uuid"] = res.RefreshTokenUUID
		rtClaims["id"] = args.ID
		rtClaims["fullname"] = args.Name
		rtClaims["email"] = args.Email
		rtClaims["exp"] = res.RefreshTokenExpires

		refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
		res.RefreshToken, err = refreshToken.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
		if err != nil {
			log.Print(err)
			err = errors.New("[repository][token] while signed string refresh token")
			return res, err
		}
	}

	return res, err
}
