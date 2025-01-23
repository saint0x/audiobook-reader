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

