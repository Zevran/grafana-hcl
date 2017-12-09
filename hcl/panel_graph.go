package hcl

type GraphPanelConfig struct {
	Name            string                           `hcl:",key"`
	AliasColors     interface{}                      `hcl:"alias_colors"`
	Bars            bool                             `hcl:"bars"`
	Fill            int32                            `hcl:"fill"`
	Legend          LegendGraphPanelConfig           `hcl:"legent"`
	LeftYAxisLabel  string                           `hcl:"left_y_axis_Label"`
	RightYAxisLabel string                           `hcl:"righty_axis_label"`
	Lines           bool                             `hcl:"lines"`
	Linewidth       uint                             `hcl:"line_width"`
	NullPointMode   string                           `hcl:"null_point_mode"`
	Percentage      bool                             `hcl:"percentage"`
	Pointradius     int                              `hcl:"point_radius"`
	Points          bool                             `hcl:"points"`
	SeriesOverrides []SeriesOverrideGraphPanelConfig `hcl:"series_overrides"`
	Stack           bool                             `hcl:"stack"`
	SteppedLine     bool                             `hcl:"stepped_line"`
	Targets         []TargetPanelConfig              `hcl:"targets"`
	TimeFrom        string                           `hcl:"time_from"`
	TimeShift       string                           `hcl:"time_shift"`
	Tooltip         TooltipConfig                    `hcl:"tooltip"`
	XAxis           bool                             `hcl:"x_axis"`
	YAxis           bool                             `hcl:"y_axis"`
	YFormats        []string                         `hcl:"y_formats"`
	Xaxes           AxisGraphPanelConfig             `hcl:"x_axe"`
	Yaxes           []AxisGraphPanelConfig           `hcl:"y_axes"`
	Decimals        uint                             `hcl:"decimals"`
}

type LegendGraphPanelConfig struct {
	AlignAsTable bool  `hcl:"align_as_table"`
	Avg          bool  `hcl:"avg"`
	Current      bool  `hcl:"current"`
	HideEmpty    bool  `hcl:"hide_empty"`
	HideZero     bool  `hcl:"hide_zero"`
	Max          bool  `hcl:"max"`
	Min          bool  `hcl:"min"`
	RightSide    bool  `hcl:"right_side"`
	Show         bool  `hcl:"show"`
	Total        bool  `hcl:"total"`
	Values       bool  `hcl:"values"`
	SideWidth    *uint `hcl:"side_widthy"`
}

type SeriesOverrideGraphPanelConfig struct {
	Alias         string            `hcl:"alias"`
	Bars          *bool             `hcl:"bars"`
	Color         *string           `hcl:"color"`
	Fill          *int              `hcl:"fill"`
	FillBelowTo   *string           `hcl:"fill_below_to"`
	Legend        *bool             `hcl:"legend"`
	Lines         *bool             `hcl:"lines"`
	Stack         *BoolStringConfig `hcl:"stack"`
	Transform     *string           `hcl:"transform"`
	YAxis         *int              `hcl:"y_axis"`
	ZIndex        *int              `hcl:"z_index"`
	NullPointMode *string           `hcl:"null_point_mode"`
}

type AxisGraphPanelConfig struct {
	Format  string           `hcl:"format"`
	LogBase int              `hcl:"logBase"`
	Max     *IntStringConfig `hcl:"max"`
	Min     *IntStringConfig `hcl:"min"`
	Show    bool             `hcl:"show"`
}
