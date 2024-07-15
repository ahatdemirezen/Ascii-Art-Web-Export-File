package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/ahatdemirezen/Ascii-Art-Web-Export-File/printasciiart"
)

func AsciiArtWriter(argStr string, bannerType string) string {
	sepArgs := strings.Split(argStr, "\n")

	file, err := os.Open(bannerType + ".txt")
	if err != nil {
		fmt.Println("Error")
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}

	res := printasciiart.PrintAsciiArt(sepArgs, lines)
	return res
}

func AsciiToWeb(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":

		http.ServeFile(w, r, "templates/form.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		str := r.FormValue("str")
		banner := r.FormValue("banner")
		if str == "" {
			str = "String Is Empty!"
		}
		_, err := os.Open(banner + ".txt")
		if err != nil {
			fmt.Fprintf(w, "%s\n", "Banner is not exist!")
			return
		}
		for _, char := range str {
			if char >= 1 && char <= 127 {
				continue
			} else {
				str = "Invalid Character."
			}
		}

		res := AsciiArtWriter(str, banner)
		download := r.FormValue("download")

		if download == "" {
			// Write the ASCII art to the response body
			fmt.Fprintf(w, "%s\n", res)
			return
		}

		w.Header().Set("Content-Disposition", "attachment; filename=ascii_art.txt")
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\n", res)
		return

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.Handle("/templates/style.css", http.StripPrefix("/templates/", http.FileServer(http.Dir("templates"))))
	http.HandleFunc("/", AsciiToWeb)
	fmt.Printf("Starting server for testing HTTP POST...\n")

	if err := http.ListenAndServe(":8080", http.DefaultServeMux); err != nil {
		fmt.Println("Something went wrong!")
		return
	}
}
