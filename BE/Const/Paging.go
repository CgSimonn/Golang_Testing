package Const

type Paging struct {
	PageIndex int64 `json:"index" form:"index"`
	PageSize  int64 `json:"size" form:"size"`
}
