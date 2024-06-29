package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var openDataAPIKey = os.Getenv("OPEN_DATA_API_KEY")

func main() {
	go update()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8000", nil)

}

func update() {
	updateTicker := time.NewTicker(8 * time.Second)
	defer updateTicker.Stop()

	for range updateTicker.C {

		var trainOccupancySum int32
		var avgTrainOccupancy int32

		var metroOccupancySum int32
		var avgMetroOccupancy int32

		var busOccupancySum int32
		var avgBusOccupancy int32

		var ferryOccupancySum int32
		var avgFerryOccupancy int32

		var cbdLightRailOccupancySum int32
		var avgCbdLightRailOccupancy int32

		trains := GetTrains()
		trainGauge.Set(float64(len(trains)))

		metros := GetMetros()
		metroGauge.Set(float64(len(metros)))

		buses := GetBuses()
		busGauge.Set(float64(len(buses)))

		ferries := GetFerries()
		ferryGauge.Set(float64(len(ferries)))

		cbdLightRail := GetLightRail()
		cbdLightRailGauge.Set(float64(len(cbdLightRail)))

		for _, v := range trains {
			trainOccupancySum += v.Occupancy
		}

		for _, v := range metros {
			metroOccupancySum += v.Occupancy
		}

		for _, v := range buses {
			busOccupancySum += v.Occupancy
		}

		for _, v := range ferries {
			ferryOccupancySum += v.Occupancy
		}

		for _, v := range cbdLightRail {
			cbdLightRailOccupancySum += v.Occupancy
		}

		if len(trains) > 0 {
			avgTrainOccupancy = trainOccupancySum / int32(len(trains))
			trainOccupancyGauge.Set(float64(avgTrainOccupancy))
		}

		if len(metros) > 0 {
			avgMetroOccupancy = metroOccupancySum / int32(len(metros))
			metroOccupancyGauge.Set(float64(avgMetroOccupancy))
		}

		if len(buses) > 0 {
			avgBusOccupancy = busOccupancySum / int32(len(buses))
			busOccupancyGauge.Set(float64(avgBusOccupancy))
		}

		if len(ferries) > 0 {
			avgFerryOccupancy = ferryOccupancySum / int32(len(ferries))
			ferryOccupancyGauge.Set(float64(avgFerryOccupancy))
		}

		if len(cbdLightRail) > 0 {
			avgCbdLightRailOccupancy = cbdLightRailOccupancySum / int32(len(cbdLightRail))
			cbdLightRailOccupancyGauge.Set(float64(avgCbdLightRailOccupancy))
		}

	}

}
