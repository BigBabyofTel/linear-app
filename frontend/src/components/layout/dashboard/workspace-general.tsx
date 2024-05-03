'use client'
import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";

export default function WorkspaceGeneral() {
    // TODO: Need to implement setting a logo, updating workspace name/url, functionality to update button, ability to delete workspace.
    const [workspaceName, setWorkspaceName] = useState("Full Stack Group");
    const [workspaceURL, setWorkspaceURL] = useState("full-stack-group");
    return (
        <div className="flex flex-col gap-6 w-2/5">
            <div>
                <h1 className="text-2xl font-semibold">
                    Workspace
                </h1>
                <p className="text-sm text-muted-foreground">
                    Manage your workspace settings. Your workspace is in the <span className="font-semibold text-white"> United States</span> region.
                </p>
                <hr className="w-full h-[1px] mt-6" />
            </div>
            <div>
                <h2>Logo</h2>
                <div className="flex items-center justify-center size-16 bg-blue-500 rounded mt-6">
                    <label 
                        htmlFor="imageUpload" 
                        className="text-xl">
                        FG
                    </label>
                    <Input 
                        type="file" 
                        accept="image/*" 
                        id="imageUpload" 
                        style={{display: "none"}} 
                    />
                </div>
                <p className="mt-2 text-sm text-muted-foreground">
                    Pick a logo for your workspace. Recommend size is 256x256px.
                </p>
                <hr className="w-full h-[1px] mt-6" />
            </div>
            <div>
                <h1 className="text-md font-semibold">
                    General
                </h1>
                <div className="flex flex-col">
                    <label 
                        htmlFor="workspace-name" 
                        className="text-sm font-semibold">
                        Workspace name
                    </label>
                    <Input 
                        type="text" 
                        id="workspace-name" 
                        className="w-5/12 rounded h-8 pl-2 text-sm text-white" 
                        value={workspaceName} 
                        onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setWorkspaceName(e.target.value)}  
                    />
                </div>
                <div className="flex flex-col mt-6">
                    <label 
                        htmlFor="workspace-url" 
                        className="text-sm font-semibold">
                        Workspace URL
                    </label>
                    <Input 
                        type="text" 
                        id="workspace-url" 
                        className="w-5/12 rounded h-8 pl-2 text-sm text-white" 
                        value={"linear.app/" + workspaceURL} 
                        onChange={(e: { target: { value: React.SetStateAction<string>; }; }) => setWorkspaceURL(e.target.value)}
                    />
                </div>
                <Button className="w-[80px] h-[35px] bg-indigo-400 mt-6 rounded text-sm hover:brightness-110">
                    Update
                </Button>
                <hr className="w-full h-[1px] mt-6" />
            </div>
            <div>
                <h1 className="text-md font-semibold">
                    Delete workspace
                </h1>
                <p className="mt-4 text-sm text-muted-foreground">
                    If you want to permanently delete this workspace and all of its data, including but not limited to users, issues, and comments, you can do so below.
                </p>
                <Button variant="destructive" className="w-[200px] h-[35px] mt-4 rounded text-sm hover:brightness-110">
                    Delete this workspace
                </Button>
            </div>
        </div>
    )
}