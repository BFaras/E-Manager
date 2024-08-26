import {
    clerkMiddleware,
    createRouteMatcher
  } from '@clerk/nextjs/server';
import { NextResponse } from 'next/server';
  
  const isProtectedRoute = createRouteMatcher([
    '/api/:path'
  ]);

  const isIgnoredRoute = createRouteMatcher([
    '/api/webhook', 
    '/api/stores',
  ]);
  
  export default clerkMiddleware((auth, req) => {
    if (isIgnoredRoute(req)) {
      return NextResponse.next();
    }
    if (isProtectedRoute(req)) auth().protect();
  });
  
  export const config = {
    matcher: ['/((?!.*\\..*|_next).*)', '/', '/(api|trpc)(.*)'],
  };