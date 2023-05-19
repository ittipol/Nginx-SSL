import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {

  const body = await request.json()

  console.log('========================================')
  console.log(body)
  console.log('========================================')

  try {
    const res = await externalApi.post('/login', body)

    const response = NextResponse.json(res.data,{
      status: res.status
    })

    Object.keys(res.headers).forEach(
      (key) => {
          response.headers.set(key, res.headers[key])
      }
    )

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