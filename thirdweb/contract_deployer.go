package thirdweb

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/thirdweb-dev/go-sdk/internal/abi"
)

type ContractDeployer struct {
	*ProviderHandler
	factory *abi.TWFactory
	helper  *contractHelper
	storage storage
}

func newContractDeployer(provider *ethclient.Client, privateKey string, storage storage) (*ContractDeployer, error) {
	handler, err := NewProviderHandler(provider, privateKey)
	if err != nil {
		return nil, err
	}

	chainId, err := provider.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	factoryAddress, err := getContractAddressByChainId(ChainID(chainId.Int64()), "TWFactory")
	if err != nil {
		return nil, err
	}

	helper, err := newContractHelper(common.HexToAddress(factoryAddress), provider, privateKey)
	if err != nil {
		return nil, err
	}

	factory, err := abi.NewTWFactory(common.HexToAddress(factoryAddress), provider)
	if err != nil {
		return nil, err
	}

	contractDeployer := &ContractDeployer{
		handler,
		factory,
		helper,
		storage,
	}

	return contractDeployer, nil
}

func (deployer *ContractDeployer) DeployNFTCollection(metadata *NFTCollectionContractMetadata) (string, error) {
	return deployer.deployContract("nft-collection", metadata)
}

func (deployer *ContractDeployer) DeployEditionCollection(metadata *EditionContractMetadata) (string, error) {
	return deployer.deployContract("edition", metadata)
}

func (deployer *ContractDeployer) DeployToken(metadata *TokenContractMetadata) (string, error) {
	return deployer.deployContract("token", metadata)
}

func (deployer *ContractDeployer) DeployNFTDropCollection(metadata *NFTDropContractMetadata) (string, error) {
	return deployer.deployContract("nft-drop", metadata)
}

func (deployer *ContractDeployer) DeployEditionDropCollection(metadata *EditionDropContractMetadata) (string, error) {
	return deployer.deployContract("edition-drop", metadata)
}

func (deployer *ContractDeployer) DeployMultiwrap(metadata *MultiwrapContractMetadata) (string, error) {
	return deployer.deployContract("multiwrap", metadata)
}

func (deployer *ContractDeployer) deployContract(contractType string, metadata interface{}) (string, error) {
	contractUri, err := deployer.storage.Upload(
		metadata, deployer.helper.getAddress().String(),
		deployer.helper.GetSignerAddress().String(),
	)
	if err != nil {
		return "", err
	}

	encodedType, err := deployer.getEncodedType(contractType)
	if err != nil {
		return "", err
	}

	encodedFunc, err := deployer.getEncodedFunc(contractType, metadata, contractUri)
	if err != nil {
		return "", err
	}

	opts, err := deployer.helper.getTxOptions()
	if err != nil {
		return "", err
	}

	tx, err := deployer.factory.DeployProxy(opts, encodedType, encodedFunc)
	if err != nil {
		return "", err
	}

	receipt, err := deployer.helper.awaitTx(tx.Hash())
	if err != nil {
		return "", err
	}

	fmt.Println(receipt)

	return "", nil
}

func (deployer *ContractDeployer) getEncodedType(contractType string) ([32]byte, error) {
	var remoteName string

	switch contractType {
	case "nft-collection":
		remoteName = "TokenERC721"
	case "edition":
		remoteName = "TokenERC1155"
	case "token":
		remoteName = "TokenERC20"
	case "nft-drop":
		remoteName = "DropERC721"
	case "edition-drop":
		remoteName = "DropERC1155"
	case "multiwrap":
		remoteName = "Multiwrap"
	default:
		return [32]byte{}, fmt.Errorf("Unsupported contract type: %s", contractType)
	}

	contractBytes := []byte(remoteName)
	encodedType := [32]byte{}
	copy(encodedType[:], contractBytes)

	return encodedType, nil
}

func (deployer *ContractDeployer) getEncodedFunc(contractType string, metadata interface{}, contractUri string) ([]byte, error) {
	trustedForwarders, err := deployer.getDefaultTrustForwarders()
	if err != nil {
		return nil, err
	}

	switch contractType {
	case "nft-collection", "nft-drop":
		meta, ok := metadata.(*NFTDropContractMetadata)
		if !ok {
			return nil, err
		}

		contract, err := abi.NewDropERC721(common.HexToAddress(zeroAddress), deployer.GetProvider())
		if err != nil {
			return nil, err
		}

		tx, err := contract.Initialize(
			&bind.TransactOpts{},
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
			common.HexToAddress(meta.PlatformFeeRecipient),
		)
		if err != nil {
			return nil, err
		}

		return tx.Data(), nil
	case "edition", "edition-drop":
		meta, ok := metadata.(*EditionDropContractMetadata)
		if !ok {
			return nil, err
		}

		contract, err := abi.NewDropERC1155(common.HexToAddress(zeroAddress), deployer.GetProvider())
		if err != nil {
			return nil, err
		}

		tx, err := contract.Initialize(
			&bind.TransactOpts{},
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
			common.HexToAddress(meta.PlatformFeeRecipient),
		)
		if err != nil {
			return nil, err
		}

		return tx.Data(), nil
	case "token":
		meta, ok := metadata.(*TokenContractMetadata)
		if !ok {
			return nil, err
		}

		contract, err := abi.NewTokenERC20(common.HexToAddress(zeroAddress), deployer.GetProvider())
		if err != nil {
			return nil, err
		}

		tx, err := contract.Initialize(
			&bind.TransactOpts{},
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.PrimarySaleRecipient),
			common.HexToAddress(meta.PlatformFeeRecipient),
			big.NewInt(int64(meta.PlatformFeeBasisPoints)),
		)
		if err != nil {
			return nil, err
		}

		return tx.Data(), nil
	case "multiwrap":
		meta, ok := metadata.(*MultiwrapContractMetadata)
		if !ok {
			return nil, err
		}

		contract, err := abi.NewMultiwrap(common.HexToAddress(zeroAddress), deployer.GetProvider())
		if err != nil {
			return nil, err
		}

		tx, err := contract.Initialize(
			&bind.TransactOpts{},
			deployer.GetSignerAddress(),
			meta.Name,
			meta.Symbol,
			contractUri,
			trustedForwarders,
			common.HexToAddress(meta.FeeRecipient),
			big.NewInt(int64(meta.SellerFeeBasisPoints)),
		)
		if err != nil {
			return nil, err
		}

		return tx.Data(), nil
	default:
		return nil, fmt.Errorf("Unsupported contract type: %s", contractType)
	}
}

func (deployer *ContractDeployer) getDefaultTrustForwarders() ([]common.Address, error) {
	chainId, err := deployer.GetProvider().ChainID(context.Background())
	if err != nil {
		return []common.Address{}, err
	}

	biconomyForwarder, err := getContractAddressByChainId(ChainID(chainId.Int64()), "BiconomyForwarder")
	if err != nil {
		return []common.Address{}, err
	}

	if biconomyForwarder != zeroAddress {
		return []common.Address{
			common.HexToAddress(ozDefenderForwarderAddress),
			common.HexToAddress(biconomyForwarder),
		}, nil
	}

	return []common.Address{
		common.HexToAddress(ozDefenderForwarderAddress),
	}, nil
}
