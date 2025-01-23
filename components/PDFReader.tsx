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

