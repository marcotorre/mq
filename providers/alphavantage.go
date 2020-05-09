package providers

type JsonResp struct {
	Fields JsonRespFields `json:"Global Quote"`
}

type JsonRespFields struct {
	Price      string `json:"05. price"`
	ChgPercent string `json:"10. change percent"`
}
