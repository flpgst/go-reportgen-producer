# Go-Report Producer

Using [XO](https://github.com/xo/xo#xo) to generate idiomatic code for different languages code based on a database schema.

## Quickstart

Install XO

```console
$ go install github.com/xo/xo@latest
```

Create a folder to store XO models in the app root (usually named as models)

```console
$ mkdir -p models
```

Generate code from your database schema. Default output folder is models. Use `-o folderName` flag if your models folder has a different name.

```console
$ xo schema yourDatabaseConnectionString/yourDatabaseName --src xo-tpl
```

> The `--src xo-tpl` flag parameter will provide the XO template that only allows Read operations from the target database.

## How it works

The application root has files according to the supported databases.

Create your select statements at the file related to your database.

To write your select statements use the functions created by XO at models folder.

Ensure to have indexes in your database to every table field you want to create a filter.
