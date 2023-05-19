import { ResponseResult } from "./response"

export interface LoginBody {
    email: string,
    password: string
}

export interface LoginResponseResult extends ResponseResult {
    data: LoginResponseData
}

export interface LoginResponseData {
    accessToken: string
}