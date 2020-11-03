package main

import (
	"flag"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

const Description = "Create a new security group"

func main() {

	createSecurityGroup()
}

func creds() *ec2.EC2 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	ec2sess := ec2.New(sess)

	return ec2sess
}

func createSecurityGroup() *ec2.CreateSecurityGroupOutput {
	groupPtr := flag.String("securitygroupname", "securitygroupname", "")
	vpcPtr := flag.String("vpcid", "securitygroupname", "")

	flag.Parse()

	secgroup, err := creds().CreateSecurityGroup(&ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(*groupPtr),
		VpcId:       aws.String(*vpcPtr),
		Description: aws.String(Description),
	})
	if err != nil {
		log.Println(err)
	}

	return secgroup
}
