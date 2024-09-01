import classNames from "classnames";
import { useEffect, useRef } from "react";

type Props = {
  chats: { content: string, role: string }[];
  isLoading: boolean;
};

const ChatList = ({ chats, isLoading }: Props) => {
  const container = useRef<HTMLDivElement>(null);

  useEffect(() => {
    if (chats.length > 0) {
      container.current?.scrollTo({
        top: container.current.scrollHeight,
        behavior: "smooth",
      });
    }
  }, [chats]);

  return (
    <div ref={container} className="w-full h-full flex-1 overflow-auto py-2">
      <div className="chat chat-start">
        <div className={classNames("chat-bubble")}></div>
      </div>
      {chats.map((chat, i) => {
        return (
          <div
            key={i}
            className={classNames(
              "chat",
              chat.role === "user" ? "chat-end" : "chat-start"
            )}
          >
            <div
              className={classNames(
                "chat-bubble",
                chat.role === "user" && "chat-bubble-primary",
                "whitespace-pre-wrap"
              )}
            >
              {chat.content}
            </div>
          </div>
        );
      })}
      {isLoading && (
        <div className="chat chat-start">
          <div className="chat-image avatar">
            <div className="w-10 rounded-full">
              {/* <Image src={Icon} alt="icon" /> */}
            </div>
          </div>
          <div className="chat-bubble h-full">
            <div className="flex justify-between gap-2 items-center h-full">
              {[...Array(3)].map((_, i) => (
                <div
                  key={i}
                  className={`bg-primary-content w-2 h-2 rounded-full animate-bounce`}
                  style={{
                    animationDelay: `${(i + 1) * 100}ms`,
                  }}
                ></div>
              ))}
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default ChatList;