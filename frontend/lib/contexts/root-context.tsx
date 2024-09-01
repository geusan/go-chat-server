import { createContext, Dispatch, SetStateAction, useContext, useState } from "react";

type RootContextProps = {
  chatrooms: any[];
  setChatrooms: Dispatch<SetStateAction<any[]>>;
  isLoggedIn: boolean;
  setIsLoggedIn: Dispatch<SetStateAction<boolean>>;
}

export const RootContext = createContext<RootContextProps>({
  chatrooms: [],
  setChatrooms: () => {},
  isLoggedIn: false,
  setIsLoggedIn: () => {},
});

export const useRootContextState = (): RootContextProps => {
  const [chatrooms, setChatrooms] = useState<any[]>([]);
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  return {
    chatrooms,
    setChatrooms,
    isLoggedIn,
    setIsLoggedIn,
  };
};

export const useRootContext = () => useContext(RootContext);
