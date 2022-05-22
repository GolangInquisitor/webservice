# Test web service
 
## Instruction
1. Install PostgreSQL
2. psql  -U postgres -d <base> /scripts/postgres/scoltest.sql
3. build project cmd\server_app.main.go
4. place the file config.yml next to the executable file 
5. change the settings in config.yml if necessary
6. run


## API
 - Create user  
 route: /user  
 method: POST  
 keys:  
  	"uuid"  
	"name"  
	"surname"  
	"midlename"  
	"gender"  
	"age"  
  