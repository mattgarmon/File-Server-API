Matt Garmon - 2018

File Browser API



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

a) Homepage 
    1) Display welcome page
    2) List possible URLs and their actions
b) Browse files
    1) Determine root directory or sub directory
    2) Read file info for each file in the reqested directory
    3) Sort by filename if requested
    4) Page out the files for easier browsing
    5) Serve the requested page with a list of files and current page number
c) Download file
    1) Determine if the file is in the root directorty or sub directory
    2) Serve the file contents with a new modification time

Note: The mux will handle all other HTTP requests and display a 404 page not
found error if neccesary.



TESTS:

Some of the tests that would be needed to properly test this API are as follows:

1)....



MISSING REQUERMENTS:

Some requirements I think are missing from this API design are as follows:

1)....
