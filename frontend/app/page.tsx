'use client'
import { useRouter } from "next/navigation"

export default function Home() {
  const router = useRouter()
  return (
    <main className="flex min-h-screen">
      <h1>Home</h1>
      <button className=' bg-red-50 h-8' onClick={()=>router.push('/login')}>Login</button>
    </main>
  )
}
