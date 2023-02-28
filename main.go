package main

import (
	"crypto/md5"
	"embed"
	"encoding/hex"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spotbun/filedrop/pkg/config"
)

//go:embed web/templates
var templatesFS embed.FS

func main() {
	log.Println("Starting server...")

	err := os.Mkdir("./uploads", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(fmt.Errorf("error creating uploads directory: %w", err))
		return
	}

	// parse templates from file system
	templates := template.Must(template.ParseFS(templatesFS, "web/templates/*.gohtml"))

	// define handler
	http.HandleFunc("/", renderTemplate(templates, "index.gohtml", config.GetConfig()))

	// upload handler
	http.HandleFunc("/upload", handleUpload)

	log.Println("Server started on http://0.0.0.0:80")

	// start server
	log.Fatal(http.ListenAndServe(":80", nil))
}

func renderTemplate(templates *template.Template, name string, data any) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// execute template
		err := templates.ExecuteTemplate(w, name, data)
		if err != nil {
			log.Println(err)
		}
	}
}

// handleUpload handles file uploads. There can be multiple files uploaded at once.
// File names should be prefixed with their hash.
func handleUpload(w http.ResponseWriter, r *http.Request) {
	// parse multipart form
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		log.Println(fmt.Errorf("error parsing multipart form: %w", err))
		return
	}

	// get files
	for _, f := range r.MultipartForm.File {
		for _, file := range f {
			log.Printf("Uploading file: %s", file.Filename)
			// open file
			f, err := file.Open()
			if err != nil {
				log.Println(fmt.Errorf("error opening file: %w", err))
				return
			}

			// calculate hash for file
			hash, err := hashFile(f)
			if err != nil {
				log.Println(fmt.Errorf("error hashing file: %w", err))
				return
			}

			err = os.Mkdir("./uploads/"+hash, 0777)
			if err != nil && !os.IsExist(err) {
				log.Fatal(fmt.Errorf("error creating hashed uploads directory: %w", err))
				return
			}

			// create new file
			newFile, err := os.Create(fmt.Sprintf("./uploads/%s/%s", hash, file.Filename))
			if err != nil {
				log.Println(fmt.Errorf("error creating new file: %w", err))
				return
			}

			// copy file
			_, err = io.Copy(newFile, f)
			if err != nil {
				log.Println(fmt.Errorf("error copying file: %w", err))
				return
			}

			f.Close()
			newFile.Close()
		}
	}
}

func hashFile(f io.Reader) (string, error) {
	hash := md5.New()
	if _, err := io.Copy(hash, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
