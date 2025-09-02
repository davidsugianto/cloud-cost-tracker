package provider

type CloudProvider string

const (
	ProviderGCP     CloudProvider = "gcp"
	ProviderAWS     CloudProvider = "aws"
	ProviderAlibaba CloudProvider = "aliyun"
)

type ProviderAccount struct {
	ID          string        `json:"id"`
	Provider    CloudProvider `json:"provider"`
	Name        string        `json:"name"`        // friendly name
	ExternalID  string        `json:"external_id"` // project/account/subscription ID
	Credentials string        `json:"-"`           // encrypted/secret store ref
	Active      bool          `json:"active"`
}
