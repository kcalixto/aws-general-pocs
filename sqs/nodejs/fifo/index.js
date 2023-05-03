import AWS from "aws-sdk"
import { v4 as uuidv4 } from 'uuid';

export const handler = async (event, ctx, callback) => {
    return new Promise(resolve => {
        try {
            console.log(event.Records[0].body);

            AWS.config.update({
                "region": "sa-east-1"
            })

            const db = new AWS.DynamoDB()

            const body = JSON.parse(event.Records[0].body)
            const pk = uuidv4();
            // console.log("pk: ", pk);

            const params = {
                TableName: "teste-sqs",
                Item: {
                    "formID_eventName": {
                        S: pk,
                    },
                    "firebolt_id": {
                        S: body?.firebolt_id ?? "",
                    },
                    "event_name": {
                        S: body?.event_name ?? "",
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
                    "event_value": {
                        S: body?.event_value ?? "",
                    },

                    "oc_start_source": {
                        S: body?.oc_start_source ?? "",
                    }
                }
            }

            db.putItem(params, (err, data) => {
                if (err) throw new Error(err)
                console.log(data);
            })

            db.getItem({
                "TableName": "teste-sqs",
                "Key": {
                    "formID_eventName": {
                        S: pk,
                    }
                },
            }, (err, data) => {
                if (err) throw new Error(err)
                console.log(data.Item);
            })

        } catch (error) {
            console.log(error);
        } finally {
            setTimeout(resolve, 500);
        }
    })
}