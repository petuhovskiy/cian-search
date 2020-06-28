package nometa

import "github.com/petuhovskiy/cian-search/cian"

type User struct {
	AccountType       string `json:"accountType"`
	AgencyName        string `json:"agencyName"`
	CianProfileStatus string `json:"cianProfileStatus"`
	CompanyName       string `json:"companyName"`
	Experience        string `json:"experience"`

	IsAgent            bool `json:"isAgent"`
	IsBuilder          bool `json:"isBuilder"`
	IsCallbackUser     bool `json:"isCallbackUser"`
	IsPassportVerified bool `json:"isPassportVerified"`
	IsSelfEmployed     bool `json:"isSelfEmployed"`
	IsSubAgent         bool `json:"isSubAgent"`

	UserTrustLevel string `json:"userTrustLevel"`

	// TODO: rating. cannot found :(
	UserType      string `json:"userType"`
	IsCianPartner bool   `json:"isCianPartner"`
}

func ConvertUser(u cian.User) User {
	return User{
		AccountType:        u.AccountType,
		AgencyName:         u.AgencyName,
		CianProfileStatus:  u.CianProfileStatus,
		CompanyName:        u.CompanyName,
		Experience:         u.Experience,
		IsAgent:            u.IsAgent,
		IsBuilder:          u.IsBuilder,
		IsCallbackUser:     u.IsCallbackUser,
		IsPassportVerified: u.IsPassportVerified,
		IsSelfEmployed:     u.IsSelfEmployed,
		IsSubAgent:         u.IsSubAgent,
		UserTrustLevel:     u.UserTrustLevel,
		UserType:           u.UserType,
		IsCianPartner:      u.IsCianPartner,
	}
}
