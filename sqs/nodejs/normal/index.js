import AWS from "aws-sdk"

export const handler = (event, ctx, callback) => {
    try {
        console.log(event.Records[0].body);

        AWS.config.update({
            "region": "sa-east-1"
        })

        const db = new AWS.DynamoDB()

        const body = JSON.parse(event.Records[0].body)
        const pk = body?.firebolt_id + "_" + body?.event_name
        const params = {
            "TableName": "teste-sqs",
            "Item": {
                "formID_eventName": {
                    S: pk,
                },
                "action": {
                    S: body?.action ?? "",
                },
                "appsflyer_id": {
                    S: body?.appsflyer_id ?? "",
                },
                "campaign": {
                    S: body?.campaign ?? "",
                },
                "event_name": {
                    S: body?.event_name ?? "",
                },
                "event_value": {
                    S: body?.event_value ?? "",
                },
                "firebolt_id": {
                    S: body?.firebolt_id ?? "",
                },
                "oc_start_source": {
                    S: body?.oc_start_source ?? "",
                }
            }
        }

        db.putItem(params)

        db.getItem({
            "TableName": "teste-sqs",
            "Key": {
                S: pk,
            }
        })

    } catch (error) {
        console.log(error);
    }
}

