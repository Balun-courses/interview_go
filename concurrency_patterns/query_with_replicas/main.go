package main

type Database interface {
	Query(query string) string // simple interface for example
}

func DistributedQuery(replicas []Database, query string) string {
	responseCh := make(chan string, 1)
	for _, replica := range replicas {
		go func() {
			select {
			case responseCh <- replica.Query(query):
			default:
				return
			}
		}()
	}

	return <-responseCh
}

func main() {
	replicas := []*PgSQLDatabase{
		NewPgSQLDatabase("127.0.0.1:5432"),
		NewPgSQLDatabase("127.0.0.2:5432"),
		NewPgSQLDatabase("127.0.0.3:5432"),
	}

	response := DistributedQuery(replicas, "query to pgsql...")
	_ = response
}
