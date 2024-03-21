```sh
sqitch deploy --target local_tradingdb
```

```sh
sqitch log --target local_tradingdb
sqitch revert --target local_tradingdb --to add_symbol_info
```