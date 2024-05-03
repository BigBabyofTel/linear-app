//
import React from "react";
import { Button } from "@/components/ui/button";
import { Modal } from "@/components/ui/modal";
import CreateWorkspace from "@/components/layout/dashboard/create-workspace";
import { UserRecord } from "@/hooks/useSession";
import { API } from "@/lib/utils";
import { cookies } from "next/headers";

export default function Overview() {
    const cookie = cookies();
    const token  = cookie.get("token");

    return (
        <div className="flex flex-col justify-center items-center w-full mt-20">
            <CreateWorkspace token={token?.value!} />
        </div>
    )
}