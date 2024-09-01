"use client";
import React, { useEffect, useState } from "react";
import ChatInput from "./chat-input";
import ChatList from "./chat-list";
import { useChatSocket } from "./use-chat-socket";
import { API_DOMAIN } from "@/lib/constants/api";
import { useParams } from "next/navigation";

function ChatMeesage({
  message,
  isOwner,
  owner,
}: React.PropsWithChildren<{
  message: string;
  isOwner?: boolean;
  owner: string;
}>) {
  return (
    <div>
      {!isOwner && <div>{owner}</div>}
      <div>{message}</div>
    </div>
  );
}

async function getChatroomUrl(_chatId: string | string[]): Promise<[any, any]> {
  const chatId = typeof _chatId === "string" ? _chatId : _chatId[0];
  const res = await fetch(API_DOMAIN + `/rooms/${chatId}/connect`);
  const data = await res.json();
  if (res.ok) {
    return [data, null];
  }
  return [null, data];
}

export default function Chatroom() {
  const { chatId } = useParams();
  const [input, onInputChange] = useState("");
  const [chatroomUrl, setChatroomUrl] = useState("");

  const {} = useChatSocket({
    onMessage: (e) => console.log(e),
    onError: (e) => console.log(e),
    onClose: (e) => console.log(e),
    chatroomUrl,
  });
  useEffect(() => {
    getChatroomUrl(chatId).then(([data, err]) => {
      if (err) {
        alert(err);
        return;
      }
      setChatroomUrl(data.url);
    });
  }, []);
  return (
    <div className="min-h-screen flex bg-base-200 justify-center">
      <div className="flex-1 overflow-y-scroll mb-16 max-w-lg">
        <ChatList
          chats={[
            { content: "Sample message 1", role: "user" },
            { content: "Sample message 2", role: "user" },
            { content: "Sample message 3", role: "owner" },
            { content: "Sample message 4", role: "user" },
            { content: "Sample message 1", role: "user" },
            { content: "Sample message 2", role: "user" },
            { content: "Sample message 3", role: "owner" },
            { content: "Sample message 4", role: "user" },
            { content: "Sample message 1", role: "user" },
            { content: "Sample message 2", role: "user" },
            { content: "Sample message 3", role: "owner" },
            { content: "Sample message 4", role: "user" },
            { content: "Sample message 1", role: "user" },
            { content: "Sample message 2", role: "user" },
            { content: "Sample message 3", role: "owner" },
            { content: "Sample message 4", role: "user" },
          ]}
          isLoading={false}
        />
      </div>
      <ChatInput
        input={input}
        onInputChange={(e) => onInputChange(e.target.value)}
        isLoading={false}
        isFinished={false}
        onSubmit={() => console.log("hey")}
      />
    </div>
  );
}
