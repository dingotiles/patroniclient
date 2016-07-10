# Golang client library for Zalando Patroni

Patroni is a dynamic configuration and process manager for PostgreSQL HA, backed by etcd, consul or zookeeper. Each Patroni process will manage the configuration, runtime and failover sequence of a single PostgreSQL server. Multiple Patroni processes can coordinate via etcd/consul/zookeeper to configure their respective PostgreSQL servers to form asyncrhonous or synchronous replication clusters.

Each Patroni process offers an HTTP API. This project is a Golang client library to the Patroni HTTP API.


## Play time with v0.90

Using registrator, we can discover all Patroni REST APIs:

```
curl -s ${ETCD_CLUSTER}/v2/keys/dingo-postgresql95-8008 | jq -r ".node.nodes[]"
curl -s ${ETCD_CLUSTER}/v2/keys/dingo-postgresql95-8008 | jq -r ".node.nodes[].value"
```

The former returns:

```
[
  {
    "key": "/dingo-postgresql95-8008/1.cell_z1.patroni1.dingo-postgresql.bosh:cf-95520c98-efbe-4b4e-880d-14a8f2bb0ff8:8008",
    "value": "10.244.21.8:32769",
    "modifiedIndex": 13,
    "createdIndex": 13
  },
  {
    "key": "/dingo-postgresql95-8008/0.cell_z1.patroni1.dingo-postgresql.bosh:cf-49a91b69-264b-4df9-a836-523072b2f778:8008",
    "value": "10.244.21.7:32769",
    "modifiedIndex": 10,
    "createdIndex": 10
  }
]
```

The latter returns the `host:port`:

```
10.244.21.7:32769
10.244.21.8:32769
```

These `host:port` can be `curl`ed to get basic information:

```
curl -s ${ETCD_CLUSTER}/v2/keys/dingo-postgresql95-8008 | jq -r ".node.nodes[].value" | xargs -L1 curl -s | jq .
```

Returns:

```
{
  "database_system_identifier": "6304830698674094206",
  "postmaster_start_time": "2016-07-08 06:03:11.407 UTC",
  "xlog": {
    "location": 100663392
  },
  "patroni": {
    "scope": "025ea0b0-710e-4da2-890d-f245a4d35259",
    "version": "0.90"
  },
  "state": "running",
  "role": "master",
  "server_version": 90503
}
{
  "database_system_identifier": "6304830698674094206",
  "postmaster_start_time": "2016-07-08 06:03:24.859 UTC",
  "xlog": {
    "received_location": 100663392,
    "replayed_timestamp": null,
    "paused": false,
    "replayed_location": 100663392
  },
  "patroni": {
    "scope": "025ea0b0-710e-4da2-890d-f245a4d35259",
    "version": "0.90"
  },
  "state": "running",
  "role": "replica",
  "server_version": 90503
}
```

As of v0.90, sadly each result does not include self-identifying information. Only that a result is a `"role": "master"` or `"role": "replica"`. For the latter, the `"xlog"` information shows that health of the replica.
