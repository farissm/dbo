package helper

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string	`json:"username"`
	Password string	`json:"password"`
}

type LoginResponse struct {
	Code 		int		`json:"code"`
	AccessToken string	`json:"access_token"`
	TokenType 	string	`json:"token_type"`
	ExpiresIn 	int 	`json:"expires_in"`
}

type Response struct {
	Code	int		`json:"code"`
	Message	string	`json:"message"`
}

func LoginAPIResponse(accessToken string, expiresIn int, code int) LoginResponse {
	jsonResponse := LoginResponse{
		AccessToken: accessToken,
		TokenType: "Bearer",
		ExpiresIn: expiresIn,
		Code: code,
	}

	return jsonResponse
}

func APIResponse(code int, message string) Response {
	jsonResponse := Response{
		Code: code,
		Message: message,
	}

	return jsonResponse
}

func HashPassword(password string) ([]byte, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }
    return hashedPassword, nil
}


var APPLICATION_NAME = "DBO TEST"
var SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))

type JWTToken struct {
	SignedToken	string
	ExpiresIn 	float64
}

type MyClaims struct {
    jwt.StandardClaims
    Username string `json:"Username"`
}

func GenerateToken(username string) (JWTToken, error){
	expiresTime := time.Now().Add(time.Minute * 60)
	expires := expiresTime.Unix()

	timeNow := time.Now()
	diffTime := expiresTime.Sub(timeNow).Seconds()

	claims := MyClaims{
        StandardClaims: jwt.StandardClaims{
            Issuer:    APPLICATION_NAME,
            ExpiresAt: expires,
        },
        Username: username,
    }

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	generateTokenJwt, err := token.SignedString(SECRET_KEY)

    jwtToken := JWTToken{}    
    jwtToken.SignedToken = generateTokenJwt
    jwtToken.ExpiresIn = diffTime

	if err != nil {
		return jwtToken, err
	}

	return jwtToken, nil
}

func ValidateToken(encodedToken string) (*jwt.Token, error){
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(SECRET_KEY), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}