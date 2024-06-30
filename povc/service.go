package povc

type Config interface {
	StorageFolder() string
}

func New(cfg Config) *Service {
	return &Service{
		cfg: cfg,
	}
}

type Service struct {
	cfg Config
}
