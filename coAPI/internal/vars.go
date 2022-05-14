package coAPI

type Change struct {
	ID    int
	Attrs Attrs
}

type Attrs struct {
	Requester      *string `json:"requester"`
	Env            *string `json:"env"`
	Type           *string `json:"type"`
	CustomerImpact *string `json:"customer_impact"`
	Description    *string `json:"description"`
	Date           *string `json:"date"`
	Link           *string `json:"link"`
	LinkBack       *string `json:"link_back"`
}
