// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "pierre.saintsorny@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "summary": "Index of the api",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/account": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get one or multiple account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_account",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Account"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_account",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "38fzefzef74723jh",
                        "name": "api_key",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "38fzefzef74723jh",
                        "name": "secret_key",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete an account",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_account",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/account": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Account Informations about one or multiple asset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTC",
                        "name": "asset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.BinanceAccount"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/futures/account": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Futures Account Informations about one or multiple asset",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTC",
                        "name": "asset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "my_account",
                        "name": "account_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FuturesAccount"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/futures/order/buy": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Buy from binance futures using market order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_first_strategy",
                        "name": "strategy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "0.01",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FuturesOrder"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/futures/order/sell": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "sell from binance futures using market order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_first_strategy",
                        "name": "strategy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "0.01",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FuturesOrder"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/futures/position": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Futures position informations",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTCUSDT",
                        "name": "symbol",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "my_account",
                        "name": "account_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.FuturesPosition"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/order/buy": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Buy from binance using market order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_first_strategy",
                        "name": "strategy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "0.01",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/order/sell": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "sell from binance using market order",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_first_strategy",
                        "name": "strategy",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "0.01",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Order"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/binance/price": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get price for a given symbol",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTCUSDT",
                        "name": "symbol",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Prices"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/convert": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Convert from asset one to two with a given quantity",
                "parameters": [
                    {
                        "type": "string",
                        "description": "USDT",
                        "name": "from",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "BTC",
                        "name": "to",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "100",
                        "name": "quantity",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.Converter"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/data/hourly": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Hourly data fetch of binance from the scrapper",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTCUSDT",
                        "name": "symbol",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "5",
                        "name": "x",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.HourlyDataFetchOnCurrency"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/strategy": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get one or multiple strategy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_first_strategy",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Strategy"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new strategy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "BTCUSDT",
                        "name": "symbol",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "my_first_strategy_BTCUSDT",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "1h",
                        "name": "interval",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "10",
                        "name": "leverage",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "undefined",
                        "name": "account_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Delete a strategy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_first_strategy",
                        "name": "name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        },
        "/strategy/result": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get the total pnl and commission for a strategy",
                "parameters": [
                    {
                        "type": "string",
                        "description": "my_strategy",
                        "name": "strategy_name",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TotalPNLAndCommission"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/models.APIError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.APIError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "models.Account": {
            "type": "object",
            "properties": {
                "apiKey": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "insertedDate": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "secretKey": {
                    "type": "string"
                }
            }
        },
        "models.BinanceAccount": {
            "type": "object",
            "properties": {
                "balances": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "asset": {
                                "type": "string"
                            },
                            "free": {
                                "type": "string"
                            },
                            "locked": {
                                "type": "string"
                            }
                        }
                    }
                },
                "canDeposit": {
                    "type": "boolean"
                },
                "canTrade": {
                    "type": "boolean"
                },
                "canWithdraw": {
                    "type": "boolean"
                },
                "makerCommission": {
                    "type": "integer"
                },
                "permissions": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "takerCommission": {
                    "type": "integer"
                },
                "updateTime": {
                    "type": "integer"
                }
            }
        },
        "models.Converter": {
            "type": "object",
            "properties": {
                "from": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                },
                "to": {
                    "type": "string"
                }
            }
        },
        "models.FuturesAccount": {
            "type": "object",
            "properties": {
                "asset": {
                    "type": "string"
                },
                "availableBalance": {
                    "type": "string"
                }
            }
        },
        "models.FuturesOrder": {
            "type": "object",
            "properties": {
                "commission": {
                    "type": "string"
                },
                "orderId": {
                    "type": "number"
                },
                "origQty": {
                    "type": "string"
                },
                "price": {
                    "type": "string"
                },
                "realizedPNL": {
                    "type": "string"
                },
                "side": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "updateTime": {
                    "type": "integer"
                }
            }
        },
        "models.FuturesPosition": {
            "type": "object",
            "properties": {
                "entryPrice": {
                    "type": "string"
                },
                "leverage": {
                    "type": "string"
                },
                "positionAmt": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "unRealizedProfit": {
                    "type": "string"
                },
                "updateTime": {
                    "$ref": "#/definitions/models.intToTime"
                }
            }
        },
        "models.HourlyDataFetchOnCurrency": {
            "type": "object",
            "properties": {
                "bollingerBands": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "chopinessIndex": {
                    "type": "number"
                },
                "close": {
                    "type": "number"
                },
                "closeTime": {
                    "type": "string"
                },
                "currencyName": {
                    "type": "string"
                },
                "high": {
                    "type": "number"
                },
                "ichimoku": {
                    "type": "array",
                    "items": {
                        "type": "number"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "insertedDate": {
                    "type": "string"
                },
                "low": {
                    "type": "number"
                },
                "ma": {
                    "type": "number"
                },
                "obv": {
                    "type": "number"
                },
                "open": {
                    "type": "number"
                },
                "openTime": {
                    "type": "string"
                },
                "rsi": {
                    "type": "number"
                },
                "sma": {
                    "type": "number"
                },
                "stochastic": {
                    "type": "number"
                },
                "supertrend": {
                    "type": "string"
                },
                "volume": {
                    "type": "number"
                }
            }
        },
        "models.Order": {
            "type": "object",
            "properties": {
                "executedQty": {
                    "type": "string"
                },
                "fills": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "commission": {
                                "type": "string"
                            },
                            "price": {
                                "type": "string"
                            },
                            "qty": {
                                "type": "string"
                            }
                        }
                    }
                },
                "orderId": {
                    "type": "number"
                },
                "side": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                },
                "transactTime": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.Prices": {
            "type": "object",
            "properties": {
                "price": {
                    "type": "string"
                },
                "symbol": {
                    "type": "string"
                }
            }
        },
        "models.Strategy": {
            "type": "object",
            "properties": {
                "accountName": {
                    "type": "string"
                },
                "currencyName": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "interval": {
                    "type": "string"
                },
                "leverage": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TotalPNLAndCommission": {
            "type": "object",
            "properties": {
                "commission": {
                    "type": "number"
                },
                "pnl": {
                    "type": "number"
                }
            }
        },
        "models.intToTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1",
	Host:             "",
	BasePath:         "/1.0",
	Schemes:          []string{},
	Title:            "Appolo API",
	Description:      "This is the documentation for the golang api of appolo-api",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
