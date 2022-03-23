package cmd

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/Macking/collection/internal/dbcore"
	"github.com/Macking/collection/internal/miniocore"
	"github.com/Macking/collection/internal/repo"
	"github.com/gorilla/mux"
	"github.com/minio/minio-go/v7"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
	"time"
)

func (a minioAPIHandlers) MinioCheckMD5Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	var targets []string
	for _, queryMd5 := range strings.Split(vars["md5"], ",") {
		if queryMd5 == "" {
			continue
		}
		targets = append(targets, queryMd5)
	}

	var files []repo.File
	db := dbcore.GetDB(ctx)
	//db.Find(nil)
	//rec := repo.File{MD5: "bf7c3fecfe5dcffceb170b2aa6d34c31", Name: "designing.pdf", Path: "/etc/mnt", Key: "/miniocore/designing.pdf", Updated: time.Now()}
	//db.Create(&rec)
	//result := db.Where("md5 IN ?", targets).Find(&files)
	db.Where("md5 IN ?", targets).Find(&files)
	//fmt.Println("result.RowsAffected ", result.RowsAffected)
	//fmt.Println(files)
	var queryResult []miniocore.ResultMessage
	for _, f := range files {
		queryResult = append(queryResult, miniocore.ResultMessage{
			MD5:    f.MD5,
			Bucket: f.Bucket,
			Key:    f.Key,
		})
	}
	jsonBytes, _ := json.Marshal(queryResult)
	writeResponseJSON(w, jsonBytes, http.StatusOK)
	return
}

func (a minioAPIHandlers) MinioUploadFileHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	bucket := vars["bucket"]
	key := vars["key"]
	path := vars["path"]

	tempPaths := strings.Split(strings.Replace(path, "\\", "/", -1), "/")
	fileName := tempPaths[len(tempPaths)-1]

	r.ParseForm()
	uploadFile, header, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println("Parse upload file error ", err)
	}
	defer uploadFile.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, uploadFile); err != nil {
		fmt.Println("Copy ", err)
		return
	}
	result := md5hash.Sum(nil)
	fileMD5 := fmt.Sprintf("%x", result)

	if checkExistFile(w, ctx, fileMD5) {
		return
	}
	saveToMinio(ctx, bucket, key, uploadFile, header)
	saveToDB(w, ctx, fileMD5, path, bucket, key, fileName)

	return
}

func checkExistFile(w http.ResponseWriter, ctx context.Context, fileMD5 string) bool {
	db := dbcore.GetDB(ctx)
	var destFile repo.File
	query := db.Where("md5 = ?", fileMD5).First(&destFile)
	if query.RowsAffected > 0 {
		queryResult := miniocore.ErrorMessage{
			Message: "File already exist in OSS",
		}
		jsonBytes, _ := json.Marshal(queryResult)
		writeResponseJSON(w, jsonBytes, http.StatusNotAcceptable)
		return true
	}
	return false
}

func saveToDB(w http.ResponseWriter, ctx context.Context, fileMD5 string, path string, bucket string, key string, fileName string) {
	db := dbcore.GetDB(ctx)
	db.Create(&repo.File{
		MD5:     fileMD5,
		Path:    path,
		Bucket:  bucket,
		Key:     key,
		Name:    fileName,
		Updated: time.Now(),
	})
	createResult := miniocore.ResultMessage{
		MD5:    fileMD5,
		Bucket: bucket,
		Key:    key,
	}
	jsonBytes, _ := json.Marshal(createResult)
	writeResponseJSON(w, jsonBytes, http.StatusOK)
}

func saveToMinio(ctx context.Context, bucket string, key string, reader io.Reader, header *multipart.FileHeader) {
	client := miniocore.GetMinioClient(ctx)
	fileName := header.Filename
	f := strings.Split(fileName, ".")
	fileExt := f[len(f)-1]
	contentType := "application/octet-stream"
	switch fileExt {
	case "pdf", "zip", "json":
		contentType = "application/" + fileExt
	case "xml":
		contentType = "text/" + fileExt
	case "txt":
		contentType = "text/plain"
	case "png", "jpeg", "tiff", "gif":
		contentType = "image/" + fileExt
	}
	client.PutObject(ctx, bucket, key, reader, header.Size, minio.PutObjectOptions{ContentType: contentType})
}
