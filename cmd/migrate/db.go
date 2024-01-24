package main

import (
	"fmt"
	"forthboxbe/app/model"
)

func main() {
	fmt.Println("Begin user model migrate...")
	model.GetUserModel().AutoMigrate(&model.User{})
	fmt.Println("Migrate user model done")

	fmt.Println("Begin verify model migrate...")
	model.GetVerifyTokenModel().AutoMigrate(&model.VerifyToken{})
	fmt.Println("Migrate verify model done")

	fmt.Println("Begin pic model migrate...")
	model.GetPicModel().AutoMigrate(&model.Pic{})
	fmt.Println("Migrate pic model done")
}

