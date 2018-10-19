package main

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestCreateUpload(t *testing.T) {
	e := echo.New()

	requestBody := new(bytes.Buffer)
	bodyWriter := multipart.NewWriter(requestBody)

	fileWriter, err := bodyWriter.CreateFormFile("file", "filename.test")
	if assert.NoError(t, err) {
		fileWriter.Write([]byte("test"))
	}

	bodyWriter.Close()

	request := httptest.NewRequest(http.MethodPost, "/uploads", requestBody)
	request.Header.Set(echo.HeaderContentType, bodyWriter.FormDataContentType())
	recorder := httptest.NewRecorder()
	context := e.NewContext(request, recorder)
	
	err = handleUpload(context)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, recorder.Code)
		assert.Contains(t, recorder.Body.String(), "id:")
		assert.Contains(t, recorder.Body.String(), "sha1:")
	}
}
