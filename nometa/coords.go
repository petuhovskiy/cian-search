package nometa

import "github.com/petuhovskiy/cian-search/cian"

type Coords struct {
	Lng float64
	Lat float64
}

func ConvertCoords(c cian.Coordinates) Coords {
	return Coords{
		Lng: c.Lng,
		Lat: c.Lat,
	}
}
