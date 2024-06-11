package models

type Item struct {
	ServiceDescription string  `json:"serviceDescription"`
	OnuCode            int     `json:"onuCode"`
	MeasureUnit        string  `json:"measureUnit"`
	Quantity           float32 `json:"quantity"`
	Generic            string  `json:"generic"`
	Level1             string  `json:"level1"`
	Level2             string  `json:"level2"`
	Level3             string  `json:"level3"`
}

type Match struct {
	FilterName string `json:"filterName" binding:"required"`
	Keyword    string `json:"keyword" binding:"required"`
	On         string `json:"on" binding:"required"`
}

type Filter struct {
	Date    string  `json:"date" binding:"required,datetime=2006-01-02"`
	Matches []Match `json:"matches" binding:"required"`
}

type Tender struct {
	Code             string   `json:"code" binding:"required"`
	Type             string   `json:"type"`
	StateCode        int      `json:"stateCode" binding:"required"`
	EndDate          string   `json:"endDate" binding:"required,datetime=2006-01-02"`
	Name             string   `json:"name" binding:"required"`
	Description      string   `json:"description"`
	Organism         string   `json:"organism"`
	PurchasingRegion string   `json:"purchasingRegion"`
	PublicationDate  string   `json:"publicationDate"`
	Categories       []string `json:"categories"`
	Filter           Filter   `json:"filter"`
	Items            []Item   `json:"items"`
}
