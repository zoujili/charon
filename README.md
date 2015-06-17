Charon [![Build Status](https://travis-ci.org/go-soa/charon.svg)](https://travis-ci.org/go-soa/charon)
=============

<img src="/data/logo/charon.png?raw=true" width="300">

Installation
------------
1. Set you GOPATH properly (http://golang.org/doc/code.html#GOPATH)
2. `go get github.com/go-soa/auth`
3. `go get` if some dependencies are missing
4. Create `conf/{env}.xml` based on `conf/{env}.xml.dist`
5. Set `$AUTH_SERVICE_ENV` global variable to `test`, `development` or `production`

Commands
--------

#### Build
```bash
go build
```

#### Service
```bash
./auth initdb - execute data/sql/schema_{adapter}.sql against configured database.
./auth run - starts server.
./auth help [command] - display help message about available commands
```

Dependencies
------------
- PostgreSQL
- MySQL *(not supported yet)*

TODO
----
- [ ] Commands
	- [x] Initialize database
	- [x] Start server
- [ ] Views
	- [x] Registration
	- [x] Registration success
	- [x] Registration confirmation
	- [x] Login
	- [x] Logout
	- [ ] Password recovery
	- [ ] Password recovery confirmation
	- [x] 400
	- [x] 404
	- [x] 500
- [ ] REST API
	- [ ] Registration
	- [ ] Registration success
	- [ ] Registration confirmation
	- [ ] Login
	- [ ] Logout
	- [ ] Password recovery
	- [ ] Password recovery confirmation
- [ ] RPC API
	- [ ] Registration
	- [ ] Registration success
	- [ ] Registration confirmation
	- [ ] Login
	- [ ] Logout
	- [ ] Password recovery
	- [ ] Password recovery confirmation
- [ ] OAuth2 API
	- [ ] Login
	- [ ] Me
