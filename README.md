# GO Web API - Opportunities

## About

This project is based on the Create, Read, Update and Delete (CRUD) processes to serve as an object of study for the development of an Application Programming Interface (API) using the `Go` programming language and the `Gin Framework`

### Why `Go`?

`Go` The go language is very performant and has an easy-to-use syntax, It is a perfect option for developing Back-End applications, in addition to having good performance, it also has a low consumption cost compared to languages like `Java`


#### Endpoints:

- api/v1/openings - GET(returns a list of all opportunities)
- api/v1/opening?id= - GET(returns information about a specific ID)
- api/v1/opening - POST(creates a new opportunity)
- api/v1/opening/{id} - DELETE(delete an opportunity using the ID)
- api/v1/opening - PUT(update opportunity information)


all GET methods return a JSON file as in the example below

```
{
    "ID": 1, //int
	"CreatedAt": "05.06.23",//datetime
	"UpdatedAt": "09.06.23",//datetime
	"DeletedAt": "null", //datetime
	"Role": "role 1", //string 
	"Company":  "company name", //string
	"Location":  "company location", //string
	"Link":      "opportunity link", //string
	"Remote":   "false", //boolean
	"Salary": "1.000" //double
}
```
