package main

import (
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestDeleteHttp(t *testing.T) {
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

}
