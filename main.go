package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/gocarina/gocsv"
)

type cityValue struct {
	City  string `csv:"city"`
	Value string `csv:"value"`
}

// ListCityValue ...
type ListCityValue struct {
	data []cityValue
}

// Exported ...
var (
	URL = "http://ti.saude.rs.gov.br/covid19/"
	// REGEX ...
	rgxGraphCities = regexp.MustCompile(`GraphMunicipio.*\n.*,`)
	rgxGetJSON     = regexp.MustCompile(`\{.*\}`)
	rgxLables      = regexp.MustCompile(`\{labels:.*\],`)
	rgxValues      = regexp.MustCompile(`\[[\d\,]+\]`)
	rgxCities      = regexp.MustCompile(`\[[\wÀ-ú\"\s\,\]\']+,d`)
)

func getCityValue(body string) ListCityValue {
	sliceGraph := rgxGraphCities.FindString(body)
	getJSON := rgxGetJSON.FindString(sliceGraph)
	getJSON = rgxLables.FindString(getJSON)
	getValues := rgxValues.FindString(getJSON)
	getStates := rgxCities.FindString(getJSON)

	splitVals := strings.Split(getValues, ",")
	splitCity := strings.Split(getStates, ",")
	if splitCity[len(splitCity)-1] == "d" {
		splitCity = splitCity[:len(splitCity)-1]
	}
	if strings.Contains(splitVals[len(splitVals)-1], "]") {
		splitVals[len(splitVals)-1] = strings.Replace(splitVals[len(splitVals)-1], "]", "", -1)
	}
	if strings.Contains(splitCity[len(splitCity)-1], "]") {
		splitCity[len(splitCity)-1] = strings.Replace(splitCity[len(splitCity)-1], "]", "", -1)
	}
	if strings.Contains(splitVals[0], "[") {
		splitVals[0] = strings.Replace(splitVals[0], "[", "", -1)
	}
	if strings.Contains(splitCity[0], "[") {
		splitCity[0] = strings.Replace(splitCity[0], "[", "", -1)
	}
	var list ListCityValue
	if len(splitVals) == len(splitCity) {
		for i := 0; i < len(splitVals); i++ {
			splitCity[i] = strings.Replace(splitCity[i], "\"", "", -1)
			list.data = append(list.data, cityValue{
				City:  splitCity[i],
				Value: splitVals[i],
			})
		}
	}
	return list
}

func createCSV(list ListCityValue) {
	FILENAME := "data.csv"
	file, err := os.OpenFile(FILENAME, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("[ERR] Cannot open file %v :: %v", FILENAME, err)
	}

	defer func() {
		cerr := file.Close()
		if cerr != nil {
			log.Fatalf("[ERR] Cannot close file %v :: %v", FILENAME, cerr)
		}
	}()

	var vStateValue []*cityValue
	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		log.Fatalf("[ERR] Cannot start file %v :: %v", FILENAME, err)
	}

	for _, v := range list.data {
		vStateValue = append(vStateValue, &cityValue{City: v.City, Value: v.Value})
	}
	err = gocsv.MarshalFile(&vStateValue, file) // Use this to save the CSV back to the file
	if err != nil {
		log.Fatalf("[ERR] Cannot save file %v :: %v", FILENAME, err)
	}
}

func main() {
	var client http.Client
	resp, err := client.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		list := getCityValue(bodyString)
		createCSV(list)
	}
}
