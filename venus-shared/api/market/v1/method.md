# Sample code of curl

```bash
# <Inputs> corresponding to the value of Inputs Tag of each API
curl http://<ip>:<port>/rpc/v1 -X POST -H "Content-Type: application/json"  -H "Authorization: Bearer <token>"  -d '{"method": "VENUS_MARKET.<method>", "params": <Inputs>, "id": 0}'
```
# Groups

* [Market](#market)
  * [ActorDelete](#actordelete)
  * [ActorExist](#actorexist)
  * [ActorList](#actorlist)
  * [ActorSectorSize](#actorsectorsize)
  * [ActorUpsert](#actorupsert)
  * [AddFsPieceStorage](#addfspiecestorage)
  * [AddS3PieceStorage](#adds3piecestorage)
  * [AssignDeals](#assigndeals)
  * [AssignUnPackedDeals](#assignunpackeddeals)
  * [DagstoreDestroyShard](#dagstoredestroyshard)
  * [DagstoreGC](#dagstoregc)
  * [DagstoreInitializeAll](#dagstoreinitializeall)
  * [DagstoreInitializeShard](#dagstoreinitializeshard)
  * [DagstoreInitializeStorage](#dagstoreinitializestorage)
  * [DagstoreListShards](#dagstorelistshards)
  * [DagstoreRecoverShard](#dagstorerecovershard)
  * [DealsBatchImportData](#dealsbatchimportdata)
  * [DealsConsiderOfflineRetrievalDeals](#dealsconsiderofflineretrievaldeals)
  * [DealsConsiderOfflineStorageDeals](#dealsconsiderofflinestoragedeals)
  * [DealsConsiderOnlineRetrievalDeals](#dealsconsideronlineretrievaldeals)
  * [DealsConsiderOnlineStorageDeals](#dealsconsideronlinestoragedeals)
  * [DealsConsiderUnverifiedStorageDeals](#dealsconsiderunverifiedstoragedeals)
  * [DealsConsiderVerifiedStorageDeals](#dealsconsiderverifiedstoragedeals)
  * [DealsImport](#dealsimport)
  * [DealsImportData](#dealsimportdata)
  * [DealsMaxProviderCollateralMultiplier](#dealsmaxprovidercollateralmultiplier)
  * [DealsMaxPublishFee](#dealsmaxpublishfee)
  * [DealsMaxStartDelay](#dealsmaxstartdelay)
  * [DealsPieceCidBlocklist](#dealspiececidblocklist)
  * [DealsPublishMsgPeriod](#dealspublishmsgperiod)
  * [DealsSetConsiderOfflineRetrievalDeals](#dealssetconsiderofflineretrievaldeals)
  * [DealsSetConsiderOfflineStorageDeals](#dealssetconsiderofflinestoragedeals)
  * [DealsSetConsiderOnlineRetrievalDeals](#dealssetconsideronlineretrievaldeals)
  * [DealsSetConsiderOnlineStorageDeals](#dealssetconsideronlinestoragedeals)
  * [DealsSetConsiderUnverifiedStorageDeals](#dealssetconsiderunverifiedstoragedeals)
  * [DealsSetConsiderVerifiedStorageDeals](#dealssetconsiderverifiedstoragedeals)
  * [DealsSetMaxProviderCollateralMultiplier](#dealssetmaxprovidercollateralmultiplier)
  * [DealsSetMaxPublishFee](#dealssetmaxpublishfee)
  * [DealsSetMaxStartDelay](#dealssetmaxstartdelay)
  * [DealsSetPieceCidBlocklist](#dealssetpiececidblocklist)
  * [DealsSetPublishMsgPeriod](#dealssetpublishmsgperiod)
  * [GetDeals](#getdeals)
  * [GetDirectDeal](#getdirectdeal)
  * [GetDirectDealByAllocationID](#getdirectdealbyallocationid)
  * [GetRetrievalDealStatistic](#getretrievaldealstatistic)
  * [GetStorageDealStatistic](#getstoragedealstatistic)
  * [GetUnPackedDeals](#getunpackeddeals)
  * [ID](#id)
  * [ImportDirectDeal](#importdirectdeal)
  * [IndexerAnnounceAllDeals](#indexerannouncealldeals)
  * [IndexerAnnounceDeal](#indexerannouncedeal)
  * [IndexerAnnounceDealRemoved](#indexerannouncedealremoved)
  * [IndexerAnnounceLatest](#indexerannouncelatest)
  * [IndexerAnnounceLatestHttp](#indexerannouncelatesthttp)
  * [IndexerListMultihashes](#indexerlistmultihashes)
  * [ListDirectDeals](#listdirectdeals)
  * [ListPieceStorageInfos](#listpiecestorageinfos)
  * [ListenMarketEvent](#listenmarketevent)
  * [MarkDealsAsPacking](#markdealsaspacking)
  * [MarketAddBalance](#marketaddbalance)
  * [MarketCancelDataTransfer](#marketcanceldatatransfer)
  * [MarketDataTransferPath](#marketdatatransferpath)
  * [MarketDataTransferUpdates](#marketdatatransferupdates)
  * [MarketGetAsk](#marketgetask)
  * [MarketGetDeal](#marketgetdeal)
  * [MarketGetDealUpdates](#marketgetdealupdates)
  * [MarketGetReserved](#marketgetreserved)
  * [MarketGetRetrievalAsk](#marketgetretrievalask)
  * [MarketGetRetrievalDeal](#marketgetretrievaldeal)
  * [MarketImportDealData](#marketimportdealdata)
  * [MarketImportPublishedDeal](#marketimportpublisheddeal)
  * [MarketListDataTransfers](#marketlistdatatransfers)
  * [MarketListDeals](#marketlistdeals)
  * [MarketListIncompleteDeals](#marketlistincompletedeals)
  * [MarketListRetrievalAsk](#marketlistretrievalask)
  * [MarketListRetrievalDeals](#marketlistretrievaldeals)
  * [MarketListStorageAsk](#marketliststorageask)
  * [MarketMaxBalanceAddFee](#marketmaxbalanceaddfee)
  * [MarketMaxDealsPerPublishMsg](#marketmaxdealsperpublishmsg)
  * [MarketPendingDeals](#marketpendingdeals)
  * [MarketPublishPendingDeals](#marketpublishpendingdeals)
  * [MarketReleaseFunds](#marketreleasefunds)
  * [MarketReserveFunds](#marketreservefunds)
  * [MarketRestartDataTransfer](#marketrestartdatatransfer)
  * [MarketSetAsk](#marketsetask)
  * [MarketSetDataTransferPath](#marketsetdatatransferpath)
  * [MarketSetMaxBalanceAddFee](#marketsetmaxbalanceaddfee)
  * [MarketSetMaxDealsPerPublishMsg](#marketsetmaxdealsperpublishmsg)
  * [MarketSetRetrievalAsk](#marketsetretrievalask)
  * [MarketWithdraw](#marketwithdraw)
  * [MessagerGetMessage](#messagergetmessage)
  * [MessagerPushMessage](#messagerpushmessage)
  * [MessagerWaitMessage](#messagerwaitmessage)
  * [NetAddrsListen](#netaddrslisten)
  * [PaychVoucherList](#paychvoucherlist)
  * [PiecesGetCIDInfo](#piecesgetcidinfo)
  * [PiecesGetPieceInfo](#piecesgetpieceinfo)
  * [PiecesListCidInfos](#pieceslistcidinfos)
  * [PiecesListPieces](#pieceslistpieces)
  * [ReleaseDeals](#releasedeals)
  * [ReleaseDirectDeals](#releasedirectdeals)
  * [RemovePieceStorage](#removepiecestorage)
  * [ResponseMarketEvent](#responsemarketevent)
  * [SectorGetExpectedSealDuration](#sectorgetexpectedsealduration)
  * [SectorSetExpectedSealDuration](#sectorsetexpectedsealduration)
  * [UpdateDealOnPacking](#updatedealonpacking)
  * [UpdateDealStatus](#updatedealstatus)
  * [UpdateDirectDealPayloadCID](#updatedirectdealpayloadcid)
  * [UpdateDirectDealState](#updatedirectdealstate)
  * [UpdateStorageDealPayloadSize](#updatestoragedealpayloadsize)
  * [UpdateStorageDealStatus](#updatestoragedealstatus)
  * [Version](#version)

## Market

### ActorDelete


Perms: admin

Inputs:
```json
[
  "f01234"
]
```

Response: `{}`

### ActorExist


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### ActorList


Perms: read

Inputs: `[]`

Response:
```json
[
  {
    "Addr": "f01234",
    "Account": "string value"
  }
]
```

### ActorSectorSize


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `34359738368`

### ActorUpsert


Perms: admin

Inputs:
```json
[
  {
    "Addr": "f01234",
    "Account": "string value"
  }
]
```

Response: `true`

### AddFsPieceStorage


Perms: admin

Inputs:
```json
[
  "string value",
  "string value",
  true
]
```

Response: `{}`

### AddS3PieceStorage


Perms: admin

Inputs:
```json
[
  "string value",
  "string value",
  "string value",
  "string value",
  "string value",
  "string value",
  "string value",
  true
]
```

Response: `{}`

### AssignDeals


Perms: write

Inputs:
```json
[
  {
    "Miner": 1000,
    "Number": 9
  },
  34359738368,
  {
    "MaxPiece": 123,
    "MaxPieceSize": 42,
    "MinPiece": 123,
    "MinPieceSize": 42,
    "MinUsedSpace": 42,
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "SectorExpiration": 10101
  }
]
```

Response:
```json
[
  {
    "DealID": 5432,
    "PublishCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "AllocationID": 0,
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1032,
    "Client": "f01234",
    "Provider": "f01234",
    "Offset": 1032,
    "Length": 1032,
    "PayloadSize": 42,
    "StartEpoch": 10101,
    "EndEpoch": 10101
  }
]
```

### AssignUnPackedDeals


Perms: write

Inputs:
```json
[
  {
    "Miner": 1000,
    "Number": 9
  },
  34359738368,
  {
    "MaxPiece": 123,
    "MaxPieceSize": 42,
    "MinPiece": 123,
    "MinPieceSize": 42,
    "MinUsedSpace": 42,
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "SectorExpiration": 10101
  }
]
```

Response:
```json
[
  {
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1032,
    "VerifiedDeal": true,
    "Client": "f01234",
    "Provider": "f01234",
    "Label": "",
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "StoragePricePerEpoch": "0",
    "ProviderCollateral": "0",
    "ClientCollateral": "0",
    "Offset": 1032,
    "Length": 1032,
    "PayloadSize": 42,
    "DealID": 5432,
    "TotalStorageFee": "0",
    "FastRetrieval": true,
    "PublishCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    }
  }
]
```

### DagstoreDestroyShard
DagstoreDestroyShard destroy shard by key


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response: `{}`

### DagstoreGC
DagstoreGC runs garbage collection on the DAG store.


Perms: admin

Inputs: `[]`

Response:
```json
[
  {
    "Key": "string value",
    "Success": true,
    "Error": "string value"
  }
]
```

### DagstoreInitializeAll
DagstoreInitializeAll initializes all uninitialized shards in bulk,
according to the policy passed in the parameters.

It is recommended to set a maximum concurrency to avoid extreme
IO pressure if the storage subsystem has a large amount of deals.

It returns a stream of events to report progress.


Perms: admin

Inputs:
```json
[
  {
    "MaxConcurrency": 123,
    "IncludeSealed": true
  }
]
```

Response:
```json
{
  "Key": "string value",
  "Event": "string value",
  "Success": true,
  "Error": "string value",
  "Total": 123,
  "Current": 123
}
```

### DagstoreInitializeShard
DagstoreInitializeShard initializes an uninitialized shard.

Initialization consists of fetching the shard's data (deal payload) from
the storage subsystem, generating an index, and persisting the index
to facilitate later retrievals, and/or to publish to external sources.

This operation is intended to complement the initial migration. The
migration registers a shard for every unique piece CID, with lazy
initialization. Thus, shards are not initialized immediately to avoid
IO activity competing with proving. Instead, shard are initialized
when first accessed. This method forces the initialization of a shard by
accessing it and immediately releasing it. This is useful to warm up the
cache to facilitate subsequent retrievals, and to generate the indexes
to publish them externally.

This operation fails if the shard is not in ShardStateNew state.
It blocks until initialization finishes.


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response: `{}`

### DagstoreInitializeStorage
DagstoreInitializeStorage initializes all pieces in specify storage


Perms: admin

Inputs:
```json
[
  "string value",
  {
    "MaxConcurrency": 123,
    "IncludeSealed": true
  }
]
```

Response:
```json
{
  "Key": "string value",
  "Event": "string value",
  "Success": true,
  "Error": "string value",
  "Total": 123,
  "Current": 123
}
```

### DagstoreListShards
DagstoreListShards returns information about all shards known to the
DAG store. Only available on nodes running the markets subsystem.


Perms: admin

Inputs: `[]`

Response:
```json
[
  {
    "Key": "string value",
    "State": "string value",
    "Error": "string value"
  }
]
```

### DagstoreRecoverShard
DagstoreRecoverShard attempts to recover a failed shard.

This operation fails if the shard is not in ShardStateErrored state.
It blocks until recovery finishes. If recovery failed, it returns the
error.


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response: `{}`

### DealsBatchImportData


Perms: admin

Inputs:
```json
[
  {
    "Refs": [
      {
        "ProposalCID": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "UUID": "07070707-0707-0707-0707-070707070707",
        "File": "string value"
      }
    ],
    "SkipCommP": true
  }
]
```

Response:
```json
[
  {
    "Target": "string value",
    "Message": "string value"
  }
]
```

### DealsConsiderOfflineRetrievalDeals


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### DealsConsiderOfflineStorageDeals


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### DealsConsiderOnlineRetrievalDeals


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### DealsConsiderOnlineStorageDeals


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### DealsConsiderUnverifiedStorageDeals


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### DealsConsiderVerifiedStorageDeals


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `true`

### DealsImport


Perms: admin

Inputs:
```json
[
  [
    {
      "ID": "07070707-0707-0707-0707-070707070707",
      "Proposal": {
        "PieceCID": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "PieceSize": 1032,
        "VerifiedDeal": true,
        "Client": "f01234",
        "Provider": "f01234",
        "Label": "",
        "StartEpoch": 10101,
        "EndEpoch": 10101,
        "StoragePricePerEpoch": "0",
        "ProviderCollateral": "0",
        "ClientCollateral": "0"
      },
      "ClientSignature": {
        "Type": 2,
        "Data": "Ynl0ZSBhcnJheQ=="
      },
      "ProposalCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "AddFundsCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PublishCid": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "Miner": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Client": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "State": 42,
      "PiecePath": "/some/path",
      "PayloadSize": 42,
      "MetadataPath": "/some/path",
      "SlashEpoch": 10101,
      "FastRetrieval": true,
      "Message": "string value",
      "FundsReserved": "0",
      "Ref": {
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
      "AvailableForRetrieval": true,
      "DealID": 5432,
      "CreationTime": "0001-01-01T00:00:00Z",
      "TransferChannelId": {
        "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
        "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
        "ID": 3
      },
      "SectorNumber": 9,
      "Offset": 1032,
      "PieceStatus": "Undefine",
      "InboundCAR": "string value",
      "CreatedAt": 42,
      "UpdatedAt": 42
    }
  ]
]
```

Response: `{}`

### DealsImportData


Perms: admin

Inputs:
```json
[
  {
    "ProposalCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "UUID": "07070707-0707-0707-0707-070707070707",
    "File": "string value"
  },
  true
]
```

Response: `{}`

### DealsMaxProviderCollateralMultiplier


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `42`

### DealsMaxPublishFee


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `"0 FIL"`

### DealsMaxStartDelay


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `60000000000`

### DealsPieceCidBlocklist


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

### DealsPublishMsgPeriod


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `60000000000`

### DealsSetConsiderOfflineRetrievalDeals


Perms: write

Inputs:
```json
[
  "f01234",
  true
]
```

Response: `{}`

### DealsSetConsiderOfflineStorageDeals


Perms: write

Inputs:
```json
[
  "f01234",
  true
]
```

Response: `{}`

### DealsSetConsiderOnlineRetrievalDeals


Perms: write

Inputs:
```json
[
  "f01234",
  true
]
```

Response: `{}`

### DealsSetConsiderOnlineStorageDeals


Perms: write

Inputs:
```json
[
  "f01234",
  true
]
```

Response: `{}`

### DealsSetConsiderUnverifiedStorageDeals


Perms: write

Inputs:
```json
[
  "f01234",
  true
]
```

Response: `{}`

### DealsSetConsiderVerifiedStorageDeals


Perms: write

Inputs:
```json
[
  "f01234",
  true
]
```

Response: `{}`

### DealsSetMaxProviderCollateralMultiplier


Perms: write

Inputs:
```json
[
  "f01234",
  42
]
```

Response: `{}`

### DealsSetMaxPublishFee


Perms: write

Inputs:
```json
[
  "f01234",
  "0 FIL"
]
```

Response: `{}`

### DealsSetMaxStartDelay


Perms: write

Inputs:
```json
[
  "f01234",
  60000000000
]
```

Response: `{}`

### DealsSetPieceCidBlocklist


Perms: write

Inputs:
```json
[
  "f01234",
  [
    {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    }
  ]
]
```

Response: `{}`

### DealsSetPublishMsgPeriod


Perms: write

Inputs:
```json
[
  "f01234",
  60000000000
]
```

Response: `{}`

### GetDeals


Perms: read

Inputs:
```json
[
  "f01234",
  123,
  123
]
```

Response:
```json
[
  {
    "DealID": 5432,
    "SectorID": 9,
    "Offset": 1032,
    "Length": 1032,
    "Proposal": {
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1032,
      "VerifiedDeal": true,
      "Client": "f01234",
      "Provider": "f01234",
      "Label": "",
      "StartEpoch": 10101,
      "EndEpoch": 10101,
      "StoragePricePerEpoch": "0",
      "ProviderCollateral": "0",
      "ClientCollateral": "0"
    },
    "ClientSignature": {
      "Type": 2,
      "Data": "Ynl0ZSBhcnJheQ=="
    },
    "TransferType": "string value",
    "Root": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PublishCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "FastRetrieval": true,
    "Status": "Undefine"
  }
]
```

### GetDirectDeal


Perms: read

Inputs:
```json
[
  "07070707-0707-0707-0707-070707070707"
]
```

Response:
```json
{
  "ID": "07070707-0707-0707-0707-070707070707",
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PieceSize": 1032,
  "Client": "f01234",
  "Provider": "f01234",
  "PayloadSize": 42,
  "PayloadCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "State": 1,
  "AllocationID": 42,
  "ClaimID": 42,
  "SectorID": 9,
  "Offset": 1032,
  "Length": 1032,
  "StartEpoch": 10101,
  "EndEpoch": 10101,
  "Message": "string value",
  "CreatedAt": 42,
  "UpdatedAt": 42
}
```

### GetDirectDealByAllocationID


Perms: read

Inputs:
```json
[
  0
]
```

Response:
```json
{
  "ID": "07070707-0707-0707-0707-070707070707",
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PieceSize": 1032,
  "Client": "f01234",
  "Provider": "f01234",
  "PayloadSize": 42,
  "PayloadCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "State": 1,
  "AllocationID": 42,
  "ClaimID": 42,
  "SectorID": 9,
  "Offset": 1032,
  "Length": 1032,
  "StartEpoch": 10101,
  "EndEpoch": 10101,
  "Message": "string value",
  "CreatedAt": 42,
  "UpdatedAt": 42
}
```

### GetRetrievalDealStatistic
GetRetrievalDealStatistic get retrieval deal statistic information
todo address undefined is invalid, it is currently not possible to directly associate an order with a miner


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
{
  "DealsStatus": {
    "0": 9
  }
}
```

### GetStorageDealStatistic
GetStorageDealStatistic get storage deal statistic information
if set miner address to address.Undef, return all storage deal info


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
{
  "DealsStatus": {
    "42": 9
  }
}
```

### GetUnPackedDeals


Perms: read

Inputs:
```json
[
  "f01234",
  {
    "MaxPiece": 123,
    "MaxPieceSize": 42,
    "MinPiece": 123,
    "MinPieceSize": 42,
    "MinUsedSpace": 42,
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "SectorExpiration": 10101
  }
]
```

Response:
```json
[
  {
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1032,
    "VerifiedDeal": true,
    "Client": "f01234",
    "Provider": "f01234",
    "Label": "",
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "StoragePricePerEpoch": "0",
    "ProviderCollateral": "0",
    "ClientCollateral": "0",
    "Offset": 1032,
    "Length": 1032,
    "PayloadSize": 42,
    "DealID": 5432,
    "TotalStorageFee": "0",
    "FastRetrieval": true,
    "PublishCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    }
  }
]
```

### ID


Perms: read

Inputs: `[]`

Response: `"12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf"`

### ImportDirectDeal
ImportDirectDeal import direct deals


Perms: write

Inputs:
```json
[
  {
    "SkipCommP": true,
    "DealParams": [
      {
        "PayloadSize": 42,
        "PayloadCID": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "DealUUID": "07070707-0707-0707-0707-070707070707",
        "AllocationID": 42,
        "PieceCID": {
          "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
        },
        "Client": "f01234",
        "StartEpoch": 10101,
        "EndEpoch": 10101
      }
    ]
  }
]
```

Response: `{}`

### IndexerAnnounceAllDeals


Perms: admin

Inputs:
```json
[
  "f01234"
]
```

Response: `{}`

### IndexerAnnounceDeal


Perms: admin

Inputs:
```json
[
  "Ynl0ZSBhcnJheQ=="
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### IndexerAnnounceDealRemoved


Perms: admin

Inputs:
```json
[
  "Ynl0ZSBhcnJheQ=="
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### IndexerAnnounceLatest


Perms: admin

Inputs: `[]`

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### IndexerAnnounceLatestHttp


Perms: admin

Inputs:
```json
[
  [
    "string value"
  ]
]
```

Response:
```json
{
  "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
}
```

### IndexerListMultihashes


Perms: read

Inputs:
```json
[
  "Ynl0ZSBhcnJheQ=="
]
```

Response:
```json
[
  "Bw=="
]
```

### ListDirectDeals


Perms: read

Inputs:
```json
[
  {
    "Provider": "f01234",
    "Client": "f01234",
    "State": 1,
    "Offset": 123,
    "Limit": 123,
    "Asc": true
  }
]
```

Response:
```json
[
  {
    "ID": "07070707-0707-0707-0707-070707070707",
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1032,
    "Client": "f01234",
    "Provider": "f01234",
    "PayloadSize": 42,
    "PayloadCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "State": 1,
    "AllocationID": 42,
    "ClaimID": 42,
    "SectorID": 9,
    "Offset": 1032,
    "Length": 1032,
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "Message": "string value",
    "CreatedAt": 42,
    "UpdatedAt": 42
  }
]
```

### ListPieceStorageInfos


Perms: read

Inputs: `[]`

Response:
```json
{
  "FsStorage": [
    {
      "Path": "string value",
      "Name": "string value",
      "ReadOnly": true,
      "Status": {
        "Capacity": 9,
        "Available": 9,
        "Reserved": 9
      }
    }
  ],
  "S3Storage": [
    {
      "Name": "string value",
      "ReadOnly": true,
      "EndPoint": "string value",
      "Bucket": "string value",
      "SubDir": "string value",
      "Status": {
        "Capacity": 9,
        "Available": 9,
        "Reserved": 9
      }
    }
  ]
}
```

### ListenMarketEvent


Perms: read

Inputs:
```json
[
  {
    "Miner": "f01234"
  }
]
```

Response:
```json
{
  "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
  "Method": "string value",
  "Payload": "Ynl0ZSBhcnJheQ=="
}
```

### MarkDealsAsPacking


Perms: write

Inputs:
```json
[
  "f01234",
  [
    5432
  ]
]
```

Response: `{}`

### MarketAddBalance


Perms: sign

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

### MarketCancelDataTransfer
MarketCancelDataTransfer cancels a data transfer with the given transfer ID and other peer


Perms: admin

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### MarketDataTransferPath


Perms: admin

Inputs:
```json
[
  "f01234"
]
```

Response: `"string value"`

### MarketDataTransferUpdates


Perms: admin

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

### MarketGetAsk


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
{
  "Ask": {
    "Price": "0",
    "VerifiedPrice": "0",
    "MinPieceSize": 1032,
    "MaxPieceSize": 1032,
    "Miner": "f01234",
    "Timestamp": 10101,
    "Expiry": 10101,
    "SeqNo": 42
  },
  "Signature": {
    "Type": 2,
    "Data": "Ynl0ZSBhcnJheQ=="
  },
  "CreatedAt": 42,
  "UpdatedAt": 42
}
```

### MarketGetDeal


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
  "ID": "07070707-0707-0707-0707-070707070707",
  "Proposal": {
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1032,
    "VerifiedDeal": true,
    "Client": "f01234",
    "Provider": "f01234",
    "Label": "",
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "StoragePricePerEpoch": "0",
    "ProviderCollateral": "0",
    "ClientCollateral": "0"
  },
  "ClientSignature": {
    "Type": 2,
    "Data": "Ynl0ZSBhcnJheQ=="
  },
  "ProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "AddFundsCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PublishCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Miner": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Client": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "State": 42,
  "PiecePath": "/some/path",
  "PayloadSize": 42,
  "MetadataPath": "/some/path",
  "SlashEpoch": 10101,
  "FastRetrieval": true,
  "Message": "string value",
  "FundsReserved": "0",
  "Ref": {
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
  "AvailableForRetrieval": true,
  "DealID": 5432,
  "CreationTime": "0001-01-01T00:00:00Z",
  "TransferChannelId": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "SectorNumber": 9,
  "Offset": 1032,
  "PieceStatus": "Undefine",
  "InboundCAR": "string value",
  "CreatedAt": 42,
  "UpdatedAt": 42
}
```

### MarketGetDealUpdates


Perms: admin

Inputs: `[]`

Response:
```json
{
  "ID": "07070707-0707-0707-0707-070707070707",
  "Proposal": {
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PieceSize": 1032,
    "VerifiedDeal": true,
    "Client": "f01234",
    "Provider": "f01234",
    "Label": "",
    "StartEpoch": 10101,
    "EndEpoch": 10101,
    "StoragePricePerEpoch": "0",
    "ProviderCollateral": "0",
    "ClientCollateral": "0"
  },
  "ClientSignature": {
    "Type": 2,
    "Data": "Ynl0ZSBhcnJheQ=="
  },
  "ProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "AddFundsCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PublishCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Miner": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Client": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "State": 42,
  "PiecePath": "/some/path",
  "PayloadSize": 42,
  "MetadataPath": "/some/path",
  "SlashEpoch": 10101,
  "FastRetrieval": true,
  "Message": "string value",
  "FundsReserved": "0",
  "Ref": {
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
  "AvailableForRetrieval": true,
  "DealID": 5432,
  "CreationTime": "0001-01-01T00:00:00Z",
  "TransferChannelId": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "SectorNumber": 9,
  "Offset": 1032,
  "PieceStatus": "Undefine",
  "InboundCAR": "string value",
  "CreatedAt": 42,
  "UpdatedAt": 42
}
```

### MarketGetReserved


Perms: sign

Inputs:
```json
[
  "f01234"
]
```

Response: `"0"`

### MarketGetRetrievalAsk


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
{
  "PricePerByte": "0",
  "UnsealPrice": "0",
  "PaymentInterval": 42,
  "PaymentIntervalIncrease": 42
}
```

### MarketGetRetrievalDeal


Perms: read

Inputs:
```json
[
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  42
]
```

Response:
```json
{
  "PayloadCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ID": 5,
  "Selector": {
    "Node": null
  },
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PricePerByte": "0",
  "PaymentInterval": 42,
  "PaymentIntervalIncrease": 42,
  "UnsealPrice": "0",
  "StoreID": 42,
  "SelStorageProposalCid": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "ChannelID": {
    "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "ID": 3
  },
  "Status": 0,
  "Receiver": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "TotalSent": 42,
  "FundsReceived": "0",
  "Message": "string value",
  "CurrentInterval": 42,
  "LegacyProtocol": true,
  "CreatedAt": 42,
  "UpdatedAt": 42
}
```

### MarketImportDealData


Perms: admin

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "string value"
]
```

Response: `{}`

### MarketImportPublishedDeal


Perms: write

Inputs:
```json
[
  {
    "ID": "07070707-0707-0707-0707-070707070707",
    "Proposal": {
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1032,
      "VerifiedDeal": true,
      "Client": "f01234",
      "Provider": "f01234",
      "Label": "",
      "StartEpoch": 10101,
      "EndEpoch": 10101,
      "StoragePricePerEpoch": "0",
      "ProviderCollateral": "0",
      "ClientCollateral": "0"
    },
    "ClientSignature": {
      "Type": 2,
      "Data": "Ynl0ZSBhcnJheQ=="
    },
    "ProposalCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "AddFundsCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PublishCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Miner": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Client": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "State": 42,
    "PiecePath": "/some/path",
    "PayloadSize": 42,
    "MetadataPath": "/some/path",
    "SlashEpoch": 10101,
    "FastRetrieval": true,
    "Message": "string value",
    "FundsReserved": "0",
    "Ref": {
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
    "AvailableForRetrieval": true,
    "DealID": 5432,
    "CreationTime": "0001-01-01T00:00:00Z",
    "TransferChannelId": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "SectorNumber": 9,
    "Offset": 1032,
    "PieceStatus": "Undefine",
    "InboundCAR": "string value",
    "CreatedAt": 42,
    "UpdatedAt": 42
  }
]
```

Response: `{}`

### MarketListDataTransfers


Perms: admin

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

### MarketListDeals


Perms: read

Inputs:
```json
[
  [
    "f01234"
  ]
]
```

Response:
```json
[
  {
    "Proposal": {
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1032,
      "VerifiedDeal": true,
      "Client": "f01234",
      "Provider": "f01234",
      "Label": "",
      "StartEpoch": 10101,
      "EndEpoch": 10101,
      "StoragePricePerEpoch": "0",
      "ProviderCollateral": "0",
      "ClientCollateral": "0"
    },
    "State": {
      "SectorNumber": 9,
      "SectorStartEpoch": 10101,
      "LastUpdatedEpoch": 10101,
      "SlashEpoch": 10101
    }
  }
]
```

### MarketListIncompleteDeals


Perms: read

Inputs:
```json
[
  {
    "Miner": "f01234",
    "State": 12,
    "Client": "string value",
    "DiscardFailedDeal": true,
    "DealID": 5432,
    "PieceCID": "string value",
    "Offset": 123,
    "Limit": 123,
    "Asc": true
  }
]
```

Response:
```json
[
  {
    "ID": "07070707-0707-0707-0707-070707070707",
    "Proposal": {
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      },
      "PieceSize": 1032,
      "VerifiedDeal": true,
      "Client": "f01234",
      "Provider": "f01234",
      "Label": "",
      "StartEpoch": 10101,
      "EndEpoch": 10101,
      "StoragePricePerEpoch": "0",
      "ProviderCollateral": "0",
      "ClientCollateral": "0"
    },
    "ClientSignature": {
      "Type": 2,
      "Data": "Ynl0ZSBhcnJheQ=="
    },
    "ProposalCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "AddFundsCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PublishCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "Miner": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "Client": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "State": 42,
    "PiecePath": "/some/path",
    "PayloadSize": 42,
    "MetadataPath": "/some/path",
    "SlashEpoch": 10101,
    "FastRetrieval": true,
    "Message": "string value",
    "FundsReserved": "0",
    "Ref": {
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
    "AvailableForRetrieval": true,
    "DealID": 5432,
    "CreationTime": "0001-01-01T00:00:00Z",
    "TransferChannelId": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "SectorNumber": 9,
    "Offset": 1032,
    "PieceStatus": "Undefine",
    "InboundCAR": "string value",
    "CreatedAt": 42,
    "UpdatedAt": 42
  }
]
```

### MarketListRetrievalAsk


Perms: read

Inputs: `[]`

Response:
```json
[
  {
    "Miner": "f01234",
    "PricePerByte": "0",
    "UnsealPrice": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "CreatedAt": 42,
    "UpdatedAt": 42
  }
]
```

### MarketListRetrievalDeals


Perms: read

Inputs:
```json
[
  {
    "Receiver": "string value",
    "PayloadCID": "string value",
    "Status": 12,
    "DiscardFailedDeal": true,
    "Offset": 123,
    "Limit": 123
  }
]
```

Response:
```json
[
  {
    "PayloadCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "ID": 5,
    "Selector": {
      "Node": null
    },
    "PieceCID": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "PricePerByte": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42,
    "UnsealPrice": "0",
    "StoreID": 42,
    "SelStorageProposalCid": {
      "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
    },
    "ChannelID": {
      "Initiator": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "Responder": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
      "ID": 3
    },
    "Status": 0,
    "Receiver": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
    "TotalSent": 42,
    "FundsReceived": "0",
    "Message": "string value",
    "CurrentInterval": 42,
    "LegacyProtocol": true,
    "CreatedAt": 42,
    "UpdatedAt": 42
  }
]
```

### MarketListStorageAsk


Perms: read

Inputs: `[]`

Response:
```json
[
  {
    "Ask": {
      "Price": "0",
      "VerifiedPrice": "0",
      "MinPieceSize": 1032,
      "MaxPieceSize": 1032,
      "Miner": "f01234",
      "Timestamp": 10101,
      "Expiry": 10101,
      "SeqNo": 42
    },
    "Signature": {
      "Type": 2,
      "Data": "Ynl0ZSBhcnJheQ=="
    },
    "CreatedAt": 42,
    "UpdatedAt": 42
  }
]
```

### MarketMaxBalanceAddFee


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `"0 FIL"`

### MarketMaxDealsPerPublishMsg


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `42`

### MarketPendingDeals


Perms: write

Inputs: `[]`

Response:
```json
[
  {
    "Deals": [
      {
        "Proposal": {
          "PieceCID": {
            "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
          },
          "PieceSize": 1032,
          "VerifiedDeal": true,
          "Client": "f01234",
          "Provider": "f01234",
          "Label": "",
          "StartEpoch": 10101,
          "EndEpoch": 10101,
          "StoragePricePerEpoch": "0",
          "ProviderCollateral": "0",
          "ClientCollateral": "0"
        },
        "ClientSignature": {
          "Type": 2,
          "Data": "Ynl0ZSBhcnJheQ=="
        }
      }
    ],
    "PublishPeriodStart": "0001-01-01T00:00:00Z",
    "PublishPeriod": 60000000000
  }
]
```

### MarketPublishPendingDeals


Perms: admin

Inputs: `[]`

Response: `{}`

### MarketReleaseFunds


Perms: sign

Inputs:
```json
[
  "f01234",
  "0"
]
```

Response: `{}`

### MarketReserveFunds


Perms: sign

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

### MarketRestartDataTransfer
MarketRestartDataTransfer attempts to restart a data transfer with the given transfer ID and other peer


Perms: admin

Inputs:
```json
[
  3,
  "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  true
]
```

Response: `{}`

### MarketSetAsk


Perms: admin

Inputs:
```json
[
  "f01234",
  "0",
  "0",
  10101,
  1032,
  1032
]
```

Response: `{}`

### MarketSetDataTransferPath


Perms: admin

Inputs:
```json
[
  "f01234",
  "string value"
]
```

Response: `{}`

### MarketSetMaxBalanceAddFee


Perms: write

Inputs:
```json
[
  "f01234",
  "0 FIL"
]
```

Response: `{}`

### MarketSetMaxDealsPerPublishMsg


Perms: write

Inputs:
```json
[
  "f01234",
  42
]
```

Response: `{}`

### MarketSetRetrievalAsk


Perms: admin

Inputs:
```json
[
  "f01234",
  {
    "PricePerByte": "0",
    "UnsealPrice": "0",
    "PaymentInterval": 42,
    "PaymentIntervalIncrease": 42
  }
]
```

Response: `{}`

### MarketWithdraw


Perms: sign

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
messager


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

### NetAddrsListen


Perms: read

Inputs: `[]`

Response:
```json
{
  "ID": "12D3KooWGzxzKZYveHXtpG6AsrUJBcWxHBFS2HsEoGTxrMLvKXtf",
  "Addrs": [
    "/ip4/52.36.61.156/tcp/1347/p2p/12D3KooWFETiESTf1v4PGUvtnxMAcEFMzLZbJGg4tjWfGEimYior"
  ]
}
```

### PaychVoucherList
Paych


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response:
```json
[
  {
    "ChannelAddr": "f01234",
    "TimeLockMin": 10101,
    "TimeLockMax": 10101,
    "SecretHash": "Ynl0ZSBhcnJheQ==",
    "Extra": {
      "Actor": "f01234",
      "Method": 1,
      "Data": "Ynl0ZSBhcnJheQ=="
    },
    "Lane": 42,
    "Nonce": 42,
    "Amount": "0",
    "MinSettleHeight": 10101,
    "Merges": [
      {
        "Lane": 42,
        "Nonce": 42
      }
    ],
    "Signature": {
      "Type": 2,
      "Data": "Ynl0ZSBhcnJheQ=="
    }
  }
]
```

### PiecesGetCIDInfo


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
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "PieceBlockLocations": [
    {
      "RelOffset": 42,
      "BlockSize": 42,
      "PieceCID": {
        "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
      }
    }
  ]
}
```

### PiecesGetPieceInfo


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
  "PieceCID": {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  "Deals": [
    {
      "DealID": 5432,
      "SectorID": 9,
      "Offset": 1032,
      "Length": 1032
    }
  ]
}
```

### PiecesListCidInfos


Perms: read

Inputs: `[]`

Response:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

### PiecesListPieces


Perms: read

Inputs: `[]`

Response:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

### ReleaseDeals
ReleaseDeals is used to release the deals that have been assigned by AssignUnPackedDeals method.


Perms: write

Inputs:
```json
[
  "f01234",
  [
    5432
  ]
]
```

Response: `{}`

### ReleaseDirectDeals


Perms: write

Inputs:
```json
[
  "f01234",
  [
    0
  ]
]
```

Response: `{}`

### RemovePieceStorage


Perms: admin

Inputs:
```json
[
  "string value"
]
```

Response: `{}`

### ResponseMarketEvent
market event


Perms: read

Inputs:
```json
[
  {
    "Id": "e26f1e5c-47f7-4561-a11d-18fab6e748af",
    "Payload": "Ynl0ZSBhcnJheQ==",
    "Error": "string value"
  }
]
```

Response: `{}`

### SectorGetExpectedSealDuration
SectorGetExpectedSealDuration gets the time that a newly-created sector
waits for more deals before it starts sealing


Perms: read

Inputs:
```json
[
  "f01234"
]
```

Response: `60000000000`

### SectorSetExpectedSealDuration
SectorSetExpectedSealDuration sets the expected time for a sector to seal


Perms: write

Inputs:
```json
[
  "f01234",
  60000000000
]
```

Response: `{}`

### UpdateDealOnPacking


Perms: write

Inputs:
```json
[
  "f01234",
  5432,
  9,
  1032
]
```

Response: `{}`

### UpdateDealStatus


Perms: write

Inputs:
```json
[
  "f01234",
  5432,
  "Undefine",
  42
]
```

Response: `{}`

### UpdateDirectDealPayloadCID


Perms: write

Inputs:
```json
[
  "07070707-0707-0707-0707-070707070707",
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  }
]
```

Response: `{}`

### UpdateDirectDealState


Perms: write

Inputs:
```json
[
  "07070707-0707-0707-0707-070707070707",
  1
]
```

Response: `{}`

### UpdateStorageDealPayloadSize


Perms: write

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  42
]
```

Response: `{}`

### UpdateStorageDealStatus


Perms: write

Inputs:
```json
[
  {
    "/": "bafy2bzacea3wsdh6y3a36tb3skempjoxqpuyompjbmfeyf34fi3uy6uue42v4"
  },
  42,
  "Undefine"
]
```

Response: `{}`

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

