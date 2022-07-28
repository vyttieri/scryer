# Scryer
"One who divines, sees or predicts the future by means of a scrying tool; especially a crystal ball."
![Scryer](scryer.png)

### TODO
- Make `golangci-lint` run properly in SublimeText
- Error handling for OneStepGPS API
- Multiple `.gitignore` business
- How to integrate JS dev server and Go? (Currently using Vite proxy request)
- Pool or Persist MySQL Client

### Run
- `go run main.go` or with air installed `air`
- `go run main.go -- migrate` to run DB migrations
- `npm run dev` runs the Vite dev server
- `mysql.server start` runs MySQL on MacOS

### Description

### Development Plan
1. ~Set up framework for Go web-server~
2. ~Write code to hit OneStepGPS API~
3. ~Initial front-end with Vue JS~
    1. ~Ingest device data into store? (or Vue equivalent)~
    2. ~List of devices w/ name, current position, and active state or drive status.~
    3. ~Display devices on Google Map~
    4. ~Add device refresh~
4. ~Add additional front-end functionality~
    - ~sort order~
    - ~hide specific devices from view~
    - ~centering~
5. ~Save user preferences on front-end~
    - SortOrder
6. Set up user login
    - JWT authentication on backend
    -
7. Save user preferences on back-end
8. Add user-uploaded icons functionality
9. Prettify frontend
10. Write some tests?

### Structure

### Setup
- `npm install`
- `ONESTEPGPS_API_KEY` needs to be set in your local `ENV`
    - `export ONESTEPGPS_API_KEY=[API_KEY]`
- `JWT_SECRET_KEY` needs to be set in your local `ENV`; recommended 32 characters long alphanumeric randomly generated key
    - `export JWT_SECRET_KEY=[SECRET_KEY]`
- `VITE_GOOGLE_MAPS_API_KEY` needs to be set in `scryer-frontend/.env`
- In MySQL: `CREATE DATABASE scryer;`; `CREATE USER 'scryer'@'localhost' IDENTIFIED BY 'onestepgpsr00lz';` `GRANT ALL PRIVILEGES ON scryer.* TO 'scryer'@'localhost' WITH GRANT OPTION;`

### Limitations

### Improvements
- It would be great to define a more generalized API in which the client could request the specific fields they wanted to retrieve from the API. Since this project is limited in scope I opted to build an endpoint that serves exactly the data the client needs.
- It would be best to break this out into two repositories.