# go-sqlcipher for Mira

This is Mira's maintained fork of `mutecomm/go-sqlcipher`: a self-contained
`database/sql` driver that bundles SQLCipher rather than relying on a system
SQLite library.

## What this fork changes

`v4.17.0-mira.2` updates the bundled database engine from SQLCipher 4.4.2 to
SQLCipher 4.17.0 (SQLite 3.53.3). It also:

- uses SQLCipher's maintained OpenSSL provider instead of the unmaintained
  bundled LibTomCrypt integration;
- defines SQLCipher's required initialization and shutdown hooks;
- supports FTS5 when the consuming program is built with `-tags fts5` (or
  `-tags sqlite_fts5`);
- quotes passphrases safely before issuing `PRAGMA key`.

The package links against OpenSSL's `libcrypto`. Build hosts therefore need the
OpenSSL development headers and library (for example, `libssl-dev` on Debian or
Ubuntu).

## Installation

```sh
go get github.com/benoitpetit/go-sqlcipher/v4@v4.17.0-mira.2
```

Import the driver for its `database/sql` registration:

```go
import _ "github.com/benoitpetit/go-sqlcipher/v4"
```

Build applications that need full-text search with:

```sh
go build -tags fts5 ./...
```

## Opening an encrypted database

Pass the key as `_pragma_key` in the DSN. Always URL-escape a passphrase;
this preserves spaces, quotes, and other special characters.

```go
key := url.QueryEscape("correct horse battery staple")
db, err := sql.Open("sqlite3", "mira.db?_pragma_key="+key+"&_pragma_cipher_page_size=4096")
if err != nil {
	log.Fatal(err)
}
defer db.Close()
```

For a 32-byte hexadecimal key:

```go
key := "2DD29CA851E7B56E4697B0E1F08507293D761A05CE4D1B628663F411A8086D99"
db, err := sql.Open("sqlite3", "mira.db?_pragma_key=x'"+key+"'")
```

`sqlite3.IsEncrypted(path)` reports whether a file appears encrypted. It does
not validate that a supplied key is correct; issue a query after opening to do
that.

## Compatibility and migration

SQLCipher 4 databases remain in the SQLCipher 4 format. This fork tests opening
an existing SQLCipher 4 fixture, but every production database should be backed
up and tested in a staging environment before upgrading. SQLCipher 3 and 4 are
not format-compatible; follow SQLCipher's [migration guide](https://www.zetetic.net/sqlcipher/sqlcipher-api/#Migrating_Databases)
for that transition.

The upstream APIs come from [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3)
and [SQLCipher](https://github.com/sqlcipher/sqlcipher). This fork is scoped to
Mira's SQLCipher build and maintenance needs.

## License

The originating packages retain their respective licenses. See [LICENSE](LICENSE).
