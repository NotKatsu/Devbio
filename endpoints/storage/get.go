package storage

import (
	"devbio/database"
	ReturnModule "devbio/modules/return_module"
	"fmt"
	"net/http"
	"strings"
)

func GetRequest(w http.ResponseWriter, r *http.Request) {
	segments := strings.Split(r.URL.Path, "/")

	username := segments[len(segments)-1]

	var image []byte
	var imageDatabaseError error

	image, imageDatabaseError = database.GetProfileImageByUsername(username)

	if imageDatabaseError != nil {
		ReturnModule.InternalServerError(w, r)
	}

	w.Header().Set("Content-Type", "image/jpeg")
	w.Header().Set("Content-Length", fmt.Sprint(len(image)))

	_, imageDatabaseError = w.Write(image)

	if imageDatabaseError != nil {
		ReturnModule.InternalServerError(w, r)
	}
}
