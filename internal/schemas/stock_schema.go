package schemas

type StockFilter struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type StockResponse struct {
	Items []StockResponseItem `json:"items"`
}

type StockResponseItem struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}
