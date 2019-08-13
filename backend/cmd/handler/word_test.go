package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func Test_GetWord_With_1_Should_Be_Hello_World_In_Thai(t *testing.T) {
	expected := `{text:"สวัสดีชาวโลก"}`

	request := httptest.NewRequest(http.MethodGet, "/api/v1/word/1", nil)
	recorder := httptest.NewRecorder()

	testRoute := gin.Default()
	testRoute.GET("/api/v1/word/1", GetWordHandler)
	testRoute.ServeHTTP(recorder, request)

	response := recorder.Result()

	actual, _ := ioutil.ReadAll(response.Body)

	if string(actual) != expected {
		t.Errorf("actual %s is not equals to expected %s", actual, expected)
	}

}
