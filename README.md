# GoMail

## What is GoMail?
It is a simple and fast email API written in Go! It uses goroutines to deal with high request demands

## How to use
- Clone the repository
- Add .env file with the following variables:
    ```
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=<SMTP PORT (USUALLY 587)>
    SMTP_USER=<YOUR SENDER EMAIL>
    SMTP_PASS=<YOUR APP PASSWORD>
    API_PORT=<PORT YOU WANT TO USE THE API ON>
    ```
    > **Note**: see [App Password](https://myaccount.google.com/apppasswords)

- Run with `go run .` (or `go run main.go`)