import fs from "fs";

const BASE_DOC_URL = "https://docs.thirdweb.com/go/";

const contracts = [
  {
    name: "EditionDrop",
    files: ["edition_drop.md", "erc1155.md"],
    docName: "edition_drop",
  },
  {
    name: "Edition",
    files: ["edition.md", "erc1155.md"],
    docName: "edition",
  },
  {
    name: "NFTCollection",
    files: ["nft_collection.md", "erc721.md"],
    docName: "nft_collection",
  },
  {
    name: "NFTDrop",
    files: ["nft_drop.md", "erc721.md"],
    docName: "nft_drop",
  },
  {
    name: "Token",
    files: ["token.md", "erc20.md"],
    docName: "token",
  },
]

async function main() {
  const snippets = {};

  contracts.map((contractData) => {
    const contractName = contractData.name
    const contractReference = `${BASE_DOC_URL}${contractData.docName}`;

    const contract = {
      name: contractName,
      summary: "",
      description: "",
      methods: [],
      properties: [],
      reference: contractReference,
    };

    contractData.files.map((fileName) => {
      const functions = fs.readFileSync(`./docs/${fileName}`).toString().split(/### func \\\(\\.+\) \[.+\]/);
      functions.map((fn) => {
        const functionNames = fn.match(/(?<=func \(.+\) )\S+(?=\()/)
        const examples = fn.match(/(?<=#### Example\n\n\`\`\`\n)(\n|.)+(?=\`\`\`)/)

        if (examples?.length && functionNames?.length) {  
          const fnReference = `${contractReference}#func-${contractData.docName.replace("_", "")}-${functionNames[0].toLowerCase()}`;
          
          const fnData = {
            name: functionNames[0],
            summary: "",
            example: examples[0],
            reference: fnReference,
          }
  
          contract.methods.push(fnData)
        }
      })
    });

    snippets[contractName] = contract; 
  })

  fs.writeFile(`./docs/snippets.json`, JSON.stringify(snippets, null, 2), (err) => {
    if (err) throw err;
  });
}

main()
