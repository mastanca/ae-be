# Accounting Notebook BE

## Usage

``` shell script
make run
```

## Test
```shell script
make run-tests
```

## Endpoints

#### `POST` /transactions

Commit a new transaction

Expected body:
```json
{
  "type": "debit",
  "amount": 500
}
```

Response:

Status: 201
```json
{
    "id": "985aa1e7-b9e0-4786-b69a-2825eb26e30b",
    "type": "debit",
    "amount": 500,
    "effective_date": "2020-11-08T16:04:28.219161-03:00"
}
```

Status: 409
```json
{
    "reason": "insufficient funds"
}
```

#### `GET` /transactions

Get all transactions

Response:

Status: 200
```json
[
    {
        "id": "985aa1e7-b9e0-4786-b69a-2825eb26e30b",
        "type": "debit",
        "amount": 500,
        "effective_date": "2020-11-08T16:04:28.219161-03:00"
    }
]
```

Status: 404
```json
{
    "reason": "nonexistent account"
}
```

#### `GET` /transactions/:id

Get details of the given transaction

Response:

Status: 200
```json
{
    "id": "985aa1e7-b9e0-4786-b69a-2825eb26e30b",
    "type": "debit",
    "amount": 500,
    "effective_date": "2020-11-08T16:04:28.219161-03:00"
}
```

Status: 404
```json
{
  "reason": "transaction not found"
}
```

#### `GET` /

Get user's balance

Response:

Status: 200
```json
{
  "balance": 500
}
```

Status: 404
```json
{
  "reason": "non existent account"
}
```
