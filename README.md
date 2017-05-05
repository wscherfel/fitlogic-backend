# fitlogic-backend
Backend for FitLogic product, project for MPR subject on FIT VUT summer semester 2016/2017.

## Usage
If you only wish to run server, simply execute binary on any linux system.
These properties can be set without recompiling server:
- Port the server runs on
- Secret that is used to create and decode JWTs
- TimeFormat that is used in Projects and Risks

If you wish to make changes to code you have to have Go set up and the project saved in the right path ($GOPATH/github.com/wscherfel/fitlogic-backend) otherwise imports won't work.

## Project structure

### Package controllers
This package contains controllers. Each controller has its own structure (e.g. `UserController`) and its methods are handlers of endpoints. There are 4 controllers present:
- UserController handles users endpoints
- ProjectController handles projects endpoints
- RiskController handles risks endpoints
- CmController handles countermeasures endpoints - currently not used

### Package common
This package contains returned errors, types (e.g. `IDsRequest`) and functions (e.g. working with JWTs) used in all controllers.

### Package models
This package contains models for DB.

There are 4 models present:
- User
- Project
- Risk
- CounterMeasure which is currently not used.

### Package access
This package contains data access objects for each of models. It is represented by a structure named `{ModelName}DAO`.

### Package cmd
This package contains `main` function. Routing of endpoints and connection to DB is done here.

## Project compilation
Compile project using those commands:

`cd $GOPATH/github.com/wscherfel/fitlogic-backend/`

`go build ./cmd/fitlogic`
