package migration

import (

	"fmt"
    "demo/app/models"
    "demo/pkg/db_service"

)



func Migrate() {

	 fmt.Println("[Migrate] Migratation ")
     db :=  db_service.Connect();

     db.AutoMigrate(&models.Employ{})

}
