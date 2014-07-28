experiment using [vegeta](https://github.com/tsenart/vegeta) as a lib

it uses a fork [https://github.com/thierryhenrio/vegeta](https://github.com/thierryhenrio/vegeta) which is a hack to enable sending unique request

what pains me is to rewrite a main whenever we want different requests ...
in this context [wrk](https://github.com/wg/wrk) looks like a simpler approach

usage
=====

```
./goku | ../vegeta/vegeta report
```

`goku` does not have any options ... yet
option is change in main and recompile ( sic :)
