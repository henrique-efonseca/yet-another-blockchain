package commands

import (
	"fmt"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/network"
	"github.com/spf13/cobra"
)

var peerAddress string

var addPeerCmd = &cobra.Command{
	Use:   "add-peer",
	Short: "Add a new peer to the network",
	Run: func(cmd *cobra.Command, args []string) {
		if peerAddress == "" {
			fmt.Println("Please provide a peer address using the --address flag.")
			return
		}
		err := network.GlobalNetwork.AddPeer(peerAddress)
		if err != nil {
			fmt.Printf("Failed to add peer: %v\n", err)
		} else {
			fmt.Printf("Peer %s added successfully\n", peerAddress)
		}
	},
}

func init() {
	addPeerCmd.Flags().StringVar(&peerAddress, "address", "", "Address of the peer to add")
	rootCmd.AddCommand(addPeerCmd)
}
