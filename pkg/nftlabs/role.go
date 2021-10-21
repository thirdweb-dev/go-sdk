package nftlabs

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

type Role string

const (
	AdminRole  Role = ""
	MinterRole      = "MINTER_ROLE"
	TransferRole = "TRANSFER_ROLE"
	PauserRole = "PAUSER_ROLE"
)

func RoleFromString(str string) (Role, error) {
	switch str {
	case "admin":
		return AdminRole, nil
	case "minter":
		return MinterRole, nil
	case "pauser":
		return PauserRole, nil
	case "transfer":
		return TransferRole, nil
	default:
		return "", errors.New(fmt.Sprintf("Role not found %v", str))
	}
}


type defaultModule interface {
	GrantRole(role Role, address string) error
	RevokeRole(role Role, address string) error
}

type defaultModuleImpl struct {
	sdk *Sdk
}

func (role *Role) getHash() [32]byte {
	var roleHash []byte
	if *role == AdminRole {
		roleHash = []byte("0x00000000000000000000000000000000")
	} else {
		roleHash = crypto.Keccak256([]byte(fmt.Sprintf("0x%v", role)))
	}
	r := [32]byte{}
	copy(r[:], roleHash)
	return r
}

func (manager *defaultModuleImpl) GrantRole(role Role, address string) error {
	var module interface{} = manager
	adr := common.HexToAddress(address)

	r := role.getHash()

	opts := &bind.TransactOpts{
		NoSend: false,
		From:   manager.sdk.getSignerAddress(),
		Signer: manager.sdk.getSigner(),
	}

	var transaction *types.Transaction
	if m, ok := module.(Currency); ok {
		if tx, err :=  m.getModule().GrantRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	} else if m, ok := module.(Nft); ok {
		if tx, err :=  m.getModule().GrantRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	} else if m, ok := module.(NftCollection); ok {
		if tx, err :=  m.getModule().GrantRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	} else if m, ok := module.(Currency); ok {
		if tx, err :=  m.getModule().GrantRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	}

	return waitForTx(manager.sdk.client, transaction.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
}

func (manager *defaultModuleImpl) RevokeRole(role Role, address string) error {
	var module interface{} = manager
	adr := common.HexToAddress(address)

	roleHash := crypto.Keccak256([]byte(fmt.Sprintf("0x%v", role)))
	r := [32]byte{}
	copy(r[:], roleHash)

	opts := &bind.TransactOpts{
		NoSend: false,
		From:   manager.sdk.getSignerAddress(),
		Signer: manager.sdk.getSigner(),
	}

	var transaction *types.Transaction
	if m, ok := module.(Currency); ok {
		if tx, err :=  m.getModule().RevokeRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	} else if m, ok := module.(Nft); ok {
		if tx, err :=  m.getModule().RevokeRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	} else if m, ok := module.(NftCollection); ok {
		if tx, err :=  m.getModule().RevokeRole(opts, r, adr); err != nil {
			return err
		} else {
			transaction = tx
		}
	} else if m, ok := module.(Currency); ok {
		if strings.ToLower(address) == strings.ToLower(manager.sdk.getSignerAddress().String()) {
			if tx, err :=  m.getModule().RenounceRole(opts, r, adr); err != nil {
				return err
			} else {
				transaction = tx
			}
		} else {
			if tx, err :=  m.getModule().RevokeRole(opts, r, adr); err != nil {
				 return err
			} else {
				 transaction = tx
			}
		}
	}

	return waitForTx(manager.sdk.client, transaction.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
}
