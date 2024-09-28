package main

import (
	"context"
	pb "distanceCalculator/distanceCalculator/distance1"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.DistanceCalculatorServiceServer
}

func (s *server) GetDistance(ctx context.Context, req *pb.DistanceRequest) (*pb.DistanceResponse, error) {
	fromAddress := req.GetFromAddress()
	toAddress := req.GetToAddress()

	distance, err := calculateDistance(fromAddress, toAddress)
	if err != nil {
		return nil, err
	}

	return &pb.DistanceResponse{Distance: distance}, nil
}

type NominatimResponse []struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type OSRMResponse struct {
	Routes []struct {
		Distance float64 `json:"distance"` // Distance in meters
		Time     float64 `json:"duration"` // Time in seconds
	} `json:"routes"`
}

func main() {
	// fromAddress := "4463 Maryland Ave, St. Louis, MO 63108"
	// toAddress := "14375 Manchester Rd, Manchester, MO 63011"

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server listening on port", lis.Addr().String())

	s := grpc.NewServer()
	pb.RegisterDistanceCalculatorServiceServer(s, &server{})

	// Register reflection service on gRPC server.
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// Calculate the distance between the two addresses

	// distance, err := calculateDistance(fromAddress, toAddress)
	// if err != nil {
	// 	fmt.Println("Error calculating distance:", err)
	// 	return
	// }

	// fmt.Printf("Distance: %.2f Miles\n", distance)
}

func calculateDistance(fromAddress, toAddress string) (float64, error) {
	fromCoords, err := geocodeAddress(fromAddress)
	if err != nil {
		return 0, err
	}

	toCoords, err := geocodeAddress(toAddress)
	if err != nil {
		return 0, err
	}

	//return fromCoords[0] + toCoords[0], nil
	//return haversineDistance(fromCoords[0], fromCoords[1], toCoords[0], toCoords[1]), nil
	return getOSRMRouteDistance(fromCoords, toCoords)
}

func geocodeAddress(address string) ([]float64, error) {

	encodedAddress := url.QueryEscape(address)
	urlString := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1", encodedAddress)

	fmt.Println("urlString ", urlString)

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return nil, err
	}

	// Set User-Agent header
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; AcmeInc/1.0)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: status code %d", resp.StatusCode)
	}

	fmt.Println("resp ", resp)
	if err != nil {
		return nil, err
	}
	var data NominatimResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("address not found")
	}

	lat, err := strconv.ParseFloat(data[0].Lat, 64)
	if err != nil {
		return nil, err
	}

	lon, err := strconv.ParseFloat(data[0].Lon, 64)
	if err != nil {
		return nil, err
	}

	fmt.Println("lat ", lat)
	fmt.Println("lon ", lon)

	return []float64{lat, lon}, nil
}

func getOSRMRouteDistance(fromCoords, toCoords []float64) (float64, error) {
	osrmServer := "http://router.project-osrm.org" // Replace with the actual OSRM server URL
	urlString := fmt.Sprintf("%s/route/v1/driving/%f,%f;%f,%f?overview=false",
		osrmServer, fromCoords[1], fromCoords[0], toCoords[1], toCoords[0])

	fmt.Println("urlString ", urlString)

	// Create a new HTTP client
	client := &http.Client{}

	// Create a new request
	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return 0, err
	}

	// Set the User-Agent header
	req.Header.Set("User-Agent", "My-Go-Distance-Calculator/1.0") // Customize this

	// Make the request using the client
	resp, err := doRequestWithRetry(client, req, 3)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data OSRMResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return 0, err
	}

	if len(data.Routes) == 0 {
		return 0, fmt.Errorf("no route found")
	}

	distanceInMeters := data.Routes[0].Distance
	distanceInMiles := distanceInMeters / 1609.34
	timeInMinutes := data.Routes[0].Time / 60

	fmt.Println("timeInMinutes ", timeInMinutes)

	return distanceInMiles, nil
}

func doRequestWithRetry(client *http.Client, req *http.Request, retries int) (*http.Response, error) {
	var resp *http.Response
	var err error

	for i := 0; i < retries; i++ {
		resp, err = client.Do(req)
		if err == nil {
			return resp, nil
		}
		time.Sleep(2 * time.Second) // Wait before retrying
	}

	return nil, err
}
