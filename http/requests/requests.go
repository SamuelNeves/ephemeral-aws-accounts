package requests

type ApplyRequest struct {
	AccountName string `json:"account-name"`
	ParentId    string `json:"parent-id"`

	ResourcesToCreate string
}
