# Quotes service

## Description

It's a service that shows random quotes, allows to like them and show quote that is pretty similar to a specified one

## Additional programs

1. [Taskfile](https://taskfile.dev/installation/) (Optional)
2. docker-compose or podman-compose
3. [Postman](https://www.postman.com/downloads/)

## How to run

1. Copy the [.env.example](.env.example) file to `.env` and change variables to what you need

   ```shell
   cp .env.example .env && nano .env
   ```

   1.1. Change envs in [Taskfile.yml](./Taskfile.yml) file

2. Start the postgres database
   ```shell
   export POSTGRES_ADDR=5432 POSTGRES_USER=postgres POSTGRES_PASSWORD=postgres POSTGRES_DB=quotes
   cd containers && docker-compose -f database.yml up -d && cd ..
   ```   

   Or

   ```shell
   task db
   ```

3. Run the following command:

    ```shell
    go run main.go
    ```

4. Or you can run `Run` configuration if you are using `Goland`
5. Or you can run the following taskfile command:

    ```shell
    task run
    ```

6. Import [Quotes](./postman/Quotes.postman_collection.json) collection into `Postman` program
7. Send request throw `Postman`
