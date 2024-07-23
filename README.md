# Yet Another Blockchain Framework (YABF)

Yet Another Blockchain Framework (YABF) is a modular and extensible descentralized blockchain framework designed to help developers build custom blockchains with ease. 
Developers can create and manage custom blockchains using YABF SDKs, CLI, APIs and configuration files.
YABF supports multiple consensus mechanisms, smart contracts, wallet integration, and a wide range of blockchain features, making it the perfect foundation for creating bespoke blockchain solutions.

##

## Features

- **Modular Architecture**: Plug-and-play functionality for consensus algorithms, storage solutions, and networking components.
- **Consensus Mechanisms**: Support for Proof of Work (PoW), Proof of Stake (PoS), and other consensus mechanisms.
- **Smart Contracts**: Built-in support for deploying and executing smart contracts.
- **Wallet Integration**: Create and manage wallets, including transaction functionalities.
- **Configurable via Files**: Customize blockchain features using configuration files.
- **API and SDK**: Comprehensive API and SDK for building applications on top of the blockchain.
- **Security Features**: Robust security measures to protect against common blockchain attacks and vulnerabilities.
- **Privacy and Confidentiality**: Privacy-preserving techniques like Zero-Knowledge Proofs (ZKPs) and confidential transactions.
- **Interoperability**: Cross-chain communication protocols for interoperability with other blockchains.
- **Governance**: On-chain governance mechanisms for protocol upgrades and decision-making.
- **Scalability Solutions**: Layer 2 solutions, sidechains, and sharding techniques.
- **Data Storage**: Integration with decentralized storage solutions like IPFS and Filecoin.
- **Tokenization**: Create custom tokens and digital assets on the blockchain.
- **Decentralized Applications (DApps)**: Build decentralized applications on top of the blockchain.
- **Governance and Upgrades**: On-chain governance mechanisms for protocol upgrades and decision-making.

<br>

## Use Cases

1. **Digital Identity Management**: Secure, decentralized digital identity management system.
2. **Decentralized Voting System**: Transparent and secure blockchain-based voting system.
3. **Real Estate Property Management**: Immutable records for property transactions and ownership.
4. **Supply Chain Finance**: Blockchain-based platform for supply chain financing.
5. **Decentralized Cloud Storage**: Rent and access decentralized data storage.
6. **Tokenized Loyalty Programs**: Blockchain-based loyalty program for businesses.
7. **Intellectual Property Protection**: Register and protect intellectual property rights.
8. **Decentralized Autonomous Organization (DAO)**: Create and manage DAOs with on-chain governance.
9. **Healthcare Data Management**: Secure and interoperable healthcare data management.
10. **Decentralized Energy Trading**: Peer-to-peer energy trading platform.

<br>

## Installation

### Prerequisites

- Go 1.16 or higher

### Clone the Repository

```bash
git clone https://github.com/henrique-efonseca/yet-another-blockchain-framework.git
cd yet-another-blockchain-framework
```

### Install Dependencies

```bash
go mod tidy
```

### Build the Project

```bash
go build
```

### Configuration

Create a configuration file (config.yaml) to customize the blockchain features.

```yaml
network:
  name: "CustomBlockchain"
  port: 8080
  peers: []
consensus:
  type: "PoW" # Options: PoW, PoS, PBFT
  difficulty: 3
block:
  maxSize: 1 # in MB
smartContracts:
  enabled: true
wallet:
  enabled: true
  balance: 1000
```

### Running the Blockchain

#### Initialize the Blockchain

```bash
./yet-another-blockchain-framework init
```
<br>

#### Start Multiple Blockchain Nodes

```bash
./yet-another-blockchain-framework start
```
```bash
./yet-another-blockchain-framework start --config=config2.yaml
```



<br>

## API Reference

The blockchain exposes a RESTful API for interacting with the blockchain. The API documentation can be found [here](api-reference.md).

Here are some examples of API endpoints:

- **Get Blockchain Info**: Get information about the blockchain.
```bash
curl http://localhost:8080/api/v1/blockchain/info
```

- **Get All Blocks**: Get all blocks in the blockchain.
```bash
curl http://localhost:8080/api/v1/blockchain/blocks
```

- **Mine Block**: Mine a new block.
```bash
curl -X POST http://localhost:8080/api/v1/blockchain/mine
```

- **Create Wallet**: Create a new wallet.
```bash
curl -X POST http://localhost:8080/api/v1/wallet/create
```

- **Create a Smart Contract**: Create a new smart contract.
```bash
curl -X POST http://localhost:8080/api/v1/smart-contracts/create -d '{"code": "function add(a, b) { return a + b; }"}'
```

<br>

## SDK Reference

There are multiple SDKs available for interacting with the blockchain.
You can find the SDK documentation [here](sdk-reference.md).
Here is an example of using the Go SDK:


Install the SDK:

```bash
go get github.com/henrique-efonseca/yet-another-blockchain-framework/sdk/go
```

Use the SDK in your Go application:

```go
package main

import (
  "fmt"
  "github.com/henrique-efonseca/yet-another-blockchain-framework/sdk/go"
)

func main() {
  client := sdk.NewClient("http://localhost:8080")
  
  // Get blockchain info
  info, err := client.GetBlockchainInfo()
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Println(info)
}
```

<br>








