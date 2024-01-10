'use client'
import Button from "@/components/button";
import FriendCard from "@/components/friendCard";
import TextBox from "@/components/textbox";
import { BACKEND_DOMAIN } from "@/constants/config";
import axios from "axios";
import { IoMdArrowRoundBack } from "react-icons/io";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import FriendReqeustCard from "@/components/FriendRequestCard";



export default function AddFriend() {
    const [friendTag, setFriendTag] = useState("")
    const [friendRequests, setfriendRequests] = useState([])
    const router = useRouter()
    //sendFriendRequest
    const sendFriendRequest = async () => {
        let url = `${BACKEND_DOMAIN}/friend/request`
        let paylaod = {
            "tag": friendTag
        }
        return await axios.post(url, paylaod, {withCredentials:true}).then((response) => response.data).catch((error) => alert(error.response.data.message))

    }
    //get all my friend requests
    const getFriendRequests = async () => {
        let url = `${BACKEND_DOMAIN}/friend/request`
        let responseBody = await axios.get(url, {withCredentials:true}).then((response) => response.data).catch((error) => alert(error.response.data.message))
        setfriendRequests(responseBody)
    }
    const removeFriendRequest=(id:string)=>{
        setfriendRequests(friendRequests.filter((friendData: { "_id": string, "tag": string, "name": string, "email": string, "friend_request":string })=>{friendData._id != id}))

    }
    const acceptFriendRequest = async (id:string) => {
        let url = `${BACKEND_DOMAIN}/friend/request/${id}/accept`
        await axios.get(url, {withCredentials:true}).then((response) => response.data).catch((error) => alert(error.response.data.message))
        removeFriendRequest(id)

    }
    const declineFriendRequest = async (id:string) => {
        let url = `${BACKEND_DOMAIN}/friend/request/${id}/decline`
        await axios.get(url, {withCredentials:true}).then((response) => response.data).catch((error) => alert(error.response.data.message))
        removeFriendRequest(id)
        
    }
    useEffect(() => {
        getFriendRequests()
    }, [])

    return (<div className="flex flex-col p-5 gap-5">
        <div className=" h-14 flex items-center ">
            <button className="hover:bg-gray-200 p-3 rounded-full active:bg-gray-300" onClick={() => { router.replace("/friends") }}><IoMdArrowRoundBack className="h-6 w-6" /></button>
            <div className="flex flex-col">
                <span className=" font-bold text-lg">Add Friends</span>
                <span className=" font-light text-sm  text-gray-400">you can add friends using thier email</span>
            </div>
        </div>
        <div className="flex flex-col gap-10">
            <div className="flex flex-row gap-3">
                <TextBox name="Friend Tag" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setFriendTag(e.target.value)} />
                <Button name="Send Friend Request" width="auto" onClick={() => { sendFriendRequest() }} />
            </div>
            <div className="flex flex-col gap-3 ">
                <span className=" font-bold text-lg">Friend Requests</span>
                <span className=" border-b-2"></span>
                <div className=" overflow-y-auto">
                    {
                        friendRequests.length>0 ? friendRequests.map((friendData: { "_id": string, "tag": string, "name": string, "email": string, "friend_request":string }) => {
                            return <FriendReqeustCard id={friendData._id} name={friendData.name} tag={friendData.tag} email={friendData.email} friendRequest={friendData.friend_request} acceptCallback={acceptFriendRequest} declineCallback={declineFriendRequest}/>
                        }) : <span className=" text-sm text-gray-400"> No New Requests</span>
                    }
                </div>
            </div>
        </div>
    </div>)
}