/*
 * ITU
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"TeacherAddTeacherPost",
		strings.ToUpper("Post"),
		"/v1/teacher/addTeacher",
		TeacherAddTeacherPost,
	},

	Route{
		"TeacherGetTeacherGet",
		strings.ToUpper("Get"),
		"/v1/teacher/getTeacher",
		TeacherGetTeacherGet,
	},

	Route{
		"TeacherRemoveTeacherDelete",
		strings.ToUpper("Delete"),
		"/v1/teacher/removeTeacher",
		TeacherRemoveTeacherDelete,
	},

	Route{
		"TeacherUpdateTeacherPut",
		strings.ToUpper("Put"),
		"/v1/teacher/updateTeacher",
		TeacherUpdateTeacherPut,
	},
}
