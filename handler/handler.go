package handler

import (
	"ascii-art-web/logic"
	"html/template"
	"log"
	"net/http"
)

var ts = template.Must(template.ParseFiles("templates/index.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Println("Wrong Method: needs GET method")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if err := ts.Execute(w, nil); err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Print("Wrong Method: POST Method allowed only")
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/ascii-art" {
		http.Error(w, "Path Not Found", http.StatusNotFound)
		return
	}

	text := r.FormValue("text")
	banner := r.FormValue("banner")

	if text == "" {
		log.Print("Empty Banner")
		http.Error(w, "Input a string .", http.StatusBadRequest)
		return
	}

	if banner == "" {
		log.Print("Empty banner")
		http.Error(w, "Select a banner", http.StatusBadRequest)
		return
	}

	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		log.Print("Empty banner")
		http.Error(w, "Wrong banner", http.StatusBadRequest)
		return
	}

	bannerPath := "banner/" + banner + ".txt"
	log.Println(bannerPath)
	data, err := logic.Generate(text, bannerPath)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "input contains non-ASCII characters", http.StatusBadRequest)
		return
	}
	log.Println(data)

	if err := ts.Execute(w, data); err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
