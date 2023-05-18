import { ResponseData } from '@/models/response';
import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

// Method Get
// export async function GET(request: Request) {}

export async function POST(request: Request) {

  const body = await request.json()

  console.log('========================================')
  console.log(body)
  console.log('========================================')

  try {
    const res = await externalApi.post('register', body)

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