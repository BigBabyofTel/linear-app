import { deleteAccount } from "@/lib/actions";
import { Button } from "../ui/button";
import { Card, CardContent, CardHeader } from "../ui/card";
import { Modal } from "../ui/modal";

type deleteAccountProps = {
    closeFn: () => void;
    isOpen: boolean;
  };


export default function DeleteAccount({isOpen, closeFn}: deleteAccountProps) {
    
    return (
<Modal isOpen={isOpen} onClose={closeFn}>

<Card>
    <CardHeader><span className="text-2xl">Are you sure?</span>This is process is permenant. All Data will be deleted and cannot be restored</CardHeader>
    <CardContent><Button onClick={() => deleteAccount()}>Delete</Button></CardContent>
</Card>
</Modal>
    )
}