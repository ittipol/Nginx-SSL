"use client"

import { useAppDispatch, useAppSelector } from "@/redux/hook";
import { AppDispatch } from "@/redux/store";

const Header = () => {

    const dispatch:AppDispatch = useAppDispatch()
    const name = useAppSelector((state) => state.user.name);

    return (
        <header className="mb-5">
            <div>Header</div>
            <div>Name: {name}</div>
            <hr />
        </header>
    )
}

export default Header