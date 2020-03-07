package cmd

import (
	"fmt"

	"github.com/tpphu/week2-exercise/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/urfave/cli/v2"
)

var Seed = cli.Command{
	Name:  "seed",
	Usage: "Seeding data to database",
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
				URL: urls[i],
			})
		}

		fmt.Println("This command used to seed and create db")
		return nil
	},
}
