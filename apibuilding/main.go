package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for course -seperate file
type Course struct {
	Courseid    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}
type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware /helper -seperate file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}
func main() {
	fmt.Println("Api learning with go lang")
	// route
	r := mux.NewRouter()

	//seeding
	courses = append(courses, Course{Courseid: "1", CourseName: "React",
		CoursePrice: 100,
		Author:      &Author{Fullname: "Rahul", Website: "rahul.com"}})
	courses = append(courses, Course{Courseid: "2", CourseName: "Mern stack",
		CoursePrice: 399, Author: &Author{Fullname: "Suvadip", Website: "mern.dev"}})
	//routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllcourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen to port
	log.Fatal(http.ListenAndServe(":4000", r))

}

//controllers -seperate file

// server home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go api building</h1>"))
}

func getAllcourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get one course")
	w.Header().Set("Content-Type", "application/json")
	//get id from req
	params := mux.Vars(r)
	//loop through course and find matching course id
	for _, course := range courses {
		if course.Courseid == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("no course found with id")
	return

}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("create one course")
	w.Header().Set("Content-Type", "application/json")
	// if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// what about data like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {

		json.NewEncoder(w).Encode("no data inside json")
		return
	}
	//generate a unique id in string
	rand.Seed(time.Now().UnixNano())
	course.Courseid = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")
	//first get id
	params := mux.Vars(r)
	//loop through course and find matching course id and remove and add with same id
	for index, course := range courses {
		if course.Courseid == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.Courseid = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("no course found with id")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")
	//first get id
	params := mux.Vars(r)
	//loop through course and find matching course id and remove and add with same id
	for index, course := range courses {
		if course.Courseid == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}

	}
	json.NewEncoder(w).Encode("no course found with id")
}
