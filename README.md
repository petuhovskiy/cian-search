# cian-search

## Usage

```
go run main.go [-query=query.json] [-result=result.json]
```

The command takes query.json as a query, fetches and saves result to result.json.

## Pages

If request contains $page instead of page number, it gets replaced by pages 1 to +inf, until empty page is received. 

## Output format

Output format is plain json array, with json objects. Single array element represents single offer.
