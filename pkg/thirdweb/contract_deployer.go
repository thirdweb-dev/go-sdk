package thirdweb

import "github.com/ethereum/go-ethereum/ethclient"

type ContractDeployer struct {
	*ProviderHandler
	factory any
	storage Storage
}

func NewContractDeployer(provider *ethclient.Client, privateKey string, storage Storage) (*ContractDeployer, error) {
	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	contractDeployer := &ContractDeployer{
		handler,
		[]any{},
		storage,
	}

	return contractDeployer, nil
}

func DeployNFTCollection(metadata *NFTCollectionContractMetadata) (string, error) {
	return "", nil
}

func DeployEditionCollection(metadata *EditionContractMetadata) (string, error) {
	return "", nil
}

func DeployNFTDropCollection(metadata *NFTDropContractMetadata) (string, error) {
	return "", nil
}

func deployContract(contractType string, metadata interface{}) (string, error) {
	return "", nil
}
