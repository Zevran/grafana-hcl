package hcl

type TablePanelConfig struct {
	Columns   map[string]interface{} `hcl:"columns"`
	Sort      SortTablePanelConfig   `hcl:"sort"`
	Styles    map[string]interface{} `hcl:"styles"`
	Transform string                 `hcl:"transform"`
	Targets   []TargetPanelConfig    `hcl:"targets"`
	Scroll    bool                   `hcl:"scroll"`
}

type SortTablePanelConfig struct {
	Col  uint `hcl:"col"`
	Desc bool `hcl:"desc"`
}
