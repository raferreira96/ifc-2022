package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/raferreira96/ifc-2022/codepix/application/grpc"
	"github.com/raferreira96/ifc-2022/codepix/infrastructure/db"
)

var database *gorm.DB

func main() {
	database = db.ConnectDB(os.Getenv("env"))
	grpc.StartGrpcServer(database, 50051)
}
