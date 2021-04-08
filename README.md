# Finite Mock Server

The public facing API through which connectors are exposed as a single abstract API.

## Up and running

Start the server with the command below:

~~~bash
$ ./finite-mock-server
2021/04/07 17:31:30 version 0.0.1, commit 6e1b03d15d7f5aeb6453d14043fe6eace379534a, built at 2021-04-07T16:16:36Z by goreleaser
2021/04/07 17:31:30 connected to sqlite3 database
2021/04/07 17:31:30 database migration run successfully
2021/04/07 17:31:30 created account with Id d4234ac6-c550-4503-b21b-a13557c48cbb
2021/04/07 17:31:30 created account with Id 055a00dd-2288-4c01-8d86-67b7d5d6e8f0
2021/04/07 17:31:30 created account with Id 0dea92e1-a495-4600-9683-7a66995310fc
2021/04/07 17:31:30 created account with Id 173f00f1-f12b-4a95-b0ea-b393d224f1ac
2021/04/07 17:31:30 created account with Id 40429067-3006-4fb9-b8ae-ad5111317719
2021/04/07 17:31:30 mock server started on 0.0.0.0:8080
2021/04/07 17:31:30 specs available on http://0.0.0.0:8080/swagger-ui
~~~

Note the Swagger spec is available on [http://0.0.0.0:8080/swagger-ui](http://0.0.0.0:8080/swagger-ui).

The value required for the access token (`X-TOKEN`) defaults to `123456789`; use the example `curl` command below to query one of the accounts:

~~~bash
$ curl -H "Accept: application/json" \
    -H "X-TOKEN: 123456789" \
    http://0.0.0.0:8080/cloud/v1/account/d4234ac6-c550-4503-b21b-a13557c48cbb
~~~

## Limitations

Currently the build (using `goreleaser`) is not able to cross-compile a binary for `linux-amd64` with support for sqlite3; the error message upon start is:

~~~log
failed creating schema resources: Binary was compiled with 'CGO_ENABLED=0', go-sqlite3 requires cgo to work. This is a stub
~~~

As a workaround, use an external database (Postgres / Mysql) as described below.

## Configuration options

The server can be configured using environment variables:
  * `export KEY=VALUE`
  * create a `.env` file with the KEY=VALUE pairs per line

| Key | Default value | Description |
|--- |--- |--- |
| MOCK_SERVER_SECRET | 123456789 | The (static) token used to secure the API; you need to pass this via the HTTP header `X-TOKEN`. |
| MOCK_SERVER_LISTEN_ADDRESS | 0.0.0.0 | The network address the server listens on |
| MOCK_SERVER_LISTEN_PORT | 8080 | The TCP network port the server listens on |
| TEST_DATA_ACCOUNTS | 5 | The number of fake accounts the server generates upon startup. When using a persistent database this is only done when the database contains no data|
| DB_VENDOR | sqlite3 | By default, an in-memory database is used; supported values: `postgres` / `mysql` |
| DB_HOST | - | Database hostname or IP address |
| DB_PORT | - | Database port - defaults to 5432 for postgres, 3306 for mysql |
| DB_DATABASE_NAME | - | Name of the database |
| DB_USER_NAME | - | Username to connect to the database |
| DB_USER_PASSWORD | - | Password of the user to connect to the database |

## Examples

### Postgres

Create a `docker-compose.yaml` file for Postgres and bring up the database using `docker-compose up -d`:

~~~yaml
version: '3.8'

services:
  postgres:
    image: postgres
    restart: always
    environment:
      POSTGRES_DB: finite
      POSTGRES_USER: mock
      POSTGRES_PASSWORD: secret
      PGDATA: /var/lib/postgresql/data/pgdata
    ports:
      - 5432:5432
    expose:
      - 5432
    volumes:
      - finite-mock-postgres:/var/lib/postgresql/data/pgdata

volumes:
  finite-mock-postgres:
~~~

Create the following `.env` file in the same directory as the `finite-mock-server` binary:

~~~bash
#
# Database settings
DB_VENDOR=postgres              # postgres, mysql or sqlite3 (in memory)
DB_HOST=localhost
DB_PORT=5432                    # 5432 for postgres, 3306 for mysql
DB_DATABASE_NAME=finite
DB_USER_NAME=mock
DB_USER_PASSWORD=secret

#
# Mock server settings
MOCK_SERVER_LISTEN_ADDRESS=0.0.0.0
MOCK_SERVER_LISTEN_PORT=8080
MOCK_SERVER_SECRET=VXNlclNlc3Npb25Ub2tlbjoxMjM0MTIzNA==

TEST_DATA_ACCOUNTS=50
~~~

Next, start the mock-server using `./finite-mock-server`.

### MySQL

Create a `docker-compose.yaml` file for MySQL and bring up the database using `docker-compose up -d`:

~~~yaml
version: '3.8'

services:
  mysql:
    image: mysql:5.7
    restart: always
    environment:
      MYSQL_DATABASE: finite
      MYSQL_USER: mock
      MYSQL_PASSWORD: secret
      MYSQL_ROOT_PASSWORD: secret
    ports:
      - 3306:3306
    expose:
      - 3306
    volumes:
      - finite-mock-mysql:/var/lib/mysql

volumes:
  finite-mock-mysql:
~~~

Create the following `.env` file in the same directory as the `finite-mock-server` binary:

~~~bash
#
# Database settings
DB_VENDOR=mysql
DB_HOST=localhost
DB_PORT=3306
DB_DATABASE_NAME=finite
DB_USER_NAME=mock
DB_USER_PASSWORD=secret

#
# Mock server settings
MOCK_SERVER_LISTEN_ADDRESS=0.0.0.0
MOCK_SERVER_LISTEN_PORT=8080
MOCK_SERVER_SECRET=VXNlclNlc3Npb25Ub2tlbjoxMjM0MTIzNA==

TEST_DATA_ACCOUNTS=50
~~~

Next, start the mock-server using `./finite-mock-server`.

### Cleanup

~~~bash
# Bring down the database (but keep the data)
$ docker-compose down --remove-orphans
~~~
