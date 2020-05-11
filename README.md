# mq - get stock market quotes

0.0.1

### Summary

mq is a command line tool for getting market quotes

AlphaVantage is the quote provider, claim your free api key [here](https://www.alphavantage.co/support/#api-key)

```sh
$ mq -s GOOG,TSLA,BTC,AAPL -k {YOUR_ALPHAVANTAGE_API_KEY}

Symbol    Price ($)  Change today (%)  
------    ---------  ----------------  
GOOG      724.12     -0.16%   
AAPL      100.41     +0.79%   
TSLA      225.12     +2.52%   
BTCUSD=X  452.05     +0.90%   

```

### Installation

mq requires the Go Tools (v1.6+ tested, but others are probably fine). Instructions are [here](https://golang.org/doc/install) for installation.

`make` builds the binary and `make install` places it on your path.

```sh
$ git clone {CLONE_URL}
$ cd mq
$ make && make install
```

### Usage

```sh
mq -s <COMMA SEPERATED SYMBOLS> -f <FILE OF NEWLINE SEPERATED SYMBOLS>
```

### Other stuff

Contributions are welcome.

Future plans for development:
* Better error handling
* Add secrets management for API key through a .env file
* Add unit tests