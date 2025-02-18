## okp4d query ibc channel unreceived-acks

Query all the unreceived acks associated with a channel

### Synopsis

Given a list of acknowledgement sequences from counterparty, determine if an ack on the counterparty chain has been received on the executing chain.

The return value represents:

- Unreceived packet acknowledgement: packet commitment exists on original sending (executing) chain and ack exists on receiving chain.

```
okp4d query ibc channel unreceived-acks [port-id] [channel-id] [flags]
```

### Examples

```
okp4d query ibc channel unreceived-acks [port-id] [channel-id] --sequences=1,2,3
```

### Options

```
      --height int             Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help                   help for unreceived-acks
      --node string            &lt;host&gt;:&lt;port&gt; to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string          Output format (text|json) (default "text")
      --sequences int64Slice   comma separated list of packet sequence numbers (default [])
```

### Options inherited from parent commands

```
      --chain-id string   The network chain ID (default "okp4d")
```

### SEE ALSO

- [okp4d query ibc channel](okp4d_query_ibc_channel.md)	 - IBC channel query subcommands
