# Jockey Image

Generated: 02-07-2025 at 21:27:20

## Repository Structure

```
pdf-player
│   ├── contexts
│   ├── styles
│   ├── tailwind.config.ts
│   ├── types
│       └── index.ts
│   ├── package-lock.json
│   ├── app
│   ├── tailwind.config.js
│   ├── backend.txt
│   ├── lib
│       └── uploadthing.ts
│   ├── utils
│       └── audioCache.ts
│   ├── ereader.db
│   ├── public
│   │   ├── placeholder-logo.png
│       └── placeholder-user.jpg
│   ├── There Is No Antimemetics Division.pdf
│   ├── components.json
│   ├── postcss.config.mjs
│   ├── backend
│   │   ├── models.go
│   │   ├── go.sum
│   │   ├── db.go
│   │   ├── README.md
│   │   ├── config.go
│   │   ├── go.mod
│   │   ├── storage.go
│   │   ├── schema.sql
│   │   ├── tts.go
│       └── jockey-img
│           └── backend_01-22-25.md
│   ├── components
│   │   ├── theme-provider.tsx
│   │   ├── PDFReader.tsx
│   │   ├── PDFUploader.tsx
│   │   ├── pdf-reader.tsx
│   │   ├── HomePage.tsx
│   │   ├── book-drawer.tsx
│       └── ui
│       │   ├── context-menu.tsx
│       │   ├── resizable.tsx
│       │   ├── alert-dialog.tsx
│       │   ├── carousel.tsx
│       │   ├── sheet.tsx
│       │   ├── toggle-group.tsx
│       │   ├── checkbox.tsx
│       │   ├── toast.tsx
│       │   ├── radio-group.tsx
│       │   ├── navigation-menu.tsx
│       │   ├── accordion.tsx
│       │   ├── dialog.tsx
│       │   ├── dropdown-menu.tsx
│       │   ├── aspect-ratio.tsx
│       │   ├── table.tsx
│       │   ├── command.tsx
│       │   ├── slider.tsx
│       │   ├── switch.tsx
│       │   ├── sonner.tsx
│       │   ├── card.tsx
│       │   ├── tooltip.tsx
│       │   ├── breadcrumb.tsx
│       │   ├── select.tsx
│       │   ├── form.tsx
│       │   ├── scroll-area.tsx
│       │   ├── tabs.tsx
│       │   ├── use-mobile.tsx
│       │   ├── alert.tsx
│       │   ├── use-toast.ts
│       │   ├── progress.tsx
│       │   ├── chart.tsx
│           └── toaster.tsx
│   ├── hooks
│   │   ├── useBooks.ts
│       └── use-toast.ts
│   ├── next.config.mjs
│   ├── tsconfig.json
│   ├── package.json
    └── test.txt
```

## File: /Users/saint/Desktop/pdf-player/types/index.ts

```ts
export interface Book {
  id: string
  title: string
  author: string
  coverUrl: string
  fileUrl: string
  status: 'pending' | 'processing' | 'ready'
  currentPage?: number
  pageCount?: number
  language?: string
  createdAt?: string
  updatedAt?: string
}

export interface AudioSegment {
  id: string
  bookId: string
  content: string
  audioUrl: string
  status: 'pending' | 'processing' | 'completed' | 'error'
  createdAt?: string
  updatedAt?: string
}

export interface UploadResponse {
  id: string
  status: 'pending' | 'processing' | 'ready'
} 
```

## File: /Users/saint/Desktop/pdf-player/tailwind.config.js

```js
/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ["class"],
  content: ["./pages/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}", "./app/**/*.{ts,tsx}", "./src/**/*.{ts,tsx}"],
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      colors: {
        border: "hsl(var(--border))",
        input: "hsl(var(--input))",
        ring: "hsl(var(--ring))",
        background: "hsl(var(--background))",
        foreground: "hsl(var(--foreground))",
        primary: {
          DEFAULT: "hsl(var(--primary))",
          foreground: "hsl(var(--primary-foreground))",
        },
        secondary: {
          DEFAULT: "hsl(var(--secondary))",
          foreground: "hsl(var(--secondary-foreground))",
        },
        destructive: {
          DEFAULT: "hsl(var(--destructive))",
          foreground: "hsl(var(--destructive-foreground))",
        },
        muted: {
          DEFAULT: "hsl(var(--muted))",
          foreground: "hsl(var(--muted-foreground))",
        },
        accent: {
          DEFAULT: "hsl(var(--accent))",
          foreground: "hsl(var(--accent-foreground))",
        },
        popover: {
          DEFAULT: "hsl(var(--popover))",
          foreground: "hsl(var(--popover-foreground))",
        },
        card: {
          DEFAULT: "hsl(var(--card))",
          foreground: "hsl(var(--card-foreground))",
        },
      },
      borderRadius: {
        lg: "var(--radius)",
        md: "calc(var(--radius) - 2px)",
        sm: "calc(var(--radius) - 4px)",
      },
      keyframes: {
        "accordion-down": {
          from: { height: 0 },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: 0 },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
}

```

## File: /Users/saint/Desktop/pdf-player/contexts/BooksContext.tsx

```tsx
"use client"

import { createContext, useContext, ReactNode } from 'react'
import { useBooks as useBooksHook } from '@/hooks/useBooks'
import { Book } from '@/types'

interface BooksContextType {
  books: Book[]
  loading: boolean
  error: string | null
  addBook: (book: Book) => Promise<void>
  updateBook: (id: string, updates: Partial<Book>) => Promise<void>
  getBookById: (id: string) => Book | undefined
  refreshBooks: () => Promise<void>
}

const BooksContext = createContext<BooksContextType | undefined>(undefined)

export function BooksProvider({ children }: { children: ReactNode }) {
  const booksData = useBooksHook()

  return (
    <BooksContext.Provider value={booksData}>
      {children}
    </BooksContext.Provider>
  )
}

export function useBooks() {
  const context = useContext(BooksContext)
  if (context === undefined) {
    throw new Error('useBooks must be used within a BooksProvider')
  }
  return context
} 
```

## File: /Users/saint/Desktop/pdf-player/contexts/TTSContext.tsx

```tsx
"use client"

import { createContext, useContext, ReactNode, useState, useCallback, useRef, useEffect } from 'react'
import { Book, AudioSegment } from '@/types'
import { audioCache } from '@/utils/audioCache'

interface TTSContextType {
  isPlaying: boolean
  currentTime: number
  duration: number
  currentSegment: number
  segments: AudioSegment[]
  currentBook: Book | null
  play: () => void
  pause: () => void
  stop: () => void
  skipForward: () => void
  skipBackward: () => void
  setAudioSource: (url: string) => void
  seekTo: (time: number) => void
  setSpeed: (speed: number) => void
  loadBookAudio: (book: Book) => Promise<void>
  playSegment: (segmentIndex: number) => void
}

const TTSContext = createContext<TTSContextType | undefined>(undefined)

export function TTSProvider({ children }: { children: ReactNode }) {
  const [isPlaying, setIsPlaying] = useState(false)
  const [currentTime, setCurrentTime] = useState(0)
  const [duration, setDuration] = useState(0)
  const [currentSegment, setCurrentSegment] = useState(0)
  const [playbackSpeed, setPlaybackSpeed] = useState(1)
  const [segments, setSegments] = useState<AudioSegment[]>([])
  const [currentBook, setCurrentBook] = useState<Book | null>(null)
  const audioRef = useRef<HTMLAudioElement | null>(null)

  const play = useCallback(() => {
    if (audioRef.current) {
      audioRef.current.play()
      setIsPlaying(true)
    }
  }, [])

  const pause = useCallback(() => {
    if (audioRef.current) {
      audioRef.current.pause()
      setIsPlaying(false)
    }
  }, [])

  const stop = useCallback(() => {
    if (audioRef.current) {
      audioRef.current.pause()
      audioRef.current.currentTime = 0
      setIsPlaying(false)
      setCurrentTime(0)
    }
  }, [])

  const skipForward = useCallback(() => {
    if (audioRef.current) {
      const newTime = audioRef.current.currentTime + 10
      if (newTime < audioRef.current.duration) {
        audioRef.current.currentTime = newTime
      } else {
        // Move to next segment if available
        const nextSegment = currentSegment + 1
        if (nextSegment < segments.length) {
          playSegment(nextSegment)
        }
      }
    }
  }, [currentSegment, segments])

  const skipBackward = useCallback(() => {
    if (audioRef.current) {
      const newTime = audioRef.current.currentTime - 10
      if (newTime > 0) {
        audioRef.current.currentTime = newTime
      } else {
        // Move to previous segment if available
        const prevSegment = currentSegment - 1
        if (prevSegment >= 0) {
          playSegment(prevSegment)
        }
      }
    }
  }, [currentSegment])

  const setAudioSource = useCallback((url: string) => {
    if (audioRef.current) {
      audioRef.current.src = url
      audioRef.current.load()
      setCurrentTime(0)
      setDuration(0)
      setIsPlaying(false)
    }
  }, [])

  const seekTo = useCallback((time: number) => {
    if (audioRef.current) {
      audioRef.current.currentTime = time
    }
  }, [])

  const setSpeed = useCallback((speed: number) => {
    if (audioRef.current) {
      audioRef.current.playbackRate = speed
      setPlaybackSpeed(speed)
    }
  }, [])

  const preloadNextSegments = useCallback((currentIndex: number) => {
    const nextSegments = segments.slice(
      currentIndex + 1,
      currentIndex + 4
    )
    nextSegments.forEach(segment => {
      audioCache.preload(segment.audioUrl, segment.id)
    })
  }, [segments])

  const loadBookAudio = useCallback(async (book: Book) => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/books/${book.id}/audio-segments`)
      if (!response.ok) throw new Error('Failed to fetch audio segments')
      const audioSegments: AudioSegment[] = await response.json()
      
      // Filter only completed segments
      const completedSegments = audioSegments.filter(
        segment => segment.status === 'completed'
      )
      
      setSegments(completedSegments)
      setCurrentBook(book)
      setCurrentSegment(0)
      
      // Clear previous cache
      audioCache.clear()
      
      // Preload first few segments
      completedSegments.slice(0, 3).forEach(segment => {
        audioCache.preload(segment.audioUrl, segment.id)
      })
      
      // Set first segment as current
      if (completedSegments.length > 0) {
        setAudioSource(completedSegments[0].audioUrl)
      }
    } catch (err) {
      console.error('Error loading book audio:', err)
    }
  }, [setAudioSource])

  const playSegment = useCallback((segmentIndex: number) => {
    if (segmentIndex >= 0 && segmentIndex < segments.length) {
      const segment = segments[segmentIndex]
      const cachedAudio = audioCache.get(segment.id)
      
      setCurrentSegment(segmentIndex)
      if (cachedAudio) {
        if (audioRef.current) {
          audioRef.current.src = cachedAudio.src
          audioRef.current.load()
        }
      } else {
        setAudioSource(segment.audioUrl)
      }
      play()
    }
  }, [segments, setAudioSource, play])

  // Handle segment transitions
  useEffect(() => {
    if (audioRef.current) {
      const handleEnded = () => {
        const nextSegment = currentSegment + 1
        if (nextSegment < segments.length) {
          playSegment(nextSegment)
        } else {
          setIsPlaying(false)
          setCurrentTime(0)
        }
      }

      audioRef.current.addEventListener('ended', handleEnded)
      return () => {
        audioRef.current?.removeEventListener('ended', handleEnded)
      }
    }
  }, [currentSegment, segments, playSegment])

  // Preload next segments when current segment changes
  useEffect(() => {
    if (currentSegment >= 0) {
      preloadNextSegments(currentSegment)
    }
  }, [currentSegment, preloadNextSegments])

  // Initialize audio element
  useEffect(() => {
    if (typeof window !== 'undefined' && !audioRef.current) {
      audioRef.current = new Audio()
      audioRef.current.addEventListener('timeupdate', () => {
        setCurrentTime(audioRef.current?.currentTime || 0)
      })
      audioRef.current.addEventListener('loadedmetadata', () => {
        setDuration(audioRef.current?.duration || 0)
      })
    }

    return () => {
      if (audioRef.current) {
        audioRef.current.src = ''
        audioRef.current = null
      }
      audioCache.clear()
    }
  }, [])

  return (
    <TTSContext.Provider
      value={{
        isPlaying,
        currentTime,
        duration,
        currentSegment,
        segments,
        currentBook,
        play,
        pause,
        stop,
        skipForward,
        skipBackward,
        setAudioSource,
        seekTo,
        setSpeed,
        loadBookAudio,
        playSegment,
      }}
    >
      {children}
    </TTSContext.Provider>
  )
}

export function useTTS() {
  const context = useContext(TTSContext)
  if (context === undefined) {
    throw new Error('useTTS must be used within a TTSProvider')
  }
  return context
} 
```

## File: /Users/saint/Desktop/pdf-player/app/layout.tsx

```tsx
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

```

## File: /Users/saint/Desktop/pdf-player/app/api/uploadthing/core.ts

```ts
import { createUploadthing, type FileRouter } from "uploadthing/next"

const f = createUploadthing()

// Verify all required environment variables are present
const requiredEnvVars = {
  UPLOADTHING_SECRET: process.env.UPLOADTHING_SECRET,
  UPLOADTHING_APP_ID: process.env.UPLOADTHING_APP_ID,
  UPLOADTHING_TOKEN: process.env.UPLOADTHING_TOKEN,
}

// Log environment variable status (without exposing values)
console.log("UploadThing Environment Variables Status:", 
  Object.keys(requiredEnvVars).reduce((acc, key) => ({
    ...acc,
    [key]: requiredEnvVars[key as keyof typeof requiredEnvVars] ? "✓ Set" : "✗ Missing"
  }), {})
)

export const ourFileRouter = {
  pdfUploader: f({ pdf: { maxFileSize: "32MB" } })
    .middleware(async () => {
      // This code runs on your server before upload
      console.log("UploadThing middleware executing...")
      
      // Verify environment variables
      const missingVars = Object.entries(requiredEnvVars)
        .filter(([_, value]) => !value)
        .map(([key]) => key)

      if (missingVars.length > 0) {
        const error = `Missing required environment variables: ${missingVars.join(", ")}`
        console.error(error)
        throw new Error(error)
      }

      console.log("All environment variables verified")
      return { uploadedAt: new Date() }
    })
    .onUploadComplete(async ({ metadata, file }) => {
      // This code RUNS ON YOUR SERVER after upload
      console.log("Upload complete. File details:", {
        url: file.url,
        name: file.name,
        size: file.size,
        uploadedAt: metadata.uploadedAt
      })
      
      return { url: file.url }
    }),
} satisfies FileRouter

export type OurFileRouter = typeof ourFileRouter 
```

## File: /Users/saint/Desktop/pdf-player/app/api/uploadthing/route.ts

```ts
import { createRouteHandler } from "uploadthing/next"
import { ourFileRouter } from "./core"

// Export routes for Next App Router
export const { GET, POST } = createRouteHandler({
  router: ourFileRouter,
}) 
```

## File: /Users/saint/Desktop/pdf-player/app/page.tsx

```tsx
import { Metadata } from "next"
import dynamic from "next/dynamic"

export const metadata: Metadata = {
  title: "PDF Player",
  description: "A modern PDF reader with text-to-speech capabilities",
}

const HomePage = dynamic(() => import("../components/HomePage"), {
  ssr: false,
})

export default function Page() {
  return <HomePage />
}

```

## File: /Users/saint/Desktop/pdf-player/app/globals.css

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
  :root {
    --background: 0 0% 100%;
    --foreground: 0 0% 3.9%;
    --card: 0 0% 100%;
    --card-foreground: 0 0% 3.9%;
    --popover: 0 0% 100%;
    --popover-foreground: 0 0% 3.9%;
    --primary: 0 0% 9%;
    --primary-foreground: 0 0% 98%;
    --secondary: 0 0% 96.1%;
    --secondary-foreground: 0 0% 9%;
    --muted: 0 0% 96.1%;
    --muted-foreground: 0 0% 45.1%;
    --accent: 0 0% 96.1%;
    --accent-foreground: 0 0% 9%;
    --destructive: 0 84.2% 60.2%;
    --destructive-foreground: 0 0% 98%;
    --border: 0 0% 89.8%;
    --input: 0 0% 89.8%;
    --ring: 0 0% 3.9%;
    --radius: 1rem;
  }

  .dark {
    --background: 0 0% 3.9%;
    --foreground: 0 0% 98%;
    --card: 0 0% 3.9%;
    --card-foreground: 0 0% 98%;
    --popover: 0 0% 3.9%;
    --popover-foreground: 0 0% 98%;
    --primary: 0 0% 98%;
    --primary-foreground: 0 0% 9%;
    --secondary: 0 0% 14.9%;
    --secondary-foreground: 0 0% 98%;
    --muted: 0 0% 14.9%;
    --muted-foreground: 0 0% 63.9%;
    --accent: 0 0% 14.9%;
    --accent-foreground: 0 0% 98%;
    --destructive: 0 62.8% 30.6%;
    --destructive-foreground: 0 0% 98%;
    --border: 0 0% 14.9%;
    --input: 0 0% 14.9%;
    --ring: 0 0% 83.1%;
  }
}

@layer utilities {
  .glass-effect {
    @apply bg-white/10 dark:bg-black/10 backdrop-blur-xl border border-white/20 dark:border-white/10;
  }

  .glass-effect-strong {
    @apply bg-white/20 dark:bg-black/20 backdrop-blur-2xl border border-white/30 dark:border-white/20;
  }
}

.progress-gradient {
  background: linear-gradient(90deg, rgba(255, 255, 255, 0.2) 0%, rgba(255, 255, 255, 0.1) 100%);
}

.player-shadow {
  box-shadow: 0 8px 32px 0 rgba(31, 38, 135, 0.07);
}

```

## File: /Users/saint/Desktop/pdf-player/postcss.config.mjs

```mjs
/** @type {import('postcss-load-config').Config} */
const config = {
  plugins: {
    tailwindcss: {},
  },
};

export default config;
```

## File: /Users/saint/Desktop/pdf-player/next.config.mjs

```mjs
let userConfig = undefined
try {
  userConfig = await import('./v0-user-next.config')
} catch (e) {
  // ignore error
}

/** @type {import('next').NextConfig} */
const nextConfig = {
  eslint: {
    ignoreDuringBuilds: true,
  },
  typescript: {
    ignoreBuildErrors: true,
  },
  images: {
    unoptimized: true,
  },
  experimental: {
    webpackBuildWorker: true,
    parallelServerBuildTraces: true,
    parallelServerCompiles: true,
  },
}

mergeConfig(nextConfig, userConfig)

function mergeConfig(nextConfig, userConfig) {
  if (!userConfig) {
    return
  }

  for (const key in userConfig) {
    if (
      typeof nextConfig[key] === 'object' &&
      !Array.isArray(nextConfig[key])
    ) {
      nextConfig[key] = {
        ...nextConfig[key],
        ...userConfig[key],
      }
    } else {
      nextConfig[key] = userConfig[key]
    }
  }
}

export default nextConfig
```

## File: /Users/saint/Desktop/pdf-player/utils/audioCache.ts

```ts
"use client"

class AudioCache {
  private cache: Map<string, HTMLAudioElement>
  private maxSize: number

  constructor(maxSize = 10) {
    this.cache = new Map()
    this.maxSize = maxSize
  }

  get(segmentId: string): HTMLAudioElement | undefined {
    return this.cache.get(segmentId)
  }

  set(segmentId: string, audio: HTMLAudioElement): void {
    if (this.cache.size >= this.maxSize) {
      // Remove oldest entry
      const firstKey = this.cache.keys().next().value
      const oldAudio = this.cache.get(firstKey)
      if (oldAudio) {
        oldAudio.src = '' // Clear source to help garbage collection
      }
      this.cache.delete(firstKey)
    }
    this.cache.set(segmentId, audio)
  }

  preload(segmentUrl: string, segmentId: string): void {
    if (!this.cache.has(segmentId)) {
      const audio = new Audio()
      audio.src = segmentUrl
      audio.load() // Preload without playing
      this.set(segmentId, audio)
    }
  }

  clear(): void {
    this.cache.forEach(audio => {
      audio.src = '' // Clear source to help garbage collection
    })
    this.cache.clear()
  }
}

export const audioCache = new AudioCache() 
```

## File: /Users/saint/Desktop/pdf-player/backend/models.go

```go
package main

import (
	"database/sql"
	"time"
)

// Book represents a PDF book in the system
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	CoverURL    string    `json:"coverUrl"`
	FileURL     string    `json:"fileUrl"`
	PageCount   int       `json:"pageCount"`
	CurrentPage int       `json:"currentPage"`
	Language    string    `json:"language"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Categories  []string  `json:"categories,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
}

// ReadingProgress tracks a user's reading progress for a book
type ReadingProgress struct {
	ID                string  `json:"id"`
	BookID            string  `json:"bookId"`
	UserID            string  `json:"userId"`
	CurrentPage       int     `json:"currentPage"`
	TotalPages        int     `json:"totalPages"`
	CompletionPercent float64 `json:"completionPercent"`
	LastReadAt        string  `json:"lastReadAt"`
}

// Bookmark represents a user's bookmark in a book
type Bookmark struct {
	ID         string    `json:"id"`
	BookID     string    `json:"bookId"`
	UserID     string    `json:"userId"`
	PageNumber int       `json:"pageNumber"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// AudioSegment represents a segment of text and its corresponding audio
type AudioSegment struct {
	ID        string    `json:"id"`
	BookID    string    `json:"bookId"`
	Content   string    `json:"content"`
	AudioURL  string    `json:"audioUrl"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// Category represents a book category
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Tag represents a book tag
type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Request/Response structures
type UpdateProgressRequest struct {
	CurrentPage int     `json:"currentPage"`
	TotalPages  int     `json:"totalPages"`
	Completion  float64 `json:"completion"`
}

type CreateBookmarkRequest struct {
	PageNumber int    `json:"pageNumber"`
	Note       string `json:"note"`
}

type GenerateAudioRequest struct {
	BookID    string `json:"bookId"`
	StartPage int    `json:"startPage"`
	EndPage   int    `json:"endPage"`
}

// UpdateBook updates an existing book in the database
func UpdateBook(db *sql.DB, book *Book) error {
	query := `
		UPDATE books 
		SET title = ?, author = ?, file_url = ?, 
			page_count = ?, language = ?, updated_at = ?
		WHERE id = ?`

	_, err := db.Exec(query,
		book.Title, book.Author, book.FileURL,
		book.PageCount, book.Language, book.UpdatedAt,
		book.ID)

	return err
}

// SaveAudioSegment saves a new audio segment to the database
func SaveAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, content, audio_url,
			status, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		segment.CreatedAt,
		segment.UpdatedAt,
	)

	return err
}
```

## File: /Users/saint/Desktop/pdf-player/backend/config.go

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port       string
	BackendURL string
	Env        string

	// Database
	DBPath string

	// File Storage
	UploadDir     string
	PDFDir        string
	CoverDir      string
	MaxUploadSize int64

	// Replicate API
	ReplicateAPIToken  string
	ReplicateAPIURL    string
	KokoroModelVersion string

	// UploadThing
	UploadThingURL    string
	UploadThingToken  string
	UploadThingSecret string
	UploadThingAppID  string

	// TTS
	TTSMaxChunkSize   int
	TTSDefaultLang    string
	TTSDefaultSpeaker int
	TTSDefaultSpeed   float64
	TTSTopK           int
	TTSTopP           float64
	TTSTemperature    float64

	// CORS
	AllowedOrigins []string
}

var AppConfig Config

// LoadConfig loads the configuration from environment variables
func LoadConfig() error {
	// Load .env file if it exists
	godotenv.Load()

	AppConfig = Config{
		// Server
		Port:       getEnv("PORT", "8080"),
		BackendURL: getEnv("BACKEND_URL", "http://localhost:8080"),
		Env:        getEnv("ENV", "development"),

		// Database
		DBPath: getEnv("DB_PATH", "./ereader.db"),

		// File Storage
		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		PDFDir:        getEnv("PDF_DIR", "./uploads/pdfs"),
		CoverDir:      getEnv("COVER_DIR", "./uploads/covers"),
		MaxUploadSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 32) * 1024 * 1024, // Convert MB to bytes

		// Replicate API
		ReplicateAPIToken:  getEnvRequired("REPLICATE_API_TOKEN"),
		ReplicateAPIURL:    getEnv("REPLICATE_API_URL", "https://api.replicate.com/v1"),
		KokoroModelVersion: getEnv("KOKORO_MODEL_VERSION", "82c9a78a6fff5fa0e4acbd0f8eb453e0c4b6d5d90e316b544b8c7e04c42fba42"),

		// UploadThing
		UploadThingURL:    getEnv("UPLOADTHING_URL", "https://uploadthing.com/api"),
		UploadThingToken:  getEnvRequired("UPLOADTHING_TOKEN"),
		UploadThingSecret: getEnvRequired("UPLOADTHING_SECRET"),
		UploadThingAppID:  getEnvRequired("UPLOADTHING_APP_ID"),

		// TTS
		TTSMaxChunkSize:   getEnvAsInt("TTS_MAX_CHUNK_SIZE", 1000),
		TTSDefaultLang:    getEnv("TTS_DEFAULT_LANGUAGE", "en"),
		TTSDefaultSpeaker: getEnvAsInt("TTS_DEFAULT_SPEAKER_ID", 0),
		TTSDefaultSpeed:   getEnvAsFloat64("TTS_DEFAULT_SPEED", 1.0),
		TTSTopK:           getEnvAsInt("TTS_TOP_K", 50),
		TTSTopP:           getEnvAsFloat64("TTS_TOP_P", 0.8),
		TTSTemperature:    getEnvAsFloat64("TTS_TEMPERATURE", 0.8),

		// CORS
		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
	}

	// Create required directories
	for _, dir := range []string{AppConfig.UploadDir, AppConfig.PDFDir, AppConfig.CoverDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return nil
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsFloat64(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}
```

## File: /Users/saint/Desktop/pdf-player/backend/go.mod

```mod
module backend

go 1.21

require (
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06
	github.com/mattn/go-sqlite3 v1.14.24
)

require github.com/gorilla/websocket v1.5.3
```

## File: /Users/saint/Desktop/pdf-player/backend/jockey-img/backend_01-22-25.md

````md
# Jockey Image

Generated: 01-22-2025 at 11:12:42

## Repository Structure

```
backend
│   ├── models.go
│   ├── main.go
│   ├── config.go
│   ├── go.mod
│   ├── db.go
│   ├── pdf.go
│   ├── tts.go
│   ├── storage.go
│   ├── utils.go
│   ├── go.sum
│   ├── schema.sql
│   ├── README.md
    └── pdf-player
```

## File: /Users/saint/Desktop/pdf-player/backend/models.go

```go
package main

import "time"

// Book represents a PDF book in the system
type Book struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	CoverURL    string    `json:"coverUrl"`
	Content     string    `json:"content"`
	FilePath    string    `json:"filePath"`
	PageCount   int       `json:"pageCount"`
	CurrentPage int       `json:"currentPage"`
	Language    string    `json:"language"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Categories  []string  `json:"categories,omitempty"`
	Tags        []string  `json:"tags,omitempty"`
}

// ReadingProgress tracks a user's reading progress for a book
type ReadingProgress struct {
	ID                string  `json:"id"`
	BookID            string  `json:"bookId"`
	UserID            string  `json:"userId"`
	CurrentPage       int     `json:"currentPage"`
	TotalPages        int     `json:"totalPages"`
	CompletionPercent float64 `json:"completionPercent"`
	LastReadAt        string  `json:"lastReadAt"`
}

// Bookmark represents a user's bookmark in a book
type Bookmark struct {
	ID         string    `json:"id"`
	BookID     string    `json:"bookId"`
	UserID     string    `json:"userId"`
	PageNumber int       `json:"pageNumber"`
	Note       string    `json:"note"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// AudioSegment represents a segment of text converted to audio
type AudioSegment struct {
	ID            string    `json:"id"`
	BookID        string    `json:"bookId"`
	SegmentNumber int       `json:"segmentNumber"`
	Content       string    `json:"content"`
	AudioURL      string    `json:"audioUrl"`
	Duration      float64   `json:"duration"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
}

// Category represents a book category
type Category struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
}

// Tag represents a book tag
type Tag struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

// Request/Response structures
type UpdateProgressRequest struct {
	CurrentPage int     `json:"currentPage"`
	TotalPages  int     `json:"totalPages"`
	Completion  float64 `json:"completion"`
}

type CreateBookmarkRequest struct {
	PageNumber int    `json:"pageNumber"`
	Note       string `json:"note"`
}

type GenerateAudioRequest struct {
	BookID    string `json:"bookId"`
	StartPage int    `json:"startPage"`
	EndPage   int    `json:"endPage"`
}
```

## File: /Users/saint/Desktop/pdf-player/backend/config.go

```go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	// Server
	Port       string
	BackendURL string
	Env        string

	// Database
	DBPath string

	// File Storage
	UploadDir     string
	PDFDir        string
	CoverDir      string
	MaxUploadSize int64

	// Replicate API
	ReplicateAPIToken  string
	ReplicateAPIURL    string
	KokoroModelVersion string

	// TTS
	TTSMaxChunkSize   int
	TTSDefaultLang    string
	TTSDefaultSpeaker int
	TTSDefaultSpeed   float64
	TTSTopK           int
	TTSTopP           float64
	TTSTemperature    float64

	// CORS
	AllowedOrigins []string
}

var AppConfig Config

// LoadConfig loads the configuration from environment variables
func LoadConfig() error {
	// Load .env file if it exists
	godotenv.Load()

	AppConfig = Config{
		// Server
		Port:       getEnv("PORT", "8080"),
		BackendURL: getEnv("BACKEND_URL", "http://localhost:8080"),
		Env:        getEnv("ENV", "development"),

		// Database
		DBPath: getEnv("DB_PATH", "./ereader.db"),

		// File Storage
		UploadDir:     getEnv("UPLOAD_DIR", "./uploads"),
		PDFDir:        getEnv("PDF_DIR", "./uploads/pdfs"),
		CoverDir:      getEnv("COVER_DIR", "./uploads/covers"),
		MaxUploadSize: getEnvAsInt64("MAX_UPLOAD_SIZE", 32) * 1024 * 1024, // Convert MB to bytes

		// Replicate API
		ReplicateAPIToken:  getEnvRequired("REPLICATE_API_TOKEN"),
		ReplicateAPIURL:    getEnv("REPLICATE_API_URL", "https://api.replicate.com/v1"),
		KokoroModelVersion: getEnv("KOKORO_MODEL_VERSION", "82c9a78a6fff5fa0e4acbd0f8eb453e0c4b6d5d90e316b544b8c7e04c42fba42"),

		// TTS
		TTSMaxChunkSize:   getEnvAsInt("TTS_MAX_CHUNK_SIZE", 1000),
		TTSDefaultLang:    getEnv("TTS_DEFAULT_LANGUAGE", "en"),
		TTSDefaultSpeaker: getEnvAsInt("TTS_DEFAULT_SPEAKER_ID", 0),
		TTSDefaultSpeed:   getEnvAsFloat64("TTS_DEFAULT_SPEED", 1.0),
		TTSTopK:           getEnvAsInt("TTS_TOP_K", 50),
		TTSTopP:           getEnvAsFloat64("TTS_TOP_P", 0.8),
		TTSTemperature:    getEnvAsFloat64("TTS_TEMPERATURE", 0.8),

		// CORS
		AllowedOrigins: strings.Split(getEnv("ALLOWED_ORIGINS", "http://localhost:3000"), ","),
	}

	// Create required directories
	for _, dir := range []string{AppConfig.UploadDir, AppConfig.PDFDir, AppConfig.CoverDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return nil
}

// Helper functions to get environment variables
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvRequired(key string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		panic(fmt.Sprintf("Required environment variable %s is not set", key))
	}
	return value
}

func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsInt64(key string, defaultValue int64) int64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
		return value
	}
	return defaultValue
}

func getEnvAsFloat64(key string, defaultValue float64) float64 {
	valueStr := getEnv(key, "")
	if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
		return value
	}
	return defaultValue
}
```

## File: /Users/saint/Desktop/pdf-player/backend/go.mod

```mod
module github.com/saint/pdf-player

go 1.21

require (
	github.com/gorilla/mux v1.8.1
	github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06
	github.com/mattn/go-sqlite3 v1.14.24
)

require github.com/joho/godotenv v1.5.1

require github.com/google/uuid v1.6.0
```

## File: /Users/saint/Desktop/pdf-player/backend/db.go

```go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

// InitDB initializes the database connection and creates tables
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./ereader.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Read and execute schema.sql
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("error reading schema file: %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("error creating schema: %v", err)
	}

	return nil
}

// GetBookByID retrieves a book by its ID
func GetBookByID(db *sql.DB, id string) (*Book, error) {
	var book Book
	query := `
		SELECT id, title, author, cover_url, content, file_path, 
			page_count, current_page, language, 
			created_at, updated_at
		FROM books WHERE id = ?
	`

	err := db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.CoverURL,
		&book.Content,
		&book.FilePath,
		&book.PageCount,
		&book.CurrentPage,
		&book.Language,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("book not found")
	}
	if err != nil {
		return nil, err
	}

	// Get categories
	rows, err := db.Query(`
		SELECT c.name FROM categories c 
		JOIN book_categories bc ON c.id = bc.category_id 
		WHERE bc.book_id = ?`, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var category string
			rows.Scan(&category)
			book.Categories = append(book.Categories, category)
		}
	}

	// Get tags
	rows, err = db.Query(`
		SELECT t.name FROM tags t 
		JOIN book_tags bt ON t.id = bt.tag_id 
		WHERE bt.book_id = ?`, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var tag string
			rows.Scan(&tag)
			book.Tags = append(book.Tags, tag)
		}
	}

	return &book, nil
}

// SaveBook saves a new book to the database
func SaveBook(db *sql.DB, book *Book) error {
	query := `
		INSERT INTO books (
			id, title, author, cover_url, content, file_path,
			page_count, current_page, language, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		book.ID,
		book.Title,
		book.Author,
		book.CoverURL,
		book.Content,
		book.FilePath,
		book.PageCount,
		book.CurrentPage,
		book.Language,
		formatDateTime(book.CreatedAt),
		formatDateTime(book.UpdatedAt),
	)

	return err
}

// UpdateReadingProgress updates or creates a reading progress record
func UpdateReadingProgress(db *sql.DB, progress *ReadingProgress) error {
	query := `
		INSERT INTO reading_progress (id, book_id, user_id, current_page, total_pages, completion_percent, last_read_at)
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(book_id, user_id) DO UPDATE SET
			current_page = excluded.current_page,
			total_pages = excluded.total_pages,
			completion_percent = excluded.completion_percent,
			last_read_at = CURRENT_TIMESTAMP
	`

	_, err := db.Exec(query,
		progress.ID,
		progress.BookID,
		progress.UserID,
		progress.CurrentPage,
		progress.TotalPages,
		progress.CompletionPercent,
	)

	return err
}

// GetReadingProgress retrieves reading progress for a book and user
func GetReadingProgress(db *sql.DB, bookID, userID string) (*ReadingProgress, error) {
	var progress ReadingProgress
	query := `
		SELECT id, book_id, user_id, current_page, total_pages, completion_percent, last_read_at
		FROM reading_progress
		WHERE book_id = ? AND user_id = ?
	`

	err := db.QueryRow(query, bookID, userID).Scan(
		&progress.ID,
		&progress.BookID,
		&progress.UserID,
		&progress.CurrentPage,
		&progress.TotalPages,
		&progress.CompletionPercent,
		&progress.LastReadAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

// CreateBookmark creates a new bookmark
func CreateBookmark(db *sql.DB, bookmark *Bookmark) error {
	query := `
		INSERT INTO bookmarks (id, book_id, user_id, page_number, note, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`

	_, err := db.Exec(query,
		bookmark.ID,
		bookmark.BookID,
		bookmark.UserID,
		bookmark.PageNumber,
		bookmark.Note,
	)

	return err
}

// GetBookmarks retrieves all bookmarks for a book and user
func GetBookmarks(db *sql.DB, bookID, userID string) ([]Bookmark, error) {
	query := `
		SELECT id, book_id, user_id, page_number, note, created_at, updated_at
		FROM bookmarks
		WHERE book_id = ? AND user_id = ?
		ORDER BY page_number ASC
	`

	rows, err := db.Query(query, bookID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []Bookmark
	for rows.Next() {
		var bookmark Bookmark
		err := rows.Scan(
			&bookmark.ID,
			&bookmark.BookID,
			&bookmark.UserID,
			&bookmark.PageNumber,
			&bookmark.Note,
			&bookmark.CreatedAt,
			&bookmark.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

// DeleteBookmark deletes a bookmark
func DeleteBookmark(db *sql.DB, id string) error {
	query := "DELETE FROM bookmarks WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

// UpdateBookmark updates a bookmark's note
func UpdateBookmark(db *sql.DB, bookmark *Bookmark) error {
	query := `
		UPDATE bookmarks 
		SET note = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND book_id = ? AND user_id = ?
	`

	result, err := db.Exec(query, bookmark.Note, bookmark.ID, bookmark.BookID, bookmark.UserID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("bookmark not found")
	}

	return nil
}

// CreateAudioSegment creates a new audio segment
func CreateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, segment_number, content, audio_url, duration, status, created_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
	`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.SegmentNumber,
		segment.Content,
		segment.AudioURL,
		segment.Duration,
		segment.Status,
	)

	return err
}

// UpdateAudioSegment updates an existing audio segment
func UpdateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		UPDATE audio_segments 
		SET audio_url = ?, duration = ?, status = ?
		WHERE id = ? AND book_id = ?
	`

	result, err := db.Exec(query,
		segment.AudioURL,
		segment.Duration,
		segment.Status,
		segment.ID,
		segment.BookID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("audio segment not found")
	}

	return nil
}

// Category operations
func CreateCategory(db *sql.DB, category *Category) error {
	category.ID = generateID()
	category.CreatedAt = time.Now()

	_, err := db.Exec(`
		INSERT INTO categories (id, name, description, created_at)
		VALUES (?, ?, ?, ?)`,
		category.ID, category.Name, category.Description, category.CreatedAt)

	return err
}

func AddBookToCategory(db *sql.DB, bookID, categoryID string) error {
	_, err := db.Exec(`
		INSERT INTO book_categories (book_id, category_id)
		VALUES (?, ?)`, bookID, categoryID)

	return err
}

// Tag operations
func CreateTag(db *sql.DB, tag *Tag) error {
	tag.ID = generateID()
	tag.CreatedAt = time.Now()

	_, err := db.Exec(`
		INSERT INTO tags (id, name, created_at)
		VALUES (?, ?, ?)`,
		tag.ID, tag.Name, tag.CreatedAt)

	return err
}

func AddBookTag(db *sql.DB, bookID, tagID string) error {
	_, err := db.Exec(`
		INSERT INTO book_tags (book_id, tag_id)
		VALUES (?, ?)`, bookID, tagID)

	return err
}
```

## File: /Users/saint/Desktop/pdf-player/backend/schema.sql

```sql
-- Books table with enhanced metadata
CREATE TABLE IF NOT EXISTS books (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT NOT NULL,
    cover_url TEXT NOT NULL,
    content TEXT NOT NULL,
    file_path TEXT,
    page_count INTEGER DEFAULT 0,
    current_page INTEGER DEFAULT 0,
    language TEXT DEFAULT 'en',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories for organizing books
CREATE TABLE IF NOT EXISTS categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book-Category relationship (many-to-many)
CREATE TABLE IF NOT EXISTS book_categories (
    book_id TEXT,
    category_id TEXT,
    PRIMARY KEY (book_id, category_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- Tags for flexible book organization
CREATE TABLE IF NOT EXISTS tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book-Tag relationship (many-to-many)
CREATE TABLE IF NOT EXISTS book_tags (
    book_id TEXT,
    tag_id TEXT,
    PRIMARY KEY (book_id, tag_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- Audio synthesis tracking
CREATE TABLE IF NOT EXISTS audio_segments (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    segment_number INTEGER NOT NULL,
    content TEXT NOT NULL,
    audio_url TEXT,
    duration FLOAT,
    status TEXT CHECK (status IN ('pending', 'processing', 'completed', 'failed')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Reading progress tracking
CREATE TABLE IF NOT EXISTS reading_progress (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    current_page INTEGER DEFAULT 0,
    completion_percentage FLOAT DEFAULT 0,
    last_read_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    UNIQUE(book_id)
);

-- Bookmarks
CREATE TABLE IF NOT EXISTS bookmarks (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    page_number INTEGER NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
CREATE INDEX IF NOT EXISTS idx_reading_progress_book ON reading_progress(book_id);
CREATE INDEX IF NOT EXISTS idx_audio_segments_book ON audio_segments(book_id);
CREATE INDEX IF NOT EXISTS idx_bookmarks_book ON bookmarks(book_id); 
```

## File: /Users/saint/Desktop/pdf-player/backend/pdf.go

```go
package main

import (
	"fmt"
	"mime/multipart"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/ledongthuc/pdf"
)

// processPDF processes a PDF file and returns a Book object
func processPDF(file multipart.File, filename string) (*Book, error) {
	// Save PDF file temporarily
	filePath, err := fileStorage.SavePDF(file, &multipart.FileHeader{
		Filename: filename,
	})
	if err != nil {
		return nil, fmt.Errorf("error saving PDF: %v", err)
	}

	// Open PDF file for reading
	f, r, err := pdf.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer f.Close()

	// Get total number of pages
	totalPages := r.NumPage()

	// Extract text from all pages
	var content strings.Builder
	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		page := r.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return nil, fmt.Errorf("error extracting text from page %d: %v", pageNum, err)
		}
		content.WriteString(text)
	}

	// Create book object
	book := &Book{
		ID:          uuid.New().String(),
		Title:       strings.TrimSuffix(filename, ".pdf"),
		Author:      "Unknown", // Can be updated later
		FilePath:    filePath,
		Content:     content.String(),
		PageCount:   totalPages,
		CurrentPage: 0,
		Language:    "en", // Default to English, can be updated later
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return book, nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/tts.go

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// TTSRequest represents a request to the Replicate API
type TTSRequest struct {
	Version string   `json:"version"`
	Input   TTSInput `json:"input"`
}

// TTSInput represents the input parameters for the TTS model
type TTSInput struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

// TTSResponse represents a response from the Replicate API
type TTSResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Output string `json:"output"`
	Error  string `json:"error"`
}

// processAudioSegment processes a text segment and generates audio
func processAudioSegment(segment *AudioSegment) error {
	// Get Kokoro model version from environment
	modelVersion := os.Getenv("KOKORO_MODEL_VERSION")
	if modelVersion == "" {
		return fmt.Errorf("KOKORO_MODEL_VERSION not set")
	}

	// Prepare request
	reqBody := TTSRequest{
		Version: modelVersion,
		Input: TTSInput{
			Text:     segment.Content,
			Language: "en", // Default to English for now
		},
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", os.Getenv("REPLICATE_API_URL")+"/predictions", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	var ttsResp TTSResponse
	if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Check for error
	if ttsResp.Error != "" {
		return fmt.Errorf("TTS error: %s", ttsResp.Error)
	}

	// Poll for completion
	for ttsResp.Status != "succeeded" {
		time.Sleep(2 * time.Second)

		// Get prediction status
		req, err = http.NewRequest("GET", os.Getenv("REPLICATE_API_URL")+"/predictions/"+ttsResp.ID, nil)
		if err != nil {
			return fmt.Errorf("error creating status request: %v", err)
		}

		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return fmt.Errorf("error checking status: %v", err)
		}

		if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
			resp.Body.Close()
			return fmt.Errorf("error decoding status: %v", err)
		}
		resp.Body.Close()

		if ttsResp.Status == "failed" {
			return fmt.Errorf("TTS processing failed: %s", ttsResp.Error)
		}
	}

	// Download audio file
	resp, err = http.Get(ttsResp.Output)
	if err != nil {
		return fmt.Errorf("error downloading audio: %v", err)
	}
	defer resp.Body.Close()

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "tts-*.mp3")
	if err != nil {
		return fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy audio to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("error saving audio: %v", err)
	}

	// Save audio file using storage
	audioFile, err := os.Open(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("error opening temp file: %v", err)
	}
	defer audioFile.Close()

	// Save audio and get URL
	audioURL, err := fileStorage.SaveAudio(audioFile, "output.mp3")
	if err != nil {
		return fmt.Errorf("error saving audio file: %v", err)
	}

	// Update segment with audio URL and status
	segment.AudioURL = audioURL
	segment.Status = "completed"

	// Update database
	if err := UpdateAudioSegment(db, segment); err != nil {
		return fmt.Errorf("error updating segment: %v", err)
	}

	return nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/go.sum

```sum
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/gorilla/mux v1.8.1 h1:TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY=
github.com/gorilla/mux v1.8.1/go.mod h1:AKf9I4AEqPTmMytcMc0KkNouC66V3BtZ4qD5fmWSiMQ=
github.com/joho/godotenv v1.5.1 h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=
github.com/joho/godotenv v1.5.1/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06 h1:kacRlPN7EN++tVpGUorNGPn/4DnB7/DfTY82AOn6ccU=
github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06/go.mod h1:imJHygn/1yfhB7XSJJKlFZKl/J+dCPAknuiaGOshXAs=
github.com/mattn/go-sqlite3 v1.14.24 h1:tpSp2G2KyMnnQu99ngJ47EIkWVmliIizyZBfPrBWDRM=
github.com/mattn/go-sqlite3 v1.14.24/go.mod h1:Uh1q+B4BYcTPb+yiD3kU8Ct7aC0hY9fxUwlHK0RXw+Y=
```

## File: /Users/saint/Desktop/pdf-player/backend/README.md

````md
# PDF Player Backend

This is the backend service for the PDF Player application. It provides APIs for uploading PDFs, extracting text, managing books, and tracking reading progress.

## Prerequisites

- Go 1.16 or later
- SQLite3

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run *.go
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

## API Endpoints

### Books
- **POST** `/api/upload` - Upload a PDF file
  - Content-Type: `multipart/form-data`
  - Form field: `file` (PDF file)
  - Returns: Book object with extracted text

- **GET** `/api/books` - Get all books
  - Returns: Array of book objects

- **GET** `/api/book/{id}` - Get a specific book
  - Returns: Single book object

### Reading Progress
- **PUT** `/api/book/{id}/progress` - Update reading progress
  - Body: `{ "currentPage": number, "completion": number }`

- **GET** `/api/book/{id}/progress` - Get reading progress
  - Returns: Reading progress object

### Bookmarks
- **POST** `/api/book/{id}/bookmarks` - Create a bookmark
  - Body: `{ "pageNumber": number, "note": string }`
  - Returns: Created bookmark object

- **GET** `/api/book/{id}/bookmarks` - Get all bookmarks for a book
  - Returns: Array of bookmark objects

### Audio Segments
- **POST** `/api/book/{id}/audio` - Create an audio segment
  - Body: `{ "startPage": number, "endPage": number }`
  - Returns: Created audio segment object

- **PUT** `/api/book/{id}/audio/{segmentId}` - Update audio segment status
  - Body: `{ "status": string, "audioUrl": string, "duration": number }`

### Categories
- **POST** `/api/categories` - Create a category
  - Body: `{ "name": string, "description": string }`
  - Returns: Created category object

- **PUT** `/api/book/{id}/categories/{categoryId}` - Add book to category

### Tags
- **POST** `/api/tags` - Create a tag
  - Body: `{ "name": string }`
  - Returns: Created tag object

- **PUT** `/api/book/{id}/tags/{tagId}` - Add tag to book

## Data Models

### Book
```json
{
  "id": "string",
  "title": "string",
  "author": "string",
  "coverUrl": "string",
  "content": "string",
  "filePath": "string",
  "pageCount": number,
  "currentPage": number,
  "language": "string",
  "createdAt": "datetime",
  "updatedAt": "datetime",
  "categories": ["string"],
  "tags": ["string"]
}
```

### Reading Progress
```json
{
  "id": "string",
  "bookId": "string",
  "currentPage": number,
  "completionPercent": number,
  "lastReadAt": "datetime"
}
```

### Bookmark
```json
{
  "id": "string",
  "bookId": "string",
  "pageNumber": number,
  "note": "string",
  "createdAt": "datetime"
}
```

### Audio Segment
```json
{
  "id": "string",
  "bookId": "string",
  "segmentNumber": number,
  "content": "string",
  "audioUrl": "string",
  "duration": number,
  "status": "string",
  "createdAt": "datetime"
}
```

## Future Enhancements

1. PDF metadata extraction for better book information
2. Custom book cover image upload
3. Text-to-speech synthesis integration using Kokoro TTS
4. Book categories and tags for better organization
5. Full-text search functionality 
````

## File: /Users/saint/Desktop/pdf-player/backend/storage.go

```go
package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// FileStorage handles file operations
type FileStorage struct {
	baseDir  string
	pdfDir   string
	coverDir string
	audioDir string
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(baseDir string) (*FileStorage, error) {
	fs := &FileStorage{
		baseDir:  baseDir,
		pdfDir:   filepath.Join(baseDir, "pdfs"),
		coverDir: filepath.Join(baseDir, "covers"),
		audioDir: filepath.Join(baseDir, "audio"),
	}

	// Create directories if they don't exist
	for _, dir := range []string{fs.pdfDir, fs.coverDir, fs.audioDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return fs, nil
}

// SavePDF saves a PDF file and returns its path
func (fs *FileStorage) SavePDF(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	if !strings.EqualFold(ext, ".pdf") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	filename := uuid.New().String() + ext
	filePath := filepath.Join(fs.pdfDir, filename)

	return fs.saveFile(file, filePath)
}

// SaveCover saves a cover image and returns its URL
func (fs *FileStorage) SaveCover(file multipart.File, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if !strings.EqualFold(ext, ".jpg") && !strings.EqualFold(ext, ".jpeg") && !strings.EqualFold(ext, ".png") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFilename := uuid.New().String() + ext
	filePath := filepath.Join(fs.coverDir, newFilename)

	_, err := fs.saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	// Return URL-friendly path
	return "/uploads/covers/" + newFilename, nil
}

// SaveAudio saves an audio file and returns its URL
func (fs *FileStorage) SaveAudio(file multipart.File, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if !strings.EqualFold(ext, ".mp3") && !strings.EqualFold(ext, ".wav") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFilename := uuid.New().String() + ext
	filePath := filepath.Join(fs.audioDir, newFilename)

	_, err := fs.saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	// Return URL-friendly path
	return "/uploads/audio/" + newFilename, nil
}

// saveFile is a helper function to save a file
func (fs *FileStorage) saveFile(file multipart.File, filePath string) (string, error) {
	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(filePath) // Clean up on error
		return "", fmt.Errorf("error copying file: %v", err)
	}

	return filePath, nil
}

// DeleteFile deletes a file at the given path
func (fs *FileStorage) DeleteFile(filePath string) error {
	// Ensure the file is within our base directory
	if !strings.HasPrefix(filePath, fs.baseDir) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	return os.Remove(filePath)
}

func (fs *FileStorage) GetFilePath(relativePath string) string {
	// Remove leading slash if present
	if len(relativePath) > 0 && relativePath[0] == '/' {
		relativePath = relativePath[1:]
	}

	return filepath.Join(fs.baseDir, relativePath)
}
```

## File: /Users/saint/Desktop/pdf-player/backend/utils.go

```go
package main

import (
	"time"

	"github.com/google/uuid"
)

// Helper function to format datetime for SQLite
func formatDateTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

// generateID generates a unique ID using UUID
func generateID() string {
	return uuid.New().String()
}
```

## File: /Users/saint/Desktop/pdf-player/backend/main.go

```go
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

// BookBasic represents the basic book information used in the initial implementation
type BookBasic struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	CoverURL string `json:"coverUrl"`
	Content  string `json:"content"`
}

var db *sql.DB
var fileStorage *FileStorage

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize database
	var err error
	db, err = sql.Open("sqlite3", "books.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create tables
	if err := InitDB(); err != nil {
		log.Fatal(err)
	}

	// Initialize file storage
	fileStorage, err = NewFileStorage("uploads")
	if err != nil {
		log.Fatal(err)
	}

	// Create router
	router := mux.NewRouter()

	// Configure CORS
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			allowedOrigins := strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",")
			origin := r.Header.Get("Origin")
			for _, allowed := range allowedOrigins {
				if origin == allowed {
					w.Header().Set("Access-Control-Allow-Origin", origin)
					break
				}
			}
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	// File upload routes
	router.HandleFunc("/api/upload/pdf", uploadPDFHandler).Methods("POST")
	router.HandleFunc("/api/upload/cover", uploadCoverHandler).Methods("POST")

	// Book routes
	router.HandleFunc("/api/books/{id}", getBookHandler).Methods("GET")
	router.HandleFunc("/api/books", getBooksHandler).Methods("GET")

	// Reading progress routes
	router.HandleFunc("/api/progress", updateProgressHandler).Methods("POST")
	router.HandleFunc("/api/progress/{bookId}", getProgressHandler).Methods("GET")

	// Bookmark routes
	router.HandleFunc("/api/bookmarks", createBookmarkHandler).Methods("POST")
	router.HandleFunc("/api/bookmarks/{bookId}", getBookmarksHandler).Methods("GET")
	router.HandleFunc("/api/bookmarks/{id}", updateBookmarkHandler).Methods("PUT")
	router.HandleFunc("/api/bookmarks/{id}", deleteBookmarkHandler).Methods("DELETE")

	// Audio segment routes
	router.HandleFunc("/api/audio/generate", generateAudioHandler).Methods("POST")
	router.HandleFunc("/api/audio/{bookId}", getAudioSegmentsHandler).Methods("GET")

	// Category and tag routes
	router.HandleFunc("/api/categories", getCategoriesHandler).Methods("GET")
	router.HandleFunc("/api/tags", getTagsHandler).Methods("GET")

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
	// Parse multipart form with 32MB limit
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Process PDF and save book
	book, err := processPDF(file, header.Filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error processing PDF: %v", err), http.StatusInternalServerError)
		return
	}

	// Save book to database
	if err := SaveBook(db, book); err != nil {
		http.Error(w, "Error saving book", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func uploadCoverHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	coverURL, err := fileStorage.SaveCover(file, header.Filename)
	if err != nil {
		http.Error(w, "Error saving cover", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"coverUrl": coverURL})
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, cover_url, page_count, language, created_at
		FROM books
		ORDER BY created_at DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.CoverURL,
			&book.PageCount,
			&book.Language,
			&book.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func updateProgressHandler(w http.ResponseWriter, r *http.Request) {
	var progress ReadingProgress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if progress.ID == "" {
		progress.ID = uuid.New().String()
	}

	if err := UpdateReadingProgress(db, &progress); err != nil {
		http.Error(w, "Error updating progress", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	progress, err := GetReadingProgress(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving progress", http.StatusInternalServerError)
		return
	}

	if progress == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func createBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if bookmark.ID == "" {
		bookmark.ID = uuid.New().String()
	}

	if err := CreateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error creating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	bookmarks, err := GetBookmarks(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving bookmarks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmarks)
}

func updateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookmark.ID = vars["id"]
	if err := UpdateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error updating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := DeleteBookmark(db, vars["id"]); err != nil {
		http.Error(w, "Error deleting bookmark", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateAudioHandler(w http.ResponseWriter, r *http.Request) {
	var segment AudioSegment
	if err := json.NewDecoder(r.Body).Decode(&segment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if segment.ID == "" {
		segment.ID = uuid.New().String()
	}

	// Set initial status
	segment.Status = "pending"
	segment.CreatedAt = time.Now()

	if err := CreateAudioSegment(db, &segment); err != nil {
		http.Error(w, "Error creating audio segment", http.StatusInternalServerError)
		return
	}

	// Start TTS processing in background
	go processAudioSegment(&segment)

	json.NewEncoder(w).Encode(segment)
}

func getAudioSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	query := `
		SELECT id, book_id, segment_number, content, audio_url, duration, status, created_at
		FROM audio_segments
		WHERE book_id = ?
		ORDER BY segment_number ASC
	`

	rows, err := db.Query(query, vars["bookId"])
	if err != nil {
		http.Error(w, "Error retrieving audio segments", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var segments []AudioSegment
	for rows.Next() {
		var segment AudioSegment
		err := rows.Scan(
			&segment.ID,
			&segment.BookID,
			&segment.SegmentNumber,
			&segment.Content,
			&segment.AudioURL,
			&segment.Duration,
			&segment.Status,
			&segment.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning audio segments", http.StatusInternalServerError)
			return
		}
		segments = append(segments, segment)
	}

	json.NewEncoder(w).Encode(segments)
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, description, created_at FROM categories ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning categories", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func getTagsHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, created_at FROM tags ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving tags", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning tags", http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}
```



---

> 📸 Generated with [Jockey CLI](https://github.com/saint0x/jockey-cli)
````

## File: /Users/saint/Desktop/pdf-player/backend/db.go

```go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"
)

// InitDB initializes the database connection and creates tables
func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./ereader.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	// Read and execute schema.sql
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("error reading schema file: %v", err)
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return fmt.Errorf("error creating schema: %v", err)
	}

	// Create audio_segments table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS audio_segments (
			id TEXT PRIMARY KEY,
			book_id TEXT NOT NULL,
			segment_number INTEGER NOT NULL,
			content TEXT NOT NULL,
			audio_url TEXT,
			duration REAL,
			status TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL,
			FOREIGN KEY (book_id) REFERENCES books(id)
		)
	`)
	if err != nil {
		return fmt.Errorf("error creating audio_segments table: %v", err)
	}

	return nil
}

// GetBookByID retrieves a book by its ID
func GetBookByID(db *sql.DB, id string) (*Book, error) {
	var book Book
	query := `
		SELECT id, title, author, cover_url, file_url, 
			page_count, current_page, language, status,
			created_at, updated_at
		FROM books WHERE id = ?
	`

	err := db.QueryRow(query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.CoverURL,
		&book.FileURL,
		&book.PageCount,
		&book.CurrentPage,
		&book.Language,
		&book.Status,
		&book.CreatedAt,
		&book.UpdatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("book not found")
	}
	if err != nil {
		return nil, err
	}

	// Get categories
	rows, err := db.Query(`
		SELECT c.name FROM categories c 
		JOIN book_categories bc ON c.id = bc.category_id 
		WHERE bc.book_id = ?`, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var category string
			rows.Scan(&category)
			book.Categories = append(book.Categories, category)
		}
	}

	// Get tags
	rows, err = db.Query(`
		SELECT t.name FROM tags t 
		JOIN book_tags bt ON t.id = bt.tag_id 
		WHERE bt.book_id = ?`, id)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var tag string
			rows.Scan(&tag)
			book.Tags = append(book.Tags, tag)
		}
	}

	return &book, nil
}

// SaveBook saves a new book to the database
func SaveBook(db *sql.DB, book *Book) error {
	query := `
		INSERT INTO books (
			id, title, author, cover_url, file_url,
			page_count, current_page, language, status,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		book.ID,
		book.Title,
		book.Author,
		book.CoverURL,
		book.FileURL,
		book.PageCount,
		book.CurrentPage,
		book.Language,
		book.Status,
		formatDateTime(book.CreatedAt),
		formatDateTime(book.UpdatedAt),
	)

	return err
}

// UpdateReadingProgress updates or creates a reading progress record
func UpdateReadingProgress(db *sql.DB, progress *ReadingProgress) error {
	query := `
		INSERT INTO reading_progress (id, book_id, user_id, current_page, total_pages, completion_percent, last_read_at)
		VALUES (?, ?, ?, ?, ?, ?, CURRENT_TIMESTAMP)
		ON CONFLICT(book_id, user_id) DO UPDATE SET
			current_page = excluded.current_page,
			total_pages = excluded.total_pages,
			completion_percent = excluded.completion_percent,
			last_read_at = CURRENT_TIMESTAMP
	`

	_, err := db.Exec(query,
		progress.ID,
		progress.BookID,
		progress.UserID,
		progress.CurrentPage,
		progress.TotalPages,
		progress.CompletionPercent,
	)

	return err
}

// GetReadingProgress retrieves reading progress for a book and user
func GetReadingProgress(db *sql.DB, bookID, userID string) (*ReadingProgress, error) {
	var progress ReadingProgress
	query := `
		SELECT id, book_id, user_id, current_page, total_pages, completion_percent, last_read_at
		FROM reading_progress
		WHERE book_id = ? AND user_id = ?
	`

	err := db.QueryRow(query, bookID, userID).Scan(
		&progress.ID,
		&progress.BookID,
		&progress.UserID,
		&progress.CurrentPage,
		&progress.TotalPages,
		&progress.CompletionPercent,
		&progress.LastReadAt,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &progress, nil
}

// CreateBookmark creates a new bookmark
func CreateBookmark(db *sql.DB, bookmark *Bookmark) error {
	query := `
		INSERT INTO bookmarks (id, book_id, user_id, page_number, note, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
	`

	_, err := db.Exec(query,
		bookmark.ID,
		bookmark.BookID,
		bookmark.UserID,
		bookmark.PageNumber,
		bookmark.Note,
	)

	return err
}

// GetBookmarks retrieves all bookmarks for a book and user
func GetBookmarks(db *sql.DB, bookID, userID string) ([]Bookmark, error) {
	query := `
		SELECT id, book_id, user_id, page_number, note, created_at, updated_at
		FROM bookmarks
		WHERE book_id = ? AND user_id = ?
		ORDER BY page_number ASC
	`

	rows, err := db.Query(query, bookID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookmarks []Bookmark
	for rows.Next() {
		var bookmark Bookmark
		err := rows.Scan(
			&bookmark.ID,
			&bookmark.BookID,
			&bookmark.UserID,
			&bookmark.PageNumber,
			&bookmark.Note,
			&bookmark.CreatedAt,
			&bookmark.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		bookmarks = append(bookmarks, bookmark)
	}

	return bookmarks, nil
}

// DeleteBookmark deletes a bookmark
func DeleteBookmark(db *sql.DB, id string) error {
	query := "DELETE FROM bookmarks WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}

// UpdateBookmark updates a bookmark's note
func UpdateBookmark(db *sql.DB, bookmark *Bookmark) error {
	query := `
		UPDATE bookmarks 
		SET note = ?, updated_at = CURRENT_TIMESTAMP
		WHERE id = ? AND book_id = ? AND user_id = ?
	`

	result, err := db.Exec(query, bookmark.Note, bookmark.ID, bookmark.BookID, bookmark.UserID)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return fmt.Errorf("bookmark not found")
	}

	return nil
}

// GetAudioSegments retrieves all audio segments for a book
func GetAudioSegments(bookID string) ([]AudioSegment, error) {
	query := `
		SELECT id, book_id, content, audio_url, status, created_at, updated_at
		FROM audio_segments
		WHERE book_id = ?
		ORDER BY created_at ASC
	`

	rows, err := db.Query(query, bookID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var segments []AudioSegment
	for rows.Next() {
		var segment AudioSegment
		err := rows.Scan(
			&segment.ID,
			&segment.BookID,
			&segment.Content,
			&segment.AudioURL,
			&segment.Status,
			&segment.CreatedAt,
			&segment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

// CreateAudioSegment creates a new audio segment
func CreateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, content, audio_url, status,
			created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		formatDateTime(segment.CreatedAt),
		formatDateTime(segment.UpdatedAt),
	)

	return err
}

// UpdateAudioSegment updates an existing audio segment
func UpdateAudioSegment(db *sql.DB, segment *AudioSegment) error {
	query := `
		UPDATE audio_segments 
		SET content = ?, audio_url = ?, status = ?, updated_at = ?
		WHERE id = ? AND book_id = ?
	`

	_, err := db.Exec(query,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		formatDateTime(segment.UpdatedAt),
		segment.ID,
		segment.BookID,
	)

	return err
}

// Category operations
func CreateCategory(db *sql.DB, category *Category) error {
	category.ID = generateID()
	category.CreatedAt = time.Now()

	_, err := db.Exec(`
		INSERT INTO categories (id, name, description, created_at)
		VALUES (?, ?, ?, ?)`,
		category.ID, category.Name, category.Description, category.CreatedAt)

	return err
}

func AddBookToCategory(db *sql.DB, bookID, categoryID string) error {
	_, err := db.Exec(`
		INSERT INTO book_categories (book_id, category_id)
		VALUES (?, ?)`, bookID, categoryID)

	return err
}

// Tag operations
func CreateTag(db *sql.DB, tag *Tag) error {
	tag.ID = generateID()
	tag.CreatedAt = time.Now()

	_, err := db.Exec(`
		INSERT INTO tags (id, name, created_at)
		VALUES (?, ?, ?)`,
		tag.ID, tag.Name, tag.CreatedAt)

	return err
}

func AddBookTag(db *sql.DB, bookID, tagID string) error {
	_, err := db.Exec(`
		INSERT INTO book_tags (book_id, tag_id)
		VALUES (?, ?)`, bookID, tagID)

	return err
}

// CleanupDuplicateBooks removes duplicate books keeping only the latest version
func CleanupDuplicateBooks(db *sql.DB) error {
	// First, get all books grouped by title, keeping only the latest version
	query := `
		WITH RankedBooks AS (
			SELECT *,
				ROW_NUMBER() OVER (PARTITION BY title ORDER BY created_at DESC) as rn
			FROM books
		)
		DELETE FROM books 
		WHERE id IN (
			SELECT id 
			FROM RankedBooks 
			WHERE rn > 1
		)
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error cleaning up duplicate books: %v", err)
	}

	// Also clean up orphaned audio segments
	_, err = db.Exec(`
		DELETE FROM audio_segments 
		WHERE book_id NOT IN (SELECT id FROM books)
	`)
	if err != nil {
		return fmt.Errorf("error cleaning up orphaned audio segments: %v", err)
	}

	return nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/schema.sql

```sql
-- Books table with enhanced metadata
CREATE TABLE IF NOT EXISTS books (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    author TEXT,
    cover_url TEXT,
    file_url TEXT NOT NULL,
    page_count INTEGER DEFAULT 0,
    current_page INTEGER DEFAULT 0,
    language TEXT DEFAULT 'en',
    status TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Categories for organizing books
CREATE TABLE IF NOT EXISTS categories (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book-Category relationship (many-to-many)
CREATE TABLE IF NOT EXISTS book_categories (
    book_id TEXT,
    category_id TEXT,
    PRIMARY KEY (book_id, category_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

-- Tags for flexible book organization
CREATE TABLE IF NOT EXISTS tags (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Book-Tag relationship (many-to-many)
CREATE TABLE IF NOT EXISTS book_tags (
    book_id TEXT,
    tag_id TEXT,
    PRIMARY KEY (book_id, tag_id),
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

-- Audio synthesis tracking
CREATE TABLE IF NOT EXISTS audio_segments (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    content TEXT NOT NULL,
    audio_url TEXT,
    status TEXT NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Reading progress tracking
CREATE TABLE IF NOT EXISTS reading_progress (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    current_page INTEGER DEFAULT 0,
    completion_percentage FLOAT DEFAULT 0,
    last_read_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE,
    UNIQUE(book_id)
);

-- Bookmarks
CREATE TABLE IF NOT EXISTS bookmarks (
    id TEXT PRIMARY KEY,
    book_id TEXT NOT NULL,
    page_number INTEGER NOT NULL,
    note TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (book_id) REFERENCES books(id) ON DELETE CASCADE
);

-- Create indexes for better query performance
CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
CREATE INDEX IF NOT EXISTS idx_reading_progress_book ON reading_progress(book_id);
CREATE INDEX IF NOT EXISTS idx_audio_segments_book_id ON audio_segments(book_id);
CREATE INDEX IF NOT EXISTS idx_bookmarks_book ON bookmarks(book_id); 
```

## File: /Users/saint/Desktop/pdf-player/backend/pdf.go

```go
package main

import (
	"database/sql"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	"github.com/ledongthuc/pdf"
)

// processPDF processes a PDF file and returns a Book object
func processPDF(file multipart.File, filename string) (*Book, error) {
	// Create temporary file
	tmpFile, err := os.CreateTemp("", "book-*.pdf")
	if err != nil {
		return nil, fmt.Errorf("error creating temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy uploaded file to temp file
	if _, err := io.Copy(tmpFile, file); err != nil {
		return nil, fmt.Errorf("error copying file: %v", err)
	}

	// Open PDF file for reading
	pdfFile, reader, err := pdf.Open(tmpFile.Name())
	if err != nil {
		return nil, fmt.Errorf("error opening PDF: %v", err)
	}
	defer pdfFile.Close()

	// Get total number of pages
	totalPages := reader.NumPage()

	// Create book object
	book := &Book{
		ID:          uuid.New().String(),
		Title:       strings.TrimSuffix(filename, ".pdf"),
		Author:      "Unknown", // Can be updated later
		FileURL:     "",        // Will be set after uploading to UploadThing
		PageCount:   totalPages,
		CurrentPage: 0,
		Language:    "en", // Default to English, can be updated later
		Status:      "pending",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return book, nil
}

func processBook(db *sql.DB, book *Book) error {
	// Download PDF from UploadThing
	req, err := http.NewRequest("GET", book.FileURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	// Add UploadThing authentication
	req.Header.Set("Authorization", "Bearer "+AppConfig.UploadThingToken)

	// Make the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download PDF: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download PDF: status code %d", resp.StatusCode)
	}

	// Create temporary file
	tmpFile, err := os.CreateTemp("", "book-*.pdf")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	// Copy PDF to temp file
	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("failed to copy PDF to temp file: %v", err)
	}

	// Open PDF file
	pdfFile, reader, err := pdf.Open(tmpFile.Name())
	if err != nil {
		return fmt.Errorf("failed to open PDF: %v", err)
	}
	defer pdfFile.Close()

	// Update book page count
	totalPages := reader.NumPage()
	book.PageCount = totalPages
	book.UpdatedAt = time.Now()
	if err := UpdateBook(db, book); err != nil {
		return fmt.Errorf("failed to update book: %v", err)
	}

	// Process each page
	for pageNum := 1; pageNum <= totalPages; pageNum++ {
		page := reader.Page(pageNum)
		if page.V.IsNull() {
			continue
		}

		text, err := page.GetPlainText(nil)
		if err != nil {
			return fmt.Errorf("failed to get text from page %d: %v", pageNum, err)
		}

		// Create audio segment
		segment := &AudioSegment{
			ID:        uuid.New().String(),
			BookID:    book.ID,
			Content:   text,
			Status:    "pending",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		if err := SaveAudioSegment(db, segment); err != nil {
			return fmt.Errorf("failed to save audio segment: %v", err)
		}

		// Start TTS processing in background
		go processAudioSegment(db, segment)
	}

	return nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/tts.go

```go
package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// TTSRequest represents a request to the Replicate API
type TTSRequest struct {
	Version string   `json:"version"`
	Input   TTSInput `json:"input"`
}

// TTSInput represents the input parameters for the TTS model
type TTSInput struct {
	Text     string `json:"text"`
	Language string `json:"language"`
}

// TTSResponse represents a response from the Replicate API
type TTSResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Output string `json:"output"`
	Error  string `json:"error"`
}

// processAudioSegment processes a text segment and generates audio
func processAudioSegment(db *sql.DB, segment *AudioSegment) error {
	// Get Kokoro model version from environment
	modelVersion := os.Getenv("KOKORO_MODEL_VERSION")
	if modelVersion == "" {
		return fmt.Errorf("KOKORO_MODEL_VERSION not set")
	}

	// Prepare request
	reqBody := TTSRequest{
		Version: modelVersion,
		Input: TTSInput{
			Text:     segment.Content,
			Language: "en", // Default to English for now
		},
	}

	// Convert request to JSON
	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return fmt.Errorf("error marshaling request: %v", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", os.Getenv("REPLICATE_API_URL")+"/predictions", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	var ttsResp TTSResponse
	if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	// Check for error
	if ttsResp.Error != "" {
		return fmt.Errorf("TTS error: %s", ttsResp.Error)
	}

	// Poll for completion
	for ttsResp.Status != "succeeded" {
		time.Sleep(2 * time.Second)

		// Get prediction status
		req, err = http.NewRequest("GET", os.Getenv("REPLICATE_API_URL")+"/predictions/"+ttsResp.ID, nil)
		if err != nil {
			return fmt.Errorf("error creating status request: %v", err)
		}

		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return fmt.Errorf("error checking status: %v", err)
		}

		if err := json.NewDecoder(resp.Body).Decode(&ttsResp); err != nil {
			resp.Body.Close()
			return fmt.Errorf("error decoding status: %v", err)
		}
		resp.Body.Close()

		if ttsResp.Status == "failed" {
			return fmt.Errorf("TTS processing failed: %s", ttsResp.Error)
		}
	}

	// Download audio file from Replicate
	resp, err = http.Get(ttsResp.Output)
	if err != nil {
		return fmt.Errorf("error downloading audio: %v", err)
	}
	defer resp.Body.Close()

	// Create multipart form data for UploadThing
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.mp3", segment.ID))
	if err != nil {
		return fmt.Errorf("error creating form file: %v", err)
	}

	// Copy audio data to form file
	if _, err := io.Copy(part, resp.Body); err != nil {
		return fmt.Errorf("error copying audio data: %v", err)
	}

	// Close multipart writer
	writer.Close()

	// Create request to UploadThing
	uploadReq, err := http.NewRequest("POST", os.Getenv("UPLOADTHING_URL")+"/upload", body)
	if err != nil {
		return fmt.Errorf("error creating upload request: %v", err)
	}

	// Set headers for UploadThing
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("Authorization", "Bearer "+os.Getenv("UPLOADTHING_TOKEN"))

	// Send upload request
	uploadResp, err := client.Do(uploadReq)
	if err != nil {
		return fmt.Errorf("error uploading to UploadThing: %v", err)
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode != http.StatusOK {
		return fmt.Errorf("UploadThing error: %s", uploadResp.Status)
	}

	// Parse UploadThing response
	var uploadResult struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(uploadResp.Body).Decode(&uploadResult); err != nil {
		return fmt.Errorf("error decoding UploadThing response: %v", err)
	}

	// Update segment with UploadThing URL and status
	segment.AudioURL = uploadResult.URL
	segment.Status = "completed"

	// Update database
	if err := UpdateAudioSegment(db, segment); err != nil {
		return fmt.Errorf("error updating segment: %v", err)
	}

	return nil
}

// generateTTS generates audio for the given text using Kokoro TTS
func generateTTS(text string) (string, error) {
	// Prepare request to Replicate API
	reqBody := map[string]interface{}{
		"version": os.Getenv("KOKORO_MODEL_VERSION"),
		"input": map[string]interface{}{
			"text":       text,
			"language":   "en",
			"speaker_id": 0,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %v", err)
	}

	// Create request
	req, err := http.NewRequest("POST", os.Getenv("REPLICATE_API_URL")+"/predictions", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code %d: %s", resp.StatusCode, body)
	}

	// Parse response
	var prediction struct {
		ID     string `json:"id"`
		Status string `json:"status"`
		Output string `json:"output"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	// Poll for completion
	for prediction.Status != "succeeded" {
		time.Sleep(1 * time.Second)

		req, err = http.NewRequest("GET", os.Getenv("REPLICATE_API_URL")+"/predictions/"+prediction.ID, nil)
		if err != nil {
			return "", fmt.Errorf("error creating poll request: %v", err)
		}
		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return "", fmt.Errorf("error polling: %v", err)
		}
		defer resp.Body.Close()

		if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
			return "", fmt.Errorf("error decoding poll response: %v", err)
		}

		if prediction.Status == "failed" {
			return "", fmt.Errorf("TTS generation failed")
		}
	}

	// Download the audio file from Replicate
	audioResp, err := http.Get(prediction.Output)
	if err != nil {
		return "", fmt.Errorf("error downloading audio: %v", err)
	}
	defer audioResp.Body.Close()

	// Create multipart form data for UploadThing
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Create form file
	part, err := writer.CreateFormFile("file", fmt.Sprintf("%s.mp3", generateID()))
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}

	// Copy audio data to form file
	if _, err := io.Copy(part, audioResp.Body); err != nil {
		return "", fmt.Errorf("error copying audio data: %v", err)
	}

	// Close multipart writer
	writer.Close()

	// Create request to UploadThing
	uploadReq, err := http.NewRequest("POST", os.Getenv("UPLOADTHING_URL")+"/upload", body)
	if err != nil {
		return "", fmt.Errorf("error creating upload request: %v", err)
	}

	// Set headers for UploadThing
	uploadReq.Header.Set("Content-Type", writer.FormDataContentType())
	uploadReq.Header.Set("Authorization", "Bearer "+os.Getenv("UPLOADTHING_TOKEN"))

	// Send upload request
	uploadResp, err := client.Do(uploadReq)
	if err != nil {
		return "", fmt.Errorf("error uploading to UploadThing: %v", err)
	}
	defer uploadResp.Body.Close()

	if uploadResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("UploadThing error: %s", uploadResp.Status)
	}

	// Parse UploadThing response
	var uploadResult struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(uploadResp.Body).Decode(&uploadResult); err != nil {
		return "", fmt.Errorf("error decoding UploadThing response: %v", err)
	}

	return uploadResult.URL, nil
}
```

## File: /Users/saint/Desktop/pdf-player/backend/go.sum

```sum
github.com/google/uuid v1.6.0 h1:NIvaJDMOsjHA8n1jAhLSgzrAzy1Hgr+hNrb57e+94F0=
github.com/google/uuid v1.6.0/go.mod h1:TIyPZe4MgqvfeYDBFedMoGGpEw/LqOeaOT+nhxU+yHo=
github.com/gorilla/mux v1.8.1 h1:TuBL49tXwgrFYWhqrNgrUNEY92u81SPhu7sTdzQEiWY=
github.com/gorilla/mux v1.8.1/go.mod h1:AKf9I4AEqPTmMytcMc0KkNouC66V3BtZ4qD5fmWSiMQ=
github.com/gorilla/websocket v1.5.3 h1:saDtZ6Pbx/0u+bgYQ3q96pZgCzfhKXGPqt7kZ72aNNg=
github.com/gorilla/websocket v1.5.3/go.mod h1:YR8l580nyteQvAITg2hZ9XVh4b55+EU/adAjf1fMHhE=
github.com/joho/godotenv v1.5.1 h1:7eLL/+HRGLY0ldzfGMeQkb7vMd0as4CfYvUVzLqw0N0=
github.com/joho/godotenv v1.5.1/go.mod h1:f4LDr5Voq0i2e/R5DDNOoa2zzDfwtkZa6DnEwAbqwq4=
github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06 h1:kacRlPN7EN++tVpGUorNGPn/4DnB7/DfTY82AOn6ccU=
github.com/ledongthuc/pdf v0.0.0-20240201131950-da5b75280b06/go.mod h1:imJHygn/1yfhB7XSJJKlFZKl/J+dCPAknuiaGOshXAs=
github.com/mattn/go-sqlite3 v1.14.24 h1:tpSp2G2KyMnnQu99ngJ47EIkWVmliIizyZBfPrBWDRM=
github.com/mattn/go-sqlite3 v1.14.24/go.mod h1:Uh1q+B4BYcTPb+yiD3kU8Ct7aC0hY9fxUwlHK0RXw+Y=
```

## File: /Users/saint/Desktop/pdf-player/backend/README.md

````md
# PDF Player Backend

This is the backend service for the PDF Player application. It provides APIs for uploading PDFs, extracting text, managing books, and tracking reading progress.

## Prerequisites

- Go 1.16 or later
- SQLite3

## Setup

1. Install dependencies:
```bash
go mod download
```

2. Run the server:
```bash
go run *.go
```

The server will start on port 8080 by default. You can change the port by setting the `PORT` environment variable.

## API Endpoints

### Books
- **POST** `/api/upload` - Upload a PDF file
  - Content-Type: `multipart/form-data`
  - Form field: `file` (PDF file)
  - Returns: Book object with extracted text

- **GET** `/api/books` - Get all books
  - Returns: Array of book objects

- **GET** `/api/book/{id}` - Get a specific book
  - Returns: Single book object

### Reading Progress
- **PUT** `/api/book/{id}/progress` - Update reading progress
  - Body: `{ "currentPage": number, "completion": number }`

- **GET** `/api/book/{id}/progress` - Get reading progress
  - Returns: Reading progress object

### Bookmarks
- **POST** `/api/book/{id}/bookmarks` - Create a bookmark
  - Body: `{ "pageNumber": number, "note": string }`
  - Returns: Created bookmark object

- **GET** `/api/book/{id}/bookmarks` - Get all bookmarks for a book
  - Returns: Array of bookmark objects

### Audio Segments
- **POST** `/api/book/{id}/audio` - Create an audio segment
  - Body: `{ "startPage": number, "endPage": number }`
  - Returns: Created audio segment object

- **PUT** `/api/book/{id}/audio/{segmentId}` - Update audio segment status
  - Body: `{ "status": string, "audioUrl": string, "duration": number }`

### Categories
- **POST** `/api/categories` - Create a category
  - Body: `{ "name": string, "description": string }`
  - Returns: Created category object

- **PUT** `/api/book/{id}/categories/{categoryId}` - Add book to category

### Tags
- **POST** `/api/tags` - Create a tag
  - Body: `{ "name": string }`
  - Returns: Created tag object

- **PUT** `/api/book/{id}/tags/{tagId}` - Add tag to book

## Data Models

### Book
```json
{
  "id": "string",
  "title": "string",
  "author": "string",
  "coverUrl": "string",
  "content": "string",
  "filePath": "string",
  "pageCount": number,
  "currentPage": number,
  "language": "string",
  "createdAt": "datetime",
  "updatedAt": "datetime",
  "categories": ["string"],
  "tags": ["string"]
}
```

### Reading Progress
```json
{
  "id": "string",
  "bookId": "string",
  "currentPage": number,
  "completionPercent": number,
  "lastReadAt": "datetime"
}
```

### Bookmark
```json
{
  "id": "string",
  "bookId": "string",
  "pageNumber": number,
  "note": "string",
  "createdAt": "datetime"
}
```

### Audio Segment
```json
{
  "id": "string",
  "bookId": "string",
  "segmentNumber": number,
  "content": "string",
  "audioUrl": "string",
  "duration": number,
  "status": "string",
  "createdAt": "datetime"
}
```

## Future Enhancements

1. PDF metadata extraction for better book information
2. Custom book cover image upload
3. Text-to-speech synthesis integration using Kokoro TTS
4. Book categories and tags for better organization
5. Full-text search functionality 
````

## File: /Users/saint/Desktop/pdf-player/backend/storage.go

```go
package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

// FileStorage handles file operations
type FileStorage struct {
	baseDir  string
	pdfDir   string
	coverDir string
	audioDir string
}

// NewFileStorage creates a new FileStorage instance
func NewFileStorage(baseDir string) (*FileStorage, error) {
	fs := &FileStorage{
		baseDir:  baseDir,
		pdfDir:   filepath.Join(baseDir, "pdfs"),
		coverDir: filepath.Join(baseDir, "covers"),
		audioDir: filepath.Join(baseDir, "audio"),
	}

	// Create directories if they don't exist
	for _, dir := range []string{fs.pdfDir, fs.coverDir, fs.audioDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, fmt.Errorf("error creating directory %s: %v", dir, err)
		}
	}

	return fs, nil
}

// SavePDF saves a PDF file and returns its path
func (fs *FileStorage) SavePDF(file multipart.File, header *multipart.FileHeader) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	if !strings.EqualFold(ext, ".pdf") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	filename := uuid.New().String() + ext
	filePath := filepath.Join(fs.pdfDir, filename)

	return fs.saveFile(file, filePath)
}

// SaveCover saves a cover image and returns its URL
func (fs *FileStorage) SaveCover(file multipart.File, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if !strings.EqualFold(ext, ".jpg") && !strings.EqualFold(ext, ".jpeg") && !strings.EqualFold(ext, ".png") {
		return "", fmt.Errorf("invalid file type: %s", ext)
	}

	newFilename := uuid.New().String() + ext
	filePath := filepath.Join(fs.coverDir, newFilename)

	_, err := fs.saveFile(file, filePath)
	if err != nil {
		return "", err
	}

	// Return URL-friendly path
	return "/uploads/covers/" + newFilename, nil
}

// SaveAudio saves an audio file and returns its URL
func (fs *FileStorage) SaveAudio(file io.Reader, filename string) (string, error) {
	// Generate unique filename
	ext := filepath.Ext(filename)
	if ext == "" {
		ext = ".mp3"
	}
	filename = uuid.New().String() + ext

	// Create audio directory if it doesn't exist
	audioDir := filepath.Join(fs.baseDir, "audio")
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		return "", fmt.Errorf("error creating audio directory: %v", err)
	}

	// Create file
	filePath := filepath.Join(audioDir, filename)
	f, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer f.Close()

	// Copy file contents
	if _, err := io.Copy(f, file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}

	// Return URL path relative to server root
	return fmt.Sprintf("/audio/%s", filename), nil
}

// saveFile is a helper function to save a file
func (fs *FileStorage) saveFile(file multipart.File, filePath string) (string, error) {
	// Create destination file
	dst, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("error creating file: %v", err)
	}
	defer dst.Close()

	// Copy file contents
	if _, err := io.Copy(dst, file); err != nil {
		os.Remove(filePath) // Clean up on error
		return "", fmt.Errorf("error copying file: %v", err)
	}

	return filePath, nil
}

// DeleteFile deletes a file at the given path
func (fs *FileStorage) DeleteFile(filePath string) error {
	// Ensure the file is within our base directory
	if !strings.HasPrefix(filePath, fs.baseDir) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	return os.Remove(filePath)
}

func (fs *FileStorage) GetFilePath(relativePath string) string {
	// Remove leading slash if present
	if len(relativePath) > 0 && relativePath[0] == '/' {
		relativePath = relativePath[1:]
	}

	return filepath.Join(fs.baseDir, relativePath)
}
```

## File: /Users/saint/Desktop/pdf-player/backend/utils.go

```go
package main

import (
	"time"

	"github.com/google/uuid"
)

// Helper function to format datetime for SQLite
func formatDateTime(t time.Time) string {
	return t.UTC().Format("2006-01-02 15:04:05")
}

// generateID generates a unique ID using UUID
func generateID() string {
	return uuid.New().String()
}
```

## File: /Users/saint/Desktop/pdf-player/backend/main.go

```go
package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var fileStorage *FileStorage

// Add WebSocket upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for now
	},
}

// Add WebSocket connections map
var wsConnections = make(map[string][]*websocket.Conn)

func main() {
	// Load configuration
	if err := LoadConfig(); err != nil {
		log.Fatal("Error loading configuration:", err)
	}

	// Initialize database
	var err error
	db, err = sql.Open("sqlite3", AppConfig.DBPath)
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	// Create tables
	if err := InitDB(); err != nil {
		log.Fatal("Error initializing database:", err)
	}

	// Clean up any duplicate books
	if err := CleanupDuplicateBooks(db); err != nil {
		log.Printf("Warning: Error cleaning up duplicates: %v", err)
	}

	// Initialize file storage
	fileStorage, err = NewFileStorage(AppConfig.UploadDir)
	if err != nil {
		log.Fatal("Error initializing file storage:", err)
	}

	// Create audio directory if it doesn't exist
	audioDir := filepath.Join(AppConfig.UploadDir, "audio")
	if err := os.MkdirAll(audioDir, 0755); err != nil {
		log.Fatal("Error creating audio directory:", err)
	}

	// Initialize router
	router := mux.NewRouter()

	// Add CORS middleware
	router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set CORS headers for all responses
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-User-ID")
			w.Header().Set("Access-Control-Max-Age", "3600")

			// Handle preflight requests
			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	})

	// Serve static audio files
	audioDir = filepath.Join(AppConfig.UploadDir, "audio")
	router.PathPrefix("/audio/").Handler(http.StripPrefix("/audio/", http.FileServer(http.Dir(audioDir))))

	// File upload routes
	router.HandleFunc("/api/upload", uploadPDFHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/upload/pdf", uploadPDFHandler).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/upload/cover", uploadCoverHandler).Methods("POST", "OPTIONS")

	// Book routes
	router.HandleFunc("/api/books/{id}", getBookHandler).Methods("GET")
	router.HandleFunc("/api/books", getBooksHandler).Methods("GET")
	router.HandleFunc("/api/books/{id}/status", getBookStatusHandler).Methods("GET")
	router.HandleFunc("/api/books/{id}/update-url", updateBookURLHandler).Methods("POST")
	router.HandleFunc("/api/books/{id}/process", processBookHandler).Methods("POST")

	// Reading progress routes
	router.HandleFunc("/api/progress", updateProgressHandler).Methods("POST")
	router.HandleFunc("/api/progress/{bookId}", getProgressHandler).Methods("GET")

	// Bookmark routes
	router.HandleFunc("/api/bookmarks", createBookmarkHandler).Methods("POST")
	router.HandleFunc("/api/bookmarks/{bookId}", getBookmarksHandler).Methods("GET")
	router.HandleFunc("/api/bookmarks/{id}", updateBookmarkHandler).Methods("PUT")
	router.HandleFunc("/api/bookmarks/{id}", deleteBookmarkHandler).Methods("DELETE")

	// Audio segment routes
	router.HandleFunc("/api/books/{id}/audio-segments", getAudioSegmentsHandler).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/books/{id}/generate-audio", generateBookAudioHandler).Methods("POST")
	router.HandleFunc("/api/audio/generate", generateAudioHandler).Methods("POST")

	// Category and tag routes
	router.HandleFunc("/api/categories", getCategoriesHandler).Methods("GET")
	router.HandleFunc("/api/tags", getTagsHandler).Methods("GET")

	// WebSocket routes
	router.HandleFunc("/ws/books/{id}", wsHandler)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func uploadPDFHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[Upload] Starting new upload request. Method: %s, Content-Type: %s", r.Method, r.Header.Get("Content-Type"))

	var req struct {
		FileURL string `json:"fileUrl"`
		Title   string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("[Upload] Error parsing request: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("[Upload] Received request for title: %s, URL: %s", req.Title, req.FileURL)

	// Create initial book record
	book := &Book{
		ID:        uuid.New().String(),
		Title:     req.Title,
		FileURL:   req.FileURL,
		Status:    "processing",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	log.Printf("[Upload] Created book record with ID: %s", book.ID)

	// Save initial book record
	if err := SaveBook(db, book); err != nil {
		log.Printf("[Upload] Error saving book: %v", err)
		http.Error(w, fmt.Sprintf("Error saving book: %v", err), http.StatusInternalServerError)
		return
	}
	log.Printf("[Upload] Successfully saved book to database")

	// Start processing in background immediately
	go func() {
		log.Printf("[Processing] Starting background processing for book: %s", book.ID)

		// Process PDF first to create segments
		log.Printf("[PDF] Starting PDF processing for book: %s", book.ID)
		if err := processBook(db, book); err != nil {
			log.Printf("[PDF] Error processing PDF: %v", err)
			book.Status = "error"
			UpdateBook(db, book)
			return
		}
		log.Printf("[PDF] Successfully processed PDF for book: %s", book.ID)

		// Now start TTS processing
		log.Printf("[TTS] Starting TTS processing for book: %s", book.ID)
		if err := processTextToSpeech(book); err != nil {
			log.Printf("[TTS] Error processing TTS: %v", err)
			book.Status = "error"
			UpdateBook(db, book)
			return
		}
		log.Printf("[TTS] Successfully processed TTS for book: %s", book.ID)

		log.Printf("[Processing] All processing completed successfully for book: %s", book.ID)
		// Update book status to ready
		book.Status = "ready"
		book.UpdatedAt = time.Now()
		if err := UpdateBook(db, book); err != nil {
			log.Printf("[Processing] Error updating book status: %v", err)
		}
		log.Printf("[Processing] Book status updated to ready: %s", book.ID)
	}()

	log.Printf("[Upload] Returning response for book: %s", book.ID)
	// Return immediate response with book ID
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     book.ID,
		"status": book.Status,
	})
}

// Add WebSocket handler
func wsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[WS] Error upgrading connection: %v", err)
		return
	}

	// Add connection to the map
	wsConnections[bookID] = append(wsConnections[bookID], conn)
	log.Printf("[WS] New connection established for book: %s", bookID)

	// Clean up on disconnect
	go func() {
		<-r.Context().Done()
		wsConnections[bookID] = removeConn(wsConnections[bookID], conn)
		log.Printf("[WS] Connection closed for book: %s", bookID)
	}()
}

func removeConn(conns []*websocket.Conn, conn *websocket.Conn) []*websocket.Conn {
	for i, c := range conns {
		if c == conn {
			return append(conns[:i], conns[i+1:]...)
		}
	}
	return conns
}

// Update processTextToSpeech to notify clients
func processTextToSpeech(book *Book) error {
	log.Printf("[TTS] Starting TTS processing for book: %s", book.ID)

	segments, err := GetAudioSegments(book.ID)
	if err != nil {
		log.Printf("[TTS] Error getting audio segments: %v", err)
		return fmt.Errorf("error getting audio segments: %v", err)
	}
	log.Printf("[TTS] Found %d segments to process", len(segments))

	for i, segment := range segments {
		if segment.Status != "pending" {
			log.Printf("[TTS] Skipping segment %d/%d (ID: %s) - status: %s", i+1, len(segments), segment.ID, segment.Status)
			continue
		}

		if len(strings.TrimSpace(segment.Content)) == 0 {
			log.Printf("[TTS] Skipping empty segment %d/%d (ID: %s)", i+1, len(segments), segment.ID)
			segment.Status = "skipped"
			segment.UpdatedAt = time.Now()
			if err := UpdateAudioSegment(db, &segment); err != nil {
				log.Printf("[TTS] Error updating empty segment status: %v", err)
			}
			continue
		}

		log.Printf("[TTS] Processing segment %d/%d (ID: %s) - Content length: %d", i+1, len(segments), segment.ID, len(segment.Content))
		audioData, err := generateTTSAudio(segment.Content)
		if err != nil {
			log.Printf("[TTS] Error generating audio: %v", err)
			segment.Status = "error"
			UpdateAudioSegment(db, &segment)
			return fmt.Errorf("error generating TTS: %v", err)
		}
		log.Printf("[TTS] Successfully generated audio for segment: %s", segment.ID)

		// Save audio to temporary file for immediate playback
		audioFileName := fmt.Sprintf("tts-%s.mp3", uuid.New().String())
		audioPath := filepath.Join(AppConfig.UploadDir, "audio", audioFileName)
		if err := os.WriteFile(audioPath, audioData, 0644); err != nil {
			log.Printf("[TTS] Error writing audio file: %v", err)
			return fmt.Errorf("error writing audio file: %v", err)
		}
		log.Printf("[TTS] Saved audio file: %s", audioPath)

		// Update segment with local URL and notify frontend immediately
		segment.AudioURL = "/audio/" + audioFileName
		segment.Status = "completed"
		segment.UpdatedAt = time.Now()
		if err := UpdateAudioSegment(db, &segment); err != nil {
			log.Printf("[TTS] Error updating audio segment: %v", err)
			return fmt.Errorf("error updating audio segment: %v", err)
		}

		// Notify WebSocket clients about the new audio
		if conns, ok := wsConnections[book.ID]; ok {
			notification := map[string]interface{}{
				"type":    "audio_ready",
				"segment": segment,
			}
			for _, conn := range conns {
				if err := conn.WriteJSON(notification); err != nil {
					log.Printf("[WS] Error sending notification: %v", err)
				}
			}
			log.Printf("[WS] Notified clients about new audio segment: %s", segment.ID)
		}

		// Upload to UploadThing in background
		go func(segmentID string, audioPath string, audioData []byte) {
			log.Printf("[Upload] Starting UploadThing upload for segment: %s", segmentID)
			uploadURL, err := uploadToUploadThing(audioPath)
			if err != nil {
				log.Printf("[Upload] Error uploading to UploadThing: %v", err)
				return
			}
			log.Printf("[Upload] Successfully uploaded to UploadThing: %s", uploadURL)

			// Update segment with UploadThing URL
			segment.AudioURL = uploadURL
			segment.UpdatedAt = time.Now()
			if err := UpdateAudioSegment(db, &segment); err != nil {
				log.Printf("[Upload] Error updating segment with UploadThing URL: %v", err)
				return
			}

			// Clean up local file
			os.Remove(audioPath)
			log.Printf("[Upload] Cleaned up local file: %s", audioPath)
		}(segment.ID, audioPath, audioData)
	}

	log.Printf("[TTS] Completed TTS processing for all segments of book: %s", book.ID)
	return nil
}

func generateTTSAudio(text string) ([]byte, error) {
	text = strings.TrimSpace(text)
	if text == "" {
		return nil, fmt.Errorf("cannot generate audio for empty text")
	}

	log.Printf("[TTS] Starting TTS generation for text length: %d", len(text))

	// Call Replicate API to generate audio
	replicateURL := AppConfig.ReplicateAPIURL + "/predictions"

	requestBody := map[string]interface{}{
		"version": AppConfig.KokoroModelVersion,
		"input": map[string]interface{}{
			"text":     text,
			"language": "en",
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling request: %v", err)
	}

	// Create initial prediction
	req, err := http.NewRequest("POST", replicateURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	var prediction struct {
		ID     string   `json:"id"`
		Status string   `json:"status"`
		Output []string `json:"output"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&prediction); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	// Poll for completion
	maxAttempts := 60 // 2 minutes total
	for i := 0; i < maxAttempts; i++ {
		time.Sleep(2 * time.Second)

		req, err = http.NewRequest("GET", replicateURL+"/"+prediction.ID, nil)
		if err != nil {
			return nil, fmt.Errorf("error creating poll request: %v", err)
		}
		req.Header.Set("Authorization", "Token "+os.Getenv("REPLICATE_API_TOKEN"))

		resp, err = client.Do(req)
		if err != nil {
			return nil, fmt.Errorf("error polling prediction: %v", err)
		}

		var result struct {
			Status string   `json:"status"`
			Output []string `json:"output"`
			Error  string   `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return nil, fmt.Errorf("error decoding poll response: %v", err)
		}
		resp.Body.Close()

		log.Printf("[TTS] Poll status: %s", result.Status)

		switch result.Status {
		case "succeeded":
			if len(result.Output) == 0 {
				return nil, fmt.Errorf("no output from model")
			}

			// Download the audio file
			audioResp, err := http.Get(result.Output[0])
			if err != nil {
				return nil, fmt.Errorf("error downloading audio: %v", err)
			}
			defer audioResp.Body.Close()

			return io.ReadAll(audioResp.Body)

		case "failed":
			return nil, fmt.Errorf("prediction failed: %s", result.Error)

		case "canceled":
			return nil, fmt.Errorf("prediction was canceled")

		case "completed":
			if len(result.Output) == 0 {
				return nil, fmt.Errorf("no output from model")
			}

			// Download the audio file
			audioResp, err := http.Get(result.Output[0])
			if err != nil {
				return nil, fmt.Errorf("error downloading audio: %v", err)
			}
			defer audioResp.Body.Close()

			return io.ReadAll(audioResp.Body)
		}
	}

	return nil, fmt.Errorf("prediction timed out after %d attempts", maxAttempts)
}

func uploadCoverHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	coverURL, err := fileStorage.SaveCover(file, header.Filename)
	if err != nil {
		http.Error(w, "Error saving cover", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"coverUrl": coverURL})
}

func getBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	query := `
		SELECT id, title, author, cover_url, page_count, language, created_at
		FROM books
		ORDER BY created_at DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving books", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		err := rows.Scan(
			&book.ID,
			&book.Title,
			&book.Author,
			&book.CoverURL,
			&book.PageCount,
			&book.Language,
			&book.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning books", http.StatusInternalServerError)
			return
		}
		books = append(books, book)
	}

	json.NewEncoder(w).Encode(books)
}

func updateProgressHandler(w http.ResponseWriter, r *http.Request) {
	var progress ReadingProgress
	if err := json.NewDecoder(r.Body).Decode(&progress); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if progress.ID == "" {
		progress.ID = uuid.New().String()
	}

	if err := UpdateReadingProgress(db, &progress); err != nil {
		http.Error(w, "Error updating progress", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func getProgressHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	progress, err := GetReadingProgress(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving progress", http.StatusInternalServerError)
		return
	}

	if progress == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(progress)
}

func createBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if bookmark.ID == "" {
		bookmark.ID = uuid.New().String()
	}

	if err := CreateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error creating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func getBookmarksHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := r.Header.Get("X-User-ID") // Get user ID from header

	bookmarks, err := GetBookmarks(db, vars["bookId"], userID)
	if err != nil {
		http.Error(w, "Error retrieving bookmarks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmarks)
}

func updateBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookmark Bookmark
	if err := json.NewDecoder(r.Body).Decode(&bookmark); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	bookmark.ID = vars["id"]
	if err := UpdateBookmark(db, &bookmark); err != nil {
		http.Error(w, "Error updating bookmark", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookmark)
}

func deleteBookmarkHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if err := DeleteBookmark(db, vars["id"]); err != nil {
		http.Error(w, "Error deleting bookmark", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateAudioHandler(w http.ResponseWriter, r *http.Request) {
	var segment AudioSegment
	if err := json.NewDecoder(r.Body).Decode(&segment); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Generate ID if not provided
	if segment.ID == "" {
		segment.ID = uuid.New().String()
	}

	// Set initial status
	segment.Status = "pending"
	segment.CreatedAt = time.Now()
	segment.UpdatedAt = time.Now()

	if err := CreateAudioSegment(db, &segment); err != nil {
		http.Error(w, "Error creating audio segment", http.StatusInternalServerError)
		return
	}

	// Start TTS processing in background
	go processAudioSegment(db, &segment)

	json.NewEncoder(w).Encode(segment)
}

func getAudioSegmentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	segments, err := GetAudioSegments(bookID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to get audio segments: %v", err), http.StatusInternalServerError)
		return
	}

	// Initialize empty array if segments is nil
	if segments == nil {
		segments = []AudioSegment{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(segments)
}

func getCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, description, created_at FROM categories ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving categories", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning categories", http.StatusInternalServerError)
			return
		}
		categories = append(categories, category)
	}

	json.NewEncoder(w).Encode(categories)
}

func getTagsHandler(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, created_at FROM tags ORDER BY name"
	rows, err := db.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving tags", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tags []Tag
	for rows.Next() {
		var tag Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
		)
		if err != nil {
			http.Error(w, "Error scanning tags", http.StatusInternalServerError)
			return
		}
		tags = append(tags, tag)
	}

	json.NewEncoder(w).Encode(tags)
}

func getBookStatusHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":     book.ID,
		"status": book.Status,
	})
}

func updateBookURLHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var req struct {
		CloudURL string `json:"cloudUrl"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	book, err := GetBookByID(db, vars["id"])
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	book.FileURL = req.CloudURL
	if err := UpdateBook(db, book); err != nil {
		http.Error(w, "Error updating book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func generateBookAudioHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Get the book
	book, err := GetBookByID(db, bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Start TTS processing in background
	go func() {
		if err := processTextToSpeech(book); err != nil {
			log.Printf("Error processing TTS for book %s: %v", book.ID, err)
			return
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "processing",
		"message": "Audio generation started",
	})
}

func processBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID := vars["id"]

	// Get the book
	book, err := GetBookByID(db, bookID)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Start processing in background
	go func() {
		if err := processBook(db, book); err != nil {
			log.Printf("Error processing book: %v", err)
			return
		}
	}()

	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "processing",
		"message": "Audio generation started",
	})
}

func uploadToUploadThing(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create multipart form
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// Add file
	part, err := writer.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return "", fmt.Errorf("error creating form file: %v", err)
	}
	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("error copying file: %v", err)
	}
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", AppConfig.UploadThingURL, body)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", "Bearer "+AppConfig.UploadThingToken)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error uploading to UploadThing: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("UploadThing error: %s", resp.Status)
	}

	// Parse response
	var result struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("error decoding response: %v", err)
	}

	return result.URL, nil
}
```

## File: /Users/saint/Desktop/pdf-player/ereader.db

```db
```

## File: /Users/saint/Desktop/pdf-player/tailwind.config.ts

```ts
import type { Config } from "tailwindcss";

const config: Config = {
    darkMode: ["class"],
    content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
    "*.{js,ts,jsx,tsx,mdx}"
  ],
  theme: {
  	extend: {
  		colors: {
  			background: 'hsl(var(--background))',
  			foreground: 'hsl(var(--foreground))',
  			card: {
  				DEFAULT: 'hsl(var(--card))',
  				foreground: 'hsl(var(--card-foreground))'
  			},
  			popover: {
  				DEFAULT: 'hsl(var(--popover))',
  				foreground: 'hsl(var(--popover-foreground))'
  			},
  			primary: {
  				DEFAULT: 'hsl(var(--primary))',
  				foreground: 'hsl(var(--primary-foreground))'
  			},
  			secondary: {
  				DEFAULT: 'hsl(var(--secondary))',
  				foreground: 'hsl(var(--secondary-foreground))'
  			},
  			muted: {
  				DEFAULT: 'hsl(var(--muted))',
  				foreground: 'hsl(var(--muted-foreground))'
  			},
  			accent: {
  				DEFAULT: 'hsl(var(--accent))',
  				foreground: 'hsl(var(--accent-foreground))'
  			},
  			destructive: {
  				DEFAULT: 'hsl(var(--destructive))',
  				foreground: 'hsl(var(--destructive-foreground))'
  			},
  			border: 'hsl(var(--border))',
  			input: 'hsl(var(--input))',
  			ring: 'hsl(var(--ring))',
  			chart: {
  				'1': 'hsl(var(--chart-1))',
  				'2': 'hsl(var(--chart-2))',
  				'3': 'hsl(var(--chart-3))',
  				'4': 'hsl(var(--chart-4))',
  				'5': 'hsl(var(--chart-5))'
  			},
  			sidebar: {
  				DEFAULT: 'hsl(var(--sidebar-background))',
  				foreground: 'hsl(var(--sidebar-foreground))',
  				primary: 'hsl(var(--sidebar-primary))',
  				'primary-foreground': 'hsl(var(--sidebar-primary-foreground))',
  				accent: 'hsl(var(--sidebar-accent))',
  				'accent-foreground': 'hsl(var(--sidebar-accent-foreground))',
  				border: 'hsl(var(--sidebar-border))',
  				ring: 'hsl(var(--sidebar-ring))'
  			}
  		},
  		borderRadius: {
  			lg: 'var(--radius)',
  			md: 'calc(var(--radius) - 2px)',
  			sm: 'calc(var(--radius) - 4px)'
  		},
  		keyframes: {
  			'accordion-down': {
  				from: {
  					height: '0'
  				},
  				to: {
  					height: 'var(--radix-accordion-content-height)'
  				}
  			},
  			'accordion-up': {
  				from: {
  					height: 'var(--radix-accordion-content-height)'
  				},
  				to: {
  					height: '0'
  				}
  			}
  		},
  		animation: {
  			'accordion-down': 'accordion-down 0.2s ease-out',
  			'accordion-up': 'accordion-up 0.2s ease-out'
  		}
  	}
  },
  plugins: [require("tailwindcss-animate")],
};
export default config;
```

## File: /Users/saint/Desktop/pdf-player/styles/globals.css

```css
@tailwind base;
@tailwind components;
@tailwind utilities;

body {
  font-family: Arial, Helvetica, sans-serif;
}

@layer utilities {
  .text-balance {
    text-wrap: balance;
  }
}

@layer base {
  :root {
    --background: 0 0% 100%;
    --foreground: 0 0% 3.9%;
    --card: 0 0% 100%;
    --card-foreground: 0 0% 3.9%;
    --popover: 0 0% 100%;
    --popover-foreground: 0 0% 3.9%;
    --primary: 0 0% 9%;
    --primary-foreground: 0 0% 98%;
    --secondary: 0 0% 96.1%;
    --secondary-foreground: 0 0% 9%;
    --muted: 0 0% 96.1%;
    --muted-foreground: 0 0% 45.1%;
    --accent: 0 0% 96.1%;
    --accent-foreground: 0 0% 9%;
    --destructive: 0 84.2% 60.2%;
    --destructive-foreground: 0 0% 98%;
    --border: 0 0% 89.8%;
    --input: 0 0% 89.8%;
    --ring: 0 0% 3.9%;
    --chart-1: 12 76% 61%;
    --chart-2: 173 58% 39%;
    --chart-3: 197 37% 24%;
    --chart-4: 43 74% 66%;
    --chart-5: 27 87% 67%;
    --radius: 0.5rem;
    --sidebar-background: 0 0% 98%;
    --sidebar-foreground: 240 5.3% 26.1%;
    --sidebar-primary: 240 5.9% 10%;
    --sidebar-primary-foreground: 0 0% 98%;
    --sidebar-accent: 240 4.8% 95.9%;
    --sidebar-accent-foreground: 240 5.9% 10%;
    --sidebar-border: 220 13% 91%;
    --sidebar-ring: 217.2 91.2% 59.8%;
  }
  .dark {
    --background: 0 0% 3.9%;
    --foreground: 0 0% 98%;
    --card: 0 0% 3.9%;
    --card-foreground: 0 0% 98%;
    --popover: 0 0% 3.9%;
    --popover-foreground: 0 0% 98%;
    --primary: 0 0% 98%;
    --primary-foreground: 0 0% 9%;
    --secondary: 0 0% 14.9%;
    --secondary-foreground: 0 0% 98%;
    --muted: 0 0% 14.9%;
    --muted-foreground: 0 0% 63.9%;
    --accent: 0 0% 14.9%;
    --accent-foreground: 0 0% 98%;
    --destructive: 0 62.8% 30.6%;
    --destructive-foreground: 0 0% 98%;
    --border: 0 0% 14.9%;
    --input: 0 0% 14.9%;
    --ring: 0 0% 83.1%;
    --chart-1: 220 70% 50%;
    --chart-2: 160 60% 45%;
    --chart-3: 30 80% 55%;
    --chart-4: 280 65% 60%;
    --chart-5: 340 75% 55%;
    --sidebar-background: 240 5.9% 10%;
    --sidebar-foreground: 240 4.8% 95.9%;
    --sidebar-primary: 224.3 76.3% 48%;
    --sidebar-primary-foreground: 0 0% 100%;
    --sidebar-accent: 240 3.7% 15.9%;
    --sidebar-accent-foreground: 240 4.8% 95.9%;
    --sidebar-border: 240 3.7% 15.9%;
    --sidebar-ring: 217.2 91.2% 59.8%;
  }
}

@layer base {
  * {
    @apply border-border;
  }
  body {
    @apply bg-background text-foreground;
  }
}
```

## File: /Users/saint/Desktop/pdf-player/components/theme-provider.tsx

```tsx
"use client"

import * as React from "react"
import { ThemeProvider as NextThemesProvider } from "next-themes"
import type { ThemeProviderProps } from "next-themes"

export function ThemeProvider({ children, ...props }: ThemeProviderProps) {
  return <NextThemesProvider {...props}>{children}</NextThemesProvider>
}

```

## File: /Users/saint/Desktop/pdf-player/components/ui/aspect-ratio.tsx

```tsx
"use client"

import * as AspectRatioPrimitive from "@radix-ui/react-aspect-ratio"

const AspectRatio = AspectRatioPrimitive.Root

export { AspectRatio }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/alert-dialog.tsx

```tsx
"use client"

import * as React from "react"
import * as AlertDialogPrimitive from "@radix-ui/react-alert-dialog"

import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/ui/button"

const AlertDialog = AlertDialogPrimitive.Root

const AlertDialogTrigger = AlertDialogPrimitive.Trigger

const AlertDialogPortal = AlertDialogPrimitive.Portal

const AlertDialogOverlay = React.forwardRef<
  React.ElementRef<typeof AlertDialogPrimitive.Overlay>,
  React.ComponentPropsWithoutRef<typeof AlertDialogPrimitive.Overlay>
>(({ className, ...props }, ref) => (
  <AlertDialogPrimitive.Overlay
    className={cn(
      "fixed inset-0 z-50 bg-black/80  data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
      className
    )}
    {...props}
    ref={ref}
  />
))
AlertDialogOverlay.displayName = AlertDialogPrimitive.Overlay.displayName

const AlertDialogContent = React.forwardRef<
  React.ElementRef<typeof AlertDialogPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof AlertDialogPrimitive.Content>
>(({ className, ...props }, ref) => (
  <AlertDialogPortal>
    <AlertDialogOverlay />
    <AlertDialogPrimitive.Content
      ref={ref}
      className={cn(
        "fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg",
        className
      )}
      {...props}
    />
  </AlertDialogPortal>
))
AlertDialogContent.displayName = AlertDialogPrimitive.Content.displayName

const AlertDialogHeader = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn(
      "flex flex-col space-y-2 text-center sm:text-left",
      className
    )}
    {...props}
  />
)
AlertDialogHeader.displayName = "AlertDialogHeader"

const AlertDialogFooter = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn(
      "flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
      className
    )}
    {...props}
  />
)
AlertDialogFooter.displayName = "AlertDialogFooter"

const AlertDialogTitle = React.forwardRef<
  React.ElementRef<typeof AlertDialogPrimitive.Title>,
  React.ComponentPropsWithoutRef<typeof AlertDialogPrimitive.Title>
>(({ className, ...props }, ref) => (
  <AlertDialogPrimitive.Title
    ref={ref}
    className={cn("text-lg font-semibold", className)}
    {...props}
  />
))
AlertDialogTitle.displayName = AlertDialogPrimitive.Title.displayName

const AlertDialogDescription = React.forwardRef<
  React.ElementRef<typeof AlertDialogPrimitive.Description>,
  React.ComponentPropsWithoutRef<typeof AlertDialogPrimitive.Description>
>(({ className, ...props }, ref) => (
  <AlertDialogPrimitive.Description
    ref={ref}
    className={cn("text-sm text-muted-foreground", className)}
    {...props}
  />
))
AlertDialogDescription.displayName =
  AlertDialogPrimitive.Description.displayName

const AlertDialogAction = React.forwardRef<
  React.ElementRef<typeof AlertDialogPrimitive.Action>,
  React.ComponentPropsWithoutRef<typeof AlertDialogPrimitive.Action>
>(({ className, ...props }, ref) => (
  <AlertDialogPrimitive.Action
    ref={ref}
    className={cn(buttonVariants(), className)}
    {...props}
  />
))
AlertDialogAction.displayName = AlertDialogPrimitive.Action.displayName

const AlertDialogCancel = React.forwardRef<
  React.ElementRef<typeof AlertDialogPrimitive.Cancel>,
  React.ComponentPropsWithoutRef<typeof AlertDialogPrimitive.Cancel>
>(({ className, ...props }, ref) => (
  <AlertDialogPrimitive.Cancel
    ref={ref}
    className={cn(
      buttonVariants({ variant: "outline" }),
      "mt-2 sm:mt-0",
      className
    )}
    {...props}
  />
))
AlertDialogCancel.displayName = AlertDialogPrimitive.Cancel.displayName

export {
  AlertDialog,
  AlertDialogPortal,
  AlertDialogOverlay,
  AlertDialogTrigger,
  AlertDialogContent,
  AlertDialogHeader,
  AlertDialogFooter,
  AlertDialogTitle,
  AlertDialogDescription,
  AlertDialogAction,
  AlertDialogCancel,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/use-mobile.tsx

```tsx
import * as React from "react"

const MOBILE_BREAKPOINT = 768

export function useIsMobile() {
  const [isMobile, setIsMobile] = React.useState<boolean | undefined>(undefined)

  React.useEffect(() => {
    const mql = window.matchMedia(`(max-width: ${MOBILE_BREAKPOINT - 1}px)`)
    const onChange = () => {
      setIsMobile(window.innerWidth < MOBILE_BREAKPOINT)
    }
    mql.addEventListener("change", onChange)
    setIsMobile(window.innerWidth < MOBILE_BREAKPOINT)
    return () => mql.removeEventListener("change", onChange)
  }, [])

  return !!isMobile
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/pagination.tsx

```tsx
import * as React from "react"
import { ChevronLeft, ChevronRight, MoreHorizontal } from "lucide-react"

import { cn } from "@/lib/utils"
import { ButtonProps, buttonVariants } from "@/components/ui/button"

const Pagination = ({ className, ...props }: React.ComponentProps<"nav">) => (
  <nav
    role="navigation"
    aria-label="pagination"
    className={cn("mx-auto flex w-full justify-center", className)}
    {...props}
  />
)
Pagination.displayName = "Pagination"

const PaginationContent = React.forwardRef<
  HTMLUListElement,
  React.ComponentProps<"ul">
>(({ className, ...props }, ref) => (
  <ul
    ref={ref}
    className={cn("flex flex-row items-center gap-1", className)}
    {...props}
  />
))
PaginationContent.displayName = "PaginationContent"

const PaginationItem = React.forwardRef<
  HTMLLIElement,
  React.ComponentProps<"li">
>(({ className, ...props }, ref) => (
  <li ref={ref} className={cn("", className)} {...props} />
))
PaginationItem.displayName = "PaginationItem"

type PaginationLinkProps = {
  isActive?: boolean
} & Pick<ButtonProps, "size"> &
  React.ComponentProps<"a">

const PaginationLink = ({
  className,
  isActive,
  size = "icon",
  ...props
}: PaginationLinkProps) => (
  <a
    aria-current={isActive ? "page" : undefined}
    className={cn(
      buttonVariants({
        variant: isActive ? "outline" : "ghost",
        size,
      }),
      className
    )}
    {...props}
  />
)
PaginationLink.displayName = "PaginationLink"

const PaginationPrevious = ({
  className,
  ...props
}: React.ComponentProps<typeof PaginationLink>) => (
  <PaginationLink
    aria-label="Go to previous page"
    size="default"
    className={cn("gap-1 pl-2.5", className)}
    {...props}
  >
    <ChevronLeft className="h-4 w-4" />
    <span>Previous</span>
  </PaginationLink>
)
PaginationPrevious.displayName = "PaginationPrevious"

const PaginationNext = ({
  className,
  ...props
}: React.ComponentProps<typeof PaginationLink>) => (
  <PaginationLink
    aria-label="Go to next page"
    size="default"
    className={cn("gap-1 pr-2.5", className)}
    {...props}
  >
    <span>Next</span>
    <ChevronRight className="h-4 w-4" />
  </PaginationLink>
)
PaginationNext.displayName = "PaginationNext"

const PaginationEllipsis = ({
  className,
  ...props
}: React.ComponentProps<"span">) => (
  <span
    aria-hidden
    className={cn("flex h-9 w-9 items-center justify-center", className)}
    {...props}
  >
    <MoreHorizontal className="h-4 w-4" />
    <span className="sr-only">More pages</span>
  </span>
)
PaginationEllipsis.displayName = "PaginationEllipsis"

export {
  Pagination,
  PaginationContent,
  PaginationEllipsis,
  PaginationItem,
  PaginationLink,
  PaginationNext,
  PaginationPrevious,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/tabs.tsx

```tsx
"use client"

import * as React from "react"
import * as TabsPrimitive from "@radix-ui/react-tabs"

import { cn } from "@/lib/utils"

const Tabs = TabsPrimitive.Root

const TabsList = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.List>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.List>
>(({ className, ...props }, ref) => (
  <TabsPrimitive.List
    ref={ref}
    className={cn(
      "inline-flex h-10 items-center justify-center rounded-md bg-muted p-1 text-muted-foreground",
      className
    )}
    {...props}
  />
))
TabsList.displayName = TabsPrimitive.List.displayName

const TabsTrigger = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.Trigger>
>(({ className, ...props }, ref) => (
  <TabsPrimitive.Trigger
    ref={ref}
    className={cn(
      "inline-flex items-center justify-center whitespace-nowrap rounded-sm px-3 py-1.5 text-sm font-medium ring-offset-background transition-all focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=active]:bg-background data-[state=active]:text-foreground data-[state=active]:shadow-sm",
      className
    )}
    {...props}
  />
))
TabsTrigger.displayName = TabsPrimitive.Trigger.displayName

const TabsContent = React.forwardRef<
  React.ElementRef<typeof TabsPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof TabsPrimitive.Content>
>(({ className, ...props }, ref) => (
  <TabsPrimitive.Content
    ref={ref}
    className={cn(
      "mt-2 ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2",
      className
    )}
    {...props}
  />
))
TabsContent.displayName = TabsPrimitive.Content.displayName

export { Tabs, TabsList, TabsTrigger, TabsContent }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/card.tsx

```tsx
import * as React from "react"

import { cn } from "@/lib/utils"

const Card = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn(
      "rounded-lg border bg-card text-card-foreground shadow-sm",
      className
    )}
    {...props}
  />
))
Card.displayName = "Card"

const CardHeader = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("flex flex-col space-y-1.5 p-6", className)}
    {...props}
  />
))
CardHeader.displayName = "CardHeader"

const CardTitle = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn(
      "text-2xl font-semibold leading-none tracking-tight",
      className
    )}
    {...props}
  />
))
CardTitle.displayName = "CardTitle"

const CardDescription = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("text-sm text-muted-foreground", className)}
    {...props}
  />
))
CardDescription.displayName = "CardDescription"

const CardContent = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div ref={ref} className={cn("p-6 pt-0", className)} {...props} />
))
CardContent.displayName = "CardContent"

const CardFooter = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("flex items-center p-6 pt-0", className)}
    {...props}
  />
))
CardFooter.displayName = "CardFooter"

export { Card, CardHeader, CardFooter, CardTitle, CardDescription, CardContent }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/slider.tsx

```tsx
"use client"

import * as React from "react"
import * as SliderPrimitive from "@radix-ui/react-slider"

import { cn } from "@/lib/utils"

const Slider = React.forwardRef<
  React.ElementRef<typeof SliderPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof SliderPrimitive.Root>
>(({ className, ...props }, ref) => (
  <SliderPrimitive.Root
    ref={ref}
    className={cn(
      "relative flex w-full touch-none select-none items-center",
      className
    )}
    {...props}
  >
    <SliderPrimitive.Track className="relative h-2 w-full grow overflow-hidden rounded-full bg-secondary">
      <SliderPrimitive.Range className="absolute h-full bg-primary" />
    </SliderPrimitive.Track>
    <SliderPrimitive.Thumb className="block h-5 w-5 rounded-full border-2 border-primary bg-background ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50" />
  </SliderPrimitive.Root>
))
Slider.displayName = SliderPrimitive.Root.displayName

export { Slider }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/popover.tsx

```tsx
"use client"

import * as React from "react"
import * as PopoverPrimitive from "@radix-ui/react-popover"

import { cn } from "@/lib/utils"

const Popover = PopoverPrimitive.Root

const PopoverTrigger = PopoverPrimitive.Trigger

const PopoverContent = React.forwardRef<
  React.ElementRef<typeof PopoverPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof PopoverPrimitive.Content>
>(({ className, align = "center", sideOffset = 4, ...props }, ref) => (
  <PopoverPrimitive.Portal>
    <PopoverPrimitive.Content
      ref={ref}
      align={align}
      sideOffset={sideOffset}
      className={cn(
        "z-50 w-72 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
        className
      )}
      {...props}
    />
  </PopoverPrimitive.Portal>
))
PopoverContent.displayName = PopoverPrimitive.Content.displayName

export { Popover, PopoverTrigger, PopoverContent }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/progress.tsx

```tsx
"use client"

import * as React from "react"
import * as ProgressPrimitive from "@radix-ui/react-progress"

import { cn } from "@/lib/utils"

const Progress = React.forwardRef<
  React.ElementRef<typeof ProgressPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof ProgressPrimitive.Root>
>(({ className, value, ...props }, ref) => (
  <ProgressPrimitive.Root
    ref={ref}
    className={cn(
      "relative h-4 w-full overflow-hidden rounded-full bg-secondary",
      className
    )}
    {...props}
  >
    <ProgressPrimitive.Indicator
      className="h-full w-full flex-1 bg-primary transition-all"
      style={{ transform: `translateX(-${100 - (value || 0)}%)` }}
    />
  </ProgressPrimitive.Root>
))
Progress.displayName = ProgressPrimitive.Root.displayName

export { Progress }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/toaster.tsx

```tsx
"use client"

import { useToast } from "@/hooks/use-toast"
import {
  Toast,
  ToastClose,
  ToastDescription,
  ToastProvider,
  ToastTitle,
  ToastViewport,
} from "@/components/ui/toast"

export function Toaster() {
  const { toasts } = useToast()

  return (
    <ToastProvider>
      {toasts.map(function ({ id, title, description, action, ...props }) {
        return (
          <Toast key={id} {...props}>
            <div className="grid gap-1">
              {title && <ToastTitle>{title}</ToastTitle>}
              {description && (
                <ToastDescription>{description}</ToastDescription>
              )}
            </div>
            {action}
            <ToastClose />
          </Toast>
        )
      })}
      <ToastViewport />
    </ToastProvider>
  )
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/input-otp.tsx

```tsx
"use client"

import * as React from "react"
import { OTPInput, OTPInputContext } from "input-otp"
import { Dot } from "lucide-react"

import { cn } from "@/lib/utils"

const InputOTP = React.forwardRef<
  React.ElementRef<typeof OTPInput>,
  React.ComponentPropsWithoutRef<typeof OTPInput>
>(({ className, containerClassName, ...props }, ref) => (
  <OTPInput
    ref={ref}
    containerClassName={cn(
      "flex items-center gap-2 has-[:disabled]:opacity-50",
      containerClassName
    )}
    className={cn("disabled:cursor-not-allowed", className)}
    {...props}
  />
))
InputOTP.displayName = "InputOTP"

const InputOTPGroup = React.forwardRef<
  React.ElementRef<"div">,
  React.ComponentPropsWithoutRef<"div">
>(({ className, ...props }, ref) => (
  <div ref={ref} className={cn("flex items-center", className)} {...props} />
))
InputOTPGroup.displayName = "InputOTPGroup"

const InputOTPSlot = React.forwardRef<
  React.ElementRef<"div">,
  React.ComponentPropsWithoutRef<"div"> & { index: number }
>(({ index, className, ...props }, ref) => {
  const inputOTPContext = React.useContext(OTPInputContext)
  const { char, hasFakeCaret, isActive } = inputOTPContext.slots[index]

  return (
    <div
      ref={ref}
      className={cn(
        "relative flex h-10 w-10 items-center justify-center border-y border-r border-input text-sm transition-all first:rounded-l-md first:border-l last:rounded-r-md",
        isActive && "z-10 ring-2 ring-ring ring-offset-background",
        className
      )}
      {...props}
    >
      {char}
      {hasFakeCaret && (
        <div className="pointer-events-none absolute inset-0 flex items-center justify-center">
          <div className="h-4 w-px animate-caret-blink bg-foreground duration-1000" />
        </div>
      )}
    </div>
  )
})
InputOTPSlot.displayName = "InputOTPSlot"

const InputOTPSeparator = React.forwardRef<
  React.ElementRef<"div">,
  React.ComponentPropsWithoutRef<"div">
>(({ ...props }, ref) => (
  <div ref={ref} role="separator" {...props}>
    <Dot />
  </div>
))
InputOTPSeparator.displayName = "InputOTPSeparator"

export { InputOTP, InputOTPGroup, InputOTPSlot, InputOTPSeparator }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/chart.tsx

```tsx
"use client"

import * as React from "react"
import * as RechartsPrimitive from "recharts"

import { cn } from "@/lib/utils"

// Format: { THEME_NAME: CSS_SELECTOR }
const THEMES = { light: "", dark: ".dark" } as const

export type ChartConfig = {
  [k in string]: {
    label?: React.ReactNode
    icon?: React.ComponentType
  } & (
    | { color?: string; theme?: never }
    | { color?: never; theme: Record<keyof typeof THEMES, string> }
  )
}

type ChartContextProps = {
  config: ChartConfig
}

const ChartContext = React.createContext<ChartContextProps | null>(null)

function useChart() {
  const context = React.useContext(ChartContext)

  if (!context) {
    throw new Error("useChart must be used within a <ChartContainer />")
  }

  return context
}

const ChartContainer = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div"> & {
    config: ChartConfig
    children: React.ComponentProps<
      typeof RechartsPrimitive.ResponsiveContainer
    >["children"]
  }
>(({ id, className, children, config, ...props }, ref) => {
  const uniqueId = React.useId()
  const chartId = `chart-${id || uniqueId.replace(/:/g, "")}`

  return (
    <ChartContext.Provider value={{ config }}>
      <div
        data-chart={chartId}
        ref={ref}
        className={cn(
          "flex aspect-video justify-center text-xs [&_.recharts-cartesian-axis-tick_text]:fill-muted-foreground [&_.recharts-cartesian-grid_line[stroke='#ccc']]:stroke-border/50 [&_.recharts-curve.recharts-tooltip-cursor]:stroke-border [&_.recharts-dot[stroke='#fff']]:stroke-transparent [&_.recharts-layer]:outline-none [&_.recharts-polar-grid_[stroke='#ccc']]:stroke-border [&_.recharts-radial-bar-background-sector]:fill-muted [&_.recharts-rectangle.recharts-tooltip-cursor]:fill-muted [&_.recharts-reference-line_[stroke='#ccc']]:stroke-border [&_.recharts-sector[stroke='#fff']]:stroke-transparent [&_.recharts-sector]:outline-none [&_.recharts-surface]:outline-none",
          className
        )}
        {...props}
      >
        <ChartStyle id={chartId} config={config} />
        <RechartsPrimitive.ResponsiveContainer>
          {children}
        </RechartsPrimitive.ResponsiveContainer>
      </div>
    </ChartContext.Provider>
  )
})
ChartContainer.displayName = "Chart"

const ChartStyle = ({ id, config }: { id: string; config: ChartConfig }) => {
  const colorConfig = Object.entries(config).filter(
    ([_, config]) => config.theme || config.color
  )

  if (!colorConfig.length) {
    return null
  }

  return (
    <style
      dangerouslySetInnerHTML={{
        __html: Object.entries(THEMES)
          .map(
            ([theme, prefix]) => `
${prefix} [data-chart=${id}] {
${colorConfig
  .map(([key, itemConfig]) => {
    const color =
      itemConfig.theme?.[theme as keyof typeof itemConfig.theme] ||
      itemConfig.color
    return color ? `  --color-${key}: ${color};` : null
  })
  .join("\n")}
}
`
          )
          .join("\n"),
      }}
    />
  )
}

const ChartTooltip = RechartsPrimitive.Tooltip

const ChartTooltipContent = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<typeof RechartsPrimitive.Tooltip> &
    React.ComponentProps<"div"> & {
      hideLabel?: boolean
      hideIndicator?: boolean
      indicator?: "line" | "dot" | "dashed"
      nameKey?: string
      labelKey?: string
    }
>(
  (
    {
      active,
      payload,
      className,
      indicator = "dot",
      hideLabel = false,
      hideIndicator = false,
      label,
      labelFormatter,
      labelClassName,
      formatter,
      color,
      nameKey,
      labelKey,
    },
    ref
  ) => {
    const { config } = useChart()

    const tooltipLabel = React.useMemo(() => {
      if (hideLabel || !payload?.length) {
        return null
      }

      const [item] = payload
      const key = `${labelKey || item.dataKey || item.name || "value"}`
      const itemConfig = getPayloadConfigFromPayload(config, item, key)
      const value =
        !labelKey && typeof label === "string"
          ? config[label as keyof typeof config]?.label || label
          : itemConfig?.label

      if (labelFormatter) {
        return (
          <div className={cn("font-medium", labelClassName)}>
            {labelFormatter(value, payload)}
          </div>
        )
      }

      if (!value) {
        return null
      }

      return <div className={cn("font-medium", labelClassName)}>{value}</div>
    }, [
      label,
      labelFormatter,
      payload,
      hideLabel,
      labelClassName,
      config,
      labelKey,
    ])

    if (!active || !payload?.length) {
      return null
    }

    const nestLabel = payload.length === 1 && indicator !== "dot"

    return (
      <div
        ref={ref}
        className={cn(
          "grid min-w-[8rem] items-start gap-1.5 rounded-lg border border-border/50 bg-background px-2.5 py-1.5 text-xs shadow-xl",
          className
        )}
      >
        {!nestLabel ? tooltipLabel : null}
        <div className="grid gap-1.5">
          {payload.map((item, index) => {
            const key = `${nameKey || item.name || item.dataKey || "value"}`
            const itemConfig = getPayloadConfigFromPayload(config, item, key)
            const indicatorColor = color || item.payload.fill || item.color

            return (
              <div
                key={item.dataKey}
                className={cn(
                  "flex w-full flex-wrap items-stretch gap-2 [&>svg]:h-2.5 [&>svg]:w-2.5 [&>svg]:text-muted-foreground",
                  indicator === "dot" && "items-center"
                )}
              >
                {formatter && item?.value !== undefined && item.name ? (
                  formatter(item.value, item.name, item, index, item.payload)
                ) : (
                  <>
                    {itemConfig?.icon ? (
                      <itemConfig.icon />
                    ) : (
                      !hideIndicator && (
                        <div
                          className={cn(
                            "shrink-0 rounded-[2px] border-[--color-border] bg-[--color-bg]",
                            {
                              "h-2.5 w-2.5": indicator === "dot",
                              "w-1": indicator === "line",
                              "w-0 border-[1.5px] border-dashed bg-transparent":
                                indicator === "dashed",
                              "my-0.5": nestLabel && indicator === "dashed",
                            }
                          )}
                          style={
                            {
                              "--color-bg": indicatorColor,
                              "--color-border": indicatorColor,
                            } as React.CSSProperties
                          }
                        />
                      )
                    )}
                    <div
                      className={cn(
                        "flex flex-1 justify-between leading-none",
                        nestLabel ? "items-end" : "items-center"
                      )}
                    >
                      <div className="grid gap-1.5">
                        {nestLabel ? tooltipLabel : null}
                        <span className="text-muted-foreground">
                          {itemConfig?.label || item.name}
                        </span>
                      </div>
                      {item.value && (
                        <span className="font-mono font-medium tabular-nums text-foreground">
                          {item.value.toLocaleString()}
                        </span>
                      )}
                    </div>
                  </>
                )}
              </div>
            )
          })}
        </div>
      </div>
    )
  }
)
ChartTooltipContent.displayName = "ChartTooltip"

const ChartLegend = RechartsPrimitive.Legend

const ChartLegendContent = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div"> &
    Pick<RechartsPrimitive.LegendProps, "payload" | "verticalAlign"> & {
      hideIcon?: boolean
      nameKey?: string
    }
>(
  (
    { className, hideIcon = false, payload, verticalAlign = "bottom", nameKey },
    ref
  ) => {
    const { config } = useChart()

    if (!payload?.length) {
      return null
    }

    return (
      <div
        ref={ref}
        className={cn(
          "flex items-center justify-center gap-4",
          verticalAlign === "top" ? "pb-3" : "pt-3",
          className
        )}
      >
        {payload.map((item) => {
          const key = `${nameKey || item.dataKey || "value"}`
          const itemConfig = getPayloadConfigFromPayload(config, item, key)

          return (
            <div
              key={item.value}
              className={cn(
                "flex items-center gap-1.5 [&>svg]:h-3 [&>svg]:w-3 [&>svg]:text-muted-foreground"
              )}
            >
              {itemConfig?.icon && !hideIcon ? (
                <itemConfig.icon />
              ) : (
                <div
                  className="h-2 w-2 shrink-0 rounded-[2px]"
                  style={{
                    backgroundColor: item.color,
                  }}
                />
              )}
              {itemConfig?.label}
            </div>
          )
        })}
      </div>
    )
  }
)
ChartLegendContent.displayName = "ChartLegend"

// Helper to extract item config from a payload.
function getPayloadConfigFromPayload(
  config: ChartConfig,
  payload: unknown,
  key: string
) {
  if (typeof payload !== "object" || payload === null) {
    return undefined
  }

  const payloadPayload =
    "payload" in payload &&
    typeof payload.payload === "object" &&
    payload.payload !== null
      ? payload.payload
      : undefined

  let configLabelKey: string = key

  if (
    key in payload &&
    typeof payload[key as keyof typeof payload] === "string"
  ) {
    configLabelKey = payload[key as keyof typeof payload] as string
  } else if (
    payloadPayload &&
    key in payloadPayload &&
    typeof payloadPayload[key as keyof typeof payloadPayload] === "string"
  ) {
    configLabelKey = payloadPayload[
      key as keyof typeof payloadPayload
    ] as string
  }

  return configLabelKey in config
    ? config[configLabelKey]
    : config[key as keyof typeof config]
}

export {
  ChartContainer,
  ChartTooltip,
  ChartTooltipContent,
  ChartLegend,
  ChartLegendContent,
  ChartStyle,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/hover-card.tsx

```tsx
"use client"

import * as React from "react"
import * as HoverCardPrimitive from "@radix-ui/react-hover-card"

import { cn } from "@/lib/utils"

const HoverCard = HoverCardPrimitive.Root

const HoverCardTrigger = HoverCardPrimitive.Trigger

const HoverCardContent = React.forwardRef<
  React.ElementRef<typeof HoverCardPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof HoverCardPrimitive.Content>
>(({ className, align = "center", sideOffset = 4, ...props }, ref) => (
  <HoverCardPrimitive.Content
    ref={ref}
    align={align}
    sideOffset={sideOffset}
    className={cn(
      "z-50 w-64 rounded-md border bg-popover p-4 text-popover-foreground shadow-md outline-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
      className
    )}
    {...props}
  />
))
HoverCardContent.displayName = HoverCardPrimitive.Content.displayName

export { HoverCard, HoverCardTrigger, HoverCardContent }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/sheet.tsx

```tsx
"use client"

import * as React from "react"
import * as SheetPrimitive from "@radix-ui/react-dialog"
import { cva, type VariantProps } from "class-variance-authority"
import { X } from "lucide-react"

import { cn } from "@/lib/utils"

const Sheet = SheetPrimitive.Root

const SheetTrigger = SheetPrimitive.Trigger

const SheetClose = SheetPrimitive.Close

const SheetPortal = SheetPrimitive.Portal

const SheetOverlay = React.forwardRef<
  React.ElementRef<typeof SheetPrimitive.Overlay>,
  React.ComponentPropsWithoutRef<typeof SheetPrimitive.Overlay>
>(({ className, ...props }, ref) => (
  <SheetPrimitive.Overlay
    className={cn(
      "fixed inset-0 z-50 bg-black/80  data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
      className
    )}
    {...props}
    ref={ref}
  />
))
SheetOverlay.displayName = SheetPrimitive.Overlay.displayName

const sheetVariants = cva(
  "fixed z-50 gap-4 bg-background p-6 shadow-lg transition ease-in-out data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:duration-300 data-[state=open]:duration-500",
  {
    variants: {
      side: {
        top: "inset-x-0 top-0 border-b data-[state=closed]:slide-out-to-top data-[state=open]:slide-in-from-top",
        bottom:
          "inset-x-0 bottom-0 border-t data-[state=closed]:slide-out-to-bottom data-[state=open]:slide-in-from-bottom",
        left: "inset-y-0 left-0 h-full w-3/4 border-r data-[state=closed]:slide-out-to-left data-[state=open]:slide-in-from-left sm:max-w-sm",
        right:
          "inset-y-0 right-0 h-full w-3/4  border-l data-[state=closed]:slide-out-to-right data-[state=open]:slide-in-from-right sm:max-w-sm",
      },
    },
    defaultVariants: {
      side: "right",
    },
  }
)

interface SheetContentProps
  extends React.ComponentPropsWithoutRef<typeof SheetPrimitive.Content>,
    VariantProps<typeof sheetVariants> {}

const SheetContent = React.forwardRef<
  React.ElementRef<typeof SheetPrimitive.Content>,
  SheetContentProps
>(({ side = "right", className, children, ...props }, ref) => (
  <SheetPortal>
    <SheetOverlay />
    <SheetPrimitive.Content
      ref={ref}
      className={cn(sheetVariants({ side }), className)}
      {...props}
    >
      {children}
      <SheetPrimitive.Close className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-secondary">
        <X className="h-4 w-4" />
        <span className="sr-only">Close</span>
      </SheetPrimitive.Close>
    </SheetPrimitive.Content>
  </SheetPortal>
))
SheetContent.displayName = SheetPrimitive.Content.displayName

const SheetHeader = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn(
      "flex flex-col space-y-2 text-center sm:text-left",
      className
    )}
    {...props}
  />
)
SheetHeader.displayName = "SheetHeader"

const SheetFooter = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn(
      "flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
      className
    )}
    {...props}
  />
)
SheetFooter.displayName = "SheetFooter"

const SheetTitle = React.forwardRef<
  React.ElementRef<typeof SheetPrimitive.Title>,
  React.ComponentPropsWithoutRef<typeof SheetPrimitive.Title>
>(({ className, ...props }, ref) => (
  <SheetPrimitive.Title
    ref={ref}
    className={cn("text-lg font-semibold text-foreground", className)}
    {...props}
  />
))
SheetTitle.displayName = SheetPrimitive.Title.displayName

const SheetDescription = React.forwardRef<
  React.ElementRef<typeof SheetPrimitive.Description>,
  React.ComponentPropsWithoutRef<typeof SheetPrimitive.Description>
>(({ className, ...props }, ref) => (
  <SheetPrimitive.Description
    ref={ref}
    className={cn("text-sm text-muted-foreground", className)}
    {...props}
  />
))
SheetDescription.displayName = SheetPrimitive.Description.displayName

export {
  Sheet,
  SheetPortal,
  SheetOverlay,
  SheetTrigger,
  SheetClose,
  SheetContent,
  SheetHeader,
  SheetFooter,
  SheetTitle,
  SheetDescription,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/scroll-area.tsx

```tsx
"use client"

import * as React from "react"
import * as ScrollAreaPrimitive from "@radix-ui/react-scroll-area"

import { cn } from "@/lib/utils"

const ScrollArea = React.forwardRef<
  React.ElementRef<typeof ScrollAreaPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof ScrollAreaPrimitive.Root>
>(({ className, children, ...props }, ref) => (
  <ScrollAreaPrimitive.Root
    ref={ref}
    className={cn("relative overflow-hidden", className)}
    {...props}
  >
    <ScrollAreaPrimitive.Viewport className="h-full w-full rounded-[inherit]">
      {children}
    </ScrollAreaPrimitive.Viewport>
    <ScrollBar />
    <ScrollAreaPrimitive.Corner />
  </ScrollAreaPrimitive.Root>
))
ScrollArea.displayName = ScrollAreaPrimitive.Root.displayName

const ScrollBar = React.forwardRef<
  React.ElementRef<typeof ScrollAreaPrimitive.ScrollAreaScrollbar>,
  React.ComponentPropsWithoutRef<typeof ScrollAreaPrimitive.ScrollAreaScrollbar>
>(({ className, orientation = "vertical", ...props }, ref) => (
  <ScrollAreaPrimitive.ScrollAreaScrollbar
    ref={ref}
    orientation={orientation}
    className={cn(
      "flex touch-none select-none transition-colors",
      orientation === "vertical" &&
        "h-full w-2.5 border-l border-l-transparent p-[1px]",
      orientation === "horizontal" &&
        "h-2.5 flex-col border-t border-t-transparent p-[1px]",
      className
    )}
    {...props}
  >
    <ScrollAreaPrimitive.ScrollAreaThumb className="relative flex-1 rounded-full bg-border" />
  </ScrollAreaPrimitive.ScrollAreaScrollbar>
))
ScrollBar.displayName = ScrollAreaPrimitive.ScrollAreaScrollbar.displayName

export { ScrollArea, ScrollBar }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/resizable.tsx

```tsx
"use client"

import { GripVertical } from "lucide-react"
import * as ResizablePrimitive from "react-resizable-panels"

import { cn } from "@/lib/utils"

const ResizablePanelGroup = ({
  className,
  ...props
}: React.ComponentProps<typeof ResizablePrimitive.PanelGroup>) => (
  <ResizablePrimitive.PanelGroup
    className={cn(
      "flex h-full w-full data-[panel-group-direction=vertical]:flex-col",
      className
    )}
    {...props}
  />
)

const ResizablePanel = ResizablePrimitive.Panel

const ResizableHandle = ({
  withHandle,
  className,
  ...props
}: React.ComponentProps<typeof ResizablePrimitive.PanelResizeHandle> & {
  withHandle?: boolean
}) => (
  <ResizablePrimitive.PanelResizeHandle
    className={cn(
      "relative flex w-px items-center justify-center bg-border after:absolute after:inset-y-0 after:left-1/2 after:w-1 after:-translate-x-1/2 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring focus-visible:ring-offset-1 data-[panel-group-direction=vertical]:h-px data-[panel-group-direction=vertical]:w-full data-[panel-group-direction=vertical]:after:left-0 data-[panel-group-direction=vertical]:after:h-1 data-[panel-group-direction=vertical]:after:w-full data-[panel-group-direction=vertical]:after:-translate-y-1/2 data-[panel-group-direction=vertical]:after:translate-x-0 [&[data-panel-group-direction=vertical]>div]:rotate-90",
      className
    )}
    {...props}
  >
    {withHandle && (
      <div className="z-10 flex h-4 w-3 items-center justify-center rounded-sm border bg-border">
        <GripVertical className="h-2.5 w-2.5" />
      </div>
    )}
  </ResizablePrimitive.PanelResizeHandle>
)

export { ResizablePanelGroup, ResizablePanel, ResizableHandle }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/label.tsx

```tsx
"use client"

import * as React from "react"
import * as LabelPrimitive from "@radix-ui/react-label"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const labelVariants = cva(
  "text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
)

const Label = React.forwardRef<
  React.ElementRef<typeof LabelPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof LabelPrimitive.Root> &
    VariantProps<typeof labelVariants>
>(({ className, ...props }, ref) => (
  <LabelPrimitive.Root
    ref={ref}
    className={cn(labelVariants(), className)}
    {...props}
  />
))
Label.displayName = LabelPrimitive.Root.displayName

export { Label }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/sonner.tsx

```tsx
"use client"

import { useTheme } from "next-themes"
import { Toaster as Sonner } from "sonner"

type ToasterProps = React.ComponentProps<typeof Sonner>

const Toaster = ({ ...props }: ToasterProps) => {
  const { theme = "system" } = useTheme()

  return (
    <Sonner
      theme={theme as ToasterProps["theme"]}
      className="toaster group"
      toastOptions={{
        classNames: {
          toast:
            "group toast group-[.toaster]:bg-background group-[.toaster]:text-foreground group-[.toaster]:border-border group-[.toaster]:shadow-lg",
          description: "group-[.toast]:text-muted-foreground",
          actionButton:
            "group-[.toast]:bg-primary group-[.toast]:text-primary-foreground",
          cancelButton:
            "group-[.toast]:bg-muted group-[.toast]:text-muted-foreground",
        },
      }}
      {...props}
    />
  )
}

export { Toaster }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/navigation-menu.tsx

```tsx
import * as React from "react"
import * as NavigationMenuPrimitive from "@radix-ui/react-navigation-menu"
import { cva } from "class-variance-authority"
import { ChevronDown } from "lucide-react"

import { cn } from "@/lib/utils"

const NavigationMenu = React.forwardRef<
  React.ElementRef<typeof NavigationMenuPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof NavigationMenuPrimitive.Root>
>(({ className, children, ...props }, ref) => (
  <NavigationMenuPrimitive.Root
    ref={ref}
    className={cn(
      "relative z-10 flex max-w-max flex-1 items-center justify-center",
      className
    )}
    {...props}
  >
    {children}
    <NavigationMenuViewport />
  </NavigationMenuPrimitive.Root>
))
NavigationMenu.displayName = NavigationMenuPrimitive.Root.displayName

const NavigationMenuList = React.forwardRef<
  React.ElementRef<typeof NavigationMenuPrimitive.List>,
  React.ComponentPropsWithoutRef<typeof NavigationMenuPrimitive.List>
>(({ className, ...props }, ref) => (
  <NavigationMenuPrimitive.List
    ref={ref}
    className={cn(
      "group flex flex-1 list-none items-center justify-center space-x-1",
      className
    )}
    {...props}
  />
))
NavigationMenuList.displayName = NavigationMenuPrimitive.List.displayName

const NavigationMenuItem = NavigationMenuPrimitive.Item

const navigationMenuTriggerStyle = cva(
  "group inline-flex h-10 w-max items-center justify-center rounded-md bg-background px-4 py-2 text-sm font-medium transition-colors hover:bg-accent hover:text-accent-foreground focus:bg-accent focus:text-accent-foreground focus:outline-none disabled:pointer-events-none disabled:opacity-50 data-[active]:bg-accent/50 data-[state=open]:bg-accent/50"
)

const NavigationMenuTrigger = React.forwardRef<
  React.ElementRef<typeof NavigationMenuPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof NavigationMenuPrimitive.Trigger>
>(({ className, children, ...props }, ref) => (
  <NavigationMenuPrimitive.Trigger
    ref={ref}
    className={cn(navigationMenuTriggerStyle(), "group", className)}
    {...props}
  >
    {children}{" "}
    <ChevronDown
      className="relative top-[1px] ml-1 h-3 w-3 transition duration-200 group-data-[state=open]:rotate-180"
      aria-hidden="true"
    />
  </NavigationMenuPrimitive.Trigger>
))
NavigationMenuTrigger.displayName = NavigationMenuPrimitive.Trigger.displayName

const NavigationMenuContent = React.forwardRef<
  React.ElementRef<typeof NavigationMenuPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof NavigationMenuPrimitive.Content>
>(({ className, ...props }, ref) => (
  <NavigationMenuPrimitive.Content
    ref={ref}
    className={cn(
      "left-0 top-0 w-full data-[motion^=from-]:animate-in data-[motion^=to-]:animate-out data-[motion^=from-]:fade-in data-[motion^=to-]:fade-out data-[motion=from-end]:slide-in-from-right-52 data-[motion=from-start]:slide-in-from-left-52 data-[motion=to-end]:slide-out-to-right-52 data-[motion=to-start]:slide-out-to-left-52 md:absolute md:w-auto ",
      className
    )}
    {...props}
  />
))
NavigationMenuContent.displayName = NavigationMenuPrimitive.Content.displayName

const NavigationMenuLink = NavigationMenuPrimitive.Link

const NavigationMenuViewport = React.forwardRef<
  React.ElementRef<typeof NavigationMenuPrimitive.Viewport>,
  React.ComponentPropsWithoutRef<typeof NavigationMenuPrimitive.Viewport>
>(({ className, ...props }, ref) => (
  <div className={cn("absolute left-0 top-full flex justify-center")}>
    <NavigationMenuPrimitive.Viewport
      className={cn(
        "origin-top-center relative mt-1.5 h-[var(--radix-navigation-menu-viewport-height)] w-full overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-lg data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-90 md:w-[var(--radix-navigation-menu-viewport-width)]",
        className
      )}
      ref={ref}
      {...props}
    />
  </div>
))
NavigationMenuViewport.displayName =
  NavigationMenuPrimitive.Viewport.displayName

const NavigationMenuIndicator = React.forwardRef<
  React.ElementRef<typeof NavigationMenuPrimitive.Indicator>,
  React.ComponentPropsWithoutRef<typeof NavigationMenuPrimitive.Indicator>
>(({ className, ...props }, ref) => (
  <NavigationMenuPrimitive.Indicator
    ref={ref}
    className={cn(
      "top-full z-[1] flex h-1.5 items-end justify-center overflow-hidden data-[state=visible]:animate-in data-[state=hidden]:animate-out data-[state=hidden]:fade-out data-[state=visible]:fade-in",
      className
    )}
    {...props}
  >
    <div className="relative top-[60%] h-2 w-2 rotate-45 rounded-tl-sm bg-border shadow-md" />
  </NavigationMenuPrimitive.Indicator>
))
NavigationMenuIndicator.displayName =
  NavigationMenuPrimitive.Indicator.displayName

export {
  navigationMenuTriggerStyle,
  NavigationMenu,
  NavigationMenuList,
  NavigationMenuItem,
  NavigationMenuContent,
  NavigationMenuTrigger,
  NavigationMenuLink,
  NavigationMenuIndicator,
  NavigationMenuViewport,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/accordion.tsx

```tsx
"use client"

import * as React from "react"
import * as AccordionPrimitive from "@radix-ui/react-accordion"
import { ChevronDown } from "lucide-react"

import { cn } from "@/lib/utils"

const Accordion = AccordionPrimitive.Root

const AccordionItem = React.forwardRef<
  React.ElementRef<typeof AccordionPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof AccordionPrimitive.Item>
>(({ className, ...props }, ref) => (
  <AccordionPrimitive.Item
    ref={ref}
    className={cn("border-b", className)}
    {...props}
  />
))
AccordionItem.displayName = "AccordionItem"

const AccordionTrigger = React.forwardRef<
  React.ElementRef<typeof AccordionPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof AccordionPrimitive.Trigger>
>(({ className, children, ...props }, ref) => (
  <AccordionPrimitive.Header className="flex">
    <AccordionPrimitive.Trigger
      ref={ref}
      className={cn(
        "flex flex-1 items-center justify-between py-4 font-medium transition-all hover:underline [&[data-state=open]>svg]:rotate-180",
        className
      )}
      {...props}
    >
      {children}
      <ChevronDown className="h-4 w-4 shrink-0 transition-transform duration-200" />
    </AccordionPrimitive.Trigger>
  </AccordionPrimitive.Header>
))
AccordionTrigger.displayName = AccordionPrimitive.Trigger.displayName

const AccordionContent = React.forwardRef<
  React.ElementRef<typeof AccordionPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof AccordionPrimitive.Content>
>(({ className, children, ...props }, ref) => (
  <AccordionPrimitive.Content
    ref={ref}
    className="overflow-hidden text-sm transition-all data-[state=closed]:animate-accordion-up data-[state=open]:animate-accordion-down"
    {...props}
  >
    <div className={cn("pb-4 pt-0", className)}>{children}</div>
  </AccordionPrimitive.Content>
))

AccordionContent.displayName = AccordionPrimitive.Content.displayName

export { Accordion, AccordionItem, AccordionTrigger, AccordionContent }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/drawer.tsx

```tsx
"use client"

import * as React from "react"
import { Drawer as DrawerPrimitive } from "vaul"

import { cn } from "@/lib/utils"

const Drawer = ({
  shouldScaleBackground = true,
  ...props
}: React.ComponentProps<typeof DrawerPrimitive.Root>) => (
  <DrawerPrimitive.Root
    shouldScaleBackground={shouldScaleBackground}
    {...props}
  />
)
Drawer.displayName = "Drawer"

const DrawerTrigger = DrawerPrimitive.Trigger

const DrawerPortal = DrawerPrimitive.Portal

const DrawerClose = DrawerPrimitive.Close

const DrawerOverlay = React.forwardRef<
  React.ElementRef<typeof DrawerPrimitive.Overlay>,
  React.ComponentPropsWithoutRef<typeof DrawerPrimitive.Overlay>
>(({ className, ...props }, ref) => (
  <DrawerPrimitive.Overlay
    ref={ref}
    className={cn("fixed inset-0 z-50 bg-black/80", className)}
    {...props}
  />
))
DrawerOverlay.displayName = DrawerPrimitive.Overlay.displayName

const DrawerContent = React.forwardRef<
  React.ElementRef<typeof DrawerPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DrawerPrimitive.Content>
>(({ className, children, ...props }, ref) => (
  <DrawerPortal>
    <DrawerOverlay />
    <DrawerPrimitive.Content
      ref={ref}
      className={cn(
        "fixed inset-x-0 bottom-0 z-50 mt-24 flex h-auto flex-col rounded-t-[10px] border bg-background",
        className
      )}
      {...props}
    >
      <div className="mx-auto mt-4 h-2 w-[100px] rounded-full bg-muted" />
      {children}
    </DrawerPrimitive.Content>
  </DrawerPortal>
))
DrawerContent.displayName = "DrawerContent"

const DrawerHeader = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn("grid gap-1.5 p-4 text-center sm:text-left", className)}
    {...props}
  />
)
DrawerHeader.displayName = "DrawerHeader"

const DrawerFooter = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn("mt-auto flex flex-col gap-2 p-4", className)}
    {...props}
  />
)
DrawerFooter.displayName = "DrawerFooter"

const DrawerTitle = React.forwardRef<
  React.ElementRef<typeof DrawerPrimitive.Title>,
  React.ComponentPropsWithoutRef<typeof DrawerPrimitive.Title>
>(({ className, ...props }, ref) => (
  <DrawerPrimitive.Title
    ref={ref}
    className={cn(
      "text-lg font-semibold leading-none tracking-tight",
      className
    )}
    {...props}
  />
))
DrawerTitle.displayName = DrawerPrimitive.Title.displayName

const DrawerDescription = React.forwardRef<
  React.ElementRef<typeof DrawerPrimitive.Description>,
  React.ComponentPropsWithoutRef<typeof DrawerPrimitive.Description>
>(({ className, ...props }, ref) => (
  <DrawerPrimitive.Description
    ref={ref}
    className={cn("text-sm text-muted-foreground", className)}
    {...props}
  />
))
DrawerDescription.displayName = DrawerPrimitive.Description.displayName

export {
  Drawer,
  DrawerPortal,
  DrawerOverlay,
  DrawerTrigger,
  DrawerClose,
  DrawerContent,
  DrawerHeader,
  DrawerFooter,
  DrawerTitle,
  DrawerDescription,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/tooltip.tsx

```tsx
"use client"

import * as React from "react"
import * as TooltipPrimitive from "@radix-ui/react-tooltip"

import { cn } from "@/lib/utils"

const TooltipProvider = TooltipPrimitive.Provider

const Tooltip = TooltipPrimitive.Root

const TooltipTrigger = TooltipPrimitive.Trigger

const TooltipContent = React.forwardRef<
  React.ElementRef<typeof TooltipPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof TooltipPrimitive.Content>
>(({ className, sideOffset = 4, ...props }, ref) => (
  <TooltipPrimitive.Content
    ref={ref}
    sideOffset={sideOffset}
    className={cn(
      "z-50 overflow-hidden rounded-md border bg-popover px-3 py-1.5 text-sm text-popover-foreground shadow-md animate-in fade-in-0 zoom-in-95 data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=closed]:zoom-out-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
      className
    )}
    {...props}
  />
))
TooltipContent.displayName = TooltipPrimitive.Content.displayName

export { Tooltip, TooltipTrigger, TooltipContent, TooltipProvider }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/alert.tsx

```tsx
import * as React from "react"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const alertVariants = cva(
  "relative w-full rounded-lg border p-4 [&>svg~*]:pl-7 [&>svg+div]:translate-y-[-3px] [&>svg]:absolute [&>svg]:left-4 [&>svg]:top-4 [&>svg]:text-foreground",
  {
    variants: {
      variant: {
        default: "bg-background text-foreground",
        destructive:
          "border-destructive/50 text-destructive dark:border-destructive [&>svg]:text-destructive",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  }
)

const Alert = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement> & VariantProps<typeof alertVariants>
>(({ className, variant, ...props }, ref) => (
  <div
    ref={ref}
    role="alert"
    className={cn(alertVariants({ variant }), className)}
    {...props}
  />
))
Alert.displayName = "Alert"

const AlertTitle = React.forwardRef<
  HTMLParagraphElement,
  React.HTMLAttributes<HTMLHeadingElement>
>(({ className, ...props }, ref) => (
  <h5
    ref={ref}
    className={cn("mb-1 font-medium leading-none tracking-tight", className)}
    {...props}
  />
))
AlertTitle.displayName = "AlertTitle"

const AlertDescription = React.forwardRef<
  HTMLParagraphElement,
  React.HTMLAttributes<HTMLParagraphElement>
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    className={cn("text-sm [&_p]:leading-relaxed", className)}
    {...props}
  />
))
AlertDescription.displayName = "AlertDescription"

export { Alert, AlertTitle, AlertDescription }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/use-toast.ts

```ts
"use client"

// Inspired by react-hot-toast library
import * as React from "react"

import type {
  ToastActionElement,
  ToastProps,
} from "@/components/ui/toast"

const TOAST_LIMIT = 1
const TOAST_REMOVE_DELAY = 1000000

type ToasterToast = ToastProps & {
  id: string
  title?: React.ReactNode
  description?: React.ReactNode
  action?: ToastActionElement
}

const actionTypes = {
  ADD_TOAST: "ADD_TOAST",
  UPDATE_TOAST: "UPDATE_TOAST",
  DISMISS_TOAST: "DISMISS_TOAST",
  REMOVE_TOAST: "REMOVE_TOAST",
} as const

let count = 0

function genId() {
  count = (count + 1) % Number.MAX_SAFE_INTEGER
  return count.toString()
}

type ActionType = typeof actionTypes

type Action =
  | {
      type: ActionType["ADD_TOAST"]
      toast: ToasterToast
    }
  | {
      type: ActionType["UPDATE_TOAST"]
      toast: Partial<ToasterToast>
    }
  | {
      type: ActionType["DISMISS_TOAST"]
      toastId?: ToasterToast["id"]
    }
  | {
      type: ActionType["REMOVE_TOAST"]
      toastId?: ToasterToast["id"]
    }

interface State {
  toasts: ToasterToast[]
}

const toastTimeouts = new Map<string, ReturnType<typeof setTimeout>>()

const addToRemoveQueue = (toastId: string) => {
  if (toastTimeouts.has(toastId)) {
    return
  }

  const timeout = setTimeout(() => {
    toastTimeouts.delete(toastId)
    dispatch({
      type: "REMOVE_TOAST",
      toastId: toastId,
    })
  }, TOAST_REMOVE_DELAY)

  toastTimeouts.set(toastId, timeout)
}

export const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case "ADD_TOAST":
      return {
        ...state,
        toasts: [action.toast, ...state.toasts].slice(0, TOAST_LIMIT),
      }

    case "UPDATE_TOAST":
      return {
        ...state,
        toasts: state.toasts.map((t) =>
          t.id === action.toast.id ? { ...t, ...action.toast } : t
        ),
      }

    case "DISMISS_TOAST": {
      const { toastId } = action

      // ! Side effects ! - This could be extracted into a dismissToast() action,
      // but I'll keep it here for simplicity
      if (toastId) {
        addToRemoveQueue(toastId)
      } else {
        state.toasts.forEach((toast) => {
          addToRemoveQueue(toast.id)
        })
      }

      return {
        ...state,
        toasts: state.toasts.map((t) =>
          t.id === toastId || toastId === undefined
            ? {
                ...t,
                open: false,
              }
            : t
        ),
      }
    }
    case "REMOVE_TOAST":
      if (action.toastId === undefined) {
        return {
          ...state,
          toasts: [],
        }
      }
      return {
        ...state,
        toasts: state.toasts.filter((t) => t.id !== action.toastId),
      }
  }
}

const listeners: Array<(state: State) => void> = []

let memoryState: State = { toasts: [] }

function dispatch(action: Action) {
  memoryState = reducer(memoryState, action)
  listeners.forEach((listener) => {
    listener(memoryState)
  })
}

type Toast = Omit<ToasterToast, "id">

function toast({ ...props }: Toast) {
  const id = genId()

  const update = (props: ToasterToast) =>
    dispatch({
      type: "UPDATE_TOAST",
      toast: { ...props, id },
    })
  const dismiss = () => dispatch({ type: "DISMISS_TOAST", toastId: id })

  dispatch({
    type: "ADD_TOAST",
    toast: {
      ...props,
      id,
      open: true,
      onOpenChange: (open) => {
        if (!open) dismiss()
      },
    },
  })

  return {
    id: id,
    dismiss,
    update,
  }
}

function useToast() {
  const [state, setState] = React.useState<State>(memoryState)

  React.useEffect(() => {
    listeners.push(setState)
    return () => {
      const index = listeners.indexOf(setState)
      if (index > -1) {
        listeners.splice(index, 1)
      }
    }
  }, [state])

  return {
    ...state,
    toast,
    dismiss: (toastId?: string) => dispatch({ type: "DISMISS_TOAST", toastId }),
  }
}

export { useToast, toast }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/switch.tsx

```tsx
"use client"

import * as React from "react"
import * as SwitchPrimitives from "@radix-ui/react-switch"

import { cn } from "@/lib/utils"

const Switch = React.forwardRef<
  React.ElementRef<typeof SwitchPrimitives.Root>,
  React.ComponentPropsWithoutRef<typeof SwitchPrimitives.Root>
>(({ className, ...props }, ref) => (
  <SwitchPrimitives.Root
    className={cn(
      "peer inline-flex h-6 w-11 shrink-0 cursor-pointer items-center rounded-full border-2 border-transparent transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 focus-visible:ring-offset-background disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=unchecked]:bg-input",
      className
    )}
    {...props}
    ref={ref}
  >
    <SwitchPrimitives.Thumb
      className={cn(
        "pointer-events-none block h-5 w-5 rounded-full bg-background shadow-lg ring-0 transition-transform data-[state=checked]:translate-x-5 data-[state=unchecked]:translate-x-0"
      )}
    />
  </SwitchPrimitives.Root>
))
Switch.displayName = SwitchPrimitives.Root.displayName

export { Switch }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/calendar.tsx

```tsx
"use client"

import * as React from "react"
import { ChevronLeft, ChevronRight } from "lucide-react"
import { DayPicker } from "react-day-picker"

import { cn } from "@/lib/utils"
import { buttonVariants } from "@/components/ui/button"

export type CalendarProps = React.ComponentProps<typeof DayPicker>

function Calendar({
  className,
  classNames,
  showOutsideDays = true,
  ...props
}: CalendarProps) {
  return (
    <DayPicker
      showOutsideDays={showOutsideDays}
      className={cn("p-3", className)}
      classNames={{
        months: "flex flex-col sm:flex-row space-y-4 sm:space-x-4 sm:space-y-0",
        month: "space-y-4",
        caption: "flex justify-center pt-1 relative items-center",
        caption_label: "text-sm font-medium",
        nav: "space-x-1 flex items-center",
        nav_button: cn(
          buttonVariants({ variant: "outline" }),
          "h-7 w-7 bg-transparent p-0 opacity-50 hover:opacity-100"
        ),
        nav_button_previous: "absolute left-1",
        nav_button_next: "absolute right-1",
        table: "w-full border-collapse space-y-1",
        head_row: "flex",
        head_cell:
          "text-muted-foreground rounded-md w-9 font-normal text-[0.8rem]",
        row: "flex w-full mt-2",
        cell: "h-9 w-9 text-center text-sm p-0 relative [&:has([aria-selected].day-range-end)]:rounded-r-md [&:has([aria-selected].day-outside)]:bg-accent/50 [&:has([aria-selected])]:bg-accent first:[&:has([aria-selected])]:rounded-l-md last:[&:has([aria-selected])]:rounded-r-md focus-within:relative focus-within:z-20",
        day: cn(
          buttonVariants({ variant: "ghost" }),
          "h-9 w-9 p-0 font-normal aria-selected:opacity-100"
        ),
        day_range_end: "day-range-end",
        day_selected:
          "bg-primary text-primary-foreground hover:bg-primary hover:text-primary-foreground focus:bg-primary focus:text-primary-foreground",
        day_today: "bg-accent text-accent-foreground",
        day_outside:
          "day-outside text-muted-foreground aria-selected:bg-accent/50 aria-selected:text-muted-foreground",
        day_disabled: "text-muted-foreground opacity-50",
        day_range_middle:
          "aria-selected:bg-accent aria-selected:text-accent-foreground",
        day_hidden: "invisible",
        ...classNames,
      }}
      components={{
        IconLeft: ({ ...props }) => <ChevronLeft className="h-4 w-4" />,
        IconRight: ({ ...props }) => <ChevronRight className="h-4 w-4" />,
      }}
      {...props}
    />
  )
}
Calendar.displayName = "Calendar"

export { Calendar }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/breadcrumb.tsx

```tsx
import * as React from "react"
import { Slot } from "@radix-ui/react-slot"
import { ChevronRight, MoreHorizontal } from "lucide-react"

import { cn } from "@/lib/utils"

const Breadcrumb = React.forwardRef<
  HTMLElement,
  React.ComponentPropsWithoutRef<"nav"> & {
    separator?: React.ReactNode
  }
>(({ ...props }, ref) => <nav ref={ref} aria-label="breadcrumb" {...props} />)
Breadcrumb.displayName = "Breadcrumb"

const BreadcrumbList = React.forwardRef<
  HTMLOListElement,
  React.ComponentPropsWithoutRef<"ol">
>(({ className, ...props }, ref) => (
  <ol
    ref={ref}
    className={cn(
      "flex flex-wrap items-center gap-1.5 break-words text-sm text-muted-foreground sm:gap-2.5",
      className
    )}
    {...props}
  />
))
BreadcrumbList.displayName = "BreadcrumbList"

const BreadcrumbItem = React.forwardRef<
  HTMLLIElement,
  React.ComponentPropsWithoutRef<"li">
>(({ className, ...props }, ref) => (
  <li
    ref={ref}
    className={cn("inline-flex items-center gap-1.5", className)}
    {...props}
  />
))
BreadcrumbItem.displayName = "BreadcrumbItem"

const BreadcrumbLink = React.forwardRef<
  HTMLAnchorElement,
  React.ComponentPropsWithoutRef<"a"> & {
    asChild?: boolean
  }
>(({ asChild, className, ...props }, ref) => {
  const Comp = asChild ? Slot : "a"

  return (
    <Comp
      ref={ref}
      className={cn("transition-colors hover:text-foreground", className)}
      {...props}
    />
  )
})
BreadcrumbLink.displayName = "BreadcrumbLink"

const BreadcrumbPage = React.forwardRef<
  HTMLSpanElement,
  React.ComponentPropsWithoutRef<"span">
>(({ className, ...props }, ref) => (
  <span
    ref={ref}
    role="link"
    aria-disabled="true"
    aria-current="page"
    className={cn("font-normal text-foreground", className)}
    {...props}
  />
))
BreadcrumbPage.displayName = "BreadcrumbPage"

const BreadcrumbSeparator = ({
  children,
  className,
  ...props
}: React.ComponentProps<"li">) => (
  <li
    role="presentation"
    aria-hidden="true"
    className={cn("[&>svg]:w-3.5 [&>svg]:h-3.5", className)}
    {...props}
  >
    {children ?? <ChevronRight />}
  </li>
)
BreadcrumbSeparator.displayName = "BreadcrumbSeparator"

const BreadcrumbEllipsis = ({
  className,
  ...props
}: React.ComponentProps<"span">) => (
  <span
    role="presentation"
    aria-hidden="true"
    className={cn("flex h-9 w-9 items-center justify-center", className)}
    {...props}
  >
    <MoreHorizontal className="h-4 w-4" />
    <span className="sr-only">More</span>
  </span>
)
BreadcrumbEllipsis.displayName = "BreadcrumbElipssis"

export {
  Breadcrumb,
  BreadcrumbList,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbPage,
  BreadcrumbSeparator,
  BreadcrumbEllipsis,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/radio-group.tsx

```tsx
"use client"

import * as React from "react"
import * as RadioGroupPrimitive from "@radix-ui/react-radio-group"
import { Circle } from "lucide-react"

import { cn } from "@/lib/utils"

const RadioGroup = React.forwardRef<
  React.ElementRef<typeof RadioGroupPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof RadioGroupPrimitive.Root>
>(({ className, ...props }, ref) => {
  return (
    <RadioGroupPrimitive.Root
      className={cn("grid gap-2", className)}
      {...props}
      ref={ref}
    />
  )
})
RadioGroup.displayName = RadioGroupPrimitive.Root.displayName

const RadioGroupItem = React.forwardRef<
  React.ElementRef<typeof RadioGroupPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof RadioGroupPrimitive.Item>
>(({ className, ...props }, ref) => {
  return (
    <RadioGroupPrimitive.Item
      ref={ref}
      className={cn(
        "aspect-square h-4 w-4 rounded-full border border-primary text-primary ring-offset-background focus:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50",
        className
      )}
      {...props}
    >
      <RadioGroupPrimitive.Indicator className="flex items-center justify-center">
        <Circle className="h-2.5 w-2.5 fill-current text-current" />
      </RadioGroupPrimitive.Indicator>
    </RadioGroupPrimitive.Item>
  )
})
RadioGroupItem.displayName = RadioGroupPrimitive.Item.displayName

export { RadioGroup, RadioGroupItem }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/command.tsx

```tsx
"use client"

import * as React from "react"
import { type DialogProps } from "@radix-ui/react-dialog"
import { Command as CommandPrimitive } from "cmdk"
import { Search } from "lucide-react"

import { cn } from "@/lib/utils"
import { Dialog, DialogContent } from "@/components/ui/dialog"

const Command = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive>
>(({ className, ...props }, ref) => (
  <CommandPrimitive
    ref={ref}
    className={cn(
      "flex h-full w-full flex-col overflow-hidden rounded-md bg-popover text-popover-foreground",
      className
    )}
    {...props}
  />
))
Command.displayName = CommandPrimitive.displayName

const CommandDialog = ({ children, ...props }: DialogProps) => {
  return (
    <Dialog {...props}>
      <DialogContent className="overflow-hidden p-0 shadow-lg">
        <Command className="[&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:font-medium [&_[cmdk-group-heading]]:text-muted-foreground [&_[cmdk-group]:not([hidden])_~[cmdk-group]]:pt-0 [&_[cmdk-group]]:px-2 [&_[cmdk-input-wrapper]_svg]:h-5 [&_[cmdk-input-wrapper]_svg]:w-5 [&_[cmdk-input]]:h-12 [&_[cmdk-item]]:px-2 [&_[cmdk-item]]:py-3 [&_[cmdk-item]_svg]:h-5 [&_[cmdk-item]_svg]:w-5">
          {children}
        </Command>
      </DialogContent>
    </Dialog>
  )
}

const CommandInput = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Input>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Input>
>(({ className, ...props }, ref) => (
  <div className="flex items-center border-b px-3" cmdk-input-wrapper="">
    <Search className="mr-2 h-4 w-4 shrink-0 opacity-50" />
    <CommandPrimitive.Input
      ref={ref}
      className={cn(
        "flex h-11 w-full rounded-md bg-transparent py-3 text-sm outline-none placeholder:text-muted-foreground disabled:cursor-not-allowed disabled:opacity-50",
        className
      )}
      {...props}
    />
  </div>
))

CommandInput.displayName = CommandPrimitive.Input.displayName

const CommandList = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.List>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.List>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.List
    ref={ref}
    className={cn("max-h-[300px] overflow-y-auto overflow-x-hidden", className)}
    {...props}
  />
))

CommandList.displayName = CommandPrimitive.List.displayName

const CommandEmpty = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Empty>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Empty>
>((props, ref) => (
  <CommandPrimitive.Empty
    ref={ref}
    className="py-6 text-center text-sm"
    {...props}
  />
))

CommandEmpty.displayName = CommandPrimitive.Empty.displayName

const CommandGroup = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Group>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Group>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Group
    ref={ref}
    className={cn(
      "overflow-hidden p-1 text-foreground [&_[cmdk-group-heading]]:px-2 [&_[cmdk-group-heading]]:py-1.5 [&_[cmdk-group-heading]]:text-xs [&_[cmdk-group-heading]]:font-medium [&_[cmdk-group-heading]]:text-muted-foreground",
      className
    )}
    {...props}
  />
))

CommandGroup.displayName = CommandPrimitive.Group.displayName

const CommandSeparator = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Separator>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Separator>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Separator
    ref={ref}
    className={cn("-mx-1 h-px bg-border", className)}
    {...props}
  />
))
CommandSeparator.displayName = CommandPrimitive.Separator.displayName

const CommandItem = React.forwardRef<
  React.ElementRef<typeof CommandPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof CommandPrimitive.Item>
>(({ className, ...props }, ref) => (
  <CommandPrimitive.Item
    ref={ref}
    className={cn(
      "relative flex cursor-default gap-2 select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none data-[disabled=true]:pointer-events-none data-[selected='true']:bg-accent data-[selected=true]:text-accent-foreground data-[disabled=true]:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0",
      className
    )}
    {...props}
  />
))

CommandItem.displayName = CommandPrimitive.Item.displayName

const CommandShortcut = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement>) => {
  return (
    <span
      className={cn(
        "ml-auto text-xs tracking-widest text-muted-foreground",
        className
      )}
      {...props}
    />
  )
}
CommandShortcut.displayName = "CommandShortcut"

export {
  Command,
  CommandDialog,
  CommandInput,
  CommandList,
  CommandEmpty,
  CommandGroup,
  CommandItem,
  CommandShortcut,
  CommandSeparator,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/toggle-group.tsx

```tsx
"use client"

import * as React from "react"
import * as ToggleGroupPrimitive from "@radix-ui/react-toggle-group"
import { type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"
import { toggleVariants } from "@/components/ui/toggle"

const ToggleGroupContext = React.createContext<
  VariantProps<typeof toggleVariants>
>({
  size: "default",
  variant: "default",
})

const ToggleGroup = React.forwardRef<
  React.ElementRef<typeof ToggleGroupPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof ToggleGroupPrimitive.Root> &
    VariantProps<typeof toggleVariants>
>(({ className, variant, size, children, ...props }, ref) => (
  <ToggleGroupPrimitive.Root
    ref={ref}
    className={cn("flex items-center justify-center gap-1", className)}
    {...props}
  >
    <ToggleGroupContext.Provider value={{ variant, size }}>
      {children}
    </ToggleGroupContext.Provider>
  </ToggleGroupPrimitive.Root>
))

ToggleGroup.displayName = ToggleGroupPrimitive.Root.displayName

const ToggleGroupItem = React.forwardRef<
  React.ElementRef<typeof ToggleGroupPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof ToggleGroupPrimitive.Item> &
    VariantProps<typeof toggleVariants>
>(({ className, children, variant, size, ...props }, ref) => {
  const context = React.useContext(ToggleGroupContext)

  return (
    <ToggleGroupPrimitive.Item
      ref={ref}
      className={cn(
        toggleVariants({
          variant: context.variant || variant,
          size: context.size || size,
        }),
        className
      )}
      {...props}
    >
      {children}
    </ToggleGroupPrimitive.Item>
  )
})

ToggleGroupItem.displayName = ToggleGroupPrimitive.Item.displayName

export { ToggleGroup, ToggleGroupItem }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/avatar.tsx

```tsx
"use client"

import * as React from "react"
import * as AvatarPrimitive from "@radix-ui/react-avatar"

import { cn } from "@/lib/utils"

const Avatar = React.forwardRef<
  React.ElementRef<typeof AvatarPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof AvatarPrimitive.Root>
>(({ className, ...props }, ref) => (
  <AvatarPrimitive.Root
    ref={ref}
    className={cn(
      "relative flex h-10 w-10 shrink-0 overflow-hidden rounded-full",
      className
    )}
    {...props}
  />
))
Avatar.displayName = AvatarPrimitive.Root.displayName

const AvatarImage = React.forwardRef<
  React.ElementRef<typeof AvatarPrimitive.Image>,
  React.ComponentPropsWithoutRef<typeof AvatarPrimitive.Image>
>(({ className, ...props }, ref) => (
  <AvatarPrimitive.Image
    ref={ref}
    className={cn("aspect-square h-full w-full", className)}
    {...props}
  />
))
AvatarImage.displayName = AvatarPrimitive.Image.displayName

const AvatarFallback = React.forwardRef<
  React.ElementRef<typeof AvatarPrimitive.Fallback>,
  React.ComponentPropsWithoutRef<typeof AvatarPrimitive.Fallback>
>(({ className, ...props }, ref) => (
  <AvatarPrimitive.Fallback
    ref={ref}
    className={cn(
      "flex h-full w-full items-center justify-center rounded-full bg-muted",
      className
    )}
    {...props}
  />
))
AvatarFallback.displayName = AvatarPrimitive.Fallback.displayName

export { Avatar, AvatarImage, AvatarFallback }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/menubar.tsx

```tsx
"use client"

import * as React from "react"
import * as MenubarPrimitive from "@radix-ui/react-menubar"
import { Check, ChevronRight, Circle } from "lucide-react"

import { cn } from "@/lib/utils"

const MenubarMenu = MenubarPrimitive.Menu

const MenubarGroup = MenubarPrimitive.Group

const MenubarPortal = MenubarPrimitive.Portal

const MenubarSub = MenubarPrimitive.Sub

const MenubarRadioGroup = MenubarPrimitive.RadioGroup

const Menubar = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.Root>
>(({ className, ...props }, ref) => (
  <MenubarPrimitive.Root
    ref={ref}
    className={cn(
      "flex h-10 items-center space-x-1 rounded-md border bg-background p-1",
      className
    )}
    {...props}
  />
))
Menubar.displayName = MenubarPrimitive.Root.displayName

const MenubarTrigger = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.Trigger>
>(({ className, ...props }, ref) => (
  <MenubarPrimitive.Trigger
    ref={ref}
    className={cn(
      "flex cursor-default select-none items-center rounded-sm px-3 py-1.5 text-sm font-medium outline-none focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground",
      className
    )}
    {...props}
  />
))
MenubarTrigger.displayName = MenubarPrimitive.Trigger.displayName

const MenubarSubTrigger = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.SubTrigger>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.SubTrigger> & {
    inset?: boolean
  }
>(({ className, inset, children, ...props }, ref) => (
  <MenubarPrimitive.SubTrigger
    ref={ref}
    className={cn(
      "flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground",
      inset && "pl-8",
      className
    )}
    {...props}
  >
    {children}
    <ChevronRight className="ml-auto h-4 w-4" />
  </MenubarPrimitive.SubTrigger>
))
MenubarSubTrigger.displayName = MenubarPrimitive.SubTrigger.displayName

const MenubarSubContent = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.SubContent>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.SubContent>
>(({ className, ...props }, ref) => (
  <MenubarPrimitive.SubContent
    ref={ref}
    className={cn(
      "z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
      className
    )}
    {...props}
  />
))
MenubarSubContent.displayName = MenubarPrimitive.SubContent.displayName

const MenubarContent = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.Content>
>(
  (
    { className, align = "start", alignOffset = -4, sideOffset = 8, ...props },
    ref
  ) => (
    <MenubarPrimitive.Portal>
      <MenubarPrimitive.Content
        ref={ref}
        align={align}
        alignOffset={alignOffset}
        sideOffset={sideOffset}
        className={cn(
          "z-50 min-w-[12rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
          className
        )}
        {...props}
      />
    </MenubarPrimitive.Portal>
  )
)
MenubarContent.displayName = MenubarPrimitive.Content.displayName

const MenubarItem = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.Item> & {
    inset?: boolean
  }
>(({ className, inset, ...props }, ref) => (
  <MenubarPrimitive.Item
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      inset && "pl-8",
      className
    )}
    {...props}
  />
))
MenubarItem.displayName = MenubarPrimitive.Item.displayName

const MenubarCheckboxItem = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.CheckboxItem>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.CheckboxItem>
>(({ className, children, checked, ...props }, ref) => (
  <MenubarPrimitive.CheckboxItem
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    checked={checked}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <MenubarPrimitive.ItemIndicator>
        <Check className="h-4 w-4" />
      </MenubarPrimitive.ItemIndicator>
    </span>
    {children}
  </MenubarPrimitive.CheckboxItem>
))
MenubarCheckboxItem.displayName = MenubarPrimitive.CheckboxItem.displayName

const MenubarRadioItem = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.RadioItem>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.RadioItem>
>(({ className, children, ...props }, ref) => (
  <MenubarPrimitive.RadioItem
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <MenubarPrimitive.ItemIndicator>
        <Circle className="h-2 w-2 fill-current" />
      </MenubarPrimitive.ItemIndicator>
    </span>
    {children}
  </MenubarPrimitive.RadioItem>
))
MenubarRadioItem.displayName = MenubarPrimitive.RadioItem.displayName

const MenubarLabel = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.Label>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.Label> & {
    inset?: boolean
  }
>(({ className, inset, ...props }, ref) => (
  <MenubarPrimitive.Label
    ref={ref}
    className={cn(
      "px-2 py-1.5 text-sm font-semibold",
      inset && "pl-8",
      className
    )}
    {...props}
  />
))
MenubarLabel.displayName = MenubarPrimitive.Label.displayName

const MenubarSeparator = React.forwardRef<
  React.ElementRef<typeof MenubarPrimitive.Separator>,
  React.ComponentPropsWithoutRef<typeof MenubarPrimitive.Separator>
>(({ className, ...props }, ref) => (
  <MenubarPrimitive.Separator
    ref={ref}
    className={cn("-mx-1 my-1 h-px bg-muted", className)}
    {...props}
  />
))
MenubarSeparator.displayName = MenubarPrimitive.Separator.displayName

const MenubarShortcut = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement>) => {
  return (
    <span
      className={cn(
        "ml-auto text-xs tracking-widest text-muted-foreground",
        className
      )}
      {...props}
    />
  )
}
MenubarShortcut.displayname = "MenubarShortcut"

export {
  Menubar,
  MenubarMenu,
  MenubarTrigger,
  MenubarContent,
  MenubarItem,
  MenubarSeparator,
  MenubarLabel,
  MenubarCheckboxItem,
  MenubarRadioGroup,
  MenubarRadioItem,
  MenubarPortal,
  MenubarSubContent,
  MenubarSubTrigger,
  MenubarGroup,
  MenubarSub,
  MenubarShortcut,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/dialog.tsx

```tsx
"use client"

import * as React from "react"
import * as DialogPrimitive from "@radix-ui/react-dialog"
import { X } from "lucide-react"

import { cn } from "@/lib/utils"

const Dialog = DialogPrimitive.Root

const DialogTrigger = DialogPrimitive.Trigger

const DialogPortal = DialogPrimitive.Portal

const DialogClose = DialogPrimitive.Close

const DialogOverlay = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Overlay>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Overlay>
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Overlay
    ref={ref}
    className={cn(
      "fixed inset-0 z-50 bg-black/80  data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0",
      className
    )}
    {...props}
  />
))
DialogOverlay.displayName = DialogPrimitive.Overlay.displayName

const DialogContent = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Content>
>(({ className, children, ...props }, ref) => (
  <DialogPortal>
    <DialogOverlay />
    <DialogPrimitive.Content
      ref={ref}
      className={cn(
        "fixed left-[50%] top-[50%] z-50 grid w-full max-w-lg translate-x-[-50%] translate-y-[-50%] gap-4 border bg-background p-6 shadow-lg duration-200 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[state=closed]:slide-out-to-left-1/2 data-[state=closed]:slide-out-to-top-[48%] data-[state=open]:slide-in-from-left-1/2 data-[state=open]:slide-in-from-top-[48%] sm:rounded-lg",
        className
      )}
      {...props}
    >
      {children}
      <DialogPrimitive.Close className="absolute right-4 top-4 rounded-sm opacity-70 ring-offset-background transition-opacity hover:opacity-100 focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none data-[state=open]:bg-accent data-[state=open]:text-muted-foreground">
        <X className="h-4 w-4" />
        <span className="sr-only">Close</span>
      </DialogPrimitive.Close>
    </DialogPrimitive.Content>
  </DialogPortal>
))
DialogContent.displayName = DialogPrimitive.Content.displayName

const DialogHeader = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn(
      "flex flex-col space-y-1.5 text-center sm:text-left",
      className
    )}
    {...props}
  />
)
DialogHeader.displayName = "DialogHeader"

const DialogFooter = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) => (
  <div
    className={cn(
      "flex flex-col-reverse sm:flex-row sm:justify-end sm:space-x-2",
      className
    )}
    {...props}
  />
)
DialogFooter.displayName = "DialogFooter"

const DialogTitle = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Title>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Title>
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Title
    ref={ref}
    className={cn(
      "text-lg font-semibold leading-none tracking-tight",
      className
    )}
    {...props}
  />
))
DialogTitle.displayName = DialogPrimitive.Title.displayName

const DialogDescription = React.forwardRef<
  React.ElementRef<typeof DialogPrimitive.Description>,
  React.ComponentPropsWithoutRef<typeof DialogPrimitive.Description>
>(({ className, ...props }, ref) => (
  <DialogPrimitive.Description
    ref={ref}
    className={cn("text-sm text-muted-foreground", className)}
    {...props}
  />
))
DialogDescription.displayName = DialogPrimitive.Description.displayName

export {
  Dialog,
  DialogPortal,
  DialogOverlay,
  DialogClose,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogFooter,
  DialogTitle,
  DialogDescription,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/badge.tsx

```tsx
import * as React from "react"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const badgeVariants = cva(
  "inline-flex items-center rounded-full border px-2.5 py-0.5 text-xs font-semibold transition-colors focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2",
  {
    variants: {
      variant: {
        default:
          "border-transparent bg-primary text-primary-foreground hover:bg-primary/80",
        secondary:
          "border-transparent bg-secondary text-secondary-foreground hover:bg-secondary/80",
        destructive:
          "border-transparent bg-destructive text-destructive-foreground hover:bg-destructive/80",
        outline: "text-foreground",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  }
)

export interface BadgeProps
  extends React.HTMLAttributes<HTMLDivElement>,
    VariantProps<typeof badgeVariants> {}

function Badge({ className, variant, ...props }: BadgeProps) {
  return (
    <div className={cn(badgeVariants({ variant }), className)} {...props} />
  )
}

export { Badge, badgeVariants }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/sidebar.tsx

```tsx
"use client"

import * as React from "react"
import { Slot } from "@radix-ui/react-slot"
import { VariantProps, cva } from "class-variance-authority"
import { PanelLeft } from "lucide-react"

import { useIsMobile } from "@/hooks/use-mobile"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { Separator } from "@/components/ui/separator"
import { Sheet, SheetContent } from "@/components/ui/sheet"
import { Skeleton } from "@/components/ui/skeleton"
import {
  Tooltip,
  TooltipContent,
  TooltipProvider,
  TooltipTrigger,
} from "@/components/ui/tooltip"

const SIDEBAR_COOKIE_NAME = "sidebar:state"
const SIDEBAR_COOKIE_MAX_AGE = 60 * 60 * 24 * 7
const SIDEBAR_WIDTH = "16rem"
const SIDEBAR_WIDTH_MOBILE = "18rem"
const SIDEBAR_WIDTH_ICON = "3rem"
const SIDEBAR_KEYBOARD_SHORTCUT = "b"

type SidebarContext = {
  state: "expanded" | "collapsed"
  open: boolean
  setOpen: (open: boolean) => void
  openMobile: boolean
  setOpenMobile: (open: boolean) => void
  isMobile: boolean
  toggleSidebar: () => void
}

const SidebarContext = React.createContext<SidebarContext | null>(null)

function useSidebar() {
  const context = React.useContext(SidebarContext)
  if (!context) {
    throw new Error("useSidebar must be used within a SidebarProvider.")
  }

  return context
}

const SidebarProvider = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div"> & {
    defaultOpen?: boolean
    open?: boolean
    onOpenChange?: (open: boolean) => void
  }
>(
  (
    {
      defaultOpen = true,
      open: openProp,
      onOpenChange: setOpenProp,
      className,
      style,
      children,
      ...props
    },
    ref
  ) => {
    const isMobile = useIsMobile()
    const [openMobile, setOpenMobile] = React.useState(false)

    // This is the internal state of the sidebar.
    // We use openProp and setOpenProp for control from outside the component.
    const [_open, _setOpen] = React.useState(defaultOpen)
    const open = openProp ?? _open
    const setOpen = React.useCallback(
      (value: boolean | ((value: boolean) => boolean)) => {
        const openState = typeof value === "function" ? value(open) : value
        if (setOpenProp) {
          setOpenProp(openState)
        } else {
          _setOpen(openState)
        }

        // This sets the cookie to keep the sidebar state.
        document.cookie = `${SIDEBAR_COOKIE_NAME}=${openState}; path=/; max-age=${SIDEBAR_COOKIE_MAX_AGE}`
      },
      [setOpenProp, open]
    )

    // Helper to toggle the sidebar.
    const toggleSidebar = React.useCallback(() => {
      return isMobile
        ? setOpenMobile((open) => !open)
        : setOpen((open) => !open)
    }, [isMobile, setOpen, setOpenMobile])

    // Adds a keyboard shortcut to toggle the sidebar.
    React.useEffect(() => {
      const handleKeyDown = (event: KeyboardEvent) => {
        if (
          event.key === SIDEBAR_KEYBOARD_SHORTCUT &&
          (event.metaKey || event.ctrlKey)
        ) {
          event.preventDefault()
          toggleSidebar()
        }
      }

      window.addEventListener("keydown", handleKeyDown)
      return () => window.removeEventListener("keydown", handleKeyDown)
    }, [toggleSidebar])

    // We add a state so that we can do data-state="expanded" or "collapsed".
    // This makes it easier to style the sidebar with Tailwind classes.
    const state = open ? "expanded" : "collapsed"

    const contextValue = React.useMemo<SidebarContext>(
      () => ({
        state,
        open,
        setOpen,
        isMobile,
        openMobile,
        setOpenMobile,
        toggleSidebar,
      }),
      [state, open, setOpen, isMobile, openMobile, setOpenMobile, toggleSidebar]
    )

    return (
      <SidebarContext.Provider value={contextValue}>
        <TooltipProvider delayDuration={0}>
          <div
            style={
              {
                "--sidebar-width": SIDEBAR_WIDTH,
                "--sidebar-width-icon": SIDEBAR_WIDTH_ICON,
                ...style,
              } as React.CSSProperties
            }
            className={cn(
              "group/sidebar-wrapper flex min-h-svh w-full has-[[data-variant=inset]]:bg-sidebar",
              className
            )}
            ref={ref}
            {...props}
          >
            {children}
          </div>
        </TooltipProvider>
      </SidebarContext.Provider>
    )
  }
)
SidebarProvider.displayName = "SidebarProvider"

const Sidebar = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div"> & {
    side?: "left" | "right"
    variant?: "sidebar" | "floating" | "inset"
    collapsible?: "offcanvas" | "icon" | "none"
  }
>(
  (
    {
      side = "left",
      variant = "sidebar",
      collapsible = "offcanvas",
      className,
      children,
      ...props
    },
    ref
  ) => {
    const { isMobile, state, openMobile, setOpenMobile } = useSidebar()

    if (collapsible === "none") {
      return (
        <div
          className={cn(
            "flex h-full w-[--sidebar-width] flex-col bg-sidebar text-sidebar-foreground",
            className
          )}
          ref={ref}
          {...props}
        >
          {children}
        </div>
      )
    }

    if (isMobile) {
      return (
        <Sheet open={openMobile} onOpenChange={setOpenMobile} {...props}>
          <SheetContent
            data-sidebar="sidebar"
            data-mobile="true"
            className="w-[--sidebar-width] bg-sidebar p-0 text-sidebar-foreground [&>button]:hidden"
            style={
              {
                "--sidebar-width": SIDEBAR_WIDTH_MOBILE,
              } as React.CSSProperties
            }
            side={side}
          >
            <div className="flex h-full w-full flex-col">{children}</div>
          </SheetContent>
        </Sheet>
      )
    }

    return (
      <div
        ref={ref}
        className="group peer hidden md:block text-sidebar-foreground"
        data-state={state}
        data-collapsible={state === "collapsed" ? collapsible : ""}
        data-variant={variant}
        data-side={side}
      >
        {/* This is what handles the sidebar gap on desktop */}
        <div
          className={cn(
            "duration-200 relative h-svh w-[--sidebar-width] bg-transparent transition-[width] ease-linear",
            "group-data-[collapsible=offcanvas]:w-0",
            "group-data-[side=right]:rotate-180",
            variant === "floating" || variant === "inset"
              ? "group-data-[collapsible=icon]:w-[calc(var(--sidebar-width-icon)_+_theme(spacing.4))]"
              : "group-data-[collapsible=icon]:w-[--sidebar-width-icon]"
          )}
        />
        <div
          className={cn(
            "duration-200 fixed inset-y-0 z-10 hidden h-svh w-[--sidebar-width] transition-[left,right,width] ease-linear md:flex",
            side === "left"
              ? "left-0 group-data-[collapsible=offcanvas]:left-[calc(var(--sidebar-width)*-1)]"
              : "right-0 group-data-[collapsible=offcanvas]:right-[calc(var(--sidebar-width)*-1)]",
            // Adjust the padding for floating and inset variants.
            variant === "floating" || variant === "inset"
              ? "p-2 group-data-[collapsible=icon]:w-[calc(var(--sidebar-width-icon)_+_theme(spacing.4)_+2px)]"
              : "group-data-[collapsible=icon]:w-[--sidebar-width-icon] group-data-[side=left]:border-r group-data-[side=right]:border-l",
            className
          )}
          {...props}
        >
          <div
            data-sidebar="sidebar"
            className="flex h-full w-full flex-col bg-sidebar group-data-[variant=floating]:rounded-lg group-data-[variant=floating]:border group-data-[variant=floating]:border-sidebar-border group-data-[variant=floating]:shadow"
          >
            {children}
          </div>
        </div>
      </div>
    )
  }
)
Sidebar.displayName = "Sidebar"

const SidebarTrigger = React.forwardRef<
  React.ElementRef<typeof Button>,
  React.ComponentProps<typeof Button>
>(({ className, onClick, ...props }, ref) => {
  const { toggleSidebar } = useSidebar()

  return (
    <Button
      ref={ref}
      data-sidebar="trigger"
      variant="ghost"
      size="icon"
      className={cn("h-7 w-7", className)}
      onClick={(event) => {
        onClick?.(event)
        toggleSidebar()
      }}
      {...props}
    >
      <PanelLeft />
      <span className="sr-only">Toggle Sidebar</span>
    </Button>
  )
})
SidebarTrigger.displayName = "SidebarTrigger"

const SidebarRail = React.forwardRef<
  HTMLButtonElement,
  React.ComponentProps<"button">
>(({ className, ...props }, ref) => {
  const { toggleSidebar } = useSidebar()

  return (
    <button
      ref={ref}
      data-sidebar="rail"
      aria-label="Toggle Sidebar"
      tabIndex={-1}
      onClick={toggleSidebar}
      title="Toggle Sidebar"
      className={cn(
        "absolute inset-y-0 z-20 hidden w-4 -translate-x-1/2 transition-all ease-linear after:absolute after:inset-y-0 after:left-1/2 after:w-[2px] hover:after:bg-sidebar-border group-data-[side=left]:-right-4 group-data-[side=right]:left-0 sm:flex",
        "[[data-side=left]_&]:cursor-w-resize [[data-side=right]_&]:cursor-e-resize",
        "[[data-side=left][data-state=collapsed]_&]:cursor-e-resize [[data-side=right][data-state=collapsed]_&]:cursor-w-resize",
        "group-data-[collapsible=offcanvas]:translate-x-0 group-data-[collapsible=offcanvas]:after:left-full group-data-[collapsible=offcanvas]:hover:bg-sidebar",
        "[[data-side=left][data-collapsible=offcanvas]_&]:-right-2",
        "[[data-side=right][data-collapsible=offcanvas]_&]:-left-2",
        className
      )}
      {...props}
    />
  )
})
SidebarRail.displayName = "SidebarRail"

const SidebarInset = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"main">
>(({ className, ...props }, ref) => {
  return (
    <main
      ref={ref}
      className={cn(
        "relative flex min-h-svh flex-1 flex-col bg-background",
        "peer-data-[variant=inset]:min-h-[calc(100svh-theme(spacing.4))] md:peer-data-[variant=inset]:m-2 md:peer-data-[state=collapsed]:peer-data-[variant=inset]:ml-2 md:peer-data-[variant=inset]:ml-0 md:peer-data-[variant=inset]:rounded-xl md:peer-data-[variant=inset]:shadow",
        className
      )}
      {...props}
    />
  )
})
SidebarInset.displayName = "SidebarInset"

const SidebarInput = React.forwardRef<
  React.ElementRef<typeof Input>,
  React.ComponentProps<typeof Input>
>(({ className, ...props }, ref) => {
  return (
    <Input
      ref={ref}
      data-sidebar="input"
      className={cn(
        "h-8 w-full bg-background shadow-none focus-visible:ring-2 focus-visible:ring-sidebar-ring",
        className
      )}
      {...props}
    />
  )
})
SidebarInput.displayName = "SidebarInput"

const SidebarHeader = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div">
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-sidebar="header"
      className={cn("flex flex-col gap-2 p-2", className)}
      {...props}
    />
  )
})
SidebarHeader.displayName = "SidebarHeader"

const SidebarFooter = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div">
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-sidebar="footer"
      className={cn("flex flex-col gap-2 p-2", className)}
      {...props}
    />
  )
})
SidebarFooter.displayName = "SidebarFooter"

const SidebarSeparator = React.forwardRef<
  React.ElementRef<typeof Separator>,
  React.ComponentProps<typeof Separator>
>(({ className, ...props }, ref) => {
  return (
    <Separator
      ref={ref}
      data-sidebar="separator"
      className={cn("mx-2 w-auto bg-sidebar-border", className)}
      {...props}
    />
  )
})
SidebarSeparator.displayName = "SidebarSeparator"

const SidebarContent = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div">
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-sidebar="content"
      className={cn(
        "flex min-h-0 flex-1 flex-col gap-2 overflow-auto group-data-[collapsible=icon]:overflow-hidden",
        className
      )}
      {...props}
    />
  )
})
SidebarContent.displayName = "SidebarContent"

const SidebarGroup = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div">
>(({ className, ...props }, ref) => {
  return (
    <div
      ref={ref}
      data-sidebar="group"
      className={cn("relative flex w-full min-w-0 flex-col p-2", className)}
      {...props}
    />
  )
})
SidebarGroup.displayName = "SidebarGroup"

const SidebarGroupLabel = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div"> & { asChild?: boolean }
>(({ className, asChild = false, ...props }, ref) => {
  const Comp = asChild ? Slot : "div"

  return (
    <Comp
      ref={ref}
      data-sidebar="group-label"
      className={cn(
        "duration-200 flex h-8 shrink-0 items-center rounded-md px-2 text-xs font-medium text-sidebar-foreground/70 outline-none ring-sidebar-ring transition-[margin,opa] ease-linear focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0",
        "group-data-[collapsible=icon]:-mt-8 group-data-[collapsible=icon]:opacity-0",
        className
      )}
      {...props}
    />
  )
})
SidebarGroupLabel.displayName = "SidebarGroupLabel"

const SidebarGroupAction = React.forwardRef<
  HTMLButtonElement,
  React.ComponentProps<"button"> & { asChild?: boolean }
>(({ className, asChild = false, ...props }, ref) => {
  const Comp = asChild ? Slot : "button"

  return (
    <Comp
      ref={ref}
      data-sidebar="group-action"
      className={cn(
        "absolute right-3 top-3.5 flex aspect-square w-5 items-center justify-center rounded-md p-0 text-sidebar-foreground outline-none ring-sidebar-ring transition-transform hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 [&>svg]:size-4 [&>svg]:shrink-0",
        // Increases the hit area of the button on mobile.
        "after:absolute after:-inset-2 after:md:hidden",
        "group-data-[collapsible=icon]:hidden",
        className
      )}
      {...props}
    />
  )
})
SidebarGroupAction.displayName = "SidebarGroupAction"

const SidebarGroupContent = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div">
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    data-sidebar="group-content"
    className={cn("w-full text-sm", className)}
    {...props}
  />
))
SidebarGroupContent.displayName = "SidebarGroupContent"

const SidebarMenu = React.forwardRef<
  HTMLUListElement,
  React.ComponentProps<"ul">
>(({ className, ...props }, ref) => (
  <ul
    ref={ref}
    data-sidebar="menu"
    className={cn("flex w-full min-w-0 flex-col gap-1", className)}
    {...props}
  />
))
SidebarMenu.displayName = "SidebarMenu"

const SidebarMenuItem = React.forwardRef<
  HTMLLIElement,
  React.ComponentProps<"li">
>(({ className, ...props }, ref) => (
  <li
    ref={ref}
    data-sidebar="menu-item"
    className={cn("group/menu-item relative", className)}
    {...props}
  />
))
SidebarMenuItem.displayName = "SidebarMenuItem"

const sidebarMenuButtonVariants = cva(
  "peer/menu-button flex w-full items-center gap-2 overflow-hidden rounded-md p-2 text-left text-sm outline-none ring-sidebar-ring transition-[width,height,padding] hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 active:bg-sidebar-accent active:text-sidebar-accent-foreground disabled:pointer-events-none disabled:opacity-50 group-has-[[data-sidebar=menu-action]]/menu-item:pr-8 aria-disabled:pointer-events-none aria-disabled:opacity-50 data-[active=true]:bg-sidebar-accent data-[active=true]:font-medium data-[active=true]:text-sidebar-accent-foreground data-[state=open]:hover:bg-sidebar-accent data-[state=open]:hover:text-sidebar-accent-foreground group-data-[collapsible=icon]:!size-8 group-data-[collapsible=icon]:!p-2 [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0",
  {
    variants: {
      variant: {
        default: "hover:bg-sidebar-accent hover:text-sidebar-accent-foreground",
        outline:
          "bg-background shadow-[0_0_0_1px_hsl(var(--sidebar-border))] hover:bg-sidebar-accent hover:text-sidebar-accent-foreground hover:shadow-[0_0_0_1px_hsl(var(--sidebar-accent))]",
      },
      size: {
        default: "h-8 text-sm",
        sm: "h-7 text-xs",
        lg: "h-12 text-sm group-data-[collapsible=icon]:!p-0",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
)

const SidebarMenuButton = React.forwardRef<
  HTMLButtonElement,
  React.ComponentProps<"button"> & {
    asChild?: boolean
    isActive?: boolean
    tooltip?: string | React.ComponentProps<typeof TooltipContent>
  } & VariantProps<typeof sidebarMenuButtonVariants>
>(
  (
    {
      asChild = false,
      isActive = false,
      variant = "default",
      size = "default",
      tooltip,
      className,
      ...props
    },
    ref
  ) => {
    const Comp = asChild ? Slot : "button"
    const { isMobile, state } = useSidebar()

    const button = (
      <Comp
        ref={ref}
        data-sidebar="menu-button"
        data-size={size}
        data-active={isActive}
        className={cn(sidebarMenuButtonVariants({ variant, size }), className)}
        {...props}
      />
    )

    if (!tooltip) {
      return button
    }

    if (typeof tooltip === "string") {
      tooltip = {
        children: tooltip,
      }
    }

    return (
      <Tooltip>
        <TooltipTrigger asChild>{button}</TooltipTrigger>
        <TooltipContent
          side="right"
          align="center"
          hidden={state !== "collapsed" || isMobile}
          {...tooltip}
        />
      </Tooltip>
    )
  }
)
SidebarMenuButton.displayName = "SidebarMenuButton"

const SidebarMenuAction = React.forwardRef<
  HTMLButtonElement,
  React.ComponentProps<"button"> & {
    asChild?: boolean
    showOnHover?: boolean
  }
>(({ className, asChild = false, showOnHover = false, ...props }, ref) => {
  const Comp = asChild ? Slot : "button"

  return (
    <Comp
      ref={ref}
      data-sidebar="menu-action"
      className={cn(
        "absolute right-1 top-1.5 flex aspect-square w-5 items-center justify-center rounded-md p-0 text-sidebar-foreground outline-none ring-sidebar-ring transition-transform hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 peer-hover/menu-button:text-sidebar-accent-foreground [&>svg]:size-4 [&>svg]:shrink-0",
        // Increases the hit area of the button on mobile.
        "after:absolute after:-inset-2 after:md:hidden",
        "peer-data-[size=sm]/menu-button:top-1",
        "peer-data-[size=default]/menu-button:top-1.5",
        "peer-data-[size=lg]/menu-button:top-2.5",
        "group-data-[collapsible=icon]:hidden",
        showOnHover &&
          "group-focus-within/menu-item:opacity-100 group-hover/menu-item:opacity-100 data-[state=open]:opacity-100 peer-data-[active=true]/menu-button:text-sidebar-accent-foreground md:opacity-0",
        className
      )}
      {...props}
    />
  )
})
SidebarMenuAction.displayName = "SidebarMenuAction"

const SidebarMenuBadge = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div">
>(({ className, ...props }, ref) => (
  <div
    ref={ref}
    data-sidebar="menu-badge"
    className={cn(
      "absolute right-1 flex h-5 min-w-5 items-center justify-center rounded-md px-1 text-xs font-medium tabular-nums text-sidebar-foreground select-none pointer-events-none",
      "peer-hover/menu-button:text-sidebar-accent-foreground peer-data-[active=true]/menu-button:text-sidebar-accent-foreground",
      "peer-data-[size=sm]/menu-button:top-1",
      "peer-data-[size=default]/menu-button:top-1.5",
      "peer-data-[size=lg]/menu-button:top-2.5",
      "group-data-[collapsible=icon]:hidden",
      className
    )}
    {...props}
  />
))
SidebarMenuBadge.displayName = "SidebarMenuBadge"

const SidebarMenuSkeleton = React.forwardRef<
  HTMLDivElement,
  React.ComponentProps<"div"> & {
    showIcon?: boolean
  }
>(({ className, showIcon = false, ...props }, ref) => {
  // Random width between 50 to 90%.
  const width = React.useMemo(() => {
    return `${Math.floor(Math.random() * 40) + 50}%`
  }, [])

  return (
    <div
      ref={ref}
      data-sidebar="menu-skeleton"
      className={cn("rounded-md h-8 flex gap-2 px-2 items-center", className)}
      {...props}
    >
      {showIcon && (
        <Skeleton
          className="size-4 rounded-md"
          data-sidebar="menu-skeleton-icon"
        />
      )}
      <Skeleton
        className="h-4 flex-1 max-w-[--skeleton-width]"
        data-sidebar="menu-skeleton-text"
        style={
          {
            "--skeleton-width": width,
          } as React.CSSProperties
        }
      />
    </div>
  )
})
SidebarMenuSkeleton.displayName = "SidebarMenuSkeleton"

const SidebarMenuSub = React.forwardRef<
  HTMLUListElement,
  React.ComponentProps<"ul">
>(({ className, ...props }, ref) => (
  <ul
    ref={ref}
    data-sidebar="menu-sub"
    className={cn(
      "mx-3.5 flex min-w-0 translate-x-px flex-col gap-1 border-l border-sidebar-border px-2.5 py-0.5",
      "group-data-[collapsible=icon]:hidden",
      className
    )}
    {...props}
  />
))
SidebarMenuSub.displayName = "SidebarMenuSub"

const SidebarMenuSubItem = React.forwardRef<
  HTMLLIElement,
  React.ComponentProps<"li">
>(({ ...props }, ref) => <li ref={ref} {...props} />)
SidebarMenuSubItem.displayName = "SidebarMenuSubItem"

const SidebarMenuSubButton = React.forwardRef<
  HTMLAnchorElement,
  React.ComponentProps<"a"> & {
    asChild?: boolean
    size?: "sm" | "md"
    isActive?: boolean
  }
>(({ asChild = false, size = "md", isActive, className, ...props }, ref) => {
  const Comp = asChild ? Slot : "a"

  return (
    <Comp
      ref={ref}
      data-sidebar="menu-sub-button"
      data-size={size}
      data-active={isActive}
      className={cn(
        "flex h-7 min-w-0 -translate-x-px items-center gap-2 overflow-hidden rounded-md px-2 text-sidebar-foreground outline-none ring-sidebar-ring hover:bg-sidebar-accent hover:text-sidebar-accent-foreground focus-visible:ring-2 active:bg-sidebar-accent active:text-sidebar-accent-foreground disabled:pointer-events-none disabled:opacity-50 aria-disabled:pointer-events-none aria-disabled:opacity-50 [&>span:last-child]:truncate [&>svg]:size-4 [&>svg]:shrink-0 [&>svg]:text-sidebar-accent-foreground",
        "data-[active=true]:bg-sidebar-accent data-[active=true]:text-sidebar-accent-foreground",
        size === "sm" && "text-xs",
        size === "md" && "text-sm",
        "group-data-[collapsible=icon]:hidden",
        className
      )}
      {...props}
    />
  )
})
SidebarMenuSubButton.displayName = "SidebarMenuSubButton"

export {
  Sidebar,
  SidebarContent,
  SidebarFooter,
  SidebarGroup,
  SidebarGroupAction,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarHeader,
  SidebarInput,
  SidebarInset,
  SidebarMenu,
  SidebarMenuAction,
  SidebarMenuBadge,
  SidebarMenuButton,
  SidebarMenuItem,
  SidebarMenuSkeleton,
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  SidebarProvider,
  SidebarRail,
  SidebarSeparator,
  SidebarTrigger,
  useSidebar,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/table.tsx

```tsx
import * as React from "react"

import { cn } from "@/lib/utils"

const Table = React.forwardRef<
  HTMLTableElement,
  React.HTMLAttributes<HTMLTableElement>
>(({ className, ...props }, ref) => (
  <div className="relative w-full overflow-auto">
    <table
      ref={ref}
      className={cn("w-full caption-bottom text-sm", className)}
      {...props}
    />
  </div>
))
Table.displayName = "Table"

const TableHeader = React.forwardRef<
  HTMLTableSectionElement,
  React.HTMLAttributes<HTMLTableSectionElement>
>(({ className, ...props }, ref) => (
  <thead ref={ref} className={cn("[&_tr]:border-b", className)} {...props} />
))
TableHeader.displayName = "TableHeader"

const TableBody = React.forwardRef<
  HTMLTableSectionElement,
  React.HTMLAttributes<HTMLTableSectionElement>
>(({ className, ...props }, ref) => (
  <tbody
    ref={ref}
    className={cn("[&_tr:last-child]:border-0", className)}
    {...props}
  />
))
TableBody.displayName = "TableBody"

const TableFooter = React.forwardRef<
  HTMLTableSectionElement,
  React.HTMLAttributes<HTMLTableSectionElement>
>(({ className, ...props }, ref) => (
  <tfoot
    ref={ref}
    className={cn(
      "border-t bg-muted/50 font-medium [&>tr]:last:border-b-0",
      className
    )}
    {...props}
  />
))
TableFooter.displayName = "TableFooter"

const TableRow = React.forwardRef<
  HTMLTableRowElement,
  React.HTMLAttributes<HTMLTableRowElement>
>(({ className, ...props }, ref) => (
  <tr
    ref={ref}
    className={cn(
      "border-b transition-colors hover:bg-muted/50 data-[state=selected]:bg-muted",
      className
    )}
    {...props}
  />
))
TableRow.displayName = "TableRow"

const TableHead = React.forwardRef<
  HTMLTableCellElement,
  React.ThHTMLAttributes<HTMLTableCellElement>
>(({ className, ...props }, ref) => (
  <th
    ref={ref}
    className={cn(
      "h-12 px-4 text-left align-middle font-medium text-muted-foreground [&:has([role=checkbox])]:pr-0",
      className
    )}
    {...props}
  />
))
TableHead.displayName = "TableHead"

const TableCell = React.forwardRef<
  HTMLTableCellElement,
  React.TdHTMLAttributes<HTMLTableCellElement>
>(({ className, ...props }, ref) => (
  <td
    ref={ref}
    className={cn("p-4 align-middle [&:has([role=checkbox])]:pr-0", className)}
    {...props}
  />
))
TableCell.displayName = "TableCell"

const TableCaption = React.forwardRef<
  HTMLTableCaptionElement,
  React.HTMLAttributes<HTMLTableCaptionElement>
>(({ className, ...props }, ref) => (
  <caption
    ref={ref}
    className={cn("mt-4 text-sm text-muted-foreground", className)}
    {...props}
  />
))
TableCaption.displayName = "TableCaption"

export {
  Table,
  TableHeader,
  TableBody,
  TableFooter,
  TableHead,
  TableRow,
  TableCell,
  TableCaption,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/separator.tsx

```tsx
"use client"

import * as React from "react"
import * as SeparatorPrimitive from "@radix-ui/react-separator"

import { cn } from "@/lib/utils"

const Separator = React.forwardRef<
  React.ElementRef<typeof SeparatorPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof SeparatorPrimitive.Root>
>(
  (
    { className, orientation = "horizontal", decorative = true, ...props },
    ref
  ) => (
    <SeparatorPrimitive.Root
      ref={ref}
      decorative={decorative}
      orientation={orientation}
      className={cn(
        "shrink-0 bg-border",
        orientation === "horizontal" ? "h-[1px] w-full" : "h-full w-[1px]",
        className
      )}
      {...props}
    />
  )
)
Separator.displayName = SeparatorPrimitive.Root.displayName

export { Separator }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/button.tsx

```tsx
import * as React from "react"
import { Slot } from "@radix-ui/react-slot"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const buttonVariants = cva(
  "inline-flex items-center justify-center gap-2 whitespace-nowrap rounded-md text-sm font-medium ring-offset-background transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0",
  {
    variants: {
      variant: {
        default: "bg-primary text-primary-foreground hover:bg-primary/90",
        destructive:
          "bg-destructive text-destructive-foreground hover:bg-destructive/90",
        outline:
          "border border-input bg-background hover:bg-accent hover:text-accent-foreground",
        secondary:
          "bg-secondary text-secondary-foreground hover:bg-secondary/80",
        ghost: "hover:bg-accent hover:text-accent-foreground",
        link: "text-primary underline-offset-4 hover:underline",
      },
      size: {
        default: "h-10 px-4 py-2",
        sm: "h-9 rounded-md px-3",
        lg: "h-11 rounded-md px-8",
        icon: "h-10 w-10",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
)

export interface ButtonProps
  extends React.ButtonHTMLAttributes<HTMLButtonElement>,
    VariantProps<typeof buttonVariants> {
  asChild?: boolean
}

const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ className, variant, size, asChild = false, ...props }, ref) => {
    const Comp = asChild ? Slot : "button"
    return (
      <Comp
        className={cn(buttonVariants({ variant, size, className }))}
        ref={ref}
        {...props}
      />
    )
  }
)
Button.displayName = "Button"

export { Button, buttonVariants }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/toggle.tsx

```tsx
"use client"

import * as React from "react"
import * as TogglePrimitive from "@radix-ui/react-toggle"
import { cva, type VariantProps } from "class-variance-authority"

import { cn } from "@/lib/utils"

const toggleVariants = cva(
  "inline-flex items-center justify-center rounded-md text-sm font-medium ring-offset-background transition-colors hover:bg-muted hover:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 data-[state=on]:bg-accent data-[state=on]:text-accent-foreground [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0 gap-2",
  {
    variants: {
      variant: {
        default: "bg-transparent",
        outline:
          "border border-input bg-transparent hover:bg-accent hover:text-accent-foreground",
      },
      size: {
        default: "h-10 px-3 min-w-10",
        sm: "h-9 px-2.5 min-w-9",
        lg: "h-11 px-5 min-w-11",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "default",
    },
  }
)

const Toggle = React.forwardRef<
  React.ElementRef<typeof TogglePrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof TogglePrimitive.Root> &
    VariantProps<typeof toggleVariants>
>(({ className, variant, size, ...props }, ref) => (
  <TogglePrimitive.Root
    ref={ref}
    className={cn(toggleVariants({ variant, size, className }))}
    {...props}
  />
))

Toggle.displayName = TogglePrimitive.Root.displayName

export { Toggle, toggleVariants }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/toast.tsx

```tsx
"use client"

import * as React from "react"
import * as ToastPrimitives from "@radix-ui/react-toast"
import { cva, type VariantProps } from "class-variance-authority"
import { X } from "lucide-react"

import { cn } from "@/lib/utils"

const ToastProvider = ToastPrimitives.Provider

const ToastViewport = React.forwardRef<
  React.ElementRef<typeof ToastPrimitives.Viewport>,
  React.ComponentPropsWithoutRef<typeof ToastPrimitives.Viewport>
>(({ className, ...props }, ref) => (
  <ToastPrimitives.Viewport
    ref={ref}
    className={cn(
      "fixed top-0 z-[100] flex max-h-screen w-full flex-col-reverse p-4 sm:bottom-0 sm:right-0 sm:top-auto sm:flex-col md:max-w-[420px]",
      className
    )}
    {...props}
  />
))
ToastViewport.displayName = ToastPrimitives.Viewport.displayName

const toastVariants = cva(
  "group pointer-events-auto relative flex w-full items-center justify-between space-x-4 overflow-hidden rounded-md border p-6 pr-8 shadow-lg transition-all data-[swipe=cancel]:translate-x-0 data-[swipe=end]:translate-x-[var(--radix-toast-swipe-end-x)] data-[swipe=move]:translate-x-[var(--radix-toast-swipe-move-x)] data-[swipe=move]:transition-none data-[state=open]:animate-in data-[state=closed]:animate-out data-[swipe=end]:animate-out data-[state=closed]:fade-out-80 data-[state=closed]:slide-out-to-right-full data-[state=open]:slide-in-from-top-full data-[state=open]:sm:slide-in-from-bottom-full",
  {
    variants: {
      variant: {
        default: "border bg-background text-foreground",
        destructive:
          "destructive group border-destructive bg-destructive text-destructive-foreground",
      },
    },
    defaultVariants: {
      variant: "default",
    },
  }
)

const Toast = React.forwardRef<
  React.ElementRef<typeof ToastPrimitives.Root>,
  React.ComponentPropsWithoutRef<typeof ToastPrimitives.Root> &
    VariantProps<typeof toastVariants>
>(({ className, variant, ...props }, ref) => {
  return (
    <ToastPrimitives.Root
      ref={ref}
      className={cn(toastVariants({ variant }), className)}
      {...props}
    />
  )
})
Toast.displayName = ToastPrimitives.Root.displayName

const ToastAction = React.forwardRef<
  React.ElementRef<typeof ToastPrimitives.Action>,
  React.ComponentPropsWithoutRef<typeof ToastPrimitives.Action>
>(({ className, ...props }, ref) => (
  <ToastPrimitives.Action
    ref={ref}
    className={cn(
      "inline-flex h-8 shrink-0 items-center justify-center rounded-md border bg-transparent px-3 text-sm font-medium ring-offset-background transition-colors hover:bg-secondary focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:pointer-events-none disabled:opacity-50 group-[.destructive]:border-muted/40 group-[.destructive]:hover:border-destructive/30 group-[.destructive]:hover:bg-destructive group-[.destructive]:hover:text-destructive-foreground group-[.destructive]:focus:ring-destructive",
      className
    )}
    {...props}
  />
))
ToastAction.displayName = ToastPrimitives.Action.displayName

const ToastClose = React.forwardRef<
  React.ElementRef<typeof ToastPrimitives.Close>,
  React.ComponentPropsWithoutRef<typeof ToastPrimitives.Close>
>(({ className, ...props }, ref) => (
  <ToastPrimitives.Close
    ref={ref}
    className={cn(
      "absolute right-2 top-2 rounded-md p-1 text-foreground/50 opacity-0 transition-opacity hover:text-foreground focus:opacity-100 focus:outline-none focus:ring-2 group-hover:opacity-100 group-[.destructive]:text-red-300 group-[.destructive]:hover:text-red-50 group-[.destructive]:focus:ring-red-400 group-[.destructive]:focus:ring-offset-red-600",
      className
    )}
    toast-close=""
    {...props}
  >
    <X className="h-4 w-4" />
  </ToastPrimitives.Close>
))
ToastClose.displayName = ToastPrimitives.Close.displayName

const ToastTitle = React.forwardRef<
  React.ElementRef<typeof ToastPrimitives.Title>,
  React.ComponentPropsWithoutRef<typeof ToastPrimitives.Title>
>(({ className, ...props }, ref) => (
  <ToastPrimitives.Title
    ref={ref}
    className={cn("text-sm font-semibold", className)}
    {...props}
  />
))
ToastTitle.displayName = ToastPrimitives.Title.displayName

const ToastDescription = React.forwardRef<
  React.ElementRef<typeof ToastPrimitives.Description>,
  React.ComponentPropsWithoutRef<typeof ToastPrimitives.Description>
>(({ className, ...props }, ref) => (
  <ToastPrimitives.Description
    ref={ref}
    className={cn("text-sm opacity-90", className)}
    {...props}
  />
))
ToastDescription.displayName = ToastPrimitives.Description.displayName

type ToastProps = React.ComponentPropsWithoutRef<typeof Toast>

type ToastActionElement = React.ReactElement<typeof ToastAction>

export {
  type ToastProps,
  type ToastActionElement,
  ToastProvider,
  ToastViewport,
  Toast,
  ToastTitle,
  ToastDescription,
  ToastClose,
  ToastAction,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/checkbox.tsx

```tsx
"use client"

import * as React from "react"
import * as CheckboxPrimitive from "@radix-ui/react-checkbox"
import { Check } from "lucide-react"

import { cn } from "@/lib/utils"

const Checkbox = React.forwardRef<
  React.ElementRef<typeof CheckboxPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof CheckboxPrimitive.Root>
>(({ className, ...props }, ref) => (
  <CheckboxPrimitive.Root
    ref={ref}
    className={cn(
      "peer h-4 w-4 shrink-0 rounded-sm border border-primary ring-offset-background focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 data-[state=checked]:bg-primary data-[state=checked]:text-primary-foreground",
      className
    )}
    {...props}
  >
    <CheckboxPrimitive.Indicator
      className={cn("flex items-center justify-center text-current")}
    >
      <Check className="h-4 w-4" />
    </CheckboxPrimitive.Indicator>
  </CheckboxPrimitive.Root>
))
Checkbox.displayName = CheckboxPrimitive.Root.displayName

export { Checkbox }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/collapsible.tsx

```tsx
"use client"

import * as CollapsiblePrimitive from "@radix-ui/react-collapsible"

const Collapsible = CollapsiblePrimitive.Root

const CollapsibleTrigger = CollapsiblePrimitive.CollapsibleTrigger

const CollapsibleContent = CollapsiblePrimitive.CollapsibleContent

export { Collapsible, CollapsibleTrigger, CollapsibleContent }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/dropdown-menu.tsx

```tsx
"use client"

import * as React from "react"
import * as DropdownMenuPrimitive from "@radix-ui/react-dropdown-menu"
import { Check, ChevronRight, Circle } from "lucide-react"

import { cn } from "@/lib/utils"

const DropdownMenu = DropdownMenuPrimitive.Root

const DropdownMenuTrigger = DropdownMenuPrimitive.Trigger

const DropdownMenuGroup = DropdownMenuPrimitive.Group

const DropdownMenuPortal = DropdownMenuPrimitive.Portal

const DropdownMenuSub = DropdownMenuPrimitive.Sub

const DropdownMenuRadioGroup = DropdownMenuPrimitive.RadioGroup

const DropdownMenuSubTrigger = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.SubTrigger>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.SubTrigger> & {
    inset?: boolean
  }
>(({ className, inset, children, ...props }, ref) => (
  <DropdownMenuPrimitive.SubTrigger
    ref={ref}
    className={cn(
      "flex cursor-default gap-2 select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent data-[state=open]:bg-accent [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0",
      inset && "pl-8",
      className
    )}
    {...props}
  >
    {children}
    <ChevronRight className="ml-auto" />
  </DropdownMenuPrimitive.SubTrigger>
))
DropdownMenuSubTrigger.displayName =
  DropdownMenuPrimitive.SubTrigger.displayName

const DropdownMenuSubContent = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.SubContent>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.SubContent>
>(({ className, ...props }, ref) => (
  <DropdownMenuPrimitive.SubContent
    ref={ref}
    className={cn(
      "z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-lg data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
      className
    )}
    {...props}
  />
))
DropdownMenuSubContent.displayName =
  DropdownMenuPrimitive.SubContent.displayName

const DropdownMenuContent = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.Content>
>(({ className, sideOffset = 4, ...props }, ref) => (
  <DropdownMenuPrimitive.Portal>
    <DropdownMenuPrimitive.Content
      ref={ref}
      sideOffset={sideOffset}
      className={cn(
        "z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
        className
      )}
      {...props}
    />
  </DropdownMenuPrimitive.Portal>
))
DropdownMenuContent.displayName = DropdownMenuPrimitive.Content.displayName

const DropdownMenuItem = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.Item> & {
    inset?: boolean
  }
>(({ className, inset, ...props }, ref) => (
  <DropdownMenuPrimitive.Item
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center gap-2 rounded-sm px-2 py-1.5 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0",
      inset && "pl-8",
      className
    )}
    {...props}
  />
))
DropdownMenuItem.displayName = DropdownMenuPrimitive.Item.displayName

const DropdownMenuCheckboxItem = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.CheckboxItem>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.CheckboxItem>
>(({ className, children, checked, ...props }, ref) => (
  <DropdownMenuPrimitive.CheckboxItem
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    checked={checked}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <DropdownMenuPrimitive.ItemIndicator>
        <Check className="h-4 w-4" />
      </DropdownMenuPrimitive.ItemIndicator>
    </span>
    {children}
  </DropdownMenuPrimitive.CheckboxItem>
))
DropdownMenuCheckboxItem.displayName =
  DropdownMenuPrimitive.CheckboxItem.displayName

const DropdownMenuRadioItem = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.RadioItem>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.RadioItem>
>(({ className, children, ...props }, ref) => (
  <DropdownMenuPrimitive.RadioItem
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none transition-colors focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <DropdownMenuPrimitive.ItemIndicator>
        <Circle className="h-2 w-2 fill-current" />
      </DropdownMenuPrimitive.ItemIndicator>
    </span>
    {children}
  </DropdownMenuPrimitive.RadioItem>
))
DropdownMenuRadioItem.displayName = DropdownMenuPrimitive.RadioItem.displayName

const DropdownMenuLabel = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.Label>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.Label> & {
    inset?: boolean
  }
>(({ className, inset, ...props }, ref) => (
  <DropdownMenuPrimitive.Label
    ref={ref}
    className={cn(
      "px-2 py-1.5 text-sm font-semibold",
      inset && "pl-8",
      className
    )}
    {...props}
  />
))
DropdownMenuLabel.displayName = DropdownMenuPrimitive.Label.displayName

const DropdownMenuSeparator = React.forwardRef<
  React.ElementRef<typeof DropdownMenuPrimitive.Separator>,
  React.ComponentPropsWithoutRef<typeof DropdownMenuPrimitive.Separator>
>(({ className, ...props }, ref) => (
  <DropdownMenuPrimitive.Separator
    ref={ref}
    className={cn("-mx-1 my-1 h-px bg-muted", className)}
    {...props}
  />
))
DropdownMenuSeparator.displayName = DropdownMenuPrimitive.Separator.displayName

const DropdownMenuShortcut = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement>) => {
  return (
    <span
      className={cn("ml-auto text-xs tracking-widest opacity-60", className)}
      {...props}
    />
  )
}
DropdownMenuShortcut.displayName = "DropdownMenuShortcut"

export {
  DropdownMenu,
  DropdownMenuTrigger,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuCheckboxItem,
  DropdownMenuRadioItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuShortcut,
  DropdownMenuGroup,
  DropdownMenuPortal,
  DropdownMenuSub,
  DropdownMenuSubContent,
  DropdownMenuSubTrigger,
  DropdownMenuRadioGroup,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/select.tsx

```tsx
"use client"

import * as React from "react"
import * as SelectPrimitive from "@radix-ui/react-select"
import { Check, ChevronDown, ChevronUp } from "lucide-react"

import { cn } from "@/lib/utils"

const Select = SelectPrimitive.Root

const SelectGroup = SelectPrimitive.Group

const SelectValue = SelectPrimitive.Value

const SelectTrigger = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.Trigger>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.Trigger>
>(({ className, children, ...props }, ref) => (
  <SelectPrimitive.Trigger
    ref={ref}
    className={cn(
      "flex h-10 w-full items-center justify-between rounded-md border border-input bg-background px-3 py-2 text-sm ring-offset-background placeholder:text-muted-foreground focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 [&>span]:line-clamp-1",
      className
    )}
    {...props}
  >
    {children}
    <SelectPrimitive.Icon asChild>
      <ChevronDown className="h-4 w-4 opacity-50" />
    </SelectPrimitive.Icon>
  </SelectPrimitive.Trigger>
))
SelectTrigger.displayName = SelectPrimitive.Trigger.displayName

const SelectScrollUpButton = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.ScrollUpButton>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.ScrollUpButton>
>(({ className, ...props }, ref) => (
  <SelectPrimitive.ScrollUpButton
    ref={ref}
    className={cn(
      "flex cursor-default items-center justify-center py-1",
      className
    )}
    {...props}
  >
    <ChevronUp className="h-4 w-4" />
  </SelectPrimitive.ScrollUpButton>
))
SelectScrollUpButton.displayName = SelectPrimitive.ScrollUpButton.displayName

const SelectScrollDownButton = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.ScrollDownButton>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.ScrollDownButton>
>(({ className, ...props }, ref) => (
  <SelectPrimitive.ScrollDownButton
    ref={ref}
    className={cn(
      "flex cursor-default items-center justify-center py-1",
      className
    )}
    {...props}
  >
    <ChevronDown className="h-4 w-4" />
  </SelectPrimitive.ScrollDownButton>
))
SelectScrollDownButton.displayName =
  SelectPrimitive.ScrollDownButton.displayName

const SelectContent = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.Content>
>(({ className, children, position = "popper", ...props }, ref) => (
  <SelectPrimitive.Portal>
    <SelectPrimitive.Content
      ref={ref}
      className={cn(
        "relative z-50 max-h-96 min-w-[8rem] overflow-hidden rounded-md border bg-popover text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
        position === "popper" &&
          "data-[side=bottom]:translate-y-1 data-[side=left]:-translate-x-1 data-[side=right]:translate-x-1 data-[side=top]:-translate-y-1",
        className
      )}
      position={position}
      {...props}
    >
      <SelectScrollUpButton />
      <SelectPrimitive.Viewport
        className={cn(
          "p-1",
          position === "popper" &&
            "h-[var(--radix-select-trigger-height)] w-full min-w-[var(--radix-select-trigger-width)]"
        )}
      >
        {children}
      </SelectPrimitive.Viewport>
      <SelectScrollDownButton />
    </SelectPrimitive.Content>
  </SelectPrimitive.Portal>
))
SelectContent.displayName = SelectPrimitive.Content.displayName

const SelectLabel = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.Label>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.Label>
>(({ className, ...props }, ref) => (
  <SelectPrimitive.Label
    ref={ref}
    className={cn("py-1.5 pl-8 pr-2 text-sm font-semibold", className)}
    {...props}
  />
))
SelectLabel.displayName = SelectPrimitive.Label.displayName

const SelectItem = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.Item>
>(({ className, children, ...props }, ref) => (
  <SelectPrimitive.Item
    ref={ref}
    className={cn(
      "relative flex w-full cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <SelectPrimitive.ItemIndicator>
        <Check className="h-4 w-4" />
      </SelectPrimitive.ItemIndicator>
    </span>

    <SelectPrimitive.ItemText>{children}</SelectPrimitive.ItemText>
  </SelectPrimitive.Item>
))
SelectItem.displayName = SelectPrimitive.Item.displayName

const SelectSeparator = React.forwardRef<
  React.ElementRef<typeof SelectPrimitive.Separator>,
  React.ComponentPropsWithoutRef<typeof SelectPrimitive.Separator>
>(({ className, ...props }, ref) => (
  <SelectPrimitive.Separator
    ref={ref}
    className={cn("-mx-1 my-1 h-px bg-muted", className)}
    {...props}
  />
))
SelectSeparator.displayName = SelectPrimitive.Separator.displayName

export {
  Select,
  SelectGroup,
  SelectValue,
  SelectTrigger,
  SelectContent,
  SelectLabel,
  SelectItem,
  SelectSeparator,
  SelectScrollUpButton,
  SelectScrollDownButton,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/textarea.tsx

```tsx
import * as React from "react"

import { cn } from "@/lib/utils"

const Textarea = React.forwardRef<
  HTMLTextAreaElement,
  React.ComponentProps<"textarea">
>(({ className, ...props }, ref) => {
  return (
    <textarea
      className={cn(
        "flex min-h-[80px] w-full rounded-md border border-input bg-background px-3 py-2 text-base ring-offset-background placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
        className
      )}
      ref={ref}
      {...props}
    />
  )
})
Textarea.displayName = "Textarea"

export { Textarea }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/input.tsx

```tsx
import * as React from "react"

import { cn } from "@/lib/utils"

const Input = React.forwardRef<HTMLInputElement, React.ComponentProps<"input">>(
  ({ className, type, ...props }, ref) => {
    return (
      <input
        type={type}
        className={cn(
          "flex h-10 w-full rounded-md border border-input bg-background px-3 py-2 text-base ring-offset-background file:border-0 file:bg-transparent file:text-sm file:font-medium file:text-foreground placeholder:text-muted-foreground focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-ring focus-visible:ring-offset-2 disabled:cursor-not-allowed disabled:opacity-50 md:text-sm",
          className
        )}
        ref={ref}
        {...props}
      />
    )
  }
)
Input.displayName = "Input"

export { Input }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/skeleton.tsx

```tsx
import { cn } from "@/lib/utils"

function Skeleton({
  className,
  ...props
}: React.HTMLAttributes<HTMLDivElement>) {
  return (
    <div
      className={cn("animate-pulse rounded-md bg-muted", className)}
      {...props}
    />
  )
}

export { Skeleton }
```

## File: /Users/saint/Desktop/pdf-player/components/ui/context-menu.tsx

```tsx
"use client"

import * as React from "react"
import * as ContextMenuPrimitive from "@radix-ui/react-context-menu"
import { Check, ChevronRight, Circle } from "lucide-react"

import { cn } from "@/lib/utils"

const ContextMenu = ContextMenuPrimitive.Root

const ContextMenuTrigger = ContextMenuPrimitive.Trigger

const ContextMenuGroup = ContextMenuPrimitive.Group

const ContextMenuPortal = ContextMenuPrimitive.Portal

const ContextMenuSub = ContextMenuPrimitive.Sub

const ContextMenuRadioGroup = ContextMenuPrimitive.RadioGroup

const ContextMenuSubTrigger = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.SubTrigger>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.SubTrigger> & {
    inset?: boolean
  }
>(({ className, inset, children, ...props }, ref) => (
  <ContextMenuPrimitive.SubTrigger
    ref={ref}
    className={cn(
      "flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[state=open]:bg-accent data-[state=open]:text-accent-foreground",
      inset && "pl-8",
      className
    )}
    {...props}
  >
    {children}
    <ChevronRight className="ml-auto h-4 w-4" />
  </ContextMenuPrimitive.SubTrigger>
))
ContextMenuSubTrigger.displayName = ContextMenuPrimitive.SubTrigger.displayName

const ContextMenuSubContent = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.SubContent>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.SubContent>
>(({ className, ...props }, ref) => (
  <ContextMenuPrimitive.SubContent
    ref={ref}
    className={cn(
      "z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
      className
    )}
    {...props}
  />
))
ContextMenuSubContent.displayName = ContextMenuPrimitive.SubContent.displayName

const ContextMenuContent = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.Content>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.Content>
>(({ className, ...props }, ref) => (
  <ContextMenuPrimitive.Portal>
    <ContextMenuPrimitive.Content
      ref={ref}
      className={cn(
        "z-50 min-w-[8rem] overflow-hidden rounded-md border bg-popover p-1 text-popover-foreground shadow-md animate-in fade-in-80 data-[state=open]:animate-in data-[state=closed]:animate-out data-[state=closed]:fade-out-0 data-[state=open]:fade-in-0 data-[state=closed]:zoom-out-95 data-[state=open]:zoom-in-95 data-[side=bottom]:slide-in-from-top-2 data-[side=left]:slide-in-from-right-2 data-[side=right]:slide-in-from-left-2 data-[side=top]:slide-in-from-bottom-2",
        className
      )}
      {...props}
    />
  </ContextMenuPrimitive.Portal>
))
ContextMenuContent.displayName = ContextMenuPrimitive.Content.displayName

const ContextMenuItem = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.Item>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.Item> & {
    inset?: boolean
  }
>(({ className, inset, ...props }, ref) => (
  <ContextMenuPrimitive.Item
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm px-2 py-1.5 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      inset && "pl-8",
      className
    )}
    {...props}
  />
))
ContextMenuItem.displayName = ContextMenuPrimitive.Item.displayName

const ContextMenuCheckboxItem = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.CheckboxItem>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.CheckboxItem>
>(({ className, children, checked, ...props }, ref) => (
  <ContextMenuPrimitive.CheckboxItem
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    checked={checked}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <ContextMenuPrimitive.ItemIndicator>
        <Check className="h-4 w-4" />
      </ContextMenuPrimitive.ItemIndicator>
    </span>
    {children}
  </ContextMenuPrimitive.CheckboxItem>
))
ContextMenuCheckboxItem.displayName =
  ContextMenuPrimitive.CheckboxItem.displayName

const ContextMenuRadioItem = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.RadioItem>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.RadioItem>
>(({ className, children, ...props }, ref) => (
  <ContextMenuPrimitive.RadioItem
    ref={ref}
    className={cn(
      "relative flex cursor-default select-none items-center rounded-sm py-1.5 pl-8 pr-2 text-sm outline-none focus:bg-accent focus:text-accent-foreground data-[disabled]:pointer-events-none data-[disabled]:opacity-50",
      className
    )}
    {...props}
  >
    <span className="absolute left-2 flex h-3.5 w-3.5 items-center justify-center">
      <ContextMenuPrimitive.ItemIndicator>
        <Circle className="h-2 w-2 fill-current" />
      </ContextMenuPrimitive.ItemIndicator>
    </span>
    {children}
  </ContextMenuPrimitive.RadioItem>
))
ContextMenuRadioItem.displayName = ContextMenuPrimitive.RadioItem.displayName

const ContextMenuLabel = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.Label>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.Label> & {
    inset?: boolean
  }
>(({ className, inset, ...props }, ref) => (
  <ContextMenuPrimitive.Label
    ref={ref}
    className={cn(
      "px-2 py-1.5 text-sm font-semibold text-foreground",
      inset && "pl-8",
      className
    )}
    {...props}
  />
))
ContextMenuLabel.displayName = ContextMenuPrimitive.Label.displayName

const ContextMenuSeparator = React.forwardRef<
  React.ElementRef<typeof ContextMenuPrimitive.Separator>,
  React.ComponentPropsWithoutRef<typeof ContextMenuPrimitive.Separator>
>(({ className, ...props }, ref) => (
  <ContextMenuPrimitive.Separator
    ref={ref}
    className={cn("-mx-1 my-1 h-px bg-border", className)}
    {...props}
  />
))
ContextMenuSeparator.displayName = ContextMenuPrimitive.Separator.displayName

const ContextMenuShortcut = ({
  className,
  ...props
}: React.HTMLAttributes<HTMLSpanElement>) => {
  return (
    <span
      className={cn(
        "ml-auto text-xs tracking-widest text-muted-foreground",
        className
      )}
      {...props}
    />
  )
}
ContextMenuShortcut.displayName = "ContextMenuShortcut"

export {
  ContextMenu,
  ContextMenuTrigger,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuCheckboxItem,
  ContextMenuRadioItem,
  ContextMenuLabel,
  ContextMenuSeparator,
  ContextMenuShortcut,
  ContextMenuGroup,
  ContextMenuPortal,
  ContextMenuSub,
  ContextMenuSubContent,
  ContextMenuSubTrigger,
  ContextMenuRadioGroup,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/form.tsx

```tsx
"use client"

import * as React from "react"
import * as LabelPrimitive from "@radix-ui/react-label"
import { Slot } from "@radix-ui/react-slot"
import {
  Controller,
  ControllerProps,
  FieldPath,
  FieldValues,
  FormProvider,
  useFormContext,
} from "react-hook-form"

import { cn } from "@/lib/utils"
import { Label } from "@/components/ui/label"

const Form = FormProvider

type FormFieldContextValue<
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>
> = {
  name: TName
}

const FormFieldContext = React.createContext<FormFieldContextValue>(
  {} as FormFieldContextValue
)

const FormField = <
  TFieldValues extends FieldValues = FieldValues,
  TName extends FieldPath<TFieldValues> = FieldPath<TFieldValues>
>({
  ...props
}: ControllerProps<TFieldValues, TName>) => {
  return (
    <FormFieldContext.Provider value={{ name: props.name }}>
      <Controller {...props} />
    </FormFieldContext.Provider>
  )
}

const useFormField = () => {
  const fieldContext = React.useContext(FormFieldContext)
  const itemContext = React.useContext(FormItemContext)
  const { getFieldState, formState } = useFormContext()

  const fieldState = getFieldState(fieldContext.name, formState)

  if (!fieldContext) {
    throw new Error("useFormField should be used within <FormField>")
  }

  const { id } = itemContext

  return {
    id,
    name: fieldContext.name,
    formItemId: `${id}-form-item`,
    formDescriptionId: `${id}-form-item-description`,
    formMessageId: `${id}-form-item-message`,
    ...fieldState,
  }
}

type FormItemContextValue = {
  id: string
}

const FormItemContext = React.createContext<FormItemContextValue>(
  {} as FormItemContextValue
)

const FormItem = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  const id = React.useId()

  return (
    <FormItemContext.Provider value={{ id }}>
      <div ref={ref} className={cn("space-y-2", className)} {...props} />
    </FormItemContext.Provider>
  )
})
FormItem.displayName = "FormItem"

const FormLabel = React.forwardRef<
  React.ElementRef<typeof LabelPrimitive.Root>,
  React.ComponentPropsWithoutRef<typeof LabelPrimitive.Root>
>(({ className, ...props }, ref) => {
  const { error, formItemId } = useFormField()

  return (
    <Label
      ref={ref}
      className={cn(error && "text-destructive", className)}
      htmlFor={formItemId}
      {...props}
    />
  )
})
FormLabel.displayName = "FormLabel"

const FormControl = React.forwardRef<
  React.ElementRef<typeof Slot>,
  React.ComponentPropsWithoutRef<typeof Slot>
>(({ ...props }, ref) => {
  const { error, formItemId, formDescriptionId, formMessageId } = useFormField()

  return (
    <Slot
      ref={ref}
      id={formItemId}
      aria-describedby={
        !error
          ? `${formDescriptionId}`
          : `${formDescriptionId} ${formMessageId}`
      }
      aria-invalid={!!error}
      {...props}
    />
  )
})
FormControl.displayName = "FormControl"

const FormDescription = React.forwardRef<
  HTMLParagraphElement,
  React.HTMLAttributes<HTMLParagraphElement>
>(({ className, ...props }, ref) => {
  const { formDescriptionId } = useFormField()

  return (
    <p
      ref={ref}
      id={formDescriptionId}
      className={cn("text-sm text-muted-foreground", className)}
      {...props}
    />
  )
})
FormDescription.displayName = "FormDescription"

const FormMessage = React.forwardRef<
  HTMLParagraphElement,
  React.HTMLAttributes<HTMLParagraphElement>
>(({ className, children, ...props }, ref) => {
  const { error, formMessageId } = useFormField()
  const body = error ? String(error?.message) : children

  if (!body) {
    return null
  }

  return (
    <p
      ref={ref}
      id={formMessageId}
      className={cn("text-sm font-medium text-destructive", className)}
      {...props}
    >
      {body}
    </p>
  )
})
FormMessage.displayName = "FormMessage"

export {
  useFormField,
  Form,
  FormItem,
  FormLabel,
  FormControl,
  FormDescription,
  FormMessage,
  FormField,
}
```

## File: /Users/saint/Desktop/pdf-player/components/ui/carousel.tsx

```tsx
"use client"

import * as React from "react"
import useEmblaCarousel, {
  type UseEmblaCarouselType,
} from "embla-carousel-react"
import { ArrowLeft, ArrowRight } from "lucide-react"

import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"

type CarouselApi = UseEmblaCarouselType[1]
type UseCarouselParameters = Parameters<typeof useEmblaCarousel>
type CarouselOptions = UseCarouselParameters[0]
type CarouselPlugin = UseCarouselParameters[1]

type CarouselProps = {
  opts?: CarouselOptions
  plugins?: CarouselPlugin
  orientation?: "horizontal" | "vertical"
  setApi?: (api: CarouselApi) => void
}

type CarouselContextProps = {
  carouselRef: ReturnType<typeof useEmblaCarousel>[0]
  api: ReturnType<typeof useEmblaCarousel>[1]
  scrollPrev: () => void
  scrollNext: () => void
  canScrollPrev: boolean
  canScrollNext: boolean
} & CarouselProps

const CarouselContext = React.createContext<CarouselContextProps | null>(null)

function useCarousel() {
  const context = React.useContext(CarouselContext)

  if (!context) {
    throw new Error("useCarousel must be used within a <Carousel />")
  }

  return context
}

const Carousel = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement> & CarouselProps
>(
  (
    {
      orientation = "horizontal",
      opts,
      setApi,
      plugins,
      className,
      children,
      ...props
    },
    ref
  ) => {
    const [carouselRef, api] = useEmblaCarousel(
      {
        ...opts,
        axis: orientation === "horizontal" ? "x" : "y",
      },
      plugins
    )
    const [canScrollPrev, setCanScrollPrev] = React.useState(false)
    const [canScrollNext, setCanScrollNext] = React.useState(false)

    const onSelect = React.useCallback((api: CarouselApi) => {
      if (!api) {
        return
      }

      setCanScrollPrev(api.canScrollPrev())
      setCanScrollNext(api.canScrollNext())
    }, [])

    const scrollPrev = React.useCallback(() => {
      api?.scrollPrev()
    }, [api])

    const scrollNext = React.useCallback(() => {
      api?.scrollNext()
    }, [api])

    const handleKeyDown = React.useCallback(
      (event: React.KeyboardEvent<HTMLDivElement>) => {
        if (event.key === "ArrowLeft") {
          event.preventDefault()
          scrollPrev()
        } else if (event.key === "ArrowRight") {
          event.preventDefault()
          scrollNext()
        }
      },
      [scrollPrev, scrollNext]
    )

    React.useEffect(() => {
      if (!api || !setApi) {
        return
      }

      setApi(api)
    }, [api, setApi])

    React.useEffect(() => {
      if (!api) {
        return
      }

      onSelect(api)
      api.on("reInit", onSelect)
      api.on("select", onSelect)

      return () => {
        api?.off("select", onSelect)
      }
    }, [api, onSelect])

    return (
      <CarouselContext.Provider
        value={{
          carouselRef,
          api: api,
          opts,
          orientation:
            orientation || (opts?.axis === "y" ? "vertical" : "horizontal"),
          scrollPrev,
          scrollNext,
          canScrollPrev,
          canScrollNext,
        }}
      >
        <div
          ref={ref}
          onKeyDownCapture={handleKeyDown}
          className={cn("relative", className)}
          role="region"
          aria-roledescription="carousel"
          {...props}
        >
          {children}
        </div>
      </CarouselContext.Provider>
    )
  }
)
Carousel.displayName = "Carousel"

const CarouselContent = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  const { carouselRef, orientation } = useCarousel()

  return (
    <div ref={carouselRef} className="overflow-hidden">
      <div
        ref={ref}
        className={cn(
          "flex",
          orientation === "horizontal" ? "-ml-4" : "-mt-4 flex-col",
          className
        )}
        {...props}
      />
    </div>
  )
})
CarouselContent.displayName = "CarouselContent"

const CarouselItem = React.forwardRef<
  HTMLDivElement,
  React.HTMLAttributes<HTMLDivElement>
>(({ className, ...props }, ref) => {
  const { orientation } = useCarousel()

  return (
    <div
      ref={ref}
      role="group"
      aria-roledescription="slide"
      className={cn(
        "min-w-0 shrink-0 grow-0 basis-full",
        orientation === "horizontal" ? "pl-4" : "pt-4",
        className
      )}
      {...props}
    />
  )
})
CarouselItem.displayName = "CarouselItem"

const CarouselPrevious = React.forwardRef<
  HTMLButtonElement,
  React.ComponentProps<typeof Button>
>(({ className, variant = "outline", size = "icon", ...props }, ref) => {
  const { orientation, scrollPrev, canScrollPrev } = useCarousel()

  return (
    <Button
      ref={ref}
      variant={variant}
      size={size}
      className={cn(
        "absolute  h-8 w-8 rounded-full",
        orientation === "horizontal"
          ? "-left-12 top-1/2 -translate-y-1/2"
          : "-top-12 left-1/2 -translate-x-1/2 rotate-90",
        className
      )}
      disabled={!canScrollPrev}
      onClick={scrollPrev}
      {...props}
    >
      <ArrowLeft className="h-4 w-4" />
      <span className="sr-only">Previous slide</span>
    </Button>
  )
})
CarouselPrevious.displayName = "CarouselPrevious"

const CarouselNext = React.forwardRef<
  HTMLButtonElement,
  React.ComponentProps<typeof Button>
>(({ className, variant = "outline", size = "icon", ...props }, ref) => {
  const { orientation, scrollNext, canScrollNext } = useCarousel()

  return (
    <Button
      ref={ref}
      variant={variant}
      size={size}
      className={cn(
        "absolute h-8 w-8 rounded-full",
        orientation === "horizontal"
          ? "-right-12 top-1/2 -translate-y-1/2"
          : "-bottom-12 left-1/2 -translate-x-1/2 rotate-90",
        className
      )}
      disabled={!canScrollNext}
      onClick={scrollNext}
      {...props}
    >
      <ArrowRight className="h-4 w-4" />
      <span className="sr-only">Next slide</span>
    </Button>
  )
})
CarouselNext.displayName = "CarouselNext"

export {
  type CarouselApi,
  Carousel,
  CarouselContent,
  CarouselItem,
  CarouselPrevious,
  CarouselNext,
}
```

## File: /Users/saint/Desktop/pdf-player/components/PDFReader.tsx

```tsx
"use client"

import { useState, useRef, useEffect } from "react"
import { Play, Pause, Repeat, Shuffle, Volume2 } from "lucide-react"
import { useTTS } from '@/contexts/TTSContext'

interface PDFReaderProps {
  text: string
  title?: string
}

export default function PDFReader({ text, title = "Uploaded PDF" }: PDFReaderProps) {
  const { 
    isPlaying,
    currentTime,
    duration,
    currentSegment,
    segments,
    currentBook,
    play,
    pause,
    skipForward,
    skipBackward,
    seekTo,
    setSpeed
  } = useTTS()
  const [progress, setProgress] = useState(0)
  const [volume, setVolume] = useState(0.75)
  const [isRepeat, setIsRepeat] = useState(false)
  const [isShuffle, setIsShuffle] = useState(false)
  const audioRef = useRef<HTMLAudioElement | null>(null)
  const textRef = useRef<HTMLDivElement | null>(null)

  const formatTime = (time: number) => {
    const minutes = Math.floor(time / 60)
    const seconds = Math.floor(time % 60)
    return `${minutes}:${seconds.toString().padStart(2, "0")}`
  }

  const handleTimeUpdate = () => {
    if (audioRef.current) {
      const current = audioRef.current.currentTime
      const duration = audioRef.current.duration
      setProgress((current / duration) * 100)
    }
  }

  const handlePlayPause = () => {
    if (isPlaying) {
      pause()
    } else {
      play()
    }
  }

  const handleSeek = (time: number) => {
    seekTo(time)
  }

  const handleVolumeChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const newVolume = parseFloat(e.target.value)
    setVolume(newVolume)
    if (audioRef.current) {
      audioRef.current.volume = newVolume
    }
  }

  const handleRepeatToggle = () => {
    setIsRepeat(!isRepeat)
  }

  const handleShuffleToggle = () => {
    setIsShuffle(!isShuffle)
  }

  // Update current segment highlight based on playback
  useEffect(() => {
    if (segments[currentSegment]) {
      const segment = segments[currentSegment]
      const textContent = textRef.current
      if (textContent) {
        // Remove previous highlights
        const highlights = textContent.getElementsByClassName('highlight')
        Array.from(highlights).forEach(el => {
          el.classList.remove('highlight')
        })

        // Add highlight to current segment
        const segmentText = segment.content
        const range = document.createRange()
        const textNodes = Array.from(textContent.childNodes).filter(node => node.nodeType === Node.TEXT_NODE)
        
        for (const node of textNodes) {
          const text = node.textContent || ''
          const index = text.indexOf(segmentText)
          if (index >= 0) {
            range.setStart(node, index)
            range.setEnd(node, index + segmentText.length)
            const span = document.createElement('span')
            span.className = 'highlight'
            range.surroundContents(span)
            break
          }
        }
      }
    }
  }, [currentSegment, segments])

  // Handle audio end
  useEffect(() => {
    const handleEnded = () => {
      if (isRepeat) {
        playSegment(currentSegment)
      } else if (isShuffle) {
        const nextSegment = Math.floor(Math.random() * segments.length)
        playSegment(nextSegment)
      } else {
        const nextSegment = currentSegment + 1
        if (nextSegment < segments.length) {
          playSegment(nextSegment)
        }
      }
    }

    if (audioRef.current) {
      audioRef.current.addEventListener('ended', handleEnded)
      return () => {
        audioRef.current?.removeEventListener('ended', handleEnded)
      }
    }
  }, [isRepeat, isShuffle, currentSegment, segments])

  // Set initial volume
  useEffect(() => {
    if (audioRef.current) {
      audioRef.current.volume = volume
    }
  }, [volume])

  return (
    <div className="fixed inset-0 flex items-center justify-center p-4 bg-gradient-to-b from-gray-100/50 to-gray-200/50 dark:from-gray-900/50 dark:to-gray-800/50">
      <div className="w-full max-w-md">
        <div className="glass-effect rounded-3xl overflow-hidden player-shadow">
          <div className="p-6 space-y-8">
            {/* Large PDF Cover Preview */}
            <div className="aspect-[4/3] w-full rounded-2xl overflow-hidden glass-effect-strong">
              <div className="w-full h-full bg-gradient-to-br from-gray-200 to-gray-300 dark:from-gray-700 dark:to-gray-800 flex items-center justify-center relative group">
                <div className="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-300" />
                {/* PDF Preview Image would go here */}
                <div className="absolute inset-0 flex items-center justify-center">
                  <span className="text-2xl font-medium text-gray-500 dark:text-gray-400">PDF</span>
                </div>
                {/* Play button overlay */}
                <button
                  onClick={handlePlayPause}
                  className="absolute invisible group-hover:visible opacity-0 group-hover:opacity-100 transition-all duration-300 p-4 rounded-full glass-effect-strong text-white hover:scale-105"
                >
                  {isPlaying ? <Pause className="w-12 h-12" /> : <Play className="w-12 h-12" />}
                </button>
              </div>
            </div>

            {/* PDF Title and Status */}
            <div className="flex items-center justify-between">
              <div className="flex-1 min-w-0">
                <h2 className="text-xl font-semibold text-gray-900 dark:text-white truncate">{title}</h2>
                <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">
                  {currentBook && `Page ${currentBook.currentPage} of ${currentBook.pageCount}`}
                </p>
              </div>
              <div className="w-6 h-6 rounded-full bg-emerald-400/20 flex items-center justify-center ml-4">
                <div className="w-2 h-2 rounded-full bg-emerald-400"></div>
              </div>
            </div>

            {/* Progress bar */}
            <div className="space-y-2">
              <div className="h-1.5 rounded-full progress-gradient overflow-hidden">
                <div
                  className="h-full bg-white/30 transition-all duration-200 ease-out"
                  style={{ width: `${progress}%` }}
                ></div>
              </div>
              <div className="flex justify-between text-sm text-gray-400">
                <span>{formatTime(currentTime)}</span>
                <span>{formatTime(duration)}</span>
              </div>
            </div>

            {/* Controls */}
            <div className="flex items-center justify-between">
              <button 
                className={`p-2 ${isShuffle ? 'text-white' : 'text-gray-400'} hover:text-white transition-colors`}
                onClick={handleShuffleToggle}
              >
                <Shuffle className="w-5 h-5" />
              </button>
              <div className="flex items-center space-x-6">
                <button
                  onClick={skipBackward}
                  className="p-3 text-gray-400 hover:text-white transition-colors flex flex-col items-center"
                >
                  <svg
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    className="w-7 h-7"
                  >
                    <path
                      d="M12.5 8C9.85 8 7.45 8.99 5.6 10.6L2 7V16H11L7.38 12.38C8.77 11.22 10.54 10.5 12.5 10.5C16.04 10.5 19.05 12.81 20.1 16L22.47 15.22C21.08 11.03 17.15 8 12.5 8Z"
                      fill="currentColor"
                    />
                  </svg>
                  <span className="text-xs mt-1">10</span>
                </button>
                <button
                  onClick={handlePlayPause}
                  className="p-5 rounded-full glass-effect-strong text-white hover:scale-105 transition-transform"
                >
                  {isPlaying ? <Pause className="w-10 h-10" /> : <Play className="w-10 h-10" />}
                </button>
                <button
                  onClick={skipForward}
                  className="p-3 text-gray-400 hover:text-white transition-colors flex flex-col items-center"
                >
                  <svg
                    width="24"
                    height="24"
                    viewBox="0 0 24 24"
                    fill="none"
                    xmlns="http://www.w3.org/2000/svg"
                    className="w-7 h-7"
                  >
                    <path
                      d="M18.4 10.6C16.55 8.99 14.15 8 11.5 8C6.85 8 2.92 11.03 1.54 15.22L3.9 16C4.95 12.81 7.96 10.5 11.5 10.5C13.45 10.5 15.23 11.22 16.62 12.38L13 16H22V7L18.4 10.6Z"
                      fill="currentColor"
                    />
                  </svg>
                  <span className="text-xs mt-1">10</span>
                </button>
              </div>
              <button 
                className={`p-2 ${isRepeat ? 'text-white' : 'text-gray-400'} hover:text-white transition-colors`}
                onClick={handleRepeatToggle}
              >
                <Repeat className="w-5 h-5" />
              </button>
            </div>

            {/* Volume control */}
            <div className="flex items-center space-x-3">
              <Volume2 className="w-5 h-5 text-gray-400" />
              <div className="flex-1 h-1.5 rounded-full progress-gradient">
                <input
                  type="range"
                  min="0"
                  max="1"
                  step="0.01"
                  value={volume}
                  onChange={handleVolumeChange}
                  className="w-full h-full appearance-none bg-transparent [&::-webkit-slider-thumb]:appearance-none [&::-webkit-slider-thumb]:w-3 [&::-webkit-slider-thumb]:h-3 [&::-webkit-slider-thumb]:rounded-full [&::-webkit-slider-thumb]:bg-white [&::-webkit-slider-thumb]:cursor-pointer"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
      <div ref={textRef} className="hidden">{text}</div>
    </div>
  )
}

```

## File: /Users/saint/Desktop/pdf-player/components/pdf-reader.tsx

```tsx
"use client"

import { useState, useRef, useEffect } from "react"
import { Play, Pause, Repeat, Shuffle, Volume2 } from "lucide-react"

interface PDFReaderProps {
  text: string
  title?: string
}

export default function PDFReader({ text, title = "Uploaded PDF" }: PDFReaderProps) {
  const [isPlaying, setIsPlaying] = useState(false)
  const [progress, setProgress] = useState(0)
  const [currentTime, setCurrentTime] = useState("0:00")
  const [duration, setDuration] = useState("-2:13") // Simulated duration
  const audioRef = useRef<HTMLAudioElement | null>(null)

  // ... (rest of the component remains the same)

  return (
    <div className="fixed inset-0 flex items-center justify-center p-4 bg-gradient-to-b from-gray-100/50 to-gray-200/50 dark:from-gray-900/50 dark:to-gray-800/50">
      <div className="w-full max-w-md">
        <div className="glass-effect rounded-3xl overflow-hidden player-shadow">
          <div className="p-6 space-y-8">
            {/* Large PDF Cover Preview */}
            <div className="aspect-[4/3] w-full rounded-2xl overflow-hidden glass-effect-strong">
              <div className="w-full h-full bg-gradient-to-br from-gray-200 to-gray-300 dark:from-gray-700 dark:to-gray-800 flex items-center justify-center relative group">
                <div className="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-300" />
                {/* PDF Preview Image would go here */}
                <div className="absolute inset-0 flex items-center justify-center">
                  <span className="text-2xl font-medium text-gray-500 dark:text-gray-400">PDF</span>
                </div>
                {/* Play button overlay */}
                <button
                  onClick={togglePlayPause}
                  className="absolute invisible group-hover:visible opacity-0 group-hover:opacity-100 transition-all duration-300 p-4 rounded-full glass-effect-strong text-white hover:scale-105"
                >
                  {isPlaying ? <Pause className="w-12 h-12" /> : <Play className="w-12 h-12" />}
                </button>
              </div>
            </div>

            {/* PDF Title and Status */}
            <div className="flex items-center justify-between">
              <div className="flex-1 min-w-0">
                <h2 className="text-xl font-semibold text-gray-900 dark:text-white truncate">{title}</h2>
                <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">Chapter 1 • Page 12</p>
              </div>
              <div className="w-6 h-6 rounded-full bg-emerald-400/20 flex items-center justify-center ml-4">
                <div className="w-2 h-2 rounded-full bg-emerald-400"></div>
              </div>
            </div>

            {/* ... (rest of the component remains the same) */}
          </div>
        </div>
      </div>
    </div>
  )
}

```

## File: /Users/saint/Desktop/pdf-player/components/PDFUploader.tsx

```tsx
"use client"

import { useState } from "react"
import { Upload, File } from "lucide-react"
import { useUploadThing } from "@/lib/uploadthing"
import { generateClientDropzoneAccept } from "uploadthing/client"
import React from "react"
import { useTTS } from '@/contexts/TTSContext'
import { Book } from "@/types"

export default function PDFUploader({ onUploadComplete }: { onUploadComplete: (text: string) => void }) {
  const [file, setFile] = useState<File | null>(null)
  const [isUploading, setIsUploading] = useState(false)
  const [uploadProgress, setUploadProgress] = useState(0)
  const { startUpload, permittedFileInfo } = useUploadThing("pdfUploader", {
    onUploadProgress: (progress) => {
      setUploadProgress(progress)
    },
    onUploadBegin: () => {
      console.log("Starting UploadThing upload...")
    },
    onUploadError: (error) => {
      console.error("Upload error:", error)
      setIsUploading(false)
      setUploadProgress(0)
    }
  })
  const { loadBookAudio } = useTTS()

  const handleFileChange = (event: React.ChangeEvent<HTMLInputElement>) => {
    const selectedFile = event.target.files?.[0]
    if (selectedFile && selectedFile.type === "application/pdf") {
      setFile(selectedFile)
    } else {
      alert("Please select a valid PDF file.")
    }
  }

  const handleUpload = async () => {
    if (!file) return
    try {
      setIsUploading(true)
      console.log("Starting upload process...")

      // Upload to UploadThing first
      const [res] = await startUpload([file])
      if (!res) throw new Error('Upload failed')
      
      console.log("UploadThing upload complete, creating book record...")

      // Create book record with the URL
      const createResponse = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/upload`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          fileUrl: res.url,
          title: file.name.replace(/\.pdf$/, '') // Remove .pdf extension for title
        })
      })

      if (!createResponse.ok) {
        throw new Error('Failed to create book record')
      }

      const { id } = await createResponse.json()
      console.log("Book record created, transitioning to player...")

      // Load audio segments for the book - this will start polling for new segments
      await loadBookAudio({ id } as Book)
      
      onUploadComplete('')
    } catch (error) {
      console.error('Upload error:', error)
    } finally {
      setIsUploading(false)
      setUploadProgress(0)
    }
  }

  // Pre-compute the accept prop to avoid generating during render
  const acceptProp = React.useMemo(() => {
    if (!permittedFileInfo?.config) return { "application/pdf": [".pdf"] }
    return generateClientDropzoneAccept(permittedFileInfo.config)
  }, [permittedFileInfo?.config])

  return (
    <div className="w-full max-w-md mx-auto">
      <div className="glass-effect rounded-3xl overflow-hidden player-shadow">
        <div className="p-6 space-y-6">
          {/* PDF Upload Area */}
          <div className="aspect-square w-full rounded-2xl overflow-hidden glass-effect-strong">
            <label
              htmlFor="dropzone-file"
              className="w-full h-full bg-gradient-to-br from-gray-200 to-gray-300 dark:from-gray-700 dark:to-gray-800 flex flex-col items-center justify-center relative group cursor-pointer"
            >
              <div className="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-300" />
              <input
                id="dropzone-file"
                type="file"
                className="hidden"
                onChange={handleFileChange}
                accept={acceptProp}
              />
              {file ? (
                <File className="w-20 h-20 text-gray-400 mb-4" />
              ) : (
                <Upload className="w-20 h-20 text-gray-400 mb-4" />
              )}
              <p className="text-lg font-medium text-gray-500 dark:text-gray-400 text-center px-4">
                {file ? file.name : "Click to upload PDF"}
              </p>
              <p className="text-sm text-gray-500 dark:text-gray-400 mt-2">
                {file ? `${(file.size / 1024 / 1024).toFixed(2)} MB` : "PDF (MAX. 32MB)"}
              </p>
            </label>
          </div>

          {/* Upload Button and Progress */}
          <div className="space-y-4">
            <button
              onClick={handleUpload}
              disabled={!file || isUploading}
              className="w-full p-3 rounded-full glass-effect-strong text-white hover:scale-105 transition-transform disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center"
            >
              {isUploading ? "Uploading..." : "Upload PDF"}
              {!isUploading && <Upload className="w-5 h-5 ml-2" />}
            </button>
            {isUploading && (
              <div className="space-y-2">
                <div className="h-1.5 rounded-full progress-gradient overflow-hidden">
                  <div
                    className="h-full bg-white/30 transition-all duration-200 ease-out"
                    style={{ width: `${uploadProgress}%` }}
                  ></div>
                </div>
                <div className="flex justify-between text-sm text-gray-400">
                  <span>{uploadProgress}%</span>
                  <span>Uploading...</span>
                </div>
              </div>
            )}
          </div>
        </div>
      </div>
    </div>
  )
}

```

## File: /Users/saint/Desktop/pdf-player/components/book-drawer.tsx

```tsx
"use client"

import * as React from "react"
import { ChevronUp } from "lucide-react"
import { motion, AnimatePresence } from "framer-motion"
import { Book } from "@/types"
import { useState } from "react"
import { useTTS } from "@/contexts/TTSContext"

interface BookDrawerProps {
  books: Book[]
  onSelectBook: (book: Book) => void
  loading?: boolean
  error?: string | null
}

export function BookDrawer({ books, onSelectBook, loading, error }: BookDrawerProps) {
  const [isOpen, setIsOpen] = useState(false)
  const { loadBookAudio } = useTTS()

  const handleBookSelect = async (book: Book) => {
    try {
      await loadBookAudio(book)
      onSelectBook(book)
    } catch (err) {
      console.error('Error loading book audio:', err)
      // Still allow book selection even if audio fails
      onSelectBook(book)
    }
  }

  return (
    <div className="fixed bottom-0 left-0 right-0 z-50">
      <div
        className="h-1 w-16 bg-gray-300 dark:bg-gray-700 rounded-full mx-auto mb-2 cursor-pointer"
        onClick={() => setIsOpen(!isOpen)}
      />
      <AnimatePresence>
        {isOpen && (
          <motion.div
            initial={{ y: "100%" }}
            animate={{ y: 0 }}
            exit={{ y: "100%" }}
            transition={{ type: "spring", stiffness: 300, damping: 30 }}
            className="bg-gradient-to-b from-gray-100/90 to-gray-200/90 dark:from-gray-900/90 dark:to-gray-800/90 backdrop-blur-xl rounded-t-3xl overflow-hidden"
          >
            <div className="p-4 space-y-4">
              <div className="flex justify-between items-center">
                <h2 className="text-xl font-semibold text-gray-900 dark:text-white">Your Books</h2>
                <ChevronUp
                  className="w-6 h-6 text-gray-500 dark:text-gray-400 cursor-pointer"
                  onClick={() => setIsOpen(false)}
                />
              </div>
              <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4 max-h-[60vh] overflow-y-auto pb-4">
                {loading ? (
                  // Loading skeletons
                  Array.from({ length: 4 }).map((_, i) => (
                    <div
                      key={`skeleton-${i}`}
                      className="aspect-[3/4] rounded-xl overflow-hidden glass-effect-strong animate-pulse"
                    >
                      <div className="w-full h-full bg-gray-300 dark:bg-gray-700" />
                    </div>
                  ))
                ) : error ? (
                  <div className="col-span-full text-center text-red-500 dark:text-red-400">
                    {error}
                  </div>
                ) : books.length === 0 ? (
                  <div className="col-span-full text-center text-gray-500 dark:text-gray-400">
                    No books yet. Upload your first PDF!
                  </div>
                ) : (
                  books.map((book) => (
                    <div
                      key={book.id}
                      className="aspect-[3/4] rounded-xl overflow-hidden glass-effect-strong cursor-pointer hover:scale-105 transition-transform duration-200"
                      onClick={() => handleBookSelect(book)}
                    >
                      <div className="w-full h-full bg-gradient-to-br from-gray-200 to-gray-300 dark:from-gray-700 dark:to-gray-800 relative group">
                        <img
                          src={book.coverUrl || "/placeholder.svg"}
                          alt={book.title}
                          className="w-full h-full object-cover"
                        />
                        <div className="absolute inset-0 bg-black/0 group-hover:bg-black/20 transition-colors duration-300" />
                        <div className="absolute bottom-0 left-0 right-0 p-2 bg-gradient-to-t from-black/70 to-transparent">
                          <p className="text-sm font-medium text-white truncate">{book.title}</p>
                          <p className="text-xs text-gray-300 truncate">{book.author}</p>
                        </div>
                      </div>
                    </div>
                  ))
                )}
              </div>
            </div>
          </motion.div>
        )}
      </AnimatePresence>
    </div>
  )
}

```

## File: /Users/saint/Desktop/pdf-player/components/HomePage.tsx

```tsx
"use client"

import { useState } from "react"
import { Moon, Sun } from "lucide-react"
import { Button } from "@/components/ui/button"
import { useTheme } from "next-themes"
import PDFUploader from "@/components/PDFUploader"
import PDFReader from "@/components/PDFReader"
import { BookDrawer } from "@/components/book-drawer"
import { Switch } from "@/components/ui/switch"
import { Book } from "@/types"
import { useBooks } from "@/contexts/BooksContext"

export default function HomePage() {
  const [currentBook, setCurrentBook] = useState<Book | null>(null)
  const { theme, setTheme } = useTheme()
  const { books, loading, error, addBook, refreshBooks } = useBooks()

  const toggleTheme = () => {
    setTheme(theme === "dark" ? "light" : "dark")
  }

  const handleUploadComplete = (text: string) => {
    const newBook: Book = {
      id: currentBook?.id || '', // The ID will be set by the backend
      title: currentBook?.title || 'Uploaded Book',
      author: currentBook?.author || "Unknown",
      coverUrl: currentBook?.coverUrl || "/placeholder.svg?height=400&width=300",
      content: text,
      status: 'processed'
    }
    addBook(newBook)
    setCurrentBook(newBook)
  }

  const handleSelectBook = (book: Book) => {
    setCurrentBook(book)
  }

  return (
    <div className="min-h-screen bg-gradient-to-b from-gray-100 to-gray-200 dark:from-gray-900 dark:to-gray-800 transition-colors duration-300">
      <div className="fixed top-4 right-4 z-50 p-2 rounded-full glass-effect">
        <Switch
          checked={theme === "dark"}
          onCheckedChange={(checked) => setTheme(checked ? "dark" : "light")}
          className="glass-effect"
        />
      </div>
      <main className="flex min-h-screen flex-col items-center justify-center p-4">
        <div className="w-full max-w-md transition-all duration-500 ease-in-out">
          {!currentBook ? (
            <PDFUploader onUploadComplete={handleUploadComplete} />
          ) : (
            <PDFReader text={currentBook.content} title={currentBook.title} />
          )}
        </div>
      </main>
      <BookDrawer 
        books={books} 
        onSelectBook={handleSelectBook} 
        loading={loading}
        error={error}
      />
    </div>
  )
} 
```

## File: /Users/saint/Desktop/pdf-player/public/placeholder-logo.svg

```svg
<svg xmlns="http://www.w3.org/2000/svg" width="215" height="48" fill="none"><path fill="#000" d="M57.588 9.6h6L73.828 38h-5.2l-2.36-6.88h-11.36L52.548 38h-5.2l10.24-28.4Zm7.16 17.16-4.16-12.16-4.16 12.16h8.32Zm23.694-2.24c-.186-1.307-.706-2.32-1.56-3.04-.853-.72-1.866-1.08-3.04-1.08-1.68 0-2.986.613-3.92 1.84-.906 1.227-1.36 2.947-1.36 5.16s.454 3.933 1.36 5.16c.934 1.227 2.24 1.84 3.92 1.84 1.254 0 2.307-.373 3.16-1.12.854-.773 1.387-1.867 1.6-3.28l5.12.24c-.186 1.68-.733 3.147-1.64 4.4-.906 1.227-2.08 2.173-3.52 2.84-1.413.667-2.986 1-4.72 1-2.08 0-3.906-.453-5.48-1.36-1.546-.907-2.76-2.2-3.64-3.88-.853-1.68-1.28-3.627-1.28-5.84 0-2.24.427-4.187 1.28-5.84.88-1.68 2.094-2.973 3.64-3.88 1.574-.907 3.4-1.36 5.48-1.36 1.68 0 3.227.32 4.64.96 1.414.64 2.56 1.56 3.44 2.76.907 1.2 1.454 2.6 1.64 4.2l-5.12.28Zm11.486-7.72.12 3.4c.534-1.227 1.307-2.173 2.32-2.84 1.04-.693 2.267-1.04 3.68-1.04 1.494 0 2.76.387 3.8 1.16 1.067.747 1.827 1.813 2.28 3.2.507-1.44 1.294-2.52 2.36-3.24 1.094-.747 2.414-1.12 3.96-1.12 1.414 0 2.64.307 3.68.92s1.84 1.52 2.4 2.72c.56 1.2.84 2.667.84 4.4V38h-4.96V25.92c0-1.813-.293-3.187-.88-4.12-.56-.96-1.413-1.44-2.56-1.44-.906 0-1.68.213-2.32.64-.64.427-1.133 1.053-1.48 1.88-.32.827-.48 1.84-.48 3.04V38h-4.56V25.92c0-1.2-.133-2.213-.4-3.04-.24-.827-.626-1.453-1.16-1.88-.506-.427-1.133-.64-1.88-.64-.906 0-1.68.227-2.32.68-.64.427-1.133 1.053-1.48 1.88-.32.827-.48 1.827-.48 3V38h-4.96V16.8h4.48Zm26.723 10.6c0-2.24.427-4.187 1.28-5.84.854-1.68 2.067-2.973 3.64-3.88 1.574-.907 3.4-1.36 5.48-1.36 1.84 0 3.494.413 4.96 1.24 1.467.827 2.64 2.08 3.52 3.76.88 1.653 1.347 3.693 1.4 6.12v1.32h-15.08c.107 1.813.614 3.227 1.52 4.24.907.987 2.134 1.48 3.68 1.48.987 0 1.88-.253 2.68-.76a4.803 4.803 0 0 0 1.84-2.2l5.08.36c-.64 2.027-1.84 3.64-3.6 4.84-1.733 1.173-3.733 1.76-6 1.76-2.08 0-3.906-.453-5.48-1.36-1.573-.907-2.786-2.2-3.64-3.88-.853-1.68-1.28-3.627-1.28-5.84Zm15.16-2.04c-.213-1.733-.76-3.013-1.64-3.84-.853-.827-1.893-1.24-3.12-1.24-1.44 0-2.6.453-3.48 1.36-.88.88-1.44 2.12-1.68 3.72h9.92ZM163.139 9.6V38h-5.04V9.6h5.04Zm8.322 7.2.24 5.88-.64-.36c.32-2.053 1.094-3.56 2.32-4.52 1.254-.987 2.787-1.48 4.6-1.48 2.32 0 4.107.733 5.36 2.2 1.254 1.44 1.88 3.387 1.88 5.84V38h-4.96V25.92c0-1.253-.12-2.28-.36-3.08-.24-.8-.64-1.413-1.2-1.84-.533-.427-1.253-.64-2.16-.64-1.44 0-2.573.48-3.4 1.44-.8.933-1.2 2.307-1.2 4.12V38h-4.96V16.8h4.48Zm30.003 7.72c-.186-1.307-.706-2.32-1.56-3.04-.853-.72-1.866-1.08-3.04-1.08-1.68 0-2.986.613-3.92 1.84-.906 1.227-1.36 2.947-1.36 5.16s.454 3.933 1.36 5.16c.934 1.227 2.24 1.84 3.92 1.84 1.254 0 2.307-.373 3.16-1.12.854-.773 1.387-1.867 1.6-3.28l5.12.24c-.186 1.68-.733 3.147-1.64 4.4-.906 1.227-2.08 2.173-3.52 2.84-1.413.667-2.986 1-4.72 1-2.08 0-3.906-.453-5.48-1.36-1.546-.907-2.76-2.2-3.64-3.88-.853-1.68-1.28-3.627-1.28-5.84 0-2.24.427-4.187 1.28-5.84.88-1.68 2.094-2.973 3.64-3.88 1.574-.907 3.4-1.36 5.48-1.36 1.68 0 3.227.32 4.64.96 1.414.64 2.56 1.56 3.44 2.76.907 1.2 1.454 2.6 1.64 4.2l-5.12.28Zm11.443 8.16V38h-5.6v-5.32h5.6Z"/><path fill="#171717" fill-rule="evenodd" d="m7.839 40.783 16.03-28.054L20 6 0 40.783h7.839Zm8.214 0H40L27.99 19.894l-4.02 7.032 3.976 6.914H20.02l-3.967 6.943Z" clip-rule="evenodd"/></svg>
```

## File: /Users/saint/Desktop/pdf-player/public/placeholder-logo.png

```png
�PNG

   IHDR      �   M��   0PLTE                                                Z?   tRNS � �@��`P0p���w  �IDATx��ؽJ3Q�7'��%�|?� ���E�l�7���(X�D������w`����[�*t����D���mD�}��4;;�DDDDDDDDDDDD_�_İ��!�y�`�_�:��;Ļ�'|�	��;.I"����3*5����J�1�� �T��FI��	��=��3܃�2~�b���0��U9\��]�4�#w0��Gt\&1�?21,���o!e�m��ĻR�����5�� ؽAJ�9��R)�5�0.FFASaǃ�T�#|�K���I�������1�
M������N"��$����G�V�T���T^^��A�$S��h(�������G]co"J׸^^�'�=���%�	�W�6Ы�W��w�a�߇*�^^�YG�c���`'F����������������^ 5_�,�S�%    IEND�B`�
```

## File: /Users/saint/Desktop/pdf-player/public/placeholder-user.jpg

```jpg
���� JFIF      �� C 
	
		
$ &%# #"(-90(*6+"#2D26;=@@@&0FKE>J9?@=�� C=)#)==================================================��  � � ��             ��                 ��     ـ                                                           |�r4�-�       ̈"x�'�0      �Í��8�H�N�      q�����Q��      �����V�`=       �($q"_��               
�S8�P��0     VFbP��!
Io40     ��[?p #�|�@    !.E�  3��4p    Bq  �Z    s   �                                                          �� C 		       AQR�!1@Ua��02Tq���56cps�� "#$PS�����  ? ��R���,�����
�n�k��n8rZ�����9Vv� �V��ms$9zWʏh�-+@Z�2�PGE������EY9��i�Ͻ�S	��O��Ȕ��_I��W髵�}�����B�ՎT��>%r�[e/,W�D}��D�>b�e>�v�Z�p&�*VS��V�sV�c �:��~K������� C��:��k�'An|ʶ�}\��C� �f����a�;�h��J����q!i=���"�q�NF�IZ�`�wĝ5hAj�	
�RXl䎉�lk���@I�%l��Ն���FDY-����Eq�i����O�I�_�2b�lǋ�Yu��k���AO����٣��ܭ�n��cam�jN�j���VL�}�;��oކ6��շs��,���ք���l�i����l{I�O��(!%J $���n�-@G����n��ܮi!�괁G�:�^��n�g3l�F%���9]�Pq��)�:���@�*ɍmׅ�VLY'�s+�z���V�m �J9��S�_���#��;�����Ǌ!5�#�q\�M@��@�]yz�����A;e�k��@� s�^���G����\�5F��(�� S��Ly���c�i8�����o�8T��i�N��7D����t-� p�3`r� qr�;|�.��bTG��[i H��͚-�
��Oj�H����M�ؒFE�{�3X�n���e� �R3/�~�����
�� a����!�j&@^r�����Y�������l�Z? �7땵��)ki�w��\.�u�����X��\.�u�����X��\.�u�����X��\.�u��p�M����o(N��3�Vg�����Z�s��%�\�]q}d�k\_Y5���MG�����Q��q}d�k\_Y5���MG�����Q��q}d�kV|5���8���//����                ��� ? ��                ��� ? ��
```

## File: /Users/saint/Desktop/pdf-player/public/placeholder.jpg

```jpg
���� JFIF   H H  �� �Exif  MM *                  J       R(       �i       Z       H      H    �       �       �           �� 8Photoshop 3.0 8BIM      8BIM%     ��ُ ��	���B~��    ��           	
�� � s !1"AQ2aq#� �B�R3�$b0�r�C�4��S@%c5�s�PD���&T6d�t�`҄�p�'E7e�Uu��Å��Fv��GVf�	
()*89:HIJWXYZghijwxyz�����������������������������������������������������������        	
�� � � ! 1A0"2Q@3#aBqR4�P$��C�b5S��%`�D�r��c6p&ET�'��	
()*789:FGHIJUVWXYZdefghijstuvwxyz����������������������������������������������������������������������������� C 

	
")$+*($''-2@7-0=0''8L9=CEHIH+6OUNFT@GHE�� C!!E.'.EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE��    �k��  �� ?�� ?��  ?�� 3     !1AQaq��������� 0@P`p���������  ?!���    �� 3   	 !1AQa q𑁡�����0@P`p��������� ?��� ?���  ?���
```

## File: /Users/saint/Desktop/pdf-player/public/placeholder.svg

```svg
<svg xmlns="http://www.w3.org/2000/svg" width="1200" height="1200" fill="none"><rect width="1200" height="1200" fill="#EAEAEA" rx="3"/><g opacity=".5"><g opacity=".5"><path fill="#FAFAFA" d="M600.709 736.5c-75.454 0-136.621-61.167-136.621-136.62 0-75.454 61.167-136.621 136.621-136.621 75.453 0 136.62 61.167 136.62 136.621 0 75.453-61.167 136.62-136.62 136.62Z"/><path stroke="#C9C9C9" stroke-width="2.418" d="M600.709 736.5c-75.454 0-136.621-61.167-136.621-136.62 0-75.454 61.167-136.621 136.621-136.621 75.453 0 136.62 61.167 136.62 136.621 0 75.453-61.167 136.62-136.62 136.62Z"/></g><path stroke="url(#a)" stroke-width="2.418" d="M0-1.209h553.581" transform="scale(1 -1) rotate(45 1163.11 91.165)"/><path stroke="url(#b)" stroke-width="2.418" d="M404.846 598.671h391.726"/><path stroke="url(#c)" stroke-width="2.418" d="M599.5 795.742V404.017"/><path stroke="url(#d)" stroke-width="2.418" d="m795.717 796.597-391.441-391.44"/><path fill="#fff" d="M600.709 656.704c-31.384 0-56.825-25.441-56.825-56.824 0-31.384 25.441-56.825 56.825-56.825 31.383 0 56.824 25.441 56.824 56.825 0 31.383-25.441 56.824-56.824 56.824Z"/><g clip-path="url(#e)"><path fill="#666" fill-rule="evenodd" d="M616.426 586.58h-31.434v16.176l3.553-3.554.531-.531h9.068l.074-.074 8.463-8.463h2.565l7.18 7.181V586.58Zm-15.715 14.654 3.698 3.699 1.283 1.282-2.565 2.565-1.282-1.283-5.2-5.199h-6.066l-5.514 5.514-.073.073v2.876a2.418 2.418 0 0 0 2.418 2.418h26.598a2.418 2.418 0 0 0 2.418-2.418v-8.317l-8.463-8.463-7.181 7.181-.071.072Zm-19.347 5.442v4.085a6.045 6.045 0 0 0 6.046 6.045h26.598a6.044 6.044 0 0 0 6.045-6.045v-7.108l1.356-1.355-1.282-1.283-.074-.073v-17.989h-38.689v23.43l-.146.146.146.147Z" clip-rule="evenodd"/></g><path stroke="#C9C9C9" stroke-width="2.418" d="M600.709 656.704c-31.384 0-56.825-25.441-56.825-56.824 0-31.384 25.441-56.825 56.825-56.825 31.383 0 56.824 25.441 56.824 56.825 0 31.383-25.441 56.824-56.824 56.824Z"/></g><defs><linearGradient id="a" x1="554.061" x2="-.48" y1=".083" y2=".087" gradientUnits="userSpaceOnUse"><stop stop-color="#C9C9C9" stop-opacity="0"/><stop offset=".208" stop-color="#C9C9C9"/><stop offset=".792" stop-color="#C9C9C9"/><stop offset="1" stop-color="#C9C9C9" stop-opacity="0"/></linearGradient><linearGradient id="b" x1="796.912" x2="404.507" y1="599.963" y2="599.965" gradientUnits="userSpaceOnUse"><stop stop-color="#C9C9C9" stop-opacity="0"/><stop offset=".208" stop-color="#C9C9C9"/><stop offset=".792" stop-color="#C9C9C9"/><stop offset="1" stop-color="#C9C9C9" stop-opacity="0"/></linearGradient><linearGradient id="c" x1="600.792" x2="600.794" y1="403.677" y2="796.082" gradientUnits="userSpaceOnUse"><stop stop-color="#C9C9C9" stop-opacity="0"/><stop offset=".208" stop-color="#C9C9C9"/><stop offset=".792" stop-color="#C9C9C9"/><stop offset="1" stop-color="#C9C9C9" stop-opacity="0"/></linearGradient><linearGradient id="d" x1="404.85" x2="796.972" y1="403.903" y2="796.02" gradientUnits="userSpaceOnUse"><stop stop-color="#C9C9C9" stop-opacity="0"/><stop offset=".208" stop-color="#C9C9C9"/><stop offset=".792" stop-color="#C9C9C9"/><stop offset="1" stop-color="#C9C9C9" stop-opacity="0"/></linearGradient><clipPath id="e"><path fill="#fff" d="M581.364 580.535h38.689v38.689h-38.689z"/></clipPath></defs></svg>
```

## File: /Users/saint/Desktop/pdf-player/package-lock.json

```json
{
  "name": "my-v0-project",
  "version": "0.1.0",
  "lockfileVersion": 3,
  "requires": true,
  "packages": {
    "": {
      "name": "my-v0-project",
      "version": "0.1.0",
      "dependencies": {
        "@hookform/resolvers": "^3.9.1",
        "@radix-ui/react-accordion": "^1.2.2",
        "@radix-ui/react-alert-dialog": "^1.1.4",
        "@radix-ui/react-aspect-ratio": "^1.1.1",
        "@radix-ui/react-avatar": "^1.1.2",
        "@radix-ui/react-checkbox": "^1.1.3",
        "@radix-ui/react-collapsible": "^1.1.2",
        "@radix-ui/react-context-menu": "^2.2.4",
        "@radix-ui/react-dialog": "^1.1.4",
        "@radix-ui/react-dropdown-menu": "^2.1.4",
        "@radix-ui/react-hover-card": "^1.1.4",
        "@radix-ui/react-label": "^2.1.1",
        "@radix-ui/react-menubar": "^1.1.4",
        "@radix-ui/react-navigation-menu": "^1.2.3",
        "@radix-ui/react-popover": "^1.1.4",
        "@radix-ui/react-progress": "^1.1.1",
        "@radix-ui/react-radio-group": "^1.2.2",
        "@radix-ui/react-scroll-area": "^1.2.2",
        "@radix-ui/react-select": "^2.1.4",
        "@radix-ui/react-separator": "^1.1.1",
        "@radix-ui/react-slider": "^1.2.2",
        "@radix-ui/react-slot": "^1.1.1",
        "@radix-ui/react-switch": "^1.1.2",
        "@radix-ui/react-tabs": "^1.1.2",
        "@radix-ui/react-toast": "^1.2.4",
        "@radix-ui/react-toggle": "^1.1.1",
        "@radix-ui/react-toggle-group": "^1.1.1",
        "@radix-ui/react-tooltip": "^1.1.6",
        "@uploadthing/react": "^7.1.5",
        "autoprefixer": "^10.4.20",
        "class-variance-authority": "^0.7.1",
        "clsx": "^2.1.1",
        "cmdk": "1.0.4",
        "date-fns": "4.1.0",
        "embla-carousel-react": "8.5.1",
        "framer-motion": "latest",
        "input-otp": "1.4.1",
        "lucide-react": "^0.454.0",
        "next": "14.2.16",
        "next-themes": "latest",
        "react": "^18",
        "react-day-picker": "8.10.1",
        "react-dom": "^18",
        "react-hook-form": "^7.54.1",
        "react-resizable-panels": "^2.1.7",
        "recharts": "2.15.0",
        "sonner": "^1.7.1",
        "tailwind-merge": "^2.5.5",
        "tailwindcss-animate": "^1.0.7",
        "uploadthing": "^7.4.4",
        "vaul": "^0.9.6",
        "zod": "^3.24.1"
      },
      "devDependencies": {
        "@types/node": "^22",
        "@types/react": "^18",
        "@types/react-dom": "^18",
        "postcss": "^8",
        "tailwindcss": "^3.4.17",
        "typescript": "^5"
      }
    },
    "node_modules/@alloc/quick-lru": {
      "version": "5.2.0",
      "resolved": "https://registry.npmjs.org/@alloc/quick-lru/-/quick-lru-5.2.0.tgz",
      "integrity": "sha512-UrcABB+4bUrFABwbluTIBErXwvbsU/V7TZWfmbgJfbkwiBuziS9gxdODUyuiecfdGQ85jglMW6juS3+z5TsKLw==",
      "engines": {
        "node": ">=10"
      },
      "funding": {
        "url": "https://github.com/sponsors/sindresorhus"
      }
    },
    "node_modules/@babel/runtime": {
      "version": "7.26.0",
      "resolved": "https://registry.npmjs.org/@babel/runtime/-/runtime-7.26.0.tgz",
      "integrity": "sha512-FDSOghenHTiToteC/QRlv2q3DhPZ/oOXTBoirfWNx1Cx3TMVcGWQtMMmQcSvb/JjpNeGzx8Pq/b4fKEJuWm1sw==",
      "dependencies": {
        "regenerator-runtime": "^0.14.0"
      },
      "engines": {
        "node": ">=6.9.0"
      }
    },
    "node_modules/@effect/platform": {
      "version": "0.70.7",
      "resolved": "https://registry.npmjs.org/@effect/platform/-/platform-0.70.7.tgz",
      "integrity": "sha512-TbNwj/mOJhycPygbmicGBS7CNtv5Z8WVheRbLUdP3oPAe/nbSOJVLc8ZPvOejhquF/1vJMKuqY5MWfkcBpvi/g==",
      "dependencies": {
        "find-my-way-ts": "^0.1.5",
        "multipasta": "^0.2.5"
      },
      "peerDependencies": {
        "effect": "^3.11.5"
      }
    },
    "node_modules/@floating-ui/core": {
      "version": "1.6.9",
      "resolved": "https://registry.npmjs.org/@floating-ui/core/-/core-1.6.9.tgz",
      "integrity": "sha512-uMXCuQ3BItDUbAMhIXw7UPXRfAlOAvZzdK9BWpE60MCn+Svt3aLn9jsPTi/WNGlRUu2uI0v5S7JiIUsbsvh3fw==",
      "dependencies": {
        "@floating-ui/utils": "^0.2.9"
      }
    },
    "node_modules/@floating-ui/dom": {
      "version": "1.6.13",
      "resolved": "https://registry.npmjs.org/@floating-ui/dom/-/dom-1.6.13.tgz",
      "integrity": "sha512-umqzocjDgNRGTuO7Q8CU32dkHkECqI8ZdMZ5Swb6QAM0t5rnlrN3lGo1hdpscRd3WS8T6DKYK4ephgIH9iRh3w==",
      "dependencies": {
        "@floating-ui/core": "^1.6.0",
        "@floating-ui/utils": "^0.2.9"
      }
    },
    "node_modules/@floating-ui/react-dom": {
      "version": "2.1.2",
      "resolved": "https://registry.npmjs.org/@floating-ui/react-dom/-/react-dom-2.1.2.tgz",
      "integrity": "sha512-06okr5cgPzMNBy+Ycse2A6udMi4bqwW/zgBF/rwjcNqWkyr82Mcg8b0vjX8OJpZFy/FKjJmw6wV7t44kK6kW7A==",
      "dependencies": {
        "@floating-ui/dom": "^1.0.0"
      },
      "peerDependencies": {
        "react": ">=16.8.0",
        "react-dom": ">=16.8.0"
      }
    },
    "node_modules/@floating-ui/utils": {
      "version": "0.2.9",
      "resolved": "https://registry.npmjs.org/@floating-ui/utils/-/utils-0.2.9.tgz",
      "integrity": "sha512-MDWhGtE+eHw5JW7lq4qhc5yRLS11ERl1c7Z6Xd0a58DozHES6EnNNwUWbMiG4J9Cgj053Bhk8zvlhFYKVhULwg=="
    },
    "node_modules/@hookform/resolvers": {
      "version": "3.10.0",
      "resolved": "https://registry.npmjs.org/@hookform/resolvers/-/resolvers-3.10.0.tgz",
      "integrity": "sha512-79Dv+3mDF7i+2ajj7SkypSKHhl1cbln1OGavqrsF7p6mbUv11xpqpacPsGDCTRvCSjEEIez2ef1NveSVL3b0Ag==",
      "peerDependencies": {
        "react-hook-form": "^7.0.0"
      }
    },
    "node_modules/@isaacs/cliui": {
      "version": "8.0.2",
      "resolved": "https://registry.npmjs.org/@isaacs/cliui/-/cliui-8.0.2.tgz",
      "integrity": "sha512-O8jcjabXaleOG9DQ0+ARXWZBTfnP4WNAqzuiJK7ll44AmxGKv/J2M4TPjxjY3znBCfvBXFzucm1twdyFybFqEA==",
      "dependencies": {
        "string-width": "^5.1.2",
        "string-width-cjs": "npm:string-width@^4.2.0",
        "strip-ansi": "^7.0.1",
        "strip-ansi-cjs": "npm:strip-ansi@^6.0.1",
        "wrap-ansi": "^8.1.0",
        "wrap-ansi-cjs": "npm:wrap-ansi@^7.0.0"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/@jridgewell/gen-mapping": {
      "version": "0.3.8",
      "resolved": "https://registry.npmjs.org/@jridgewell/gen-mapping/-/gen-mapping-0.3.8.tgz",
      "integrity": "sha512-imAbBGkb+ebQyxKgzv5Hu2nmROxoDOXHh80evxdoXNOrvAnVx7zimzc1Oo5h9RlfV4vPXaE2iM5pOFbvOCClWA==",
      "dependencies": {
        "@jridgewell/set-array": "^1.2.1",
        "@jridgewell/sourcemap-codec": "^1.4.10",
        "@jridgewell/trace-mapping": "^0.3.24"
      },
      "engines": {
        "node": ">=6.0.0"
      }
    },
    "node_modules/@jridgewell/resolve-uri": {
      "version": "3.1.2",
      "resolved": "https://registry.npmjs.org/@jridgewell/resolve-uri/-/resolve-uri-3.1.2.tgz",
      "integrity": "sha512-bRISgCIjP20/tbWSPWMEi54QVPRZExkuD9lJL+UIxUKtwVJA8wW1Trb1jMs1RFXo1CBTNZ/5hpC9QvmKWdopKw==",
      "engines": {
        "node": ">=6.0.0"
      }
    },
    "node_modules/@jridgewell/set-array": {
      "version": "1.2.1",
      "resolved": "https://registry.npmjs.org/@jridgewell/set-array/-/set-array-1.2.1.tgz",
      "integrity": "sha512-R8gLRTZeyp03ymzP/6Lil/28tGeGEzhx1q2k703KGWRAI1VdvPIXdG70VJc2pAMw3NA6JKL5hhFu1sJX0Mnn/A==",
      "engines": {
        "node": ">=6.0.0"
      }
    },
    "node_modules/@jridgewell/sourcemap-codec": {
      "version": "1.5.0",
      "resolved": "https://registry.npmjs.org/@jridgewell/sourcemap-codec/-/sourcemap-codec-1.5.0.tgz",
      "integrity": "sha512-gv3ZRaISU3fjPAgNsriBRqGWQL6quFx04YMPW/zD8XMLsU32mhCCbfbO6KZFLjvYpCZ8zyDEgqsgf+PwPaM7GQ=="
    },
    "node_modules/@jridgewell/trace-mapping": {
      "version": "0.3.25",
      "resolved": "https://registry.npmjs.org/@jridgewell/trace-mapping/-/trace-mapping-0.3.25.tgz",
      "integrity": "sha512-vNk6aEwybGtawWmy/PzwnGDOjCkLWSD2wqvjGGAgOAwCGWySYXfYoxt00IJkTF+8Lb57DwOb3Aa0o9CApepiYQ==",
      "dependencies": {
        "@jridgewell/resolve-uri": "^3.1.0",
        "@jridgewell/sourcemap-codec": "^1.4.14"
      }
    },
    "node_modules/@next/env": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/env/-/env-14.2.16.tgz",
      "integrity": "sha512-fLrX5TfJzHCbnZ9YUSnGW63tMV3L4nSfhgOQ0iCcX21Pt+VSTDuaLsSuL8J/2XAiVA5AnzvXDpf6pMs60QxOag=="
    },
    "node_modules/@next/swc-darwin-arm64": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-darwin-arm64/-/swc-darwin-arm64-14.2.16.tgz",
      "integrity": "sha512-uFT34QojYkf0+nn6MEZ4gIWQ5aqGF11uIZ1HSxG+cSbj+Mg3+tYm8qXYd3dKN5jqKUm5rBVvf1PBRO/MeQ6rxw==",
      "cpu": [
        "arm64"
      ],
      "optional": true,
      "os": [
        "darwin"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-darwin-x64": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-darwin-x64/-/swc-darwin-x64-14.2.16.tgz",
      "integrity": "sha512-mCecsFkYezem0QiZlg2bau3Xul77VxUD38b/auAjohMA22G9KTJneUYMv78vWoCCFkleFAhY1NIvbyjj1ncG9g==",
      "cpu": [
        "x64"
      ],
      "optional": true,
      "os": [
        "darwin"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-linux-arm64-gnu": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-linux-arm64-gnu/-/swc-linux-arm64-gnu-14.2.16.tgz",
      "integrity": "sha512-yhkNA36+ECTC91KSyZcgWgKrYIyDnXZj8PqtJ+c2pMvj45xf7y/HrgI17hLdrcYamLfVt7pBaJUMxADtPaczHA==",
      "cpu": [
        "arm64"
      ],
      "optional": true,
      "os": [
        "linux"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-linux-arm64-musl": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-linux-arm64-musl/-/swc-linux-arm64-musl-14.2.16.tgz",
      "integrity": "sha512-X2YSyu5RMys8R2lA0yLMCOCtqFOoLxrq2YbazFvcPOE4i/isubYjkh+JCpRmqYfEuCVltvlo+oGfj/b5T2pKUA==",
      "cpu": [
        "arm64"
      ],
      "optional": true,
      "os": [
        "linux"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-linux-x64-gnu": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-linux-x64-gnu/-/swc-linux-x64-gnu-14.2.16.tgz",
      "integrity": "sha512-9AGcX7VAkGbc5zTSa+bjQ757tkjr6C/pKS7OK8cX7QEiK6MHIIezBLcQ7gQqbDW2k5yaqba2aDtaBeyyZh1i6Q==",
      "cpu": [
        "x64"
      ],
      "optional": true,
      "os": [
        "linux"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-linux-x64-musl": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-linux-x64-musl/-/swc-linux-x64-musl-14.2.16.tgz",
      "integrity": "sha512-Klgeagrdun4WWDaOizdbtIIm8khUDQJ/5cRzdpXHfkbY91LxBXeejL4kbZBrpR/nmgRrQvmz4l3OtttNVkz2Sg==",
      "cpu": [
        "x64"
      ],
      "optional": true,
      "os": [
        "linux"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-win32-arm64-msvc": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-win32-arm64-msvc/-/swc-win32-arm64-msvc-14.2.16.tgz",
      "integrity": "sha512-PwW8A1UC1Y0xIm83G3yFGPiOBftJK4zukTmk7DI1CebyMOoaVpd8aSy7K6GhobzhkjYvqS/QmzcfsWG2Dwizdg==",
      "cpu": [
        "arm64"
      ],
      "optional": true,
      "os": [
        "win32"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-win32-ia32-msvc": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-win32-ia32-msvc/-/swc-win32-ia32-msvc-14.2.16.tgz",
      "integrity": "sha512-jhPl3nN0oKEshJBNDAo0etGMzv0j3q3VYorTSFqH1o3rwv1MQRdor27u1zhkgsHPNeY1jxcgyx1ZsCkDD1IHgg==",
      "cpu": [
        "ia32"
      ],
      "optional": true,
      "os": [
        "win32"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@next/swc-win32-x64-msvc": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/@next/swc-win32-x64-msvc/-/swc-win32-x64-msvc-14.2.16.tgz",
      "integrity": "sha512-OA7NtfxgirCjfqt+02BqxC3MIgM/JaGjw9tOe4fyZgPsqfseNiMPnCRP44Pfs+Gpo9zPN+SXaFsgP6vk8d571A==",
      "cpu": [
        "x64"
      ],
      "optional": true,
      "os": [
        "win32"
      ],
      "engines": {
        "node": ">= 10"
      }
    },
    "node_modules/@nodelib/fs.scandir": {
      "version": "2.1.5",
      "resolved": "https://registry.npmjs.org/@nodelib/fs.scandir/-/fs.scandir-2.1.5.tgz",
      "integrity": "sha512-vq24Bq3ym5HEQm2NKCr3yXDwjc7vTsEThRDnkp2DK9p1uqLR+DHurm/NOTo0KG7HYHU7eppKZj3MyqYuMBf62g==",
      "dependencies": {
        "@nodelib/fs.stat": "2.0.5",
        "run-parallel": "^1.1.9"
      },
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/@nodelib/fs.stat": {
      "version": "2.0.5",
      "resolved": "https://registry.npmjs.org/@nodelib/fs.stat/-/fs.stat-2.0.5.tgz",
      "integrity": "sha512-RkhPPp2zrqDAQA/2jNhnztcPAlv64XdhIp7a7454A5ovI7Bukxgt7MX7udwAu3zg1DcpPU0rz3VV1SeaqvY4+A==",
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/@nodelib/fs.walk": {
      "version": "1.2.8",
      "resolved": "https://registry.npmjs.org/@nodelib/fs.walk/-/fs.walk-1.2.8.tgz",
      "integrity": "sha512-oGB+UxlgWcgQkgwo8GcEGwemoTFt3FIO9ababBmaGwXIoBKZ+GTy0pP185beGg7Llih/NSHSV2XAs1lnznocSg==",
      "dependencies": {
        "@nodelib/fs.scandir": "2.1.5",
        "fastq": "^1.6.0"
      },
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/@pkgjs/parseargs": {
      "version": "0.11.0",
      "resolved": "https://registry.npmjs.org/@pkgjs/parseargs/-/parseargs-0.11.0.tgz",
      "integrity": "sha512-+1VkjdD0QBLPodGrJUeqarH8VAIvQODIbwh9XpP5Syisf7YoQgsJKPNFoqqLQlu+VQ/tVSshMR6loPMn8U+dPg==",
      "optional": true,
      "engines": {
        "node": ">=14"
      }
    },
    "node_modules/@radix-ui/number": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/number/-/number-1.1.0.tgz",
      "integrity": "sha512-V3gRzhVNU1ldS5XhAPTom1fOIo4ccrjjJgmE+LI2h/WaFpHmx0MQApT+KZHnx8abG6Avtfcz4WoEciMnpFT3HQ=="
    },
    "node_modules/@radix-ui/primitive": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/primitive/-/primitive-1.1.1.tgz",
      "integrity": "sha512-SJ31y+Q/zAyShtXJc8x83i9TYdbAfHZ++tUZnvjJJqFjzsdUnKsxPL6IEtBlxKkU7yzer//GQtZSV4GbldL3YA=="
    },
    "node_modules/@radix-ui/react-accordion": {
      "version": "1.2.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-accordion/-/react-accordion-1.2.2.tgz",
      "integrity": "sha512-b1oh54x4DMCdGsB4/7ahiSrViXxaBwRPotiZNnYXjLha9vfuURSAZErki6qjDoSIV0eXx5v57XnTGVtGwnfp2g==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collapsible": "1.1.2",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-alert-dialog": {
      "version": "1.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-alert-dialog/-/react-alert-dialog-1.1.4.tgz",
      "integrity": "sha512-A6Kh23qZDLy3PSU4bh2UJZznOrUdHImIXqF8YtUa6CN73f8EOO9XlXSCd9IHyPvIquTaa/kwaSWzZTtUvgXVGw==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-dialog": "1.1.4",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-slot": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-arrow": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-arrow/-/react-arrow-1.1.1.tgz",
      "integrity": "sha512-NaVpZfmv8SKeZbn4ijN2V3jlHA9ngBG16VnIIm22nUR0Yk8KUALyBxT3KYEUnNuch9sTE8UTsS3whzBgKOL30w==",
      "dependencies": {
        "@radix-ui/react-primitive": "2.0.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-aspect-ratio": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-aspect-ratio/-/react-aspect-ratio-1.1.1.tgz",
      "integrity": "sha512-kNU4FIpcFMBLkOUcgeIteH06/8JLBcYY6Le1iKenDGCYNYFX3TQqCZjzkOsz37h7r94/99GTb7YhEr98ZBJibw==",
      "dependencies": {
        "@radix-ui/react-primitive": "2.0.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-avatar": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-avatar/-/react-avatar-1.1.2.tgz",
      "integrity": "sha512-GaC7bXQZ5VgZvVvsJ5mu/AEbjYLnhhkoidOboC50Z6FFlLA03wG2ianUoH+zgDQ31/9gCF59bE4+2bBgTyMiig==",
      "dependencies": {
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-checkbox": {
      "version": "1.1.3",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-checkbox/-/react-checkbox-1.1.3.tgz",
      "integrity": "sha512-HD7/ocp8f1B3e6OHygH0n7ZKjONkhciy1Nh0yuBgObqThc3oyx+vuMfFHKAknXRHHWVE9XvXStxJFyjUmB8PIw==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-previous": "1.1.0",
        "@radix-ui/react-use-size": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-collapsible": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-collapsible/-/react-collapsible-1.1.2.tgz",
      "integrity": "sha512-PliMB63vxz7vggcyq0IxNYk8vGDrLXVWw4+W4B8YnwI1s18x7YZYqlG9PLX7XxAJUi0g2DxP4XKJMFHh/iVh9A==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-collection": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-collection/-/react-collection-1.1.1.tgz",
      "integrity": "sha512-LwT3pSho9Dljg+wY2KN2mrrh6y3qELfftINERIzBUO9e0N+t0oMTyn3k9iv+ZqgrwGkRnLpNJrsMv9BZlt2yuA==",
      "dependencies": {
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-slot": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-compose-refs": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-compose-refs/-/react-compose-refs-1.1.1.tgz",
      "integrity": "sha512-Y9VzoRDSJtgFMUCoiZBDVo084VQ5hfpXxVE+NgkdNsjiDBByiImMZKKhxMwCbdHvhlENG6a833CbFkOQvTricw==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-context": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-context/-/react-context-1.1.1.tgz",
      "integrity": "sha512-UASk9zi+crv9WteK/NU4PLvOoL3OuE6BWVKNF6hPRBtYBDXQ2u5iu3O59zUlJiTVvkyuycnqrztsHVJwcK9K+Q==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-context-menu": {
      "version": "2.2.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-context-menu/-/react-context-menu-2.2.4.tgz",
      "integrity": "sha512-ap4wdGwK52rJxGkwukU1NrnEodsUFQIooANKu+ey7d6raQ2biTcEf8za1zr0mgFHieevRTB2nK4dJeN8pTAZGQ==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-menu": "2.1.4",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-dialog": {
      "version": "1.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-dialog/-/react-dialog-1.1.4.tgz",
      "integrity": "sha512-Ur7EV1IwQGCyaAuyDRiOLA5JIUZxELJljF+MbM/2NC0BYwfuRrbpS30BiQBJrVruscgUkieKkqXYDOoByaxIoA==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-focus-guards": "1.1.1",
        "@radix-ui/react-focus-scope": "1.1.1",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-slot": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "aria-hidden": "^1.1.1",
        "react-remove-scroll": "^2.6.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-direction": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-direction/-/react-direction-1.1.0.tgz",
      "integrity": "sha512-BUuBvgThEiAXh2DWu93XsT+a3aWrGqolGlqqw5VU1kG7p/ZH2cuDlM1sRLNnY3QcBS69UIz2mcKhMxDsdewhjg==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-dismissable-layer": {
      "version": "1.1.3",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-dismissable-layer/-/react-dismissable-layer-1.1.3.tgz",
      "integrity": "sha512-onrWn/72lQoEucDmJnr8uczSNTujT0vJnA/X5+3AkChVPowr8n1yvIKIabhWyMQeMvvmdpsvcyDqx3X1LEXCPg==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-escape-keydown": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-dropdown-menu": {
      "version": "2.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-dropdown-menu/-/react-dropdown-menu-2.1.4.tgz",
      "integrity": "sha512-iXU1Ab5ecM+yEepGAWK8ZhMyKX4ubFdCNtol4sT9D0OVErG9PNElfx3TQhjw7n7BC5nFVz68/5//clWy+8TXzA==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-menu": "2.1.4",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-focus-guards": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-focus-guards/-/react-focus-guards-1.1.1.tgz",
      "integrity": "sha512-pSIwfrT1a6sIoDASCSpFwOasEwKTZWDw/iBdtnqKO7v6FeOzYJ7U53cPzYFVR3geGGXgVHaH+CdngrrAzqUGxg==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-focus-scope": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-focus-scope/-/react-focus-scope-1.1.1.tgz",
      "integrity": "sha512-01omzJAYRxXdG2/he/+xy+c8a8gCydoQ1yOxnWNcRhrrBW5W+RQJ22EK1SaO8tb3WoUsuEw7mJjBozPzihDFjA==",
      "dependencies": {
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-hover-card": {
      "version": "1.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-hover-card/-/react-hover-card-1.1.4.tgz",
      "integrity": "sha512-QSUUnRA3PQ2UhvoCv3eYvMnCAgGQW+sTu86QPuNb+ZMi+ZENd6UWpiXbcWDQ4AEaKF9KKpCHBeaJz9Rw6lRlaQ==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-popper": "1.2.1",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-id": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-id/-/react-id-1.1.0.tgz",
      "integrity": "sha512-EJUrI8yYh7WOjNOqpoJaf1jlFIH2LvtgAl+YcFqNCa+4hj64ZXmPkAKOFs/ukjz3byN6bdb/AVUqHkI8/uWWMA==",
      "dependencies": {
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-label": {
      "version": "2.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-label/-/react-label-2.1.1.tgz",
      "integrity": "sha512-UUw5E4e/2+4kFMH7+YxORXGWggtY6sM8WIwh5RZchhLuUg2H1hc98Py+pr8HMz6rdaYrK2t296ZEjYLOCO5uUw==",
      "dependencies": {
        "@radix-ui/react-primitive": "2.0.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-menu": {
      "version": "2.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-menu/-/react-menu-2.1.4.tgz",
      "integrity": "sha512-BnOgVoL6YYdHAG6DtXONaR29Eq4nvbi8rutrV/xlr3RQCMMb3yqP85Qiw/3NReozrSW+4dfLkK+rc1hb4wPU/A==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-focus-guards": "1.1.1",
        "@radix-ui/react-focus-scope": "1.1.1",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-popper": "1.2.1",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-roving-focus": "1.1.1",
        "@radix-ui/react-slot": "1.1.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "aria-hidden": "^1.1.1",
        "react-remove-scroll": "^2.6.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-menubar": {
      "version": "1.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-menubar/-/react-menubar-1.1.4.tgz",
      "integrity": "sha512-+KMpi7VAZuB46+1LD7a30zb5IxyzLgC8m8j42gk3N4TUCcViNQdX8FhoH1HDvYiA8quuqcek4R4bYpPn/SY1GA==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-menu": "2.1.4",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-roving-focus": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-navigation-menu": {
      "version": "1.2.3",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-navigation-menu/-/react-navigation-menu-1.2.3.tgz",
      "integrity": "sha512-IQWAsQ7dsLIYDrn0WqPU+cdM7MONTv9nqrLVYoie3BPiabSfUVDe6Fr+oEt0Cofsr9ONDcDe9xhmJbL1Uq1yKg==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0",
        "@radix-ui/react-use-previous": "1.1.0",
        "@radix-ui/react-visually-hidden": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-popover": {
      "version": "1.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-popover/-/react-popover-1.1.4.tgz",
      "integrity": "sha512-aUACAkXx8LaFymDma+HQVji7WhvEhpFJ7+qPz17Nf4lLZqtreGOFRiNQWQmhzp7kEWg9cOyyQJpdIMUMPc/CPw==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-focus-guards": "1.1.1",
        "@radix-ui/react-focus-scope": "1.1.1",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-popper": "1.2.1",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-slot": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "aria-hidden": "^1.1.1",
        "react-remove-scroll": "^2.6.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-popper": {
      "version": "1.2.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-popper/-/react-popper-1.2.1.tgz",
      "integrity": "sha512-3kn5Me69L+jv82EKRuQCXdYyf1DqHwD2U/sxoNgBGCB7K9TRc3bQamQ+5EPM9EvyPdli0W41sROd+ZU1dTCztw==",
      "dependencies": {
        "@floating-ui/react-dom": "^2.0.0",
        "@radix-ui/react-arrow": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0",
        "@radix-ui/react-use-rect": "1.1.0",
        "@radix-ui/react-use-size": "1.1.0",
        "@radix-ui/rect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-portal": {
      "version": "1.1.3",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-portal/-/react-portal-1.1.3.tgz",
      "integrity": "sha512-NciRqhXnGojhT93RPyDaMPfLH3ZSl4jjIFbZQ1b/vxvZEdHsBZ49wP9w8L3HzUQwep01LcWtkUvm0OVB5JAHTw==",
      "dependencies": {
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-presence": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-presence/-/react-presence-1.1.2.tgz",
      "integrity": "sha512-18TFr80t5EVgL9x1SwF/YGtfG+l0BS0PRAlCWBDoBEiDQjeKgnNZRVJp/oVBl24sr3Gbfwc/Qpj4OcWTQMsAEg==",
      "dependencies": {
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-primitive": {
      "version": "2.0.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-primitive/-/react-primitive-2.0.1.tgz",
      "integrity": "sha512-sHCWTtxwNn3L3fH8qAfnF3WbUZycW93SM1j3NFDzXBiz8D6F5UTTy8G1+WFEaiCdvCVRJWj6N2R4Xq6HdiHmDg==",
      "dependencies": {
        "@radix-ui/react-slot": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-progress": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-progress/-/react-progress-1.1.1.tgz",
      "integrity": "sha512-6diOawA84f/eMxFHcWut0aE1C2kyE9dOyCTQOMRR2C/qPiXz/X0SaiA/RLbapQaXUCmy0/hLMf9meSccD1N0pA==",
      "dependencies": {
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-radio-group": {
      "version": "1.2.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-radio-group/-/react-radio-group-1.2.2.tgz",
      "integrity": "sha512-E0MLLGfOP0l8P/NxgVzfXJ8w3Ch8cdO6UDzJfDChu4EJDy+/WdO5LqpdY8PYnCErkmZH3gZhDL1K7kQ41fAHuQ==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-roving-focus": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-previous": "1.1.0",
        "@radix-ui/react-use-size": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-roving-focus": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-roving-focus/-/react-roving-focus-1.1.1.tgz",
      "integrity": "sha512-QE1RoxPGJ/Nm8Qmk0PxP8ojmoaS67i0s7hVssS7KuI2FQoc/uzVlZsqKfQvxPE6D8hICCPHJ4D88zNhT3OOmkw==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-scroll-area": {
      "version": "1.2.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-scroll-area/-/react-scroll-area-1.2.2.tgz",
      "integrity": "sha512-EFI1N/S3YxZEW/lJ/H1jY3njlvTd8tBmgKEn4GHi51+aMm94i6NmAJstsm5cu3yJwYqYc93gpCPm21FeAbFk6g==",
      "dependencies": {
        "@radix-ui/number": "1.1.0",
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-select": {
      "version": "2.1.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-select/-/react-select-2.1.4.tgz",
      "integrity": "sha512-pOkb2u8KgO47j/h7AylCj7dJsm69BXcjkrvTqMptFqsE2i0p8lHkfgneXKjAgPzBMivnoMyt8o4KiV4wYzDdyQ==",
      "dependencies": {
        "@radix-ui/number": "1.1.0",
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-focus-guards": "1.1.1",
        "@radix-ui/react-focus-scope": "1.1.1",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-popper": "1.2.1",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-slot": "1.1.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0",
        "@radix-ui/react-use-previous": "1.1.0",
        "@radix-ui/react-visually-hidden": "1.1.1",
        "aria-hidden": "^1.1.1",
        "react-remove-scroll": "^2.6.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-separator": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-separator/-/react-separator-1.1.1.tgz",
      "integrity": "sha512-RRiNRSrD8iUiXriq/Y5n4/3iE8HzqgLHsusUSg5jVpU2+3tqcUFPJXHDymwEypunc2sWxDUS3UC+rkZRlHedsw==",
      "dependencies": {
        "@radix-ui/react-primitive": "2.0.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-slider": {
      "version": "1.2.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-slider/-/react-slider-1.2.2.tgz",
      "integrity": "sha512-sNlU06ii1/ZcbHf8I9En54ZPW0Vil/yPVg4vQMcFNjrIx51jsHbFl1HYHQvCIWJSr1q0ZmA+iIs/ZTv8h7HHSA==",
      "dependencies": {
        "@radix-ui/number": "1.1.0",
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0",
        "@radix-ui/react-use-previous": "1.1.0",
        "@radix-ui/react-use-size": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-slot": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-slot/-/react-slot-1.1.1.tgz",
      "integrity": "sha512-RApLLOcINYJA+dMVbOju7MYv1Mb2EBp2nH4HdDzXTSyaR5optlm6Otrz1euW3HbdOR8UmmFK06TD+A9frYWv+g==",
      "dependencies": {
        "@radix-ui/react-compose-refs": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-switch": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-switch/-/react-switch-1.1.2.tgz",
      "integrity": "sha512-zGukiWHjEdBCRyXvKR6iXAQG6qXm2esuAD6kDOi9Cn+1X6ev3ASo4+CsYaD6Fov9r/AQFekqnD/7+V0Cs6/98g==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-previous": "1.1.0",
        "@radix-ui/react-use-size": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-tabs": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-tabs/-/react-tabs-1.1.2.tgz",
      "integrity": "sha512-9u/tQJMcC2aGq7KXpGivMm1mgq7oRJKXphDwdypPd/j21j/2znamPU8WkXgnhUaTrSFNIt8XhOyCAupg8/GbwQ==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-roving-focus": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-toast": {
      "version": "1.2.4",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-toast/-/react-toast-1.2.4.tgz",
      "integrity": "sha512-Sch9idFJHJTMH9YNpxxESqABcAFweJG4tKv+0zo0m5XBvUSL8FM5xKcJLFLXononpePs8IclyX1KieL5SDUNgA==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-collection": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-callback-ref": "1.1.0",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-use-layout-effect": "1.1.0",
        "@radix-ui/react-visually-hidden": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-toggle": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-toggle/-/react-toggle-1.1.1.tgz",
      "integrity": "sha512-i77tcgObYr743IonC1hrsnnPmszDRn8p+EGUsUt+5a/JFn28fxaM88Py6V2mc8J5kELMWishI0rLnuGLFD/nnQ==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-toggle-group": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-toggle-group/-/react-toggle-group-1.1.1.tgz",
      "integrity": "sha512-OgDLZEA30Ylyz8YSXvnGqIHtERqnUt1KUYTKdw/y8u7Ci6zGiJfXc02jahmcSNK3YcErqioj/9flWC9S1ihfwg==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-direction": "1.1.0",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-roving-focus": "1.1.1",
        "@radix-ui/react-toggle": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-tooltip": {
      "version": "1.1.6",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-tooltip/-/react-tooltip-1.1.6.tgz",
      "integrity": "sha512-TLB5D8QLExS1uDn7+wH/bjEmRurNMTzNrtq7IjaS4kjion9NtzsTGkvR5+i7yc9q01Pi2KMM2cN3f8UG4IvvXA==",
      "dependencies": {
        "@radix-ui/primitive": "1.1.1",
        "@radix-ui/react-compose-refs": "1.1.1",
        "@radix-ui/react-context": "1.1.1",
        "@radix-ui/react-dismissable-layer": "1.1.3",
        "@radix-ui/react-id": "1.1.0",
        "@radix-ui/react-popper": "1.2.1",
        "@radix-ui/react-portal": "1.1.3",
        "@radix-ui/react-presence": "1.1.2",
        "@radix-ui/react-primitive": "2.0.1",
        "@radix-ui/react-slot": "1.1.1",
        "@radix-ui/react-use-controllable-state": "1.1.0",
        "@radix-ui/react-visually-hidden": "1.1.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-callback-ref": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-callback-ref/-/react-use-callback-ref-1.1.0.tgz",
      "integrity": "sha512-CasTfvsy+frcFkbXtSJ2Zu9JHpN8TYKxkgJGWbjiZhFivxaeW7rMeZt7QELGVLaYVfFMsKHjb7Ak0nMEe+2Vfw==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-controllable-state": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-controllable-state/-/react-use-controllable-state-1.1.0.tgz",
      "integrity": "sha512-MtfMVJiSr2NjzS0Aa90NPTnvTSg6C/JLCV7ma0W6+OMV78vd8OyRpID+Ng9LxzsPbLeuBnWBA1Nq30AtBIDChw==",
      "dependencies": {
        "@radix-ui/react-use-callback-ref": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-escape-keydown": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-escape-keydown/-/react-use-escape-keydown-1.1.0.tgz",
      "integrity": "sha512-L7vwWlR1kTTQ3oh7g1O0CBF3YCyyTj8NmhLR+phShpyA50HCfBFKVJTpshm9PzLiKmehsrQzTYTpX9HvmC9rhw==",
      "dependencies": {
        "@radix-ui/react-use-callback-ref": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-layout-effect": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-layout-effect/-/react-use-layout-effect-1.1.0.tgz",
      "integrity": "sha512-+FPE0rOdziWSrH9athwI1R0HDVbWlEhd+FR+aSDk4uWGmSJ9Z54sdZVDQPZAinJhJXwfT+qnj969mCsT2gfm5w==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-previous": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-previous/-/react-use-previous-1.1.0.tgz",
      "integrity": "sha512-Z/e78qg2YFnnXcW88A4JmTtm4ADckLno6F7OXotmkQfeuCVaKuYzqAATPhVzl3delXE7CxIV8shofPn3jPc5Og==",
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-rect": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-rect/-/react-use-rect-1.1.0.tgz",
      "integrity": "sha512-0Fmkebhr6PiseyZlYAOtLS+nb7jLmpqTrJyv61Pe68MKYW6OWdRE2kI70TaYY27u7H0lajqM3hSMMLFq18Z7nQ==",
      "dependencies": {
        "@radix-ui/rect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-use-size": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-use-size/-/react-use-size-1.1.0.tgz",
      "integrity": "sha512-XW3/vWuIXHa+2Uwcc2ABSfcCledmXhhQPlGbfcRXbiUQI5Icjcg19BGCZVKKInYbvUCut/ufbbLLPFC5cbb1hw==",
      "dependencies": {
        "@radix-ui/react-use-layout-effect": "1.1.0"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/react-visually-hidden": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/@radix-ui/react-visually-hidden/-/react-visually-hidden-1.1.1.tgz",
      "integrity": "sha512-vVfA2IZ9q/J+gEamvj761Oq1FpWgCDaNOOIfbPVp2MVPLEomUr5+Vf7kJGwQ24YxZSlQVar7Bes8kyTo5Dshpg==",
      "dependencies": {
        "@radix-ui/react-primitive": "2.0.1"
      },
      "peerDependencies": {
        "@types/react": "*",
        "@types/react-dom": "*",
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        },
        "@types/react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/@radix-ui/rect": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/@radix-ui/rect/-/rect-1.1.0.tgz",
      "integrity": "sha512-A9+lCBZoaMJlVKcRBz2YByCG+Cp2t6nAnMnNba+XiWxnj6r4JUFqfsgwocMBZU9LPtdxC6wB56ySYpc7LQIoJg=="
    },
    "node_modules/@standard-schema/spec": {
      "version": "1.0.0-beta.3",
      "resolved": "https://registry.npmjs.org/@standard-schema/spec/-/spec-1.0.0-beta.3.tgz",
      "integrity": "sha512-0ifF3BjA1E8SY9C+nUew8RefNOIq0cDlYALPty4rhUm8Rrl6tCM8hBT4bhGhx7I7iXD0uAgt50lgo8dD73ACMw=="
    },
    "node_modules/@swc/counter": {
      "version": "0.1.3",
      "resolved": "https://registry.npmjs.org/@swc/counter/-/counter-0.1.3.tgz",
      "integrity": "sha512-e2BR4lsJkkRlKZ/qCHPw9ZaSxc0MVUd7gtbtaB7aMvHeJVYe8sOB8DBZkP2DtISHGSku9sCK6T6cnY0CtXrOCQ=="
    },
    "node_modules/@swc/helpers": {
      "version": "0.5.5",
      "resolved": "https://registry.npmjs.org/@swc/helpers/-/helpers-0.5.5.tgz",
      "integrity": "sha512-KGYxvIOXcceOAbEk4bi/dVLEK9z8sZ0uBB3Il5b1rhfClSpcX0yfRO0KmTkqR2cnQDymwLB+25ZyMzICg/cm/A==",
      "dependencies": {
        "@swc/counter": "^0.1.3",
        "tslib": "^2.4.0"
      }
    },
    "node_modules/@types/d3-array": {
      "version": "3.2.1",
      "resolved": "https://registry.npmjs.org/@types/d3-array/-/d3-array-3.2.1.tgz",
      "integrity": "sha512-Y2Jn2idRrLzUfAKV2LyRImR+y4oa2AntrgID95SHJxuMUrkNXmanDSed71sRNZysveJVt1hLLemQZIady0FpEg=="
    },
    "node_modules/@types/d3-color": {
      "version": "3.1.3",
      "resolved": "https://registry.npmjs.org/@types/d3-color/-/d3-color-3.1.3.tgz",
      "integrity": "sha512-iO90scth9WAbmgv7ogoq57O9YpKmFBbmoEoCHDB2xMBY0+/KVrqAaCDyCE16dUspeOvIxFFRI+0sEtqDqy2b4A=="
    },
    "node_modules/@types/d3-ease": {
      "version": "3.0.2",
      "resolved": "https://registry.npmjs.org/@types/d3-ease/-/d3-ease-3.0.2.tgz",
      "integrity": "sha512-NcV1JjO5oDzoK26oMzbILE6HW7uVXOHLQvHshBUW4UMdZGfiY6v5BeQwh9a9tCzv+CeefZQHJt5SRgK154RtiA=="
    },
    "node_modules/@types/d3-interpolate": {
      "version": "3.0.4",
      "resolved": "https://registry.npmjs.org/@types/d3-interpolate/-/d3-interpolate-3.0.4.tgz",
      "integrity": "sha512-mgLPETlrpVV1YRJIglr4Ez47g7Yxjl1lj7YKsiMCb27VJH9W8NVM6Bb9d8kkpG/uAQS5AmbA48q2IAolKKo1MA==",
      "dependencies": {
        "@types/d3-color": "*"
      }
    },
    "node_modules/@types/d3-path": {
      "version": "3.1.0",
      "resolved": "https://registry.npmjs.org/@types/d3-path/-/d3-path-3.1.0.tgz",
      "integrity": "sha512-P2dlU/q51fkOc/Gfl3Ul9kicV7l+ra934qBFXCFhrZMOL6du1TM0pm1ThYvENukyOn5h9v+yMJ9Fn5JK4QozrQ=="
    },
    "node_modules/@types/d3-scale": {
      "version": "4.0.8",
      "resolved": "https://registry.npmjs.org/@types/d3-scale/-/d3-scale-4.0.8.tgz",
      "integrity": "sha512-gkK1VVTr5iNiYJ7vWDI+yUFFlszhNMtVeneJ6lUTKPjprsvLLI9/tgEGiXJOnlINJA8FyA88gfnQsHbybVZrYQ==",
      "dependencies": {
        "@types/d3-time": "*"
      }
    },
    "node_modules/@types/d3-shape": {
      "version": "3.1.7",
      "resolved": "https://registry.npmjs.org/@types/d3-shape/-/d3-shape-3.1.7.tgz",
      "integrity": "sha512-VLvUQ33C+3J+8p+Daf+nYSOsjB4GXp19/S/aGo60m9h1v6XaxjiT82lKVWJCfzhtuZ3yD7i/TPeC/fuKLLOSmg==",
      "dependencies": {
        "@types/d3-path": "*"
      }
    },
    "node_modules/@types/d3-time": {
      "version": "3.0.4",
      "resolved": "https://registry.npmjs.org/@types/d3-time/-/d3-time-3.0.4.tgz",
      "integrity": "sha512-yuzZug1nkAAaBlBBikKZTgzCeA+k1uy4ZFwWANOfKw5z5LRhV0gNA7gNkKm7HoK+HRN0wX3EkxGk0fpbWhmB7g=="
    },
    "node_modules/@types/d3-timer": {
      "version": "3.0.2",
      "resolved": "https://registry.npmjs.org/@types/d3-timer/-/d3-timer-3.0.2.tgz",
      "integrity": "sha512-Ps3T8E8dZDam6fUyNiMkekK3XUsaUEik+idO9/YjPtfj2qruF8tFBXS7XhtE4iIXBLxhmLjP3SXpLhVf21I9Lw=="
    },
    "node_modules/@types/node": {
      "version": "22.10.7",
      "resolved": "https://registry.npmjs.org/@types/node/-/node-22.10.7.tgz",
      "integrity": "sha512-V09KvXxFiutGp6B7XkpaDXlNadZxrzajcY50EuoLIpQ6WWYCSvf19lVIazzfIzQvhUN2HjX12spLojTnhuKlGg==",
      "dev": true,
      "dependencies": {
        "undici-types": "~6.20.0"
      }
    },
    "node_modules/@types/prop-types": {
      "version": "15.7.14",
      "resolved": "https://registry.npmjs.org/@types/prop-types/-/prop-types-15.7.14.tgz",
      "integrity": "sha512-gNMvNH49DJ7OJYv+KAKn0Xp45p8PLl6zo2YnvDIbTd4J6MER2BmWN49TG7n9LvkyihINxeKW8+3bfS2yDC9dzQ==",
      "devOptional": true
    },
    "node_modules/@types/react": {
      "version": "18.3.18",
      "resolved": "https://registry.npmjs.org/@types/react/-/react-18.3.18.tgz",
      "integrity": "sha512-t4yC+vtgnkYjNSKlFx1jkAhH8LgTo2N/7Qvi83kdEaUtMDiwpbLAktKDaAMlRcJ5eSxZkH74eEGt1ky31d7kfQ==",
      "devOptional": true,
      "dependencies": {
        "@types/prop-types": "*",
        "csstype": "^3.0.2"
      }
    },
    "node_modules/@types/react-dom": {
      "version": "18.3.5",
      "resolved": "https://registry.npmjs.org/@types/react-dom/-/react-dom-18.3.5.tgz",
      "integrity": "sha512-P4t6saawp+b/dFrUr2cvkVsfvPguwsxtH6dNIYRllMsefqFzkZk5UIjzyDOv5g1dXIPdG4Sp1yCR4Z6RCUsG/Q==",
      "devOptional": true,
      "peerDependencies": {
        "@types/react": "^18.0.0"
      }
    },
    "node_modules/@uploadthing/mime-types": {
      "version": "0.3.4",
      "resolved": "https://registry.npmjs.org/@uploadthing/mime-types/-/mime-types-0.3.4.tgz",
      "integrity": "sha512-EB0o0a4y++UJFMLqS8LDVhSQgXUllG6t5fwl5cbmOM0Uay8YY5Nc/JjFwzJ8wccdIKz4oYwQW7EOSfUyaMPbfw=="
    },
    "node_modules/@uploadthing/react": {
      "version": "7.1.5",
      "resolved": "https://registry.npmjs.org/@uploadthing/react/-/react-7.1.5.tgz",
      "integrity": "sha512-D0lP4YnUWw5LFHrLuL+sq26wlul9dGE70G+oOQ+qkM1ZqSeKSGvIdAIOAQUBj2qQi4Re7o0VpF3+3qBF3svN6A==",
      "dependencies": {
        "@uploadthing/shared": "7.1.5",
        "file-selector": "0.6.0"
      },
      "peerDependencies": {
        "next": "*",
        "react": "^17.0.2 || ^18.0.0 || ^19.0.0",
        "uploadthing": "^7.2.0"
      },
      "peerDependenciesMeta": {
        "next": {
          "optional": true
        }
      }
    },
    "node_modules/@uploadthing/shared": {
      "version": "7.1.5",
      "resolved": "https://registry.npmjs.org/@uploadthing/shared/-/shared-7.1.5.tgz",
      "integrity": "sha512-5ePSiR2CPA5a8gfPqsHpE/uPwGG9iPqhxcFpeLGnEuExPqIi5ZY7upZgpc+6T1gMxwfbP0O5X4043TTNk3q/pA==",
      "dependencies": {
        "@uploadthing/mime-types": "0.3.4",
        "effect": "3.11.5",
        "sqids": "^0.3.0"
      }
    },
    "node_modules/ansi-regex": {
      "version": "6.1.0",
      "resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-6.1.0.tgz",
      "integrity": "sha512-7HSX4QQb4CspciLpVFwyRe79O3xsIZDDLER21kERQ71oaPodF8jL725AgJMFAYbooIqolJoRLuM81SpeUkpkvA==",
      "engines": {
        "node": ">=12"
      },
      "funding": {
        "url": "https://github.com/chalk/ansi-regex?sponsor=1"
      }
    },
    "node_modules/ansi-styles": {
      "version": "6.2.1",
      "resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-6.2.1.tgz",
      "integrity": "sha512-bN798gFfQX+viw3R7yrGWRqnrN2oRkEkUjjl4JNn4E8GxxbjtG3FbrEIIY3l8/hrwUwIeCZvi4QuOTP4MErVug==",
      "engines": {
        "node": ">=12"
      },
      "funding": {
        "url": "https://github.com/chalk/ansi-styles?sponsor=1"
      }
    },
    "node_modules/any-promise": {
      "version": "1.3.0",
      "resolved": "https://registry.npmjs.org/any-promise/-/any-promise-1.3.0.tgz",
      "integrity": "sha512-7UvmKalWRt1wgjL1RrGxoSJW/0QZFIegpeGvZG9kjp8vrRu55XTHbwnqq2GpXm9uLbcuhxm3IqX9OB4MZR1b2A=="
    },
    "node_modules/anymatch": {
      "version": "3.1.3",
      "resolved": "https://registry.npmjs.org/anymatch/-/anymatch-3.1.3.tgz",
      "integrity": "sha512-KMReFUr0B4t+D+OBkjR3KYqvocp2XaSzO55UcB6mgQMd3KbcE+mWTyvVV7D/zsdEbNnV6acZUutkiHQXvTr1Rw==",
      "dependencies": {
        "normalize-path": "^3.0.0",
        "picomatch": "^2.0.4"
      },
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/arg": {
      "version": "5.0.2",
      "resolved": "https://registry.npmjs.org/arg/-/arg-5.0.2.tgz",
      "integrity": "sha512-PYjyFOLKQ9y57JvQ6QLo8dAgNqswh8M1RMJYdQduT6xbWSgK36P/Z/v+p888pM69jMMfS8Xd8F6I1kQ/I9HUGg=="
    },
    "node_modules/aria-hidden": {
      "version": "1.2.4",
      "resolved": "https://registry.npmjs.org/aria-hidden/-/aria-hidden-1.2.4.tgz",
      "integrity": "sha512-y+CcFFwelSXpLZk/7fMB2mUbGtX9lKycf1MWJ7CaTIERyitVlyQx6C+sxcROU2BAJ24OiZyK+8wj2i8AlBoS3A==",
      "dependencies": {
        "tslib": "^2.0.0"
      },
      "engines": {
        "node": ">=10"
      }
    },
    "node_modules/autoprefixer": {
      "version": "10.4.20",
      "resolved": "https://registry.npmjs.org/autoprefixer/-/autoprefixer-10.4.20.tgz",
      "integrity": "sha512-XY25y5xSv/wEoqzDyXXME4AFfkZI0P23z6Fs3YgymDnKJkCGOnkL0iTxCa85UTqaSgfcqyf3UA6+c7wUvx/16g==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/postcss/"
        },
        {
          "type": "tidelift",
          "url": "https://tidelift.com/funding/github/npm/autoprefixer"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "browserslist": "^4.23.3",
        "caniuse-lite": "^1.0.30001646",
        "fraction.js": "^4.3.7",
        "normalize-range": "^0.1.2",
        "picocolors": "^1.0.1",
        "postcss-value-parser": "^4.2.0"
      },
      "bin": {
        "autoprefixer": "bin/autoprefixer"
      },
      "engines": {
        "node": "^10 || ^12 || >=14"
      },
      "peerDependencies": {
        "postcss": "^8.1.0"
      }
    },
    "node_modules/balanced-match": {
      "version": "1.0.2",
      "resolved": "https://registry.npmjs.org/balanced-match/-/balanced-match-1.0.2.tgz",
      "integrity": "sha512-3oSeUO0TMV67hN1AmbXsK4yaqU7tjiHlbxRDZOpH0KW9+CeX4bRAaX0Anxt0tx2MrpRpWwQaPwIlISEJhYU5Pw=="
    },
    "node_modules/binary-extensions": {
      "version": "2.3.0",
      "resolved": "https://registry.npmjs.org/binary-extensions/-/binary-extensions-2.3.0.tgz",
      "integrity": "sha512-Ceh+7ox5qe7LJuLHoY0feh3pHuUDHAcRUeyL2VYghZwfpkNIy/+8Ocg0a3UuSoYzavmylwuLWQOf3hl0jjMMIw==",
      "engines": {
        "node": ">=8"
      },
      "funding": {
        "url": "https://github.com/sponsors/sindresorhus"
      }
    },
    "node_modules/brace-expansion": {
      "version": "2.0.1",
      "resolved": "https://registry.npmjs.org/brace-expansion/-/brace-expansion-2.0.1.tgz",
      "integrity": "sha512-XnAIvQ8eM+kC6aULx6wuQiwVsnzsi9d3WxzV3FpWTGA19F621kwdbsAcFKXgKUHZWsy+mY6iL1sHTxWEFCytDA==",
      "dependencies": {
        "balanced-match": "^1.0.0"
      }
    },
    "node_modules/braces": {
      "version": "3.0.3",
      "resolved": "https://registry.npmjs.org/braces/-/braces-3.0.3.tgz",
      "integrity": "sha512-yQbXgO/OSZVD2IsiLlro+7Hf6Q18EJrKSEsdoMzKePKXct3gvD8oLcOQdIzGupr5Fj+EDe8gO/lxc1BzfMpxvA==",
      "dependencies": {
        "fill-range": "^7.1.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/browserslist": {
      "version": "4.24.4",
      "resolved": "https://registry.npmjs.org/browserslist/-/browserslist-4.24.4.tgz",
      "integrity": "sha512-KDi1Ny1gSePi1vm0q4oxSF8b4DR44GF4BbmS2YdhPLOEqd8pDviZOGH/GsmRwoWJ2+5Lr085X7naowMwKHDG1A==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/browserslist"
        },
        {
          "type": "tidelift",
          "url": "https://tidelift.com/funding/github/npm/browserslist"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "caniuse-lite": "^1.0.30001688",
        "electron-to-chromium": "^1.5.73",
        "node-releases": "^2.0.19",
        "update-browserslist-db": "^1.1.1"
      },
      "bin": {
        "browserslist": "cli.js"
      },
      "engines": {
        "node": "^6 || ^7 || ^8 || ^9 || ^10 || ^11 || ^12 || >=13.7"
      }
    },
    "node_modules/busboy": {
      "version": "1.6.0",
      "resolved": "https://registry.npmjs.org/busboy/-/busboy-1.6.0.tgz",
      "integrity": "sha512-8SFQbg/0hQ9xy3UNTB0YEnsNBbWfhf7RtnzpL7TkBiTBRfrQ9Fxcnz7VJsleJpyp6rVLvXiuORqjlHi5q+PYuA==",
      "dependencies": {
        "streamsearch": "^1.1.0"
      },
      "engines": {
        "node": ">=10.16.0"
      }
    },
    "node_modules/camelcase-css": {
      "version": "2.0.1",
      "resolved": "https://registry.npmjs.org/camelcase-css/-/camelcase-css-2.0.1.tgz",
      "integrity": "sha512-QOSvevhslijgYwRx6Rv7zKdMF8lbRmx+uQGx2+vDc+KI/eBnsy9kit5aj23AgGu3pa4t9AgwbnXWqS+iOY+2aA==",
      "engines": {
        "node": ">= 6"
      }
    },
    "node_modules/caniuse-lite": {
      "version": "1.0.30001695",
      "resolved": "https://registry.npmjs.org/caniuse-lite/-/caniuse-lite-1.0.30001695.tgz",
      "integrity": "sha512-vHyLade6wTgI2u1ec3WQBxv+2BrTERV28UXQu9LO6lZ9pYeMk34vjXFLOxo1A4UBA8XTL4njRQZdno/yYaSmWw==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/browserslist"
        },
        {
          "type": "tidelift",
          "url": "https://tidelift.com/funding/github/npm/caniuse-lite"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ]
    },
    "node_modules/chokidar": {
      "version": "3.6.0",
      "resolved": "https://registry.npmjs.org/chokidar/-/chokidar-3.6.0.tgz",
      "integrity": "sha512-7VT13fmjotKpGipCW9JEQAusEPE+Ei8nl6/g4FBAmIm0GOOLMua9NDDo/DWp0ZAxCr3cPq5ZpBqmPAQgDda2Pw==",
      "dependencies": {
        "anymatch": "~3.1.2",
        "braces": "~3.0.2",
        "glob-parent": "~5.1.2",
        "is-binary-path": "~2.1.0",
        "is-glob": "~4.0.1",
        "normalize-path": "~3.0.0",
        "readdirp": "~3.6.0"
      },
      "engines": {
        "node": ">= 8.10.0"
      },
      "funding": {
        "url": "https://paulmillr.com/funding/"
      },
      "optionalDependencies": {
        "fsevents": "~2.3.2"
      }
    },
    "node_modules/chokidar/node_modules/glob-parent": {
      "version": "5.1.2",
      "resolved": "https://registry.npmjs.org/glob-parent/-/glob-parent-5.1.2.tgz",
      "integrity": "sha512-AOIgSQCepiJYwP3ARnGx+5VnTu2HBYdzbGP45eLw1vr3zB3vZLeyed1sC9hnbcOc9/SrMyM5RPQrkGz4aS9Zow==",
      "dependencies": {
        "is-glob": "^4.0.1"
      },
      "engines": {
        "node": ">= 6"
      }
    },
    "node_modules/class-variance-authority": {
      "version": "0.7.1",
      "resolved": "https://registry.npmjs.org/class-variance-authority/-/class-variance-authority-0.7.1.tgz",
      "integrity": "sha512-Ka+9Trutv7G8M6WT6SeiRWz792K5qEqIGEGzXKhAE6xOWAY6pPH8U+9IY3oCMv6kqTmLsv7Xh/2w2RigkePMsg==",
      "dependencies": {
        "clsx": "^2.1.1"
      },
      "funding": {
        "url": "https://polar.sh/cva"
      }
    },
    "node_modules/client-only": {
      "version": "0.0.1",
      "resolved": "https://registry.npmjs.org/client-only/-/client-only-0.0.1.tgz",
      "integrity": "sha512-IV3Ou0jSMzZrd3pZ48nLkT9DA7Ag1pnPzaiQhpW7c3RbcqqzvzzVu+L8gfqMp/8IM2MQtSiqaCxrrcfu8I8rMA=="
    },
    "node_modules/clsx": {
      "version": "2.1.1",
      "resolved": "https://registry.npmjs.org/clsx/-/clsx-2.1.1.tgz",
      "integrity": "sha512-eYm0QWBtUrBWZWG0d386OGAw16Z995PiOVo2B7bjWSbHedGl5e0ZWaq65kOGgUSNesEIDkB9ISbTg/JK9dhCZA==",
      "engines": {
        "node": ">=6"
      }
    },
    "node_modules/cmdk": {
      "version": "1.0.4",
      "resolved": "https://registry.npmjs.org/cmdk/-/cmdk-1.0.4.tgz",
      "integrity": "sha512-AnsjfHyHpQ/EFeAnG216WY7A5LiYCoZzCSygiLvfXC3H3LFGCprErteUcszaVluGOhuOTbJS3jWHrSDYPBBygg==",
      "dependencies": {
        "@radix-ui/react-dialog": "^1.1.2",
        "@radix-ui/react-id": "^1.1.0",
        "@radix-ui/react-primitive": "^2.0.0",
        "use-sync-external-store": "^1.2.2"
      },
      "peerDependencies": {
        "react": "^18 || ^19 || ^19.0.0-rc",
        "react-dom": "^18 || ^19 || ^19.0.0-rc"
      }
    },
    "node_modules/color-convert": {
      "version": "2.0.1",
      "resolved": "https://registry.npmjs.org/color-convert/-/color-convert-2.0.1.tgz",
      "integrity": "sha512-RRECPsj7iu/xb5oKYcsFHSppFNnsj/52OVTRKb4zP5onXwVF3zVmmToNcOfGC+CRDpfK/U584fMg38ZHCaElKQ==",
      "dependencies": {
        "color-name": "~1.1.4"
      },
      "engines": {
        "node": ">=7.0.0"
      }
    },
    "node_modules/color-name": {
      "version": "1.1.4",
      "resolved": "https://registry.npmjs.org/color-name/-/color-name-1.1.4.tgz",
      "integrity": "sha512-dOy+3AuW3a2wNbZHIuMZpTcgjGuLU/uBL/ubcZF9OXbDo8ff4O8yVp5Bf0efS8uEoYo5q4Fx7dY9OgQGXgAsQA=="
    },
    "node_modules/commander": {
      "version": "4.1.1",
      "resolved": "https://registry.npmjs.org/commander/-/commander-4.1.1.tgz",
      "integrity": "sha512-NOKm8xhkzAjzFx8B2v5OAHT+u5pRQc2UCa2Vq9jYL/31o2wi9mxBA7LIFs3sV5VSC49z6pEhfbMULvShKj26WA==",
      "engines": {
        "node": ">= 6"
      }
    },
    "node_modules/cross-spawn": {
      "version": "7.0.6",
      "resolved": "https://registry.npmjs.org/cross-spawn/-/cross-spawn-7.0.6.tgz",
      "integrity": "sha512-uV2QOWP2nWzsy2aMp8aRibhi9dlzF5Hgh5SHaB9OiTGEyDTiJJyx0uy51QXdyWbtAHNua4XJzUKca3OzKUd3vA==",
      "dependencies": {
        "path-key": "^3.1.0",
        "shebang-command": "^2.0.0",
        "which": "^2.0.1"
      },
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/cssesc": {
      "version": "3.0.0",
      "resolved": "https://registry.npmjs.org/cssesc/-/cssesc-3.0.0.tgz",
      "integrity": "sha512-/Tb/JcjK111nNScGob5MNtsntNM1aCNUDipB/TkwZFhyDrrE47SOx/18wF2bbjgc3ZzCSKW1T5nt5EbFoAz/Vg==",
      "bin": {
        "cssesc": "bin/cssesc"
      },
      "engines": {
        "node": ">=4"
      }
    },
    "node_modules/csstype": {
      "version": "3.1.3",
      "resolved": "https://registry.npmjs.org/csstype/-/csstype-3.1.3.tgz",
      "integrity": "sha512-M1uQkMl8rQK/szD0LNhtqxIPLpimGm8sOBwU7lLnCpSbTyY3yeU1Vc7l4KT5zT4s/yOxHH5O7tIuuLOCnLADRw=="
    },
    "node_modules/d3-array": {
      "version": "3.2.4",
      "resolved": "https://registry.npmjs.org/d3-array/-/d3-array-3.2.4.tgz",
      "integrity": "sha512-tdQAmyA18i4J7wprpYq8ClcxZy3SC31QMeByyCFyRt7BVHdREQZ5lpzoe5mFEYZUWe+oq8HBvk9JjpibyEV4Jg==",
      "dependencies": {
        "internmap": "1 - 2"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-color": {
      "version": "3.1.0",
      "resolved": "https://registry.npmjs.org/d3-color/-/d3-color-3.1.0.tgz",
      "integrity": "sha512-zg/chbXyeBtMQ1LbD/WSoW2DpC3I0mpmPdW+ynRTj/x2DAWYrIY7qeZIHidozwV24m4iavr15lNwIwLxRmOxhA==",
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-ease": {
      "version": "3.0.1",
      "resolved": "https://registry.npmjs.org/d3-ease/-/d3-ease-3.0.1.tgz",
      "integrity": "sha512-wR/XK3D3XcLIZwpbvQwQ5fK+8Ykds1ip7A2Txe0yxncXSdq1L9skcG7blcedkOX+ZcgxGAmLX1FrRGbADwzi0w==",
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-format": {
      "version": "3.1.0",
      "resolved": "https://registry.npmjs.org/d3-format/-/d3-format-3.1.0.tgz",
      "integrity": "sha512-YyUI6AEuY/Wpt8KWLgZHsIU86atmikuoOmCfommt0LYHiQSPjvX2AcFc38PX0CBpr2RCyZhjex+NS/LPOv6YqA==",
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-interpolate": {
      "version": "3.0.1",
      "resolved": "https://registry.npmjs.org/d3-interpolate/-/d3-interpolate-3.0.1.tgz",
      "integrity": "sha512-3bYs1rOD33uo8aqJfKP3JWPAibgw8Zm2+L9vBKEHJ2Rg+viTR7o5Mmv5mZcieN+FRYaAOWX5SJATX6k1PWz72g==",
      "dependencies": {
        "d3-color": "1 - 3"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-path": {
      "version": "3.1.0",
      "resolved": "https://registry.npmjs.org/d3-path/-/d3-path-3.1.0.tgz",
      "integrity": "sha512-p3KP5HCf/bvjBSSKuXid6Zqijx7wIfNW+J/maPs+iwR35at5JCbLUT0LzF1cnjbCHWhqzQTIN2Jpe8pRebIEFQ==",
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-scale": {
      "version": "4.0.2",
      "resolved": "https://registry.npmjs.org/d3-scale/-/d3-scale-4.0.2.tgz",
      "integrity": "sha512-GZW464g1SH7ag3Y7hXjf8RoUuAFIqklOAq3MRl4OaWabTFJY9PN/E1YklhXLh+OQ3fM9yS2nOkCoS+WLZ6kvxQ==",
      "dependencies": {
        "d3-array": "2.10.0 - 3",
        "d3-format": "1 - 3",
        "d3-interpolate": "1.2.0 - 3",
        "d3-time": "2.1.1 - 3",
        "d3-time-format": "2 - 4"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-shape": {
      "version": "3.2.0",
      "resolved": "https://registry.npmjs.org/d3-shape/-/d3-shape-3.2.0.tgz",
      "integrity": "sha512-SaLBuwGm3MOViRq2ABk3eLoxwZELpH6zhl3FbAoJ7Vm1gofKx6El1Ib5z23NUEhF9AsGl7y+dzLe5Cw2AArGTA==",
      "dependencies": {
        "d3-path": "^3.1.0"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-time": {
      "version": "3.1.0",
      "resolved": "https://registry.npmjs.org/d3-time/-/d3-time-3.1.0.tgz",
      "integrity": "sha512-VqKjzBLejbSMT4IgbmVgDjpkYrNWUYJnbCGo874u7MMKIWsILRX+OpX/gTk8MqjpT1A/c6HY2dCA77ZN0lkQ2Q==",
      "dependencies": {
        "d3-array": "2 - 3"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-time-format": {
      "version": "4.1.0",
      "resolved": "https://registry.npmjs.org/d3-time-format/-/d3-time-format-4.1.0.tgz",
      "integrity": "sha512-dJxPBlzC7NugB2PDLwo9Q8JiTR3M3e4/XANkreKSUxF8vvXKqm1Yfq4Q5dl8budlunRVlUUaDUgFt7eA8D6NLg==",
      "dependencies": {
        "d3-time": "1 - 3"
      },
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/d3-timer": {
      "version": "3.0.1",
      "resolved": "https://registry.npmjs.org/d3-timer/-/d3-timer-3.0.1.tgz",
      "integrity": "sha512-ndfJ/JxxMd3nw31uyKoY2naivF+r29V+Lc0svZxe1JvvIRmi8hUsrMvdOwgS1o6uBHmiz91geQ0ylPP0aj1VUA==",
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/date-fns": {
      "version": "4.1.0",
      "resolved": "https://registry.npmjs.org/date-fns/-/date-fns-4.1.0.tgz",
      "integrity": "sha512-Ukq0owbQXxa/U3EGtsdVBkR1w7KOQ5gIBqdH2hkvknzZPYvBxb/aa6E8L7tmjFtkwZBu3UXBbjIgPo/Ez4xaNg==",
      "funding": {
        "type": "github",
        "url": "https://github.com/sponsors/kossnocorp"
      }
    },
    "node_modules/decimal.js-light": {
      "version": "2.5.1",
      "resolved": "https://registry.npmjs.org/decimal.js-light/-/decimal.js-light-2.5.1.tgz",
      "integrity": "sha512-qIMFpTMZmny+MMIitAB6D7iVPEorVw6YQRWkvarTkT4tBeSLLiHzcwj6q0MmYSFCiVpiqPJTJEYIrpcPzVEIvg=="
    },
    "node_modules/detect-node-es": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/detect-node-es/-/detect-node-es-1.1.0.tgz",
      "integrity": "sha512-ypdmJU/TbBby2Dxibuv7ZLW3Bs1QEmM7nHjEANfohJLvE0XVujisn1qPJcZxg+qDucsr+bP6fLD1rPS3AhJ7EQ=="
    },
    "node_modules/didyoumean": {
      "version": "1.2.2",
      "resolved": "https://registry.npmjs.org/didyoumean/-/didyoumean-1.2.2.tgz",
      "integrity": "sha512-gxtyfqMg7GKyhQmb056K7M3xszy/myH8w+B4RT+QXBQsvAOdc3XymqDDPHx1BgPgsdAA5SIifona89YtRATDzw=="
    },
    "node_modules/dlv": {
      "version": "1.1.3",
      "resolved": "https://registry.npmjs.org/dlv/-/dlv-1.1.3.tgz",
      "integrity": "sha512-+HlytyjlPKnIG8XuRG8WvmBP8xs8P71y+SKKS6ZXWoEgLuePxtDoUEiH7WkdePWrQ5JBpE6aoVqfZfJUQkjXwA=="
    },
    "node_modules/dom-helpers": {
      "version": "5.2.1",
      "resolved": "https://registry.npmjs.org/dom-helpers/-/dom-helpers-5.2.1.tgz",
      "integrity": "sha512-nRCa7CK3VTrM2NmGkIy4cbK7IZlgBE/PYMn55rrXefr5xXDP0LdtfPnblFDoVdcAfslJ7or6iqAUnx0CCGIWQA==",
      "dependencies": {
        "@babel/runtime": "^7.8.7",
        "csstype": "^3.0.2"
      }
    },
    "node_modules/eastasianwidth": {
      "version": "0.2.0",
      "resolved": "https://registry.npmjs.org/eastasianwidth/-/eastasianwidth-0.2.0.tgz",
      "integrity": "sha512-I88TYZWc9XiYHRQ4/3c5rjjfgkjhLyW2luGIheGERbNQ6OY7yTybanSpDXZa8y7VUP9YmDcYa+eyq4ca7iLqWA=="
    },
    "node_modules/effect": {
      "version": "3.11.5",
      "resolved": "https://registry.npmjs.org/effect/-/effect-3.11.5.tgz",
      "integrity": "sha512-oSzaR/S/2A/qDTnDqMWxQUNSjCG2sRLB4NEvTu+l9RqE122MTgKXOWzw0x4MHsdovRTzAihfkpgBj2aLFnH2+w==",
      "dependencies": {
        "fast-check": "^3.21.0"
      }
    },
    "node_modules/electron-to-chromium": {
      "version": "1.5.84",
      "resolved": "https://registry.npmjs.org/electron-to-chromium/-/electron-to-chromium-1.5.84.tgz",
      "integrity": "sha512-I+DQ8xgafao9Ha6y0qjHHvpZ9OfyA1qKlkHkjywxzniORU2awxyz7f/iVJcULmrF2yrM3nHQf+iDjJtbbexd/g=="
    },
    "node_modules/embla-carousel": {
      "version": "8.5.1",
      "resolved": "https://registry.npmjs.org/embla-carousel/-/embla-carousel-8.5.1.tgz",
      "integrity": "sha512-JUb5+FOHobSiWQ2EJNaueCNT/cQU9L6XWBbWmorWPQT9bkbk+fhsuLr8wWrzXKagO3oWszBO7MSx+GfaRk4E6A=="
    },
    "node_modules/embla-carousel-react": {
      "version": "8.5.1",
      "resolved": "https://registry.npmjs.org/embla-carousel-react/-/embla-carousel-react-8.5.1.tgz",
      "integrity": "sha512-z9Y0K84BJvhChXgqn2CFYbfEi6AwEr+FFVVKm/MqbTQ2zIzO1VQri6w67LcfpVF0AjbhwVMywDZqY4alYkjW5w==",
      "dependencies": {
        "embla-carousel": "8.5.1",
        "embla-carousel-reactive-utils": "8.5.1"
      },
      "peerDependencies": {
        "react": "^16.8.0 || ^17.0.1 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      }
    },
    "node_modules/embla-carousel-reactive-utils": {
      "version": "8.5.1",
      "resolved": "https://registry.npmjs.org/embla-carousel-reactive-utils/-/embla-carousel-reactive-utils-8.5.1.tgz",
      "integrity": "sha512-n7VSoGIiiDIc4MfXF3ZRTO59KDp820QDuyBDGlt5/65+lumPHxX2JLz0EZ23hZ4eg4vZGUXwMkYv02fw2JVo/A==",
      "peerDependencies": {
        "embla-carousel": "8.5.1"
      }
    },
    "node_modules/emoji-regex": {
      "version": "9.2.2",
      "resolved": "https://registry.npmjs.org/emoji-regex/-/emoji-regex-9.2.2.tgz",
      "integrity": "sha512-L18DaJsXSUk2+42pv8mLs5jJT2hqFkFE4j21wOmgbUqsZ2hL72NsUU785g9RXgo3s0ZNgVl42TiHp3ZtOv/Vyg=="
    },
    "node_modules/escalade": {
      "version": "3.2.0",
      "resolved": "https://registry.npmjs.org/escalade/-/escalade-3.2.0.tgz",
      "integrity": "sha512-WUj2qlxaQtO4g6Pq5c29GTcWGDyd8itL8zTlipgECz3JesAiiOKotd8JU6otB3PACgG6xkJUyVhboMS+bje/jA==",
      "engines": {
        "node": ">=6"
      }
    },
    "node_modules/eventemitter3": {
      "version": "4.0.7",
      "resolved": "https://registry.npmjs.org/eventemitter3/-/eventemitter3-4.0.7.tgz",
      "integrity": "sha512-8guHBZCwKnFhYdHr2ysuRWErTwhoN2X8XELRlrRwpmfeY2jjuUN4taQMsULKUVo1K4DvZl+0pgfyoysHxvmvEw=="
    },
    "node_modules/fast-check": {
      "version": "3.23.2",
      "resolved": "https://registry.npmjs.org/fast-check/-/fast-check-3.23.2.tgz",
      "integrity": "sha512-h5+1OzzfCC3Ef7VbtKdcv7zsstUQwUDlYpUTvjeUsJAssPgLn7QzbboPtL5ro04Mq0rPOsMzl7q5hIbRs2wD1A==",
      "funding": [
        {
          "type": "individual",
          "url": "https://github.com/sponsors/dubzzz"
        },
        {
          "type": "opencollective",
          "url": "https://opencollective.com/fast-check"
        }
      ],
      "dependencies": {
        "pure-rand": "^6.1.0"
      },
      "engines": {
        "node": ">=8.0.0"
      }
    },
    "node_modules/fast-equals": {
      "version": "5.2.2",
      "resolved": "https://registry.npmjs.org/fast-equals/-/fast-equals-5.2.2.tgz",
      "integrity": "sha512-V7/RktU11J3I36Nwq2JnZEM7tNm17eBJz+u25qdxBZeCKiX6BkVSZQjwWIr+IobgnZy+ag73tTZgZi7tr0LrBw==",
      "engines": {
        "node": ">=6.0.0"
      }
    },
    "node_modules/fast-glob": {
      "version": "3.3.3",
      "resolved": "https://registry.npmjs.org/fast-glob/-/fast-glob-3.3.3.tgz",
      "integrity": "sha512-7MptL8U0cqcFdzIzwOTHoilX9x5BrNqye7Z/LuC7kCMRio1EMSyqRK3BEAUD7sXRq4iT4AzTVuZdhgQ2TCvYLg==",
      "dependencies": {
        "@nodelib/fs.stat": "^2.0.2",
        "@nodelib/fs.walk": "^1.2.3",
        "glob-parent": "^5.1.2",
        "merge2": "^1.3.0",
        "micromatch": "^4.0.8"
      },
      "engines": {
        "node": ">=8.6.0"
      }
    },
    "node_modules/fast-glob/node_modules/glob-parent": {
      "version": "5.1.2",
      "resolved": "https://registry.npmjs.org/glob-parent/-/glob-parent-5.1.2.tgz",
      "integrity": "sha512-AOIgSQCepiJYwP3ARnGx+5VnTu2HBYdzbGP45eLw1vr3zB3vZLeyed1sC9hnbcOc9/SrMyM5RPQrkGz4aS9Zow==",
      "dependencies": {
        "is-glob": "^4.0.1"
      },
      "engines": {
        "node": ">= 6"
      }
    },
    "node_modules/fastq": {
      "version": "1.18.0",
      "resolved": "https://registry.npmjs.org/fastq/-/fastq-1.18.0.tgz",
      "integrity": "sha512-QKHXPW0hD8g4UET03SdOdunzSouc9N4AuHdsX8XNcTsuz+yYFILVNIX4l9yHABMhiEI9Db0JTTIpu0wB+Y1QQw==",
      "dependencies": {
        "reusify": "^1.0.4"
      }
    },
    "node_modules/file-selector": {
      "version": "0.6.0",
      "resolved": "https://registry.npmjs.org/file-selector/-/file-selector-0.6.0.tgz",
      "integrity": "sha512-QlZ5yJC0VxHxQQsQhXvBaC7VRJ2uaxTf+Tfpu4Z/OcVQJVpZO+DGU0rkoVW5ce2SccxugvpBJoMvUs59iILYdw==",
      "dependencies": {
        "tslib": "^2.4.0"
      },
      "engines": {
        "node": ">= 12"
      }
    },
    "node_modules/fill-range": {
      "version": "7.1.1",
      "resolved": "https://registry.npmjs.org/fill-range/-/fill-range-7.1.1.tgz",
      "integrity": "sha512-YsGpe3WHLK8ZYi4tWDg2Jy3ebRz2rXowDxnld4bkQB00cc/1Zw9AWnC0i9ztDJitivtQvaI9KaLyKrc+hBW0yg==",
      "dependencies": {
        "to-regex-range": "^5.0.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/find-my-way-ts": {
      "version": "0.1.5",
      "resolved": "https://registry.npmjs.org/find-my-way-ts/-/find-my-way-ts-0.1.5.tgz",
      "integrity": "sha512-4GOTMrpGQVzsCH2ruUn2vmwzV/02zF4q+ybhCIrw/Rkt3L8KWcycdC6aJMctJzwN4fXD4SD5F/4B9Sksh5rE0A=="
    },
    "node_modules/foreground-child": {
      "version": "3.3.0",
      "resolved": "https://registry.npmjs.org/foreground-child/-/foreground-child-3.3.0.tgz",
      "integrity": "sha512-Ld2g8rrAyMYFXBhEqMz8ZAHBi4J4uS1i/CxGMDnjyFWddMXLVcDp051DZfu+t7+ab7Wv6SMqpWmyFIj5UbfFvg==",
      "dependencies": {
        "cross-spawn": "^7.0.0",
        "signal-exit": "^4.0.1"
      },
      "engines": {
        "node": ">=14"
      },
      "funding": {
        "url": "https://github.com/sponsors/isaacs"
      }
    },
    "node_modules/fraction.js": {
      "version": "4.3.7",
      "resolved": "https://registry.npmjs.org/fraction.js/-/fraction.js-4.3.7.tgz",
      "integrity": "sha512-ZsDfxO51wGAXREY55a7la9LScWpwv9RxIrYABrlvOFBlH/ShPnrtsXeuUIfXKKOVicNxQ+o8JTbJvjS4M89yew==",
      "engines": {
        "node": "*"
      },
      "funding": {
        "type": "patreon",
        "url": "https://github.com/sponsors/rawify"
      }
    },
    "node_modules/framer-motion": {
      "version": "12.0.1",
      "resolved": "https://registry.npmjs.org/framer-motion/-/framer-motion-12.0.1.tgz",
      "integrity": "sha512-u6p0Qc4cY/AEQAtrC7qiYlXla39qnWoI4JXY7OCNBDXwJ5yRBD8HU+RhaOqqziw2m/b0BDh32f44W94+wXonMQ==",
      "dependencies": {
        "motion-dom": "^12.0.0",
        "motion-utils": "^12.0.0",
        "tslib": "^2.4.0"
      },
      "peerDependencies": {
        "@emotion/is-prop-valid": "*",
        "react": "^18.0.0 || ^19.0.0",
        "react-dom": "^18.0.0 || ^19.0.0"
      },
      "peerDependenciesMeta": {
        "@emotion/is-prop-valid": {
          "optional": true
        },
        "react": {
          "optional": true
        },
        "react-dom": {
          "optional": true
        }
      }
    },
    "node_modules/fsevents": {
      "version": "2.3.3",
      "resolved": "https://registry.npmjs.org/fsevents/-/fsevents-2.3.3.tgz",
      "integrity": "sha512-5xoDfX+fL7faATnagmWPpbFtwh/R77WmMMqqHGS65C3vvB0YHrgF+B1YmZ3441tMj5n63k0212XNoJwzlhffQw==",
      "hasInstallScript": true,
      "optional": true,
      "os": [
        "darwin"
      ],
      "engines": {
        "node": "^8.16.0 || ^10.6.0 || >=11.0.0"
      }
    },
    "node_modules/function-bind": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/function-bind/-/function-bind-1.1.2.tgz",
      "integrity": "sha512-7XHNxH7qX9xG5mIwxkhumTox/MIRNcOgDrxWsMt2pAr23WHp6MrRlN7FBSFpCpr+oVO0F744iUgR82nJMfG2SA==",
      "funding": {
        "url": "https://github.com/sponsors/ljharb"
      }
    },
    "node_modules/get-nonce": {
      "version": "1.0.1",
      "resolved": "https://registry.npmjs.org/get-nonce/-/get-nonce-1.0.1.tgz",
      "integrity": "sha512-FJhYRoDaiatfEkUK8HKlicmu/3SGFD51q3itKDGoSTysQJBnfOcxU5GxnhE1E6soB76MbT0MBtnKJuXyAx+96Q==",
      "engines": {
        "node": ">=6"
      }
    },
    "node_modules/glob": {
      "version": "10.4.5",
      "resolved": "https://registry.npmjs.org/glob/-/glob-10.4.5.tgz",
      "integrity": "sha512-7Bv8RF0k6xjo7d4A/PxYLbUCfb6c+Vpd2/mB2yRDlew7Jb5hEXiCD9ibfO7wpk8i4sevK6DFny9h7EYbM3/sHg==",
      "dependencies": {
        "foreground-child": "^3.1.0",
        "jackspeak": "^3.1.2",
        "minimatch": "^9.0.4",
        "minipass": "^7.1.2",
        "package-json-from-dist": "^1.0.0",
        "path-scurry": "^1.11.1"
      },
      "bin": {
        "glob": "dist/esm/bin.mjs"
      },
      "funding": {
        "url": "https://github.com/sponsors/isaacs"
      }
    },
    "node_modules/glob-parent": {
      "version": "6.0.2",
      "resolved": "https://registry.npmjs.org/glob-parent/-/glob-parent-6.0.2.tgz",
      "integrity": "sha512-XxwI8EOhVQgWp6iDL+3b0r86f4d6AX6zSU55HfB4ydCEuXLXc5FcYeOu+nnGftS4TEju/11rt4KJPTMgbfmv4A==",
      "dependencies": {
        "is-glob": "^4.0.3"
      },
      "engines": {
        "node": ">=10.13.0"
      }
    },
    "node_modules/graceful-fs": {
      "version": "4.2.11",
      "resolved": "https://registry.npmjs.org/graceful-fs/-/graceful-fs-4.2.11.tgz",
      "integrity": "sha512-RbJ5/jmFcNNCcDV5o9eTnBLJ/HszWV0P73bc+Ff4nS/rJj+YaS6IGyiOL0VoBYX+l1Wrl3k63h/KrH+nhJ0XvQ=="
    },
    "node_modules/hasown": {
      "version": "2.0.2",
      "resolved": "https://registry.npmjs.org/hasown/-/hasown-2.0.2.tgz",
      "integrity": "sha512-0hJU9SCPvmMzIBdZFqNPXWa6dqh7WdH0cII9y+CyS8rG3nL48Bclra9HmKhVVUHyPWNH5Y7xDwAB7bfgSjkUMQ==",
      "dependencies": {
        "function-bind": "^1.1.2"
      },
      "engines": {
        "node": ">= 0.4"
      }
    },
    "node_modules/input-otp": {
      "version": "1.4.1",
      "resolved": "https://registry.npmjs.org/input-otp/-/input-otp-1.4.1.tgz",
      "integrity": "sha512-+yvpmKYKHi9jIGngxagY9oWiiblPB7+nEO75F2l2o4vs+6vpPZZmUl4tBNYuTCvQjhvEIbdNeJu70bhfYP2nbw==",
      "peerDependencies": {
        "react": "^16.8 || ^17.0 || ^18.0 || ^19.0.0 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17.0 || ^18.0 || ^19.0.0 || ^19.0.0-rc"
      }
    },
    "node_modules/internmap": {
      "version": "2.0.3",
      "resolved": "https://registry.npmjs.org/internmap/-/internmap-2.0.3.tgz",
      "integrity": "sha512-5Hh7Y1wQbvY5ooGgPbDaL5iYLAPzMTUrjMulskHLH6wnv/A+1q5rgEaiuqEjB+oxGXIVZs1FF+R/KPN3ZSQYYg==",
      "engines": {
        "node": ">=12"
      }
    },
    "node_modules/is-binary-path": {
      "version": "2.1.0",
      "resolved": "https://registry.npmjs.org/is-binary-path/-/is-binary-path-2.1.0.tgz",
      "integrity": "sha512-ZMERYes6pDydyuGidse7OsHxtbI7WVeUEozgR/g7rd0xUimYNlvZRE/K2MgZTjWy725IfelLeVcEM97mmtRGXw==",
      "dependencies": {
        "binary-extensions": "^2.0.0"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/is-core-module": {
      "version": "2.16.1",
      "resolved": "https://registry.npmjs.org/is-core-module/-/is-core-module-2.16.1.tgz",
      "integrity": "sha512-UfoeMA6fIJ8wTYFEUjelnaGI67v6+N7qXJEvQuIGa99l4xsCruSYOVSQ0uPANn4dAzm8lkYPaKLrrijLq7x23w==",
      "dependencies": {
        "hasown": "^2.0.2"
      },
      "engines": {
        "node": ">= 0.4"
      },
      "funding": {
        "url": "https://github.com/sponsors/ljharb"
      }
    },
    "node_modules/is-extglob": {
      "version": "2.1.1",
      "resolved": "https://registry.npmjs.org/is-extglob/-/is-extglob-2.1.1.tgz",
      "integrity": "sha512-SbKbANkN603Vi4jEZv49LeVJMn4yGwsbzZworEoyEiutsN3nJYdbO36zfhGJ6QEDpOZIFkDtnq5JRxmvl3jsoQ==",
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/is-fullwidth-code-point": {
      "version": "3.0.0",
      "resolved": "https://registry.npmjs.org/is-fullwidth-code-point/-/is-fullwidth-code-point-3.0.0.tgz",
      "integrity": "sha512-zymm5+u+sCsSWyD9qNaejV3DFvhCKclKdizYaJUuHA83RLjb7nSuGnddCHGv0hk+KY7BMAlsWeK4Ueg6EV6XQg==",
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/is-glob": {
      "version": "4.0.3",
      "resolved": "https://registry.npmjs.org/is-glob/-/is-glob-4.0.3.tgz",
      "integrity": "sha512-xelSayHH36ZgE7ZWhli7pW34hNbNl8Ojv5KVmkJD4hBdD3th8Tfk9vYasLM+mXWOZhFkgZfxhLSnrwRr4elSSg==",
      "dependencies": {
        "is-extglob": "^2.1.1"
      },
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/is-number": {
      "version": "7.0.0",
      "resolved": "https://registry.npmjs.org/is-number/-/is-number-7.0.0.tgz",
      "integrity": "sha512-41Cifkg6e8TylSpdtTpeLVMqvSBEVzTttHvERD741+pnZ8ANv0004MRL43QKPDlK9cGvNp6NZWZUBlbGXYxxng==",
      "engines": {
        "node": ">=0.12.0"
      }
    },
    "node_modules/isexe": {
      "version": "2.0.0",
      "resolved": "https://registry.npmjs.org/isexe/-/isexe-2.0.0.tgz",
      "integrity": "sha512-RHxMLp9lnKHGHRng9QFhRCMbYAcVpn69smSGcq3f36xjgVVWThj4qqLbTLlq7Ssj8B+fIQ1EuCEGI2lKsyQeIw=="
    },
    "node_modules/jackspeak": {
      "version": "3.4.3",
      "resolved": "https://registry.npmjs.org/jackspeak/-/jackspeak-3.4.3.tgz",
      "integrity": "sha512-OGlZQpz2yfahA/Rd1Y8Cd9SIEsqvXkLVoSw/cgwhnhFMDbsQFeZYoJJ7bIZBS9BcamUW96asq/npPWugM+RQBw==",
      "dependencies": {
        "@isaacs/cliui": "^8.0.2"
      },
      "funding": {
        "url": "https://github.com/sponsors/isaacs"
      },
      "optionalDependencies": {
        "@pkgjs/parseargs": "^0.11.0"
      }
    },
    "node_modules/jiti": {
      "version": "1.21.7",
      "resolved": "https://registry.npmjs.org/jiti/-/jiti-1.21.7.tgz",
      "integrity": "sha512-/imKNG4EbWNrVjoNC/1H5/9GFy+tqjGBHCaSsN+P2RnPqjsLmv6UD3Ej+Kj8nBWaRAwyk7kK5ZUc+OEatnTR3A==",
      "bin": {
        "jiti": "bin/jiti.js"
      }
    },
    "node_modules/js-tokens": {
      "version": "4.0.0",
      "resolved": "https://registry.npmjs.org/js-tokens/-/js-tokens-4.0.0.tgz",
      "integrity": "sha512-RdJUflcE3cUzKiMqQgsCu06FPu9UdIJO0beYbPhHN4k6apgJtifcoCtT9bcxOpYBtpD2kCM6Sbzg4CausW/PKQ=="
    },
    "node_modules/lilconfig": {
      "version": "3.1.3",
      "resolved": "https://registry.npmjs.org/lilconfig/-/lilconfig-3.1.3.tgz",
      "integrity": "sha512-/vlFKAoH5Cgt3Ie+JLhRbwOsCQePABiU3tJ1egGvyQ+33R/vcwM2Zl2QR/LzjsBeItPt3oSVXapn+m4nQDvpzw==",
      "engines": {
        "node": ">=14"
      },
      "funding": {
        "url": "https://github.com/sponsors/antonk52"
      }
    },
    "node_modules/lines-and-columns": {
      "version": "1.2.4",
      "resolved": "https://registry.npmjs.org/lines-and-columns/-/lines-and-columns-1.2.4.tgz",
      "integrity": "sha512-7ylylesZQ/PV29jhEDl3Ufjo6ZX7gCqJr5F7PKrqc93v7fzSymt1BpwEU8nAUXs8qzzvqhbjhK5QZg6Mt/HkBg=="
    },
    "node_modules/lodash": {
      "version": "4.17.21",
      "resolved": "https://registry.npmjs.org/lodash/-/lodash-4.17.21.tgz",
      "integrity": "sha512-v2kDEe57lecTulaDIuNTPy3Ry4gLGJ6Z1O3vE1krgXZNrsQ+LFTGHVxVjcXPs17LhbZVGedAJv8XZ1tvj5FvSg=="
    },
    "node_modules/loose-envify": {
      "version": "1.4.0",
      "resolved": "https://registry.npmjs.org/loose-envify/-/loose-envify-1.4.0.tgz",
      "integrity": "sha512-lyuxPGr/Wfhrlem2CL/UcnUc1zcqKAImBDzukY7Y5F/yQiNdko6+fRLevlw1HgMySw7f611UIY408EtxRSoK3Q==",
      "dependencies": {
        "js-tokens": "^3.0.0 || ^4.0.0"
      },
      "bin": {
        "loose-envify": "cli.js"
      }
    },
    "node_modules/lru-cache": {
      "version": "10.4.3",
      "resolved": "https://registry.npmjs.org/lru-cache/-/lru-cache-10.4.3.tgz",
      "integrity": "sha512-JNAzZcXrCt42VGLuYz0zfAzDfAvJWW6AfYlDBQyDV5DClI2m5sAmK+OIO7s59XfsRsWHp02jAJrRadPRGTt6SQ=="
    },
    "node_modules/lucide-react": {
      "version": "0.454.0",
      "resolved": "https://registry.npmjs.org/lucide-react/-/lucide-react-0.454.0.tgz",
      "integrity": "sha512-hw7zMDwykCLnEzgncEEjHeA6+45aeEzRYuKHuyRSOPkhko+J3ySGjGIzu+mmMfDFG1vazHepMaYFYHbTFAZAAQ==",
      "peerDependencies": {
        "react": "^16.5.1 || ^17.0.0 || ^18.0.0 || ^19.0.0-rc"
      }
    },
    "node_modules/merge2": {
      "version": "1.4.1",
      "resolved": "https://registry.npmjs.org/merge2/-/merge2-1.4.1.tgz",
      "integrity": "sha512-8q7VEgMJW4J8tcfVPy8g09NcQwZdbwFEqhe/WZkoIzjn/3TGDwtOCYtXGxA3O8tPzpczCCDgv+P2P5y00ZJOOg==",
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/micromatch": {
      "version": "4.0.8",
      "resolved": "https://registry.npmjs.org/micromatch/-/micromatch-4.0.8.tgz",
      "integrity": "sha512-PXwfBhYu0hBCPw8Dn0E+WDYb7af3dSLVWKi3HGv84IdF4TyFoC0ysxFd0Goxw7nSv4T/PzEJQxsYsEiFCKo2BA==",
      "dependencies": {
        "braces": "^3.0.3",
        "picomatch": "^2.3.1"
      },
      "engines": {
        "node": ">=8.6"
      }
    },
    "node_modules/minimatch": {
      "version": "9.0.5",
      "resolved": "https://registry.npmjs.org/minimatch/-/minimatch-9.0.5.tgz",
      "integrity": "sha512-G6T0ZX48xgozx7587koeX9Ys2NYy6Gmv//P89sEte9V9whIapMNF4idKxnW2QtCcLiTWlb/wfCabAtAFWhhBow==",
      "dependencies": {
        "brace-expansion": "^2.0.1"
      },
      "engines": {
        "node": ">=16 || 14 >=14.17"
      },
      "funding": {
        "url": "https://github.com/sponsors/isaacs"
      }
    },
    "node_modules/minipass": {
      "version": "7.1.2",
      "resolved": "https://registry.npmjs.org/minipass/-/minipass-7.1.2.tgz",
      "integrity": "sha512-qOOzS1cBTWYF4BH8fVePDBOO9iptMnGUEZwNc/cMWnTV2nVLZ7VoNWEPHkYczZA0pdoA7dl6e7FL659nX9S2aw==",
      "engines": {
        "node": ">=16 || 14 >=14.17"
      }
    },
    "node_modules/motion-dom": {
      "version": "12.0.0",
      "resolved": "https://registry.npmjs.org/motion-dom/-/motion-dom-12.0.0.tgz",
      "integrity": "sha512-CvYd15OeIR6kHgMdonCc1ihsaUG4MYh/wrkz8gZ3hBX/uamyZCXN9S9qJoYF03GqfTt7thTV/dxnHYX4+55vDg==",
      "dependencies": {
        "motion-utils": "^12.0.0"
      }
    },
    "node_modules/motion-utils": {
      "version": "12.0.0",
      "resolved": "https://registry.npmjs.org/motion-utils/-/motion-utils-12.0.0.tgz",
      "integrity": "sha512-MNFiBKbbqnmvOjkPyOKgHUp3Q6oiokLkI1bEwm5QA28cxMZrv0CbbBGDNmhF6DIXsi1pCQBSs0dX8xjeER1tmA=="
    },
    "node_modules/multipasta": {
      "version": "0.2.5",
      "resolved": "https://registry.npmjs.org/multipasta/-/multipasta-0.2.5.tgz",
      "integrity": "sha512-c8eMDb1WwZcE02WVjHoOmUVk7fnKU/RmUcosHACglrWAuPQsEJv+E8430sXj6jNc1jHw0zrS16aCjQh4BcEb4A=="
    },
    "node_modules/mz": {
      "version": "2.7.0",
      "resolved": "https://registry.npmjs.org/mz/-/mz-2.7.0.tgz",
      "integrity": "sha512-z81GNO7nnYMEhrGh9LeymoE4+Yr0Wn5McHIZMK5cfQCl+NDX08sCZgUc9/6MHni9IWuFLm1Z3HTCXu2z9fN62Q==",
      "dependencies": {
        "any-promise": "^1.0.0",
        "object-assign": "^4.0.1",
        "thenify-all": "^1.0.0"
      }
    },
    "node_modules/nanoid": {
      "version": "3.3.8",
      "resolved": "https://registry.npmjs.org/nanoid/-/nanoid-3.3.8.tgz",
      "integrity": "sha512-WNLf5Sd8oZxOm+TzppcYk8gVOgP+l58xNy58D0nbUnOxOWRWvlcCV4kUF7ltmI6PsrLl/BgKEyS4mqsGChFN0w==",
      "funding": [
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "bin": {
        "nanoid": "bin/nanoid.cjs"
      },
      "engines": {
        "node": "^10 || ^12 || ^13.7 || ^14 || >=15.0.1"
      }
    },
    "node_modules/next": {
      "version": "14.2.16",
      "resolved": "https://registry.npmjs.org/next/-/next-14.2.16.tgz",
      "integrity": "sha512-LcO7WnFu6lYSvCzZoo1dB+IO0xXz5uEv52HF1IUN0IqVTUIZGHuuR10I5efiLadGt+4oZqTcNZyVVEem/TM5nA==",
      "dependencies": {
        "@next/env": "14.2.16",
        "@swc/helpers": "0.5.5",
        "busboy": "1.6.0",
        "caniuse-lite": "^1.0.30001579",
        "graceful-fs": "^4.2.11",
        "postcss": "8.4.31",
        "styled-jsx": "5.1.1"
      },
      "bin": {
        "next": "dist/bin/next"
      },
      "engines": {
        "node": ">=18.17.0"
      },
      "optionalDependencies": {
        "@next/swc-darwin-arm64": "14.2.16",
        "@next/swc-darwin-x64": "14.2.16",
        "@next/swc-linux-arm64-gnu": "14.2.16",
        "@next/swc-linux-arm64-musl": "14.2.16",
        "@next/swc-linux-x64-gnu": "14.2.16",
        "@next/swc-linux-x64-musl": "14.2.16",
        "@next/swc-win32-arm64-msvc": "14.2.16",
        "@next/swc-win32-ia32-msvc": "14.2.16",
        "@next/swc-win32-x64-msvc": "14.2.16"
      },
      "peerDependencies": {
        "@opentelemetry/api": "^1.1.0",
        "@playwright/test": "^1.41.2",
        "react": "^18.2.0",
        "react-dom": "^18.2.0",
        "sass": "^1.3.0"
      },
      "peerDependenciesMeta": {
        "@opentelemetry/api": {
          "optional": true
        },
        "@playwright/test": {
          "optional": true
        },
        "sass": {
          "optional": true
        }
      }
    },
    "node_modules/next-themes": {
      "version": "0.4.4",
      "resolved": "https://registry.npmjs.org/next-themes/-/next-themes-0.4.4.tgz",
      "integrity": "sha512-LDQ2qIOJF0VnuVrrMSMLrWGjRMkq+0mpgl6e0juCLqdJ+oo8Q84JRWT6Wh11VDQKkMMe+dVzDKLWs5n87T+PkQ==",
      "peerDependencies": {
        "react": "^16.8 || ^17 || ^18 || ^19 || ^19.0.0-rc",
        "react-dom": "^16.8 || ^17 || ^18 || ^19 || ^19.0.0-rc"
      }
    },
    "node_modules/next/node_modules/postcss": {
      "version": "8.4.31",
      "resolved": "https://registry.npmjs.org/postcss/-/postcss-8.4.31.tgz",
      "integrity": "sha512-PS08Iboia9mts/2ygV3eLpY5ghnUcfLV/EXTOW1E2qYxJKGGBUtNjN76FYHnMs36RmARn41bC0AZmn+rR0OVpQ==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/postcss/"
        },
        {
          "type": "tidelift",
          "url": "https://tidelift.com/funding/github/npm/postcss"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "nanoid": "^3.3.6",
        "picocolors": "^1.0.0",
        "source-map-js": "^1.0.2"
      },
      "engines": {
        "node": "^10 || ^12 || >=14"
      }
    },
    "node_modules/node-releases": {
      "version": "2.0.19",
      "resolved": "https://registry.npmjs.org/node-releases/-/node-releases-2.0.19.tgz",
      "integrity": "sha512-xxOWJsBKtzAq7DY0J+DTzuz58K8e7sJbdgwkbMWQe8UYB6ekmsQ45q0M/tJDsGaZmbC+l7n57UV8Hl5tHxO9uw=="
    },
    "node_modules/normalize-path": {
      "version": "3.0.0",
      "resolved": "https://registry.npmjs.org/normalize-path/-/normalize-path-3.0.0.tgz",
      "integrity": "sha512-6eZs5Ls3WtCisHWp9S2GUy8dqkpGi4BVSz3GaqiE6ezub0512ESztXUwUB6C6IKbQkY2Pnb/mD4WYojCRwcwLA==",
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/normalize-range": {
      "version": "0.1.2",
      "resolved": "https://registry.npmjs.org/normalize-range/-/normalize-range-0.1.2.tgz",
      "integrity": "sha512-bdok/XvKII3nUpklnV6P2hxtMNrCboOjAcyBuQnWEhO665FwrSNRxU+AqpsyvO6LgGYPspN+lu5CLtw4jPRKNA==",
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/object-assign": {
      "version": "4.1.1",
      "resolved": "https://registry.npmjs.org/object-assign/-/object-assign-4.1.1.tgz",
      "integrity": "sha512-rJgTQnkUnH1sFw8yT6VSU3zD3sWmu6sZhIseY8VX+GRu3P6F7Fu+JNDoXfklElbLJSnc3FUQHVe4cU5hj+BcUg==",
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/object-hash": {
      "version": "3.0.0",
      "resolved": "https://registry.npmjs.org/object-hash/-/object-hash-3.0.0.tgz",
      "integrity": "sha512-RSn9F68PjH9HqtltsSnqYC1XXoWe9Bju5+213R98cNGttag9q9yAOTzdbsqvIa7aNm5WffBZFpWYr2aWrklWAw==",
      "engines": {
        "node": ">= 6"
      }
    },
    "node_modules/package-json-from-dist": {
      "version": "1.0.1",
      "resolved": "https://registry.npmjs.org/package-json-from-dist/-/package-json-from-dist-1.0.1.tgz",
      "integrity": "sha512-UEZIS3/by4OC8vL3P2dTXRETpebLI2NiI5vIrjaD/5UtrkFX/tNbwjTSRAGC/+7CAo2pIcBaRgWmcBBHcsaCIw=="
    },
    "node_modules/path-key": {
      "version": "3.1.1",
      "resolved": "https://registry.npmjs.org/path-key/-/path-key-3.1.1.tgz",
      "integrity": "sha512-ojmeN0qd+y0jszEtoY48r0Peq5dwMEkIlCOu6Q5f41lfkswXuKtYrhgoTpLnyIcHm24Uhqx+5Tqm2InSwLhE6Q==",
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/path-parse": {
      "version": "1.0.7",
      "resolved": "https://registry.npmjs.org/path-parse/-/path-parse-1.0.7.tgz",
      "integrity": "sha512-LDJzPVEEEPR+y48z93A0Ed0yXb8pAByGWo/k5YYdYgpY2/2EsOsksJrq7lOHxryrVOn1ejG6oAp8ahvOIQD8sw=="
    },
    "node_modules/path-scurry": {
      "version": "1.11.1",
      "resolved": "https://registry.npmjs.org/path-scurry/-/path-scurry-1.11.1.tgz",
      "integrity": "sha512-Xa4Nw17FS9ApQFJ9umLiJS4orGjm7ZzwUrwamcGQuHSzDyth9boKDaycYdDcZDuqYATXw4HFXgaqWTctW/v1HA==",
      "dependencies": {
        "lru-cache": "^10.2.0",
        "minipass": "^5.0.0 || ^6.0.2 || ^7.0.0"
      },
      "engines": {
        "node": ">=16 || 14 >=14.18"
      },
      "funding": {
        "url": "https://github.com/sponsors/isaacs"
      }
    },
    "node_modules/picocolors": {
      "version": "1.1.1",
      "resolved": "https://registry.npmjs.org/picocolors/-/picocolors-1.1.1.tgz",
      "integrity": "sha512-xceH2snhtb5M9liqDsmEw56le376mTZkEX/jEb/RxNFyegNul7eNslCXP9FDj/Lcu0X8KEyMceP2ntpaHrDEVA=="
    },
    "node_modules/picomatch": {
      "version": "2.3.1",
      "resolved": "https://registry.npmjs.org/picomatch/-/picomatch-2.3.1.tgz",
      "integrity": "sha512-JU3teHTNjmE2VCGFzuY8EXzCDVwEqB2a8fsIvwaStHhAWJEeVd1o1QD80CU6+ZdEXXSLbSsuLwJjkCBWqRQUVA==",
      "engines": {
        "node": ">=8.6"
      },
      "funding": {
        "url": "https://github.com/sponsors/jonschlinkert"
      }
    },
    "node_modules/pify": {
      "version": "2.3.0",
      "resolved": "https://registry.npmjs.org/pify/-/pify-2.3.0.tgz",
      "integrity": "sha512-udgsAY+fTnvv7kI7aaxbqwWNb0AHiB0qBO89PZKPkoTmGOgdbrHDKD+0B2X4uTfJ/FT1R09r9gTsjUjNJotuog==",
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/pirates": {
      "version": "4.0.6",
      "resolved": "https://registry.npmjs.org/pirates/-/pirates-4.0.6.tgz",
      "integrity": "sha512-saLsH7WeYYPiD25LDuLRRY/i+6HaPYr6G1OUlN39otzkSTxKnubR9RTxS3/Kk50s1g2JTgFwWQDQyplC5/SHZg==",
      "engines": {
        "node": ">= 6"
      }
    },
    "node_modules/postcss": {
      "version": "8.5.1",
      "resolved": "https://registry.npmjs.org/postcss/-/postcss-8.5.1.tgz",
      "integrity": "sha512-6oz2beyjc5VMn/KV1pPw8fliQkhBXrVn1Z3TVyqZxU8kZpzEKhBdmCFqI6ZbmGtamQvQGuU1sgPTk8ZrXDD7jQ==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/postcss/"
        },
        {
          "type": "tidelift",
          "url": "https://tidelift.com/funding/github/npm/postcss"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "nanoid": "^3.3.8",
        "picocolors": "^1.1.1",
        "source-map-js": "^1.2.1"
      },
      "engines": {
        "node": "^10 || ^12 || >=14"
      }
    },
    "node_modules/postcss-import": {
      "version": "15.1.0",
      "resolved": "https://registry.npmjs.org/postcss-import/-/postcss-import-15.1.0.tgz",
      "integrity": "sha512-hpr+J05B2FVYUAXHeK1YyI267J/dDDhMU6B6civm8hSY1jYJnBXxzKDKDswzJmtLHryrjhnDjqqp/49t8FALew==",
      "dependencies": {
        "postcss-value-parser": "^4.0.0",
        "read-cache": "^1.0.0",
        "resolve": "^1.1.7"
      },
      "engines": {
        "node": ">=14.0.0"
      },
      "peerDependencies": {
        "postcss": "^8.0.0"
      }
    },
    "node_modules/postcss-js": {
      "version": "4.0.1",
      "resolved": "https://registry.npmjs.org/postcss-js/-/postcss-js-4.0.1.tgz",
      "integrity": "sha512-dDLF8pEO191hJMtlHFPRa8xsizHaM82MLfNkUHdUtVEV3tgTp5oj+8qbEqYM57SLfc74KSbw//4SeJma2LRVIw==",
      "dependencies": {
        "camelcase-css": "^2.0.1"
      },
      "engines": {
        "node": "^12 || ^14 || >= 16"
      },
      "funding": {
        "type": "opencollective",
        "url": "https://opencollective.com/postcss/"
      },
      "peerDependencies": {
        "postcss": "^8.4.21"
      }
    },
    "node_modules/postcss-load-config": {
      "version": "4.0.2",
      "resolved": "https://registry.npmjs.org/postcss-load-config/-/postcss-load-config-4.0.2.tgz",
      "integrity": "sha512-bSVhyJGL00wMVoPUzAVAnbEoWyqRxkjv64tUl427SKnPrENtq6hJwUojroMz2VB+Q1edmi4IfrAPpami5VVgMQ==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/postcss/"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "lilconfig": "^3.0.0",
        "yaml": "^2.3.4"
      },
      "engines": {
        "node": ">= 14"
      },
      "peerDependencies": {
        "postcss": ">=8.0.9",
        "ts-node": ">=9.0.0"
      },
      "peerDependenciesMeta": {
        "postcss": {
          "optional": true
        },
        "ts-node": {
          "optional": true
        }
      }
    },
    "node_modules/postcss-nested": {
      "version": "6.2.0",
      "resolved": "https://registry.npmjs.org/postcss-nested/-/postcss-nested-6.2.0.tgz",
      "integrity": "sha512-HQbt28KulC5AJzG+cZtj9kvKB93CFCdLvog1WFLf1D+xmMvPGlBstkpTEZfK5+AN9hfJocyBFCNiqyS48bpgzQ==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/postcss/"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "postcss-selector-parser": "^6.1.1"
      },
      "engines": {
        "node": ">=12.0"
      },
      "peerDependencies": {
        "postcss": "^8.2.14"
      }
    },
    "node_modules/postcss-selector-parser": {
      "version": "6.1.2",
      "resolved": "https://registry.npmjs.org/postcss-selector-parser/-/postcss-selector-parser-6.1.2.tgz",
      "integrity": "sha512-Q8qQfPiZ+THO/3ZrOrO0cJJKfpYCagtMUkXbnEfmgUjwXg6z/WBeOyS9APBBPCTSiDV+s4SwQGu8yFsiMRIudg==",
      "dependencies": {
        "cssesc": "^3.0.0",
        "util-deprecate": "^1.0.2"
      },
      "engines": {
        "node": ">=4"
      }
    },
    "node_modules/postcss-value-parser": {
      "version": "4.2.0",
      "resolved": "https://registry.npmjs.org/postcss-value-parser/-/postcss-value-parser-4.2.0.tgz",
      "integrity": "sha512-1NNCs6uurfkVbeXG4S8JFT9t19m45ICnif8zWLd5oPSZ50QnwMfK+H3jv408d4jw/7Bttv5axS5IiHoLaVNHeQ=="
    },
    "node_modules/prop-types": {
      "version": "15.8.1",
      "resolved": "https://registry.npmjs.org/prop-types/-/prop-types-15.8.1.tgz",
      "integrity": "sha512-oj87CgZICdulUohogVAR7AjlC0327U4el4L6eAvOqCeudMDVU0NThNaV+b9Df4dXgSP1gXMTnPdhfe/2qDH5cg==",
      "dependencies": {
        "loose-envify": "^1.4.0",
        "object-assign": "^4.1.1",
        "react-is": "^16.13.1"
      }
    },
    "node_modules/prop-types/node_modules/react-is": {
      "version": "16.13.1",
      "resolved": "https://registry.npmjs.org/react-is/-/react-is-16.13.1.tgz",
      "integrity": "sha512-24e6ynE2H+OKt4kqsOvNd8kBpV65zoxbA4BVsEOB3ARVWQki/DHzaUoC5KuON/BiccDaCCTZBuOcfZs70kR8bQ=="
    },
    "node_modules/pure-rand": {
      "version": "6.1.0",
      "resolved": "https://registry.npmjs.org/pure-rand/-/pure-rand-6.1.0.tgz",
      "integrity": "sha512-bVWawvoZoBYpp6yIoQtQXHZjmz35RSVHnUOTefl8Vcjr8snTPY1wnpSPMWekcFwbxI6gtmT7rSYPFvz71ldiOA==",
      "funding": [
        {
          "type": "individual",
          "url": "https://github.com/sponsors/dubzzz"
        },
        {
          "type": "opencollective",
          "url": "https://opencollective.com/fast-check"
        }
      ]
    },
    "node_modules/queue-microtask": {
      "version": "1.2.3",
      "resolved": "https://registry.npmjs.org/queue-microtask/-/queue-microtask-1.2.3.tgz",
      "integrity": "sha512-NuaNSa6flKT5JaSYQzJok04JzTL1CA6aGhv5rfLW3PgqA+M2ChpZQnAC8h8i4ZFkBS8X5RqkDBHA7r4hej3K9A==",
      "funding": [
        {
          "type": "github",
          "url": "https://github.com/sponsors/feross"
        },
        {
          "type": "patreon",
          "url": "https://www.patreon.com/feross"
        },
        {
          "type": "consulting",
          "url": "https://feross.org/support"
        }
      ]
    },
    "node_modules/react": {
      "version": "18.3.1",
      "resolved": "https://registry.npmjs.org/react/-/react-18.3.1.tgz",
      "integrity": "sha512-wS+hAgJShR0KhEvPJArfuPVN1+Hz1t0Y6n5jLrGQbkb4urgPE/0Rve+1kMB1v/oWgHgm4WIcV+i7F2pTVj+2iQ==",
      "dependencies": {
        "loose-envify": "^1.1.0"
      },
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/react-day-picker": {
      "version": "8.10.1",
      "resolved": "https://registry.npmjs.org/react-day-picker/-/react-day-picker-8.10.1.tgz",
      "integrity": "sha512-TMx7fNbhLk15eqcMt+7Z7S2KF7mfTId/XJDjKE8f+IUcFn0l08/kI4FiYTL/0yuOLmEcbR4Fwe3GJf/NiiMnPA==",
      "funding": {
        "type": "individual",
        "url": "https://github.com/sponsors/gpbl"
      },
      "peerDependencies": {
        "date-fns": "^2.28.0 || ^3.0.0",
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0"
      }
    },
    "node_modules/react-dom": {
      "version": "18.3.1",
      "resolved": "https://registry.npmjs.org/react-dom/-/react-dom-18.3.1.tgz",
      "integrity": "sha512-5m4nQKp+rZRb09LNH59GM4BxTh9251/ylbKIbpe7TpGxfJ+9kv6BLkLBXIjjspbgbnIBNqlI23tRnTWT0snUIw==",
      "dependencies": {
        "loose-envify": "^1.1.0",
        "scheduler": "^0.23.2"
      },
      "peerDependencies": {
        "react": "^18.3.1"
      }
    },
    "node_modules/react-hook-form": {
      "version": "7.54.2",
      "resolved": "https://registry.npmjs.org/react-hook-form/-/react-hook-form-7.54.2.tgz",
      "integrity": "sha512-eHpAUgUjWbZocoQYUHposymRb4ZP6d0uwUnooL2uOybA9/3tPUvoAKqEWK1WaSiTxxOfTpffNZP7QwlnM3/gEg==",
      "engines": {
        "node": ">=18.0.0"
      },
      "funding": {
        "type": "opencollective",
        "url": "https://opencollective.com/react-hook-form"
      },
      "peerDependencies": {
        "react": "^16.8.0 || ^17 || ^18 || ^19"
      }
    },
    "node_modules/react-is": {
      "version": "18.3.1",
      "resolved": "https://registry.npmjs.org/react-is/-/react-is-18.3.1.tgz",
      "integrity": "sha512-/LLMVyas0ljjAtoYiPqYiL8VWXzUUdThrmU5+n20DZv+a+ClRoevUzw5JxU+Ieh5/c87ytoTBV9G1FiKfNJdmg=="
    },
    "node_modules/react-remove-scroll": {
      "version": "2.6.3",
      "resolved": "https://registry.npmjs.org/react-remove-scroll/-/react-remove-scroll-2.6.3.tgz",
      "integrity": "sha512-pnAi91oOk8g8ABQKGF5/M9qxmmOPxaAnopyTHYfqYEwJhyFrbbBtHuSgtKEoH0jpcxx5o3hXqH1mNd9/Oi+8iQ==",
      "dependencies": {
        "react-remove-scroll-bar": "^2.3.7",
        "react-style-singleton": "^2.2.3",
        "tslib": "^2.1.0",
        "use-callback-ref": "^1.3.3",
        "use-sidecar": "^1.1.3"
      },
      "engines": {
        "node": ">=10"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/react-remove-scroll-bar": {
      "version": "2.3.8",
      "resolved": "https://registry.npmjs.org/react-remove-scroll-bar/-/react-remove-scroll-bar-2.3.8.tgz",
      "integrity": "sha512-9r+yi9+mgU33AKcj6IbT9oRCO78WriSj6t/cF8DWBZJ9aOGPOTEDvdUDz1FwKim7QXWwmHqtdHnRJfhAxEG46Q==",
      "dependencies": {
        "react-style-singleton": "^2.2.2",
        "tslib": "^2.0.0"
      },
      "engines": {
        "node": ">=10"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/react-resizable-panels": {
      "version": "2.1.7",
      "resolved": "https://registry.npmjs.org/react-resizable-panels/-/react-resizable-panels-2.1.7.tgz",
      "integrity": "sha512-JtT6gI+nURzhMYQYsx8DKkx6bSoOGFp7A3CwMrOb8y5jFHFyqwo9m68UhmXRw57fRVJksFn1TSlm3ywEQ9vMgA==",
      "peerDependencies": {
        "react": "^16.14.0 || ^17.0.0 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc",
        "react-dom": "^16.14.0 || ^17.0.0 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      }
    },
    "node_modules/react-smooth": {
      "version": "4.0.4",
      "resolved": "https://registry.npmjs.org/react-smooth/-/react-smooth-4.0.4.tgz",
      "integrity": "sha512-gnGKTpYwqL0Iii09gHobNolvX4Kiq4PKx6eWBCYYix+8cdw+cGo3do906l1NBPKkSWx1DghC1dlWG9L2uGd61Q==",
      "dependencies": {
        "fast-equals": "^5.0.1",
        "prop-types": "^15.8.1",
        "react-transition-group": "^4.4.5"
      },
      "peerDependencies": {
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0",
        "react-dom": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0"
      }
    },
    "node_modules/react-style-singleton": {
      "version": "2.2.3",
      "resolved": "https://registry.npmjs.org/react-style-singleton/-/react-style-singleton-2.2.3.tgz",
      "integrity": "sha512-b6jSvxvVnyptAiLjbkWLE/lOnR4lfTtDAl+eUC7RZy+QQWc6wRzIV2CE6xBuMmDxc2qIihtDCZD5NPOFl7fRBQ==",
      "dependencies": {
        "get-nonce": "^1.0.0",
        "tslib": "^2.0.0"
      },
      "engines": {
        "node": ">=10"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/react-transition-group": {
      "version": "4.4.5",
      "resolved": "https://registry.npmjs.org/react-transition-group/-/react-transition-group-4.4.5.tgz",
      "integrity": "sha512-pZcd1MCJoiKiBR2NRxeCRg13uCXbydPnmB4EOeRrY7480qNWO8IIgQG6zlDkm6uRMsURXPuKq0GWtiM59a5Q6g==",
      "dependencies": {
        "@babel/runtime": "^7.5.5",
        "dom-helpers": "^5.0.1",
        "loose-envify": "^1.4.0",
        "prop-types": "^15.6.2"
      },
      "peerDependencies": {
        "react": ">=16.6.0",
        "react-dom": ">=16.6.0"
      }
    },
    "node_modules/read-cache": {
      "version": "1.0.0",
      "resolved": "https://registry.npmjs.org/read-cache/-/read-cache-1.0.0.tgz",
      "integrity": "sha512-Owdv/Ft7IjOgm/i0xvNDZ1LrRANRfew4b2prF3OWMQLxLfu3bS8FVhCsrSCMK4lR56Y9ya+AThoTpDCTxCmpRA==",
      "dependencies": {
        "pify": "^2.3.0"
      }
    },
    "node_modules/readdirp": {
      "version": "3.6.0",
      "resolved": "https://registry.npmjs.org/readdirp/-/readdirp-3.6.0.tgz",
      "integrity": "sha512-hOS089on8RduqdbhvQ5Z37A0ESjsqz6qnRcffsMU3495FuTdqSm+7bhJ29JvIOsBDEEnan5DPu9t3To9VRlMzA==",
      "dependencies": {
        "picomatch": "^2.2.1"
      },
      "engines": {
        "node": ">=8.10.0"
      }
    },
    "node_modules/recharts": {
      "version": "2.15.0",
      "resolved": "https://registry.npmjs.org/recharts/-/recharts-2.15.0.tgz",
      "integrity": "sha512-cIvMxDfpAmqAmVgc4yb7pgm/O1tmmkl/CjrvXuW+62/+7jj/iF9Ykm+hb/UJt42TREHMyd3gb+pkgoa2MxgDIw==",
      "dependencies": {
        "clsx": "^2.0.0",
        "eventemitter3": "^4.0.1",
        "lodash": "^4.17.21",
        "react-is": "^18.3.1",
        "react-smooth": "^4.0.0",
        "recharts-scale": "^0.4.4",
        "tiny-invariant": "^1.3.1",
        "victory-vendor": "^36.6.8"
      },
      "engines": {
        "node": ">=14"
      },
      "peerDependencies": {
        "react": "^16.0.0 || ^17.0.0 || ^18.0.0 || ^19.0.0",
        "react-dom": "^16.0.0 || ^17.0.0 || ^18.0.0 || ^19.0.0"
      }
    },
    "node_modules/recharts-scale": {
      "version": "0.4.5",
      "resolved": "https://registry.npmjs.org/recharts-scale/-/recharts-scale-0.4.5.tgz",
      "integrity": "sha512-kivNFO+0OcUNu7jQquLXAxz1FIwZj8nrj+YkOKc5694NbjCvcT6aSZiIzNzd2Kul4o4rTto8QVR9lMNtxD4G1w==",
      "dependencies": {
        "decimal.js-light": "^2.4.1"
      }
    },
    "node_modules/regenerator-runtime": {
      "version": "0.14.1",
      "resolved": "https://registry.npmjs.org/regenerator-runtime/-/regenerator-runtime-0.14.1.tgz",
      "integrity": "sha512-dYnhHh0nJoMfnkZs6GmmhFknAGRrLznOu5nc9ML+EJxGvrx6H7teuevqVqCuPcPK//3eDrrjQhehXVx9cnkGdw=="
    },
    "node_modules/resolve": {
      "version": "1.22.10",
      "resolved": "https://registry.npmjs.org/resolve/-/resolve-1.22.10.tgz",
      "integrity": "sha512-NPRy+/ncIMeDlTAsuqwKIiferiawhefFJtkNSW0qZJEqMEb+qBt/77B/jGeeek+F0uOeN05CDa6HXbbIgtVX4w==",
      "dependencies": {
        "is-core-module": "^2.16.0",
        "path-parse": "^1.0.7",
        "supports-preserve-symlinks-flag": "^1.0.0"
      },
      "bin": {
        "resolve": "bin/resolve"
      },
      "engines": {
        "node": ">= 0.4"
      },
      "funding": {
        "url": "https://github.com/sponsors/ljharb"
      }
    },
    "node_modules/reusify": {
      "version": "1.0.4",
      "resolved": "https://registry.npmjs.org/reusify/-/reusify-1.0.4.tgz",
      "integrity": "sha512-U9nH88a3fc/ekCF1l0/UP1IosiuIjyTh7hBvXVMHYgVcfGvt897Xguj2UOLDeI5BG2m7/uwyaLVT6fbtCwTyzw==",
      "engines": {
        "iojs": ">=1.0.0",
        "node": ">=0.10.0"
      }
    },
    "node_modules/run-parallel": {
      "version": "1.2.0",
      "resolved": "https://registry.npmjs.org/run-parallel/-/run-parallel-1.2.0.tgz",
      "integrity": "sha512-5l4VyZR86LZ/lDxZTR6jqL8AFE2S0IFLMP26AbjsLVADxHdhB/c0GUsH+y39UfCi3dzz8OlQuPmnaJOMoDHQBA==",
      "funding": [
        {
          "type": "github",
          "url": "https://github.com/sponsors/feross"
        },
        {
          "type": "patreon",
          "url": "https://www.patreon.com/feross"
        },
        {
          "type": "consulting",
          "url": "https://feross.org/support"
        }
      ],
      "dependencies": {
        "queue-microtask": "^1.2.2"
      }
    },
    "node_modules/scheduler": {
      "version": "0.23.2",
      "resolved": "https://registry.npmjs.org/scheduler/-/scheduler-0.23.2.tgz",
      "integrity": "sha512-UOShsPwz7NrMUqhR6t0hWjFduvOzbtv7toDH1/hIrfRNIDBnnBWd0CwJTGvTpngVlmwGCdP9/Zl/tVrDqcuYzQ==",
      "dependencies": {
        "loose-envify": "^1.1.0"
      }
    },
    "node_modules/shebang-command": {
      "version": "2.0.0",
      "resolved": "https://registry.npmjs.org/shebang-command/-/shebang-command-2.0.0.tgz",
      "integrity": "sha512-kHxr2zZpYtdmrN1qDjrrX/Z1rR1kG8Dx+gkpK1G4eXmvXswmcE1hTWBWYUzlraYw1/yZp6YuDY77YtvbN0dmDA==",
      "dependencies": {
        "shebang-regex": "^3.0.0"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/shebang-regex": {
      "version": "3.0.0",
      "resolved": "https://registry.npmjs.org/shebang-regex/-/shebang-regex-3.0.0.tgz",
      "integrity": "sha512-7++dFhtcx3353uBaq8DDR4NuxBetBzC7ZQOhmTQInHEd6bSrXdiEyzCvG07Z44UYdLShWUyXt5M/yhz8ekcb1A==",
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/signal-exit": {
      "version": "4.1.0",
      "resolved": "https://registry.npmjs.org/signal-exit/-/signal-exit-4.1.0.tgz",
      "integrity": "sha512-bzyZ1e88w9O1iNJbKnOlvYTrWPDl46O1bG0D3XInv+9tkPrxrN8jUUTiFlDkkmKWgn1M6CfIA13SuGqOa9Korw==",
      "engines": {
        "node": ">=14"
      },
      "funding": {
        "url": "https://github.com/sponsors/isaacs"
      }
    },
    "node_modules/sonner": {
      "version": "1.7.2",
      "resolved": "https://registry.npmjs.org/sonner/-/sonner-1.7.2.tgz",
      "integrity": "sha512-zMbseqjrOzQD1a93lxahm+qMGxWovdMxBlkTbbnZdNqVLt4j+amF9PQxUCL32WfztOFt9t9ADYkejAL3jF9iNA==",
      "peerDependencies": {
        "react": "^18.0.0 || ^19.0.0 || ^19.0.0-rc",
        "react-dom": "^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      }
    },
    "node_modules/source-map-js": {
      "version": "1.2.1",
      "resolved": "https://registry.npmjs.org/source-map-js/-/source-map-js-1.2.1.tgz",
      "integrity": "sha512-UXWMKhLOwVKb728IUtQPXxfYU+usdybtUrK/8uGE8CQMvrhOpwvzDBwj0QhSL7MQc7vIsISBG8VQ8+IDQxpfQA==",
      "engines": {
        "node": ">=0.10.0"
      }
    },
    "node_modules/sqids": {
      "version": "0.3.0",
      "resolved": "https://registry.npmjs.org/sqids/-/sqids-0.3.0.tgz",
      "integrity": "sha512-lOQK1ucVg+W6n3FhRwwSeUijxe93b51Bfz5PMRMihVf1iVkl82ePQG7V5vwrhzB11v0NtsR25PSZRGiSomJaJw=="
    },
    "node_modules/streamsearch": {
      "version": "1.1.0",
      "resolved": "https://registry.npmjs.org/streamsearch/-/streamsearch-1.1.0.tgz",
      "integrity": "sha512-Mcc5wHehp9aXz1ax6bZUyY5afg9u2rv5cqQI3mRrYkGC8rW2hM02jWuwjtL++LS5qinSyhj2QfLyNsuc+VsExg==",
      "engines": {
        "node": ">=10.0.0"
      }
    },
    "node_modules/string-width": {
      "version": "5.1.2",
      "resolved": "https://registry.npmjs.org/string-width/-/string-width-5.1.2.tgz",
      "integrity": "sha512-HnLOCR3vjcY8beoNLtcjZ5/nxn2afmME6lhrDrebokqMap+XbeW8n9TXpPDOqdGK5qcI3oT0GKTW6wC7EMiVqA==",
      "dependencies": {
        "eastasianwidth": "^0.2.0",
        "emoji-regex": "^9.2.2",
        "strip-ansi": "^7.0.1"
      },
      "engines": {
        "node": ">=12"
      },
      "funding": {
        "url": "https://github.com/sponsors/sindresorhus"
      }
    },
    "node_modules/string-width-cjs": {
      "name": "string-width",
      "version": "4.2.3",
      "resolved": "https://registry.npmjs.org/string-width/-/string-width-4.2.3.tgz",
      "integrity": "sha512-wKyQRQpjJ0sIp62ErSZdGsjMJWsap5oRNihHhu6G7JVO/9jIB6UyevL+tXuOqrng8j/cxKTWyWUwvSTriiZz/g==",
      "dependencies": {
        "emoji-regex": "^8.0.0",
        "is-fullwidth-code-point": "^3.0.0",
        "strip-ansi": "^6.0.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/string-width-cjs/node_modules/ansi-regex": {
      "version": "5.0.1",
      "resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.1.tgz",
      "integrity": "sha512-quJQXlTSUGL2LH9SUXo8VwsY4soanhgo6LNSm84E1LBcE8s3O0wpdiRzyR9z/ZZJMlMWv37qOOb9pdJlMUEKFQ==",
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/string-width-cjs/node_modules/emoji-regex": {
      "version": "8.0.0",
      "resolved": "https://registry.npmjs.org/emoji-regex/-/emoji-regex-8.0.0.tgz",
      "integrity": "sha512-MSjYzcWNOA0ewAHpz0MxpYFvwg6yjy1NG3xteoqz644VCo/RPgnr1/GGt+ic3iJTzQ8Eu3TdM14SawnVUmGE6A=="
    },
    "node_modules/string-width-cjs/node_modules/strip-ansi": {
      "version": "6.0.1",
      "resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.1.tgz",
      "integrity": "sha512-Y38VPSHcqkFrCpFnQ9vuSXmquuv5oXOKpGeT6aGrr3o3Gc9AlVa6JBfUSOCnbxGGZF+/0ooI7KrPuUSztUdU5A==",
      "dependencies": {
        "ansi-regex": "^5.0.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/strip-ansi": {
      "version": "7.1.0",
      "resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-7.1.0.tgz",
      "integrity": "sha512-iq6eVVI64nQQTRYq2KtEg2d2uU7LElhTJwsH4YzIHZshxlgZms/wIc4VoDQTlG/IvVIrBKG06CrZnp0qv7hkcQ==",
      "dependencies": {
        "ansi-regex": "^6.0.1"
      },
      "engines": {
        "node": ">=12"
      },
      "funding": {
        "url": "https://github.com/chalk/strip-ansi?sponsor=1"
      }
    },
    "node_modules/strip-ansi-cjs": {
      "name": "strip-ansi",
      "version": "6.0.1",
      "resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.1.tgz",
      "integrity": "sha512-Y38VPSHcqkFrCpFnQ9vuSXmquuv5oXOKpGeT6aGrr3o3Gc9AlVa6JBfUSOCnbxGGZF+/0ooI7KrPuUSztUdU5A==",
      "dependencies": {
        "ansi-regex": "^5.0.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/strip-ansi-cjs/node_modules/ansi-regex": {
      "version": "5.0.1",
      "resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.1.tgz",
      "integrity": "sha512-quJQXlTSUGL2LH9SUXo8VwsY4soanhgo6LNSm84E1LBcE8s3O0wpdiRzyR9z/ZZJMlMWv37qOOb9pdJlMUEKFQ==",
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/styled-jsx": {
      "version": "5.1.1",
      "resolved": "https://registry.npmjs.org/styled-jsx/-/styled-jsx-5.1.1.tgz",
      "integrity": "sha512-pW7uC1l4mBZ8ugbiZrcIsiIvVx1UmTfw7UkC3Um2tmfUq9Bhk8IiyEIPl6F8agHgjzku6j0xQEZbfA5uSgSaCw==",
      "dependencies": {
        "client-only": "0.0.1"
      },
      "engines": {
        "node": ">= 12.0.0"
      },
      "peerDependencies": {
        "react": ">= 16.8.0 || 17.x.x || ^18.0.0-0"
      },
      "peerDependenciesMeta": {
        "@babel/core": {
          "optional": true
        },
        "babel-plugin-macros": {
          "optional": true
        }
      }
    },
    "node_modules/sucrase": {
      "version": "3.35.0",
      "resolved": "https://registry.npmjs.org/sucrase/-/sucrase-3.35.0.tgz",
      "integrity": "sha512-8EbVDiu9iN/nESwxeSxDKe0dunta1GOlHufmSSXxMD2z2/tMZpDMpvXQGsc+ajGo8y2uYUmixaSRUc/QPoQ0GA==",
      "dependencies": {
        "@jridgewell/gen-mapping": "^0.3.2",
        "commander": "^4.0.0",
        "glob": "^10.3.10",
        "lines-and-columns": "^1.1.6",
        "mz": "^2.7.0",
        "pirates": "^4.0.1",
        "ts-interface-checker": "^0.1.9"
      },
      "bin": {
        "sucrase": "bin/sucrase",
        "sucrase-node": "bin/sucrase-node"
      },
      "engines": {
        "node": ">=16 || 14 >=14.17"
      }
    },
    "node_modules/supports-preserve-symlinks-flag": {
      "version": "1.0.0",
      "resolved": "https://registry.npmjs.org/supports-preserve-symlinks-flag/-/supports-preserve-symlinks-flag-1.0.0.tgz",
      "integrity": "sha512-ot0WnXS9fgdkgIcePe6RHNk1WA8+muPa6cSjeR3V8K27q9BB1rTE3R1p7Hv0z1ZyAc8s6Vvv8DIyWf681MAt0w==",
      "engines": {
        "node": ">= 0.4"
      },
      "funding": {
        "url": "https://github.com/sponsors/ljharb"
      }
    },
    "node_modules/tailwind-merge": {
      "version": "2.6.0",
      "resolved": "https://registry.npmjs.org/tailwind-merge/-/tailwind-merge-2.6.0.tgz",
      "integrity": "sha512-P+Vu1qXfzediirmHOC3xKGAYeZtPcV9g76X+xg2FD4tYgR71ewMA35Y3sCz3zhiN/dwefRpJX0yBcgwi1fXNQA==",
      "funding": {
        "type": "github",
        "url": "https://github.com/sponsors/dcastil"
      }
    },
    "node_modules/tailwindcss": {
      "version": "3.4.17",
      "resolved": "https://registry.npmjs.org/tailwindcss/-/tailwindcss-3.4.17.tgz",
      "integrity": "sha512-w33E2aCvSDP0tW9RZuNXadXlkHXqFzSkQew/aIa2i/Sj8fThxwovwlXHSPXTbAHwEIhBFXAedUhP2tueAKP8Og==",
      "dependencies": {
        "@alloc/quick-lru": "^5.2.0",
        "arg": "^5.0.2",
        "chokidar": "^3.6.0",
        "didyoumean": "^1.2.2",
        "dlv": "^1.1.3",
        "fast-glob": "^3.3.2",
        "glob-parent": "^6.0.2",
        "is-glob": "^4.0.3",
        "jiti": "^1.21.6",
        "lilconfig": "^3.1.3",
        "micromatch": "^4.0.8",
        "normalize-path": "^3.0.0",
        "object-hash": "^3.0.0",
        "picocolors": "^1.1.1",
        "postcss": "^8.4.47",
        "postcss-import": "^15.1.0",
        "postcss-js": "^4.0.1",
        "postcss-load-config": "^4.0.2",
        "postcss-nested": "^6.2.0",
        "postcss-selector-parser": "^6.1.2",
        "resolve": "^1.22.8",
        "sucrase": "^3.35.0"
      },
      "bin": {
        "tailwind": "lib/cli.js",
        "tailwindcss": "lib/cli.js"
      },
      "engines": {
        "node": ">=14.0.0"
      }
    },
    "node_modules/tailwindcss-animate": {
      "version": "1.0.7",
      "resolved": "https://registry.npmjs.org/tailwindcss-animate/-/tailwindcss-animate-1.0.7.tgz",
      "integrity": "sha512-bl6mpH3T7I3UFxuvDEXLxy/VuFxBk5bbzplh7tXI68mwMokNYd1t9qPBHlnyTwfa4JGC4zP516I1hYYtQ/vspA==",
      "peerDependencies": {
        "tailwindcss": ">=3.0.0 || insiders"
      }
    },
    "node_modules/thenify": {
      "version": "3.3.1",
      "resolved": "https://registry.npmjs.org/thenify/-/thenify-3.3.1.tgz",
      "integrity": "sha512-RVZSIV5IG10Hk3enotrhvz0T9em6cyHBLkH/YAZuKqd8hRkKhSfCGIcP2KUY0EPxndzANBmNllzWPwak+bheSw==",
      "dependencies": {
        "any-promise": "^1.0.0"
      }
    },
    "node_modules/thenify-all": {
      "version": "1.6.0",
      "resolved": "https://registry.npmjs.org/thenify-all/-/thenify-all-1.6.0.tgz",
      "integrity": "sha512-RNxQH/qI8/t3thXJDwcstUO4zeqo64+Uy/+sNVRBx4Xn2OX+OZ9oP+iJnNFqplFra2ZUVeKCSa2oVWi3T4uVmA==",
      "dependencies": {
        "thenify": ">= 3.1.0 < 4"
      },
      "engines": {
        "node": ">=0.8"
      }
    },
    "node_modules/tiny-invariant": {
      "version": "1.3.3",
      "resolved": "https://registry.npmjs.org/tiny-invariant/-/tiny-invariant-1.3.3.tgz",
      "integrity": "sha512-+FbBPE1o9QAYvviau/qC5SE3caw21q3xkvWKBtja5vgqOWIHHJ3ioaq1VPfn/Szqctz2bU/oYeKd9/z5BL+PVg=="
    },
    "node_modules/to-regex-range": {
      "version": "5.0.1",
      "resolved": "https://registry.npmjs.org/to-regex-range/-/to-regex-range-5.0.1.tgz",
      "integrity": "sha512-65P7iz6X5yEr1cwcgvQxbbIw7Uk3gOy5dIdtZ4rDveLqhrdJP+Li/Hx6tyK0NEb+2GCyneCMJiGqrADCSNk8sQ==",
      "dependencies": {
        "is-number": "^7.0.0"
      },
      "engines": {
        "node": ">=8.0"
      }
    },
    "node_modules/ts-interface-checker": {
      "version": "0.1.13",
      "resolved": "https://registry.npmjs.org/ts-interface-checker/-/ts-interface-checker-0.1.13.tgz",
      "integrity": "sha512-Y/arvbn+rrz3JCKl9C4kVNfTfSm2/mEp5FSz5EsZSANGPSlQrpRI5M4PKF+mJnE52jOO90PnPSc3Ur3bTQw0gA=="
    },
    "node_modules/tslib": {
      "version": "2.8.1",
      "resolved": "https://registry.npmjs.org/tslib/-/tslib-2.8.1.tgz",
      "integrity": "sha512-oJFu94HQb+KVduSUQL7wnpmqnfmLsOA/nAh6b6EH0wCEoK0/mPeXU6c3wKDV83MkOuHPRHtSXKKU99IBazS/2w=="
    },
    "node_modules/typescript": {
      "version": "5.7.3",
      "resolved": "https://registry.npmjs.org/typescript/-/typescript-5.7.3.tgz",
      "integrity": "sha512-84MVSjMEHP+FQRPy3pX9sTVV/INIex71s9TL2Gm5FG/WG1SqXeKyZ0k7/blY/4FdOzI12CBy1vGc4og/eus0fw==",
      "dev": true,
      "bin": {
        "tsc": "bin/tsc",
        "tsserver": "bin/tsserver"
      },
      "engines": {
        "node": ">=14.17"
      }
    },
    "node_modules/undici-types": {
      "version": "6.20.0",
      "resolved": "https://registry.npmjs.org/undici-types/-/undici-types-6.20.0.tgz",
      "integrity": "sha512-Ny6QZ2Nju20vw1SRHe3d9jVu6gJ+4e3+MMpqu7pqE5HT6WsTSlce++GQmK5UXS8mzV8DSYHrQH+Xrf2jVcuKNg==",
      "dev": true
    },
    "node_modules/update-browserslist-db": {
      "version": "1.1.2",
      "resolved": "https://registry.npmjs.org/update-browserslist-db/-/update-browserslist-db-1.1.2.tgz",
      "integrity": "sha512-PPypAm5qvlD7XMZC3BujecnaOxwhrtoFR+Dqkk5Aa/6DssiH0ibKoketaj9w8LP7Bont1rYeoV5plxD7RTEPRg==",
      "funding": [
        {
          "type": "opencollective",
          "url": "https://opencollective.com/browserslist"
        },
        {
          "type": "tidelift",
          "url": "https://tidelift.com/funding/github/npm/browserslist"
        },
        {
          "type": "github",
          "url": "https://github.com/sponsors/ai"
        }
      ],
      "dependencies": {
        "escalade": "^3.2.0",
        "picocolors": "^1.1.1"
      },
      "bin": {
        "update-browserslist-db": "cli.js"
      },
      "peerDependencies": {
        "browserslist": ">= 4.21.0"
      }
    },
    "node_modules/uploadthing": {
      "version": "7.4.4",
      "resolved": "https://registry.npmjs.org/uploadthing/-/uploadthing-7.4.4.tgz",
      "integrity": "sha512-dgV8LqaKrXSnpDkqzzLbVEmIubMPi8a5bRtIzM+rG5EbZpIZyOG0bByLxdSlkSbK5L6F+HbIGX5eDkLPFNL8dg==",
      "dependencies": {
        "@effect/platform": "0.70.7",
        "@standard-schema/spec": "1.0.0-beta.3",
        "@uploadthing/mime-types": "0.3.4",
        "@uploadthing/shared": "7.1.5",
        "effect": "3.11.5"
      },
      "engines": {
        "node": ">=18.13.0"
      },
      "peerDependencies": {
        "express": "*",
        "h3": "*",
        "tailwindcss": "^3.0.0 || ^4.0.0-beta.0"
      },
      "peerDependenciesMeta": {
        "express": {
          "optional": true
        },
        "fastify": {
          "optional": true
        },
        "h3": {
          "optional": true
        },
        "next": {
          "optional": true
        },
        "tailwindcss": {
          "optional": true
        }
      }
    },
    "node_modules/use-callback-ref": {
      "version": "1.3.3",
      "resolved": "https://registry.npmjs.org/use-callback-ref/-/use-callback-ref-1.3.3.tgz",
      "integrity": "sha512-jQL3lRnocaFtu3V00JToYz/4QkNWswxijDaCVNZRiRTO3HQDLsdu1ZtmIUvV4yPp+rvWm5j0y0TG/S61cuijTg==",
      "dependencies": {
        "tslib": "^2.0.0"
      },
      "engines": {
        "node": ">=10"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/use-sidecar": {
      "version": "1.1.3",
      "resolved": "https://registry.npmjs.org/use-sidecar/-/use-sidecar-1.1.3.tgz",
      "integrity": "sha512-Fedw0aZvkhynoPYlA5WXrMCAMm+nSWdZt6lzJQ7Ok8S6Q+VsHmHpRWndVRJ8Be0ZbkfPc5LRYH+5XrzXcEeLRQ==",
      "dependencies": {
        "detect-node-es": "^1.1.0",
        "tslib": "^2.0.0"
      },
      "engines": {
        "node": ">=10"
      },
      "peerDependencies": {
        "@types/react": "*",
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0 || ^19.0.0-rc"
      },
      "peerDependenciesMeta": {
        "@types/react": {
          "optional": true
        }
      }
    },
    "node_modules/use-sync-external-store": {
      "version": "1.4.0",
      "resolved": "https://registry.npmjs.org/use-sync-external-store/-/use-sync-external-store-1.4.0.tgz",
      "integrity": "sha512-9WXSPC5fMv61vaupRkCKCxsPxBocVnwakBEkMIHHpkTTg6icbJtg6jzgtLDm4bl3cSHAca52rYWih0k4K3PfHw==",
      "peerDependencies": {
        "react": "^16.8.0 || ^17.0.0 || ^18.0.0 || ^19.0.0"
      }
    },
    "node_modules/util-deprecate": {
      "version": "1.0.2",
      "resolved": "https://registry.npmjs.org/util-deprecate/-/util-deprecate-1.0.2.tgz",
      "integrity": "sha512-EPD5q1uXyFxJpCrLnCc1nHnq3gOa6DZBocAIiI2TaSCA7VCJ1UJDMagCzIkXNsUYfD1daK//LTEQ8xiIbrHtcw=="
    },
    "node_modules/vaul": {
      "version": "0.9.9",
      "resolved": "https://registry.npmjs.org/vaul/-/vaul-0.9.9.tgz",
      "integrity": "sha512-7afKg48srluhZwIkaU+lgGtFCUsYBSGOl8vcc8N/M3YQlZFlynHD15AE+pwrYdc826o7nrIND4lL9Y6b9WWZZQ==",
      "dependencies": {
        "@radix-ui/react-dialog": "^1.1.1"
      },
      "peerDependencies": {
        "react": "^16.8 || ^17.0 || ^18.0",
        "react-dom": "^16.8 || ^17.0 || ^18.0"
      }
    },
    "node_modules/victory-vendor": {
      "version": "36.9.2",
      "resolved": "https://registry.npmjs.org/victory-vendor/-/victory-vendor-36.9.2.tgz",
      "integrity": "sha512-PnpQQMuxlwYdocC8fIJqVXvkeViHYzotI+NJrCuav0ZYFoq912ZHBk3mCeuj+5/VpodOjPe1z0Fk2ihgzlXqjQ==",
      "dependencies": {
        "@types/d3-array": "^3.0.3",
        "@types/d3-ease": "^3.0.0",
        "@types/d3-interpolate": "^3.0.1",
        "@types/d3-scale": "^4.0.2",
        "@types/d3-shape": "^3.1.0",
        "@types/d3-time": "^3.0.0",
        "@types/d3-timer": "^3.0.0",
        "d3-array": "^3.1.6",
        "d3-ease": "^3.0.1",
        "d3-interpolate": "^3.0.1",
        "d3-scale": "^4.0.2",
        "d3-shape": "^3.1.0",
        "d3-time": "^3.0.0",
        "d3-timer": "^3.0.1"
      }
    },
    "node_modules/which": {
      "version": "2.0.2",
      "resolved": "https://registry.npmjs.org/which/-/which-2.0.2.tgz",
      "integrity": "sha512-BLI3Tl1TW3Pvl70l3yq3Y64i+awpwXqsGBYWkkqMtnbXgrMD+yj7rhW0kuEDxzJaYXGjEW5ogapKNMEKNMjibA==",
      "dependencies": {
        "isexe": "^2.0.0"
      },
      "bin": {
        "node-which": "bin/node-which"
      },
      "engines": {
        "node": ">= 8"
      }
    },
    "node_modules/wrap-ansi": {
      "version": "8.1.0",
      "resolved": "https://registry.npmjs.org/wrap-ansi/-/wrap-ansi-8.1.0.tgz",
      "integrity": "sha512-si7QWI6zUMq56bESFvagtmzMdGOtoxfR+Sez11Mobfc7tm+VkUckk9bW2UeffTGVUbOksxmSw0AA2gs8g71NCQ==",
      "dependencies": {
        "ansi-styles": "^6.1.0",
        "string-width": "^5.0.1",
        "strip-ansi": "^7.0.1"
      },
      "engines": {
        "node": ">=12"
      },
      "funding": {
        "url": "https://github.com/chalk/wrap-ansi?sponsor=1"
      }
    },
    "node_modules/wrap-ansi-cjs": {
      "name": "wrap-ansi",
      "version": "7.0.0",
      "resolved": "https://registry.npmjs.org/wrap-ansi/-/wrap-ansi-7.0.0.tgz",
      "integrity": "sha512-YVGIj2kamLSTxw6NsZjoBxfSwsn0ycdesmc4p+Q21c5zPuZ1pl+NfxVdxPtdHvmNVOQ6XSYG4AUtyt/Fi7D16Q==",
      "dependencies": {
        "ansi-styles": "^4.0.0",
        "string-width": "^4.1.0",
        "strip-ansi": "^6.0.0"
      },
      "engines": {
        "node": ">=10"
      },
      "funding": {
        "url": "https://github.com/chalk/wrap-ansi?sponsor=1"
      }
    },
    "node_modules/wrap-ansi-cjs/node_modules/ansi-regex": {
      "version": "5.0.1",
      "resolved": "https://registry.npmjs.org/ansi-regex/-/ansi-regex-5.0.1.tgz",
      "integrity": "sha512-quJQXlTSUGL2LH9SUXo8VwsY4soanhgo6LNSm84E1LBcE8s3O0wpdiRzyR9z/ZZJMlMWv37qOOb9pdJlMUEKFQ==",
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/wrap-ansi-cjs/node_modules/ansi-styles": {
      "version": "4.3.0",
      "resolved": "https://registry.npmjs.org/ansi-styles/-/ansi-styles-4.3.0.tgz",
      "integrity": "sha512-zbB9rCJAT1rbjiVDb2hqKFHNYLxgtk8NURxZ3IZwD3F6NtxbXZQCnnSi1Lkx+IDohdPlFp222wVALIheZJQSEg==",
      "dependencies": {
        "color-convert": "^2.0.1"
      },
      "engines": {
        "node": ">=8"
      },
      "funding": {
        "url": "https://github.com/chalk/ansi-styles?sponsor=1"
      }
    },
    "node_modules/wrap-ansi-cjs/node_modules/emoji-regex": {
      "version": "8.0.0",
      "resolved": "https://registry.npmjs.org/emoji-regex/-/emoji-regex-8.0.0.tgz",
      "integrity": "sha512-MSjYzcWNOA0ewAHpz0MxpYFvwg6yjy1NG3xteoqz644VCo/RPgnr1/GGt+ic3iJTzQ8Eu3TdM14SawnVUmGE6A=="
    },
    "node_modules/wrap-ansi-cjs/node_modules/string-width": {
      "version": "4.2.3",
      "resolved": "https://registry.npmjs.org/string-width/-/string-width-4.2.3.tgz",
      "integrity": "sha512-wKyQRQpjJ0sIp62ErSZdGsjMJWsap5oRNihHhu6G7JVO/9jIB6UyevL+tXuOqrng8j/cxKTWyWUwvSTriiZz/g==",
      "dependencies": {
        "emoji-regex": "^8.0.0",
        "is-fullwidth-code-point": "^3.0.0",
        "strip-ansi": "^6.0.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/wrap-ansi-cjs/node_modules/strip-ansi": {
      "version": "6.0.1",
      "resolved": "https://registry.npmjs.org/strip-ansi/-/strip-ansi-6.0.1.tgz",
      "integrity": "sha512-Y38VPSHcqkFrCpFnQ9vuSXmquuv5oXOKpGeT6aGrr3o3Gc9AlVa6JBfUSOCnbxGGZF+/0ooI7KrPuUSztUdU5A==",
      "dependencies": {
        "ansi-regex": "^5.0.1"
      },
      "engines": {
        "node": ">=8"
      }
    },
    "node_modules/yaml": {
      "version": "2.7.0",
      "resolved": "https://registry.npmjs.org/yaml/-/yaml-2.7.0.tgz",
      "integrity": "sha512-+hSoy/QHluxmC9kCIJyL/uyFmLmc+e5CFR5Wa+bpIhIj85LVb9ZH2nVnqrHoSvKogwODv0ClqZkmiSSaIH5LTA==",
      "bin": {
        "yaml": "bin.mjs"
      },
      "engines": {
        "node": ">= 14"
      }
    },
    "node_modules/zod": {
      "version": "3.24.1",
      "resolved": "https://registry.npmjs.org/zod/-/zod-3.24.1.tgz",
      "integrity": "sha512-muH7gBL9sI1nciMZV67X5fTKKBLtwpZ5VBp1vsOQzj1MhrBZ4wlVCm3gedKZWLp0Oyel8sIGfeiz54Su+OVT+A==",
      "funding": {
        "url": "https://github.com/sponsors/colinhacks"
      }
    }
  }
}
```

## File: /Users/saint/Desktop/pdf-player/package.json

```json
{
  "name": "my-v0-project",
  "version": "0.1.0",
  "private": true,
  "scripts": {
    "dev": "next dev",
    "build": "next build",
    "start": "next start",
    "lint": "next lint"
  },
  "dependencies": {
    "@hookform/resolvers": "^3.9.1",
    "@radix-ui/react-accordion": "^1.2.2",
    "@radix-ui/react-alert-dialog": "^1.1.4",
    "@radix-ui/react-aspect-ratio": "^1.1.1",
    "@radix-ui/react-avatar": "^1.1.2",
    "@radix-ui/react-checkbox": "^1.1.3",
    "@radix-ui/react-collapsible": "^1.1.2",
    "@radix-ui/react-context-menu": "^2.2.4",
    "@radix-ui/react-dialog": "^1.1.4",
    "@radix-ui/react-dropdown-menu": "^2.1.4",
    "@radix-ui/react-hover-card": "^1.1.4",
    "@radix-ui/react-label": "^2.1.1",
    "@radix-ui/react-menubar": "^1.1.4",
    "@radix-ui/react-navigation-menu": "^1.2.3",
    "@radix-ui/react-popover": "^1.1.4",
    "@radix-ui/react-progress": "^1.1.1",
    "@radix-ui/react-radio-group": "^1.2.2",
    "@radix-ui/react-scroll-area": "^1.2.2",
    "@radix-ui/react-select": "^2.1.4",
    "@radix-ui/react-separator": "^1.1.1",
    "@radix-ui/react-slider": "^1.2.2",
    "@radix-ui/react-slot": "^1.1.1",
    "@radix-ui/react-switch": "^1.1.2",
    "@radix-ui/react-tabs": "^1.1.2",
    "@radix-ui/react-toast": "^1.2.4",
    "@radix-ui/react-toggle": "^1.1.1",
    "@radix-ui/react-toggle-group": "^1.1.1",
    "@radix-ui/react-tooltip": "^1.1.6",
    "@uploadthing/react": "^7.1.5",
    "autoprefixer": "^10.4.20",
    "class-variance-authority": "^0.7.1",
    "clsx": "^2.1.1",
    "cmdk": "1.0.4",
    "date-fns": "4.1.0",
    "embla-carousel-react": "8.5.1",
    "framer-motion": "latest",
    "input-otp": "1.4.1",
    "lucide-react": "^0.454.0",
    "next": "14.2.16",
    "next-themes": "latest",
    "react": "^18",
    "react-day-picker": "8.10.1",
    "react-dom": "^18",
    "react-hook-form": "^7.54.1",
    "react-resizable-panels": "^2.1.7",
    "recharts": "2.15.0",
    "sonner": "^1.7.1",
    "tailwind-merge": "^2.5.5",
    "tailwindcss-animate": "^1.0.7",
    "uploadthing": "^7.4.4",
    "vaul": "^0.9.6",
    "zod": "^3.24.1"
  },
  "devDependencies": {
    "@types/node": "^22",
    "@types/react": "^18",
    "@types/react-dom": "^18",
    "postcss": "^8",
    "tailwindcss": "^3.4.17",
    "typescript": "^5"
  }
}
```

## File: /Users/saint/Desktop/pdf-player/backend.txt

```txt
eReader Backend Setup

1. Set up the project:
 - Create a new directory for the backend
 - Initialize a new Go module: go mod init github.com/yourusername/ereader-backend

2. Install dependencies:
 - go get github.com/gorilla/mux
 - go get github.com/ledongthuc/pdf
 - go get github.com/replicate/replicate-go
 - go get github.com/uploadthing/go
 - go get github.com/mattn/go-sqlite3

3. Set up Replicate credentials:
 - Sign up for a Replicate account at https://replicate.com
 - Obtain your API token from the Replicate dashboard
 - Set the REPLICATE_API_TOKEN environment variable with your token

4. Set up Uploadthing credentials:
 - Sign up for an Uploadthing account at https://uploadthing.com
 - Obtain your API keys from the Uploadthing dashboard
 - Set the UPLOADTHING_APP_ID and UPLOADTHING_SECRET environment variables with your keys

5. Set up SQLite database:
 - Create a new SQLite database file: touch ereader.db
 - Set up the database schema with necessary tables (books, users, etc.)
 - Implement database connection and query functions in your Go code

6. Implement the main.go file with the following features:
 - PDF upload endpoint (/api/upload)
 - Text extraction from PDF
 - Text-to-speech synthesis endpoint (/api/synthesize) using Kokoro TTS via Replicate
 - Image upload endpoint (/api/upload-image) using Uploadthing
 - Database operations for storing and retrieving book information

7. Create a PDF processing function:
 - Use github.com/ledongthuc/pdf to extract text from uploaded PDFs
 - Store the extracted text in the SQLite database

8. Implement the Kokoro TTS integration:
 - Use the Replicate Go client to interact with the Kokoro TTS model
 - Convert extracted text to speech using the Kokoro TTS model
 - Return the generated audio file URL
 - Store the audio file URL in the SQLite database

9. Set up the Uploadthing integration:
 - Implement the image upload functionality using the Uploadthing Go SDK
 - Store and retrieve image URLs for book covers in the SQLite database

10. Create API endpoints:
 - POST /api/upload: Handle PDF uploads, text extraction, and store in database
 - POST /api/synthesize: Generate speech from extracted text and update database
 - POST /api/upload-image: Handle image uploads for book covers and update database
 - GET /api/books: Retrieve a list of uploaded books with their details from database
 - GET /api/book/{id}: Retrieve a specific book's details from database

11. Implement database operations:
 - Create functions for inserting, updating, and querying book information
 - Implement user management (if required) using the SQLite database
 - Ensure proper error handling for database operations

12. Implement error handling and logging throughout the application

13. Run the backend:
 - go run main.go

Note: Ensure that the backend is running on the correct port (default: 8080) and that the frontend is configured to make API calls to the correct backend URL.

Remember to handle authentication and authorization for your API endpoints to secure your application. Also, make sure to properly manage database connections and implement connection pooling if necessary for better performance.

```

## File: /Users/saint/Desktop/pdf-player/hooks/use-mobile.tsx

```tsx
import * as React from "react"

const MOBILE_BREAKPOINT = 768

export function useIsMobile() {
  const [isMobile, setIsMobile] = React.useState<boolean | undefined>(undefined)

  React.useEffect(() => {
    const mql = window.matchMedia(`(max-width: ${MOBILE_BREAKPOINT - 1}px)`)
    const onChange = () => {
      setIsMobile(window.innerWidth < MOBILE_BREAKPOINT)
    }
    mql.addEventListener("change", onChange)
    setIsMobile(window.innerWidth < MOBILE_BREAKPOINT)
    return () => mql.removeEventListener("change", onChange)
  }, [])

  return !!isMobile
}
```

## File: /Users/saint/Desktop/pdf-player/hooks/use-toast.ts

```ts
"use client"

// Inspired by react-hot-toast library
import * as React from "react"

import type {
  ToastActionElement,
  ToastProps,
} from "@/components/ui/toast"

const TOAST_LIMIT = 1
const TOAST_REMOVE_DELAY = 1000000

type ToasterToast = ToastProps & {
  id: string
  title?: React.ReactNode
  description?: React.ReactNode
  action?: ToastActionElement
}

const actionTypes = {
  ADD_TOAST: "ADD_TOAST",
  UPDATE_TOAST: "UPDATE_TOAST",
  DISMISS_TOAST: "DISMISS_TOAST",
  REMOVE_TOAST: "REMOVE_TOAST",
} as const

let count = 0

function genId() {
  count = (count + 1) % Number.MAX_SAFE_INTEGER
  return count.toString()
}

type ActionType = typeof actionTypes

type Action =
  | {
      type: ActionType["ADD_TOAST"]
      toast: ToasterToast
    }
  | {
      type: ActionType["UPDATE_TOAST"]
      toast: Partial<ToasterToast>
    }
  | {
      type: ActionType["DISMISS_TOAST"]
      toastId?: ToasterToast["id"]
    }
  | {
      type: ActionType["REMOVE_TOAST"]
      toastId?: ToasterToast["id"]
    }

interface State {
  toasts: ToasterToast[]
}

const toastTimeouts = new Map<string, ReturnType<typeof setTimeout>>()

const addToRemoveQueue = (toastId: string) => {
  if (toastTimeouts.has(toastId)) {
    return
  }

  const timeout = setTimeout(() => {
    toastTimeouts.delete(toastId)
    dispatch({
      type: "REMOVE_TOAST",
      toastId: toastId,
    })
  }, TOAST_REMOVE_DELAY)

  toastTimeouts.set(toastId, timeout)
}

export const reducer = (state: State, action: Action): State => {
  switch (action.type) {
    case "ADD_TOAST":
      return {
        ...state,
        toasts: [action.toast, ...state.toasts].slice(0, TOAST_LIMIT),
      }

    case "UPDATE_TOAST":
      return {
        ...state,
        toasts: state.toasts.map((t) =>
          t.id === action.toast.id ? { ...t, ...action.toast } : t
        ),
      }

    case "DISMISS_TOAST": {
      const { toastId } = action

      // ! Side effects ! - This could be extracted into a dismissToast() action,
      // but I'll keep it here for simplicity
      if (toastId) {
        addToRemoveQueue(toastId)
      } else {
        state.toasts.forEach((toast) => {
          addToRemoveQueue(toast.id)
        })
      }

      return {
        ...state,
        toasts: state.toasts.map((t) =>
          t.id === toastId || toastId === undefined
            ? {
                ...t,
                open: false,
              }
            : t
        ),
      }
    }
    case "REMOVE_TOAST":
      if (action.toastId === undefined) {
        return {
          ...state,
          toasts: [],
        }
      }
      return {
        ...state,
        toasts: state.toasts.filter((t) => t.id !== action.toastId),
      }
  }
}

const listeners: Array<(state: State) => void> = []

let memoryState: State = { toasts: [] }

function dispatch(action: Action) {
  memoryState = reducer(memoryState, action)
  listeners.forEach((listener) => {
    listener(memoryState)
  })
}

type Toast = Omit<ToasterToast, "id">

function toast({ ...props }: Toast) {
  const id = genId()

  const update = (props: ToasterToast) =>
    dispatch({
      type: "UPDATE_TOAST",
      toast: { ...props, id },
    })
  const dismiss = () => dispatch({ type: "DISMISS_TOAST", toastId: id })

  dispatch({
    type: "ADD_TOAST",
    toast: {
      ...props,
      id,
      open: true,
      onOpenChange: (open) => {
        if (!open) dismiss()
      },
    },
  })

  return {
    id: id,
    dismiss,
    update,
  }
}

function useToast() {
  const [state, setState] = React.useState<State>(memoryState)

  React.useEffect(() => {
    listeners.push(setState)
    return () => {
      const index = listeners.indexOf(setState)
      if (index > -1) {
        listeners.splice(index, 1)
      }
    }
  }, [state])

  return {
    ...state,
    toast,
    dismiss: (toastId?: string) => dispatch({ type: "DISMISS_TOAST", toastId }),
  }
}

export { useToast, toast }
```

## File: /Users/saint/Desktop/pdf-player/hooks/useBooks.ts

```ts
"use client"

import { useState, useEffect } from 'react'
import { Book } from '@/types'

export function useBooks() {
  const [books, setBooks] = useState<Book[]>([])
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    fetchBooks()
  }, [])

  const fetchBooks = async () => {
    try {
      const response = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/api/books`)
      if (!response.ok) throw new Error('Failed to fetch books')
      const data = await response.json()
      setBooks(data)
      setError(null)
    } catch (err) {
      console.error('Error fetching books:', err)
      setError('Failed to load books')
    } finally {
      setLoading(false)
    }
  }

  const addBook = async (book: Book) => {
    try {
      setBooks(prev => [...prev, book])
      // Optimistic update, but we'll refresh the list to ensure consistency
      await fetchBooks()
    } catch (err) {
      console.error('Error adding book:', err)
      setError('Failed to add book')
    }
  }

  const updateBook = async (id: string, updates: Partial<Book>) => {
    try {
      setBooks(prev => prev.map(book => 
        book.id === id ? { ...book, ...updates } : book
      ))
      // Optimistic update, but we'll refresh the list to ensure consistency
      await fetchBooks()
    } catch (err) {
      console.error('Error updating book:', err)
      setError('Failed to update book')
    }
  }

  const getBookById = (id: string) => {
    return books.find(book => book.id === id)
  }

  return {
    books,
    loading,
    error,
    addBook,
    updateBook,
    getBookById,
    refreshBooks: fetchBooks
  }
} 
```

## File: /Users/saint/Desktop/pdf-player/lib/utils.ts

```ts
import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
```

## File: /Users/saint/Desktop/pdf-player/lib/uploadthing.ts

```ts
import { generateReactHelpers } from "@uploadthing/react"
import type { OurFileRouter } from "@/app/api/uploadthing/core"

export const { useUploadThing, uploadFiles } = generateReactHelpers<OurFileRouter>() 
```

## File: /Users/saint/Desktop/pdf-player/components.json

```json
{
  "$schema": "https://ui.shadcn.com/schema.json",
  "style": "default",
  "rsc": true,
  "tsx": true,
  "tailwind": {
    "config": "tailwind.config.ts",
    "css": "app/globals.css",
    "baseColor": "neutral",
    "cssVariables": true,
    "prefix": ""
  },
  "aliases": {
    "components": "@/components",
    "utils": "@/lib/utils",
    "ui": "@/components/ui",
    "lib": "@/lib",
    "hooks": "@/hooks"
  },
  "iconLibrary": "lucide"
}
```

## File: /Users/saint/Desktop/pdf-player/tsconfig.json

```json
{
  "compilerOptions": {
    "lib": ["dom", "dom.iterable", "esnext"],
    "allowJs": true,
    "target": "ES6",
    "skipLibCheck": true,
    "strict": true,
    "noEmit": true,
    "esModuleInterop": true,
    "module": "esnext",
    "moduleResolution": "bundler",
    "resolveJsonModule": true,
    "isolatedModules": true,
    "jsx": "preserve",
    "incremental": true,
    "plugins": [
      {
        "name": "next"
      }
    ],
    "paths": {
      "@/*": ["./*"]
    }
  },
  "include": ["next-env.d.ts", "**/*.ts", "**/*.tsx", ".next/types/**/*.ts"],
  "exclude": ["node_modules"]
}
```

## File: /Users/saint/Desktop/pdf-player/test.txt

```txt
This is a test file
```



---

> 📸 Generated with [Jockey CLI](https://github.com/saint0x/jockey-cli)
