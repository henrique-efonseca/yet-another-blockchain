package commands

import (
	"fmt"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the blockchain",
	Run: func(cmd *cobra.Command, args []string) {
		bc := blockchain.NewBlockchain()
		defer bc.Close()
		fmt.Println("Blockchain initialized successfully")

		// Load configuration
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file: %v\n", err)
			return
		}

		// Connect to peers
		peers := viper.GetStringSlice("peers")
		connectToPeers(bc, peers)
	},
}

func connectToPeers(bc *blockchain.Blockchain, peers []string) {
	p2pNetwork := network.NewP2PNetwork(bc)
	for _, peer := range peers {
		err := p2pNetwork.AddPeer(peer)
		if err != nil {
			fmt.Printf("Failed to connect to peer %s: %v\n", peer, err)
		} else {
			fmt.Printf("Connected to peer %s\n", peer)
		}
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
}
