// Copyright 2023-2024 Princess Beef Heavy Industries, LLC / Dave Shanley
// https://pb33f.io

package openapi

import (
	"fmt"
	"testing"

	"github.com/daveshanley/vacuum/model"
	drModel "github.com/pb33f/doctor/model"
	"github.com/pb33f/libopenapi"
	"github.com/stretchr/testify/assert"
)

func TestExamplesSchema(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          type: string
      examples:
        - id: smoked`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)

}

func TestExamplesSchema_TrainTravel(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Station:
      type: object
      properties:
        id:
          type: string
          format: uuid
          examples:
            - efdbb9d1-02c2-4bc3-afb7-6788d8782b1e
            - b2e783e1-c824-4d63-b37a-d8d698862f1d`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)

}

func TestExamplesSchema_Invalid(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          type: string
      additionalProperties: false
      examples:
        - id: smoked
          name: illegal`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "additional properties 'name' not allowed", res[0].Message)
	assert.Equal(t, "$.components.schemas['Herbs'].examples[0]", res[0].Path)

}

func TestExamplesSchema_Valid_OneOf(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          oneOf:
            - type: string
              const: smoked
            - type: integer
              const: 1
      examples:
        - id: smoked`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)

}

func TestExamplesSchema_Valid_OneOf_Int(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          oneOf:
            - type: string
              const: smoked
            - type: integer
              const: 1
      examples:
        - id: 1`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)

}

func TestExamplesSchema_Invalid_OneOf(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          oneOf:
            - type: string
              const: smoked
            - type: integer
              const: 1
      examples:
        - id: eaten`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 2)
	assert.Equal(t, "value must be 'smoked'", res[0].Message)
	assert.Equal(t, "$.components.schemas['Herbs'].examples[0]", res[0].Path)
	assert.Equal(t, "got string, want integer", res[1].Message)
	assert.Equal(t, "$.components.schemas['Herbs'].examples[0]", res[01].Path)

}

func TestExamplesSchema_ExampleProp(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          oneOf:
            - type: string
              const: smoked
            - type: integer
              const: 1
      example:
        id: smoked`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)

}

func TestExamplesSchema_ExampleProp_Failed(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    Herbs:
      type: object
      properties:
        id:
          oneOf:
            - type: string
              const: smoked
            - type: integer
              const: 1
      example:
        id: baked`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 2)
	assert.Equal(t, "value must be 'smoked'", res[0].Message)
	assert.Equal(t, "$.components.schemas['Herbs'].example", res[0].Path)
	assert.Equal(t, "got string, want integer", res[1].Message)
	assert.Equal(t, "$.components.schemas['Herbs'].example", res[1].Path)

}

func TestExamplesSchema_Param_Valid(t *testing.T) {
	yml := `openapi: 3.1
components:
  parameters:
    Herbs:
      in: header
      name: herbs
      schema:
        type: object
        properties:
          id:
            type: string
            const: spicy
      examples:
        - id: spicy`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)

}

func TestExamplesSchema_Param_Invalid(t *testing.T) {
	yml := `openapi: 3.1
components:
  parameters:
    Herbs:
      in: header
      name: herbs
      schema:
        type: object
        properties:
          id:
            type: string
            const: spicy
      examples:
        sammich:
          value:
            id: crispy`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "value must be 'spicy'", res[0].Message)
	assert.Equal(t, "$.components.parameters['Herbs'].examples['sammich']", res[0].Path)

}

func TestExamplesSchema_Header_Invalid(t *testing.T) {
	yml := `openapi: 3.1
paths:
  /herbs:
    get:
      responses:
        "200":
          headers:
            "Herbs":
              schema:
                type: string
                const: tasty
              examples:
                sammich:
                  value: crispy
                  
      `

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "value must be 'tasty'", res[0].Message)
	assert.Equal(t, "$.paths['/herbs'].get.responses['200'].headers['Herbs'].examples['sammich']", res[0].Path)

}

