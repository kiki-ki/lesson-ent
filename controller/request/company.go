package request

import "github.com/guregu/null"

type CompanyUpdateReq struct {
	Name string `json:"name"`
}

type CompanyCreateWithUserReq struct {
	CompanyName string      `json:"companyName"`
	UserName    string      `json:"userName"`
	UserEmail   string      `json:"userEmail"`
	UserComment null.String `json:"userComment"`
}
