Dashboard to summary your sales

how to run (make sure you have air install):
`air`

how to run db:
`docker-compose up`

how to run migration (make sure install goose first ya):
`goose -dir ./db/migrations postgres "host=hostname dbname=name user=root password=yayaya sslmode=disable" up;`

create new migration:
`goose create column sql`
