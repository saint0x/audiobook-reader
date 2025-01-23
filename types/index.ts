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