package main

//var baseType = graphql.NewObject(graphql.ObjectConfig{
//	Name: "Base",
//	Fields: graphql.Fields{
//		"id":        &graphql.Field{Type: graphql.ID},
//		"createdAt": &graphql.Field{Type: graphql.String},
//		"updatedAt": &graphql.Field{Type: graphql.String},
//		"deletedAt": &graphql.Field{Type: graphql.String},
//	},
//})
//customerType := graphql.NewObject(graphql.ObjectConfig{
//	Name: "Customer",
//	Fields: graphql.Fields{
//		"base":        &graphql.Field{Type: baseType}, // Добавляем тип Base в качестве поля
//		"firstName":   &graphql.Field{Type: graphql.String},
//		"lastName":    &graphql.Field{Type: graphql.String},
//		"email":       &graphql.Field{Type: graphql.String},
//		"phoneNumber": &graphql.Field{Type: graphql.String},
//		"address":     &graphql.Field{Type: graphql.ID},
//	},
//})
////customerCreateType := graphql.NewObject(graphql.ObjectConfig{
////	Name: "CustomerCreate",
////	Fields: graphql.Fields{
////		"ID": &graphql.Field{Type: graphql.ID},
////	},
////})
//
//rootQuery := graphql.NewObject(graphql.ObjectConfig{
//	Name: "Query",
//	Fields: graphql.Fields{
//		"customer": &graphql.Field{
//			Type:        customerType,
//			Description: "Get customer by ID",
//			Args: graphql.FieldConfigArgument{
//				"id": &graphql.ArgumentConfig{Type: graphql.ID},
//			},
//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//				id, ok := p.Args["id"].(string)
//				if !ok {
//					return nil, nil
//				}
//				parse, err := uuid.Parse(id)
//				if err != nil {
//					return nil, err
//				}
//				return customerRepo.GetCustomerByID(parse)
//			},
//		},
//		"allCustomers": &graphql.Field{
//			Type:        graphql.NewList(customerType),
//			Description: "Get all customers",
//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//				return customerRepo.GetAllCustomers()
//			},
//		},
//	},
//})
//CustomerInputType := graphql.NewInputObject(graphql.InputObjectConfig{
//	Name: "CustomerInput",
//	Fields: graphql.InputObjectConfigFieldMap{
//		"firstName":   &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
//		"lastName":    &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
//		"email":       &graphql.InputObjectFieldConfig{Type: graphql.NewNonNull(graphql.String)},
//		"phoneNumber": &graphql.InputObjectFieldConfig{Type: graphql.String},
//		"address":     &graphql.InputObjectFieldConfig{Type: graphql.String},
//	},
//})
//rootMutation := graphql.NewObject(graphql.ObjectConfig{
//	Name: "Mutation",
//	Fields: graphql.Fields{
//		"createCustomer": &graphql.Field{
//			Type:        customerType,
//			Description: "Create a new customer",
//			Args: graphql.FieldConfigArgument{
//				"input": &graphql.ArgumentConfig{Type: graphql.NewNonNull(CustomerInputType)},
//			},
//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//				input, ok := p.Args["input"].(map[string]interface{})
//				if !ok {
//					return nil, nil
//				}
//				firstName, _ := input["firstName"].(string)
//				lastName, _ := input["lastName"].(string)
//				phoneNumber, _ := input["phoneNumber"].(string)
//				email, _ := input["email"].(string)
//				//address, _ := input["address"].(string)
//
//				newCustomer := models.Customer{
//					Base:        models.Base{ID: uuid.New()},
//					FirstName:   firstName,
//					LastName:    lastName,
//					PhoneNumber: phoneNumber,
//					Email:       email,
//				}
//				if err := customerRepo.CreateCustomer(&newCustomer); err != nil {
//					return nil, err
//				}
//				return &newCustomer, nil
//			},
//		},
//	},
//})
//
//schema, err := graphql.NewSchema(graphql.SchemaConfig{
//	Query:    rootQuery,
//	Mutation: rootMutation,
//})
//if err != nil {
//	log.Fatal("Failed to create GraphQL schema:", err)
//}
//
//// Определение обработчика GraphQL для HTTP сервера
//h := handler.New(&handler.Config{
//	Schema:   &schema,
//	Pretty:   true,
//	GraphiQL: true,
//})
//http.Handle("/graphql", h)
//
//// Запуск HTTP сервера
//log.Println("Server running on port 8080")
//log.Fatal(http.ListenAndServe(":8080", nil))
