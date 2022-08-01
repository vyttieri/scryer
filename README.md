# Scryer
"One who divines, sees or predicts the future by means of a scrying tool; especially a crystal ball."
![Scryer](scryer.png)

### TODO
- Make `golangci-lint` run properly in SublimeText
- Error Handling (api, application, frontend); Data Validation (db, application, frontend)
- Multiple `.gitignore` business
- How to integrate JS dev server and Go? (Currently using Vite proxy request)
- ~Pool or Persist MySQL Client~; How does this work with multiple connections?

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
    - Groupings?
6. ~Set up user login~
    - ~JWT authentication on backend~
7. ~Create frontend views for registration/login~
    - ~Switch to using sessions. simpler.~
75.
    - rejig frontend device preferences:
    - - ~visible~
    - - icon
    - - sortPosition
8. ~Save user preferences on back-end~
    - ~Create user preferences table~
    - ~Send user preferences along with register action, create with user~
    - ~Send user preferences back with login action~
    - ~UPDATE user preferences on change~
9. Prettify frontend
    - draggable/sort
    - icon selector
    - modal for login/register forms
    - Layout
10. Code Cleanup

### Structure

### Setup
- `npm install`
- `ONESTEPGPS_API_KEY` needs to be set in your local `ENV`
    - `export ONESTEPGPS_API_KEY=[API_KEY]`
- `SECRET` needs to be set in your local `ENV`; `head -c20 /dev/urandom | base64`
    - `export JWT_SECRET_KEY=[SECRET_KEY]`
- `VITE_GOOGLE_MAPS_API_KEY` needs to be set in `scryer-frontend/.env`
- In MySQL: `CREATE DATABASE scryer;`; `CREATE USER 'scryer'@'localhost' IDENTIFIED BY 'onestepgpsr00lz';` `GRANT ALL PRIVILEGES ON scryer.* TO 'scryer'@'localhost' WITH GRANT OPTION;`

### Limitations
- Sending password over HTTP isn't great, would like to use SSL
- JWT stored in localStorage is vulnerable to malicious JS. Worth considering other methods of storage.

### Improvements
- It would be great to define a more generalized API in which the client could request the specific fields they wanted to retrieve from the API. Since this project is limited in scope I opted to build an endpoint that serves exactly the data the client needs.
