import { Icons } from "../icons";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Modal } from "../ui/modal";
import { Textarea } from "../ui/textarea";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useEffect } from "react";
import { toast } from "sonner";

type CreateIssueProps = {
  closeFn: () => void;
  isOpen: boolean;
};

const createIssueSchema = z.object({
  title: z.string().min(2, "Please provide at least 2 characters").trim(),
  description: z.string().trim().optional(),
});
type CreateIssueSchema = z.infer<typeof createIssueSchema>;

export function CreateIssue({ closeFn, isOpen }: CreateIssueProps) {
  const {
    register,
    handleSubmit,
    reset,
    formState: { errors },
  } = useForm<CreateIssueSchema>({
    resolver: zodResolver(createIssueSchema),
  });

  function onSubmit(data: CreateIssueSchema) {
    console.log(data);
    reset;
  }

  useEffect(() => {
    if (errors.title) {
      toast.error(errors.title.message);
    }
  }, [errors]);

  return (
    <Modal onClose={closeFn} isOpen={isOpen}>
      <form onSubmit={(e) => handleSubmit(onSubmit)(e)} className="grid gap-3">
        <div className="grid gap-1.5 p-6 ">
          <Input
            {...register("title")}
            placeholder="What should we do ?"
            className="text-lg md:min-w-[30rem]"
            variant="empty"
          />

          <Textarea
            {...register("description")}
            placeholder="Some cool description, Tu, TU"
            variant="empty"
            className="h-40 md:min-w-[30rem]"
          />
        </div>
        <div className="flex w-full items-center justify-between border border-t p-4">
          <Button variant="link" size="icon">
            <Icons.Attachment className="size-4" />
          </Button>
          <Button size="sm">Create</Button>
        </div>
      </form>
    </Modal>
  );
}