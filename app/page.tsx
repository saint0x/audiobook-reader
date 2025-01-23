import { Metadata } from "next"
import dynamic from "next/dynamic"

export const metadata: Metadata = {
  title: "PDF Player",
  description: "A modern PDF reader with text-to-speech capabilities",
}

const HomePage = dynamic(() => import("../components/HomePage"), {
  ssr: false,
})

export default function Page() {
  return <HomePage />
}

