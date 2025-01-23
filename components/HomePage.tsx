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