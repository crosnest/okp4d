## okp4d query ibc client states

Query all available light clients

### Synopsis

Query all available light clients

```
okp4d query ibc client states [flags]
```

### Examples

```
okp4d query ibc client states
```

### Options

```
      --count-total       count total number of records in client states to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for states
      --limit uint        pagination limit of client states to query for (default 100)
      --node string       &lt;host&gt;:&lt;port&gt; to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of client states to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of client states to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of client states to query for
      --reverse           results are sorted in descending order
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID (default "okp4d")
```

### SEE ALSO

* [okp4d query ibc client](okp4d_query_ibc_client.md)	 - IBC client query subcommands
