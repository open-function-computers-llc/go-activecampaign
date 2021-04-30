package activecampaign

type ContactsResponse struct {
	Contacts []Contact    `json:"contacts"`
	Meta     ResponseMeta `json:"meta"`
}
