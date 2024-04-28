package config

var (
	logger  *Logger
	secrets *Secrets
)

func Init() error {
	secrets = InitializeSecrets()
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetSecrets() *Secrets {
	return secrets
}
