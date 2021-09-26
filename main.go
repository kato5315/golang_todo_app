package main

import (
	"fmt"
	"golang_todo_app/app/models"
	"log"
)

func main() {
	fmt.Println(models.Db)
	// controllers.StartMainServer()
  user, _ := models.GetUserByEmail("aaa@aaa.com")
  fmt.Println(user)

  session, err := user.CreateSession()
  if err != nil {
    log.Println(err)
  }
  fmt.Println(session)

  valid, _ := session.CheckSession()
  fmt.Println(valid)
  }
