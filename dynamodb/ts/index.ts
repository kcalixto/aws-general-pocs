import * as AWS from "aws-sdk"

async function load() {
    try {
        const params: AWS.DynamoDB.DocumentClient.PutItemInput = {
            TableName: 'geminio-test-table',
            Item: {
                "id": { S: "123" },
                "anonymous_id": { S: "123" },
                "event_name": { S: "123" },
                "product_name": { S: "123" },
                "proposal_id": { S: "123" },
                "oc_match": { S: "123" },
                "flag_decil": { BOOL: false },
                "created_at": { S: Date.now() + "" },
            },
        };

        const dynamoDB = new AWS.DynamoDB({ region: "sa-east-1" });
        await dynamoDB.putItem(params).promise()

    } catch (e: any) {
        console.log("error: ", e.message);
    }
}
load()