# dinero

[![GoDoc](https://godoc.org/github.com/mattevans/dinero?status.svg)](https://godoc.org/github.com/mattevans/dinero)
[![Build Status](https://travis-ci.org/mattevans/dinero.svg?branch=master)](https://travis-ci.org/mattevans/dinero)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattevans/dinero)](https://goreportcard.com/report/github.com/mattevans/dinero)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mattevans/dinero/blob/master/LICENSE)

dinero is a [Go](http://golang.org) client library for accessing the Open Exchange Rates API (https://docs.openexchangerates.org/docs/).

Any forex rates requested will be cached (in-memory), keyed by base currency. With a customisable expiry window, subsequent requests will use cached data or fetch fresh data accordingly.

Installation
-----------------

`go get -u github.com/mattevans/dinero`

Usage
-----------------

**Intialize**

```go
// Init dinero client passing....
// - your OXR app ID
// - base currency code for conversions to work from
// - your preferrerd cache expiry
client := NewClient(
  os.Getenv("OPEN_EXCHANGE_APP_ID"), 
  "AUD",
  20*time.Minute,
)
```

---

## Currencies

**List**

```go
// List all currencies available.
rsp, err := client.Currencies.List()
if err != nil {
  return err
}
```

```json
[
  {
    "code": "INR",
    "name": "Indian Rupee"
  },
  {
    "code": "PYG",
    "name": "Paraguayan Guarani"
  },
  {
      "code": "AED",
      "name": "United Arab Emirates Dirham"
  },
  ...
}
```

---

## Rates

**List**

```go
// List latest forex rates. This will use AUD (defined when intializing the client) as the base.
rsp, err := client.Rates.List()
if err != nil {
  return err
}
```

```json
{
   "rates":{
      "AED": 2.702388,
      "AFN": 48.893275,
      "ALL": 95.142814,
      "AMD": 356.88691,
      ...
   },
   "base": "AUD"
}
```

---

**Get**

```go
// Get latest forex rate for NZD. This will use AUD (defined when intializing the client) as the base.
rsp, err := client.Rates.Get("NZD")
if err != nil {
  return err
}
```

```
1.045545
```
---

## Historical Rates

**List**

```go
historicalDate := time.Now().AddDate(0, -5, -3)

// List historical forex rates. This will use AUD (defined when intializing the client) as the base.
rsp, err := client.Historical.List(historicalDate)
if err != nil {
  return err
}
```

```json
{
   "rates":{
      "AED": 2.702388,
      "AFN": 48.893275,
      "ALL": 95.142814,
      "AMD": 356.88691,
      ...
   },
   "base": "AUD"
}
```

---

**Get**

```go
historicalDate := time.Now().AddDate(0, -5, -3)

// Get historical forex rate for NZD. This will use AUD (defined when intializing the client) as the base.
rsp, err := client.Rates.Get("NZD", historicalDate)
if err != nil {
  return err
}
```

```
1.045545
```

---

**Change Base Currency**

You set a base currency when you the intialize dinero client. Should you wish to change this at anytime, you can call...

```go
client.Rates.SetBaseCurrency("USD")
```

> NOTE: Changing the API `base` currency is available for Developer, Enterprise and Unlimited plan clients only.

---

**Expire**

You set your preferred cache expiry interval when you intialize the dinero client. By default, cached rates will expire themselves based on your configured interval.

You can force an expiry of the rates for your currency set by calling...

```go
client.Cache.Expire()
```

Contributing
-----------------
If you've found a bug or would like to contribute, please create an issue here on GitHub, or better yet fork the project and submit a pull request!
