# Basic golang service

## Description

It's a service that shows random quotes, allows to like them and show quote that is pretty similar to a specified one

## How to run

1. Clone this repo

    ```shell
    git clone <url>
    ```

2. Go to the cloned folder

    ```shell
    cd quotes
    ```

3. Run `go mod tidy` command
4. Copy the [.env.example](.env.example) file and change variables to what you need
5. Run the following command:

    ```shell
    go run main.go
    ```

6. Or you can run `Run` configuration if you are using `Goland`
7. Or you can run the following taskfile command:

    ```shell
    task run
    ```

8. Import [Quotes](./postman/Quotes.postman_collection.json) collection into Postman program
9. Send request throw Postman
