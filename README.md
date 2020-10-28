## Nomics API

### Prerequisites

Set your `NOMIC_API_KEY` environment variable:

```
export NOMIC_API_KEY="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
```

##### Weekly Coin Prices

Collect the weekly coin prices for current year:

```
make run
```

This will send a get request to the `sparkline` endpoint through the Nomics API

We will transform that data with the `transform.go` script.

The result will be in a file called `[currency]-prices.json`. All prices can be found in `coin-prices.json`.

It should look like this:

```
[
{"currency":"BTC","price":"8221.52708250","timestamp":"2020-01-06T00:00:00Z"},
{"currency":"BTC","price":"8714.82650065","timestamp":"2020-01-13T00:00:00Z"},
{"currency":"BTC","price":"8634.12424157","timestamp":"2020-01-20T00:00:00Z"},
{"currency":"BTC","price":"9319.96695674","timestamp":"2020-01-27T00:00:00Z"},
...
{"currency":"BTC","price":"13745.47312077","timestamp":"2020-10-26T00:00:00Z"}
]
```
