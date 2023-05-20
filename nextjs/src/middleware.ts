import { NextResponse } from 'next/server';
import type { NextRequest } from 'next/server';
 
const allowedOrigins = [
  'http://testhost:3000',
  "http://abc:3000"
];

// This function can be marked `async` if using `await` inside
export function middleware(request: NextRequest) {

    console.log('middleware =======')
    console.log(`URL: ${request.nextUrl.pathname}`)
    console.log(request.headers)

    if (request.nextUrl.pathname.startsWith('/user')) {
        // return NextResponse.rewrite(new URL('/register', request.url));
    } else if (request.nextUrl.pathname.startsWith('/api')) {

      const response = NextResponse.next()
      
      // CORS
      response.headers.set('Access-Control-Allow-Credentials', "true")

      if(request.headers.has('origin') && allowedOrigins.includes(request.headers.get('origin')!)) {
        console.log(`Access-Control-Allow-Origin, ${request.headers.get('origin')}`)
        response.headers.set('Access-Control-Allow-Origin', request.headers.get('origin')!)
      }else {
        console.log('Origin not found')
        console.log(request.nextUrl)
        return NextResponse.rewrite(new URL('/404', request.url))
      }
      
      response.headers.set('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
      response.headers.set('Access-Control-Allow-Headers', 'Content-Type, Authorization')

      return response
    }

    return NextResponse.next()
}
 
// See "Matching Paths" below to learn more
export const config = {
  matcher: ['/user/:path*','/api/:path*'],
};