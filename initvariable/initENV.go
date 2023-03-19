/*
Create time at 2023年3月18日0018下午 23:40:27
Create User at Administrator
*/

package initvariable

import (
	"github.com/joho/godotenv"
	"log"
)

func InitEnv() {
	err := godotenv.Load("/JWT-Test-Golang/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}
