package storage

import (
	"database/sql"
	"fmt"
	"time"

	"../model"
	"../proto"
)

type Voucher struct {
	DB *sql.DB
}

const REGISTER_ISOLATION = "INSERT INTO `voucher`(`code`, `discount`, `start`, `end`) " +
	"SELECT ?, ?, ?, ? " +
	"FROM locker " +
	"WHERE id = 1 AND " +
	"0 = (SELECT count(*) FROM `voucher` WHERE `code` = ? AND ? >= `start` AND ? <= `end` LIMIT 1) " +
	"FOR UPDATE"

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

func (s Voucher) RegisterIsolation(voucher *proto.VoucherReq) (bool, error) {
	tx, err := s.DB.Begin()
	isExist := false
	// Lock row = 1 cua cai locker ma khong can merge cao cau query chinh
	// va su dung lai cau atomic
	if err != nil {
		return isExist, err
	}
	start := time.Unix(voucher.Start.Seconds, 0)
	end := time.Unix(voucher.End.Seconds, 0)
	result, err := tx.Exec(REGISTER_ISOLATION,
		voucher.Code, voucher.Discount, start, end,
		voucher.Code, end, start)
	if err != nil {
		tx.Rollback()
		return isExist, err
	}

	if err = tx.Commit(); err != nil {
		return isExist, err
	}
	rowEffect, _ := result.RowsAffected()
	if rowEffect == 0 {
		isExist = true
	}
	return isExist, nil
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
