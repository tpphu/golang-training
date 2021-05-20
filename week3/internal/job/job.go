package job

import (
	"fmt"
	"time"

	"phudt/week3/internal/model"

	cli "github.com/urfave/cli/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Job(c *cli.Context) error {
	n := c.Int("n")
	fmt.Println("flag n:", n)
	x := c.Bool("x")
	fmt.Println("flag x:", x)
	fmt.Println("This is job app - sau khi refactor")
	// Bai Tap: Lam sao de lay cac gia tri nay tu Environment
	dsn := "root:root@tcp(127.0.0.1:3306)/covid?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	db.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	db.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	db.SetConnMaxLifetime(time.Hour)
	if err != nil {
		fmt.Println("Connect to db:", err)
		return err
	}
	// for i := 0; i < 1000; i++ {
	// 	name := fmt.Sprintf("Tran Phong Phu %d", i)
	// 	db.Exec(`INSERT INTO patient(fullname, address, birthday)
	// 	VALUES(?, "Ho Chi Minh", "1986-01-22")`, name)
	// }
	patient := model.NewPatient("Bach Pham", "HCM")
	patient.Fullname = "Phu DT"
	fmt.Println("Table name:", patient.TableName())
	fmt.Println("Patient Id Before:", patient.Id)
	err = db.Create(&patient).Error
	if err != nil {
		fmt.Println("Create error:", err)
	}

	fmt.Println("Patient Id After:", patient.Id)

	// Print full patient

	fmt.Println("full patient:", patient)
	fmt.Println("Lan nay chay lau")
	fmt.Println("age of patient:", patient.GetAge())
	fmt.Println("lan  nay chay nhanh")
	fmt.Println("age of patient:", patient.GetAge())

	return nil
}
