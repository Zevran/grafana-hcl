package hcl

type DataSourceConfig struct {
	Name string `hcl:",key"`
	Type string `hcl:"type"`
}
