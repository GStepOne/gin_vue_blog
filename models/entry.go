package models

type ESIDRequest struct {
	ID string `json:"id" form:"id" uri:"id"`
}

type ESIDListRequest struct {
	IDList []string `json:"id_list" binding:"required"`
}
