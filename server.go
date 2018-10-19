package main

import (
	"crypto/sha1"
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	idLength = 16
	idCharacters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func randomId() string {
	randomBytes := make([]byte, idLength)

	for i := 0; i < idLength; i++ {
		randomBytes[i] = idCharacters[rand.Intn(len(idCharacters))]
	}

	return string(randomBytes)
}

func digestFile(file *multipart.FileHeader) (string, error) {
	fileSrc, err := file.Open()
	if err != nil {
		return "", err
	}

	fileBytes, err := ioutil.ReadAll(fileSrc)
	if err != nil {
		return "", err
	}

	digest := sha1.New()
	digest.Write(fileBytes)

	return fmt.Sprintf("%x", digest.Sum(nil)), nil
}

func handleUpload(c echo.Context) error {
	id := randomId()

	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	sha1, err := digestFile(file)
	if err != nil {
		return err
	}

	return c.String(http.StatusOK, "id: " + id + ", sha1: " + sha1)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	e := echo.New()
	e.Use(middleware.Logger())
	e.Static("/", "public")
	e.POST("/uploads", handleUpload)
	e.Logger.Fatal(e.Start(":1323"))
}
