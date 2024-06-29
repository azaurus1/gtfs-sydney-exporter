package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var trainGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "train_count",
	Help: "Count of trains currently active",
})

var busGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "bus_count",
	Help: "Count of buses currently active",
})

var ferryGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "ferry_count",
	Help: "Count of ferries currently active",
})

var metroGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "metro_count",
	Help: "Count of metro currently active",
})

var cbdLightRailGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "cbd_light_rail_count",
	Help: "Count of light rail (CBD and South East) currently active",
})

var trainOccupancyGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "train_occupancy_count",
	Help: "Gauge of Train Occupancy",
})

var metroOccupancyGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "metro_occupancy_count",
	Help: "Gauge of Metro Occupancy",
})

var busOccupancyGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "bus_occupancy_count",
	Help: "Gauge of Bus Occupancy",
})

var ferryOccupancyGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "ferry_occupancy_count",
	Help: "Gauge of Ferry Occupancy",
})

var cbdLightRailOccupancyGauge = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "cbd_light_rail_occupancy_count",
	Help: "Gauge of Light Rail (CBD + South East) Occupancy",
})
