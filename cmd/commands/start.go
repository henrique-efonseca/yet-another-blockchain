package commands

import (
	"log"
	"net/http"
	"strconv"

	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/api"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/network"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiPort, p2pPort int

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the blockchain node",
	Run: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		if apiPort == 0 {
			apiPort = viper.GetInt("network.apiPort")
		}
		if p2pPort == 0 {
			p2pPort = viper.GetInt("network.p2pPort")
		}

		log.Printf("Config file loaded successfully. API Port: %d, P2P Port: %d", apiPort, p2pPort)

		if apiPort == 0 || p2pPort == 0 {
			log.Fatalf("Ports not provided. Use --api-port and --p2p-port flags or set them in the config.yaml file.")
		}

		log.Printf("Server listening on API port %d and P2P port %d", apiPort, p2pPort)

		bc := blockchain.NewBlockchain()
		defer bc.Close()
		p2p := network.NewP2PNetwork(bc)
		network.GlobalNetwork = p2p

		go p2p.StartServer(p2pPort)
		go p2p.ConnectToBootstrapNodes()

		apiInstance := api.NewAPI(bc)
		router := api.NewRouter(apiInstance)

		log.Fatal(http.ListenAndServe(":"+strconv.Itoa(apiPort), router))
	},
}

func init() {
	startCmd.Flags().IntVarP(&apiPort, "api-port", "", 0, "Port to run the API server on")
	startCmd.Flags().IntVarP(&p2pPort, "p2p-port", "", 0, "Port to run the P2P server on")
	rootCmd.AddCommand(startCmd)
}
