// go install github.com/joho/godotenv@latest

package main

import (
	"fmt"
)

func main() {
	// err := dotenv.Load(".env")
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// 	os.Exit(1)
	// }

	// secretKey := os.Getenv("APP_URL")

	fmt.Println(secretKey)

	// now do something with s3 or whatever
}
