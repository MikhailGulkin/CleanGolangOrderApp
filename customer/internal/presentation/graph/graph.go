package graph

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/pkg/env"
	"github.com/goccy/go-graphviz"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(fx.Invoke(Start))

func Start(
	diGraph fx.DotGraph,
) {
	if env.GetEnv("APP_MODE", "development") != "development" {
		return
	}
	g := graphviz.New()
	graph, _ := graphviz.ParseBytes([]byte(diGraph))
	os.WriteFile("./graph.gv", []byte(diGraph), 0644)
	g.RenderFilename(graph, graphviz.PNG, "./graph.png")
	g.RenderFilename(graph, graphviz.SVG, "./graph.svg")
}
