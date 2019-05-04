package main

import "sync"

type DB struct {
}

var instance *DB
var once sync.Once

func GetDBInstance() *DB {

	once.Do(func() {
		instance = &DB{}
	})

	return instance
}

func main() {

}
