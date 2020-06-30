package nometa

import (
	"fmt"
	"time"

	"github.com/petuhovskiy/cian-search/cian"
)

type Offer struct {
	IsRecidivist bool // ????
	HomeOwner    bool
	Building     Building
	CianID       int
	CianURL      string
	TotalArea    float64
	RoomsCount   int
	Added        time.Time
	// FlatType "rooms" ??
	PriceRur int // ?!
	Price    int // ?!
	Deposit  int // ?!
	// PaymentPeriod: (string) (len=7) "monthly",
	// Currency: (string) (len=3) "rur",
	ClientFee int // ???
	AgentFee  int // ???
	// Status: (string) (len=9) "published",
	Coords  Coords
	Address string
	// Category: (string) (len=8) "flatRent",
	Description   string // !!!!!!!!
	IsApartaments bool
	User          User
	IsPro         bool
	IsDeleted     bool
}

func ConvertOffer(o cian.OffersSerialized) Offer {
	return Offer{
		IsRecidivist:  o.IsRecidivist,
		HomeOwner:     o.IsByHomeowner,
		Building:      ConvertBuilding(o.Building),
		CianID:        o.CianID,
		CianURL:       fmt.Sprintf("https://www.cian.ru/rent/flat/%d/", o.CianID),
		TotalArea:     parseFloat(o.TotalArea),
		RoomsCount:    o.RoomsCount,
		Added:         time.Unix(int64(o.AddedTimestamp), 0),
		PriceRur:      o.BargainTerms.PriceRur,
		ClientFee:     o.BargainTerms.ClientFee,
		AgentFee:      o.BargainTerms.AgentFee,
		Price:         o.BargainTerms.Price,
		Deposit:       o.BargainTerms.Deposit,
		Coords:        ConvertCoords(o.Geo.Coordinates),
		Address:       o.Geo.UserInput,
		Description:   o.Description,
		IsApartaments: o.IsApartments,
		User:          ConvertUser(o.User),
		IsPro:         o.IsPro,
		IsDeleted:     false,
	}
}
