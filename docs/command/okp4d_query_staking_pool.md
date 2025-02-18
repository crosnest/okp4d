## okp4d query staking pool

Query the current staking pool values

### Synopsis

Query values for amounts stored in the staking pool.

Example:
$ okp4d query staking pool

```
okp4d query staking pool [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for pool
      --node string     &lt;host&gt;:&lt;port&gt; to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID (default "okp4d")
```

### SEE ALSO

* [okp4d query staking](okp4d_query_staking.md)	 - Querying commands for the staking module
