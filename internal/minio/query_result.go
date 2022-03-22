package minio

type ResultMessage struct {
	MD5    string `json:"md5,omitempty"`
	Bucket string `json:"bucket,omitempty"`
	Key    string `json:"key,omitempty"`
}

type ErrorMessage struct {
	Message string `json:"message,omitempty"`
}
