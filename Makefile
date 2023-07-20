ruby/fizzbuzz:
	docker compose run --rm ruby ruby /src/fizzbuzz.rb

ruby/person:
	docker compose run --rm ruby ruby /src/person.rb

go/fizzbuzz:
	docker compose run --rm go go run /src/fizzbuzz.go

go/person:
	docker compose run --rm go go run /src/person.go
