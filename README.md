# angryfern-backend

## Stack

- golang
- gin
- postgres

## Getting Started

### Prerequisites

- 💻
- golang 1.21 or [later](https://go.dev)
- docker

### Installation

1. Clone this repo
2. Copy `.env.template` and paste it in the same directory as `.env` and fill in the values.
3. Download dependencies by `go mod download`

### Running

1. Run `docker-compose up -d`
2. Run `go run ./cmd/.` or `make dev`
