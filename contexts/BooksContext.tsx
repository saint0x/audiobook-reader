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