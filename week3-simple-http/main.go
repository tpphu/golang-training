package main //rule => define package

// import go gin
import (
	"os"
	"strconv"
	"sync"

	"./handlers/product"
	"./models"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // De minh import thu vien cho thang gorm xai
	"github.com/jinzhu/gorm"
)

// Only Struct // Thieu class // Cho minh mot suc manh khac

// rule
func main() {
	// Step 1: New engine
	//var r *gin.Engine = gin.Default() // phuc tap
	r := gin.Default() //simple
	// 1.1 port
	// Cac viet tuong duong
	var port int = 3000 //Standard
	if os.Getenv("PORT") != "" {
		port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	// Step 2: Route & hander
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Muc dich la de lay ra mot con so khong trung
	// Bien dung chung, dc share giua cac goroutine la rat nguy hiem va han che design kieu nay

	// Thread safe
	var counter int64 = get_last_id()
	mux := sync.Mutex{}
	// Chay cham hon rat nhieu so voi cai unsafe
	r.GET("/safe-counter", func(c *gin.Context) {
		mux.Lock()
		counter = counter + 1 // Kha nang co 2 thang cung goi den ma co cung mot gia tri counter
		mux.Unlock()
		c.JSON(200, gin.H{
			"counter": counter,
		})
	})
	var counter2 int64 = get_last_id()
	r.GET("/unsafe-counter", func(c *gin.Context) {
		// Ban mot san 1 san pham 30 1k, 1 sphan, co kha ngang 2-3 nguoi mua dc san pham do
		// Xui xui moi bi
		counter2 = counter2 + 1 // Kha nang co 2 thang cung goi den ma co cung mot gia tri counter
		c.JSON(200, gin.H{
			"counter": counter2,
		})
	})

	var db *gorm.DB
	var err error
	db, err = gorm.Open("mysql", "root:root@(127.0.0.1)/gomay20?charset=utf8&parseTime=True&loc=Local")
	// if error
	if err != nil {
		panic(err)
	}
	db.DB().SetMaxIdleConns(3) // Neu ma chuong trinh khong su dung connect thi nen dong lai connection do, va chi giu lai 10 cai
	db.DB().SetMaxOpenConns(5) // Neu co nhu cau mo connection nhieu thi chi mo toi da la 20
	db.AutoMigrate(&models.Product{})
	// Nen de duoi sau error
	defer db.Close()
	// Truyen dc mot mock db o day
	productHandler := product.ProductHandler{DB: db}
	// curl -XPOST -H "Content-Type: application/json" --data '{"sku": "P1234", "price": 1000}' http://localhost:3000/product
	r.POST("/product", productHandler.Create)
	// route: product/1
	r.GET("/product/:id", productHandler.ProductGet)
	// Step 3: Start chuong trinh
	r.Run(":" + strconv.Itoa(port))
}

func get_last_id() int64 {
	return 0
}
