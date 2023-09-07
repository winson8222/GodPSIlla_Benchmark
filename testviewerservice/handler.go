package main

import (
	"context"
	"math/rand"
	"tests/constants"
	"tests/myPRF"
	viewersapi "testviewerservice/kitex_gen/viewersapi"
)

// ViewerServiceImpl implements the last service interface defined in the IDL.
type ViewerServiceImpl struct{}

// Getuniqueviewernames8020 implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames8020(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	var viewersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 80; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		viewersList = append(viewersList, randomString)
	}

	for i := 0; i < 20; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		viewersList = append(viewersList, randomString)
	}

	response := &viewersapi.Response{
		Viewerslist: viewersList,
	}
	return response, nil
}

// Getuniqueviewernames2080 implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames2080(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	var viewersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 20; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		viewersList = append(viewersList, randomString)
	}

	for i := 0; i < 80; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		viewersList = append(viewersList, randomString)
	}

	response := &viewersapi.Response{
		Viewerslist: viewersList,
	}
	return response, nil
}

// Getuniqueviewernames8099920 implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames8099920(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	var viewersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 80; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		viewersList = append(viewersList, randomString)
	}

	for i := 0; i < 99920; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		viewersList = append(viewersList, randomString)
	}

	response := &viewersapi.Response{
		Viewerslist: viewersList,
	}
	return response, nil
}

// Getuniqueviewernames2099980 implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames2099980(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	var viewersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 20; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		viewersList = append(viewersList, randomString)
	}

	for i := 0; i < 99980; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		viewersList = append(viewersList, randomString)
	}

	response := &viewersapi.Response{
		Viewerslist: viewersList,
	}
	return response, nil
}

// Getuniqueviewernames80k20k implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames80k20k(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	var viewersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 80000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		viewersList = append(viewersList, randomString)
	}

	for i := 0; i < 20000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		viewersList = append(viewersList, randomString)
	}

	response := &viewersapi.Response{
		Viewerslist: viewersList,
	}
	return response, nil
}

// Getuniqueviewernames20k80k implements the ViewerServiceImpl interface.
func (s *ViewerServiceImpl) Getuniqueviewernames20k80k(ctx context.Context, req *viewersapi.Request) (resp *viewersapi.Response, err error) {
	var viewersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 20000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		viewersList = append(viewersList, randomString)
	}

	for i := 0; i < 80000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		viewersList = append(viewersList, randomString)
	}

	response := &viewersapi.Response{
		Viewerslist: viewersList,
	}
	return response, nil
}
