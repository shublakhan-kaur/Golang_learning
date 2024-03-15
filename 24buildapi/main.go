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

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	r := mux.NewRouter()

	courses = append(courses, Course{CourseId: "1", CourseName: "React", CoursePrice: 200, Author: &Author{"Author1", "google.com"}},
		Course{"2", "Angular", 300, &Author{"Author2", "google.com"}},
		Course{"3", "Java", 400, &Author{"Author3", "google.com"}},
		Course{"4", "Spring", 300, &Author{"Author4", "google.com"}},
		Course{"5", "Spring Boot", 200, &Author{"Author5", "google.com"}})

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))
	// fmt.Println(courses)
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Golang API</h1>"))

}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course find with the given Id")
	//return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course in api")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside json")
		return
	}
	for _, value := range courses {
		if course.CourseName == value.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)

}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a course in Golang")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course is deleted")
			break
		}
	}

}
