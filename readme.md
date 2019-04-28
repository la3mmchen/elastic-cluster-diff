# elastic-cluster-diff

This is a tool to diff elasticsearch cluster.

## Why

My use case are two georedundant elasticsearch clusters that should have the same data. This architecture was born before the cross-cluster-replication was released.

Other usecases might be: Compare configs from elasticsearch cluster without digging into the config management. This should also work with ElasticAsAService as it only uses the HTTP api (depends on the provided apis).

## build and run

We do provide a docker image that contains all the stuff. Either pull from nexus (1) or build it yourself:

```bash
    $  make build
```

```bash
    $ ./bin/main
```
