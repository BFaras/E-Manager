import { ClerkProvider, RedirectToSignIn, SignInButton, SignedIn, SignedOut, UserButton } from '@clerk/nextjs'
import './globals.css';
import SignInPage from './(auth)/(routes)/sign-in/[[...sign-in]]/page';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <ClerkProvider>
          <html>
      <body>
      {children}
</body>
    </html>
  </ClerkProvider>
  )
}