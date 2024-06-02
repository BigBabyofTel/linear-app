'use client'
import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useMutation } from "@tanstack/react-query";
import { API } from "@/lib/utils";
import { toast } from "sonner";
import { useRouter } from "next/navigation";
import { GoError } from "@/types/errors.types";
import { z } from "zod";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";

type Token = {
    token: string;
}

export default function WorkspaceGeneral({token}: Token) {
    // TODO: Need to implement setting a logo, updating workspace name/url, functionality to update button, ability to delete workspace.
    const [workspaceName, setWorkspaceName] = useState("Full Stack Group");
    const [workspaceURL, setWorkspaceURL] = useState("full-stack-group");
    const [error, setError] = useState("");
    const router = useRouter();
    const updateSchema = z
    .object({
        name: z.string().min(1, "Required"),
        slug: z
            .string()
            .refine(async (slug) => await isSlugUnique(slug), {
                message: "URL must be unique",
            }),
    })

    type UpdateData = z.infer<typeof updateSchema>;
    type UpdatePayload = UpdateData;

    const {
        handleSubmit,
        formState: { errors },
    } = useForm<UpdateData>({
        resolver: zodResolver(updateSchema),
        defaultValues: {
            name: "",
            slug: ""
        }
    })

    const { mutate, isPending } = useMutation({
        mutationFn: (slug) => API.delete(`/v1/workspaces/${slug}`),
        onSuccess: () => {
            toast.success("Workspace deleted successfully.");
            void router.push("/");
        },
        onError: (e) => {
            const error = e as GoError;

            if (error.response?.status === 401) {
                toast.error("Unauthorized access unable to fulfill request.");
            } else {
                toast.error("Unable to delete workspace please try again.");
            }
        }
    })

    const isAuthenticated = () => {
        try {
            if (token) {
                // handleUpdate(payload);
                return;
            } else {
                throw new Error("User is not authenticated");
            }
        } catch (error) {
            console.log(error);
            toast.error("User is not authenticated");
        }
    }

    const handleDelete = () => {
        mutate();
    }

    const isSlugUnique = async (slug: string) : Promise<boolean> => {
        try {
            const response = await API.get(`/v1/workspaces/${slug}`);

            if (response.status === 422) {
                throw Error("Slug not unique");
            } else if (response.status !== 200) {
                throw Error("Server error");
            }
            
            return true;
        } catch (error) {
            return false;
        }
    }

    const handleUpdate = useMutation({
        mutationFn: (data: UpdatePayload) => API.patch(`v1/workspaces/{data.slug}`, 
        {
            Name: data.name,
            Slug: data.slug
        }, 
        { headers: { Authorization: `Bearer ${token}` } }),
        onSuccess: () => {
            toast.success("Workspace updated successfully.");
        },
        onError: (e) => {
            const error = e as GoError;

            toast.error(e.message);
        }
    })

    function onSubmit(data: UpdateData) {
        const payload: UpdatePayload = {
            name: data.name,
            slug: data.slug
        };

        isAuthenticated();
        handleUpdate.mutate(payload);
    }

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
                <form onSubmit={handleSubmit(onSubmit)}>
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
                </form>
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