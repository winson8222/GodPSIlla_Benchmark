## Overview

This testing folder contains the implementations for test microservices

The test file sends a HTTP Request to the client (3001)
which makes a grpc call to the middleman (7999)
which makes a http request to the API Gateway (8888)
which makes a thrift request to the TestViewerService and TestPosterService

## How to Use

### System Requirements

- Golang

### Setting up the project

1. Run binary (./{module name} for MacOS, ./{module name}.exe for Windows) for the client, gateway, middleman, testviewerservice and testposterservice
2. Run the binaries in all the folders to start up all the components
3. Ensure that the etcd cluster is running
4. Ensure that port 3001, 7999, 8003, 8002, 8888 are kept clear for testing

### Running the project

1. In /tests/main_test.go, run TestRun to execute all the tests once.
2. In /tests/main_test.go, run TestBenchmarking to execute all the tests a hundred times each and compute the average time taken.
