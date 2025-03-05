# Go Email Service

This is an email sending service in Go. It allows you to send emails with a subject and message using the command line.

## Requirements

- Go 1.18 or higher

## Installation

1. Clone the repository and go to the project directory:

    ```sh
    git clone https://github.com/rikidas/go-email-service.git
    cd go-email-service
    ```

2. Initialize the Go modules:

    ```sh
    go mod tidy
    ```

## Configuration

1. Create a [.env](http://_vscodecontentref_/1) file in the root of the project with the following variables:

    ```env
    SMTP_HOST=smtp.example.com
    SMTP_PORT=587
    SMTP_USER=your_user
    SMTP_PASS=your_password
    LOG_FILENAME=logs.log
    ```

2. Make sure the [.env](http://_vscodecontentref_/2) file is in the [.gitignore](http://_vscodecontentref_/3) to avoid uploading your credentials to the repository.

## Linux deploy

Run this commands on powershell

```
$env:GOOS="linux"
$env:GOARCH="amd64"
go build -o email-service cmd/main.go
```

For Raspberry pi 3 or before

```
$env:GOOS="linux"; $env:GOARCH="arm"; $env:GOARM="7"; go build -o email-service main.go

```

For Windows

```
$env:GOOS="windows"; $env:GOARCH="amd64"; go build -o email-service.exe main.go

```


## Usage

Only log
```sh
./email-service --message="Este es un log" --mode=log --log_filename="path to save log"
```
Only email:

```sh
./email-service --email="correo@ejemplo.com" --subject="Alerta" --message="Mensaje importante" --mode=mail --ruta_env="path_to .env" --log_filename="path to save log"
```
Both

```sh
./email-service --email="correo@ejemplo.com" --subject="Alerta" --message="Mensaje importante" --mode=both --ruta_env="path_to .env" --log_filename="path to save log"

```


## Recomendations
Before using it with a Gmail account, configure the account to use Gmail's SMTP.

Also, make sure port 587 is open on your machine. (Some cloud providers block it to prevent SPAM).