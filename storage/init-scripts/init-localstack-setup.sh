#!/bin/sh
echo "Begin: DynamoDB -> Create Table"
#
awslocal dynamodb create-table \
    --region us-east-1 \
    --cli-input-json file://dynamodb-table-definition.json
#
echo "End: DynamoDB -> Create Table"
#
echo "Begin: DynamoDB -> List Tables"
#
awslocal dynamodb list-tables
#
echo "End: DynamoDB -> List Tables"
#
echo "Begin: DynamoDB -> Load Base Establishment"
#
awslocal dynamodb batch-write-item \
  --region us-east-1 \
  --request-items file://data/base-establishment.json
echo "End: DynamoDB -> Load Base Establishment"
#
echo "Begin: DynamoDB -> Load Base User"
#
awslocal dynamodb batch-write-item \
  --region us-east-1 \
  --request-items file://data/base-user.json
echo "End: DynamoDB -> Load Base User"
#
echo "Begin: DynamoDB -> Load Base Events"
#
awslocal dynamodb batch-write-item \
  --region us-east-1 \
  --request-items file://data/base-events.json
echo "End: DynamoDB -> Load Events"
#
echo "Begin: DynamoDB -> Load Base Workshops"
#
awslocal dynamodb batch-write-item \
  --region us-east-1 \
  --request-items file://data/base-workshops.json
echo "End: DynamoDB -> Load Workshops"
#
echo "Start: DynamoDB -> Scan table"
awslocal dynamodb scan \
    --region us-east-1 \
    --table-name SERVICES_EVENTS_ADMIN_SERVICE
echo "End: DynamoDB -> Scan table"
#
echo "Begin: S3 -> Make Bucket"
#
awslocal s3 mb \
    --region us-east-1 \
    s3://event-admin-service-file-storage
#
echo "End: S3 -> Make Bucket"
#
echo "Begin: S3 -> List Buckets"
#
awslocal s3 ls
#
echo "End: S3 -> List Buckets"
