package hcl

type BoolStringConfig struct {
	Flag  bool   `hcl:"flag"`
	Value string `hcl:"value"`
}

type IntStringConfig struct {
	Value int64 `hcl:"value"`
	Valid bool  `hcl:"valid"`
}

type TooltipConfig struct {
	Shared       bool   `hcl:"shared"`
	ValueType    string `hcl:"value_type"`
	MsResolution bool   `hcl:"ms_resolution"`
	Sort         int    `hcl:"sort"`
}
