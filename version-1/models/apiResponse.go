package models

// APIResponse is the response obtained from the 3rd party API
type APIResponse struct {
	City       string    `json:"name"`
	Coord      coordData `json:"coord"`
	Clouds     cloudData `json:"clouds"`
	Wind       windData  `json:"wind"`
	Main       mainData  `json:"main"`
	Sys        sysData   `json:"sys"`
	StatusCode string    `json:"cod"`
	Message    string    `json:"message"`
}

// Nested structs for the APIResponse
type coordData struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}
type mainData struct {
	Temperature float32 `json:"temp"`
	Humidity    int     `json:"humidity"`
	Pressure    int     `json:"pressure"`
}
type windData struct {
	Speed float32 `json:"speed"`
}
type sysData struct {
	Country string `json:"country"`
	Sunrise int64  `json:"sunrise"`
	Sunset  int64  `json:"sunset"`
}
type cloudData struct {
	Cloudiness int `json:"all"`
}
