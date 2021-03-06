package main

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/guptarohit/asciigraph"
)

func createDevice(id string) DeviceStatsModel {
	devicesMap[id] = DeviceStatsModel{
		Hostname:  id,
		Timestamp: time.Now().Unix(),
		Graph:     make([]float64, 30),
	}
	return devicesMap[id]
}

func renderTemplate(w http.ResponseWriter, name string, data any) {
	err := templates.ExecuteTemplate(w, name, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func updateGraph(diff int64, graph []float64) []float64 {
	if diff < 30 {
		graph = append(graph, make([]float64, diff)...)
		graph = graph[(len(graph) - 30):]
	} else {
		graph = make([]float64, 30)
	}
	return graph
}

func plotGraph(graph []float64) string {
	var g string
	g = asciigraph.Plot(graph, asciigraph.Precision(0), asciigraph.Height(2))
	reg := regexp.MustCompile("[0-9]")
	g = reg.ReplaceAllString(g, "")
	g = strings.ReplaceAll(g, "  ┼", "")
	g = strings.ReplaceAll(g, "  ┤", "")
	return g
}
