package povc

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func (s *Service) runRubyCode(_ context.Context, codeDir string, input string) (string, error) {
	var err error

	codePath := filepath.Join(codeDir, "app.rb")
	wrappedCode := fmt.Sprintf(`require_relative './pover'

scene do
%s
end
`, input)

	if err = os.WriteFile(codePath, []byte(wrappedCode), 0o644); err != nil {
		return "", fmt.Errorf("failed to write file '%s': %w", codePath, err)
	}

	poverCode, err := os.ReadFile("app/pover.rb")
	if err != nil {
		return "", fmt.Errorf("failed to read file 'app/pover.rb': %w", err)
	}

	poverPath := filepath.Join(codeDir, "pover.rb")

	if err = os.WriteFile(poverPath, poverCode, 0o644); err != nil {
		return "", fmt.Errorf("failed to write file '%s': %w", poverPath, err)
	}

	cmd := exec.Command("ruby", "app.rb")
	cmd.Dir = codeDir

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("output=%s\n", output)
		return "", fmt.Errorf("failed to run ruby: %w", err)
	}

	return string(output), nil
}
