package nometa

import "github.com/petuhovskiy/cian-search/cian"

type Building struct {
	BuildYear   int
	FloorsCount int
}

func ConvertBuilding(b cian.Building) Building {
	return Building{
		BuildYear:   b.BuildYear,
		FloorsCount: b.FloorsCount,
	}
}


