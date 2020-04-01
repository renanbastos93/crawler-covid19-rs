![Go Build](https://github.com/renanbastos93/crawler-covid19-rs/workflows/Go%20Build/badge.svg)
![Gosec](https://github.com/renanbastos93/crawler-covid19-rs/workflows/Gosec/badge.svg)
![Go Test](https://github.com/renanbastos93/crawler-covid19-rs/workflows/Go%20Test/badge.svg)

### Crawler COVID-19 State Rio Grande do Sul at Brazil
This system gets data official website of government to extract data and generate a CSV.


#### Use
It is very simple only need clone project and executes commands below.
```bash
$ make run
# or
$ go run cmd/main.go
```


#### Using as a module
Simple code below, but you can use other methods publics as (GetData, GetCityValue, GetCityValue, ReplaceValue)
> `go get -u github.com/renanbastos93/crawler-covid19-rs ` 
```go
package main

import "github.com/renanbastos93/crawler-covid19-rs"

func main() {
  crawler.FileName = "yourfile.csv" // Default: data.csv but you can set path e.g /var/crawser/br/rs/Data.csv
  crawler.Run()
}
```
> Example without generate CSV

```go
package main

import "github.com/renanbastos93/crawler-covid19-rs"

func main() {
    body := GetData()
    cities, values := GetCityValue(body)
    cities, values = ReplaceValue(cities, values)
    list := SetCitiesValues(cities, values)
    // using list according to struct ListCityValue
}
```
