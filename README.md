# client-manager


para rodar as migrações 

```
    goose -dir internal/db/migrations up
```

para criar uma nova migração
```
    goose create add_users_table sql -dir internal/db/migrations
```

parar rodar o sqlc
```
    sqlc generate -f sqlc.yml
```