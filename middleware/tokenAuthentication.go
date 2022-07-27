package middleware

//TokenAuthMiddleware check if request has valid token
// func TokenAuthMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		err := util.TokenValid(c.Request)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, err.Error())
// 			c.Abort()
// 			return
// 		}

// 		tokenString := util.ExtractToken(c.Request)

// 		tokenAuth, err := util.ExtractTokenMetadata(tokenString)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, "unauthorized")
// 			c.Abort()
// 			return
// 		}

// 		_, err = util.FetchAuth(tokenAuth)
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, "unauthorized")
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }
