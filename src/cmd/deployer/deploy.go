package main

import (
	"github.com/danielstutzman/sync-cloudfront-logs-to-bigquery/src/lambda_deployer"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 1+1 {
		log.Fatalf("Usage: " + os.Args[0] + " up|down")
	}
	verb := os.Args[1]

	deployer := lambda_deployer.NewLambdaDeployer(lambda_deployer.Config{
		BucketName:   "cloudfront-logs-danstutzman",
		FunctionName: "SyncCloudfrontLogsToBigquery",
		RoleName:     "lambda-SyncCloudfrontLogsToBigquery-execution",
		PolicyName:   "lambda-SyncCloudfrontLogsToBigquery-execution-access",
	})

	switch verb {
	case "down":
		deployer.DeleteEverything()
	case "up":
		deployer.SetupBucket()
		deployer.DeployFunction()
	default:
		log.Fatalf("Unknown verb '%s'; expected up or down", verb)
	}

	log.Printf("%s completed successfully", os.Args[0])
}