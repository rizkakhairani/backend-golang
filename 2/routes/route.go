package routes

import (
	"backend-golang/2/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	//---------------------------------------------------------------//
	//-----------SEARCH MOVIES WITH KEYWORD AND PAGE NUMBER----------//
	//---------------------------------------------------------------//

	go e.GET("/movies", controllers.GetMoviesWithPagination)

	// Example:
	// localhost:8000/movies?keyword=batman&page=2

	
	//---------------------------------------------------------------//
	//----------------------GET MOVIE BY IMDB ID---------------------//
	//---------------------------------------------------------------//

	e.GET("/movies/:id", controllers.GetMoviesByID)
	
	// Example:
	// localhost:8000/movies/tt2313197

	return e
}