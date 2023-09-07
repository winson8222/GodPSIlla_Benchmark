package main

import (
	"context"
	"math/rand"
	postersapi "testposterservice/kitex_gen/postersapi"
	"tests/constants"
	"tests/myPRF"
)

// PosterServiceImpl implements the last service interface defined in the IDL.
type PosterServiceImpl struct{}

// Getuniqueusernames8020 implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames8020(ctx context.Context, req *postersapi.Request) (resp *postersapi.Response, err error) {
	var postersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 80; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		postersList = append(postersList, randomString)
	}

	for i := 0; i < 20; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		postersList = append(postersList, randomString)
	}

	response := &postersapi.Response{
		Posterslist: postersList,
	}
	return response, nil
}

// Getuniqueusernames2080 implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames2080(ctx context.Context, req *postersapi.Request) (resp *postersapi.Response, err error) {
	var postersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 20; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		postersList = append(postersList, randomString)
	}

	for i := 0; i < 80; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		postersList = append(postersList, randomString)
	}

	response := &postersapi.Response{
		Posterslist: postersList,
	}
	return response, nil
}

// Getuniqueusernames8099920 implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames8099920(ctx context.Context, req *postersapi.Request) (resp *postersapi.Response, err error) {
	var postersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 80; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		postersList = append(postersList, randomString)
	}

	for i := 0; i < 99920; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		postersList = append(postersList, randomString)
	}

	response := &postersapi.Response{
		Posterslist: postersList,
	}
	return response, nil
}

// Getuniqueusernames2099980 implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames2099980(ctx context.Context, req *postersapi.Request) (resp *postersapi.Response, err error) {
	var postersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 20; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		postersList = append(postersList, randomString)
	}

	for i := 0; i < 99980; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		postersList = append(postersList, randomString)
	}

	response := &postersapi.Response{
		Posterslist: postersList,
	}
	return response, nil
}

// Getuniqueusernames80k20k implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames80k20k(ctx context.Context, req *postersapi.Request) (resp *postersapi.Response, err error) {
	var postersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 80000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		postersList = append(postersList, randomString)
	}

	for i := 0; i < 20000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		postersList = append(postersList, randomString)
	}

	response := &postersapi.Response{
		Posterslist: postersList,
	}
	return response, nil
}

// Getuniqueusernames20k80k implements the PosterServiceImpl interface.
func (s *PosterServiceImpl) Getuniqueusernames20k80k(ctx context.Context, req *postersapi.Request) (resp *postersapi.Response, err error) {
	var postersList []string

	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	seedValue2 := int64(constants.Prf2Seed) // Choose a constant seed value
	source2 := rand.NewSource(seedValue2)
	rng2 := rand.New(source2)

	for i := 0; i < 20000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		postersList = append(postersList, randomString)
	}

	for i := 0; i < 80000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng2)
		postersList = append(postersList, randomString)
	}

	response := &postersapi.Response{
		Posterslist: postersList,
	}
	return response, nil
}
