package main

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"testing"
	"tests/constants"
	"time"
)

func TestRun(t *testing.T) {
	t.Run("TestSmallSmall20", TestSmallSmall20)
	t.Run("TestSmallSmall80", TestSmallSmall80)
	t.Run("TestSmallLarge20", TestSmallLarge20)
	t.Run("TestSmallLarge80", TestSmallLarge80)
	t.Run("TestLargeSmall20", TestLargeSmall20)
	t.Run("TestLargeSmall80", TestLargeSmall80)
	t.Run("TestLargeLarge20000", TestLargeLarge20000)
	t.Run("TestLargeLarge80000", TestLargeLarge80000)
}

func TestBenchmarking(t *testing.T) {
	t.Run("HundredTestSmallSmall20", HundredTestSmallSmall20)
	t.Run("HundredTestSmallSmall80", HundredTestSmallSmall80)
	t.Run("HundredTestSmallLarge20", HundredTestSmallLarge20)
	t.Run("HundredTestSmallLarge80", HundredTestSmallLarge80)
	t.Run("HundredTestLargeSmall20", HundredTestLargeSmall20)
	t.Run("HundredTestLargeSmall80", HundredTestLargeSmall80)
	t.Run("HundredTestLargeLarge20000", HundredTestLargeLarge20000)
	t.Run("HundredTestLargeLarge80000", HundredTestLargeLarge80000)
}

// Makes HTTP Request to GRPC Client / HTTP Server that is running at port 3001
func TestSmallSmall20(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames2080 PosterService getuniqueusernames2080"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestSmallSmall20 %s", elapsedTime)
}

func TestSmallSmall80(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames8020 PosterService getuniqueusernames8020"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestSmallSmall80 %s", elapsedTime)
}

func TestSmallLarge20(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames2099980 PosterService getuniqueusernames2099980"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestSmallLarge20 %s", elapsedTime)
}

func TestSmallLarge80(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames8099920 PosterService getuniqueusernames8099920"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestSmallLarge80 %s", elapsedTime)
}

func TestLargeSmall20(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames2080 PosterService getuniqueusernames2080"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestLargeSmall20 %s", elapsedTime)
}

func TestLargeSmall80(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames8020 PosterService getuniqueusernames8020"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestLargeSmall80 %s", elapsedTime)
}

func TestLargeLarge20000(t *testing.T) {
	startTime := time.Now()

	// use 100.csv as file input csvfile
	// use "PosterService getuniqueusernames ViewerService getuniqueviewernames" as svcinfo
	// use constants.GRPC_MIDDLEMAN_ADDR as upstreamurl

	svcinfo := "ViewerService getuniqueviewernames20k80k PosterService getuniqueusernames20k80k"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestLargeLarge20000 %s", elapsedTime)
}

func TestLargeLarge80000(t *testing.T) {
	startTime := time.Now()

	svcinfo := "ViewerService getuniqueviewernames80k20k PosterService getuniqueusernames80k20k"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	//parses the data into a multipart form and makes the http request
	resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response
	t.Log("Response Status:", resp.Status)
	respbody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	t.Logf("Intersection Size: %s\n", respbody)

	endTime := time.Now()

	elapsedTime := endTime.Sub(startTime)

	t.Logf("TestLargeLarge80000 %s", elapsedTime)

}

func makeHTTPRequest(svcinfo string, upstreamurl string, inputfilename string) (*http.Response, error) {
	// Create a buffer to store the multipart form data
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	// Add form fields (you can add as many fields as needed)
	writer.WriteField("svcinfo", svcinfo)
	writer.WriteField("upstreamurl", upstreamurl)

	// Add a file field
	file, err := os.Open(inputfilename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Create a form field for the CSV file
	csvField, err := writer.CreateFormFile("csvfile", inputfilename)
	if err != nil {
		fmt.Println("Error creating form file:", err)
		return nil, err
	}

	// Copy the contents of the CSV file into the form field
	_, err = io.Copy(csvField, file)
	if err != nil {
		fmt.Println("Error copying CSV file content:", err)
		return nil, err
	}

	// Don't forget to close the multipart writer
	writer.Close()

	// Create the HTTP request
	req, err := http.NewRequest("POST", constants.GRPCCLIENT_HTTPSERVER_ADDR+"/PSI", &body)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return nil, err
	}

	// Set the content type for the request
	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)

	return resp, err
}

// Benchmarking section
func HundredTestSmallSmall20(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames2080 PosterService getuniqueusernames2080"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestSmallSmall20: %s", averageTime)
}
func HundredTestSmallSmall80(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames8020 PosterService getuniqueusernames8020"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestSmallSmall80: %s", averageTime)
}
func HundredTestSmallLarge20(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames2099980 PosterService getuniqueusernames2099980"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestSmallLarge20: %s", averageTime)
}
func HundredTestSmallLarge80(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames8099920 PosterService getuniqueusernames8099920"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestSmallLarge80: %s", averageTime)
}
func HundredTestLargeSmall20(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames2080 PosterService getuniqueusernames2080"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestLargeSmall20: %s", averageTime)
}
func HundredTestLargeSmall80(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames8020 PosterService getuniqueusernames8020"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestLargeSmall80: %s", averageTime)
}
func HundredTestLargeLarge20000(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames20k80k PosterService getuniqueusernames20k80k"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestLargeLarge20000: %s", averageTime)
}
func HundredTestLargeLarge80000(t *testing.T) {
	// Number of times to run the entire test function
	numRuns := 100

	// Variable to store total elapsed time for all runs
	totalElapsedTime := time.Duration(0)

	svcinfo := "ViewerService getuniqueviewernames80k20k PosterService getuniqueusernames80k20k"
	upstreamurl := constants.GRPC_MIDDLEMAN_ADDR
	inputfilename := "100000.csv"

	for i := 0; i < numRuns; i++ {
		startTime := time.Now()

		//parses the data into a multipart form and makes the http request
		resp, err := makeHTTPRequest(svcinfo, upstreamurl, inputfilename)
		if err != nil {
			fmt.Println("Error making request:", err)
			return
		}
		defer resp.Body.Close()

		endTime := time.Now()

		// Calculate the elapsed time for this run
		elapsedTime := endTime.Sub(startTime)

		// Add the elapsed time to the total
		totalElapsedTime += elapsedTime
	}

	// Calculate the average time
	averageTime := totalElapsedTime / time.Duration(numRuns)

	// Print the average time
	t.Logf("Average Time Taken for a Single Run of TestLargeLarge80000: %s", averageTime)
}
