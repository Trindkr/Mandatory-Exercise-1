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
		"CourseAddCoursePost",
		strings.ToUpper("Post"),
		"/v1/course/addCourse",
		CourseAddCoursePost,
	},

	Route{
		"CourseGetCoursesGet",
		strings.ToUpper("Get"),
		"/v1/course/getCourses/",
		CourseGetCoursesGet,
	},

	Route{
		"CourseRemoveCourseDelete",
		strings.ToUpper("Delete"),
		"/v1/course/removeCourse",
		CourseRemoveCourseDelete,
	},

	Route{
		"CourseUpdateCoursePut",
		strings.ToUpper("Put"),
		"/v1/course/updateCourse",
		CourseUpdateCoursePut,
	},

	Route{
		"StudentAddStudentPost",
		strings.ToUpper("Post"),
		"/v1/student/addStudent",
		StudentAddStudentPost,
	},

	Route{
		"StudentGetStudentsGet",
		strings.ToUpper("Get"),
		"/v1/student/getStudents/",
		StudentGetStudentsGet,
	},

	Route{
		"StudentRemoveStudentDelete",
		strings.ToUpper("Delete"),
		"/v1/student/removeStudent",
		StudentRemoveStudentDelete,
	},

	Route{
		"StudentUpdateStudentPut",
		strings.ToUpper("Put"),
		"/v1/student/updateStudent",
		StudentUpdateStudentPut,
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