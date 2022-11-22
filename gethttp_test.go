package main

/*import (
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetHttp(t *testing.T) {
	var tcs = []struct {
		desc string
		id   int
		out  MoviesData
	}{
		{desc: "SUCCESS", id: 1, out: MoviesData{
			Code:   200,
			Status: "SUCCESS",
			Data: &Data{Movie: &Movies{
				Id:       1,
				Name:     "Silicon Valley",
				Genre:    "Comedy",
				Rating:   4.6,
				Plot:     "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.",
				Released: true,
			}},
		}},
	}

	for i, tt := range tcs {
		url := "/movie/" + strconv.Itoa(tt.id)
		req := httptest.NewRequest("GET", url, nil)
		wr := httptest.NewRecorder()

		getOneMovie(wr, req)
		if tt.out.Code != wr.Code {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", i, tt.out, wr.Code)
		}
	}

}*/

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetMovie(t *testing.T) {
	testCases := []struct {
		target    string
		expOutput MoviesData
		expError  error
	}{
		//{
		//	target: "/movies/$",
		//	expOutput: MoviesData{
		//		Error: "No movie found with the id"},
		//},
		{
			target: "/movie",
			expOutput: MoviesData{
				Data: &Data{Movie: &Movies{1, "Silicon Valley", "Comedy", 4.5, "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.", true}},
			},
		},
	}

	for _, tt := range testCases {

		r := httptest.NewRequest(http.MethodGet, tt.target, nil)
		w := httptest.NewRecorder()
		params := map[string]string{
			"id": "1",
		}
		r = mux.SetURLVars(r, params)
		getOneMovie(w, r)

		var movieResponse MoviesData
		res := w.Result()
		data, err := io.ReadAll(res.Body)

		if err != tt.expError {
			t.Fatalf("Wrong Output.")
			return
		}

		err = json.Unmarshal(data, &movieResponse)
		if err != nil {
			t.Fatalf("Wrong Output nil.")
			return
		}

		if !reflect.DeepEqual(movieResponse, tt.expOutput) {
			t.Fatalf("Wrong Output, Expected: %v, Got: %v", tt.expOutput, movieResponse)
		}
		fmt.Printf("output, Expected: %v, Got: %v\n", tt.expOutput, movieResponse)
	}
}
