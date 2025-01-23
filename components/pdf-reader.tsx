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

