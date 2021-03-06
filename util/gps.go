package util

import (
	"math"
)

func GetDistance(lat1, lng1, lat2, lng2 float64) float64 {
	radius := 6378.137
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius
}

func GetMid(lat1, lng1, lat2, lng2 float64) (lat float64, lng float64) {
	lat = (lat1 + lat2) / 2
	lng = (lng1 + lng2) / 2
	return
}
