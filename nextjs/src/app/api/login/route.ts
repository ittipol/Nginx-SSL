import { LoginResponseData, LoginResponseDataExternalApi } from '@/models/login';
import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {

  const body = await request.json()

  try {

    const res = await externalApi.post<LoginResponseDataExternalApi>('/login', body)

    const resData:LoginResponseData = {
      accessToken: res.data.accessToken
    }

    const response = NextResponse.json(resData, {
      status: res.status
    })

    // Set header from response
    // Object.keys(res.headers).forEach(
    //   (key) => {
    //       response.headers.set(key, res.headers[key])
    //   }
    // )

    response.cookies.set({
      name: 'refresh-token',
      value: res.data.refreshToken,
      httpOnly: true,
      // secure: true,
      sameSite: 'strict',
      path: '/',
      // domain: ''
    })

    return response
  } 
  catch(ex) {
    const error = ex as AxiosError

    const response = NextResponse.json(error.response?.statusText,{
      status: error.response?.status
    })

    response.cookies.delete("refresh-token")

    return response
  }
  
}

export async function OPTIONS(request: Request) {

  return NextResponse.json("",{
    status: 200
  })

}