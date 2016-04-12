package main

type Location struct {
	Disposition string     `json:"disposition" description:"Whether the location is mobile or fixed"`
	Elevation   string     `json:"ele" description:"The elevation of the device"`
	Name        string     `json:"name" description:"The name of the device"`
	Exposure    string     `json:"exposure" description:"Whether the location is indoors or outdoors"`
	Lonitude    string     `json:"lon" description:"The longitude of the device"`
	Latitude    string     `json:"lat" description:"The latitude of the device"`
	Domain      string     `json:"domain" description:"	The domain of the location, i.e. physical or virtual"`
	Waypoints   []Waypoint `json:"waypoints" description:"A list of locations for a mobile Feed"`
}
