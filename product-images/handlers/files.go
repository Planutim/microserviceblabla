package handlers

import (
	"example/files"
	"io"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	hclog "github.com/hashicorp/go-hclog"
)

//Files is a handler for reading and writing files
type Files struct {
	log   hclog.Logger
	store files.Storage
}

//NewFiles create a new File Handler
func NewFiles(s files.Storage, l hclog.Logger) *Files {
	return &Files{store: s, log: l}
}

//UploadREST implements the http.Handler interface
func (f *Files) UploadREST(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fn := vars["filename"]

	f.log.Info("Handle POST", "id", id, "filename", fn)

	// no need to check for invalid id or filename as the mux router will not send requests
	// here unless they have the correct parameters
	f.saveFile(id, fn, w, r.Body)
}
func (f *Files) UploadMultipart(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(120 * 1024)
	if err != nil {
		f.log.Error("bad request", "error", err)
		http.Error(w, "Expected multipart form data", http.StatusBadRequest)
	}

	id, idErr := strconv.Atoi(r.FormValue("id"))
	if idErr != nil {
		f.log.Error("bad request", "error", err)
		http.Error(w, "Expected integer id", http.StatusBadRequest)
		return
	}
	f.log.Info("Process form for id", "id", id)

	ff, fh, err := r.FormFile("file")
	if err != nil {
		f.log.Error("bad request", "error", err)
		http.Error(w, "Expected file", http.StatusBadRequest)
		return
	}

	f.saveFile(r.FormValue("id"), fh.Filename, w, ff)

}
func (f *Files) invalidURI(uri string, w http.ResponseWriter) {
	f.log.Error("Invalid path", "path", uri)
	http.Error(w, "Invalid file path should be in the format: /[id]/[filepath]", http.StatusBadRequest)
}

func (f *Files) saveFile(id, path string, w http.ResponseWriter, r io.ReadCloser) {
	f.log.Info("Save file for product", "id", id, "path", path)
	fp := filepath.Join(id, path)
	err := f.store.Save(fp, r)
	if err != nil {
		f.log.Error("Unable to save file", "error", err)
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
	}
}
