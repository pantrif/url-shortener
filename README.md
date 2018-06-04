# url-shortener
A golang URL Shortener with mysql support.  
Using Bijective conversion between natural numbers (IDs) and short strings

# Installation
## Using docker compose
```
docker-compose up --build
```
## Using an existing mysql

Edit .env file to add connection strings for mysql  
Run mysql_init/create_table.sql  
```
go run main.go
```



# Licence 
This module is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)
