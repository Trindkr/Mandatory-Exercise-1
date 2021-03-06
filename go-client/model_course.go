/*
 * ITU
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type Course struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	ECTS float32 `json:"ECTS,omitempty"`
	Rating float32 `json:"rating,omitempty"`
	ListOfStudents []interface{} `json:"listOfStudents,omitempty"`
	ListOfTeachers []interface{} `json:"listOfTeachers,omitempty"`
}
