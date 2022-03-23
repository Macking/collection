package miniocore

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
}

func defaultMinioConfig(cfg *MinioConfig) *MinioConfig {
	newCfg := *cfg
	newCfg.Endpoint = "192.168.0.64:9000"
	newCfg.AccessKeyID = ""
	newCfg.SecretAccessKey = ""
	return &newCfg
}
