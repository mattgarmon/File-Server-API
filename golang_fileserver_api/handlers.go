package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"bytes"
	"strconv"
    "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the file server API!\n")
	fmt.Fprintln(w, "Usage:\n")
	fmt.Fprintln(w, "URL                                               DESCRIPTION\n")
	fmt.Fprintln(w, "/files/page/{page}                                view page in root directory")
	fmt.Fprintln(w, "/files/directory/{directory}/page/{page}          view page in sub directory\n")
	fmt.Fprintln(w, "/files/sorted/page/{page}                         view page in sorted root directory")
	fmt.Fprintln(w, "/files/sorted/directory/{directory}/page/{page}   view page in sorted sub directory\n")
	fmt.Fprintln(w, "/files/{filename}                                 download file in root directory")
	fmt.Fprintln(w, "/files/directory/{directory}/{filename}           download file in sub directory")
}

//view a file page in the root directory
func files(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageRequest, err := strconv.Atoi(vars["page"])
	if err != nil {
		fmt.Fprintf(w, "Invalid page request")
		return
	}

	fmt.Fprintf(w, formatPage("./root/", pageRequest, false))
}

//view a file page in the sorted root directory
func filesSorted(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageRequest, err := strconv.Atoi(vars["page"])
	if err != nil {
		fmt.Fprintf(w, "Invalid page request")
		return
	}

	fmt.Fprintf(w, formatPage("./root/", pageRequest, true))
}

//view a file page in a sub directory
func directoryFiles(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	path := "./root_directory/" + vars["directory"] + "/"
	pageRequest, err := strconv.Atoi(vars["page"])
	if err != nil {
        fmt.Fprintf(w, "Invalid page request")
		return
	}

	fmt.Fprintf(w, formatPage(path, pageRequest, false))
}

//view a file page in a sorted sub directory
func directoryFilesSorted(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	pageRequest, err := strconv.Atoi(vars["page"])
	path := "./root/" + vars["directory"] + "/"
	if err != nil {
        fmt.Fprintf(w, "Invalid page request")
		return
	}

	fmt.Fprintf(w, formatPage(path, pageRequest, true))
}

//download a file from the root directory
func downloadFile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	filepath := "./root/" + filename
	modtime := time.Now()
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(w, "File not found")
		return
	}

	w.Header().Add("Content-Disposition", "Attachment")
	http.ServeContent(w, r, filename, modtime, bytes.NewReader(content))
}

//download a file from a directory
func directoryDownloadFile(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	filename := vars["filename"]
	directoryname := vars["directory"]
	filepath := "./root/" + directoryname + "/" + filename
	modtime := time.Now()
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Fprintf(w, "File not found")
		return
	}

	w.Header().Add("Content-Disposition", "Attachment")
	http.ServeContent(w, r, filename, modtime, bytes.NewReader(content))
}