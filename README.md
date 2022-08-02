# Scryer
"One who divines, sees or predicts the future by means of a scrying tool; especially a crystal ball."
![Scryer](scryer.png)

### Description
- Displays GPS information from OneStepGPS API in a list format as well as on a Google Maps map.
- Refreshes minutely to get up-to-date device GPS information.
- User registration/login to save device preferences.
    - Three saveable device preferences: Device visibility, device icon, and sort order (drag the devices to reorder)
    - Click the save button in top right to save preferences to database.

### Setup
1. Database:
   - A MySQL database is expected to be running with a database and database user.
       - Can be created locally:
       - `mysql -u root`
       - `CREATE DATABASE scryer_db;`
       - `CREATE USER 'scryer_user'@'localhost' IDENTIFIED BY 'onestepgpsr00lz';`
       - `GRANT ALL PRIVILEGES ON scryer_db.* TO 'scryer_user'@'localhost' WITH GRANT OPTION;`
2. Backend server:
    - `cd scryer-backend/`
    - `go install`
    - Create `.env` file (in `scryer-backend/`) with the env variables as listed in `.env_template`:
        - `ONESTEPGPS_API_KEY`
        - `SESSION_SECRET` (can generate using `cat /dev/urandom | env LC_ALL=C tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1` )
        - `DB_HOST`
        - `DB_PORT`
        - `DB_NAME`
        - `DB_USER`
        - `DB_USER_PASSWORD`
    - `go run main.go -- migrate` to run DB migrations
3. Frontend server:
    - `cd scryer-frontend/`
    - `npm install`
    - Create `.env` file (in `scryer-frontend/`) with the env variables as listed in `.env_template`:
        - `VITE_GOOGLE_MAPS_API_KEY`

### Run
- Database: `mysql.server start` (MacOS command)
- Backend: `cd scryer-backend/`; `go run main.go` or with air installed `air`
- Frontend: `cd scryer-frontend/`; `npm run dev`

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
9. ~Prettify frontend~
    - ~draggable/sort~
    - ~icon selector~
    - ~modal for login/register forms~
    - ~Layout~
10. ~Code Cleanup~

### Limitations/Improvements
- The OneStepGPS API endpoint is just a passthrough endpoint right now. It hits the API and then passes all that data back to the client. It would be good to tailor the endpoint to serve only the data the client needs to consume.
- Production readiness - At a minimum, different environments, using a real logger, adding a build process to compile JS.
- Password Hashing process could use, at a minimum, salts, for additional security. Ideally use a pre-built implementation that is secure.
- Better error handling. The error handling is pretty simple and takes a one-size-fits-all approach. There could be more diverse error handling depending on what went wrong.
- Better data validations. The data validations are more a proof of concept than anything else.
