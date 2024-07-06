package povc

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (*Service) runPOVRay(_ context.Context, codeDir string, input string) (string, error) {
	var err error

	codePath := filepath.Join(codeDir, "image.pov")

	if err = os.WriteFile(codePath, []byte(input), 0o600); err != nil { //nolint:mnd
		return "", fmt.Errorf("failed to write file '%s': %w", codePath, err)
	}

	cmd := exec.Command("povray", "+Oimage.png", "image.pov")
	cmd.Dir = codeDir

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("output: %s\n", output) //nolint:forbidigo

		return "", fmt.Errorf("failed to run povray: %w", err)
	}

	return filepath.Join(codeDir, "image.png"), nil
}
