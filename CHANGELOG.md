## v0.30.5
add TradeRequestType enum

## v0.30.4
adding account mode: internal and binance

## v0.30.3
rename field from AccountType to TestingType

## v0.30.2
fixing naming

## v0.30.1
tag versioning must have the following format: `vX.Y.Z` where `X` is the major version, `Y` is the minor version and `Z` is the patch version.

## v0.30
adding providerid, internalid fields to orders

## v0.28.0
feat: AllItems tells if it contains all records

## v0.27.1
feat: simple beats response

## v0.27.0
feat: MaxSpread to Deviation

## v0.26.0
- feat: add deals to tickerbeats response
## v0.25.1
- feat: add group_id and retcode to trade confirmation

## v0.25
- feat: add group_id to beats
- update protbuf version
## v0.24.5
- feat: queue interface
## v0.24.4
- feat: Return all Stop Loss or Take Profit changes
## v0.24.3
- feat: feat: TradeTransaction add internal_id
## v0.24.2
- feat: TestingStatisticsVariables
## v0.24.1
- feat: adding more fields
## v0.24.0
- feat: Testing Statistics Service
## v0.23.8
- feat: allowed_domain, disallowed_domain
## v0.23.7
- feat: renaming to crawler
## v0.23.6
- feat: add SCRAPE action
## v0.23.5
- feat: action enum
## v0.23.3
- fix: field type
## v0.23.2
- fix: field type
  
## v0.23.1
-feat: updating fields

## v0.23.0
- feat: enqueuer status
  
## v0.22.18
- feat: source id as ID

## v0.22.17
- feat: improve names and change api

## v0.22.16
- fix: server name

## v0.22.15
- feat: update ticker-beats protos
- feat: create ticker-scraper protos
 
## v0.22.14
- feat: add TimeGMT

## v0.22.13
- feat: Return all unsynced deals

## v0.22.12
- feat: close all open position

## v0.22.11
- feat: add TimeGMTOffset to trade transation

## v0.22.10
- feat: saving times and offsets 2.0

## v0.22.9
- feat: datetime as in64
- feat: TimeGMTOffset in all models

## v0.22.8
- feat: saving times and offsets

## v0.22.7
- feat: add TimeGMTOffset to the model

## v0.22.6
- feat: add TimeGMTOffset to the model

## v0.22.5
- feat: add deal entry field to transaction

## v0.22.4
- fix: enum values

## v0.22.3
- fix: expiration_time as int32

## v0.22.2
- fix: TickerBeatsResponse

## v0.22.1
- feat: improve names

## v0.22.0
- feat: ticker beats synch, try II

## v0.21.11
- feat: add server time to the request

## v0.21.10
- feat: add Direction field to Orders model

## v0.21.9
- feat: add TradeRequest on signal beats response

## v0.21.8
- feat: rollback SignalType

## v0.21.7
- feat: improving field name

## v0.21.6
- feat: created, updated, deleted as int64

## v0.21.5
- feat: add beats type enum

## v0.21.4
- feat: add beats type to the request

## v0.21.3
- feat: return orders and positions 

## v0.21.2
- feat: datetime as integer

## v0.21.1
- feat: add ID to the model

## v0.21.0
- feat: change on SignalResult

## v0.20.0
- feat: update transactions model

## v0.19.12
- feat: add magic number to positions

## v0.19.11
- feat: return a list of positions instead of deals

## v0.19.10
- fix: struct response name

## v0.19.9
- feat: only one ticker service

## v0.19.8
- fix: services signatures

## v0.19.7
- fix: names

## v0.19.6
- feat: implementation of TickerTrades

## v0.19.5
- feat: rolling back changes

## v0.19.4
- feat: plain int32 fields type

## v0.19.3
- feat: improve TickerBeats response

## v0.19.2
- fix: CreateTradeTransaction

## v0.19.1
fix: TransactionsService name

## v0.19.0
- feat: signal service

## v0.18.1
- feat: add server time field
## v0.18.0
- feat: update name to ticker

## v0.17.2
- feat: account_id in the request

## v0.17.1
- feat: server time of the position v2
  
## v0.17.0
- feat: server time of the position

## v0.16.2
- feat: additional fields

## v0.16.1
- fix: field name

## v0.16.0
- feat: rolling back to int64
- feat: creation_order

## v0.15.12
- feat: naming improvements

## v0.15.11
- fix: name on repeated fields
  
## v0.15.10
- fix: TradeRequestActions

## v0.15.9
- fix: change filed name to time_expiration

## v0.15.8
- fix: expiration should be int64

## v0.15.7
- feat: adding model ID 

## v0.15.6
- fix: cannot use type, changing to trade_type 2
  
## v0.15.5
- feat: better naming
  
## v0.15.4
- fix: cannot use type, changing to trade_type
- feat: changing to int64 as it could be as big as it
  
## v0.15.3
- fix: typo on TradeTransactionType
  
## v0.15.2
- feat: trade transation request

## v0.15.1
- fix: add account_id
  
## v0.15.0
- feat: trade transation request
  
## 0.14.0
- feat: Concurrent queue

## 0.11.0
- GRPC Health Check
