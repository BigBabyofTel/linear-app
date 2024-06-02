import React from "react";
import WorkspaceGeneral from "@/components/layout/dashboard/workspace-general";
import { useMutation } from "@tanstack/react-query";
import { cookies } from "next/headers";

export default function General() {
    
    // TODO: Make sure to implement DELETE workspace route(ONLY ADMIN CAN DO THIS)
    const cookie = cookies();
    const token  = cookie.get("token");
    return (
        <div className="flex flex-col justify-center items-center w-full mt-20">
            <WorkspaceGeneral token={token?.value!} />
        </div>
    )
}