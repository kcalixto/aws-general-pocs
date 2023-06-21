function load(event) {
    if (event.Headers["content-type"] == "application/x-www-form-urlencoded") {
        let body = atob(event.Body).split("&").map(function (x) {
            return x.split("=")
        }).reduce(function (previousItem, splitedNextItem) {
            const field = splitedNextItem[0];
            const value = splitedNextItem[1];

            previousItem[field] = value;

            return previousItem
        }, {});

        console.log(body);
    }

}

load({
    Headers: {
        "content-type": "application/x-www-form-urlencoded",
    },
    Body: "Z3JhbnRfdHlwZT11cm4lM0FpZXRmJTNBcGFyYW1zJTNBb2F1dGglM0FncmFudC10eXBlJTNBand0LWJlYXJlciZhc3NlcnRpb249ZXlKaGJHY2lPaUpTVXpJMU5pSXNJblI1Y0NJNklrcFhWQ0o5LmV5SmhkV1FpT2lKb2RIUndjem92TDNCeWIzaDVMbUZ3YVM1d2NtVmlZVzVqYnk1amIyMHVZbkl2WVhWMGFDOXpaWEoyWlhJdmRqRXVNUzkwYjJ0bGJpSXNJbk4xWWlJNklqRXpNakkyTnpOa0xUVTVaREl0TkdWbE55MDROMk0zTFdRNE9UaGpOVFJqTmpVd1lpSXNJbWxoZENJNklqRTJOakEyTURFeE56TWlMQ0psZUhBaU9pSXhOall3TmpBME9EYzVJaXdpYW5ScElqb2lNVFkyTURZd01UTTJORGsyTVNJc0luWmxjaUk2SWpFdU1TSjkuVmVFLWpja24yOVpfbGtMU3E5Z1FfVDYxME9zOUN0dTlsb3d2el9nUDVjWnYwYW1ZSVBQSmNsS1BZYktHS2VOYXJPbDd0Nk83ZUFYT2dmRjJfbWFvTVRvNnp4X3FfaWxhQnpDQU94ZTQ4Z2VzWjN3OUtObWZsalBPRVIxYXo5QlIxWmdKVlRpRllUVS1aTmRLQVV5dmV4b3hMaEpLMENKcWx1UjU0c1Vvb2lQYlBDTlljU205M01ORFh3ZjFoQk1KWUczY3p1c3lOYlBrYjVpNm4yT1pidWZObmpUc1M0UHI2RXBUZ3ZWOEFMa3dMbGNtZ2drSm5MS21IZzJEYXZPd0tvcWJTdzFPQ0pmZDZvZVY2OEppc2VwVUNaZ05HQzBmUDlTdGVoX3ZkRFN0OEtoVFVqRTBjZEhMSVltcDBWZ1hyMDJSWmw3VFBIeHNaZXFlS2NjLVh"
})