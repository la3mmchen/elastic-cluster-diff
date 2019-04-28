# elastic-cluster-diff

This is a tool to diff elasticsearch cluster. Right now it is only possible to detect different document counts in indices with the same name.

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


## example

```bash
  $ ./bin/main compare --config --cluster my-fancy-cluster-1:9200 --cluster my-fancy-cluster-2:9200

  Diff +my-fancy-cluster-1:9200 with -my-fancy-cluster-2:9200
    diff.Changelog{
    {
        Type: "update",
        Path: {"TotalDocuments"},
        From: int(758),
        To:   int(760),
    },
    {
        Type: "delete",
        Path: {"Indices", "1", "Name"},
        From: ".kibana_1",
        To:   nil,
    },
    {
        Type: "delete",
        Path: {"Indices", "1", "DocCount"},
        From: int(9),
        To:   nil,
    },
    {
        Type: "create",
        Path: {"Indices", "8", "Name"},
        From: nil,
        To:   ".kibana_1",
    },
    {
        Type: "create",
        Path: {"Indices", "8", "DocCount"},
        From: nil,
        To:   int(11),
    },
    (...)
```

LLAP.