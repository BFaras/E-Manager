import { ClerkProvider, RedirectToSignIn, SignInButton, SignedIn, SignedOut, UserButton } from '@clerk/nextjs'
import './globals.css';
import { ModalProvider } from '@/providers/modal-provider';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <ClerkProvider>
          <html>
      <body>
      <ModalProvider/>
      {children}
      </body>
    </html>
  </ClerkProvider>
  )
}