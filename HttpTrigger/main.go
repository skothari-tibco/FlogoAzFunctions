package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Azure/azure-functions-go/azfunc"
	_ "github.com/Azure/azure-functions-go/sample/HttpTrigger/shim"
	fl "github.com/skothari-tibco/flogoaztrigger"
)

// Run runs this Azure Function because it is specified in `function.json` as
// the entryPoint. Fields of the function's parameters are also bound to
// incoming and outgoing event properties as specified in `function.json`.

func Run(ctx azfunc.Context, req *http.Request) (*User, error) {

	// additional properties are bound to ctx by Azure Functions
	ctx.Log(azfunc.LogInformation, "function invoked: functionID %v, invocationID %v", ctx.FunctionID(), ctx.InvocationID())

	// use standard library to handle incoming request
	body, _ := ioutil.ReadAll(req.Body)

	// deserialize JSON content
	var data map[string]interface{}
	var err error
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal json: %s", err)
	}

	results, err := fl.Invoke()

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The result of flogo is ", results)
	// get query param values
	name := req.URL.Query().Get("name")

	if name == "" {
		return nil, fmt.Errorf("missing required query parameter: name")
	}

	u := &User{
		Name:     results,
		Greeting: fmt.Sprintf("Hello %s. %s\n", name, data["greeting"].(string)),
	}

	return u, nil
}

// User exemplifies a struct to be returned. You can use any struct or *struct.
type User struct {
	Name     string
	Greeting string
}
