## okp4d query slashing signing-infos

Query signing information of all validators

### Synopsis

signing infos of validators:

$ okp4d query slashing signing-infos

```
okp4d query slashing signing-infos [flags]
```

### Options

```
      --count-total       count total number of records in signing infos to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for signing-infos
      --limit uint        pagination limit of signing infos to query for (default 100)
      --node string       &lt;host&gt;:&lt;port&gt; to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of signing infos to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of signing infos to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of signing infos to query for
      --reverse           results are sorted in descending order
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID (default "okp4d")
```

### SEE ALSO

* [okp4d query slashing](okp4d_query_slashing.md)	 - Querying commands for the slashing module
