package hcl

import (
	"strconv"

	grafana "github.com/grafana-tools/sdk"
)

type GraphPanelConfig struct {
	Name            string                           `hcl:",key"`
	AliasColors     interface{}                      `hcl:"alias_colors"`
	Bars            bool                             `hcl:"bars"`
	Fill            int                              `hcl:"fill"`
	Legend          LegendGraphPanelConfig           `hcl:"legend"`
	LeftYAxisLabel  string                           `hcl:"left_y_axis_label"`
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
	Decimals        int                              `hcl:"decimals"`
}

type LegendGraphPanelConfig struct {
	AlignAsTable bool `hcl:"align_as_table"`
	Avg          bool `hcl:"show_average"`
	Current      bool `hcl:"show_current"`
	HideEmpty    bool `hcl:"hide_empty"`
	HideZero     bool `hcl:"hide_zero"`
	Max          bool `hcl:"show_max"`
	Min          bool `hcl:"show_min"`
	RightSide    bool `hcl:"right_side"`
	Show         bool `hcl:"show"`
	Total        bool `hcl:"show_total"`
	Values       bool `hcl:"values"`
	SideWidth    int  `hcl:"right_side_width"`
}

type SeriesOverrideGraphPanelConfig struct {
	Alias         string           `hcl:"alias"`
	Bars          bool             `hcl:"bars"`
	Color         string           `hcl:"color"`
	Fill          int              `hcl:"fill"`
	FillBelowTo   string           `hcl:"fill_below_to"`
	Legend        bool             `hcl:"legend"`
	Lines         bool             `hcl:"lines"`
	Stack         BoolStringConfig `hcl:"stack"`
	Transform     string           `hcl:"transform"`
	YAxis         int              `hcl:"y_axis"`
	ZIndex        int              `hcl:"z_index"`
	NullPointMode string           `hcl:"null_point_mode"`
}

type AxisGraphPanelConfig struct {
	Format  string          `hcl:"format"`
	LogBase int             `hcl:"logBase"`
	Max     IntStringConfig `hcl:"max"`
	Min     IntStringConfig `hcl:"min"`
	Show    bool            `hcl:"show"`
}

func (g *GraphPanelConfig) ValidateHCL() error {
	return nil
}

func (g *GraphPanelConfig) GenerateJSON() ([]byte, error) {
	gp := grafana.NewGraph(g.Name)

	gp.AliasColors = g.AliasColors
	gp.Bars = g.Bars
	gp.Fill = g.Fill
	gp.Legend.AlignAsTable = g.Legend.AlignAsTable
	gp.Legend.Avg = g.Legend.Avg
	gp.Legend.Current = g.Legend.Current
	gp.Legend.HideEmpty = g.Legend.HideEmpty
	gp.Legend.HideZero = g.Legend.HideZero
	gp.Legend.Max = g.Legend.Max
	gp.Legend.Min = g.Legend.Min
	gp.Legend.RightSide = g.Legend.RightSide
	gp.Legend.Show = g.Legend.Show
	gp.Legend.Total = g.Legend.Total
	gp.Legend.Values = g.Legend.Values
	sideWidth := uint(g.Legend.SideWidth)
	gp.Legend.SideWidth = &sideWidth
	gp.LeftYAxisLabel = &g.LeftYAxisLabel
	gp.RightYAxisLabel = &g.RightYAxisLabel
	gp.Lines = g.Lines
	gp.Linewidth = g.Linewidth
	// gp.NullPointMode = g.NullPointMode
	gp.Percentage = g.Percentage
	gp.Pointradius = g.Pointradius
	gp.Points = g.Points

	seriesOverrides := []grafana.SeriesOverride{}

	for index := 0; index < len(g.SeriesOverrides); index++ {
		seriesOveride := grafana.SeriesOverride{
			Alias:       g.SeriesOverrides[index].Alias,
			Bars:        &g.SeriesOverrides[index].Bars,
			Color:       &g.SeriesOverrides[index].Color,
			Fill:        &g.SeriesOverrides[index].Fill,
			FillBelowTo: &g.SeriesOverrides[index].FillBelowTo,
			Legend:      &g.SeriesOverrides[index].Legend,
			Lines:       &g.SeriesOverrides[index].Lines,
			Stack: &grafana.BoolString{
				Flag:  g.SeriesOverrides[index].Stack.Flag,
				Value: g.SeriesOverrides[index].Stack.Value,
			},
			Transform:     &g.SeriesOverrides[index].Transform,
			YAxis:         &g.SeriesOverrides[index].YAxis,
			ZIndex:        &g.SeriesOverrides[index].ZIndex,
			NullPointMode: &g.SeriesOverrides[index].NullPointMode,
		}

		seriesOverrides = append(seriesOverrides, seriesOveride)
	}

	gp.Stack = g.Stack
	gp.SteppedLine = g.SteppedLine

	targets := []grafana.Target{}

	for index := 0; index < len(g.Targets); index++ {
		target := grafana.Target{
			Datasource:     g.Targets[index].Datasource,
			Expr:           g.Targets[index].Expr,
			IntervalFactor: g.Targets[index].IntervalFactor,
			Interval:       g.Targets[index].Interval,
			Step:           g.Targets[index].Step,
			LegendFormat:   g.Targets[index].LegendFormat,
			DsType:         &g.Targets[index].DsType,
		}

		var targetMetrics []struct {
			ID    string `json:"id"`
			Field string `json:"field"`
			Type  string `json:"type"`
		}

		for mIndex := 0; mIndex < len(g.Targets[index].Metrics); mIndex++ {
			metric := struct {
				ID    string `json:"id"`
				Field string `json:"field"`
				Type  string `json:"type"`
			}{
				strconv.Itoa(mIndex),
				g.Targets[index].Metrics[mIndex].Field,
				g.Targets[index].Metrics[mIndex].Type,
			}

			targetMetrics = append(targetMetrics, metric)
		}

		target.Metrics = targetMetrics
		target.Query = g.Targets[index].Query
		target.TimeField = g.Targets[index].TimeField

		var targetBucketAggs []struct {
			ID       string `json:"id"`
			Field    string `json:"field"`
			Type     string `json:"type"`
			Settings struct {
				Interval    string `json:"interval"`
				MinDocCount int    `json:"min_doc_count"`
			} `json:"settings"`
		}

		for bIndex := 0; bIndex < len(g.Targets[index].BucketAggs); bIndex++ {
			bucketAgg := struct {
				ID       string `json:"id"`
				Field    string `json:"field"`
				Type     string `json:"type"`
				Settings struct {
					Interval    string `json:"interval"`
					MinDocCount int    `json:"min_doc_count"`
				} `json:"settings"`
			}{
				strconv.Itoa(bIndex),
				g.Targets[index].BucketAggs[bIndex].Field,
				g.Targets[index].BucketAggs[bIndex].Type,
				struct {
					Interval    string `json:"interval"`
					MinDocCount int    `json:"min_doc_count"`
				}{
					g.Targets[index].BucketAggs[bIndex].Settings.Interval,
					g.Targets[index].BucketAggs[bIndex].Settings.MinDocCount,
				},
			}

			targetBucketAggs = append(targetBucketAggs, bucketAgg)
		}

		target.Target = g.Targets[index].Target
		targets = append(targets, target)

	}

	return gp.MarshalJSON()
}
