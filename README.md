# filecoin-deal-proofs-svc

```go
go run .
```

http://localhost:9518/deal?dealID=1612483201

verify-proof -dataCID "datacid1234" -dealID 1612483201 -endEpoch 2000 -pieceCID piececid1234 -proof "0x05dd140f8ccca4857dce139bbfcf57e1d8731f91aa7529c24cc6b5b4a167384a0x19aa71c6797c1b3ddf50fc9ce54e8fa79f136bc69e0483be50a9e06a402d78210x4a3e8f8384574243b6834631e566940592478003ef356aa4fb2a83c99224eadc0x309e2a46029eee602f5426808541a20719013ea1f2763ea58da51bc4f4cec623" -provider "fprovider1" -signedEpoch 50 -startEpoch 10
