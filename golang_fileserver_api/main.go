package main

import (
    "log"
	"net/http"
    "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//handle homepage
	router.HandleFunc("/", home).Methods("GET") 

	//handle root directory
	router.HandleFunc("/files/page/{page}", files).Methods("GET")
	router.HandleFunc("/files/sorted/page/{page}", filesSorted).Methods("GET")
	router.HandleFunc("/files/{filename}", downloadFile).Methods("GET")

	//handle sub directories
	router.HandleFunc("/files/directory/{directory}/page/{page}", directoryFiles).Methods("GET")
	router.HandleFunc("/files/sorted/directory/{directory}/page/{page}", directoryFilesSorted).Methods("GET")
	router.HandleFunc("/files/directory/{directory}/{filename}", directoryDownloadFile).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}