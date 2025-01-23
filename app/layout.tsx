import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./globals.css"
import { ThemeProvider } from "@/components/theme-provider"
import { NextSSRPlugin } from "@uploadthing/react/next-ssr-plugin"
import { extractRouterConfig } from "uploadthing/server"
import { ourFileRouter } from "./api/uploadthing/core"
import { BooksProvider } from "@/contexts/BooksContext"
import { TTSProvider } from "@/contexts/TTSContext"

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "PDF Player",
  description: "A modern PDF reader with text-to-speech capabilities",
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en" suppressHydrationWarning>
      <head />
      <body className={inter.className}>
        <ThemeProvider
          attribute="class"
          defaultTheme="system"
          enableSystem
          disableTransitionOnChange
        >
          <BooksProvider>
            <TTSProvider>
              <NextSSRPlugin routerConfig={extractRouterConfig(ourFileRouter)} />
              {children}
            </TTSProvider>
          </BooksProvider>
        </ThemeProvider>
      </body>
    </html>
  )
}

