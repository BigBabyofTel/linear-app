## Backend 

### Start database container 
Makefile should work on the macbook and linux without any problems, on the windows maybe it will be necessary to install make and docker-compose by yourself.
```bash
make db-container
# or
docker-compose up -d
```
### Stop database container 
```bash
make db-container-stop 
# or
docker-compose down
```