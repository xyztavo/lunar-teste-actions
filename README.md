# lunar stack
this is a monolith stack that scales well and in the same time it has a great developer experience.

### should i use this?
for personal and small projects, this is amazing. if you want to deploy it, you really should make backups of the sqlite or replace it completely (maybe postgres or turso?).

## stack

### backend
- go (compiles fast to binary code, great developer experience)
- templ (typesafe react-like templating golang experience)

### frontend
- tailwind (styling)
- Alpine.js (for client side UX)
- Alpine AJAX (HTMX-like JS LIB)

### database
- sqlite (database, turso in prod needs no code change)
- Atlas (migrations)
- sqlc (typesafe queries)

## dependencies (Install any of those if needed)
- go-task (Taskfile.yaml runner)
- templ (to generate golang for templating)
- tailwindcss-cli (to generate output.css)
- air (golang hot reloading)
- atlas (cli needed to apply migrations)

## Usage
List tasks:

```sh
go-task
```

run a task:

```sh
go-task {CMD}
```

main tasks:
- `dev`: start the full dev environment (templ + tailwind + air)
- `build`: build a single Go binary for production
- `migrate`: sync the schema with the database using Atlas


### rename .env.example to .env.
- DB_URL="file:db.db" (this works fine for developer testing)