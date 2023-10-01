## Db setup

Start scylla cluster

```
docker-compose up -d
```

Run bash in scylla node

```
docker exec -it scylla-node1 bash
```

Importing keyspaces and tables

```
docker exec scylla-node1 cqlsh -f /users.txt

```

## Todo

- [ ] create an script to auto setup keyspaces and tables
