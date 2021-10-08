package nftlabs

type Storage interface {
	Get(uri string) ([]byte, error)
	Upload(data interface{}, contractAddress string, signerAddress string) (string, error)
	UploadBatch(data []interface{}, contractAddress string, signerAddress string) ([]string, error)
}
