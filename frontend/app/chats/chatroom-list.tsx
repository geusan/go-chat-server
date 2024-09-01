import { API_DOMAIN } from "@/lib/constants/api";
import { RootContext } from "@/lib/contexts/root-context";
import { useContext, useEffect } from "react";
import { AddChatroomButton } from "./add-chatroom-button";
import Link from "next/link";

async function getChatrooms(): Promise<[any[], boolean]> {
  const res = await fetch(API_DOMAIN + "/rooms", {
    credentials: "include",
    mode: "cors",
  });
  if (res.ok) {
    const chatrooms = await res.json();
    return [chatrooms, true];
  }
  return [[], false];
}

export function ChatroomList() {
  const rootState = useContext(RootContext);
  const { setIsLoggedIn, chatrooms, setChatrooms } = rootState;
  useEffect(() => {
    getChatrooms().then(([_chatrooms, _isLoggedIn]) => {
      setIsLoggedIn(_isLoggedIn);
      setChatrooms(_chatrooms);
    });
  }, []);
  return (
    <div className="bg-base">
      <div className="flex py-2">
        <div className="flex-1"></div>
        <div className="flex-none">
          <AddChatroomButton />
        </div>
      </div>
      <div className="divider">Chatroom list</div>
      {chatrooms.map((result, i) => (
        <li className="p-2 flex gap-4 align-items-center" key={result.id}>
          <b>{i + 1}</b>
          <div className="flex flex-col">
            <h2 className="text-lg">{result.name}</h2>
            {/* <p className="text-base-300 line-clamp-2">{result.limit}</p> */}
          </div>
          <div className="grow" />
          <form className="flex-shrink-0 flex items-center">
            <Link href={`/chats/${result.id}`}>
              <button className="btn btn-primary btn-sm" type="button">
                시작하기
              </button>
            </Link>
          </form>
        </li>
      ))}
    </div>
  );
}
