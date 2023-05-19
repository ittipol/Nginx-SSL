import { ResponseResult } from "./response"

export interface RegisterBody {
    email: string | undefined,
    password: string | undefined,
    name: string | undefined
}

export interface RegisterResponseResult extends ResponseResult {
    data: RegisterResponseData
}

export interface RegisterResponseData {
    message: string
}