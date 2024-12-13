<!--https://github.com/darsaveli/Readme-Markdown-Syntax-->

### Collaborators:

* **[Lochlann O Neill](https://github.com/lochlannoneill)**

# TCP Server and Client in Go

This project demonstrates a simple TCP server and client implementation in Go. The server listens for incoming connections and handles messages from multiple clients concurrently. The client connects to the server and exchanges messages.

## Features

- **TCP Server**:
  - Listens on port `8080`.
  - Accepts multiple client connections concurrently using goroutines.
  - Echoes received messages back to the client.

- **TCP Client**:
  - Connects to the server at `localhost:8080`.
  - Sends messages entered by the user to the server.
  - Displays the server's response.

## Requirements

- [Go](https://golang.org/) (version 1.16 or later)

## Setup

1. Clone the repository or copy the source code.
2. Ensure Go is installed on your system.

## Running the Server

1. Navigate to the directory containing the `server.go` file.
2. Run the server using the following command:

   ```bash
   go run server.go
   ```

3. The server will start listening on port `8080` and display messages when clients connect or send data.

## Running the Client

1. Open a new terminal.
2. Navigate to the directory containing the `client.go` file.
3. Run the client using the following command:

   ```bash
   go run client.go
   ```

4. Enter messages in the client terminal. The server's responses will be displayed.

## How It Works

### Server

1. The server initializes a TCP listener on port `8080`.
2. It accepts incoming connections in a loop and spawns a goroutine to handle each client.
3. Each goroutine reads messages from the client, displays them, and echoes them back.

### Client

1. The client establishes a connection to the server at `localhost:8080`.
2. It reads user input from the terminal and sends it to the server.
3. The client displays the server's response to the user.

## Example

1. Start the server:

   ```bash
   Server is listening on port 8080...
   New client connected
   Message received: Hello, server!
   ```

2. Run the client and interact:

   ```bash
   Connected to server
   Enter message: Hello, server!
   Server reply: Message received: Hello, server!
   ```

## Notes

- Ensure the server is running before starting the client.
- The client and server communicate using newline-terminated (`\n`) messages.

## License

This project is licensed under the MIT License. Feel free to use, modify, and distribute it as needed.
