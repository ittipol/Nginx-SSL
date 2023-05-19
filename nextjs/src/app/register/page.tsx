// on top of the file where you are using handleClick because all components in Next 13 by default are server components
"use client";

import { RegisterBody } from "@/models/register";
import { userRegister } from "@/redux/features/user/userSlice";
import { useAppDispatch } from "@/redux/hook";
import { AppDispatch } from "@/redux/store";
import { FormEvent, useRef, useState } from "react";
import { useRouter } from 'next/navigation';

export default function Register() {
    const [error, setError] = useState<string>("")

    const router = useRouter();
    const dispatch:AppDispatch = useAppDispatch()

    const textEmail = useRef<HTMLInputElement>(null)
    const textPassword = useRef<HTMLInputElement>(null)
    const textName = useRef<HTMLInputElement>(null)

    const register = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        console.log('Register form submitted');

        // const body:RegisterBody = {
        //     email: "abc@mail.com",
        //     password: "1234",
        //     name: "New Name"
        // }

        const body:RegisterBody = {
            email: textEmail.current?.value,
            password: textPassword.current?.value,
            name: textName.current?.value
        }

        const res = await dispatch(userRegister(body))
        // const val = res.payload as RegisterResponseResult

        // console.table(res)

        if (res.meta.requestStatus === "fulfilled") {
            router.push("/")
        } else {
            setError("Cannot register user, Try to input new email or password or name")
        }

    }

    return (
        <div className="mx-auto w-full md:w-[400px]">
            <div className="p-6 md:p-4">
                <div className="mb-4">
                    <h2 className="text-xl">Register</h2>
                    <div className="text-red-500">
                        {error}
                    </div>
                </div>
                <form onSubmit={register}>
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
        </div>
    );
}