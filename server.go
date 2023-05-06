package main

import (
    "fmt"
    "log"
    "net/http"
)

func formhandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        name := r.FormValue("name")
        email := r.FormValue("email")
        message := r.FormValue("message")

        // Do something with the form data, such as saving it to a database or sending an email
        fmt.Printf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)

        // Sending a response to the user
        fmt.Fprintf(w, "Thank you for your message!")
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func hellohandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "method is not supported", http.StatusMethodNotAllowed)
        return
    }
    fmt.Fprintf(w, "hello")
}

func main() {
    fileserver := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileserver)
    http.HandleFunc("/form", formhandler)
    http.HandleFunc("/hello", hellohandler)

    fmt.Print("starting server at port 8080\n")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
