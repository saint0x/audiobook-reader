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