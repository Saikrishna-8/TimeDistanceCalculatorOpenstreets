# TimeDistanceCalculatorOpenstreets

Run the grpc Server

go run main.go

Use the grpcurl to call the grpc Method.

grpcurl -plaintext -d '{"from_address": "4463 Maryland Ave, St. Louis, MO 63108", "to_address": "14375 Manchester Rd, Manchester, MO 63011"}' localhost:50051 distance1.DistanceCalculatorService/GetDistance
