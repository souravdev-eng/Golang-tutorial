package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware or helper -file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	// seeding data
	courses = append(courses, Course{
		CourseId: "1", CourseName: "ReactJS",
		CoursePrice: 299, Author: &Author{FullName: "Sourav", Website: "udemy.com"},
	})
	courses = append(courses, Course{
		CourseId: "2", CourseName: "Golang",
		CoursePrice: 399, Author: &Author{FullName: "Hitesh", Website: "udemy.com"},
	})
	courses = append(courses, Course{
		CourseId: "3", CourseName: "NodeJS",
		CoursePrice: 199, Author: &Author{FullName: "Stifan Grainder", Website: "udemy.com"},
	})

	// router
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers -file

// Serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Go API home rote</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find the matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with this ID")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about - {}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}
	// generate unique id, string
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	// append course into courses
	courses = append(courses, course)

	json.NewEncoder(w).Encode(course)
	return
}
func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from request
	params := mux.Vars(r)

	// loop , id, remove, add with my ID
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO Send a response when id is not fount
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
}
