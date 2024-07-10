package povc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	logpkg "github.com/sparkymat/pover/log"
)

func (s *Service) Compile(ctx context.Context, rubyCode string) (string, error) {
	log := logpkg.FromContext(ctx)

	randomUUID := uuid.New().String()

	log.Debug("Creating temp dir")

	codeDir, err := os.MkdirTemp("", "pover")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}

	defer func() {
		_ = os.RemoveAll(codeDir)
	}()

	log.Debug("Running ruby code")

	output, err := s.runRubyCode(ctx, codeDir, rubyCode)
	if err != nil {
		return "", fmt.Errorf("failed to run ruby code: %w", err)
	}

	log.Debug("Running POVRay")

	inImagePath, err := s.runPOVRay(ctx, codeDir, output)
	if err != nil {
		return "", fmt.Errorf("failed to run povray: %w", err)
	}

	outimagePath := filepath.Join(s.cfg.StorageFolder(), randomUUID+".png")

	inImage, err := os.Open(inImagePath) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}

	defer inImage.Close() //nolint:errcheck

	outImage, err := os.Create(outimagePath) //nolint:gosec
	if err != nil {
		return "", fmt.Errorf("failed to create output image file: %w", err)
	}

	defer outImage.Close() //nolint:errcheck

	log.Debug("Copying image file")

	if _, err := io.Copy(outImage, inImage); err != nil {
		return "", fmt.Errorf("failed to copy image file: %w", err)
	}

	return randomUUID + ".png", nil
}
