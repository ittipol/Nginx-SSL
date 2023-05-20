// on top of the file where you are using handleClick because all components in Next 13 by default are server components
"use client";

import { RegisterBody } from "@/models/register";
import { userLogin, userProfile, userRegister } from "@/redux/features/user/userSlice";
import { useAppDispatch, useAppSelector } from "@/redux/hook";
import { AppDispatch } from "@/redux/store";
import { FormEvent, useRef, useState } from "react";
import { useRouter } from 'next/navigation';
import { LoginBody } from "@/models/login";

export default function Login() {
    const [error, setError] = useState<string>("")

    const router = useRouter();
    const dispatch:AppDispatch = useAppDispatch()
    const accessToken = useAppSelector((state) => state.user.accessToken);

    const textEmail = useRef<HTMLInputElement>(null)
    const textPassword = useRef<HTMLInputElement>(null)

    const login = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log('Login form submitted');

        const body:LoginBody = {
            email: "abc@mail.com",
            password: "1234",
        }

        const res = await dispatch(userLogin(body))

        // console.table(res)

        if (res.meta.requestStatus === "fulfilled") {
            // router.push("/")

            await dispatch(userProfile())

        } else {
            setError("Invalid username or password")
        }
    }

    return (
        <div className="mx-auto w-full md:w-[400px]">
            <h1>{accessToken}</h1>
            <div className="p-6 md:p-4">
                <div className="mb-4">
                    <h2 className="text-xl">Login</h2>
                    <div className="text-red-500">
                        {error}
                    </div>
                </div>
                <form onSubmit={login}>
                    <div className="mb-5">
                        <input ref={textEmail} type="email" className="w-full p-2 text-black outline-0 rounded-lg" />
                    </div>
                    <div className="mb-5">
                        <input ref={textPassword} type="password" className="w-full p-2 text-black outline-0 rounded-lg" />
                    </div>
                    <div>
                        <button role="button" type="submit" className="w-1/5 p-1 rounded-lg bg-blue-800">Login</button>
                    </div>
                </form>
            </div>
            <button onClick={() => dispatch(userProfile())}>Get profile</button>
        </div>
    );
}