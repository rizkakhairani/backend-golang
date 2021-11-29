package controllers_test

import (
	"backend-golang/2/controllers"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetMoviesWithPagination(t *testing.T) {
	var testCases = []struct {
		name       	string
		path       	string
		keyword		string
		page		string
		expectCode 	int
	}{
		{
			name:       "Get Movies With Pagination",
			path:       "/movies",
			keyword: 	"Batman",
			page: 		"2",
			expectCode: http.StatusOK,
		},
		{
			name:       "Get Movies With Pagination",
			path:       "/movies",
			keyword: 	"Batman",
			page: 		"",
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		q := make(url.Values)
		q.Set("keyword", testCase.keyword)
		q.Set("page", testCase.page)
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			if assert.NoError(t, controllers.GetMoviesWithPagination(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

func TestGetMoviesWithPaginationError(t *testing.T) {
	var testCases = []struct {
		name        string
		path        string
		keyword		string
		page		string
		expectCode  int
		expectError string
	}{
		{
			name:        	"Get Movies With Pagination Invalid Keyword",
			path:        	"/movies",
			keyword: 		"",
			page: 			"1",
			expectCode:  	http.StatusBadRequest,
			expectError: 	"Invalid Keyword",
		},
		{
			name:        	"Get Movies With Pagination Invalid Page Number",
			path:        	"/movies",
			keyword: 		"Batman",
			page: 			"a",
			expectCode: 	http.StatusBadRequest,
			expectError: 	"Invalid Page Number",
		},
		{
			name:        	"Get Movies With Pagination Not Found",
			path:        	"/movies",
			keyword: 		"1",
			page: 			"1",
			expectCode: 	http.StatusNotFound,
			expectError: 	"Not Found",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		q := make(url.Values)
		q.Set("keyword", testCase.keyword)
		q.Set("page", testCase.page)
		req := httptest.NewRequest(http.MethodGet, "/?"+q.Encode(), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)

		t.Run(testCase.name, func(t *testing.T) {
			err := controllers.GetMoviesWithPagination(c)
			if assert.Error(t, err) {
				assert.Containsf(t, err.Error(), testCase.expectError, "expected error containing %s, got %s", testCase.expectError, err)
			}
		})
	}
}

func TestGetMoviesByID(t *testing.T) {
	var testCases = []struct {
		name       	string
		path       	string
		id			string
		expectCode 	int
	}{
		{
			name:       "Get Movies By ID",
			path:       "/movies",
			id: 		"tt0372784",
			expectCode: http.StatusOK,
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.id)

		t.Run(testCase.name, func(t *testing.T) {
			if assert.NoError(t, controllers.GetMoviesByID(c)) {
				assert.Equal(t, testCase.expectCode, rec.Code)
			}
		})
	}
}

func TestGetMoviesByIDError(t *testing.T) {
	var testCases = []struct {
		name        string
		path       	string
		id			string
		expectCode  int
		expectError string
	}{
		{
			name:        	"Get Movies By ID Not Found",
			path:      		"/movies",
			id: 			"1",
			expectCode: 	http.StatusNotFound,
			expectError: 	"Not Found",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testCase.path)
		c.SetParamNames("id")
		c.SetParamValues(testCase.id)

		t.Run(testCase.name, func(t *testing.T) {
			err := controllers.GetMoviesByID(c)
			if assert.Error(t, err) {
				assert.Containsf(t, err.Error(), testCase.expectError, "expected error containing %s, got %s", testCase.expectError, err)
			}
		})
	}
}