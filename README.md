# go-quote-service

The Go quote service is a super simple REST API service that manages a list of quotes.

You can add quotes and get all the quotes. That's it. 

# Usage

Run the code:

```
go run main.go
```

Hit it with your favorite HTTP client. I prefer [httpie]()

Here is how to add a quote:
```
http http://localhost:7777/quotes quote="We must be very careful when we give advice to younger people: sometimes they follow it! ~ Edsger W. Dijkstra" 

HTTP/1.1 200 OK
Content-Length: 0
Date: Sun, 24 Nov 2024 17:03:18 GMT
```

Here is how to get all quotes:

```
http http://localhost:7777/quotes                                                                                                                      

HTTP/1.1 200 OK
Content-Length: 131
Content-Type: application/json
Date: Sun, 24 Nov 2024 17:03:24 GMT

[
    "We must be very careful when we give advice to younger people: sometimes they follow it! ~ Edsger W. Dijkstra"
]
```


## In-memory

For quick testing you can use the in-memory mode. Just run the code and you're done.

## With Postgres

If you want to store the quotes in a database you need to configure a connection string.

Define the environment variable GO_QUOTE_SERVICE_CONNECTION_STRING in the format:
```
host=<host URL> port=<port> user=<user> password=<password> dbname=<db name>> sslmode=<ssl mode>
```

# Deploy to Render

This project was developed to support an article I can't find anymore about Render.

Go to your dashboard and add a new web service.

See https://render.com/docs/github

