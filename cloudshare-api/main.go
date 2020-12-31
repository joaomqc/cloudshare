package main

import (
	"net/http"

	"github.com/joaomqc/cloudshare/controllers"

	"github.com/go-chi/chi"
)

// var photosFolder string
// var photo string

// // File is the exported file info
// type File struct {
// 	Name  string `json:"name"`
// 	IsDir bool   `json:"isDir"`
// 	Size  int64  `json:"size"`
// }

// func mapFileInfo(fileInfoList []os.FileInfo) []File {
// 	files := make([]File, len(fileInfoList))

// 	for i, f := range fileInfoList {
// 		files[i] = File{
// 			Name:  f.Name(),
// 			IsDir: f.IsDir(),
// 			Size:  f.Size(),
// 		}
// 	}

// 	return files
// }

func main() {
	// photosFolder = "D:/Users/Joao/Desktop/text/facebook-joaocoelho737/photos/"

	router := chi.NewRouter()

	// router.Get("/", handleRoot)
	// router.Get("/{filename}", handleFile)

	router.Post("/user", controllers.HandlePostUser)
	router.Get("/user", controllers.HandleGetUsers)
	router.Get("/user/{username}", controllers.HandleGetUserByUsername)

	router.Post("/login", controllers.HandlePostLogin)

	http.ListenAndServe(":7070", router)
}

// func handleRoot(w http.ResponseWriter, r *http.Request) {
// 	fileInfoList, _ := ioutil.ReadDir(photosFolder)

// 	files := mapFileInfo(fileInfoList)

// 	json.NewEncoder(w).Encode(files)
// }

// func handleFile(w http.ResponseWriter, r *http.Request) {
// 	filename := chi.URLParam(r, "filename")

// 	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
// 	w.Header().Set("Content-Type", r.Header.Get("Content-Type"))

// 	reader, _ := os.Open(photosFolder + filename)

// 	io.Copy(w, reader)
// }
