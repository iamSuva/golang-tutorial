package main

import (
	"encoding/json"
	"fmt"
	"net/http"


	// "strings"
)
type Response struct{
	Message string `json:"message`
}
type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}


func main() {
	fmt.Println("this is http server")
	// http.HandleFunc("/",Hellofun)
	// http.HandleFunc("/about",About)
	// http.HandleFunc("/post",Posthandler)
	// http.HandleFunc("/post/",QueryHandler)
	// http.HandleFunc("/user/:id",DynamicParams)
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/hello", MyHellofun)
	http.HandleFunc("/form", HandleForm)
	http.HandleFunc("/json",JsonResponse)
    http.HandleFunc("/jsonpost",PostrequestJsonbody)

	 //task handler


	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	fmt.Println("server is getting started")
}
func JsonResponse(w http.ResponseWriter,req *http.Request){
   var response =Response{
	    Message: "Hello World, this is simple json",
 
   }
   json.NewEncoder(w).Encode(response)
}
func PostrequestJsonbody(w http.ResponseWriter,req *http.Request){
	fmt.Println(req.Body)
	if req.Method!=http.MethodPost{
		http.Error(w,"method not allowed",http.StatusBadRequest)
        return     
	}
   var user User
   err:=json.NewDecoder(req.Body).Decode(&user);
   if err!=nil{
	http.Error(w,"invalid json ",http.StatusBadRequest)
   }
  
   if user.Email==""{
	http.Error(w,"email is required",http.StatusBadRequest)
	return
   }
   if user.Name==""{
	http.Error(w,"name is required",http.StatusBadRequest)
	return
   }
   if user.Password==""{
	http.Error(w,"password is required",http.StatusBadRequest)
	return
   }
   var response =make(map[string]interface{})
  response["data"]=user
  response["message"]="Post request successful"
  response["method"]=req.Method
   w.Header().Set("content-type","application/json")
  w.Header().Set("mykey","suvadip maiti")
  json.NewEncoder(w).Encode(response)

}
func HandleForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Parse form err: %v", err)
		return
	}
	//form encoded
	name := r.FormValue("name")
	email := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Fprintf(w,"form is successful")
	fmt.Fprintf(w, "name: %s email: %s password: %s", name, email, password)
}
func MyHellofun(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "url not found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "welcome to server", http.StatusOK)
}
func DynamicParams(w http.ResponseWriter, req *http.Request) {
	// sid:=strings.TrimPrefix(req.URL.Path)
	fmt.Fprintf(w, "path: ", req.URL.Path)

}
func QueryHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Println("query handler")
	params := req.URL.Query()
	fmt.Fprintf(w, "query params: %v %v ", params.Get("name"), params.Get("email"))

}

func Hellofun(w http.ResponseWriter, req *http.Request) {
	//    method:=req.URL
	fmt.Fprintf(w, "welcome to server: %v %v", req.Method, req.URL)

}
func About(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "welcome to about page %v %v", req.Method, req.URL)
}
func Posthandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "welcome to post page %v %v\n", req.Method, req.URL)
	fmt.Fprintf(w, "post data: %v", req.PostForm)
}
