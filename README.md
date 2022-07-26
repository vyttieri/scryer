# Scryer
"One who divines, sees or predicts the future by means of a scrying tool; especially a crystal ball."
![Scryer](scryer.png)

### TODO
- Make `golangci-lint` run properly in SublimeText

### Run
- `go run main.go` or with air installed `air`

### Description

### Development Plan
1. Set up framework for Go web-server
2. Write code to hit OneStepGPS API
3. Initial front-end with Vue JS
    1. Ingest device data into store? (or Vue equivalent)
    2. List of devices w/ name, current position, and active state or drive status.
    3. Display devices on Google Map
4. Add additional front-end functionality
    - sort order
    - hide specific devices from view
    - filter by name?
5. Save user preferences on front-end
    - LocalStorage? Cookies?
6. Set up user login
7. Save user preferences on back-end
8. Add user-uploaded icons functionality
9. Write some tests

### Structure

### Setup
- `ONESTEPGPS_API_KEY` needs to be set in your local `ENV`

### Limitations

### Improvements
