'use client'
import Button from "@/components/button";
import FriendCard from "@/components/friendCard";
import TextBox from "@/components/textbox";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { IoMdArrowRoundBack } from "react-icons/io";
import axios from "axios";
import { BACKEND_DOMAIN } from "@/constants/config";

const getFriends = async () => {
    let url = `${BACKEND_DOMAIN}/friend/list`
    return await axios.get(url).then((response) => response.data)

}

export default function Friend() {
    const [friendName, setFriendName] = useState("")
    const [friendList, setFriendList] = useState([])
    const router = useRouter()

    useEffect(() => {
        const getFriends = async () => {
            let url = `${BACKEND_DOMAIN}/friend/list`
            let params = {
                "q": friendName
            }
            let responseBody = await axios.get(url, {withCredentials:true, params:params}).then((response) => response.data)
            setFriendList(responseBody)
        }
        getFriends()
    }, [friendName])


    return (<div className="flex flex-col p-5 gap-5">
        <div className=" h-14 flex items-center gap-3">
            <button className="hover:bg-gray-200 p-3 rounded-full active:bg-gray-300" onClick={() => { router.replace("/chat") }}><IoMdArrowRoundBack className="h-6 w-6" /></button>
            <div className="flex flex-col ">
                <span className=" font-bold text-lg">Friends</span>
                <span className=" font-light text-sm  text-gray-400">you can see all the existing Friends</span>
            </div>
        </div>
        <div className="flex flex-row gap-3">
            <TextBox name="Search" width="full" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setFriendName(e.target.value)} />
            <Button name="AddFriend" width="auto" onClick={() => { router.push("/friends/add") }} />
        </div>
        <div>
            {
                friendList.map((friendData: { "_id": string, "tag": string, "name": string, "email": string }) => {
                    return <FriendCard id={friendData._id} name={friendData.name} tag={friendData.tag} email={friendData.email} />
                })
            }
        </div>

    </div>)
}