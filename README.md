# GoMail

## What is GoMail?
It is a simple and fast email API written in Go! It is my first experience with Go language, there is nothing special about this project haha

## How to run
- Clone the repository
- Add .env file with the following variables:
    ```
    SMTP_HOST=smtp.gmail.com
    SMTP_PORT=<SMTP PORT> (USUALLY 587)
    SMTP_USER=<YOUR SENDER EMAIL>
    SMTP_PASS=<YOUR APP PASSWORD>
    API_PORT=<PORT YOU WANT TO USE THE API ON> (8080 is default)
    ```
    > **Note**: see [App Password](https://myaccount.google.com/apppasswords)

- Run with `go run .` (or `go run main.go`)

## How to use 
- Send a POST request to `http://localhost:8080/send-email` with the following JSON body structure:
    ```json
    {
        "to": "<receiver_email>"
        "subject": "<title_of_the_email>",
        "body": "body_of_the_email"
    }
    ```