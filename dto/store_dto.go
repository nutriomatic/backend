package dto

type StoreRegisterForm struct {
	StoreName     string `json:"store_name"`
	StoreUsername string `json:"store_username"`
	StoreAddress  string `json:"store_address"`
	StoreContact  string `json:"store_contact"`
}
