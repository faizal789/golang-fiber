package model

type Halaman struct{
	Size 	  	  int			 `json:"size"`
	Totalelements int64			 `json:"total"`
	Pagenumber 	  int 			 `json:"pagenumber"`
	Param map	  [string]string `json:"param"`
}


type Pagedata struct{
	Data interface{} 	`json:"data,omitempty"`
	Page Halaman 		`json:"hal"`
}

