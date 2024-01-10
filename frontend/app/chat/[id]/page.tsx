'use client'

export default function Page({params}:{ params: { id: string } }) {
  return <p>Chat: {params.id}</p>
}