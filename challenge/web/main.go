package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type statistics struct {
	numbers []float64
	mean    float64
	median  float64
}

const form = `<html><body><form action="/" method="POST">
<h1>Statistics</h1>
<h5>Compute base statistics for a given list of numbers</h5>
<label for="numbers">Numbers (comma or space-separated):</label><br>
<input type="text" name="numbers" size="30"><br />
<input type="submit" value="Calculate">
</form></html></body>`

const e = `<p class="e">%s</p>`

//var pageTop = ""
//var pageBottom = ""

// Define a root handler for requests to function homePage, and start the webserver combined with e-handling
func main() {
	http.HandleFunc("/", homePage)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Error starting the server")
	}
}

// Write an HTML header, parse the form, write form to writer and make request for numbers
func homePage(writer http.ResponseWriter, request *http.Request) {
	// write your code here
	writer.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case http.MethodGet:
		_, err := io.WriteString(writer, form)
		if err != nil {
			log.Println("Error: ", err)
		}
	case http.MethodPost:
		numbers, err, ok := processRequest(request)
		if ok {
			stat := getStats(numbers)
			statStr := formatStats(stat)
			_, err := io.WriteString(writer, statStr)
			if err != nil {
				return
			}
		} else {
			_, err := fmt.Fprintf(writer, e, err)
			if err != nil {
				return
			}
		}
	}
}

// Capture the numbers from the request, and format the data and check for errors
func processRequest(request *http.Request) ([]float64, string, bool) {
	// write your code here
	numbers := request.FormValue("numbers")
	strSlice := strings.Split(numbers, " ")
	var numSlice []float64
	for _, val := range strSlice {
		num, err := strconv.ParseFloat(val, 64)
		if err == nil {
			numSlice = append(numSlice, num)
		}
	}
	strings.Split(numbers, ", ")
	return numSlice, "", true
}

// sort the values to get mean and median
func getStats(numbers []float64) (stats statistics) {
	// write your code here
	return statistics{numbers: numbers, mean: sum(numbers) / float64(len(numbers)), median: median(numbers)}
}

// separate function to calculate the sum for mean
func sum(numbers []float64) (total float64) {
	// write your code here
	s := .0
	for _, val := range numbers {
		s += val
	}
	return s
}

// separate function to calculate the median
func median(numbers []float64) float64 {
	// write your code here
	sort.Float64s(numbers)
	return numbers[len(numbers)/2]
}

func formatStats(stats statistics) string {
	return fmt.Sprintf(`<table border="1">
<tr><th colspan="2">Results</th></tr>
<tr><td>Numbers</td><td>%v</td></tr>
<tr><td>Count</td><td>%d</td></tr>
<tr><td>Mean</td><td>%f</td></tr>
<tr><td>Median</td><td>%f</td></tr>
</table>`, stats.numbers, len(stats.numbers), stats.mean, stats.median)
}
