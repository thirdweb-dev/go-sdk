import fs from "fs";

const docPath = "./docs/doc.md"

const delimiters = {
  "## type [ERC1155]": {
    name: "erc1155.md",
    header: "## ERC1155\nThis interface is supported by the \`Edition\` contract.\n",
  },
  "## type [ERC721]": {
    name: "erc721.md",
    header: "\n## ERC721\nThis interface is supported by the \`NFTCollection\` and \`NFTDrop\` contracts.\n",
  },
  "## type [Edition]": {
    name: "edition.md",
    header: "\n## Edition\nYou can access this interface through the SDK with \`sdk.GetEdition(address)\`.\n",
  },
  "## type [EditionDrop]": {
    name: "edition_drop.md",
    header: "\n## Edition Drop\nYou can access this interface through the SDK with \`sdk.GetEditionDrop(address)\`.\n",
  },
  "## type [IpfsStorage]": {
    name: "storage.md",
    header: `## IPFS Storage\nYou can access this interface through the SDK with \`sdk.Storage\`.\n`,
  },
  "## type [NFTCollection]": {
    name: "nft_collection.md",
    header: "\n## NFT Collection\nYou can access this interface through the SDK with \`sdk.GetNFTCollection(address)\`.\n",
  },
  "## type [NFTDrop]": {
    name: "nft_drop.md",
    header: "\n## NFT Drop\nYou can access this interface through the SDK with \`sdk.GetNFTDrop(address)\`.\n",
  },
  "## type [ProviderHandler]": {
    name: "provider.md",
    header: "\n## Provider\n",
  },
  "## type [SDKOptions]": {
    name: "sdk.md",
    header: "\n## ThirdwebSDK\n",
  },
  "## type [WrappedToken]": {
    name: "finish.md",
    header: "\n## Wrapped Token\n",
  },
}

async function main() {
  const lines = fs.readFileSync(docPath).toString().split("\n");
  let file = ""
  let name = "start.md"
  let header = ""

  for (const line of lines) {
    let matched = false;

    for (const delimiter of Object.keys(delimiters)) {
      if (line.indexOf(delimiter) == 0 && !matched) {
        matched = true;
        fs.writeFile(`./docs/${name}`, header + file, (err) => {
          if (err) throw err;
        })

        const data = delimiters[delimiter]
        name = data.name
        header = data.header

        file = ""
      }
    }

    if (!matched) {
      file += "\n" + line
    }
  }


  fs.writeFile(`./docs/${name}`, file, (err) => {
    if (err) throw err;
  })

  const README = fs.readFileSync("./README.md").toString()
  fs.writeFile("./docs/index.md", README, (err) => {
    if (err) throw err;
  })
}

main()
