# alpha-quad-sim
> Small climate simulator of a particular 2-d 3-body restricted system made in Go (and hosted in Google App Engine)

It makes the following assumptions:
* The initial state is dry, since planets are aligned with the sun
* Collinearity of the 3 points in an orbital plane occurs then the area is equal to 0
* If the 3 planets are collinear and the Sun and the Sun can form a new triangle with 2 points with no area, then the Sun is also collinear with them

## Development

Prerequisites: Go 1.6 and the Google App Engine SDK are required. Also, this solution is made to be deployed into a Google App Engine project along with a Google Cloud MySQL instance. Both the solution and the database instance should belong to the same project

Other considerations are the following:
* New files should be added at the root level because App Engine is very picky with the folder structure on imports so the solution won't compile in the cloud otherwise
* The repository has to be in the **$GOPATH/src** directory
* The job that initializes the database with newly calculated values can be triggered with the job in **cron.yaml**. The job only populates the database and doesn't delete previous values
* The database connection can be configured in **app.yaml**
* Tests can be run with **go test**

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
## Deployment
Deploy it to a Go App Engine project:
```
gcloud app deploy
```

## API

The live version can be found at https://superb-webbing-163902.appspot.com and the current api usages are the following:

* https://superb-webbing-163902.appspot.com/clima/566
* https://superb-webbing-163902.appspot.com/clima/?dia=566
