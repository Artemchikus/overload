package apiserver

type Config struct {
	BindAddr string `toml:"bind_addr"` // порт на котором запускается сервер
}

// функция фозвращения деволтного конфига
func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
	}
}