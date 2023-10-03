## Scylla Workers

This service is used to reduce hot partitions on Scylladb clusters.

When the api get a request a new worker is created with the query stored.
If the service already has an worker with this query the request is subscribed to that worker
and when the database query is completed the response is sent to all subscribers.

This service is useful when the server get concurrency requests with the same database query.
With this service only one database query is executed

Example:

```
┌────────────┐  ┌────────────┐ ┌────────────┐ ┌────────────┐
│ Request #1 │  │ Request #2 │ │ Request #3 │ │ Request #4 │
│            │  │            │ │            │ │            │
│ Id: 1      │  │ Id: 1      │ │ Id: 1      │ │ Id: 2      │
└──────┬─────┘  └──────┬─────┘ └─────┬──────┘ └──────┬─────┘
       │               │             │               │
       │               │             │               │
┌──────▼───────────────▼─────────────▼──────┐ ┌──────▼─────┐
│ Worker #1                                 │ │ Worker #2  │
│                                           │ │            │
│ Query: [query statement="SELECT * FROM    │ │ Query: [...│
│ users WHERE id=? LIMIT 1 ALLOW FILTERING" │ │ values=[2]]│
│ values=[1] consistency=QUORUM]            │ │            │
└──────────────────────┬────────────────────┘ └──────┬─────┘
                       │                             │
┌──────────────────────▼─────────────────────────────▼─────┐
│                      Scylla Database                     │
└──────────────────────────────────────────────────────────┘
```
