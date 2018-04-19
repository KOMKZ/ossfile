package ossfile

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Migration struct {
	Dsn string
}

func (m Migration) InstallDb()  {
	db, err := gorm.Open("mysql", m.Dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&File{})

}

