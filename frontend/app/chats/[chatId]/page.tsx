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
  const res = await fetch(API_DOMAIN + `/rooms/${chatId}/open`, { credentials: 'include', mode: 'cors'});
  const data = await res.json();
  if (res.ok) {
    return [data, null];
  }
  return [null, data];
}

async function getMe(): Promise<[any, any]> {
  const res = await fetch(API_DOMAIN + `/me`, { credentials: 'include', mode: 'cors'});
  const data = await res.json();
  if (res.ok) {
    return [data, null];
  }
  return [null, data];
}

export default function Chatroom() {
  const { chatId } = useParams();
  const [input, onInputChange] = useState("");
  const [user, setUser] = useState<any>();
  const [chatroomUrl, setChatroomUrl] = useState("");
  const [chats, setChats] = useState<any[]>([]);
  const { send } = useChatSocket({
    onMessage: (e) => {
      const msg = JSON.parse(e.data);
      if (msg.role === user?.name) {
        return;
      }
      setChats(s => [...s, msg])
    },
    onError: (e) => alert("Error!"),
    onClose: (e) => alert("Chatting is closed"),
    chatroomUrl,
  });
  useEffect(() => {
    getChatroomUrl(chatId).then(([data, err]) => {
      if (err) {
        alert(JSON.stringify(err));
        return;
      }
      setChatroomUrl(data.url);
    });
    getMe().then(([data, err]) => {
      if (err) {
        alert(JSON.stringify(err));
        return;
      }
      setUser(data);
    })
  }, []);
  return (
    <div className="min-h-screen flex bg-base-200 justify-center">
      <div className="flex-1 overflow-y-scroll mb-16 max-w-lg">
        <ChatList
          chats={chats}
          isLoading={false}
        />
      </div>
      <ChatInput
        input={input}
        onInputChange={(e) => onInputChange(e.target.value)}
        isLoading={false}
        isFinished={false}
        onSubmit={() => {send(JSON.stringify({role: user?.name, content: input}));onInputChange('')}}
      />
    </div>
  );
}
