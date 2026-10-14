package main

import ("fmt"
		"os"
		"log"
		"strconv"
		"encoding/csv")

func GetCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
        log.Fatal("Error opening file: " + path, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
    	log.Fatal("Error reading file: " + path, err)
    }

    return records
}

type Data struct {
	km		float64
	price	float64
}

func CsvToData(records [][]string) []Data {
	records = records[1:]

	ret := make([]Data, len(records))
	var err error
	for i, row := range records {
		ret[i].km, err = strconv.ParseFloat(row[0], 64)
		if (err != nil) {
			log.Fatal("Error parsing value: " + row[1], err)
		}

		ret[i].price, err = strconv.ParseFloat(row[1], 64)
		if (err != nil) {
			log.Fatal("Error parsing value: " + row[1], err)
		}
	}
	return ret
}

func NormalizeData(data []Data) ([]Data, float64) {
	normalizedData := data
	var maxKm float64 = data[0].km

	for _, d := range data {
		if (d.km > maxKm) {
			maxKm = d.km
		}
	}

	for i, d := range data {
		normalizedData[i].km = d.km / maxKm
	}

	return normalizedData, maxKm
}

func estimatePrice(km float64, theta0 float64, theta1 float64) float64 {
	return (theta0 + (theta1 * km));
}

func gradientStep(data []Data, learningRate float64, theta0 float64, theta1 float64) ([]Data, float64, float64) {
	sum0 := 0.0;
	sum1 := 0.0;
	m := len(data);

	for i := 0; i < m; i++ {
		km := data[i].km;
		price := data[i].price;
		prediction := estimatePrice(km, theta0, theta1);
		error := prediction - price;

		sum0 += error;
		sum1 += error * km;
	}

	tmpTheta0 := theta0 - learningRate * (sum0 / float64(m));
	tmpTheta1 := theta1 - learningRate * (sum1 / float64(m));

	theta0 = tmpTheta0;
	theta1 = tmpTheta1;

	return data, theta0, theta1
}

func main() {
	if (len(os.Args) != 2) {
		fmt.Println("Usage: ./train <dataset>")
		return
	}

	data := CsvToData(GetCsv(os.Args[1]))
	normalizedData, maxKm := NormalizeData(data)

   	theta0 := 0.0;
	theta1 := 0.0;

	learningRate := 0.01;
	iterations := 10000;

	for i := 0; i < iterations; i++ {
		normalizedData, theta0, theta1 = gradientStep(normalizedData, learningRate, theta0, theta1);
	}

	fmt.Println("theta0:", theta0, "theta1:", theta1 / maxKm)
}
