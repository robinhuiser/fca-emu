# Finite Cloud API Emulator

The public facing API through which connectors are exposed as a single abstract API.

## Up and running

Start the server with the command below:

~~~bash
# As a standalone process (OSX amd64 only)
$ ./fca-emu

# As a Docker
$ docker run -p 8080:8080 tecdigitalacr.azurecr.io/utils/fca-emu
~~~

Please note:

* The Swagger spec is available on [http://0.0.0.0:8080/swagger-ui](http://0.0.0.0:8080/swagger-ui).
* The value required for the access token (`X-TOKEN`) defaults to `123456789`; 
* During startup, the emulator logs:
  * the available product types
  * the IDs for available entities and accounts

Use the example `curl` command below to query one of the accounts:

~~~bash
$ curl -H "Accept: application/json" \
    -H "X-TOKEN: 123456789" \
    http://0.0.0.0:8080/cloud/v1/account/ID_FROM_EMULATOR_LOG
~~~

## Configuration options

The emulator can be configured using environment variables below - when using the OSX binary you can create a `.env` file in the current working directory to set the values.

| Key | Default value | Description |
|--- |--- |--- |
| FCA_SECRET | 123456789 | The (static) token used to secure the API; you need to pass this via the HTTP header `X-TOKEN`. |
| EMULATOR_LISTEN_ADDRESS | 0.0.0.0 | The network address the server listens on |
| EMULATOR_LISTEN_PORT | 8080 | The TCP network port the server listens on |
| NUMBER_OF_ENTITIES | 3 | The number of entities (persons, organizations and corporates) generated during startup. Each entity will be assigned an n number of accounts by default|
| BRANCHES_PER_BANK | 50 | Number of branches generated for each of the banks |
| DB_VENDOR | sqlite3 | By default, an in-memory database is used; supported values: `postgres` / `mysql` |
| DB_HOST | - | Database hostname or IP address |
| DB_PORT | - | Database port - defaults to `5432` for postgres, `3306` for mysql |
| DB_DATABASE_NAME | - | Name of the database |
| DB_USER_NAME | - | Username to connect to the database |
| DB_USER_PASSWORD | - | Password of the user to connect to the database |

The example `docker-compose.yaml` below starts a Postgres database and configures the emulator with a custom security token and test data:

~~~yaml
version: '3.8'

services:
  postgres:
    image: postgres
    restart: always
    environment:
      POSTGRES_DB: fca
      POSTGRES_USER: emulator
      POSTGRES_PASSWORD: secret
      PGDATA: /var/lib/postgresql/data/pgdata
    ports:
      - 5432:5432
    expose:
      - 5432
    volumes:
      - postgres-volume:/var/lib/postgresql/data/pgdata
  
  fca-emu:
    image: tecdigitalacr.azurecr.io/utils/fca-emu
    restart: always
    environment:
      DB_VENDOR: postgres
      DB_HOST: postgres
      DB_DATABASE_NAME: fca
      DB_USER_NAME: emulator
      DB_USER_PASSWORD: secret
      FCA_SECRET: VXNlclNlc3Npb25Ub2tlbjoxMjM0MTIzNA==
      TEST_DATA_ACCOUNTS: 100
    ports:
      - 8080:8080  

volumes:
  mysql-volume:
~~~
