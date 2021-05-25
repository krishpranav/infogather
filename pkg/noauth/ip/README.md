# ip look up

- ip lookup struct
```go
type IPAddress struct {
	Status     string  `json:"status"`
	Continent  string  `json:"continent"`
	Country    string  `json:"country"`
	RegionName string  `json:"regionname"`
	City       string  `json:"city"`
	District   string  `json:"district"`
	Zip        string  `json:"zip"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	Timezone   string  `json:"timezone"`
	Currency   string  `json:"currency"`
	Isp        string  `json:"isp"`
	Org        string  `json:"org"`
	As         string  `json:"as"`
	Asname     string  `json:"asname"`
	Reverse    string  `json:"reverse"`
	Mobile     bool    `json:"mobile"`
	Proxy      bool    `json:"proxy"`
	Hosting    bool    `json:"hosting"`
	Query      string  `json:"query"`
} 
```

# functions
```go
func IPLookup(ip string) (*IPAddress, error) {
	resp, err := http.GetReq("http://ip-api.com/json/" + ip + "?fields=31162361")
	if err != nil {
		return nil, err
	}
	s := new(IPAddress)
	err = json.Unmarshal([]byte(resp), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
```