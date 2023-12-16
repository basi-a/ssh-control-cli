package db

import (
	"log"
	"os"
	"os/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sessions struct {
	gorm.Model
	Id 		int		`gorm:"column:id;autoIncrement:true"`
	Name	string
	User	string
	Host	string
}

const (
	// it is under $HOME
	SQLiteDBFilePath="/.config/sscc/db/" 
)

func Add()  {
	
}

func Del()  {
	
}

func Update()  {
	
}

func Get()  {
	
}

func List()  {
	
}

func InitDB()  {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	
	_, err = os.Stat(u.HomeDir+SQLiteDBFilePath)
	if os.IsNotExist(err){
		err = os.MkdirAll(u.HomeDir+SQLiteDBFilePath, 0750)
		if err != nil {
			log.Fatalln(err)
		}
		var db *gorm.DB
		db, err = gorm.Open(sqlite.Open(u.HomeDir+SQLiteDBFilePath+"sscc.db"), &gorm.Config{})
		if err != nil {
			log.Fatalln(err)
		}
		db.AutoMigrate(&Sessions{})
		if err != nil {
			log.Fatalln(err)
		}
	}else{
		log.Fatalln(err)
	}
}