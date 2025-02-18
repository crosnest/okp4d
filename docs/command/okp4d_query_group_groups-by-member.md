## okp4d query group groups-by-member

Query for groups by member address with pagination flags

```
okp4d query group groups-by-member [address] [flags]
```

### Options

```
      --count-total       count total number of records in groups-by-members to query for
      --height int        Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help              help for groups-by-member
      --limit uint        pagination limit of groups-by-members to query for (default 100)
      --node string       &lt;host&gt;:&lt;port&gt; to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
      --offset uint       pagination offset of groups-by-members to query for
  -o, --output string     Output format (text|json) (default "text")
      --page uint         pagination page of groups-by-members to query for. This sets offset to a multiple of limit (default 1)
      --page-key string   pagination page-key of groups-by-members to query for
      --reverse           results are sorted in descending order
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID (default "okp4d")
```

### SEE ALSO

* [okp4d query group](okp4d_query_group.md)	 - Querying commands for the group module
