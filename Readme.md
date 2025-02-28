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
    ```

2. Make sure the [.env](http://_vscodecontentref_/2) file is in the [.gitignore](http://_vscodecontentref_/3) to avoid uploading your credentials to the repository.

## Usage

Only log
```sh
./email-service --message="Este es un log" --mode=log
```
Only email:

```sh
./email-service --email="correo@ejemplo.com" --subject="Alerta" --message="Mensaje importante" --mode=mail
```
Both

```sh
./email-service --email="correo@ejemplo.com" --subject="Alerta" --message="Mensaje importante" --mode=both

```


## Recomendaciones
Before using it with a Gmail account, configure the account to use Gmail's SMTP.

Also, make sure port 587 is open.