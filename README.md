# Season of Serverless - Challenge Week 7

This repository contains a solution for the Season of Serverless Challenge week 7 aka ["The Recipe"](https://github.com/microsoft/Seasons-of-Serverless/blob/main/Jan-4-2021.md) via Azure Functions in ... (drumm roll) ... _Go_ using Custom handler and deployed to SAP Cloud Platform Kyma environment (what ...?)

## Solution Components

The challenge of week 7 was not 100% clear to me and also changed a bit in between. As a consequence this solution focuses a bit more on what is possible with Azure Functions and not too much of the specifics. In addition I did not want to create a Twilio account.

The solution provides an http endpoint `api/recipe/`:

* specifying the country `api/recipe/nigeria` or `api/recipe/kenya` will give you a list of dishes for Kenya and Nigeria including links to their recipes
* adding a query parameter with the name of the dish will directly hand back the dish and link to recipe if one is found. An example call would be `api/recipe/kenya?recipename=Karanga`

## Internal Design

Internally the solution is based on Azure Functions custom handler making use of Go. I have build a simple server (see [`server.go`](./server.go)) that has a handler for two routes.
Depending on the provisioning of the query parameter `recipename` either a specific recipe link is returned if found in the hard coded values or the complete list of available recipe links is returned for the country specified in the path.

But let us take this one step further, why not deploy the solution to ... let us say SAP Cloud Platform [Kyma](https://kyma-project.io/) that represents a Kubernetes-based opinionated runtime. So we build the Docker container (and learned that the func-CLI has a bug for custo handlers when it comes to Docker files) for our Azure Function in Go and then deploy it to Kyma. Last but not least we deploy it via a corresponding API Rule. For details see directory `k8s`.

## How to execute

You can run the scenario locally via `func start` on a Linux or Mac machine. For windows you need to build the `server.go` file and adopt the host.json accordingly to point to the .exe file.  Several HTTP calls are available in the file `requests.http` (required: [REST CLient extension in VSCode](https://marketplace.visualstudio.com/items?itemName=humao.rest-client)).
The last four requests in the file target the deployed version on SAP Cloud Platform.
