package main

import (
	"fmt"
	"strconv"
	"time"

	"./handler"
	"./model"
	"./repo"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

type Data struct {
	X string
}

func main() {
	// Connect vo DB MySQL
	db, err := gorm.Open("mysql", config.DBConnectString)
	if err != nil {
		panic(err)
	}
	db.LogMode(false)
	// Migration => Used only for Development Enviroment
	db.AutoMigrate(&model.Note{}, &model.User{})

	r := gin.Default()
	r.Use(func(c *gin.Context) {
		fmt.Println("Co di vao middleware 2")
		return
	})
	r.GET("/ping", func(c *gin.Context) {
		name := c.GetHeader("Name")
		d := Data{}
		c.BindJSON(&d)
		x, _ := c.Get("x")
		d.X = x.(string)
		c.JSON(200, gin.H{
			"message": "pong",
			"name":    name,
			"x":       d.X,
		})
	})

	group := r.Group("/note")
	{
		group.Use(func(c *gin.Context) {
			tokenString := c.GetHeader("Authentication")
			if tokenString == "" {
				c.AbortWithStatusJSON(404, gin.H{
					"error": "Missing token",
				})
				return
			}
			// Parse cai token => ve doi tuong claim
			token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
				// Don't forget to validate the alg is what you expect:
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
				return config.JWTSecretKey, nil
			})

			if err != nil {
				c.AbortWithStatus(401)
				return
			}

			claims, ok := token.Claims.(*jwt.StandardClaims)

			if ok && claims.Valid() == nil {
				c.Set(config.IdentityKey, claims.Id)
				c.Next()
				return
			}
			c.AbortWithStatus(401)
		})
		group.GET("/:id", func(c *gin.Context) {
			note := model.Note{}
			id := c.Param("id")
			// Route can biet la lam the nao de lay ra cai identity
			identity, _ := c.Get(config.IdentityKey)
			err := db.Where("id = ? AND author_id = ?", id, identity).Find(&note).Error
			if err != nil {
				c.JSON(404, gin.H{
					"success": false,
					"err":     err.Error(),
				})
				return
			}
			c.JSON(200, note)
		})
		group.POST("", func(c *gin.Context) {
			// 1. La chi la de get gia tri ma request tu client
			input, _ := handler.PreProcessingNoteInput(c)
			// Phan implement chinh ve viec lam viec voi DB
			noteRepo := repo.NoteRepoImpl{
				DB: db,
			}
			// Minh abstract cai handler input de ma thay the cai repo/mock dc
			note, err := handler.CreateNoteHandler(noteRepo, input)
			// Phan code nay co the refactor move di cho khac dc
			c.JSON(200, note)
		})
	}

	r.POST("/signin", func(c *gin.Context) {
		user := model.User{}
		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"err":     err.Error(),
			})
			return
		}
		password := []byte(user.Password)
		hashPassword, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
		user.Password = string(hashPassword)
		db.Create(&user)
		c.JSON(200, user)
	})
	r.POST("/login", func(c *gin.Context) {
		login := model.UserLoginForm{}
		err := c.BindJSON(&login)
		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"err":     err.Error(),
			})
			return
		}
		// 1. Tim user co cai username hoac email trung voi cai nguoi dung nhap
		// Neu khong co thi login failed
		// 2. Neu co thi minh kiem tra co dung password khong
		user := &model.User{}
		err = db.
			Where("username = ? OR email = ?", login.Login, login.Login).
			First(user).Error

		if err != nil {
			c.JSON(400, gin.H{
				"success": false,
				"err":     err.Error(),
			})
			return
		}
		if user.ID == 0 {
			c.JSON(401, gin.H{
				"success": false,
				"err":     "Not found user",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
		if err != nil {
			c.JSON(401, gin.H{
				"success": false,
				"err":     "Username or password is not correct",
			})
			return
		}
		// 3. Generate ra json va tra ve token cho nguoi dung
		claims := &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 365).Unix(),
			Issuer:    "NordicCoder",
			Id:        strconv.Itoa(int(user.ID)),
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// Sign key
		tokenString, _ := token.SignedString(jwtSecretKey)

		userLoginResponse := &model.UserLoginReponse{
			ID:       user.ID,
			Fullname: user.Fullname,
			Token:    tokenString,
		}
		c.SetCookie("Token", tokenString, 3600*24*365, "/", "", false, true)
		c.JSON(200, userLoginResponse)
	})

	r.Run(":8088") // listen and serve on 0.0.0.0:8080
}
