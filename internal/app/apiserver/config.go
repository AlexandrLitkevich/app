package apiserver

// Congig file fornat
type Config struct {
	/*
		формат toml
		BindAddr адрес где запускаем сервер
		LogLevel логирование нашего приложения
	*/
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

/*
Отдает дефлотный конфиг
*/
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
