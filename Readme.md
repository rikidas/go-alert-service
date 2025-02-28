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

Run the service from the command line passing the necessary arguments:

```sh
./email-service --email="email@example.com" --subject="Test" --message="This is a test message"
```

## Recomendaciones
Before using it with a Gmail account, configure the account to use Gmail's SMTP.

Also, make sure port 587 is open.