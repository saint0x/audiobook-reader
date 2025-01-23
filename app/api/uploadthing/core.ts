import { createUploadthing, type FileRouter } from "uploadthing/next"

const f = createUploadthing()

// Verify all required environment variables are present
const requiredEnvVars = {
  UPLOADTHING_SECRET: process.env.UPLOADTHING_SECRET,
  UPLOADTHING_APP_ID: process.env.UPLOADTHING_APP_ID,
  UPLOADTHING_TOKEN: process.env.UPLOADTHING_TOKEN,
}

// Log environment variable status (without exposing values)
console.log("UploadThing Environment Variables Status:", 
  Object.keys(requiredEnvVars).reduce((acc, key) => ({
    ...acc,
    [key]: requiredEnvVars[key as keyof typeof requiredEnvVars] ? "✓ Set" : "✗ Missing"
  }), {})
)

export const ourFileRouter = {
  pdfUploader: f({ pdf: { maxFileSize: "32MB" } })
    .middleware(async () => {
      // This code runs on your server before upload
      console.log("UploadThing middleware executing...")
      
      // Verify environment variables
      const missingVars = Object.entries(requiredEnvVars)
        .filter(([_, value]) => !value)
        .map(([key]) => key)

      if (missingVars.length > 0) {
        const error = `Missing required environment variables: ${missingVars.join(", ")}`
        console.error(error)
        throw new Error(error)
      }

      console.log("All environment variables verified")
      return { uploadedAt: new Date() }
    })
    .onUploadComplete(async ({ metadata, file }) => {
      // This code RUNS ON YOUR SERVER after upload
      console.log("Upload complete. File details:", {
        url: file.url,
        name: file.name,
        size: file.size,
        uploadedAt: metadata.uploadedAt
      })
      
      return { url: file.url }
    }),
} satisfies FileRouter

export type OurFileRouter = typeof ourFileRouter 