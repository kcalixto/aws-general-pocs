import AWS from "aws-sdk"

export const handler = async (event, ctx, callback) => {
    AWS.config.update({ region: "sa-east-1" });

    try {
        const s3 = new AWS.S3();

        // const data = s3.listObjects({
        //     Bucket: "kauacalixtolab.xyz"
        // }).promise()
        // console.log("data: ", data);

        const obj = await s3.getObject({
            Bucket: "kcalixto-private-bucket-test",
            Key: "print_01"
        }).promise()

        console.log("your obj: ", obj);
    } catch (error) {
        console.log(`request failed with: status ${error.statusCode} message: ${error.message}`);
    }
}
handler()