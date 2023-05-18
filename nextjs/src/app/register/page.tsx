// on top of the file where you are using handleClick because all components in Next 13 by default are server components
"use client";

import { RegisterBody } from "@/models/register";
import { userRegister } from "@/redux/features/user/userSlice";
import { useAppDispatch } from "@/redux/hook";
import { AppDispatch } from "@/redux/store";
import { FormEvent, useRef } from "react";
import { useRouter } from 'next/navigation';

export default function Register() {
    
    const router = useRouter();
    const dispatch:AppDispatch = useAppDispatch()

    const textEmail = useRef<HTMLInputElement>(null)
    const textPassword = useRef<HTMLInputElement>(null)
    const textName = useRef<HTMLInputElement>(null)

    const register = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log('The link was clicked.');

        const body:RegisterBody = {
            email: "abc@mail.com",
            password: "1234",
            name: "New Name"
        }

        // const body:RegisterBody = {
        //     email: textEmail.current?.value,
        //     password: textPassword.current?.value,
        //     name: textName.current?.value
        // }

        const res = await dispatch(userRegister(body))
        // const val = res.payload as ResponseData

        if (res.meta.requestStatus === "fulfilled") {
            router.push("/")
        } else {
            
        }

    }

    return (
        <div className="mx-auto w-full md:w-[400px]">
            <form onSubmit={register} className="p-6 md:p-4">
                <div className="mb-5">
                    <input ref={textEmail} type="email" required className="w-full p-2 text-black outline-0 rounded-lg" />
                </div>
                <div className="mb-5">
                    <input ref={textPassword} type="password" required className="w-full p-2 text-black outline-0 rounded-lg" />
                </div>
                <div className="mb-5">
                    <input ref={textName} type="text" required className="w-full p-2 text-black outline-0 rounded-lg" />
                </div>
                <div>
                    <button role="button" type="submit" className="w-1/5 p-1 rounded-lg bg-blue-800">Register</button>
                </div>
            </form>
        </div>
    );
}