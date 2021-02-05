# filecoin-deal-proofs-svc

```go
go run .
```

http://localhost:9518/deal?dealID=1612483201

verify-proof -dataCID "datacid1234" -dealID 1612483201 -endEpoch 2000 -pieceCID piececid1234 -proof "0x66f50d6484c9b96eee530786960d10a57d794f9163e4fef635cbfbdaee3965630x3821639c7915986e0bce6b5a97d535e4bcee57f320f9395283bb192f773f18da0x10f72897dba52b6440a58f8f843e8fbbc09ff44ffead2d2c6246b62e0f56c0270xdd061920d6106362260f488c8e5cbe7033a40133ff311e4ffdd150000ee5f27e" -provider "fprovider1" -signedEpoch 50 -startEpoch 10
