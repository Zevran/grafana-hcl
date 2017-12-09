package hcl

type TargetPanelConfig struct {
	Name       string `hcl:", key"`
	Datasource string `hcl:"datasource"`

	Expr           string `hcl:"expr"`
	IntervalFactor int    `hcl:"interval_factor"`
	Interval       string `hcl:"interval"`
	Step           int    `hcl:"step"`
	LegendFormat   string `hcl:"legend_format"`

	DsType     string                        `hcl:"ds_type"`
	Metrics    []MetricsTargetPanelConfig    `hcl:"metrics"`
	Query      string                        `hcl:"query"`
	Alias      string                        `hcl:"alias"`
	RawQuery   bool                          `hcl:"raw_query"`
	TimeField  string                        `hcl:"time_field"`
	BucketAggs []BucketAggsTargetPanelConfig `hcl:"bucket_aggs"`

	Target string `hcl:"target"`
}

type MetricsTargetPanelConfig struct {
	Field string `hcl:"field"`
	Type  string `hcl:"type"`
}

type BucketAggsTargetPanelConfig struct {
	Field    string `hcl:"field"`
	Type     string `hcl:"type"`
	Settings struct {
		Interval    string `hcl:"interval"`
		MinDocCount int    `hcl:"min_doc_count"`
	} `hcl:"settings"`
}
