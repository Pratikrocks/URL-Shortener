# URL-Shortener
URL Shortener written in Golang using Redis

## Installation 

* Using docker-compose
```azure
$ git clone git@github.com:Pratikrocks/URL-Shortener.git
$ cd URL-Shortener
$ docker-compose up

# This will start the redis and the server container with
# server listening to port 8082 and redis listening to port 6379
```

## API Usage
* To encode an url to a short url make a `POST` request `/info`
 endpoint with the url as the body.

```azure
$ curl --location --request POST 'localhost:8082/info' \                                  
--header 'Content-Type: application/json' \
--data-raw '{
"url": "https://leetcode.com/problemset/",
"expires": 1234
}'

```
`url` is the url to be encoded.
`expires` is the time in seconds after which the url will expire. [*TBD*]

In the response you will get the short url hash code of length 5.
For example, if the url is `https://leetcode.com/problemset/` then the
short url code corresponding to it might be `gb2A1`

* To decode a short url make a `GET` request to `/view/{:url_hash}` endpoint.
* Corresponding to this example to get redirected to the original url open any browser and make a 
`GET` request to `/view/gb2A1` with full url endpoints as
``localhost:8082/view/gb2A1`` which will get redirected to the original url.