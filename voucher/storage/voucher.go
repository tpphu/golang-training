package storage

import (
	"database/sql"
	"fmt"

	"../model"
)

type Voucher struct {
	DB *sql.DB
}

const REGISTER_ATOMIC = "INSERT INTO `voucher`(`code`, `discount`, `start`, `end`) " +
	"SELECT ?, ?, ?, ? " +
	"WHERE 0 = (SELECT count(*) FROM `voucher` WHERE `code` = ? AND ? >= `start` AND ? <= `end` LIMIT 1)"
const REGISTER = "INSERT INTO `voucher`(`code`, `discount`, `start`, `end`) " +
	"VALUES(?, ?, ?, ?)"

const COUNT_EXIST = "SELECT count(*) as existing " +
	"FROM `voucher` " +
	"WHERE `code` = ? AND ? >= `start` AND ? <= `end` " +
	"LIMIT 1"

func (s Voucher) IsExit(voucher model.Voucher) (bool, error) {
	count := 0
	err := s.DB.QueryRow(COUNT_EXIST, voucher.Code, voucher.End, voucher.Start).Scan(&count)
	if err != nil {
		return false, err
	}
	fmt.Printf("count %d", count)
	if count >= 1 {
		return true, nil
	}
	return false, nil
}

func (s Voucher) RegisterAtomic(voucher *model.Voucher) error {
	_, err := s.DB.Exec(REGISTER_ATOMIC,
		voucher.Code, voucher.Discount, voucher.Start, voucher.End,
		voucher.Code, voucher.End, voucher.Start)
	if err != nil {
		return err
	}
	return nil
}

func (s Voucher) Register(voucher *model.Voucher) error {
	_, err := s.DB.Exec(REGISTER, voucher.Code, voucher.Discount, voucher.Start, voucher.End)
	if err != nil {
		panic(err)
		return err
	}
	return nil
}
