(c) Matt Garmon - 2018

Golang File Browser API (REST)



OPERATION:

To compile and run the application simply use the command "go run *.go" without the
quotations from the directory containing main.go. This will then host the server 
application on localhost:8080. Simply add files to the root directory folder or
create sub directories inside of it (up to one level) to host the files on the API.

Note: The gorilla mux Go package is a dependency for this project. Get it by using the
command "go get github.com/gorilla/mux" without the quotations.



USAGE:

/                                                 homepage

/files/page/{page}                                view page in root directory
/files/directory/{directory}/page/{page}          view page in sub directory

/files/sorted/page/{page}                         view page in sorted root directory
/files/sorted/directory/{directory}/page/{page}   view page in sorted sub directory

/files/{filename}                                 download file in root directory
/files/directory/{directory}/{filename}           download file in sub directory



STRUCTURE:

The application utilizes the gorilla mux Go package to route all HTTP requests. 
When a page request is recieved the application first routes it through the mux 
to determine what the request is:

a) Homepage (GET)
    1) Display welcome page
    2) List possible URLs and their actions
b) Browse files (GET)
    1) Determine root directory or sub directory
    2) Read file info for each file in the reqested directory
    3) Sort by filename if requested
    4) Page out the files for easier browsing (defualt: 20 per page)
    5) Serve the requested page with a list of files and current page number
c) Download file (GET)
    1) Determine if the file is in the root directorty or sub directory
    2) Serve the file contents with a new modification time

Note: The mux will handle all other HTTP requests and display a 404 page not
found error if neccesary. Also, the implementation of the download file function
will allow for partial downloading on the client side.



TESTS:

Some of the tests that would be needed to verify functionality of this API are as follows:

1) Test many different filetypes and verify that they are downloaded as intended.
2) Verify that the file system works in it's extreme case. 
   (100K files, 100k directories, 10 GB files)
3) Test many URLs to verify that the routing is handled correctly for all cases.
4) Use a API testing tool such as Postman to verify that the response codes, messages, 
   and bodies are as intended.



MISSING REQUERMENTS:

Some requirements I think are missing from this API design are as follows:

1) To have a proper file system API, the client should also be able to upload files (PUT),
   but this depends on the intended usage.
2) Allow the file system work with files regardless of the filename (Not just numbers)



CLOSING REMARK:

This was my first time developing an API and using the Go programming language. I have learned 
a lot in the process and have had a lot of excitement using these tools. I am sure there are
minor flaws in this implementation, but I am hoping to continue to work on these types of program
to further my knowlege on the subject.