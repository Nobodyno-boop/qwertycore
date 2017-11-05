package qwertycore

import (
	"fmt"
	"net/http"
)

var a = ""

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	if _, err := LoadPage(title); err != nil {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
	p, _ := LoadPage(title)

	fmt.Fprintf(w, "%s", p.Body)
}

func def(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<meta http-equiv="X-UA-Compatible" content="ie=edge">
			<link rel="stylesheet" href="http://qwertycore.com/assets/css/semantic.min.css">
			<link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/semantic-ui/2.2.13/components/icon.min.css">
			<link rel="stylesheet" href="http://qwertycore.com/assets/css/style.css">
			<link rel="icon" href="http://qwertycore.com/assets/fonts/penta.ico">
			<title> QWTOOL | QwertyCore </title>
		</head>
		
		<body>
			<div class="ui middle aligned center aligned grid" id="u">
				<div class="column">
					<h2 class="ui teal image header">
						<div class="image">
							<img src="http://qwertycore.com/assets/fonts/penta.svg" alt="blabla">
						</div>
						<div class="content">
							Welcome !
						</div>
						<div class="sub header">
							<p class="white"> use the <a href="/view/">views url for see you views path !</a></p>
						   <p class="white "> Made with <i class="heart icon red pulse"></i></p>                    
						</div>
					</h2>
				</div>
			</div>
		</body>
		</html>`)
}

//
func Server(o *Oss) error {
	a = o.Qw
	http.HandleFunc("/", def)
	http.HandleFunc("/view/", viewHandler)
	z := http.ListenAndServe(":4242", nil)
	fmt.Println("The local server was started on :4242")
	fmt.Println("http://localhost:4242")
	return z
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "404")
	}
}

//
func InitDir(o *Oss) {
	o.mkdir("")
	o.mkdir("/views")
	o.mkdir("/assets")
	o.mkdir("/assets/js")
	o.mkdir("/assets/css")
	o.mkdir("/assets/tpl")
}
