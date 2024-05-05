'use client'
import React, { useState } from "react";
import Link from "next/link";
import { Button } from "@/components/ui/button";
import { Switch } from "@/components/ui/switch";

export default function WorkspaceSecurity() {
    // TODO: Need to implement the Google and Email toggles. Need to allow setting a domain.
    const [googleToggle, setGoogleToggle] = useState(false);
    const [emailToggle, setEmailToggle] = useState(false);
    return (
        <div className="flex flex-col gap-6 w-2/5">
            <div>
                <h1 className="text-2xl font-semibold">
                    Security
                </h1>
                <p className="text-sm text-muted-foreground">
                    Manage your workspace's security and how its members authenticate
                </p>
                <hr className="w-full h-[1px] mt-6" />
            </div>
            <div className="flex justify-between items-start">
                <div className="flex flex-col">
                    <h2>
                        Allowed Email Domains
                    </h2>
                    <p className="text-sm text-muted-foreground">
                    Anyone with an email address at these domains is allowed 
                    to sign up for this workspace.&nbsp;
                        <span className="font-semibold text-white">
                            <Link href={""}>
                                Learn more &#x2192;
                            </Link>
                        </span>
                    </p>
                </div>
                <Button className="w-[100px] h-[35px] bg-indigo-400 rounded text-sm hover:brightness-110 ">
                    Add domain
                </Button>
            </div>
            <hr className="w-full h-[1px] mt-2" />
            <div className="flex justify-between items-start">
                <div className="flex flex-col">
                    <h2>
                        Google
                    </h2>
                    <p className="text-sm text-muted-foreground">
                        Allow logins through Google's single sign-on functionality.
                    </p>
                </div>
                <Switch />
            </div>
            <hr className="w-full h-[1px] mt-2" />
            <div className="flex justify-between items-start">
                <div className="flex flex-col">
                    <h2>
                        Email code
                    </h2>
                    <p className="text-sm text-muted-foreground">
                        Allow passwordless logins through magic links or a code delivered over email.
                    </p>
                </div>
                <Switch />
            </div>
        </div>
    )
}