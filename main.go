package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/wveik/WebAppOnGo/models"
)

var posts map[string]*models.Post

func indexHandler (w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}
	
	t.ExecuteTemplate(w, "index", nil)
}

func writeHandler(w http.ResponseWriter, r *http.Request)  {
	t, err := template.ParseFiles("templates/write.html", "templates/header.html", "templates/footer.html")
	if err != nil{
		fmt.Fprintf(w, err.Error())
	}

	t.ExecuteTemplate(w, "write", nil)
}

func savePostHandler(w http.ResponseWriter, r *http.Request)  {
	id := r.FormValue("id")
	title := r.FormValue("title")
	content := r.FormValue("content")

	post := models.NewPost(id, title, content)
	posts[post.Id]= post
}

func main()  {
	fmt.Println("Listening on port :3000")

	posts = make(map[string]*models.Post, 0)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("./assets/"))))
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/write", writeHandler)
	http.HandleFunc("/SavePost", savePostHandler)

	http.ListenAndServe(":3000", nil)
}
