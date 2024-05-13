package config

import (
	"os"
	"testing"

	"github.com/qnify/api-server/modules/auth"
	. "github.com/qnify/api-server/utils/helper"
)

func TestReadConfigFromFile(t *testing.T) {

	const testYamlContent = `
auth:
  origin: "origin"
port: 3000
`
	var expectedConfig = Config{
		Auth: auth.AuthConfig{Origin: "origin"},
		Port: 3000,
	}

	const test_file = "test_config.yaml"
	err := os.WriteFile(test_file, []byte(testYamlContent), 0644)
	NoErr(err, t)

	defer func() {
		err := os.Remove(test_file)
		NoErr(err, t)
	}()

	actualConfig := LoadConfig(test_file)
	Check(actualConfig.Auth.Origin == expectedConfig.Auth.Origin, t)
	Check(actualConfig.Port == expectedConfig.Port, t)
}
