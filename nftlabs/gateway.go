package nftlabs

type Gateway interface {
	Get(uri string) ([]byte, error)
}
