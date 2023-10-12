# Go-Report Producer

Using [XO](https://github.com/xo/xo#xo) to generate idiomatic code for different languages code based on a database schema.

## How to use

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
$ xo schema yourDatabaseConnectionString/yourDatabaseName
```

