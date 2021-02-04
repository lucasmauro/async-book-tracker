package routes

import (
	"async-book-shelf/src/controllers"
	"net/http"
)

var bookRoutes = []Route{
	{
		URI:      "/books",
		Method:   http.MethodGet,
		Function: controllers.GetBooks,
	},
}
