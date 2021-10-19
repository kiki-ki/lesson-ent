package request

type CompanyUpdateReq struct {
	Name string `json:"name"`
}

type CompanyCreateWithUserReq struct {
	CompanyName string `json:"companyName"`
	UserName    string `json:"userName"`
	UserRole    string `json:"userRole"`
	UserComment string `json:"userComment"`
}
