# Database Migration

```shell
psql -h localhost -U music -d music -f db/schema.sql
psql -h localhost -U music -d music -f db/seed.sql
```