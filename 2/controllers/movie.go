package controllers

import (
	"backend-golang/2/constants"
	"backend-golang/2/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func GetMoviesWithPagination(c echo.Context) error {
	method := "GET"

	page := c.QueryParam("page")
	keyword := c.QueryParam("keyword")

	if page == "" {
		page = "1"
	}

	if keyword == "" {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%s status=%d error=%s", method, keyword, page, http.StatusBadRequest, "Invalid Keyword")
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Keyword")
	}

	keyword = strings.Replace(keyword, " ", "&", -1)

	pageNumber, err := strconv.Atoi(page)
	if err != nil {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%s status=%d error=%s", method, keyword, page, http.StatusBadRequest, err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Page Number")
	}
	
	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&s=%s&page=%d", constants.OMDB_KEY, keyword, pageNumber)
	
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%d status=%d error=%s", method, keyword, pageNumber, http.StatusInternalServerError, err.Error())
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }

    var client = &http.Client{}
	res, err := client.Do(req)
    if err != nil {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%d status=%d error=%s", method, keyword, pageNumber, http.StatusInternalServerError, err.Error())
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%d status=%d error=%s", method, keyword, pageNumber, http.StatusInternalServerError, err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	var response models.Response
    err = json.Unmarshal(body, &response)
    if err != nil {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%d status=%d error=%s", method, keyword, pageNumber, http.StatusInternalServerError, err.Error())
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }

	if len(response.Search) == 0 {
		log.Printf("method=%s uri=/movies?keyword=%s&page=%d status=%d error=%s", method, keyword, pageNumber, http.StatusNotFound, "Not Found")
        return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	log.Printf("method=%s uri=/movies?keyword=%s&page=%d status=%d", method, keyword, pageNumber, http.StatusOK)
	return c.JSON(http.StatusOK, response.Search)
}

func GetMoviesByID(c echo.Context) error {
	method := "GET"

	id := c.Param("id")

	url := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&i=%s", constants.OMDB_KEY, id)
	
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("method=%s uri=/movies/%s status=%d error=%v", method, id, http.StatusInternalServerError, err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }

    var client = &http.Client{}
	res, err := client.Do(req)
    if err != nil {
		log.Printf("method=%s uri=/movies/%s status=%d error=%v", method, id, http.StatusInternalServerError, err)
        return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
    }
    
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("method=%s uri=/movies/%s status=%d error=%v", method, id, http.StatusInternalServerError, err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Internal Server Error")
	}

	var response models.MovieDetail
    err = json.Unmarshal(body, &response)
    if err != nil {
		log.Printf("method=%s uri=/movies/%s status=%d error=%v", method, id, http.StatusNotFound, err)
        return echo.NewHTTPError(http.StatusNotFound, "Not Found")
    } else if response.Response == "False" {
		log.Printf("method=%s uri=/movies/%s status=%d error=%s", method, id, http.StatusNotFound, "Not Found")
        return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	log.Printf("method=%s uri=/movies/%s status=%d", method, id, http.StatusOK)
	return c.JSON(http.StatusOK, response)
}