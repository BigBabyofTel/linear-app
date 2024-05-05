import React, { useEffect, useRef } from "react";
import EditorJS, { OutputData, ToolConstructable } from "@editorjs/editorjs";
import Header from "@editorjs/header";
import AIText from "@alkhipce/editorjs-aitext"; // could not find types for this package
import { API } from "@/lib/utils";

type EditorProps = {
  onChange: (data: OutputData | undefined) => void;
};

export function Editor({ onChange }: EditorProps) {
  const editorContainerRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (!editorContainerRef.current) return;

    const editor = new EditorJS({
      holder: "editorjs",
      tools: {
        header: Header,
        aiText: {
          class: AIText as unknown as ToolConstructable,
          config: {
            callback: (text: string) => {
              console.log(text);
              return API.post("/ai/autocomplete", { text })
                .then((response) => {
                  return response.data.data?.choices[0].message.content;
                })
                .catch((error) => {
                  if (error.response) {
                    console.error("Error response:", error.response.data);
                    console.error("Error status:", error.response.status);
                    console.error("Error headers:", error.response.headers);
                  } else if (error.request) {
                    console.error("Error request:", error.request);
                  } else {
                    console.error("Error message:", error.message);
                  }
                  console.error("Error config:", error.config);
                });
            },
          },
        },
      },

      onChange: async () => {
        const data = await editor.save();
        onChange(data);
      },
    });

    return () => {
      editor.destroy;
    };
  }, []);

  return <div ref={editorContainerRef} id="editorjs"></div>;
}
