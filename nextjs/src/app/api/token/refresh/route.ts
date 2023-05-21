import { RefreshTokenResponseData, RefreshTokenResponseDataExternalApi } from '@/models/refreshToken';
import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {

  const headers = request.headers
  let axiosReqHeaders = {} 

  if(headers.has('Cookie')) {
    const refreshToken = getRefreshTokenFromCookie(headers.get('Cookie')!)
    axiosReqHeaders = {...axiosReqHeaders, Authorization: `Bearer ${refreshToken}`}
  }

  try {
    const res = await externalApi.post<RefreshTokenResponseDataExternalApi>('/token/refresh',{} , {
      headers: axiosReqHeaders
    })

    const resData:RefreshTokenResponseData = {
      accessToken: res.data.accessToken
    }

    const response = NextResponse.json(resData,{
      status: res.status
    })

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

const getRefreshTokenFromCookie = (cookie: string): string => {

  // reg ex
  const regex = /refresh-token\s*=\s*(\S+)/gm;

  let m;
  let token: string = ''

  while ((m = regex.exec(cookie)) !== null) {
      // This is necessary to avoid infinite loops with zero-width matches
      if (m.index === regex.lastIndex) {
          regex.lastIndex++;
      }
      
      // The result can be accessed through the `m`-variable.
      m.forEach((match:string, groupIndex) => {
          // console.log(`Found match, group ${groupIndex}: ${match}`);
          token = match
      });
  }

  return token

}