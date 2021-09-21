package nftlabs

type NftMetadata struct {
	Id string
	Uri string `json:"external_url"`
	Name string `json:"name"`
	Description string `json:"description"`
	Image string `json:"image"`
	Attributes interface{} `json:"attributes"`
}