# catalyst

I use existing libs :

 - Ozzo Validation, for input request validation
 - Godotenv, for env loader
 - jmoiron/sqlx for postgres driver


# For setup after cloning the repo:
> cd catalyst
> go mod tidy

# to do a unit test :
> go to the each usecase package you want to testing then run a command "go test"
> you can see the coverage testing in each usecase package by open the project with vscode, choose the testing file, right click then choose "Go:Toogle Test Coverage in Current Package"

# summary of unit test 
I have done the unit test and here are the results :
>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-gopDfYix/go-code-cover catalyst/src/app/use_cases/brand
>>ok  	catalyst/src/app/use_cases/brand	0.622s	coverage: 100.0% of statements

>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-gopDfYix/go-code-cover catalyst/src/app/use_cases/product
>>ok  	catalyst/src/app/use_cases/product	0.156s	coverage: 100.0% of statements

>Running tool: /usr/local/go/bin/go test -timeout 30s -coverprofile=/var/folders/h_/tjhvlj3n32sc9lvvfbm8x9ym0000gn/T/vscode-gopDfYix/go-code-cover catalyst/src/app/use_cases/order
>>ok  	catalyst/src/app/use_cases/order	1.090s	coverage: 100.0% of statements


# for db table :
>> in folder db, there is a .sql file with the create table command and insert command. I use postgresql for this case. you can run the command in your sql editor page.
>> if you running this project without docker, just make a connection into your localhost
>> then make new schema in db "projek", named "catalyst", after that run all the command in .sql file from brands, products, orders, then order_details
>> if you running this project with docker, make a connection into 0.0.0.0 and make a database e.g I use projek
>> then make new schema in db "projek", named "catalyst", after that run all the command in .sql file from brands, products, orders, then order_details

# the endpoint
here is the curl for the endpoint :
>curl --location --request POST 'http://localhost:8080/api/brand' \
--header 'Content-Type: application/json' \
--data-raw '{
 "name" : "brand3"
}'

>curl --location --request POST 'http://localhost:8080/api/product' \
--header 'Content-Type: application/json' \
--data-raw '{
 "name" : "p5",
 "price" : 1000,
 "brand_id" : 3
}'

>curl --location --request POST 'http://localhost:8080/api/order' \
--header 'Content-Type: application/json' \
--data-raw '{
    "data": [
        {
            "product_id": 1,
            "price": 1000,
            "qty": 2
        },
        {
            "product_id": 3,
            "price": 1000,
            "qty": 2
        }
    ]
}'

>curl --location --request GET 'http://localhost:8080/api/product?id=1' \
--header 'Content-Type: application/json' \
--data-raw '{
 "name" : "p3",
 "price" : 1000,
 "brand_id" : 4
}'

>curl --location --request GET 'http://localhost:8080/api/product/brand?id=3' \
--header 'Content-Type: application/json' \
--data-raw '{
 "name" : "p3",
 "price" : 1000,
 "brand_id" : 4
}'

>curl --location --request GET 'http://localhost:8080/api/order/2' \
--data-raw ''



> here is the postman link if you want to use postman instead : 
>> https://www.getpostman.com/collections/cd44c44140a288415043

# to run the project
if you're not using docker, just set the .env file with yoyr database credential, then cd catalyst, do go run main.go

# to run the project with docker
after clone and do some set up that explained before, do this following actions :
- set database credential in .env

- in this part :
>> DB_HOST=database (recommend to literally use "database" according to the docker-compose.yaml)
>> DB_PORT=5432  
>> DB_NAME=projek/your_db_name
>> DB_USERNAME=your_db_user
>> DB_PASSWORD=your_db_password
>> DB_SCHEMA=warung_pintar
>> DB_SSL_MODE=disable

- in this part :
>> POSTGRES_USER=postgres
>> POSTGRES_PASSWORD=your_db_password
>> POSTGRES_DB=projek
 
- cd catalyst, docker-compose up
- go to you postgresql db editor (pgAdmin, etc)
- make a new connection to 0.0.0.0
- make a new database, in this project I make a db named "projek"
- in db "projek" make a new schema named "catalyst"
- do all command to make the table and insert, you can see the command in db/account.sql and db/customer.sql
- project ready to use
