package examples

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/Oleg-React-Dev/go-httpclient/gohttp"
)

func TestMain(m *testing.M) {
	fmt.Println("about to start test cases for package 'example'")
	gohttp.StartMockServer()
	os.Exit(m.Run())
}

func TestGetEndpoints(t *testing.T) {

	t.Run("TestErrorFetchingFromGithub", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method: http.MethodGet,
			Url:    "https://api.github.com",
			Error:  errors.New("timeout getting github endpoints"),
		})
		endpoints, err := GetEndpoints()
		if endpoints != nil {
			t.Error("no endpoints expected")
		}

		if err == nil {
			t.Error("an error expected")
		}
		if err.Error() != "timeout getting github endpoints" {
			t.Error("invalid error message received")
		}
	})

	t.Run("TestErrorUnmarshalResponse", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url":123}`,
			// Error:              errors.New("unexpected end of JSON input"),
		})
		endpoints, err := GetEndpoints()
		if endpoints != nil {
			t.Error("endpoints expected")
		}

		if err == nil {
			t.Error("an error expected")
		}
		if !strings.Contains(err.Error(), "cannot unmarshal") {
			t.Error("invalid error message received")
		}
		// if endpoints != nil {
		// 	t.Error("no endpoints expected at this point")
		// }
	})

	t.Run("TestNoError", func(t *testing.T) {
		gohttp.FlushMocks()
		gohttp.AddMock(gohttp.Mock{
			Method:             http.MethodGet,
			Url:                "https://api.github.com",
			ResponseStatusCode: http.StatusOK,
			ResponseBody:       `{"current_user_url":"https://api.github.com/user"}`,
		})

		endpoints, err := GetEndpoints()
		if err != nil {
			t.Error("no error expected but got", err.Error())
		}
		if endpoints == nil {
			t.Error("endpoints expected")
		}

		if endpoints.CurrentUserUrl != "https://api.github.com/user" {
			t.Error("invalid current user url")
		}

	})

	// fmt.Println(err)
	// fmt.Println(endpoints)
}
