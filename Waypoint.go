package main

type Waypoint struct {
	Timestamp string `json:"at,required" description:"The time when the device changed location"`
	Lonitude  string `json:"lon" description:"The longitude of the device"`
	Latitude  string `json:"lat" description:"The latitude of the device"`
	Elevation string `json:"ele" description:"The elevation of the device"`
}
