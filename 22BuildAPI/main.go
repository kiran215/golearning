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

// Model For courses - file
type Course struct {
	Courseid    string  `json:"courseid"`
	Coursename  string  `json:"coursename"`
	Courseprice string  `json:"price"`
	Author      *Author `json:"author"`
}

//Author thype as pointer

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// middleware,helper - file
func (c *Course) isEmpty() bool {

	return c.Coursename == ""
}

func main() {
	fmt.Println("Hello we will Build API")
	// author := Author{"Kiran Sadaye", "Golang.dev"}
	// courseList := Course{"1", "Golang", "300", &author}
	// courses = append(courses, courseList)
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome)
	r.HandleFunc("/getAllCourses", getAllCourses).Methods("GET")
	// r.HandleFunc("/getOneCourse/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/getOneCourse", getOneCourse).Methods("GET").Queries("id", "{id}")

	r.HandleFunc("/createOneCourse", createOneCourse).Methods("POST")
	r.HandleFunc("/updateOneCourse/{id}", updateOneCourse).Methods("POST")
	r.HandleFunc("/deleteOneCourse/{id}", deleteOneCourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}

//controllers - file
//Serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Go lang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All courses")
	//Set headers
	w.Header().Set("Content-Type", "application/json")
	//below will return as json
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	//Set headers
	w.Header().Set("Content-Type", "application/json")

	//grap id from request
	params := mux.Vars(r)

	//loop through courses and find matching id user has send in request and return response

	for _, course := range courses {
		if course.Courseid == params["id"] {
			//below will return as json
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createOneCourse")
	//Set headers
	w.Header().Set("Content-Type", "application/json")

	//What is body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid request")
		return
	}

	var course Course

	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.isEmpty() {
		json.NewEncoder(w).Encode("NO data in json")
		return
	}

	for _, extcourse := range courses {
		fmt.Println(extcourse)
		if extcourse.Coursename == course.Coursename {
			json.NewEncoder(w).Encode("Coursename already exist")
			return
		}

	}

	//generate unique id
	//append course into courses

	// rand.Seed(time.Now().UnixNano())      //deprecated
	rand.NewSource(time.Now().UnixNano()) //new

	course.Courseid = strconv.Itoa(rand.Intn(100))
	// author := Author{course.Author.Fullname, course.Author.Website}
	// courseList := Course{course.Courseid, course.Coursename, course.Courseprice, &author}
	courses = append(courses, course)
	json.NewEncoder(w).Encode(courses)

	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")

	//Set headers
	w.Header().Set("Content-Type", "application/json")

	//What is body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid request")
		return
	}

	//first - grab id from rquest'
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Courseid == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)

			course.Courseid = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(courses)

			return
		}
	}
	json.NewEncoder(w).Encode("Course id not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")

	//Set headers
	w.Header().Set("Content-Type", "application/json")

	//What is body is empty

	if r.Body == nil {
		json.NewEncoder(w).Encode("Invalid request")
		return
	}

	//first - grab id from rquest'
	params := mux.Vars(r)

	for index, course := range courses {
		if course.Courseid == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break // if course id found delete it and break the loop instead of looping over complete array
		}
	}
	json.NewEncoder(w).Encode(courses)

	return

}
