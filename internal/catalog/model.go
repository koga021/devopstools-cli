package catalog

type Catalog struct {
	VersionClient string     `yaml:"versionClient"`
	VersionServer string     `yaml:"versionServer"`
	Menu          []MenuItem `yaml:"menu"`
}

type MenuItem struct {
	ID          string     `yaml:"id"`
	Name        string     `yaml:"name"`
	Description string     `yaml:"description,omitempty"`
	Items       []MenuItem `yaml:"items,omitempty"`
	Source      *Source    `yaml:"source,omitempty"`
}

type Source struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
	Ref  string `yaml:"ref"`
}

type Config struct {
	DevOpsToolsTarget string `env:"DEVOPS_TOOLS_TARGET" envDefault:""`
	DevOpsToolsEnv    string `env:"DEVOPS_TOOLS_ENV" envDefault:"production"`
}
