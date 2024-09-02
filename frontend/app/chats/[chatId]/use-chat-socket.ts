import { useEffect, useRef, useState } from "react";

type Props = {
  chatroomUrl: string;
  onMessage(e: MessageEvent<any>): void;
  onClose(e: CloseEvent): void;
  onError(e: Event): void;
};

export const useChatSocket = ({
  onMessage,
  chatroomUrl,
  onClose,
  onError,
}: Props) => {
  const [conn, setConn] = useState<WebSocket>();
  
  useEffect(() => {
    if (conn && conn.CONNECTING) return;
    if (!chatroomUrl) return;

    const ws = new WebSocket(chatroomUrl);
    ws.onclose = function (e) {
      onClose(e);
    };
    ws.onmessage = function (e) {
      onMessage(e);
    };
    ws.onerror = function (e) {
      onError(e);
    };
    setConn(ws);
    return () => {
      ws.close();
    }
  }, [chatroomUrl]);

  const close = () => { conn && conn.close() };
  const send = (msg: string) => { conn && conn.send(msg) };
  return { close, send }
};
