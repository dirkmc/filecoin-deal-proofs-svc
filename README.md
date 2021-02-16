# Filecoin Oracle on Ethereum - `web oracle`

This is an experimental proof-of-concept web service, which monitors the state of the Filecoin blockchain.

It was built as part of the Protocol Labs remote hack week held from February 1st, 2020 to February 5th, 2020.

:warning: This repository is a very rough proof-of-concept. If you want to put this on production you will also need to run the [Filecoin Sentinel](https://github.com/filecoin-project/sentinel) service, which this service would use to query state on the Filecoin blockchain :warning:

## Architecture

The experimental Filecoin Oracle consists of two parts:

1. `smart contracts` - [Solidity smart contracts for Ethereum](https://github.com/nonsense/filecoin-oracle)

2. `web oracle` - A trusted web service which monitors the state of the Filecoin blockchain

---

The web oracle continuously monitors the Filecoin blockchain, once an hour processes the state for all deals, and produces a merkle tree root hash of the serialized data. This service is backed by the [Filecoin Sentinel](https://github.com/filecoin-project/sentinel).

Users are able to query data CIDs of interest on the web oracle and get a merkle inclusion proof with all the relevant data for the data CID at that point in time:`dataCid`, `pieceCid`, `dealId`, `provider`, `startEpoch`, `endEpoch`, `signedEpoch`

## Installation

```
go install ./...
```

## Usage

```
Usage of deal-proofs:
  -chainid int
        chain id; rinkeby == 4 (default 4)
  -endpoint string
        endpoint to an ethereum node (default "https://rinkeby.infura.io/v3/xxxxx")
  -manager string
        manager address for the oracle contract (default "0x3b8Fd7cE0f4841F1C23B67b20676886ac230Be64")
  -oracle string
        oracle contract address on ethereum (default "0xd4375467f6CfB0493b5e4AF0601B3a0f2e7D2FcA")
  -production
        run in production, and send tx to ethereum network
  -prvkey string
        private key of account
  -remotedb string
        remote database (DSN)
```
