package nftlabs

type Gateway interface {
	Get(uri string) ([]byte, error)
	Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
}
