import fs from "fs";

const docPath = "./docs/doc.md"

// Delimiters are used to split the initially generated single doc file into multiple files
// We search for the structure `## type [delimiter]` and create a new file for each delimiter
const delimiters = {
  "ERC20": {
    name: "erc20.md",
    header: "ERC20",
  },
  "ERC721SignatureMinting": {
    name: "erc721_signature_minting.md",
    header: "ERC721 Signature Minting",
  },
  "ERC721": {
    name: "erc721.md",
    header: "ERC721",
  },
  "ERC1155SignatureMinting": {
    name: "erc1155_signature_minting.md",
    header: "ERC1155 Signature Minting",
  },
  "ERC1155": {
    name: "erc1155.md",
    header: "ERC1155",
  },
  "Edition": {
    name: "edition.md",
    header: "Edition",
  },
  "EditionDrop": {
    name: "edition_drop.md",
    header: "Edition Drop",
  },
  "IpfsStorage": {
    name: "storage.md",
    header: `IPFS Storage`,
  },
  "Multiwrap": {
    name: "multiwrap.md",
    header: "Multiwrap",
  },
  "NFTCollection": {
    name: "nft_collection.md",
    header: "NFT Collection",
  },
  "NFTDrop": {
    name: "nft_drop.md",
    header: "NFT Drop",
  },
  "ProviderHandler": {
    name: "provider.md",
    header: "Provider",
  },
  "SDKOptions": {
    name: "types.md",
    header: "Types",
  },
  "SmartContract": {
    name: "custom.md",
    header: "Custom Contracts",
  },
  "ThirdwebSDK": {
    name: "sdk.md",
    header: "ThirdwebSDK",
  },
  "Token": {
    name: "token.md",
    header: "Token",
  },
  "WrappedToken": {
    name: "finish.md",
    header: "",
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

      // Get the actual type delimiter to split the file by
      const typeDelimeter = `## type [${delimiter}]`;

      if (line.indexOf(typeDelimeter) == 0 && !matched) {
        matched = true;

        // Do basic formatting on doc output to make them nicer
        file = file.replaceAll("    // contains filtered or unexported fields\n", "").replaceAll("{\n}", "{}").replaceAll("\\,", ",")
        file = `\n## ${header}${file}`
        
        fs.writeFile(`./docs/${name}`, file, (err) => {
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
