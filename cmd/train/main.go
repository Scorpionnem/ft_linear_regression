package main

import (
	"encoding/csv"
	"fmt"
	"ft_linear_regression/internal/utils"
	"image/color"
	"log"
	"os"
	"strconv"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func GetCsv(path string) [][]string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Error opening file: "+path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Error reading file: "+path, err)
	}

	return records
}

type Data struct {
	km    float64
	price float64
}

func CsvToData(records [][]string) []Data {
	records = records[1:]

	ret := make([]Data, len(records))
	var err error
	for i, row := range records {
		ret[i].km, err = strconv.ParseFloat(row[0], 64)
		if err != nil {
			log.Fatal("Error parsing value: "+row[0], err)
		}

		ret[i].price, err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			log.Fatal("Error parsing value: "+row[1], err)
		}
	}
	return ret
}

func NormalizeData(data []Data) ([]Data, float64) {
	normalizedData := data
	var maxKm float64 = data[0].km

	for _, d := range data {
		if d.km > maxKm {
			maxKm = d.km
		}
	}

	for i, d := range data {
		normalizedData[i].km = d.km / maxKm
	}

	return normalizedData, maxKm
}

func gradientStep(data []Data, learningRate float64, theta0 float64, theta1 float64) ([]Data, float64, float64) {
	sum0 := 0.0
	sum1 := 0.0
	m := len(data)

	for i := 0; i < m; i++ {
		km := data[i].km
		price := data[i].price
		prediction := utils.EstimatePrice(km, theta0, theta1)
		error := prediction - price

		sum0 += error
		sum1 += error * km
	}

	tmpTheta0 := theta0 - learningRate*(sum0/float64(m))
	tmpTheta1 := theta1 - learningRate*(sum1/float64(m))

	theta0 = tmpTheta0
	theta1 = tmpTheta1

	return data, theta0, theta1
}

func plotResult(data []Data, theta0, theta1 float64) {
	p := plot.New()
	p.Title.Text = "Price / Km"
	p.X.Label.Text = "km"
	p.Y.Label.Text = "price"

	pts := make(plotter.XYs, len(data))
	for i := range pts {
		pts[i].X = data[i].km
		pts[i].Y = data[i].price
	}
	scatter, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}

	line := plotter.NewFunction(func(km float64) float64 {
		return theta0 + theta1*km
	})
	line.Color = color.RGBA{R: 255, A: 255}
	line.Width = vg.Points(2)

	p.Add(scatter, line)
	p.Legend.Add("data", scatter)
	p.Legend.Add("regression", line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "plot.png"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: ./train <dataset>")
		return
	}

	data := CsvToData(GetCsv(os.Args[1]))
	normalizedData, maxKm := NormalizeData(data)

	theta0 := 0.0
	theta1 := 0.0

	learningRate := 0.01
	iterations := 10000

	for range iterations {
		normalizedData, theta0, theta1 = gradientStep(normalizedData, learningRate, theta0, theta1)
	}

	plotResult(data, theta0, theta1)
	fmt.Println("theta0:", theta0, "theta1:", theta1/maxKm)
}
