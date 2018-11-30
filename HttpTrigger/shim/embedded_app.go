// Do not change this file, it has been generated using flogo-cli
// If you change it and rebuild the application your changes might get lost
package shim

// embedded flogo app descriptor file
const flogoJSON string = `{
  "name": "_APP_NAME_",
  "type": "flogo:app",
  "version": "0.0.1",
  "description": "My flogo application description",
  "appModel": "1.0.0",
  "imports": [
    "github.com/project-flogo/flow",
    "github.com/skothari-tibco/flogoaztrigger",
    "github.com/project-flogo/contrib/activity/actreturn",
    "github.com/project-flogo/contrib/activity/log"
  ],
  "triggers": [
    {
      "id": "my_rest_trigger",
      "ref": "github.com/skothari-tibco/flogoaztrigger",
      "handlers": [
        {
          "action": {
            "ref": "github.com/project-flogo/flow",
            "settings": {
              "flowURI": "res://flow:simple_flow"
            },
            "input": {
              "in": "=$.body"
            },
            "output" :{
              "code" : "200",
              "data": "=$.out"
            }
          }
        }
      ]
    }
  ],
  "resources": [
    {
      "id": "flow:simple_flow",
      "data": {
        "name": "simple_flow",
        "metadata": {
          "input": [
            { "name": "in", "type": "string",  "value": "test" }
          ],
          "output": [
            { "name": "out", "type": "string" }
          ]
        },
        "tasks": [
          {
            "id": "log",
            "name": "Log Message",
            "activity": {
              "ref": "github.com/project-flogo/contrib/activity/log",
              "input": {
                "message": "=$flow.in",
                "flowInfo": "false",
                "addToFlow": "false"
              }
            }
          },
          {
          	"id" :"return",
          	"name" : "Activity Return",
          	"activity":{
          		"ref" : "github.com/project-flogo/contrib/activity/actreturn",
          		"settings":{
          			"mappings":{
          				"out": "nameC"
          			}
          		}
          	}
          }
        ],
        "links": [
        	{
        		"from":"log",
        		"to":"return"
        	}
        ]
      }
    }
  ]
}
`

func init() {
	cfgJson = flogoJSON
}
