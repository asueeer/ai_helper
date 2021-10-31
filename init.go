package main

import (
	"nearby/biz/config"
	"nearby/biz/dal/db"
)

func Init() {
	config.SetDebugMode(true)
	{
		db.InitDB()
		config.InitKVConfig()
	}
}
