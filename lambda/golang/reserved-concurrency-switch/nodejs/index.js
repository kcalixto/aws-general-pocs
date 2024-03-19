export const handler = async (event, context) => {
    return {
        statusCode: 200,
        body: JSON.stringify({
            message: `Reserved concurrency is ${context?.functionReservedConcurrency || 'not set'}`
        })
    }
}