package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lightsail"
	"github.com/aws/aws-sdk-go/aws/awserr"
)

func main() {
	config := &aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), ""),
	}
	sess := session.Must(session.NewSession(config))
	svc := lightsail.New(sess)

	b1, err := svc.GetBlueprints(&lightsail.GetBlueprintsInput{})
	if err != nil {
		log.Fatal(err)
	}
	for _, bp := range b1.Blueprints {
		if *bp.Platform == lightsail.InstancePlatformLinuxUnix && *bp.Type == lightsail.BlueprintTypeOs && *bp.Group == "ubuntu" && *bp.Version == "16.04 LTS" {
			fmt.Println(*bp.BlueprintId)
			break
		}
	}

	pubKey := `ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDQbExl6xdd9bK0N7C3uh48z6tpb7YbKue8KC6op0+b/r4QtPcahxcNdYc3ZLvZVVrDENpTO4HFUQsCkxe+zaBtThmr8YgPca7SkN0USxp6v+QeTL9UBiJeW1zgW4NnxOAxrQm1FAMim/pYeTHuzJkBfKHLakK76uydZHPr40wjkqzuNSvqS3jf5HLOJKAixEB3K4ZxxZiNK3hZhBzBFO/HbnEakG/cspuLxo6Go+c4FT+i7C4dmD3jvVJB/5tbbq+nG5qAc4D604095n4/3KpvCCqbQUaOmIGfnD/YRONS+5NlMeBEErwAV/zW/KreS85Vn7FvKaZiRtkxLlNoH/mr tamal@appscode.com`
	r1, err := svc.ImportKeyPair(&lightsail.ImportKeyPairInput{
		KeyPairName:     aws.String("b2"),
		PublicKeyBase64: aws.String(string(pubKey)),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r1)

	r2, err := svc.GetOperation(&lightsail.GetOperationInput{
		OperationId: r1.Operation.Id,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r2)

	r3, err := svc.GetKeyPair(&lightsail.GetKeyPairInput{
		KeyPairName: aws.String("b2"),
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r3)

	r4, err := svc.DeleteKeyPair(&lightsail.DeleteKeyPairInput{
		KeyPairName: aws.String("b22"),
	})
	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			fmt.Println("---------------------------------------")
			log.Println("Error:", awsErr.Code(), awsErr.Message())
			fmt.Println("---------------------------------------")
		}
		log.Fatal(err)
	}
	fmt.Println(r4)
}
