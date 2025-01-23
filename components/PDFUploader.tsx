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

