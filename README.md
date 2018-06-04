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

# Usage

## Create short url
```
curl -X POST -F "url=http://www.google.com" http://localhost:8081/shorten
```
Output  
```
{"url":"localhost:8081/3"}
```

## Redirect
```
curl -v localhost:8081/3
```
Output  
```
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 8081 (#0)
> GET /3 HTTP/1.1
> Host: localhost:8081
> User-Agent: curl/7.54.0
> Accept: */*
>
< HTTP/1.1 303 See Other
< Location: http://www.google.com
< Date: Mon, 04 Jun 2018 08:03:13 GMT
< Content-Length: 48
< Content-Type: text/html; charset=utf-8
<
<a href="http://www.google.com">See Other</a>.
```

# Licence 
This module is open-sourced software licensed under the [MIT license](http://opensource.org/licenses/MIT)