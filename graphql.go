package main

import (
	"log"
	"net/http"
	"reflect"

	"github.com/aerogo/aero"
	"github.com/akyoto/color"
	"github.com/animenotifier/arn"
	"github.com/animenotifier/notify.moe/utils/gql"
	"github.com/graphql-go/graphql"
)

func init() {
	rootQueryFields := graphql.Fields{}

	for name, typ := range arn.DB.Types() {
		if typ.Kind() != reflect.Struct {
			continue
		}

		// Bind name for the closure
		typeName := name

		rootQueryFields[typeName] = &graphql.Field{
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{
					Type:         graphql.String,
					DefaultValue: "",
				},
			},
			Type: graphql.NewObject(graphql.ObjectConfig{
				Name:   typeName,
				Fields: gql.BindFields(typ),
			}),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)
				return arn.DB.Get(typeName, id)
			},
		}
	}

	// Schema
	rootQuery := graphql.ObjectConfig{
		Name:   "RootQuery",
		Fields: rootQueryFields,
	}

	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(rootQuery),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	app.Post("/graphql", func(ctx *aero.Context) string {
		ctx.Response().Header().Set("Access-Control-Allow-Origin", "*")
		body, err := ctx.Request().Body().JSONObject()

		if err != nil {
			return ctx.Error(http.StatusBadRequest, "Expected JSON data containing a query and variables", err)
		}

		query := body["query"].(string)
		variables := body["variables"].(map[string]interface{})

		params := graphql.Params{
			Schema:         schema,
			RequestString:  query,
			VariableValues: variables,
		}

		result := graphql.Do(params)

		if len(result.Errors) > 0 {
			color.Red("failed to execute graphql operation, errors: %+v", result.Errors)
		}

		return ctx.JSON(result)
	})
}
