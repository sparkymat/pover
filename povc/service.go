package povc

type Config interface {
	StorageFolder() string
}

func New(cfg Config, poverCode []byte) *Service {
	return &Service{
		cfg:       cfg,
		poverCode: poverCode,
	}
}

type Service struct {
	cfg       Config
	poverCode []byte
}
