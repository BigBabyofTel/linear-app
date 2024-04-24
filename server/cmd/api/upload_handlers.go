package main

import (
	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
	"net/http"
	"strings"
)

func (a *app) uploadImageHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		a.badRequestResponse(w, r, err)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		a.badRequestResponse(w, r, err)
		return
	}
	defer file.Close()

	c, err := a.AWS.Client()
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	ext := strings.Split(handler.Header.Get("Content-Type"), "/")[1]
	name := id.String() + "." + ext
	bucketName := a.config.AwsBucketName
	region := a.config.AwsRegion
	url := "https://" + bucketName + ".s3." + region + ".amazonaws.com/" + name

	_, err = c.PutObject(r.Context(), bucketName, name, file, handler.Size, minio.PutObjectOptions{ContentType: handler.Header.Get("Content-Type")})
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
	err = a.writeJSON(w, http.StatusCreated, envelope{"url": url}, nil)
	if err != nil {
		a.serverErrorResponse(w, r, err)
		return
	}
}
