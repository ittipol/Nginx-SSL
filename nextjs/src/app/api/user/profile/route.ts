import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function GET(request: Request) {

  const headers = request.headers

  console.log('========================================')
  console.log(headers)
  console.log('========================================')

  try {
    const res = await externalApi.get('/user/profile',{
      headers: {
        // name:"data"
      }
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