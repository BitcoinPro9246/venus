# Sample code of curl

```bash
# <Inputs> corresponding to the value of Inputs Tag of each API
curl http://<ip>:<port>/rpc/v0 -X POST -H "Content-Type: application/json"  -H "Authorization: Bearer <token>"  -d '{"method": "VENUS_MARKET_CLIENT.<method>", "params": <Inputs>, "id": 0}'
```
# Groups

* [MarketClient](#marketclient)
  * [ClientBatchDeal](#clientbatchdeal)
  * [ClientCalcCommP](#clientcalccommp)
  * [ClientCancelDataTransfer](#clientcanceldatatransfer)
  * [ClientCancelRetrievalDeal](#clientcancelretrievaldeal)
  * [ClientDataTransferUpdates](#clientdatatransferupdates)
  * [ClientDealPieceCID](#clientdealpiececid)
  * [ClientDealSize](#clientdealsize)
  * [ClientExport](#clientexport)
  * [ClientFindData](#clientfinddata)
  * [ClientGenCar](#clientgencar)
  * [ClientGetDealInfo](#clientgetdealinfo)
  * [ClientGetDealStatus](#clientgetdealstatus)
  * [ClientGetDealUpdates](#clientgetdealupdates)
  * [ClientGetRetrievalUpdates](#clientgetretrievalupdates)
  * [ClientGetVerifiedDealDistribution](#clientgetverifieddealdistribution)
  * [ClientHasLocal](#clienthaslocal)
  * [ClientImport](#clientimport)
  * [ClientListDataTransfers](#clientlistdatatransfers)
  * [ClientListDeals](#clientlistdeals)
  * [ClientListImports](#clientlistimports)
  * [ClientListOfflineDeals](#clientlistofflinedeals)
  * [ClientListRetrievals](#clientlistretrievals)
  * [ClientMinerQueryOffer](#clientminerqueryoffer)
  * [ClientQueryAsk](#clientqueryask)
  * [ClientRemoveImport](#clientremoveimport)
  * [ClientRestartDataTransfer](#clientrestartdatatransfer)
  * [ClientRetrieve](#clientretrieve)
  * [ClientRetrieveTryRestartInsufficientFunds](#clientretrievetryrestartinsufficientfunds)
  * [ClientRetrieveWait](#clientretrievewait)
  * [ClientStartDeal](#clientstartdeal)
  * [ClientStatelessDeal](#clientstatelessdeal)
  * [DefaultAddress](#defaultaddress)
  * [MarketAddBalance](#marketaddbalance)
  * [MarketGetReserved](#marketgetreserved)
  * [MarketReleaseFunds](#marketreleasefunds)
  * [MarketReserveFunds](#marketreservefunds)
  * [MarketWithdraw](#marketwithdraw)
  * [MessagerGetMessage](#messagergetmessage)
  * [MessagerPushMessage](#messagerpushmessage)
  * [MessagerWaitMessage](#messagerwaitmessage)
  * [Version](#version)

## MarketClient

### ClientBatchDeal
ClientBatchDeal proposes deals with a miner


Perms: write

Inputs:
```json
[
  [
    {
      "Data": {
        "TransferType": "string value",
        "Root": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "PieceCid": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "PieceSize": 1024,
        "RawBlockSize": 42
      },
      "Wallet": "f01234",
      "Miner": "f01234",
      "EpochPrice": "0",
      "MinBlocksDuration": 42,
      "ProviderCollateral": "0",
      "DealStartEpoch": 10101,
      "FastRetrieval": true,
      "VerifiedDeal": true
    }
  ]
]
```

Response:
```json
{
  "Results": [
    {
      "ProposalCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "Message": "string value"
    }
  ]
}
```

### ClientCalcCommP
ClientCalcCommP calculates the CommP for a specified file


Perms: write

Inputs:
```json
[
  "string value"
]
```

Response:
```json
{
  "Root": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 1024
}
```

### ClientCancelDataTransfer
ClientCancelDataTransfer cancels a data transfer with the given transfer ID and other peer


Perms: write

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### ClientCancelRetrievalDeal
ClientCancelRetrievalDeal cancels an ongoing retrieval deal based on DealID


Perms: write

Inputs:
```json
[
  5
]
```

Response: `{}`

### ClientDataTransferUpdates


Perms: write

Inputs: `[]`

Response:
```json
{
  "TransferID": 3,
  "Status": 1,
  "BaseCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "IsInitiator": true,
  "IsSender": true,
  "Voucher": "string value",
  "Message": "string value",
  "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Transferred": 42,
  "Stages": {
    "Stages": [
      {
        "Name": "string value",
        "Description": "string value",
        "CreatedTime": "0001-01-01T00:00:00Z",
        "UpdatedTime": "0001-01-01T00:00:00Z",
        "Logs": [
          {
            "Log": "string value",
            "UpdatedTime": "0001-01-01T00:00:00Z"
          }
        ]
      }
    ]
  }
}
```

### ClientDealPieceCID
ClientDealPieceCID returns the Piece CID and data size for the specified CID


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "PayloadSize": 9,
  "PieceSize": 1032,
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
}
```

### ClientDealSize
ClientDealSize calculates real deal data size


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "PayloadSize": 9,
  "PieceSize": 1032
}
```

### ClientExport
ClientExport exports a file stored in the local filestore to a system file


Perms: admin

Inputs:
```json
[
  {
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "DAGs": [
      {
        "DataSelector": "/ipld/a/b/c",
        "ExportMerkleProof": true
      }
    ],
    "FromLocalCAR": "string value",
    "DealID": 5
  },
  {
    "Path": "string value",
    "IsCAR": true
  }
]
```

Response: `{}`

### ClientFindData
ClientFindData identifies peers that have a certain file, and returns QueryOffers (one per peer).


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
[
  {
    "Err": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Piece": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Size": 42,
    "MinPrice": "0",
    "UnsealPrice": "0",
    "PricePerByte": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "Miner": "f01234",
    "MinerPeer": {
      "Address": "f01234",
      "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      }
    }
  }
]
```

### ClientGenCar
ClientGenCar generates a CAR file for the specified file.


Perms: write

Inputs:
```json
[
  {
    "Path": "string value",
    "IsCAR": true
  },
  "string value"
]
```

Response: `{}`

### ClientGetDealInfo
ClientGetDealInfo returns the latest information about a given deal.


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "ProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "State": 42,
  "Message": "string value",
  "DealStages": {
    "Stages": [
      {
        "Name": "string value",
        "Description": "string value",
        "ExpectedDuration": "string value",
        "CreatedTime": "0001-01-01T00:00:00Z",
        "UpdatedTime": "0001-01-01T00:00:00Z",
        "Logs": [
          {
            "Log": "string value",
            "UpdatedTime": "0001-01-01T00:00:00Z"
          }
        ]
      }
    ]
  },
  "Provider": "f01234",
  "DataRef": {
    "TransferType": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1024,
    "RawBlockSize": 42
  },
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 42,
  "PricePerEpoch": "0",
  "Duration": 42,
  "DealID": 5432,
  "CreationTime": "0001-01-01T00:00:00Z",
  "Verified": true,
  "TransferChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "DataTransfer": {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    }
  }
}
```

### ClientGetDealStatus
ClientGetDealStatus returns status given a code


Perms: read

Inputs:
```json
[
  42
]
```

Response: `"string value"`

### ClientGetDealUpdates
ClientGetDealUpdates returns the status of updated deals


Perms: write

Inputs: `[]`

Response:
```json
{
  "ProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "State": 42,
  "Message": "string value",
  "DealStages": {
    "Stages": [
      {
        "Name": "string value",
        "Description": "string value",
        "ExpectedDuration": "string value",
        "CreatedTime": "0001-01-01T00:00:00Z",
        "UpdatedTime": "0001-01-01T00:00:00Z",
        "Logs": [
          {
            "Log": "string value",
            "UpdatedTime": "0001-01-01T00:00:00Z"
          }
        ]
      }
    ]
  },
  "Provider": "f01234",
  "DataRef": {
    "TransferType": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1024,
    "RawBlockSize": 42
  },
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 42,
  "PricePerEpoch": "0",
  "Duration": 42,
  "DealID": 5432,
  "CreationTime": "0001-01-01T00:00:00Z",
  "Verified": true,
  "TransferChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "DataTransfer": {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    }
  }
}
```

### ClientGetRetrievalUpdates
ClientGetRetrievalUpdates returns status of updated retrieval deals


Perms: write

Inputs: `[]`

Response:
```json
{
  "PayloadCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ID": 5,
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PricePerByte": "0",
  "UnsealPrice": "0",
  "Status": 0,
  "Message": "string value",
  "Provider": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "BytesReceived": 42,
  "BytesPaidFor": 42,
  "TotalPaid": "0",
  "TransferChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "DataTransfer": {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    }
  },
  "Event": 5
}
```

### ClientGetVerifiedDealDistribution


Perms: read

Inputs:
```json
[
  [
    "f01234"
  ],
  "f01234"
]
```

Response:
```json
{
  "ProvidersDistribution": [
    {
      "Provider": "f01234",
      "Total": 42,
      "Uniq": 42,
      "UniqPieces": {
        "string value": 42
      },
      "DuplicationPercentage": 12.3
    }
  ],
  "ReplicasDistribution": [
    {
      "Client": "f01234",
      "Total": 42,
      "Uniq": 42,
      "DuplicationPercentage": 12.3,
      "ReplicasPercentage": {
        "string value": 12.3
      },
      "ReplicasDistribution": [
        {
          "Provider": "f01234",
          "Total": 42,
          "Uniq": 42,
          "UniqPieces": {
            "string value": 42
          },
          "DuplicationPercentage": 12.3
        }
      ]
    }
  ]
}
```

### ClientHasLocal
ClientHasLocal indicates whether a certain CID is locally stored.


Perms: write

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `true`

### ClientImport
ClientImport imports file under the specified path into filestore.


Perms: admin

Inputs:
```json
[
  {
    "Path": "string value",
    "IsCAR": true
  }
]
```

Response:
```json
{
  "Root": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ImportID": 1234
}
```

### ClientListDataTransfers
ClientListTransfers returns the status of all ongoing transfers of data


Perms: write

Inputs: `[]`

Response:
```json
[
  {
    "TransferID": 3,
    "Status": 1,
    "BaseCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "IsInitiator": true,
    "IsSender": true,
    "Voucher": "string value",
    "Message": "string value",
    "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Transferred": 42,
    "Stages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    }
  }
]
```

### ClientListDeals
ClientListDeals returns information about the deals made by the local client.


Perms: write

Inputs: `[]`

Response:
```json
[
  {
    "ProposalCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "State": 42,
    "Message": "string value",
    "DealStages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "ExpectedDuration": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    },
    "Provider": "f01234",
    "DataRef": {
      "TransferType": "string value",
      "Root": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1024,
      "RawBlockSize": 42
    },
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Size": 42,
    "PricePerEpoch": "0",
    "Duration": 42,
    "DealID": 5432,
    "CreationTime": "0001-01-01T00:00:00Z",
    "Verified": true,
    "TransferChannelID": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "DataTransfer": {
      "TransferID": 3,
      "Status": 1,
      "BaseCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "IsInitiator": true,
      "IsSender": true,
      "Voucher": "string value",
      "Message": "string value",
      "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Transferred": 42,
      "Stages": {
        "Stages": [
          {
            "Name": "string value",
            "Description": "string value",
            "CreatedTime": "0001-01-01T00:00:00Z",
            "UpdatedTime": "0001-01-01T00:00:00Z",
            "Logs": [
              {
                "Log": "string value",
                "UpdatedTime": "0001-01-01T00:00:00Z"
              }
            ]
          }
        ]
      }
    }
  }
]
```

### ClientListImports
ClientUnimport removes references to the specified file from filestore
ClientUnimport(path string)


Perms: write

Inputs: `[]`

Response:
```json
[
  {
    "Key": 1234,
    "Err": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Source": "string value",
    "FilePath": "string value",
    "CARPath": "string value"
  }
]
```

### ClientListOfflineDeals


Perms: read

Inputs: `[]`

Response:
```json
[
  {
    "ProposalCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "State": 42,
    "Message": "string value",
    "DealStages": {
      "Stages": [
        {
          "Name": "string value",
          "Description": "string value",
          "ExpectedDuration": "string value",
          "CreatedTime": "0001-01-01T00:00:00Z",
          "UpdatedTime": "0001-01-01T00:00:00Z",
          "Logs": [
            {
              "Log": "string value",
              "UpdatedTime": "0001-01-01T00:00:00Z"
            }
          ]
        }
      ]
    },
    "Provider": "f01234",
    "DataRef": {
      "TransferType": "string value",
      "Root": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1024,
      "RawBlockSize": 42
    },
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Size": 42,
    "PricePerEpoch": "0",
    "Duration": 42,
    "DealID": 5432,
    "CreationTime": "0001-01-01T00:00:00Z",
    "Verified": true,
    "TransferChannelID": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "DataTransfer": {
      "TransferID": 3,
      "Status": 1,
      "BaseCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "IsInitiator": true,
      "IsSender": true,
      "Voucher": "string value",
      "Message": "string value",
      "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Transferred": 42,
      "Stages": {
        "Stages": [
          {
            "Name": "string value",
            "Description": "string value",
            "CreatedTime": "0001-01-01T00:00:00Z",
            "UpdatedTime": "0001-01-01T00:00:00Z",
            "Logs": [
              {
                "Log": "string value",
                "UpdatedTime": "0001-01-01T00:00:00Z"
              }
            ]
          }
        ]
      }
    }
  }
]
```

### ClientListRetrievals


Perms: write

Inputs: `[]`

Response:
```json
[
  {
    "PayloadCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "ID": 5,
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PricePerByte": "0",
    "UnsealPrice": "0",
    "Status": 0,
    "Message": "string value",
    "Provider": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "BytesReceived": 42,
    "BytesPaidFor": 42,
    "TotalPaid": "0",
    "TransferChannelID": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "DataTransfer": {
      "TransferID": 3,
      "Status": 1,
      "BaseCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "IsInitiator": true,
      "IsSender": true,
      "Voucher": "string value",
      "Message": "string value",
      "OtherPeer": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Transferred": 42,
      "Stages": {
        "Stages": [
          {
            "Name": "string value",
            "Description": "string value",
            "CreatedTime": "0001-01-01T00:00:00Z",
            "UpdatedTime": "0001-01-01T00:00:00Z",
            "Logs": [
              {
                "Log": "string value",
                "UpdatedTime": "0001-01-01T00:00:00Z"
              }
            ]
          }
        ]
      }
    },
    "Event": 5
  }
]
```

### ClientMinerQueryOffer
ClientMinerQueryOffer returns a QueryOffer for the specific miner and file.


Perms: read

Inputs:
```json
[
  "f01234",
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "Err": "string value",
  "Root": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Piece": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Size": 42,
  "MinPrice": "0",
  "UnsealPrice": "0",
  "PricePerByte": "0",
  "PaymentInterval": 42,
  "PaymentIntervalIncrease": 42,
  "Miner": "f01234",
  "MinerPeer": {
    "Address": "f01234",
    "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    }
  }
}
```

### ClientQueryAsk
ClientQueryAsk returns a signed StorageAsk from the specified miner.


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "f01234"
]
```

Response:
```json
{
  "Price": "0",
  "VerifiedPrice": "0",
  "MinPieceSize": 1032,
  "MaxPieceSize": 1032,
  "Miner": "f01234",
  "Timestamp": 10101,
  "Expiry": 10101,
  "SeqNo": 42
}
```

### ClientRemoveImport
ClientRemoveImport removes file import


Perms: admin

Inputs:
```json
[
  1234
]
```

Response: `{}`

### ClientRestartDataTransfer
ClientRestartDataTransfer attempts to restart a data transfer with the given transfer ID and other peer


Perms: write

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### ClientRetrieve
ClientRetrieve initiates the retrieval of a file, as specified in the order.


Perms: admin

Inputs:
```json
[
  {
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Piece": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "DataSelector": "/ipld/a/b/c",
    "Size": 42,
    "Total": "0",
    "UnsealPrice": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "Client": "f01234",
    "Miner": "f01234",
    "MinerPeer": {
      "Address": "f01234",
      "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      }
    },
    "RemoteStore": "102334ec-35a3-4b36-be9f-02883844503a"
  }
]
```

Response:
```json
{
  "DealID": 5
}
```

### ClientRetrieveTryRestartInsufficientFunds
ClientRetrieveTryRestartInsufficientFunds attempts to restart stalled retrievals on a given payment channel
which are stuck due to insufficient funds


Perms: write

Inputs:
```json
[
  "f01234"
]
```

Response: `{}`

### ClientRetrieveWait
ClientRetrieveWait waits for retrieval to be complete


Perms: admin

Inputs:
```json
[
  5
]
```

Response: `{}`

### ClientStartDeal
ClientStartDeal proposes a deal with a miner.


Perms: admin

Inputs:
```json
[
  {
    "Data": {
      "TransferType": "string value",
      "Root": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1024,
      "RawBlockSize": 42
    },
    "Wallet": "f01234",
    "Miner": "f01234",
    "EpochPrice": "0",
    "MinBlocksDuration": 42,
    "ProviderCollateral": "0",
    "DealStartEpoch": 10101,
    "FastRetrieval": true,
    "VerifiedDeal": true
  }
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### ClientStatelessDeal
ClientStatelessDeal fire-and-forget-proposes an offline deal to a miner without subsequent tracking.


Perms: write

Inputs:
```json
[
  {
    "Data": {
      "TransferType": "string value",
      "Root": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1024,
      "RawBlockSize": 42
    },
    "Wallet": "f01234",
    "Miner": "f01234",
    "EpochPrice": "0",
    "MinBlocksDuration": 42,
    "ProviderCollateral": "0",
    "DealStartEpoch": 10101,
    "FastRetrieval": true,
    "VerifiedDeal": true
  }
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### DefaultAddress


Perms: read

Inputs: `[]`

Response: `"f01234"`

### MarketAddBalance


Perms: write

Inputs:
```json
[
  "f01234",
  "f01234",
  "0"
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### MarketGetReserved


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `"0"`

### MarketReleaseFunds


Perms: write

Inputs:
```json
[
  "f01234",
  "0"
]
```

Response: `{}`

### MarketReserveFunds


Perms: write

Inputs:
```json
[
  "f01234",
  "f01234",
  "0"
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### MarketWithdraw


Perms: write

Inputs:
```json
[
  "f01234",
  "f01234",
  "0"
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### MessagerGetMessage


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "CID": {
    "/": "bafy2bzacebbpdegvr3i4cosewthysg5xkxpqfn2wfcz6mv2hmoktwbdxkax4s"
  },
  "Version": 42,
  "To": "f01234",
  "From": "f01234",
  "Nonce": 42,
  "Value": "0",
  "GasLimit": 9,
  "GasFeeCap": "0",
  "GasPremium": "0",
  "Method": 1,
  "Params": "Ynl0ZSBhcnJheQ=="
}
```

### MessagerPushMessage


Perms: write

Inputs:
```json
[
  {
    "CID": {
      "/": "bafy2bzacebbpdegvr3i4cosewthysg5xkxpqfn2wfcz6mv2hmoktwbdxkax4s"
    },
    "Version": 42,
    "To": "f01234",
    "From": "f01234",
    "Nonce": 42,
    "Value": "0",
    "GasLimit": 9,
    "GasFeeCap": "0",
    "GasPremium": "0",
    "Method": 1,
    "Params": "Ynl0ZSBhcnJheQ=="
  },
  {
    "MaxFee": "0",
    "GasOverEstimation": 12.3,
    "GasOverPremium": 12.3
  }
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### MessagerWaitMessage


Perms: read

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response:
```json
{
  "Message": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Receipt": {
    "ExitCode": 0,
    "Return": "Ynl0ZSBhcnJheQ==",
    "GasUsed": 9,
    "EventsRoot": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    }
  },
  "ReturnDec": {},
  "TipSet": [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    {
      "/": "bafy2bzacebp3shtrn43k7g3unredz7fxn4gj533d3o43tqn2p2ipxxhrvchve"
    }
  ],
  "Height": 10101
}
```

### Version
Version provides information about API provider


Perms: read

Inputs: `[]`

Response:
```json
{
  "Version": "string value",
  "APIVersion": 131840
}
```

