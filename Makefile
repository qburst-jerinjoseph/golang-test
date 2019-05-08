# Contains all build and execution commands for lazy-go 

up:
	docker-compose build 
	docker-compose up -d

db-up: 
	docker-compose up --detach db 

down: 
	docker-compose down 

log:
	docker-compose logs --tail 100 -f lazy-go
db-log:
	docker-compose logs --tail 100 -f db

pgweb:			
	docker-compose up -d pgweb  