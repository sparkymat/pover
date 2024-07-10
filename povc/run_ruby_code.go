package povc

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	logpkg "github.com/sparkymat/pover/log"
)

func (s *Service) runRubyCode(ctx context.Context, codeDir string, input string) (string, error) {
	var err error

	log := logpkg.FromContext(ctx)

	codePath := filepath.Join(codeDir, "app.rb")
	wrappedCode := fmt.Sprintf(`require_relative './pover'

scene do
%s
end
`, input)

	log.Debugf("Writing code to %s", codePath)

	if err = os.WriteFile(codePath, []byte(wrappedCode), 0o600); err != nil { //nolint:mnd
		return "", fmt.Errorf("failed to write file '%s': %w", codePath, err)
	}

	poverPath := filepath.Join(codeDir, "pover.rb")

	log.Debugf("Writing pover.rb to %s", poverPath)

	if err = os.WriteFile(poverPath, s.poverCode, 0o600); err != nil { //nolint:mnd
		return "", fmt.Errorf("failed to write file '%s': %w", poverPath, err)
	}

	cmd := exec.Command("ruby", "app.rb")
	cmd.Dir = codeDir

	log.Debug("Running ruby on rubycode")

	output, err := cmd.Output()
	if err != nil {
		log.Errorf("error running ruby: %v", err)
		log.Debugf("output:\n%s", output)

		return "", fmt.Errorf("failed to run ruby: %w", err)
	}

	return string(output), nil
}
