# Sample Echo with GORM



## Configurations
Configuration used gonfig. All configs are declared in `config/config.json`

## Architecture
| Folder | Details |
| --- | ---|
| api | Holds the api endpoints |
| db | Database Initializer and DB manager |
| route | router setup |
| model | Models|



## Run 
`go run server.go`

## Create Token -> Save Data to Redis -> Using Refresh Token To Logout and Update is_login


