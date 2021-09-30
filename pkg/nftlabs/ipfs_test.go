package nftlabs

import (
	"testing"
)

func TestReplaceGatewayWithSlash(t *testing.T) {
	ipfsUrl := "ipfs://bafkreiebhfacq6ccsuywmmc44ictgglr33zipvwv3lhg3rzs7wnkwlro4e"
	gatewayUrl := "https://cloudflare-ipfs.com/ipfs/"

	expected := "https://cloudflare-ipfs.com/ipfs/bafkreiebhfacq6ccsuywmmc44ictgglr33zipvwv3lhg3rzs7wnkwlro4e"
	result := replaceIpfsPrefixWithGateway(ipfsUrl, gatewayUrl)
	if result != expected {
		t.Fatalf("Got %v expected %v", result, expected)
	}
}

func TestReplaceGatewayWithoutSlash(t *testing.T) {
	ipfsUrl := "ipfs://bafkreiebhfacq6ccsuywmmc44ictgglr33zipvwv3lhg3rzs7wnkwlro4e"
	gatewayUrl := "https://cloudflare-ipfs.com/ipfs"

	expected := "https://cloudflare-ipfs.com/ipfs/bafkreiebhfacq6ccsuywmmc44ictgglr33zipvwv3lhg3rzs7wnkwlro4e"
	result := replaceIpfsPrefixWithGateway(ipfsUrl, gatewayUrl)
	if result != expected {
		t.Fatalf("Got %v expected %v", result, expected)
	}
}
