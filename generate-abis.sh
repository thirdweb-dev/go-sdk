#!/bin/bash

abigen --alias contractURI=internalContractURI --pkg abi --abi ./internal/abi/json-abi/TokenERC20.json --out internal/abi/token_erc20_abi.go --type TokenERC20
abigen --alias contractURI=internalContractURI --pkg abi --abi ./internal/abi/json-abi/TokenERC721.json --out internal/abi/token_erc721_abi.go --type TokenERC721
abigen --alias contractURI=internalContractURI --pkg abi --abi ./internal/abi/json-abi/TokenERC1155.json --out internal/abi/token_erc1155_abi.go --type TokenERC1155
