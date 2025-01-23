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