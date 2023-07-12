package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func GetDatas(ctx echo.Context) error {
	response, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var data []map[string]any
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}

	return ctx.JSON(http.StatusOK, data)
}

func GetDataById(ctx echo.Context) error {
	id := ctx.Param("Id")

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v", id)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	var data map[string]any
	if err = json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, data)
}

type Post struct {
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

func CreateData(ctx echo.Context) error {
	post := new(Post)
	if err := ctx.Bind(&post); err != nil {
		log.Fatal(err)
	}

	postData, err := json.Marshal(&post)
	if err != nil {
		log.Fatal(err)
	}

	response, err := http.Post("https://jsonplaceholder.typicode.com/posts", "application/json", bytes.NewBuffer(postData))
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	var data map[string]any
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Fatal(err)
	}
	return ctx.JSON(http.StatusOK, data)
}

func DeleteData(ctx echo.Context) error {
	id := ctx.Param("Id")

	client := &http.Client{}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%v", id)
	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	return ctx.JSON(http.StatusOK, "Succesfully Deleted")
}
