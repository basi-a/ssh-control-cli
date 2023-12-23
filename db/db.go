package db

import (
	"log"
	"os"
	"os/user"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Sessions struct {
	Id 			int			`gorm:"column:id;primary_key;autoIncrement:true"` 
	UUID 		string      `gorm:"column:uuid;uniqueIndex"`
	SessionName	string
	UserName	string
	Host		string
	// gorm.Model
}

const (
	// it is under $HOME
	SQLiteDBFilePath="/.config/sscc/db/" 
)

func Add(sessionName, userName, host string)  error{
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open(u.HomeDir+SQLiteDBFilePath+"sscc.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	session := Sessions{SessionName: sessionName, UserName: userName, Host: host}
	result := db.Create(&session)
	return result.Error
}

func Del(uuid, sessionName string) error {
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open(u.HomeDir+SQLiteDBFilePath+"sscc.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	result := db.Where("uuid = ?", uuid).Delete(&Sessions{})
	return result.Error
}

func Update(sessionName, userName, host string)  {
	
}

func Get(sessionName, userName, host string) ([]Sessions, error){
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open(u.HomeDir+SQLiteDBFilePath+"sscc.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	var sessions []Sessions
	var cond string  
    if sessionName != "" {  
        cond += "session_name = ?"  
    }  
    if userName != "" {  
        if cond != "" {  
            cond += " AND "  
        }  
        cond += "user_name = ?"  
    }  
    if host != "" {  
        if cond != "" {  
            cond += " AND "  
        }  
        cond += "host = ?"  
    }  
    
	result := db.Where(cond, sessionName, userName, host).Find(&sessions)
    return sessions, result.Error  
}

func List() ([]Sessions, error){
	u, err := user.Current()
	if err != nil {
		log.Fatalln(err)
	}
	var db *gorm.DB
	db, err = gorm.Open(sqlite.Open(u.HomeDir+SQLiteDBFilePath+"sscc.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	var sessions []Sessions
	result := db.Find(&sessions)
	return sessions, result.Error
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