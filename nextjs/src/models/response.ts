// type responseData = {
//     value: any,
//     status: number|undefined
//   }

export interface ResponseResult {
    status: number
}

// export interface ExternalApiResponseData extends ResponseResult {
//     data: any
// }

export interface ErrorResult extends ResponseResult {
    error: string|number|undefined,
}