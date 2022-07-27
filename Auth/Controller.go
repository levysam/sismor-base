package auth

import (
	"strconv"
	"time"

	model "fiber-simple-api/Domains/Users"
	db "fiber-simple-api/database"
	"fiber-simple-api/util"

	"github.com/gofiber/fiber/v2"
)

//Login make the user Authentication
func Login(c *fiber.Ctx) error {
	var u model.User

	if err := c.BodyParser(u); err != nil {
		c.JSON("erro ao passar para o model user")
		return err
	}

	var user, err = model.GetUserByEmail(u.Email)

	//compare the user from the request, with the one we defined:
	if user.Email != u.Email || user.Password != u.Password {
		c.JSON(user)
		return err
	}

	//Create a jwt token with user ID that will expire in 15 minutes
	td, err := util.CreateJwtToken(uint64(user.Id))
	if err != nil {
		c.JSON(err.Error())
		return err
	}

	saveErr := createAuth(uint64(user.Id), td)
	if saveErr != nil {
		c.JSON(saveErr.Error())
		return saveErr
	}

	tokens := map[string]string{
		"access_token":  td.AccessToken,
		"refresh_token": td.RefreshToken,
	}

	c.JSON(tokens)
	return nil
}

func createAuth(userid uint64, td *model.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	client := db.GetRedisClient()
	defer client.Close()

	errAccess := client.Set(td.AccessUuid, strconv.Itoa(int(userid)), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}

	errRefresh := client.Set(td.RefreshUuid, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

//Logout erase the user session
// func Logout(c *fiber.Ctx) {
// 	tokenString := util.ExtractToken(c.Request)

// 	au, err := util.ExtractTokenMetadata(tokenString)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	deleted, delErr := util.DeleteAuth(au.AccessUuid)
// 	if delErr != nil || deleted == 0 { //if any goes wrong
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	c.JSON(http.StatusOK, "Successfully logged out")
// }
