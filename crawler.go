package crawler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gocarina/gocsv"
)

// CityValue ...
type CityValue struct {
	City  string `csv:"city"`
	Value uint64 `csv:"value"`
}

// ListCityValue ...
type ListCityValue struct {
	Data  []CityValue
	Total uint64
}

// Exported ...
var (
	URL      = "http://ti.saude.rs.gov.br/covid19/"
	FileName = "data.csv" // we can set path e.g /var/crawser/br/rs/Data.csv
	// REGEX ...
	rgxGraphCities = regexp.MustCompile(`GraphMunicipio.*\n.*,`)
	rgxGetJSON     = regexp.MustCompile(`\{.*\}`)
	rgxLables      = regexp.MustCompile(`\{labels:.*\],`)
	rgxValues      = regexp.MustCompile(`\[[\d\,]+\]`)
	rgxCities      = regexp.MustCompile(`\[[\p{L}\"\s\,\]\']+,d`)
)

// ReplaceValue ...
func ReplaceValue(splitCity, splitVals []string) ([]string, []string) {
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
	return splitCity, splitVals
}

// GetCityValue ...
func GetCityValue(body string) ([]string, []string) {
	sliceGraph := rgxGraphCities.FindString(body)
	getJSON := rgxGetJSON.FindString(sliceGraph)
	getJSON = rgxLables.FindString(getJSON)
	getValues := rgxValues.FindString(getJSON)
	getStates := rgxCities.FindString(getJSON)

	splitVals := strings.Split(getValues, ",")
	splitCity := strings.Split(getStates, ",")
	return splitCity, splitVals
}

// SetCitiesValues ...
func SetCitiesValues(splitCity, splitVals []string) ListCityValue {
	var list ListCityValue
	if len(splitVals) == len(splitCity) {
		var total uint64 = 0
		for i := 0; i < len(splitVals); i++ {
			splitCity[i] = strings.Replace(splitCity[i], "\"", "", -1)
			value := strings.Trim(splitVals[i], " ")
			val, err := strconv.ParseUint(value, 10, 64)
			if err != nil {
				log.Fatalf("[ERR] Cannot convert string to uint %v :: %v", val, err)
			}
			list.Data = append(list.Data, CityValue{
				City:  strings.Trim(splitCity[i], " "),
				Value: val,
			})
			total = total + val
		}
		list.Total = total
	}
	return list
}

func createCSV(list ListCityValue) {
	file, err := os.OpenFile(FileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		log.Fatalf("[ERR] Cannot open file %v :: %v", FileName, err)
	}

	defer func() {
		cerr := file.Close()
		if cerr != nil {
			log.Fatalf("[ERR] Cannot close file %v :: %v", FileName, cerr)
		}
	}()

	var vStateValue []*CityValue
	if _, err := file.Seek(0, 0); err != nil { // Go to the start of the file
		log.Fatalf("[ERR] Cannot start file %v :: %v", FileName, err)
	}

	for _, v := range list.Data {
		vStateValue = append(vStateValue, &CityValue{City: v.City, Value: v.Value})
	}
	// Append calc about total cases in RS
	vStateValue = append(vStateValue, &CityValue{City: "Total", Value: list.Total})

	err = gocsv.MarshalFile(&vStateValue, file) // Use this to save the CSV back to the file
	if err != nil {
		log.Fatalf("[ERR] Cannot save file %v :: %v", FileName, err)
	}
}

// GetData ...
func GetData() string {
	client := http.Client{Timeout: 2 * time.Second}
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
		return string(bodyBytes)
	}
	return ""
}

// Run ...
func Run() {
	body := GetData()
	cities, values := GetCityValue(body)
	cities, values = ReplaceValue(cities, values)
	list := SetCitiesValues(cities, values)
	createCSV(list)
}
