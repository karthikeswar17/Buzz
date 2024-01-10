'use client'
import Button from "@/components/button";
import TextBox from "@/components/textbox";
import { BACKEND_DOMAIN } from "@/constants/config";
import axios from "axios";
import { useRouter } from "next/navigation";
import { useState } from "react";

import { HiOutlineChatAlt } from 'react-icons/hi'
axios.defaults.withCredentials = true
export default function Login() {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const router = useRouter();
    const handleSignUp = () => {
        router.push("/register")
    }

    const onClickLogin = async () => {
        console.log('login', email, password)
        let payload = {
            "email": email,
            "password": password
        }
        let response = await axios.post(`${BACKEND_DOMAIN}/login`, payload).then((response) => {
            let responseBody = response.data
            console.log(responseBody)
            return responseBody
        })
        if(response.status == "Success"){
            router.push("/chat")
        }
    }

    return <div className="flex justify-center items-center h-screen w-screen bg-gray-100">
        <div className="flex flex-col h-fit w-fit rounded-lg p-12 relative bg-white">
        <div className="flex flex-row gap-3">
                <HiOutlineChatAlt className=" text-red-400 h-8 w-8" />
                <span className=" text-lg font-bold ">Login</span>
            </div>
            <div className="flex flex-col h-fit w-fit rounded-lg relative gap-7 bg-white py-8" >
                <TextBox name="Email" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setEmail(e.target.value)} />
                <TextBox name="Password" onChange={(e: React.ChangeEvent<HTMLInputElement>) => setPassword(e.target.value)} />
                <Button name="Sign in" width="auto" onClick={onClickLogin} />
                <div className="border-t border-gray-300 w-full"></div>
            </div>

            {/* <div className="flex flex-row self-center">
                <button className="border-2 border-gray-200 rounded-md box-border p-1 hover:bg-gray-100"><FcGoogle className="h-7 w-7" /></button>
            </div> */}
            <div className="flex flex-row self-center mt-5">
                <span className="">Dont have a account?</span>
                <a className=" text-blue-500 hover:text-blue-700" onClick={handleSignUp}>Sign Up</a>

            </div>
        </div>
    </div>
}