import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {

  const body = await request.json()

  console.log('========================================')
  console.log(body)
  console.log('========================================')

  try {
    const res = await externalApi.post('/register', body)

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