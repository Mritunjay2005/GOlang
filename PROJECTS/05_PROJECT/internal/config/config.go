package config

type HTTPServer struct {
	Addr string
}
type Config struct {
	Env         string `yaml: "env", env	:"ENV" , env-required: "true", env-default: "production"`
	StoragePath string `ymal: "storage_path", env:"STORAGE_PATH", env-required: "true" ,env-default: "storage/storage.db"`
	HTTPServer  `yaml: "http_server"`
}

func MustLoad() {

}
