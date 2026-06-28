package app

import (
	"context"

	"github.com/aravindgopan-dev/cli-auth-system/internal/config"
	"github.com/aravindgopan-dev/cli-auth-system/internal/database"
)


func main(){
	cfg:=config.Load()
	conn:=database.InitDB(cfg.DatabaseURL)
	defer conn.Close(context.Background())
	

	
}