import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function GET(request: Request) {

  const headers = request.headers
  let reqHeaders = {}

  if(headers.has('Authorization')) {
    reqHeaders = {...reqHeaders, Authorization: `${headers.get('Authorization')}`}
  }

  try {
    const res = await externalApi.get('/user/profile', {
      headers: reqHeaders
    })

    return NextResponse.json(res.data,{
      status: res.status
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