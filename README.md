# Url Shortener

This is just a code repository to train some Software Engineer skills.

The logic to create the URL shortener probably is not the best one, but I'm not concerned about it at this first moment.


# Running the project

## API

After cloning the project, navigate to the folder `src/api`. Creates a copy of the file `local.env.default` and rename the copied file to `local.env`. Fulfill the variables inside it with your local values. Maybe will be there some default values, if they don't match your local environment settings, you can change those values to match your environment. 

### Configuring the database

To run the application, we need a mysql database.
If there is one currently running, just take the connection string and fill it in the `local.env` file.

If don't have a mysql running, you can start one using the following command inside the `src/api` folder:

`docker-compose -f docker-stacks/mysql.yml` up -d

It will start a mysql using docker, and binding the container to the port 3306 in your machine.

### API database migrations 

After you have your database running, you can try these commands below (inside the `src/api` folder):

- `tools/migration_runner.go migrate up -c {connection string}`
  It will run the migrations

- `tools/migration_runner.go migrate down -c {connection string} -s {number of the last migrations to roolback}`
  It will rollback the `s` last migrations. The `s` parameter must be a number. If it's not provided, or filled with `0`, all migrations will be rollbacked.

- `tools/migration_runner.go migrate create -n {migration name}`
  It will create a empty migration file inside the `migrations` folder.  When creating a new migrations, make sure that they are idempotent.

