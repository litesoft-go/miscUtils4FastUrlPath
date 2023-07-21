package urlpath

import (
	"strconv"
	"testing"
	"utils"
)

func Test_GraphBuilderWithInitialPath(t *testing.T) {
	builder := NewGraphBuilder()
	if !check(t, builder, "/") {
		return
	}
	lb := builder.GetLinearBuilder()

	lb.AddPathEntry("fred")
	if !check(t, builder, "/fred") {
		return
	}

	lb.AddPathEntry("wilma")
	if !check(t, builder, "/fred/wilma") {
		return
	}

	zOr := lb.AddOr()
	zMap1 := zOr.NextOption().AddPathEntry("fw1").AddMap()
	zMap2 := zOr.NextOption().AddPathEntry("fw2").AddPathEntry("fw22").AddMap()
	zMap1.AddOption("m2-1").AddPathEntry("m1")
	zMap1.AddOption("m1-2").AddPathEntry("m2")
	zMap2.AddOption("m1-1").AddPathEntry("m1")
	zMap2.AddOption("m2-2").AddPathEntry("m2")
	if !check(t, builder, "",
		"/fred/wilma/┤fw1/¿m2-1/m1",
		"                 ¿m1-2/m2",
		"            ┤fw2/fw22/¿m1-1/m1",
		"                      ¿m2-2/m2",
		"") {
		return
	}
}

func Test_GraphBuilderWithInitialOr(t *testing.T) {
	builder := NewGraphBuilder()
	zOr := builder.GetLinearBuilder().AddOr()
	zMap1 := zOr.NextOption().AddPathEntry("fw1").AddMap()
	zMap2 := zOr.NextOption().AddPathEntry("fw2").AddPathEntry("fw22").AddMap()
	zMap1.AddOption("m2-1").AddPathEntry("m1")
	zMap1.AddOption("m1-2").AddPathEntry("m2")
	zMap2.AddOption("m1-1").AddPathEntry("m1")
	zMap2.AddOption("m2-2").AddPathEntry("m2")
	if !check(t, builder, "",
		"/┤fw1/¿m2-1/m1",
		"      ¿m1-2/m2",
		" ┤fw2/fw22/¿m1-1/m1",
		"           ¿m2-2/m2",
		"") {
		return
	}
}

func Test_GraphBuilderWithInitialMap(t *testing.T) {
	builder := NewGraphBuilder()
	zMap := builder.GetLinearBuilder().AddMap()
	zMap.AddOption("m2-1").AddPathEntry("m1")
	zMap.AddOption("m1-2").AddPathEntry("m2")
	if !check(t, builder, "",
		"/¿m2-1/m1",
		" ¿m1-2/m2",
		"") {
		return
	}
}

func check(t *testing.T, builder *GraphBuilder, expectedLines ...string) bool {
	actual := builder.String()
	expected := utils.LinesToString(expectedLines...)
	success := actual == expected
	if !success {
		t.Error(
			"\n--- expected (" + strconv.Itoa(len(expected)) + "):" +
				"\n" + expected +
				"\n--- actual (" + strconv.Itoa(len(actual)) + "):" +
				"\n" + actual + "\n")
	}
	return success
}
