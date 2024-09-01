import { useEffect, useState } from "react";

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
    if (chatroomUrl) {
      const connection = new WebSocket(chatroomUrl);
      setConn(connection);
      return () => {
        if (connection) connection.close();
      }
    }
  }, [chatroomUrl]);

  useEffect(() => {
    if (conn) {
      conn.onclose = function (e) {
        onClose(e);
      };
      conn.onmessage = function (e) {
        onMessage(e);
        console.log(e.data);
      };
      conn.onerror = function (e) {
        onError(e);
      };
    }
  }, [conn]);

  const close = () => { conn && conn.close() };
  const send = (msg: string) => { conn && conn.send(msg) };
  return { close, send }
};
