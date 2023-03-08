## Factor Core

Official Golang implementation of the Factor protocol.

## Building the source

Building `factor` requires both a Go (version 1.17 or later) and a C compiler. You can install
them using your favourite package manager. Once the dependencies are installed, run

```shell
make factor
```

## Executables

The `factor-core` project comes with several wrappers/executables found in the `cmd` directory.

|   Command    | Description                                                                                                                                                                                                                                                                                                                                                                                     |
| :----------: | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **`factor`** | Our main Factor CLI client. It is the entry point into the Factor network (main-net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Factor network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `factor --help` |

### Hardware Requirements

Minimum:

- CPU with 2+ cores
- 4GB RAM
- 50GB free storage space to sync the Mainnet
- 8 MBit/sec download Internet service

Recommended:

- Fast CPU with 4+ cores
- 8GB+ RAM
- High Performance SSD with at least 100GB free space
- 25+ MBit/sec download Internet service

## Starting up Factor Main network

If you want to participate in our network, you must contact FactorLabs. (contact, admin@gipmx.com)

## License

The `factor-core` is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html).
