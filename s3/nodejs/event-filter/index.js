export const handler = (event, ctx, callback) => {
    try {
        console.log("Hi, i was called!");
        console.log(JSON.stringify(event));
    } catch (error) {
        console.log(error);
    }
}

