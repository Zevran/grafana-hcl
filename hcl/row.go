package hcl

type Panel interface {
}

type RowConfig struct {
	Name      string  `hcl:",key"`
	Title     string  `hcl:"title"`
	ShowTitle bool    `hcl:"show_title"`
	Collapse  bool    `hcl:"collapse"`
	Editable  bool    `hcl:"editable"`
	Height    int32   `hcl:"height"`
	Panels    []Panel `hcl:"panels"`
}

func (d *RowConfig) ValidateHCL() error {
	return nil
}

func (d *RowConfig) GenerateJSON() ([]byte, error) {
	return nil, nil
}
