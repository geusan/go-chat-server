"use client";

import { useRef } from "react";
import { useInView } from "framer-motion";
import { AuthButton } from "./auth/auth-button";
import { RootContext, useRootContextState } from "@/lib/contexts/root-context";
import { ChatroomList } from "./chats/chatroom-list";


export default function Episodes({}: {}) {
  const inViewRef = useRef<HTMLDivElement>(null);
  const isInView = useInView(inViewRef);
  const rootState = useRootContextState();

  return (
    <RootContext.Provider value={rootState}>
      <main className="px-2 min-h-screen">
        <div className="navbar bg-base-100">
          <div className="flex-1">
            <a className="btn btn-ghost text-xl">Chatroom</a>
          </div>
          <div className="flex-none">
            <AuthButton />
          </div>
        </div>
        <ul className="flex flex-col justify-center">
          <ChatroomList />
        </ul>
      </main>
    </RootContext.Provider>
  );
}