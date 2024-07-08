package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"
	"strings"
)

type Task struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Completed   bool   `json:"completed"`
}

var Tasks []Task

func main() {
	http.HandleFunc("/gettasks", Gettaskshandler)
	http.HandleFunc("/posttasks", Posttaskshandler)
	http.HandleFunc("/puttasks/", Puttaskshandler)
	http.HandleFunc("/deletetasks/", Delettaskshandler)

	fmt.Println("server is running on port 8080")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}

}
func Gettaskshandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Tasks)
}
func Posttaskshandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var newTask Task
	err := json.NewDecoder(req.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	newTask.Id = len(Tasks) + 1
	Tasks = append(Tasks, newTask)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("myheader", "create task")
	var response = make(map[string]interface{})
	response["data"] = newTask
	response["message"] = "task added"

	json.NewEncoder(w).Encode(response)

}

func Puttaskshandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(req.URL.Path, "/puttasks/")
	idInt, _ := strconv.Atoi(id)
	var updatedTask Task
	err := json.NewDecoder(req.Body).Decode(&updatedTask)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for i, task := range Tasks {
		if task.Id == idInt {
			
			Tasks[i].Title = updatedTask.Title
			Tasks[i].Description = updatedTask.Description
			Tasks[i].Completed = updatedTask.Completed
			break
		}
	}
	updatedTask.Id=idInt
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("myheader", "update task")
	var response = make(map[string]interface{})
	response["message"] = "task updated"
	response["data"] = updatedTask
	json.NewEncoder(w).Encode(response)
}

func Delettaskshandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	id := strings.TrimPrefix(req.URL.Path, "/deletetasks/")
	idInt, _ := strconv.Atoi(id)
	for i, task := range Tasks {
		if task.Id == idInt {
			Tasks = append(Tasks[:i], Tasks[i+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("myheader", "delete task")
	var response = make(map[string]string)
	response["message"] = "task deleted id " + id
	json.NewEncoder(w).Encode(response)
}
