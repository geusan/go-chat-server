import { API_DOMAIN } from "@/lib/constants/api";
import { RootContext } from "@/lib/contexts/root-context";
import { useCallback, useContext, useId, useRef, useState } from "react";

const INIT_VALUE = "";

export function AddChatroomButton() {
  const dialogId = useId();
  const rootContext = useContext(RootContext);
  const [name, setName] = useState(INIT_VALUE);
  const [error, setError] = useState("");
  const dialog = useRef<HTMLDialogElement>(null);
  const openModal = useCallback(() => {
    if (dialog.current) {
      dialog.current.showModal();
    }
  }, []);
  const closeModal = useCallback(() => {
    if (dialog.current) {
      dialog.current.close();
      setName(INIT_VALUE);
      setError("");
    }
  }, []);
  const create = async () => {
    const res = await fetch(API_DOMAIN + "/rooms", {
      method: "POST",
      credentials: "include",
      headers: {
        "content-type": "application/json",
      },
      mode: "cors",
      body: JSON.stringify({ name }),
    });
    const data = await res.json();
    if (res.ok) {
      rootContext.setChatrooms((chatrooms) => [...chatrooms, data]);  
    } else {
      setError(data.message)
    }

  };
  return (
    <>
      <button
        type="button"
        className="btn btn-secondary btn-sm"
        onClick={openModal}
      >
        New Chat
      </button>
      <dialog id={dialogId} ref={dialog}>
        <div className="modal-box bg-base">
          <h3 className="font-bold text-lg">New Chatroom</h3>
          <form className="my-2 space-y-2">
            <label className="form-control w-full max-x-xs">
              <input
                type="text"
                className="input input-bordered flex items-center gap-2"
                placeholder="Username"
                value={name}
                onChange={(e) => setName(e.currentTarget.value)}
              />
              {error && (
                <div className="label">
                  <span className="label-text-alt text-error">{error}</span>
                </div>
              )}
            </label>
          </form>
          <div className="modal-action">
            <button className="btn" onClick={closeModal}>
              Close
            </button>
            <button className="btn btn-primary" onClick={create}>
              Create
            </button>
          </div>
        </div>
      </dialog>
    </>
  );
}
