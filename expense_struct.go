package main

type ReqExpense struct {
	AMOUNT 		float64 	 `json:"amount"`
	NOTE 		string  	 `json:"note"`
	TAGS 		[]string  	 `json:"tags"`
	TITLE 		string   	 `json:"title"`
}

type ResExpense struct {
	ID			string  	`json:"id"`
	TITLE 		string      `json:"title"`
	AMOUNT 		float64 	`json:"amount"`	
	NOTE 		string  	`json:"note"`
	TAGS 		[]string    `json:"tags"`	
}