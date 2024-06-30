package povc

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

func (s *Service) Compile(ctx context.Context, rubyCode string) (string, error) {
	randomUUID := uuid.New().String()

	codeDir, err := os.MkdirTemp("", "pover")
	if err != nil {
		return "", fmt.Errorf("failed to create temp dir: %w", err)
	}

	defer func() {
		_ = os.RemoveAll(codeDir)
	}()

	output, err := s.runRubyCode(ctx, codeDir, rubyCode)
	if err != nil {
		return "", fmt.Errorf("failed to run ruby code: %w", err)
	}

	inImagePath, err := s.runPOVRay(ctx, codeDir, output)
	if err != nil {
		return "", fmt.Errorf("failed to run povray: %w", err)
	}

	outimagePath := filepath.Join(s.cfg.StorageFolder(), randomUUID+".png")

	inImage, err := os.Open(inImagePath)
	if err != nil {
		return "", fmt.Errorf("failed to open image file: %w", err)
	}

	defer inImage.Close()

	outImage, err := os.Create(outimagePath)
	if err != nil {
		return "", fmt.Errorf("failed to create output image file: %w", err)
	}

	defer outImage.Close()

	if _, err := io.Copy(outImage, inImage); err != nil {
		return "", fmt.Errorf("failed to copy image file: %w", err)
	}

	return randomUUID + ".png", nil
}
