package hcl

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGraphPanelConfigParsing(t *testing.T) {

	expected := &Config{
		GraphPanels: []GraphPanelConfig{
			GraphPanelConfig{
				Name: "node_cpu",
				Bars: false,
				Fill: 1,
				Legend: LegendGraphPanelConfig{
					Show:         true,
					Min:          false,
					Max:          false,
					Avg:          false,
					Current:      false,
					Total:        false,
					HideEmpty:    false,
					HideZero:     false,
					AlignAsTable: false,
					RightSide:    false,
					SideWidth:    0,
				},
			},
		},
	}

	config, err := ParseConfig(testGraphPanel)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(config, expected) {
		t.Error("Graph Panel structure differed from expectation")
		fmt.Printf("%v\n", config)
		fmt.Printf("%v\n", expected)
	}

	for index := 0; index < len(config.GraphPanels); index++ {
		err = Validate(&config.GraphPanels[index])
		if err != nil {
			t.Error(err)
		}

	}
}

const testGraphPanel = `graph_panel "node_cpu" {
	bars = false
	fill = 1

	legend {
		show = true
		show_mix = false
		show_max = false
		show_average = false
		show_current = false
		show_total = false
		hide_empty = false
		hide_zero = false
		align_as_table = false
		right_side = false
		right_side_width = 0
	}
}`
