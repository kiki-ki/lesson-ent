package request

type CompanyUpdateReq struct {
	Name string `json:"name"`
}

type CompanyCreateWithUserReq struct {
	CompanyName string `json:"companyName"`
	UserName    string `json:"userName"`
	UserEmail   string `json:"userEmail"`
	UserComment string `json:"userComment"`
}
