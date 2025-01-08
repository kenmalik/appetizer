build:
	@go build -o appetizer .

reset:
	@if [ -f data.db ]; then rm data.db; fi
	@go build -o appetizer .

mock:
	@if [ -f data.db ]; then rm data.db; fi
	@sqlite3 data.db < ./scripts/init.sql
	@sqlite3 data.db < ./scripts/mock-data.sql
