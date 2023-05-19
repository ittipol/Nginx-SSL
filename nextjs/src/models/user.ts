import { ResponseResult } from "./response";

export interface UserRequest {}

export interface UserResponseResult extends ResponseResult {
    data: UserResponseData
}

export interface UserResponseData {
    name: string
}