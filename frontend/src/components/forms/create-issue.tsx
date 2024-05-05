import { Icons } from "../icons";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { Modal } from "../ui/modal";
import { Editor } from "../editor";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import * as z from "zod";
import { useEffect } from "react";
import { toast } from "sonner";
import { useMutation } from "@tanstack/react-query";
import { API } from "@/lib/utils";

type CreateIssueProps = {
  closeFn: () => void;
  isOpen: boolean;
  authToken: string;
};

const createIssueSchema = z.object({
  title: z.string().min(2, "Please provide at least 2 characters").trim(),
  description: z.any(),
});
type CreateIssueSchema = z.infer<typeof createIssueSchema>;

export function CreateIssue({ closeFn, isOpen, authToken }: CreateIssueProps) {
  const {
    register,
    handleSubmit,
    reset,
    setValue,
    formState: { errors },
  } = useForm<CreateIssueSchema>({
    resolver: zodResolver(createIssueSchema),
  });

  const { mutate } = useMutation({
    mutationFn: (data: CreateIssueSchema) => {
      return API.post("/issues", data, {
        headers: {
          Authorization: `Bearer ${authToken}`,
        },
      });
    },
    onSuccess: () => {
      toast.success("Issue created successfully");
      closeFn();
    },
    onError: (error) => {
      toast.error("Failed to create issue");
    },
  });

  function onSubmit(data: CreateIssueSchema) {
    const payload = {
      description: data.description,
      title: data.title,
      status: "Todo",
      priority: "Low",
    };
    mutate(payload);
  }

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

          <Editor
            onChange={(data) => {
              setValue("description", JSON.stringify(data));
            }}
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
