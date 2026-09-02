package sqlite3

/*
// make go-sqlite3 use embedded library without code changes
#cgo CFLAGS: -DUSE_LIBSQLITE3

// enable encryption codec in sqlite
#cgo CFLAGS: -DSQLITE_HAS_CODEC

// use memory for temporay storage in sqlite
#cgo CFLAGS: -DSQLITE_TEMP_STORE=2

// initialise and shut down SQLCipher's codec extension.
// SQLCipher 4.17+ requires these hooks to be set explicitly.
#cgo CFLAGS: -DSQLITE_EXTRA_INIT=sqlcipher_extra_init
#cgo CFLAGS: -DSQLITE_EXTRA_SHUTDOWN=sqlcipher_extra_shutdown

// use SQLCipher's maintained OpenSSL provider
#cgo CFLAGS: -DSQLCIPHER_CRYPTO_OPENSSL
#cgo LDFLAGS: -lcrypto

// disable assertions
#cgo CFLAGS: -DNDEBUG

// set operating specific sqlite flags
#cgo linux CFLAGS: -DSQLITE_OS_UNIX=1
#cgo windows CFLAGS: -DSQLITE_OS_WIN=1
*/
import "C"
