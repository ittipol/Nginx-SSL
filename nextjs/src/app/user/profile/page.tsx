"use client"

import { useAppDispatch, useAppSelector } from "@/redux/hook"
import { AppDispatch } from "@/redux/store"
import { useEffect, useState } from "react"

const UserProfile = () => {

    const [msg, setMsg] = useState<string>("")

    const dispatch:AppDispatch = useAppDispatch()
    const accessToken = useAppSelector((state) => state.user.accessToken);

    useEffect(() => {
        
    }, [])

    return (
        <div>
            <div>{msg}</div>
            <div>
                <div>Name:</div>
            </div>
        </div>
    )
}

export default UserProfile