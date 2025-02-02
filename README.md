Thanks to [This repo](https://github.com/sajaddp/list-of-cities-in-Iran) for provinces, cities sql.
first go in src folder, then run
``` psql -U your_username -d your_database -f ./data/seeder/provinces.sql``` to have all needed provinces.
then run ```go run /cmd/index.go``` to start project.
