import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function GET(request: Request) {

  const headers = request.headers
  let axiosReqHeaders = {} 

  console.log('========================================')
  console.log(headers.has('Authorization'))
  console.log(headers.get('Authorization'))
  console.log('========================================')

  axiosReqHeaders = {...axiosReqHeaders, Authorization: headers.get('Authorization')}
  // console.log(axiosReqHeaders)

  try {
    const res = await externalApi.get('/user/profile', {
      headers: axiosReqHeaders
    })

    return NextResponse.json(res.data,{
      // status: res.status
      status: 401
    })
  } 
  catch(ex) {
    const error = ex as AxiosError
    
    return NextResponse.json(error.response?.statusText,{
      status: error.response?.status
    })
  }
  
}

export async function OPTIONS(request: Request) {

  return NextResponse.json("",{
    status: 200
  })

}