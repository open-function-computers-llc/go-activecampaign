package activecampaign

type ResponseMeta struct {
	Total     string `json:"total"`
	Sortable  bool   `json:"sortable"`
	PageInput struct {
		Segmentid  int         `json:"segmentid"`
		Formid     int         `json:"formid"`
		Listid     int         `json:"listid"`
		Tagid      int         `json:"tagid"`
		Limit      int         `json:"limit"`
		Offset     int         `json:"offset"`
		Search     interface{} `json:"search"`
		Sort       interface{} `json:"sort"`
		Seriesid   int         `json:"seriesid"`
		Waitid     int         `json:"waitid"`
		Status     int         `json:"status"`
		ForceQuery int         `json:"forceQuery"`
		Cacheid    string      `json:"cacheid"`
		Email      string      `json:"email"`
	} `json:"page_input"`
}