func TestExamplesSchema_MT_Invalid(t *testing.T) {
	yml := `openapi: 3.1
paths:
  /herbs:
    get:
      responses:
        "200":
          content:
            application/json:
              schema:
                type: string
                const: tasty
              examples:
                sammich:
                  value: crispy
                  
      `

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 1)
	assert.Equal(t, "value must be 'tasty'", res[0].Message)
	assert.Equal(t, "$.paths['/herbs'].get.responses['200'].content['application/json'].examples['sammich']", res[0].Path)

}

/*
components:
  schemas:
    Test:
      type: array
      description: Test array with numbers
      items:
        type: number
      example:
        - 0 # <- This gives a warning
        - 0
        - 0
*/

func TestExamplesSchema_HandleJSONTime(t *testing.T) {
	yml := `openapi: 3.1
components:
  schemas:
    badDate:
      type: string
      description: a bad time.
      format: date-time
      example: 2022-08-07T12:12:00Z`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)
}

// https://github.com/daveshanley/vacuum/issues/615
func TestExamplesSchema_HandleArrays(t *testing.T) {
	yml := `openapi: 3.1.0
components:
  schemas: 
    Test:
      type: array
      description: Test array with numbers
      items:
        type: number
      example:          
        - 0
        - 0
        - 0`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)
}

func TestExamplesSchema_3_1_0_nullable(t *testing.T) {
	yml := `openapi: 3.1.0
components:
  schemas: 
    Test:
      type: object
      description: Test OpenAPI 3.1 nullable field using type array
      required:
        - nullable_prop
      properties:
        nullable_prop:
          type: ["string", "null"]
          nullable: true
      examples:
        - nullable_prop: string val
        - nullable_prop: null`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)
}

func TestExamplesSchema_3_0_0_nullable(t *testing.T) {
	yml := `openapi: 3.0.0
components:
  schemas: 
    Test:
      type: object
      description: Test OpenAPI 3.0 nullable field
      required:
        - nullable_prop
      properties:
        nullable_prop:
          type: string
          nullable: true
      examples:
        - nullable_prop: string val
        - nullable_prop: null`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}

	m, _ := document.BuildV3Model()
	path := "$"

	drDocument := drModel.NewDrDocument(m)

	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)

	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule

	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	assert.Len(t, res, 0)
}

func TestExamplesSchema_ErrorMessages_examples(t *testing.T) {
	yml := `openapi: 3.0.0
components:
  schemas:
    ErrorMessages:
      description: This will contain error object with HTTP responseCode and array of error objects messages
      properties:
        errorDetails:
          items:
            properties:
              code:
                description: Error message code
                type: string
              detail:
                description: Detailed error description to find out why the error occurred
                nullable: true
                type: string
              message:
                description: Information of the error Occurred
                type: string
            type: object
          type: array
        message:
          type: string
        responseCode:
          description: Error code
          example: 400
          format: int32
          maximum: 600
          minimum: 100
          type: integer
      type: object
      examples:
        - errorDetails:
            - code: "AP103"
              detail: "agent.generalAgency.nationalProducerNumber"
              message: "nationalProducerNumber should be <= 10 characters in length"
          message: "Request Payload Validation Error Occurred"
          responseCode: 400
        - errorDetails:
            - code: "GE004"
              detail: null
              message: "Invalid credentials provided"
          message: "Unauthorized"
          responseCode: 401`

	document, err := libopenapi.NewDocument([]byte(yml))
	if err != nil {
		panic(fmt.Sprintf("cannot create new document: %e", err))
	}
	m, _ := document.BuildV3Model()
	path := "$"
	drDocument := drModel.NewDrDocument(m)
	rule := buildOpenApiTestRuleAction(path, "examples_schema", "", nil)
	ctx := buildOpenApiTestContext(model.CastToRuleAction(rule.Then), nil)
	ctx.Document = document
	ctx.DrDocument = drDocument
	ctx.Rule = &rule
	def := ExamplesSchema{}
	res := def.RunRule(nil, ctx)

	// Print any validation errors for debugging
	for _, result := range res {
		t.Logf("Validation error: %s", result.Message)
	}

	assert.Len(t, res, 0)
}
