"use client"
import Button from "@/components/button";
import TextBox from "@/components/textbox";
import { useRouter } from "next/navigation";
import { useState } from "react";
import { FcGoogle } from "react-icons/fc"
import { HiOutlineChatAlt } from 'react-icons/hi'

export default function Register() {
    const router = useRouter();
    const handleSignIn = () => {
        router.replace("/login")
    }

    let [displayName, setDisplayName] = useState("")
    let [username, setUsername] = useState("")
    let [email, setEmail] = useState("")
    let [password, setPassword] = useState("")


    const onClickRegister = () => {
        console.log("Register clicked", displayName, username, email, password)
    }



    return <div className="flex justify-center items-center h-screen w-screen bg-gray-100">
        <div className="flex flex-col h-fit w-fit rounded-lg p-12 relative bg-white">
            <div className="flex flex-row gap-3">
                <HiOutlineChatAlt className=" text-red-400 h-8 w-8" />
                <span className=" text-lg font-bold ">Register</span>
            </div>
            <div className="flex flex-col h-fit w-fit rounded-lg relative gap-7 bg-white py-8" >
                <TextBox name="Display Name" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setDisplayName(e.target.value)} />
                <TextBox name="Username" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setUsername(e.target.value)} />
                <TextBox name="Email" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)} />
                <TextBox name="Password" type="password" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)} />
                <Button name="Sign Up" width="auto" onClick={onClickRegister} />
                <div className="border-t border-gray-300 w-full"></div>
            </div>
            {/* <div className="flex flex-row self-center">
                <button className="border-2 border-gray-200 rounded-md box-border p-1 hover:bg-gray-100"><FcGoogle className="h-7 w-7" /></button>
            </div> */}
            <div className="flex flex-row self-center mt-5">
                <span className="">Already have a account?</span>
                <a className=" text-blue-500 hover:text-blue-700" onClick={handleSignIn}>Sign In</a>

            </div>
        </div>
    </div>
}