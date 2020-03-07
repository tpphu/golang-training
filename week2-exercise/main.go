package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"

	"github.com/tpphu/golang-trainning/app"
	"github.com/tpphu/golang-trainning/model"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error occurs when load enviroment, %s", err.Error())
	}
	app := &cli.App{
		Name:    "spider",
		Usage:   "Crawler Service",
		Version: "1.0.0",
		Action: func(c *cli.Context) error {
			fmt.Println("Welcome to Crawler Service. To see how to use, type -h for more detail")
			return nil
		},
		Before: func(c *cli.Context) error {
			db, err := gorm.Open("mysql", c.String("database"))
			if err != nil {
				panic(err)
			}
			// db.DB().SetMaxIdleConns(2)
			// db.DB().SetMaxOpenConns(10)
			db.LogMode(c.Bool("dblog"))
			c.App.Metadata["db"] = db
			return nil
		},
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "port",
				Aliases: []string{"p"},
				Usage:   "Application's port is listen",
				EnvVars: []string{"PORT"},
				Value:   3000,
			},
			&cli.StringFlag{
				Name:    "database",
				Aliases: []string{"db"},
				Usage:   "Database URI to connect to DB",
				EnvVars: []string{"DB_URI"},
			},
			&cli.BoolFlag{
				Name:    "dblog",
				Usage:   "Application's port is listen",
				EnvVars: []string{"DB_LOG"},
				Value:   false,
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"run"},
				Usage:   "Start the service",
				Action: func(c *cli.Context) error {
					fmt.Println("Start the service")
					db := c.App.Metadata["db"].(*gorm.DB)
					app := app.NewApp(db)
					urls := app.Load()
					articles := app.Crawl(urls)
					app.InsertArticleToDb(articles)
					return nil
				},
			},
			{
				Name:    "migrate",
				Aliases: []string{"m"},
				Usage:   "Migrate schema to DB",
				Action: func(c *cli.Context) error {
					db := c.App.Metadata["db"].(*gorm.DB)
					db.DropTableIfExists(&model.Url{}, &model.Article{})
					db.AutoMigrate(&model.Url{}, &model.Article{})
					return nil
				},
			},
			{
				Name:    "seed",
				Aliases: []string{"s"},
				Usage:   "Seed data to DB",
				Action: func(c *cli.Context) error {
					db := c.App.Metadata["db"].(*gorm.DB)
					urls := []string{
						"https://vnexpress.net/the-gioi/iran-ghi-nhan-ky-luc-hon-1-200-ca-nhiem-ncov-moi-4065455.html",
						"http://tiasang.com.vn/-quan-ly-khoa-hoc/Thay-doi-trong-dau-tu-congtu-cho-khoa-hoc-20767",
						"https://www.thesaigontimes.vn/td/300790/doanh-nghiep-lam-nguy-vi-e-am-chu-cho-thue-cung-lao-dao.html",
						"https://www.thesaigontimes.vn/300783/chong-dich-nuoc-bot-khong-the-bang-nuoc-bot.html",
						"https://www.thesaigontimes.vn/td/300773/doanh-nghiep-noi-dia-dung-truoc-noi-lo-thau-tom.html",
						"https://www.thesaigontimes.vn/300849/mua-sam-qua-mang-xa-hoi-o-tphcm-chiem-50-mua-hang-truc-tuyen.html",
						"https://www.thesaigontimes.vn/300839/goi-tin-dung-285000-ti-dong-ho-tro-doanh-nghiep-trong-dich-covid-19.html",
						"https://www.thesaigontimes.vn/td/300743/nguoc-chieu-thi-truong-tang-truong-loi-nhuan-bat-dong-san-vuot-ca-ngan-hang.html",
						"https://www.thesaigontimes.vn/td/300828/khach-san-nho-o-tphcm-cam-cu-voi-khach-le-khach-theo-gio.html",
						"https://www.thesaigontimes.vn/300808/novaland-gui-don-%E2%80%98keu-cuu%E2%80%99-bo-xay-dung-tra-ve-tphcm.html",
						"https://www.thesaigontimes.vn/300840/thu-tuong-yeu-cau-giai-phap-de-khan-truong-phuc-hoi-du-lich-hang-khong-.html",
						"https://www.thesaigontimes.vn/300828/khach-san-nho-o-tphcm-cam-cu-voi-khach-le-khach-theo-gio.html",
						"https://www.thesaigontimes.vn/300821/soc-trang-co-nha-may-dien-gio-thu-2-hon-5300-ti-dong.html",
						"https://www.thesaigontimes.vn/300719/resort-moi-o-nam-hoi-an-khai-truong-trong-mua-vang-khach-vi-covid-19.html",
						"https://www.thesaigontimes.vn/300775/ung-pho-voi-dich-benh-dung-lam-theo-cam-xuc.html",
						"https://www.thesaigontimes.vn/300856/chiu-nhieu-suc-ep-facebook-can-nhac-dieu-chinh-du-an-tien-ao-libra.html",
					}
					for i := 0; i < len(urls); i++ {
						db.Create(&model.Url{
							URL:   urls[i],
							State: model.UrlStateIdle,
						})
					}
					return nil
				},
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
