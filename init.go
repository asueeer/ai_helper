package main

import (
	"ai_helper/biz/config"
	"ai_helper/biz/dal/db"
)

func Init() {
	config.SetDebugMode(true)
	{
		db.InitDB()
		config.InitKVConfig()
	}
}
