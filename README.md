# Connect in PostgreSQL with GO

This is a simple project to connect in PostgreSQL database, more details late!!!

To execute in linux

`go build`

next

`./connect.postgresql`

*Important*

In the method `rows.Scan` the order of columns matter. So you need put him in the same order of query.

The case I choose the order of select `select id, name, description from resource` and the `row.Scan(&id, &name, &description)`