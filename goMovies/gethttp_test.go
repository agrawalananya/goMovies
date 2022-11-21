package main

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGetHttp(t *testing.T) {
	var tcs = []struct {
		desc   string
		movies Movies
		out    interface{}
	}{
		{desc: "",
			movies: Movies{
				Id:       6,
				Name:     "Silicon Valley",
				Genre:    "Comedy",
				Rating:   4.5,
				Plot:     "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.",
				Released: true,
			}, out: MoviesData{
				Code:   200,
				Status: "SUCCESS",
				Data: &Data{Movie: &Movies{
					Id:       6,
					Name:     "Silicon Valley",
					Genre:    "Comedy",
					Rating:   4.6,
					Plot:     "Richard a programmer creates an app called the Pied Piper and tries to getinvestors for it. Meanwhile, five other programmers struggle to make their mark in SiliconValley.",
					Released: true,
				}},
			}},
	}

	for _, tt := range tcs {
		final, _ := json.Marshal(tt.movies)
		req := httptest.NewRequest("POST", "/movie", bytes.NewReader(final))
		resp := httptest.NewRecorder()
		createOneMovieData(resp, req)
		var finalAns = resp.Result()
		if !reflect.DeepEqual(tt.out, finalAns) {
			t.Errorf("test case failed")
		}
	}

}
