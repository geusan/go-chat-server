'use client'
import Toast from "./toast";
import useTimerState from "./user-timer-state";
import { useRef, useState } from "react";
import TextareaAutosize from "react-textarea-autosize";

type Props = {
  input: string;
  isFinished: boolean;
  isLoading: boolean;
  onInputChange: (e: React.ChangeEvent<HTMLTextAreaElement>) => void;
  onSubmit: (e: React.FormEvent<HTMLFormElement>) => void;
};

const MAX_INPUT_LENGTH = 100;

const ChatInput = ({
  input,
  onInputChange,
  isFinished,
  isLoading,
  onSubmit,
}: Props) => {
  const [toast, setToast] = useTimerState<boolean>({ initialState: false });
  const canSubmit = !isLoading && !isFinished;

  const handleInputChange = (e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const { value } = e.target;

    if (value.length > MAX_INPUT_LENGTH) {
      setToast(true);
    }

    onInputChange(e);
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    if (!canSubmit) {
      return;
    }

    onSubmit(e);
  };

  const handleKeyDown = (e: React.KeyboardEvent<HTMLTextAreaElement>) => {
    // chrome IME 버그
    if (e.keyCode === 229) {
      return;
    }

    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      // https://stackoverflow.com/a/73471976
      // .submit() 직접 트리거는 onSubmit 트리거 안함
      e.currentTarget.form?.dispatchEvent(
        new Event("submit", { cancelable: true, bubbles: true })
      );
    }
  };

  return (
    <div className="px-4 pb-4 pt-2 fixed w-full bottom-0 bg-base-200 max-w-lg">
      <form className="flex gap-2 items-center w-full" onSubmit={handleSubmit}>
        <TextareaAutosize
          placeholder="여기에 질문하세요."
          className="input w-full p-2"
          value={input}
          onChange={handleInputChange}
          onKeyDown={handleKeyDown}
          maxRows={3}
        />
        <button
          type="submit"
          className="btn btn-primary flex-shrink-0"
          disabled={isLoading || isFinished}
        >
          보내기
        </button>
      </form>
      <Toast type="error" show={!!toast}>
        질문은 {MAX_INPUT_LENGTH}자 이내로 입력가능해요!
      </Toast>
    </div>
  );
};

export default ChatInput;