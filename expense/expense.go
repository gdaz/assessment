package expense

type ExpenseBody struct {
	Id     int64    `json:"id"`
	Title  string   `json:"title"`
	Amount int64    `json:"amount"`
	Note   string   `json:"note"`
	Tags   []string `json:"tags"`
}

type Tag struct {
	Tag []string `json:"tags"`
}

type Err struct {
	Message string `json:"message"`
}
