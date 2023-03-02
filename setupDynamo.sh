BASE_URL=http://localhost:4566
PROFILE_TABLENAME=Profile

aws dynamodb --endpoint-url=$BASE_URL create-table \
    --table-name $PROFILE_TABLENAME \
    --attribute-definitions \
        AttributeName=UserId,AttributeType=S \
    --key-schema \
        AttributeName=UserId,KeyType=HASH \
--provisioned-throughput \
        ReadCapacityUnits=10,WriteCapacityUnits=5


aws --endpoint-url=$BASE_URL dynamodb put-item \
    --table-name $PROFILE_TABLENAME \
    --item \
        '{"UserId": {"S": "37d10e18-34a2-4bd2-b7bc-b8e6dd6358f1"}, "Email": {"S": "demo0@coinbase.com"}, "Name": {"S": "Ted Robinson"}, "LegalName": {"S": "Ted Robinson"}, "UserName": {"S": "d0"}, "Address": {"S": "Some Mountain, Canada"}, "DateOfBirth": {"S": "10/22/2003"}}'

aws --endpoint-url=$BASE_URL dynamodb put-item \
    --table-name $PROFILE_TABLENAME \
    --item \
        '{"UserId": {"S": "4f5a6336-8101-4634-a458-73b7f6fcf49f"}, "Email": {"S": "demo1@coinbase.com"}, "Name": {"S": "Henry Thomas"}, "LegalName": {"S": "Henry Thomas"}, "UserName": {"S": "d1"}, "Address": {"S": "Some Mountain, Canada"}, "DateOfBirth": {"S": "10/22/2003"}}'

