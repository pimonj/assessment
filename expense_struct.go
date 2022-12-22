package main

type ReqExpense struct {
	AMOUNT 		int 	`json:"amount"`
	NOTE 		string  `json:"note"`
	TAGS 		string  `json:"tags"`
	TITLE 		string  `json:"title"`
}


type ResExpense struct {
	AMOUNT 		int 	`json:"amount"`
	ID			string  `json:"id"`
	NOTE 		string  `json:"note"`
	TAGS 		string  `json:"tags"`
	TITLE 		string  `json:"title"`
}