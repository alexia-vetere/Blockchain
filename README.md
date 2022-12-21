## Blockchain</br></br>
---
This repository contains the code for a blockchain application developed in Golang. It consists of three main components:

blockchain_server: This directory contains the code for the blockchain server, which handles the creation and management of blocks in the blockchain.

wallet_server: This directory contains the code for the wallet server, which handles the creation and management of wallets for the users of the application.

cmd: This directory contains the code for the command-line interface (CLI) for the application. It allows users to interact with the blockchain and wallet servers through a terminal.

 <h4>Structure of the blockchain proyect</h4>
 ![image](https://user-images.githubusercontent.com/64023919/177795941-65e41002-8f28-4175-bd39-a9eac569bad6.png)

To run the application, you will need to have Go (Golang) installed on your system.

To install the necessary dependencies, navigate to the root directory of the repository and run the following command:

´´´
go mod download
´´´

To start the blockchain server, navigate to the blockchain_server directory and run the following command:

´´´
go run main.go
´´´
To start the wallet server, navigate to the wallet_server directory and run the following command:

´´´
go run main.go
´´´
To start the CLI, navigate to the cmd directory and run the following command:

´´´
go run main.go
´´´
The CLI will prompt you to enter a command to interact with the blockchain and wallet servers. 
Some available commands are createblock, getbalance, and send.

I hope you find this application and its source code helpful in learning about blockchain technology:

ㅤ➢ Creation of the Block, Blockchain, PoW and Mining </br>
ㅤ➢ Development of Wallet </br>
ㅤ➢ Development of the Blockchain Server API</br>
ㅤ➢ Creation the structure of the Blockchain Network </br>
ㅤ➢ Synchronizing transactions </br>

