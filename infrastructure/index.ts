import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";
import * as awsx from "@pulumi/awsx";
import {deployLambdas} from "./modules/lambdas";

// Create an AWS resource (S3 Bucket)
const bucket = new aws.s3.Bucket("my-bucket-discord-test");

// Export the name of the bucket
export const bucketName = bucket.id;

deployLambdas()