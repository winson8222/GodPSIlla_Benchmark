package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"tests/constants"
	"tests/myPRF"
)

func RandStringBytesMask(n int, rng *rand.Rand) string {
	b := make([]byte, n)
	for i := 0; i < n; i++ {
		idx := rng.Intn(len(constants.LetterBytes))
		b[i] = constants.LetterBytes[idx]
	}
	return string(b)
}

// This is where all the methods
func main() {
	Create100csv()
	Create100000csv()
}

func Create100csv() {
	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	// Create or open the CSV file for writing
	file, err := os.Create("100.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 100; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		// Write the randomString followed by a comma (CSV format)
		_, err := file.WriteString(randomString + ",")
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return
		}
	}

	fmt.Println("CSV file '100.csv' created successfully.")
}

func Create100000csv() {
	seedValue1 := int64(constants.Prf1Seed) // Choose a constant seed value
	source1 := rand.NewSource(seedValue1)
	rng1 := rand.New(source1)

	// Create or open the CSV file for writing
	file, err := os.Create("100000.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 0; i < 100000; i++ {
		randomString := myPRF.RandStringBytesMask(10, rng1)
		// Write the randomString followed by a comma (CSV format)
		_, err := file.WriteString(randomString + ",")
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return
		}
	}

	fmt.Println("CSV file '100000.csv' created successfully.")

}
