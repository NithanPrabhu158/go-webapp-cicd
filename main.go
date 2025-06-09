package main

import (
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	// Render the home html page from static folder
	http.ServeFile(w, r, "static/home.html")
}

func coursePage(w http.ResponseWriter, r *http.Request) {
	// Render the course html page
	http.ServeFile(w, r, "static/courses.html")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	// Render the about html page
	http.ServeFile(w, r, "static/about.html")
}

func contactPage(w http.ResponseWriter, r *http.Request) {
	// Render the contact html page
	http.ServeFile(w, r, "static/contact.html")
}

// func main() {

// 	http.HandleFunc("/home", homePage)
// 	http.HandleFunc("/courses", coursePage)
// 	http.HandleFunc("/about", aboutPage)
// 	http.HandleFunc("/contact", contactPage)

// 	err := http.ListenAndServe("0.0.0.0:8080", nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
    // Serve static files from 2106_soft_landing at /soft/
    fs := http.FileServer(http.Dir("landing_page"))
    http.Handle("/", http.StripPrefix("/", fs))

    // Optional: Redirect root to the landing page
    // http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    //     if r.URL.Path == "/" {
    //         http.Redirect(w, r, "/soft/index.html", http.StatusFound)
    //         return
    //     }
    //     http.NotFound(w, r)
    // })

    log.Println("Server started at http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal(err)
    }
}