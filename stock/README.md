# Stock Service

Get live stock quotes and prices for thousands of US and global stocks.


## Describe the stock service
```shell
micro describe service stock
```

### Response
```json
{
  "name": "stock",
  "version": "latest",
  "metadata": null,
  "endpoints": [
    {
      "name": "Stock.History",
      "request": {
        "name": "HistoryRequest",
        "type": "HistoryRequest",
        "values": [
          {
            "name": "stock",
            "type": "string",
            "values": null
          },
          {
            "name": "date",
            "type": "string",
            "values": null
          }
        ]
      },
      "response": {
        "name": "HistoryResponse",
        "type": "HistoryResponse",
        "values": [
          {
            "name": "symbol",
            "type": "string",
            "values": null
          },
          {
            "name": "open",
            "type": "float64",
            "values": null
          },
          {
            "name": "close",
            "type": "float64",
            "values": null
          },
          {
            "name": "high",
            "type": "float64",
            "values": null
          },
          {
            "name": "low",
            "type": "float64",
            "values": null
          },
          {
            "name": "volume",
            "type": "int32",
            "values": null
          },
          {
            "name": "date",
            "type": "string",
            "values": null
          }
        ]
      },
      "metadata": {}
    },
    {
      "name": "Stock.Price",
      "request": {
        "name": "PriceRequest",
        "type": "PriceRequest",
        "values": [
          {
            "name": "symbol",
            "type": "string",
            "values": null
          }
        ]
      },
      "response": {
        "name": "PriceResponse",
        "type": "PriceResponse",
        "values": [
          {
            "name": "symbol",
            "type": "string",
            "values": null
          },
          {
            "name": "price",
            "type": "float64",
            "values": null
          }
        ]
      },
      "metadata": {}
    },
    {
      "name": "Stock.Quote",
      "request": {
        "name": "QuoteRequest",
        "type": "QuoteRequest",
        "values": [
          {
            "name": "symbol",
            "type": "string",
            "values": null
          }
        ]
      },
      "response": {
        "name": "QuoteResponse",
        "type": "QuoteResponse",
        "values": [
          {
            "name": "symbol",
            "type": "string",
            "values": null
          },
          {
            "name": "ask_price",
            "type": "float64",
            "values": null
          },
          {
            "name": "bid_price",
            "type": "float64",
            "values": null
          },
          {
            "name": "ask_size",
            "type": "int32",
            "values": null
          },
          {
            "name": "bid_size",
            "type": "int32",
            "values": null
          },
          {
            "name": "timestamp",
            "type": "string",
            "values": null
          }
        ]
      },
      "metadata": {}
    }
  ],
  "nodes": [
    {
      "id": "stock-dba510c4-0262-4c17-98d9-69b0cbbac27c",
      "address": "192.168.0.33:65117",
      "metadata": {
        "broker": "http",
        "protocol": "mucp",
        "registry": "mdns",
        "server": "mucp",
        "transport": "http"
      }
    }
  ]
}
```

## Call the stock service
```shell
micro call stock Stock.Quote '{"symbol": "MSFT"}'
```

### Response
```json
{
  "ask_price":419.4655,
  "ask_size":1,
  "bid_price":419.2423,
  "bid_size":3,
  "symbol":"MSFT",
  "timestamp":"2025-01-11T00:59:57.001Z"
}
```

## References:
- https://finage.co.uk/pricing

