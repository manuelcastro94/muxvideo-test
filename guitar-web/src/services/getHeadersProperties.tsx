

export enum METHODS {
    GET = "get",
    POST = "post",
    PUT = "put",
    DELETE= "delete",
    PATCH= "patch",
}

export function GetHeadersProperties(method:METHODS, body?: any) {

    let headerSettings = {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Methods': 'POST,PUT,GET,DELETE,PATCH,OPTIONS',
        'Access-Control-Allow-Headers': 'Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With, Access-Control-Allow-Origin, Access-control-allow-methods, Access-Control-Request-Method',
        'Access-Control-Request-Method': 'POST,PUT,GET,DELETE,PATCH,OPTIONS',
    };

    const props: { method: METHODS, credentials?: "include", headers: any, body?: any } = {method: method, headers: headerSettings};

    props.credentials = "include";

    if (body)
        props.body = JSON.stringify(body);
    return props
}