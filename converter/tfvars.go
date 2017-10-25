package converter

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/hashicorp/hcl"
	"encoding/json"
	"github.com/hashicorp/hcl/hcl/printer"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

type TFVars interface {
	ToJson()error
	ToHCL()error
}

type TFVarsManager struct {

}

func (tfg TFVarsManager)ToJSON() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	var v interface{}
	err = hcl.Unmarshal(input, &v)
	if err != nil {
		return fmt.Errorf("unable to parse HCL: %s", err)
	}

	json, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to marshal json: %s", err)
	}

	fmt.Println(string(json))

	return nil
}

func (tfg TFVarsManager)ToHCL() error {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("unable to read from stdin: %s", err)
	}

	ast, err := jsonParser.Parse([]byte(input))
	if err != nil {
		return fmt.Errorf("unable to parse JSON: %s", err)
	}

	err = printer.Fprint(os.Stdout, ast)
	if err != nil {
		return fmt.Errorf("unable to print HCL: %s", err)
	}

	return nil
}
