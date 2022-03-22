package cmd

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/Macking/collection/internal/dbcore"
	"github.com/Macking/collection/internal/minio"
	"github.com/Macking/collection/internal/repo"
	"github.com/gorilla/mux"
	"io"
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
	//rec := repo.File{MD5: "bf7c3fecfe5dcffceb170b2aa6d34c31", Name: "designing.pdf", Path: "/etc/mnt", Key: "/minio/designing.pdf", Updated: time.Now()}
	//db.Create(&rec)
	//result := db.Where("md5 IN ?", targets).Find(&files)
	db.Where("md5 IN ?", targets).Find(&files)
	//fmt.Println("result.RowsAffected ", result.RowsAffected)
	//fmt.Println(files)
	var queryResult []minio.ResultMessage
	for _, f := range files {
		queryResult = append(queryResult, minio.ResultMessage{
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
	uploadFile, _, err := r.FormFile("uploadfile")
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

	db := dbcore.GetDB(ctx)
	var destFile repo.File
	query := db.Where("md5 = ?", fileMD5).First(&destFile)
	if query.RowsAffected > 0 {
		queryResult := minio.ErrorMessage{
			Message: "File already exist in OSS",
		}
		jsonBytes, _ := json.Marshal(queryResult)
		writeResponseJSON(w, jsonBytes, http.StatusNotAcceptable)
		return
	}
	db.Create(&repo.File{
		MD5:     fileMD5,
		Path:    path,
		Bucket:  bucket,
		Key:     key,
		Name:    fileName,
		Updated: time.Now(),
	})
	createResult := minio.ResultMessage{
		MD5:    fileMD5,
		Bucket: bucket,
		Key:    key,
	}
	jsonBytes, _ := json.Marshal(createResult)
	writeResponseJSON(w, jsonBytes, http.StatusOK)
	return
}
