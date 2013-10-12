# Introduction
This application is intended to act as a job post list. It is written in Go
using the [revel](https://github.com/robfig/revel) framework.

This app relies on [gorp](https://github.com/coopernurse/gorp) as the ORM-like
library and [goose]( https://bitbucket.org/liamstask/goose) for migrations.

# Setup

## Download and install the app
Run this in terminal:

```sh
go get github.com/gylaz/jobs
```

The application should now be installed in your `$GOPATH/src`.

## Create the database
Create `jobs_app` database using `psql` command:

```sh
psql postgres -c "CREATE DATABASE jobs_app"
```

## Run migrations

```sh
goose up
```

# Running
Run the app via the `revel` command:

```sh
revel run github.com/gylaz/jobs
```
