Install tools

```
cargo install diesel_cli --no-default-features --features postgres
```

# Generate new migration

On Mac M1
```sh
brew install libpq
echo 'export PATH="/opt/homebrew/opt/libpq/bin:$PATH"' >> ~/.zshrc
source ~/.zshrc
```

```sh
diesel migration generate migration_name
```
You will see the outputs.

```sh
Creating migrations/20160815133237_create_posts/up.sql
Creating migrations/20160815133237_create_posts/down.sql
```

Initial database setup
```sh
diesel setup
```

Run the migrations
```sh
diesel migration run
```

To redo the database, run
```sh
diesel migration redo
```

