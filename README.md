## Before using this app, you should to create `.env` file, and specify `DATABASE_FILE_PATH` and `RESULT_FILE_PATH`.
### When you did this, you should follow this steps:

### Install dependencies

You must be in the project folder

```bash
go mod tidy
```

### Build app

You should build it once, then you can run the already built app

```shell
go build ./cmd/main.go
```

or for windows

```shell
go build .\cmd\main.go
```

### Run app

All done, now you can launch app

```shell
./main
```

or for windows 

```shell
.\main.exe
```
