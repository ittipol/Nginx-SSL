import api from "@/util/axios"
import { createAsyncThunk, createSlice } from "@reduxjs/toolkit"
import { AxiosError } from "axios"
import { RegisterBody } from "@/models/register";

export interface UserState {
    isAuth: boolean,
    accessToken: string,
    name: string
}

const initialState: UserState = {
    isAuth: false,
    accessToken: '',
    name: ''
}

export const userProfile = createAsyncThunk(
    'user/me',
    async (_, thunkAPI) => {
        
        try {
            const res = await api.get<{name: string}>('profile')
            return thunkAPI.fulfillWithValue<typeof res.data>(res.data)

        }
        catch (ex) {
            const error = ex as AxiosError
            return thunkAPI.rejectWithValue(error.response?.status)
        }

    }
)

export const userLogin = createAsyncThunk(
    'user/login',
    async (body :{email:string, password:string}, thunkAPI) => {
        
        try {
            const res = await api.post<{accessToken: string}>('login', body)
            return res.data
        }
        catch (ex) {
            const error = ex as AxiosError
            return thunkAPI.rejectWithValue(error.response?.status)
        }
        
    }
)

export const userRegister = createAsyncThunk(
    'user/register',
    async (body: RegisterBody, thunkAPI) => {

        try {
            const res = await api.post<{accessToken: string}>('register', body)
            return {
                data: res.data,
                status: res.status
            }
        }
        catch (ex) {
            const error = ex as AxiosError
            return thunkAPI.rejectWithValue({
                data: error.response?.statusText,
                status: error.response?.status
            })
        }

    }
)

export const userSlice = createSlice({
    name: 'user',
    initialState,
    reducers: { },
    extraReducers(builder) {
        builder
        .addCase(userLogin.fulfilled, (state, action) => {
            state.isAuth = true
            state.accessToken = action.payload.accessToken
        })
        .addCase(userLogin.rejected, (state, _) => {
            state.isAuth = false
            state.accessToken = ''
        })

        .addCase(userRegister.fulfilled, (state, action) => {
            
        })
        .addCase(userRegister.rejected, (state, action) => {
            
        })

        .addCase(userProfile.fulfilled, (state, action) => {
            state.name = action.payload.name
        })
        .addCase(userProfile.rejected, (state, action) => {

        })
    },
})

export default userSlice.reducer