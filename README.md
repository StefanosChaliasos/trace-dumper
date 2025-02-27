# trace-dumper

A very simple tool, it can deploy contract and call contract api and get trace store into file just in one cmd.

## upgrade dependence

1. upgrade go mod

```
go get -v github.com/scroll-tech/go-ethereum@staging // change `staging` to a specific tag/branch here
go mod tidy
```

2. manaully edit `scrolltech/l2geth`'s tag in `docker/l2geth/Dockerfile` to the corresponding version

For example, `scrolltech/l2geth:prealpha-v3.1`

## make and start l2geth docker

create environment (**need to keep it running**)

```
make docker
docker run -it -p 8545:8545 -p 8546:8546 --rm trace-dumper/l2geth:latest
```

## dump traces

```
# build trace_dumper
make trace_dumper

# --help show detail about flags.
./bin/trace_dumper --help

# without `-wrap` get the origin result from sdk.
./bin/trace_dumper -dump erc20 # options: erc20, native, nft, greeter, sushi, dao, uniswapv2, multi_uniswapv2

# `-wrap` add json rpc wrap, in order to get the same struct when called by postman.
./bin/trace_dumper -dump erc20 -wrap

# Run the yul tests
./bin/trace_dumper -dump yuljson -json jsons/yul.json
```

## show trace list

```
ls -l ./tracedata/erc20_*.json
```
