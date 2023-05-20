import externalApi from '@/util/externalApi';
import { AxiosError } from 'axios';
import { NextResponse } from 'next/server';

export async function POST(request: Request) {

  const headers = request.headers
  const cookies = request.headers.get('cookie')
  let axiosReqHeaders = {} 

  console.log('>>> API Refresh Token ========================================')
//   console.log(headers.has('Cookie'))
//   console.log(headers.get('Cookie'))
// console.log(headers.has('Authorization'))
// console.log(headers.get('Authorization')) 
// console.log(headers)
// console.log(cookies)
console.log('========================================')

  // axiosReqHeaders = {...axiosReqHeaders, Authorization: headers.get('Authorization')}
//   axiosReqHeaders = {
//     Authorization: 'xxx'
//   }
//   console.log(axiosReqHeaders)

  try {
    const res = await externalApi.get('/token/refresh', {
      // headers: axiosReqHeaders
    })

    // return NextResponse.json(res.data,{
    //   status: res.status
    // })

    return NextResponse.json("Failed", {
        status: 200
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