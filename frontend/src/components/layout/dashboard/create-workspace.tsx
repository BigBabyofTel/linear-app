'use client'
import React, { useState } from "react";
import { Input } from "@/components/ui/input";
import { Modal } from "@/components/ui/modal";
import { Button } from "@/components/ui/button";
import { API } from "@/lib/utils";

type Token = {
    token: string;
  };

export default function CreateWorkspace({token}: Token) {
    const [workspaceName, setWorkspaceName] = useState("");
    const [workspaceURL, setWorkspaceURL] = useState("");
    const [image, setImage] = useState("");
    const [isOpen, setIsOpen] = useState(false);
    const [error, setError] = useState("");

    const handleCloseModal = () : void => {
        setIsOpen(false);
    }

    const handleWorkspaceName = (newName : string) : void => {
        setWorkspaceName(newName);
    }

    const handleWorkspaceURL = (newURL : string) : void => {
        setWorkspaceURL(newURL);
    }

    const isAuthenticated = () => {
        try {
            if (token) {
                createWorkspace(workspaceName, workspaceURL, image);
            } else {
                throw new Error("User is not authenticated");
            }
        } catch (error) {
            console.log(error);
        }
    }

    const createWorkspace = async(name: string, slug: string, image: string) => {
        try {
            const response = await API.post("/v1/workspaces", {
                Name: name,
                Slug: slug,
                Image: image
            }, {
                headers: {
                    Authorization: `Bearer ${token}`
                },
            });

            console.log(response.data);

        } catch (error) {
            console.log(error);
        }
    }

    const CreateWorkspace = () => {
        return (
            <div className="gap-12 flex flex-col">
                <div className="flex flex-col">
                    <label 
                        htmlFor="workspace-name" 
                        className="text-sm font-semibold">
                        Workspace name
                    </label>
                    <Input 
                        type="text" 
                        id="workspace-name" 
                        className="w-full rounded h-8 pl-2 text-sm text-white mt-1" 
                        value={workspaceName} 
                        onChange={(e) => handleWorkspaceName(e.target.value)}
                    />
                </div>
                <div className="flex flex-col">
                    <label 
                        htmlFor="workspace-url" 
                        className="text-sm font-semibold">
                        Workspace URL
                    </label>
                    <Input 
                        type="text" 
                        id="workspace-url" 
                        className="w-full rounded h-8 pl-2 text-sm text-white mt-1"
                        value={workspaceURL} 
                        onChange={(e) => handleWorkspaceURL(e.target.value)}
                    />
                </div>
                <Button className="full h-[35px] bg-indigo-400 mt-6 rounded text-sm hover:brightness-110" onClick={isAuthenticated}>Create Workspace</Button>
            </div>
        )
    }

    return (
        <div>
            <Button onClick={() => setIsOpen(true)}>
                Open Modal
            </Button>
            <Modal children={<CreateWorkspace />} isOpen={isOpen} onClose={handleCloseModal} className="h-[400px] w-3/12 flex flex-col items-center justify-center"/>
        </div>
    )
}