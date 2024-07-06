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

	if err = os.WriteFile(codePath, []byte(wrappedCode), 0o600); err != nil { //nolint:mnd
		return "", fmt.Errorf("failed to write file '%s': %w", codePath, err)
	}

	poverPath := filepath.Join(codeDir, "pover.rb")

	if err = os.WriteFile(poverPath, s.poverCode, 0o600); err != nil { //nolint:mnd
		return "", fmt.Errorf("failed to write file '%s': %w", poverPath, err)
	}

	cmd := exec.Command("ruby", "app.rb")
	cmd.Dir = codeDir

	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("output=%s\n", output) //nolint:forbidigo

		return "", fmt.Errorf("failed to run ruby: %w", err)
	}

	return string(output), nil
}
