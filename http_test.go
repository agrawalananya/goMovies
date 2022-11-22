package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	var tcs = []struct {
		desc   string
		movies Movies
		out    MoviesData
	}{
		{desc: "Success Case",
			movies: Movies{
				Id:       1,
				Name:     "Silicon Valley",
				Genre:    "Comedy",
				Rating:   4.5,
				Plot:     "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.",
				Released: true,
			},
			out: MoviesData{
				Code:   http.StatusOK,
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

		{desc: "Failure Case",
			movies: Movies{
				Id:       2,
				Name:     "Silicon Valley",
				Genre:    "Comedy",
				Rating:   4.5,
				Plot:     "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.",
				Released: true,
			},
			out: MoviesData{
				Code:   http.StatusForbidden,
				Status: "FAILURE",
				Data:   nil,
			}},
	}

	for i, tt := range tcs {

		url := "/movie"
		jsnBdy, _ := json.Marshal(tt.movies)
		buff := bytes.NewBuffer(jsnBdy)
		req := httptest.NewRequest("POST", url, buff)
		wr := httptest.NewRecorder()

		createOneMovieData(wr, req)
		if tt.out.Code != wr.Code {
			t.Errorf("TestCase[%v] Expected: \t%v\nGot: \t%v\n", i, tt.out.Code, wr.Code)
		}

		//final, _ := json.Marshal(tt.movies)
		//req := httptest.NewRequest("POST", "/movie", bytes.NewReader(final))
		//resp := httptest.NewRecorder()
		//createOneMovieData(resp, req)
		//var finalAns = resp.Body
		//if !reflect.DeepEqual(tt.out, finalAns) {
		//	t.Errorf("test case failed")
		//}
	}

}
