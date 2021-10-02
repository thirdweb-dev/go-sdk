package nftlabs

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type Role string

const (
	AdminRole  Role = "PROTOCOL_ADMIN"
	MinterRole      = "MINTER_ROLE"
)

func RoleFromString(str string) (Role, error) {
	switch str {
	case "admin":
		return AdminRole, nil
	case "minter":
		return MinterRole, nil
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

func (manager *defaultModuleImpl) GrantRole(role Role, address string) error {
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
	}

	return waitForTx(manager.sdk.client, transaction.Hash(), txWaitTimeBetweenAttempts, txMaxAttempts)
}
