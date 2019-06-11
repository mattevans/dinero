# dinero

[![GoDoc](https://godoc.org/github.com/mattevans/dinero?status.svg)](https://godoc.org/github.com/mattevans/dinero)
[![Build Status](https://travis-ci.org/mattevans/dinero.svg?branch=master)](https://travis-ci.org/mattevans/dinero)
[![Go Report Card](https://goreportcard.com/badge/github.com/mattevans/dinero)](https://goreportcard.com/report/github.com/mattevans/dinero)
[![license](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/mattevans/dinero/blob/master/LICENSE)

dinero is a [Go](http://golang.org) client library for accessing the Open Exchange Rates API (https://docs.openexchangerates.org/docs/).

Upon request of forex rates these will be cached (in-memory), keyed by base currency. With a customisable expiry window, subsequent requests will use cached data or fetch fresh data accordingly.

Installation
-----------------

`go get -u github.com/mattevans/dinero`

Usage
-----------------

**List**

```go
// Init dinero client, passing your OXR app ID and a base 
// currency for conversion to work from.
client := NewClient(
  os.Getenv("OPEN_EXCHANGE_APP_ID"), 
  "AUD",
)

// Get latest forex rates.
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
// Init dinero client, passing your OXR app ID and a base 
// currency for conversion to work from.
client := NewClient(
  os.Getenv("OPEN_EXCHANGE_APP_ID"), 
  "AUD",
)

// Set a base currency to work with.
client.Rates.SetBaseCurrency("AUD")

// Get latest forex rate for NZD using AUD as a base currency.
rsp, err := client.Rates.Single("NZD")
if err != nil {
  return err
}
```

```json
1.045545
```

---

**Change Base Currency**

You set a base currency when your intialize dinero client. Should you wish to change this, you can call...

```
client.Rates.SetBaseCurrency("USD")
```

---

**Expire**

By default, cached rates will expire themselves based on your configured cache TTL.

You can force an expiry of the rates for your currency set by calling...

```go
client.Cache.Expire()
```

Contributing
-----------------
If you've found a bug or would like to contribute, please create an issue here on GitHub, or better yet fork the project and submit a pull request!
