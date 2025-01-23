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