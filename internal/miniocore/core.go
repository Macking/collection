package miniocore

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	globalMinioClient *minio.Client
	globalMinioConfig *MinioConfig
)

func Connect(cfg *MinioConfig) {
	cfg = defaultMinioConfig(cfg)
	globalMinioConfig = cfg
	minioClient, _ := minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""),
		Secure: false,
	})
	globalMinioClient = minioClient
}

func GetMinioClient(ctx context.Context) *minio.Client {
	return globalMinioClient

}
