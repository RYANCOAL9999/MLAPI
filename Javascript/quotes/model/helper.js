export const handleResponse = (response, res) => {
    return res.status(200).json(response);
};

export const handleGRPCError = (err, res, errMsg) => {
    console.error(`could not ${errMsg}: `, err);
    return res.status(500).json({
        error: `could not ${errMsg}`,
    });
};

export const handleBadRequestError = (errMsg) => {
    console.log(`could not ${errMsg}: Bad Request`);
    return res.status(400).json({
        error: `Bad Request`,
    });
}
export const sendStubRequest = async (stub, event, request_message) => {
    try {
        const func = stub[event];
        const response = await func(request_message);
        return response;
    } catch (err) {
        console.error(`Could not ${event}: `, err);
        throw new Error(`Could not ${event}`);
    }
};
