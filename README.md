# ofriends

[![Join the chat at https://gitter.im/TheFlies/ofriends](https://badges.gitter.im/TheFlies/ofriends.svg)](https://gitter.im/TheFlies/ofriends?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

Tool to track guests visit to my hotels over the time

## Development
```
git clone https://github.com/TheFlies/ofriends.git

# Run with Docker compose
make compose
```
### Web development
```
# start mongo service (port 27017) using docker
docker run -p 27017:27017 -v /opt/data/mongo_home:/data/db --name mongo -d mongo:latest

# start api service (port 8080)
go run main.go
```
open other terminal then start the web in development mode (port 3000)
```
cd web
yarn serve
```
## Tech requirements:
- Authentication:
  + LDAP (AD) (backend authentication)
  + JWT (frontend authentication)
- Database
  + MongoDB
- Frontend
  + Full frontend framework [Vue](https://vuejs.org)
  + Component toolkit: [Element-UI](https://element.eleme.io)

## TODO
Need update
