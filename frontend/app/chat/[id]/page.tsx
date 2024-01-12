'use client'

import ChatBubble from "@/components/chatBubble"
import { BACKEND_DOMAIN } from "@/constants/config"
import axios from "axios"
import { config } from "process"
import { useEffect, useState } from "react"

export default function Conversation({params}:{ params: { id: string } }) {
  const [messages, setMessages] = useState([])
  const getMessages=async ()=>{
    const url = `${BACKEND_DOMAIN}/message/v1/conversation/${params.id}/message/list`
    let messages = await axios.get(url, {withCredentials:true}).then(response=>response.data)
    setMessages(messages)
  }
  useEffect(()=>{
    getMessages()
  }, [])
  
  return (<div>
      {messages.map(()=>(<ChatBubble message="asd" user="asd" />))}
    </div>
  )
}