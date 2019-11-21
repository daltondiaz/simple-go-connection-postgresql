# Connect in PostgreSQL with GO

This is a simple project to connect in PostgreSQL database. Today in all the languages you need handle with database, so in this period I'm studying Go Lang and create this simples step-by-step to connect in db.

## What you need?

- Go lang (Obvious, right?!)
- PostgreSQL (I used [DBeaver](https://dbeaver.io) to manage the DB its really good, I recommend!!)
- Lib to connect in postgresql: github.com/lib/pq (in terminal run: `go get github.com/lib/pq`)

## SQL

### create table

```sql
-- create table resources
CREATE TABLE public.resource (
	name text NULL,
	description text NULL,
	id int4 NOT NULL
);
```

### Insert the first item

```sql
INSERT INTO public.resource
("name", description, id)
VALUES('Eqp-108', 'Equipament 108', 1);
```

Add more items in this simples table for your tests.

## Build

To execute in linux

`go build`

next

`./connect.postgresql`

## Notes

In the method `rows.Scan` the order of columns matter. So you need put him in the same order of query. The case I choose the order of select `select id, name, description from resource` and the row.Scan variables is `row.Scan(&id, &name, &description)`

## Do you have problemns? Please create a Issue for me help u!!