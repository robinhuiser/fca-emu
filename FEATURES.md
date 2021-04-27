# FCA feature matrix
## Entity

| Method | URL | Status |
|--- |--- |--- |
| GET |  `/cloud/v1/entity/{entityId}` | Implemented |
| GET |  `/cloud/v1/entity/{entityId}/profile` | Implemented |
| PUT | `/cloud/v1/entity/{entityId}/profile` | Open |
| POST | `/cloud/v1/entities` | Open |
| POST | `/cloud/v1/entities/search` | Open|

## Products

| Method | URL | Status |
|--- |--- |--- |
| GET |  `/cloud/v1/entity/{entityId}/accounts` | Implemented |

## Accounts

| Method | URL | Status |
|--- |--- |--- |
| GET |  `​/cloud/v1/entity/{entityId}/accounts` | Implemented |
| GET |  `​/cloud​/v1​/account​/{accountId}` | Implemented |
| PUT | `/cloud/v1/account/{accountId}` | Open |
| GET |  `​/cloud/v1/account/{accountId}/balances` | Implemented |
| GET |  `​/cloud/v1/account/{accountId}/details` | Implemented |
| POST | `/cloud/v1/accounts/search` | Open |

## Cards

| Method | URL | Status |
|--- |--- |--- |
| GET |  `/cloud/v1/account/{accountId}/cards` | Open |

## Transactions

| Method | URL | Status |
|--- |--- |--- |
| GET |  `/cloud/v1/account/{accountId}/transactions` | Implemented |
| GET | `/cloud/v1/account/{accountId}/transaction/{transactionId}` | Next release |
| GET | `/cloud/v1/account/{accountId}/transaction/{transactionId}/images` | Next release |
| GET | `/cloud/v1/account/{accountId}/image/{itemId}` | Next release |
| POST | `/cloud/v1/transactions/search` | Open |

## Statements

| Method | URL | Status |
|--- |--- |--- |
| GET |  `/cloud/v1/account/{accountId}/statements` | Open |

## Statement

| Method | URL | Status |
|--- |--- |--- |
| GET |  `/cloud/v1/account/{accountId}/statement/{statementId}` | Open |
| GET |  `​/cloud​/v1​/statement​/{statementId}` | Open |
| GET |  `​/cloud​/v1​/account​/{accountId}​/statement​/{statementId}​/item​/{itemId}` | Open |
| GET |  `​/cloud​/v1​/statement​/{statementId}​/item​/{itemId}` | Open |

## Exchange

| Method | URL | Status |
|--- |--- |--- |
| GET |  `​/cloud​/v1​/entity​/{entityId}​/exchanges` | Open |
| GET | `/cloud/v1/account/{accountId}/exchanges` | Open |
| POST | `/cloud/v1/exchange` | Open |
| DELETE | `/cloud/v1/exchange/{exchangeId}` | Open |
| GET | `/cloud/v1/exchange/{exchangeId}` | Open |
| PUT | `/cloud/v1/exchange/{exchangeId}` | Open |
