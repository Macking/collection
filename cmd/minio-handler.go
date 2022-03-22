package cmd

import (
	"fmt"
	"github.com/Macking/collection/internal/dbcore"
	"github.com/Macking/collection/internal/repo"
	"github.com/gorilla/mux"
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

	fmt.Println("md5")
	db := dbcore.GetDB(ctx)
	//db.Find(nil)
	rec := repo.File{MD5: "bf7c3fecfe5dcffceb170b2aa6d34c31", Name: "designing.pdf", Path: "/etc/mnt", Key: "/minio/designing.pdf", Updated: time.Now()}
	db.Create(&rec)

}
