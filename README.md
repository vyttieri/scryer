# Scryer
"One who divines, sees or predicts the future by means of a scrying tool; especially a crystal ball."
![Scryer](scryer.png)

### TODO
- Error Handling (api, application, frontend); Data Validation (db, application, frontend)

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
    - ~draggable/sort~
    - ~icon selector~
    - ~modal for login/register forms~
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

### Limitations/Improvements
- The OneStepGPS API endpoint is just a passthrough endpoint right now. It hits the API and then passes all that data back to the client. It would be good to tailor the endpoint to serve only the data the client needs to consume.
- Production readiness - At a minimum, different environments, using a real logger, adding a build process to compile JS.
- Password Hashing process could use, at a minimum, salts, for additional security. Ideally use a pre-built implementation that is secure.
- Better error handling. The error handling is pretty simple and takes a one-size-fits-all approach. There could be more diverse error handling depending on what went wrong.
- Better data validations. The data validations are more a proof of concept than anything else.
