package graph

import (
	"github.com/MikhailGulkin/CleanGolangOrderApp/order/internal/presentation/config"
	"github.com/goccy/go-graphviz"
	"go.uber.org/fx"
	"os"
)

var Module = fx.Options(fx.Invoke(Start))

func Start(
	config config.AppConfig,
	diGraph fx.DotGraph,
) {
	if config.Mode != "development" {
		return
	}
	g := graphviz.New()
	graph, _ := graphviz.ParseBytes([]byte(diGraph))
	os.WriteFile("./graph.gv", []byte(diGraph), 0644)
	g.RenderFilename(graph, graphviz.PNG, "./graph.png")
	g.RenderFilename(graph, graphviz.SVG, "./graph.svg")
}
