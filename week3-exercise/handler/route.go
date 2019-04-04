package handler

import (
	"fmt"

	"../repo"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var jwtSecretKey []byte = []byte("ThisIsAVerySecretKey")
var identityKey = "identity"

func InitRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.GET("/ping", pingHandler)
	initNoteRoutes(engine, db)
	initUserRoutes(engine, db)
}

func initUserRoutes(engine *gin.Engine, db *gorm.DB) {
	engine.POST("/signin", func(c *gin.Context) {
		userRepository := &repo.UserRepoImpl{
			DB: db,
		}
		result, err := UserSignin(c, userRepository)
		simpleReturnHandler(c, err, result)
	})
	engine.POST("/login", func(c *gin.Context) {
		userRepository := &repo.UserRepoImpl{
			DB: db,
		}
		result, err := UserLogin(c, userRepository)
		simpleReturnHandler(c, err, result)
	})
}

func initNoteRoutes(engine *gin.Engine, db *gorm.DB) {
	groupRouter := engine.Group("/note")

	// 1. Authentication // Identity
	// 2. Lam logger/tracking
	// 3. Recovery
	// 4. Add nhieu cai middleware va no chay tuan tu
	groupRouter.Use(authenMiddleware)
	{
		groupRouter.GET("/:id", func(c *gin.Context) {
			noteRepository := &repo.NoteRepoImpl{
				DB: db,
			}
			result, err := NoteGet(c, noteRepository)
			simpleReturnHandler(c, err, result)
		})
		groupRouter.POST("", func(c *gin.Context) {
			// 1. Repo
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			// 2. Create note
			result, err := NoteCreate(c, repo)
			// 3. Handle result & err
			simpleReturnHandler(c, err, result)
		})
		groupRouter.PUT("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			err := NoteUpdate(c, repo)
			simpleReturnHandler(c, err, nil)
		})
		groupRouter.DELETE("/:id", func(c *gin.Context) {
			repo := &repo.NoteRepoImpl{
				DB: db,
			}
			err := NoteDelete(c, repo)
			simpleReturnHandler(c, err, nil)
		})
	}
}

func simpleReturnHandler(c *gin.Context, err error, result interface{}) {
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, result)
}

func simpleMiddleware(c *gin.Context) {
	// 1. Logic quan trong nhat la, xu ly ngung cai cai request
	// 2. Tao ra cac du lieu de set vao context cho cai handler dung
	// 3. Quyet dinh cho phep di tiep den cai middleware tiep theo hoac handler
	// if c.GetHeader("token") != "202cb962ac59075b964b07152d234b70" {
	// 	c.AbortWithStatus(400)
	// 	return
	// }
	fmt.Println("Print here for every request")
	c.Next()
}

func authenMiddleware(c *gin.Context) {

	tokenString := c.GetHeader("Authentication")
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return jwtSecretKey, nil
	})

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	fmt.Println("jwt-claims:", claims)
	if ok && claims.Valid() == nil {
		c.Set(identityKey, claims.Id)
		c.Next()
		return
	}
	c.AbortWithStatus(401)
}
