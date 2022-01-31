#!/bin/bash

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/Coin.json   --out internal/abi/coin_abi.go --type Coin

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/Coin.json   --out internal/abi/currency_abi.go --type Currency

abigen --pkg abi --abi ../nftlabs-protocols/abi/ERC1155.json --out internal/abi/erc1155_abi.go --type ERC1155

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/ERC165.json  --out internal/abi/erc165_abi.go --type ERC165

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/ERC20.json   --out internal/abi/erc20_abi.go --type ERC20

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/ERC721.json  --out internal/abi/erc721_abi.go --type ERC721

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/Market.json  --out internal/abi/market_abi.go --type Market 

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/SignatureMint721.json --out internal/abi/nft_abi.go --type NFT

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/NFTCollection.json --out internal/abi/nft_collection_abi.go --type NFTCollection

abigen --alias contractURI=internalContractURI --pkg abi --abi ../nftlabs-protocols/abi/Pack.json --out internal/abi/pack_abi.go --type Pack
