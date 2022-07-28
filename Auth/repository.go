package auth

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"fiber-simple-api/Auth/types"

	"github.com/golang-jwt/jwt"
)

//VerifyToken check if a token is valid
func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

//TokenValid check if token is expired
func TokenValid(r *http.Request) error {
	tokenString := ExtractToken(r)

	token, err := VerifyToken(tokenString)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

//ExtractTokenMetadata decrypt the token content
func ExtractTokenMetadata(tokenString string) (*types.AccessDetails, error) {
	token, err := VerifyToken(tokenString)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}

		userID, err := strconv.ParseInt(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}

		return &types.AccessDetails{
			AccessUuid: accessUUID,
			UserId:     userID,
		}, nil
	}

	return nil, err
}

//ExtractToken returns the request authentication token
func ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}
