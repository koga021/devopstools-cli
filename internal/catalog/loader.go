package catalog

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (*Catalog, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler catálogo: %w", err)
	}

	var result Catalog

	if err := yaml.Unmarshal(data, &result); err != nil {
		return nil, fmt.Errorf("erro ao interpretar YAML: %w", err)
	}

	return &result, nil
}
