package constants

const (
	LetterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Prf1Seed    = 12345 // seed to create prf1's rng object
	Prf2Seed    = 67890 // seed to create prf2's rng object
)

const (
	GRPCCLIENT_HTTPSERVER_PORT = "3001"
	GRPCCLIENT_HTTPSERVER_ADDR = "http://localhost:3001"
)

const (
	GRPC_MIDDLEMAN_PORT = "7999" //To be customised by the user
	GRPC_MIDDLEMAN_ADDR = "localhost:7999"
)

const (
	APIGATEWAY_URL = "localhost:8888"
)

const (
	TESTPOSTERSERVICE_PORT = "8002"
	TESTVIEWERSERVICE_PORT = "8003"
)
