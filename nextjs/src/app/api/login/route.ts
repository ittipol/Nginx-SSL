import { LoginResponseData } from '@/models/login';
import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {

  const body = await request.json()

  console.log('========================================')
  console.log(body)
  // console.log(request.headers)
  console.log('========================================')

  try {
    const res = await externalApi.post<LoginResponseData>('/login', body)

    const response = NextResponse.json(res.data,{
      status: res.status,
      // headers: {
      //   'Access-Control-Allow-Headers': 'Content-Type, Authorization, Set-Cookie'
      // }
    })

    // Set header from response
    // Object.keys(res.headers).forEach(
    //   (key) => {
    //       response.headers.set(key, res.headers[key])
    //   }
    // )

    response.cookies.set({
      name: 'refresh-token',
      value: res.data.accessToken,
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