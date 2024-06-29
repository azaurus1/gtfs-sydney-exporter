package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/azaurus1/gtfs-sydney-exporter/cmd/transit_realtime"
	"google.golang.org/protobuf/proto"
)

type TrainMetric struct {
	Occupancy int32
}

type MetroMetric struct {
	Occupancy int32
}

type BusMetric struct {
	Occupancy int32
}

type FerryMetric struct {
	Occupancy int32
}

type LightRailMetric struct {
	Occupancy int32
}

func GetTrains() []TrainMetric {

	var TrainMetricArr []TrainMetric

	req, err := http.NewRequest("GET", "https://api.transport.nsw.gov.au/v2/gtfs/vehiclepos/sydneytrains", nil)
	if err != nil {
		log.Println("error building request: ", err)
	}

	authStr := fmt.Sprintf("apikey %s", openDataAPIKey)

	req.Header.Set("Authorization", authStr)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request: ", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response into a FeedMessage
	feedMessage := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(body, feedMessage)
	if err != nil {
		log.Println(body)
		log.Fatalf("Failed to unmarshal feed message: %v", err)
	}

	for _, entity := range feedMessage.GetEntity() {
		trainMetric := TrainMetric{}
		if entity.Vehicle != nil {
			occupancyStr := entity.Vehicle.GetOccupancyStatus().String()
			trainMetric.Occupancy = (transit_realtime.TripUpdate_StopTimeUpdate_OccupancyStatus_value[occupancyStr])
		}
		TrainMetricArr = append(TrainMetricArr, trainMetric)
	}

	return TrainMetricArr
}

func GetMetros() []MetroMetric {

	var MetroMetricArr []MetroMetric

	req, err := http.NewRequest("GET", "https://api.transport.nsw.gov.au/v2/gtfs/vehiclepos/metro", nil)
	if err != nil {
		log.Println("error building request: ", err)
	}

	authStr := fmt.Sprintf("apikey %s", openDataAPIKey)

	req.Header.Set("Authorization", authStr)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request: ", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response into a FeedMessage
	feedMessage := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(body, feedMessage)
	if err != nil {
		log.Println(body)
		log.Fatalf("Failed to unmarshal feed message: %v", err)
	}

	for _, entity := range feedMessage.GetEntity() {
		metroMetric := MetroMetric{}
		if entity.Vehicle != nil {
			occupancyStr := entity.Vehicle.GetOccupancyStatus().String()
			metroMetric.Occupancy = (transit_realtime.TripUpdate_StopTimeUpdate_OccupancyStatus_value[occupancyStr])
		}
		MetroMetricArr = append(MetroMetricArr, metroMetric)
	}

	return MetroMetricArr
}

func GetBuses() []BusMetric {

	var BusMetricArr []BusMetric

	req, err := http.NewRequest("GET", "https://api.transport.nsw.gov.au/v1/gtfs/vehiclepos/buses", nil)
	if err != nil {
		log.Println("error building request: ", err)
	}

	authStr := fmt.Sprintf("apikey %s", openDataAPIKey)

	req.Header.Set("Authorization", authStr)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request: ", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response into a FeedMessage
	feedMessage := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(body, feedMessage)
	if err != nil {
		log.Println(body)
		log.Fatalf("Failed to unmarshal feed message: %v", err)
	}

	for _, entity := range feedMessage.GetEntity() {
		busMetric := BusMetric{}
		if entity.Vehicle != nil {
			occupancyStr := entity.Vehicle.GetOccupancyStatus().String()
			busMetric.Occupancy = (transit_realtime.TripUpdate_StopTimeUpdate_OccupancyStatus_value[occupancyStr])
		}
		BusMetricArr = append(BusMetricArr, busMetric)
	}

	return BusMetricArr
}

func GetFerries() []FerryMetric {

	var FerryMetricArr []FerryMetric

	req, err := http.NewRequest("GET", "https://api.transport.nsw.gov.au/v1/gtfs/vehiclepos/ferries/sydneyferries", nil)
	if err != nil {
		log.Println("error building request: ", err)
	}

	authStr := fmt.Sprintf("apikey %s", openDataAPIKey)

	req.Header.Set("Authorization", authStr)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request: ", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response into a FeedMessage
	feedMessage := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(body, feedMessage)
	if err != nil {
		log.Println(body)
		log.Fatalf("Failed to unmarshal feed message: %v", err)
	}

	for _, entity := range feedMessage.GetEntity() {
		ferryMetric := FerryMetric{}
		if entity.Vehicle != nil {
			occupancyStr := entity.Vehicle.GetOccupancyStatus().String()
			ferryMetric.Occupancy = (transit_realtime.TripUpdate_StopTimeUpdate_OccupancyStatus_value[occupancyStr])
		}
		FerryMetricArr = append(FerryMetricArr, ferryMetric)
	}

	return FerryMetricArr
}

func GetLightRail() []LightRailMetric {

	var LightRailMetricArr []LightRailMetric

	req, err := http.NewRequest("GET", "https://api.transport.nsw.gov.au/v1/gtfs/vehiclepos/lightrail/cbdandsoutheast", nil)
	if err != nil {
		log.Println("error building request: ", err)
	}

	authStr := fmt.Sprintf("apikey %s", openDataAPIKey)

	req.Header.Set("Authorization", authStr)

	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request: ", err)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	// Unmarshal the response into a FeedMessage
	feedMessage := &transit_realtime.FeedMessage{}
	err = proto.Unmarshal(body, feedMessage)
	if err != nil {
		log.Println(body)
		log.Fatalf("Failed to unmarshal feed message: %v", err)
	}

	for _, entity := range feedMessage.GetEntity() {
		lightRailMetric := LightRailMetric{}
		if entity.Vehicle != nil {
			occupancyStr := entity.Vehicle.GetOccupancyStatus().String()
			lightRailMetric.Occupancy = (transit_realtime.TripUpdate_StopTimeUpdate_OccupancyStatus_value[occupancyStr])
		}
		LightRailMetricArr = append(LightRailMetricArr, lightRailMetric)
	}

	return LightRailMetricArr
}
