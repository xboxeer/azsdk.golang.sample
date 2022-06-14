package main

import (
	"log"
	"os"
)

var (
	subscriptionID    string
	location          = "westus"
	resourceGroupName = "sample-resource-group"
	namespaceName     = "sample-sb-namespace" //namespace name need to be unique in azure, change to another name if it already exist, GUID is not allowed.
	queueName         = "sample-sb-queue"
)

func main() {
	subscriptionID = os.Getenv("AZURE_SUBSCRIPTION_ID")
	if len(subscriptionID) == 0 {
		log.Fatal("AZURE_SUBSCRIPTION_ID is not set.")
	}
	connectToAzure()
	createResourceGroup(resourceGroupName)
	createNamespace(namespaceName)
	createQueue(queueName)
	listAllQueue(namespaceName)
	deleteQueue(queueName)
	getAllResourceGroups()
}

//All the function defined below is just for demo purpose
//you can change the function signature such as add/remove parameters, return values, change function names
func connectToAzure() error {
	//TODO
}

func createResourceGroup(resrouceGroupName string) error {
	//TODO
}

func createNamespace(namespaceName string) error {
	//TODO
}

func createQueue(queueName string) error {
	//TODO
}

func listAllQueue(namespaceName string) error {
	//TODO
}

func deleteQueue(queueName string) error {
	//TODO
}

func getAllResourceGroups() error {
	//TODO
}
