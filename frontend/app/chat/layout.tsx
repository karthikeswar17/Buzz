'use client'
import TextBox from "@/components/textbox"
import { FaUserFriends } from "react-icons/fa";
import { AiOutlineSend } from "react-icons/ai"
import { useRouter } from "next/navigation";
export default function Chat({ children }: { children: React.ReactNode }) {
    const router = useRouter()
    return <div className="flex h-screen flex-row">
        <div className=" bg-white m-3 flex gap-2 flex-col w-[25%]">
            <div className="h-18 m-2 flex flex-row justify-between items-center">
                <span className="font-bold text-xl">Conversations</span>
                <button className="hover:bg-gray-200 p-3 rounded-full active:bg-gray-300" onClick={()=>{router.push("/friends")}}><FaUserFriends className="h-7 w-7" /></button>

            </div>
            <div className="w-full">
                <TextBox name="Search" width="full" height="11" onChange={(e: React.ChangeEvent<HTMLInputElement>) => console.log(e.target.value)} />
            </div>
            <div className="overflow-y-auto h-full">
            </div>
        </div>
        <div className="border-r "></div>
        <div className=" w-[80%] m-x-3 h-full flex flex-col">
            
            <div className=" border-b"></div>
            <div className="flex-1 overflow-y-auto">
                {children}
            </div>
            <div>
                <div className="flex flex-row m-3 px-3 gap-2">
                    <TextBox name="SendMessage" width="full" height="11" onChange={(e: React.ChangeEvent<HTMLInputElement>) => console.log(e.target.value)} />
                    <AiOutlineSend className=" h-11 w-11 bg-blue-400 rounded-full text-white p-2 hover:bg-blue-500 active:bg-blue-600" />
                </div>
            </div>
        </div>
    </div>

}