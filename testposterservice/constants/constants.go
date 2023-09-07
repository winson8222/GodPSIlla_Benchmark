package constants

import "tests/constants"

type Service struct {
	SERVICE_NAME        string
	UPSTREAM_URL        string
	LOAD_BALANCING_TYPE string
}

// information about etcd instance
const (
	ETCD_ADDRESS = "0.0.0.0:20000"
)

// var name is [ServiceName]
var ViewerService = Service{
	SERVICE_NAME:        "PosterService", //take straight from the idl file
	UPSTREAM_URL:        "localhost:" + constants.TESTPOSTERSERVICE_PORT,
	LOAD_BALANCING_TYPE: "ROUND_ROBIN",
}
