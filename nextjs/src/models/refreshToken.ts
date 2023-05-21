import { ResponseResult } from "./response"

export interface RefreshTokenResponseResult extends ResponseResult {
    data: RefreshTokenResponseData
}

export interface RefreshTokenResponseData {
    accessToken: string
}

export interface RefreshTokenResponseDataExternalApi {
    accessToken: string
    refreshToken: string
}