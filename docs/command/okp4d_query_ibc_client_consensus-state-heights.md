## okp4d query ibc client consensus-state-heights

Query the heights of all consensus states of a client.

### Synopsis

Query the heights of all consensus states associated with the provided client ID.

```
okp4d query ibc client consensus-state-heights [client-id] [flags]
```

### Examples

```
okp4d query ibc client consensus-state-heights [client-id]
```

### Options

```
      --count-total       count total number of records in consensus state heights to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for consensus-state-heights
      --limit uint        pagination limit of consensus state heights to query for (default 100)
      --node string       &lt;host&gt;:&lt;port&gt; to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of consensus state heights to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of consensus state heights to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of consensus state heights to query for
      --reverse           results are sorted in descending order
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID (default "okp4d")
```

### SEE ALSO

* [okp4d query ibc client](okp4d_query_ibc_client.md)	 - IBC client query subcommands
