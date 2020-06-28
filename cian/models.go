package cian

import "time"

type OffersResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}
type DegradationInfo struct {
	OfferChatsV2       bool `json:"offerChatsV2"`
	Users              bool `json:"users"`
	OfferChats         bool `json:"offerChats"`
	Collections        bool `json:"collections"`
	PreviewSavedSearch bool `json:"previewSavedSearch"`
	Favorites          bool `json:"favorites"`
	Auctions           bool `json:"auctions"`
	Ratings            bool `json:"ratings"`
	AgentAvailability  bool `json:"agentAvailability"`
	Villages           bool `json:"villages"`
	QuickLinks         bool `json:"quickLinks"`
	Notes              bool `json:"notes"`
}
type Kp struct {
}
type Room struct {
	Type  string `json:"type"`
	Value []int  `json:"value"`
}
type ForDay struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}
type Price struct {
	Type  string `json:"type"`
	Value Value  `json:"value"`
}
type EngineVersion struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}
type Currency struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}
type Region struct {
	Type  string `json:"type"`
	Value []int  `json:"value"`
}
type Value map[string]interface{}
type GeoQuery struct {
	Type  string  `json:"type"`
	Value []Value `json:"value"`
}
type Page struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}
type JSONQuery struct {
	Type          string        `json:"_type"`
	Room          Room          `json:"room"`
	ForDay        ForDay        `json:"for_day"`
	Price         Price         `json:"price"`
	EngineVersion EngineVersion `json:"engine_version"`
	Currency      Currency      `json:"currency"`
	Region        Region        `json:"region"`
	Geo           GeoQuery      `json:"geo"`
	Page          Page          `json:"page"`
}
type AvgPriceInformer struct {
	Currency string `json:"currency"`
	Price    int    `json:"price"`
}
type Items struct {
	Key             string      `json:"key"`
	BackgroundColor interface{} `json:"backgroundColor"`
	TextColor       interface{} `json:"textColor"`
	Title           string      `json:"title"`
	URL             string      `json:"url"`
	Visibility      []string    `json:"visibility"`
	Label           string      `json:"label"`
	QueryString     string      `json:"queryString"`
}
type Collections struct {
	Items []Items `json:"items"`
}
type SeoData struct {
	MetaDescription string        `json:"metaDescription"`
	Title           string        `json:"title"`
	Keywords        interface{}   `json:"keywords"`
	Noindex         bool          `json:"noindex"`
	SimilarLinks    []interface{} `json:"similarLinks"`
	Canonical       interface{}   `json:"canonical"`
	Text            interface{}   `json:"text"`
	Description     string        `json:"description"`
	H1              string        `json:"h1"`
}
type QsToUris struct {
}
type Breadcrumbs struct {
	URL   string `json:"url"`
	Title string `json:"title"`
}
type QuickLinks struct {
}
type Building struct {
	CargoLiftsCount     interface{} `json:"cargoLiftsCount"`
	BuildYear           int         `json:"buildYear"`
	FloorsCount         int         `json:"floorsCount"`
	MaterialType        string      `json:"materialType"`
	Deadline            interface{} `json:"deadline"`
	AccessType          interface{} `json:"accessType"`
	HeatingType         interface{} `json:"heatingType"`
	Type                interface{} `json:"type"`
	PassengerLiftsCount int         `json:"passengerLiftsCount"`
}
type UtilitiesTerms struct {
	Price                        int         `json:"price"`
	FlowMetersNotIncludedInPrice interface{} `json:"flowMetersNotIncludedInPrice"`
	IncludedInPrice              bool        `json:"includedInPrice"`
}
type BargainTerms struct {
	UtilitiesTerms    UtilitiesTerms `json:"utilitiesTerms"`
	PriceRur          int            `json:"priceRur"`
	AgentBonus        interface{}    `json:"agentBonus"`
	PaymentPeriod     string         `json:"paymentPeriod"`
	BargainAllowed    bool           `json:"bargainAllowed"`
	SaleType          interface{}    `json:"saleType"`
	Currency          string         `json:"currency"`
	ClientFee         int            `json:"clientFee"`
	VatType           interface{}    `json:"vatType"`
	AgentFee          int            `json:"agentFee"`
	Price             int            `json:"price"`
	PriceType         string         `json:"priceType"`
	LeaseTermType     string         `json:"leaseTermType"`
	LeaseType         interface{}    `json:"leaseType"`
	PriceForWorkplace interface{}    `json:"priceForWorkplace"`
	Deposit           int            `json:"deposit"`
	VatPrice          interface{}    `json:"vatPrice"`
	MortgageAllowed   interface{}    `json:"mortgageAllowed"`
}
type Phones struct {
	Number      string `json:"number"`
	CountryCode string `json:"countryCode"`
}
type Coordinates struct {
	Lng float64 `json:"lng"`
	Lat float64 `json:"lat"`
}
type Railways struct {
	Name         string  `json:"name"`
	DirectionIds []int   `json:"directionIds"`
	Time         int     `json:"time"`
	TravelType   string  `json:"travelType"`
	GeoType      string  `json:"geoType"`
	Distance     float64 `json:"distance"`
	ID           int     `json:"id"`
	Qs           string  `json:"qs"`
}
type Districts struct {
	Name       string `json:"name"`
	Qs         string `json:"qs"`
	LocationID int    `json:"locationId"`
	FullName   string `json:"fullName"`
	ParentID   int    `json:"parentId"`
	GeoType    string `json:"geoType"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	ID         int    `json:"id"`
}
type Address struct {
	Name             string `json:"name"`
	Qs               string `json:"qs"`
	ShortName        string `json:"shortName"`
	FullName         string `json:"fullName"`
	LocationTypeID   int    `json:"locationTypeId"`
	GeoType          string `json:"geoType"`
	Title            string `json:"title"`
	Type             string `json:"type"`
	ID               int    `json:"id"`
	IsFormingAddress bool   `json:"isFormingAddress"`
}
type Undergrounds struct {
	Name              string      `json:"name"`
	LineColor         string      `json:"lineColor"`
	CianID            interface{} `json:"cianId"`
	TransportType     string      `json:"transportType"`
	Qs                string      `json:"qs"`
	Time              int         `json:"time"`
	ReleaseYear       interface{} `json:"releaseYear"`
	IsDefault         bool        `json:"isDefault"`
	LineID            int         `json:"lineId"`
	UnderConstruction bool        `json:"underConstruction"`
	ID                int         `json:"id"`
}
type SortedGeoIds struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}
type Geo struct {
	Jk              interface{}    `json:"jk"`
	Coordinates     Coordinates    `json:"coordinates"`
	BuildingAddress interface{}    `json:"buildingAddress"`
	Railways        []Railways     `json:"railways"`
	Highways        []interface{}  `json:"highways"`
	Districts       []Districts    `json:"districts"`
	CountryID       int            `json:"countryId"`
	Address         []Address      `json:"address"`
	UserInput       string         `json:"userInput"`
	Undergrounds    []Undergrounds `json:"undergrounds"`
	SortedGeoIds    []SortedGeoIds `json:"sortedGeoIds"`
}
type Photos struct {
	ThumbnailURL  string      `json:"thumbnailUrl"`
	MiniURL       string      `json:"miniUrl"`
	FullURL       string      `json:"fullUrl"`
	IsDefault     bool        `json:"isDefault"`
	IsLayout      interface{} `json:"isLayout"`
	RotateDegree  interface{} `json:"rotateDegree"`
	Thumbnail2URL string      `json:"thumbnail2Url"`
	ID            int         `json:"id"`
}
type Specialty struct {
	AdditionalTypes []interface{} `json:"additionalTypes"`
	Specialties     []interface{} `json:"specialties"`
	Types           []interface{} `json:"types"`
}
type Notes struct {
	Offer   interface{} `json:"offer"`
	Realtor interface{} `json:"realtor"`
}
type Packages struct {
	CountryID       int    `json:"countryId"`
	DepositRequired bool   `json:"depositRequired"`
	GeoName         string `json:"geoName"`
	LocationID      int    `json:"locationId"`
	PackageTypeID   int    `json:"packageTypeId"`
	PackageTypeName string `json:"packageTypeName"`
}

type PhoneNumbers struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}
type AdfoxParams map[string]interface{}
type NewbuildingDynamicCalltracking struct {
	SiteBlockID interface{} `json:"siteBlockId"`
}

type AgentAvailability struct {
	Available     bool      `json:"available"`
	AvailableFrom time.Time `json:"availableFrom"`
	AvailableTo   time.Time `json:"availableTo"`
	Message       string    `json:"message"`
	Title         string    `json:"title"`
	UserID        int       `json:"userId"`
}
type User struct {
	AccountType        string            `json:"accountType"`
	AgencyName         string            `json:"agencyName"`
	AgentAvatarURL     string            `json:"agentAvatarUrl"`
	BillingInfo        BillingInfo       `json:"billingInfo"`
	CianProfileStatus  string            `json:"cianProfileStatus"`
	CianUserID         string            `json:"cianUserId"`
	CompanyName        string            `json:"companyName"`
	Experience         string            `json:"experience"`
	IsAgent            bool              `json:"isAgent"`
	IsBuilder          bool              `json:"isBuilder"`
	IsCallbackUser     bool              `json:"isCallbackUser"`
	IsChatsEnabled     bool              `json:"isChatsEnabled"`
	IsHidden           bool              `json:"isHidden"`
	IsPassportVerified bool              `json:"isPassportVerified"`
	IsSelfEmployed     bool              `json:"isSelfEmployed"`
	IsSubAgent         bool              `json:"isSubAgent"`
	PhoneNumbers       []PhoneNumbers    `json:"phoneNumbers"`
	ProfileURI         string            `json:"profileUri"`
	UserTrustLevel     string            `json:"userTrustLevel"`
	UserID             int               `json:"userId"`
	UserType           string            `json:"userType"`
	IsCianPartner      bool              `json:"isCianPartner"`
	CanShowOnline      bool              `json:"canShowOnline"`
	PersonalRating     interface{}       `json:"personalRating"`
	AgentAvailability  AgentAvailability `json:"agentAvailability"`
}

type BillingInfo struct {
	AccountType       int        `json:"accountType"`
	Packages          []Packages `json:"packages"`
	RemoveCompetitor  bool       `json:"removeCompetitor"`
	ReplenishDisabled bool       `json:"replenishDisabled"`
}

type MasterAgent struct {
	AbsoluteAvatarURL     string `json:"absoluteAvatarUrl"`
	AbsoluteMiniAvatarURL string `json:"absoluteMiniAvatarUrl"`
	AccountType           string `json:"accountType"`
	AvatarURL             string `json:"avatarUrl"`
	ID                    int    `json:"id"`
	IsModerationPassed    bool   `json:"isModerationPassed"`
	MiniAvatarURL         string `json:"miniAvatarUrl"`
	Name                  string `json:"name"`
	ProfileURI            string `json:"profileUri"`
}

type OffersSerialized struct {
	IsRecidivist                   bool                           `json:"isRecidivist"`
	RosreestrCheck                 interface{}                    `json:"rosreestrCheck"`
	Land                           interface{}                    `json:"land"`
	DescriptionWordsHighlighted    []interface{}                  `json:"descriptionWordsHighlighted"`
	Layout                         interface{}                    `json:"layout"`
	BasicProfiScore                int                            `json:"basicProfiScore"`
	DescriptionMinhash             []int                          `json:"descriptionMinhash"`
	Garage                         interface{}                    `json:"garage"`
	Building                       Building                       `json:"building"`
	IsImported                     bool                           `json:"isImported"`
	AllFromOffrep                  interface{}                    `json:"allFromOffrep"`
	RoomArea                       interface{}                    `json:"roomArea"`
	RoomsForSaleCount              interface{}                    `json:"roomsForSaleCount"`
	Newbuilding                    interface{}                    `json:"newbuilding"`
	RentByPartsDescription         interface{}                    `json:"rentByPartsDescription"`
	LivingArea                     interface{}                    `json:"livingArea"`
	OfferType                      string                         `json:"offerType"`
	ID                             int                            `json:"id"`
	NewbuildingVasPromotion        interface{}                    `json:"newbuildingVasPromotion"`
	LoggiasCount                   int                            `json:"loggiasCount"`
	CianID                         int                            `json:"cianId"`
	BedroomsCount                  interface{}                    `json:"bedroomsCount"`
	CoworkingOfferType             interface{}                    `json:"coworkingOfferType"`
	IsInHiddenBase                 bool                           `json:"isInHiddenBase"`
	PermittedUseType               interface{}                    `json:"permittedUseType"`
	CategoriesIds                  []int                          `json:"categoriesIds"`
	Booking                        interface{}                    `json:"booking"`
	TotalArea                      string                         `json:"totalArea"`
	FullURL                        string                         `json:"fullUrl"`
	RoomsCount                     int                            `json:"roomsCount"`
	AddedTimestamp                 int                            `json:"addedTimestamp"`
	IsRentByParts                  bool                           `json:"isRentByParts"`
	FlatType                       string                         `json:"flatType"`
	ShareAmount                    interface{}                    `json:"shareAmount"`
	BargainTerms                   BargainTerms                   `json:"bargainTerms"`
	Status                         string                         `json:"status"`
	WorkTimeInfo                   interface{}                    `json:"workTimeInfo"`
	OfficeType                     interface{}                    `json:"officeType"`
	IsColorized                    bool                           `json:"isColorized"`
	Kp                             interface{}                    `json:"kp"`
	Phones                         []Phones                       `json:"phones"`
	PublishedUserID                int                            `json:"publishedUserId"`
	Title                          interface{}                    `json:"title"`
	AccessType                     interface{}                    `json:"accessType"`
	AreaParts                      interface{}                    `json:"areaParts"`
	DemolishedInMoscowProgramm     bool                           `json:"demolishedInMoscowProgramm"`
	IsAuction                      bool                           `json:"isAuction"`
	IsPaid                         bool                           `json:"isPaid"`
	JkURL                          interface{}                    `json:"jkUrl"`
	IsByHomeowner                  bool                           `json:"isByHomeowner"`
	Geo                            Geo                            `json:"geo"`
	UserID                         int                            `json:"userId"`
	BalconiesCount                 interface{}                    `json:"balconiesCount"`
	Category                       string                         `json:"category"`
	FromDeveloper                  interface{}                    `json:"fromDeveloper"`
	Description                    string                         `json:"description"`
	PromoInfo                      interface{}                    `json:"promoInfo"`
	Videos                         []interface{}                  `json:"videos"`
	IsTop3                         bool                           `json:"isTop3"`
	Photos                         []Photos                       `json:"photos"`
	Added                          string                         `json:"added"`
	MinArea                        interface{}                    `json:"minArea"`
	WorkplaceCount                 interface{}                    `json:"workplaceCount"`
	DealType                       string                         `json:"dealType"`
	KitchenArea                    interface{}                    `json:"kitchenArea"`
	CianUserID                     int                            `json:"cianUserId"`
	IsApartments                   bool                           `json:"isApartments"`
	BusinessShoppingCenter         interface{}                    `json:"businessShoppingCenter"`
	IsPremium                      bool                           `json:"isPremium"`
	IsExcludedFromAction           bool                           `json:"isExcludedFromAction"`
	Specialty                      Specialty                      `json:"specialty"`
	Similar                        interface{}                    `json:"similar"`
	CreationDate                   string                         `json:"creationDate"`
	HumanizedTimedelta             string                         `json:"humanizedTimedelta"`
	FloorNumber                    int                            `json:"floorNumber"`
	IsDuplicatedDescription        bool                           `json:"isDuplicatedDescription"`
	IsFavorite                     bool                           `json:"isFavorite"`
	Notes                          Notes                          `json:"notes"`
	User                           User                           `json:"user,omitempty"`
	IsPro                          bool                           `json:"isPro"`
	GaLabel                        string                         `json:"gaLabel"`
	AdfoxParams                    AdfoxParams                    `json:"adfoxParams,omitempty"`
	ChatID                         interface{}                    `json:"chatId"`
	Chat                           interface{}                    `json:"chat"`
	Auction                        interface{}                    `json:"auction"`
	IsNew                          bool                           `json:"isNew"`
	SavedSearchLabels              []interface{}                  `json:"savedSearchLabels"`
	IsDealRequestSubstitutionPhone bool                           `json:"isDealRequestSubstitutionPhone"`
	NewbuildingDynamicCalltracking NewbuildingDynamicCalltracking `json:"newbuildingDynamicCalltracking"`
	HomeownerProofs                []string                       `json:"homeownerProofs,omitempty"`
}
type Data struct {
	FullURL                        string             `json:"fullUrl"`
	TopHitsLinks                   []interface{}      `json:"topHitsLinks"`
	MaxAuctionBet                  int                `json:"maxAuctionBet"`
	MaxAuctionService              string             `json:"maxAuctionService"`
	DegradationInfo                DegradationInfo    `json:"degradationInfo"`
	AggregatedCount                int                `json:"aggregatedCount"`
	SuggestOffersSerializedList    []interface{}      `json:"suggestOffersSerializedList"`
	NewbuildingDynamicCalltracking interface{}        `json:"newbuildingDynamicCalltracking"`
	Kp                             Kp                 `json:"kp"`
	JSONQuery                      JSONQuery          `json:"jsonQuery"`
	MlRankingGUID                  interface{}        `json:"mlRankingGuid"`
	ExtendedOffersCount            int                `json:"extendedOffersCount"`
	AvgPriceInformer               AvgPriceInformer   `json:"avgPriceInformer"`
	SearchRequestID                string             `json:"searchRequestId"`
	Collections                    Collections        `json:"collections"`
	SeoLinks                       []interface{}      `json:"seoLinks"`
	SeoData                        SeoData            `json:"seoData"`
	QsToUris                       QsToUris           `json:"qsToUris"`
	RedirectData                   interface{}        `json:"redirectData"`
	Breadcrumbs                    []Breadcrumbs      `json:"breadcrumbs"`
	MlRankingModelVersion          interface{}        `json:"mlRankingModelVersion"`
	OfferCount                     int                `json:"offerCount"`
	ExtensionTypes                 []interface{}      `json:"extensionTypes"`
	LastModified                   interface{}        `json:"lastModified"`
	MlRankingResponseCached        interface{}        `json:"mlRankingResponseCached"`
	SearchUUID                     interface{}        `json:"searchUuid"`
	QuickLinks                     QuickLinks         `json:"quickLinks"`
	QueryString                    string             `json:"queryString"`
	OffersSerialized               []OffersSerialized `json:"offersSerialized"`
}
