package apiserver

import "github.com/AlexandrLitkevich/app/internal/app/store"

// Congig file fornat
type Config struct {
	/*
		формат toml
		BindAddr адрес где запускаем сервер
		LogLevel логирование нашего приложения
		Store config data base
	*/
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

/*
Отдает дефлотный конфиг
*/
func NewConfig() *Config {
	return &Config{
		// Деволтные значения
		BindAddr: ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
