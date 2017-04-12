# alpha-quad-sim
> Small climate simulator of a particular 2-d 3-body restricted system made in Go (and hosted in Google App Engine)

It makes the following assumptions:
* The initial state is dry, since planets are aligned with the sun
* Collinearity of the 3 points in an orbital plane occurs then the area is equal to 0
* If the 3 planets are collinear and the Sun and the Sun can form a new triangle with 2 points with no area, then the Sun is also collinear with them
* It takes the continuation of several continued days with the same climate as a period. With each climate transition it increments the period number of that climate

If the area is *tending* to 0 then the climate is on it's peak rain point. We explain it with the following heuristic:
1. The largest possible side is 3000, given by the line between the most distant opposite planets V and B
2. The largest inscribed triangle inside a circle is the equilateral triangle, and considering we can’t construct one with these orbits, the largest one for this case will be an isosceles triangle
3. Following points i and ii and solving via Pythagoras theorem, we can deduct that the largest perimeter tends (but it's not equal) to 6000. This is because the point at which the triangle becomes isosceles is when the area tends to 0 and the sides converge into a line


## Development

Prerequisites: Go 1.6 and the Google App Engine SDK are required. Also, this solution is made to be deployed into a Google App Engine project along with a Google Cloud MySQL instance. Both the solution and the database instance should belong to the same project

Other considerations are the following:
* The job that initializes the database with newly calculated values can be triggered with the job in **cron.yaml**. The job only populates the database and doesn't delete previous values
* The database connection can be configured in **app.yaml**

## Database
The MySQL database connection can be configured via the environment variables in **app.yaml**

The database name is *climateregistry* and the only table used is *climates*. It's structure is the folloing:
```
+-------------+-------------+------+-----+---------+----------------+
| Field       | Type        | Null | Key | Default | Extra          |
+-------------+-------------+------+-----+---------+----------------+
| climate     | varchar(20) | YES  |     | NULL    |                |
| climate_day | int(11)     | YES  |     | NULL    |                |
| id          | int(11)     | NO   | PRI | NULL    | auto_increment |
+-------------+-------------+------+-----+---------+----------------+
```
## Deployment/Testing
The repository has to be in the **$GOPATH/src** directory

Build:
```
go build
```
Unit tests can be run with the following command:
```
go test
```
Deploy it to a Go App Engine project:
```
gcloud app deploy
```

The app can be run locally, allowing to test the api locally via localhost:3000/8080. It will also print stats for the default duration of 10 years:
```
./alpha-quad-sim offline
```

## API

The live version can be found at https://superb-webbing-163902.appspot.com and the current api usages are the following:

* https://superb-webbing-163902.appspot.com/clima/566
* https://superb-webbing-163902.appspot.com/clima?dia=566
* https://superb-webbing-163902.appspot.com/clima/?dia=566
