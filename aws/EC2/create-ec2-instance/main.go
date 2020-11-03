package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	createEC2Instance()
}

func creds() *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2sess := ec2.New(sess)

	return ec2sess
}

func createEC2Instance() *ec2.Reservation {
	imageIDPtr := flag.String("imageid", "", "")
	instanceTypePtr := flag.String("instancetype", "", "")

	flag.Parse()

	ec2, err := creds().RunInstances(&ec2.RunInstancesInput{
		ImageId:      aws.String(*imageIDPtr),
		InstanceType: aws.String(*instanceTypePtr),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		log.Println(err)
	}

	return ec2
}
