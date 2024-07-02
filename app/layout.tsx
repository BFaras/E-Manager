import { ClerkProvider, RedirectToSignIn, SignInButton, SignedIn, SignedOut, UserButton } from '@clerk/nextjs'
import './globals.css';
import { ModalProvider } from '@/providers/modal-provider';
import { ToasterProvider } from '@/providers/toast-provider';

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <ClerkProvider>
          <html>
      <body>
      <ToasterProvider></ToasterProvider>
      <ModalProvider/>
      {children}
      </body>
    </html>
  </ClerkProvider>
  )
}