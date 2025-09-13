package migrate

import (
	"fmt"
	"tosshy-blog-api/db"
	"tosshy-blog-api/model"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{})
}
