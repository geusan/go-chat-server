import { API_DOMAIN } from "@/lib/constants/api";
import { PasswordIcon, UsernameIcon } from "@/lib/icons";
import { useCallback, useId, useRef, useState } from "react";

export function SignupButton() {
  const dialogId = useId();
  const dialog = useRef<HTMLDialogElement>(null);
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");
  const openModal = useCallback(() => {
    if (dialog.current) {
      dialog.current.showModal();
    }
  }, []);
  const closeModal = useCallback(() => {
    if (dialog.current) {
      dialog.current.close();
      setUsername("");
      setPassword("");
      setError("");
    }
  }, []);
  const signup = async () => {
    const res = await fetch(API_DOMAIN + "/register", {
      method: "POST",
      headers: {
        "content-type": "application/json",
      },
      body: JSON.stringify({ name: username, password }),
    }).catch((e) => {
      setError(String(e));
      return e;
    });
    const body = await res.json();
    if (Math.ceil(res.status / 100) > 2) {
      setError(body.message);
    } else {
      alert("가입이 완료되었습니다.");
      closeModal();
    }
  };
  return (
    <>
      <a className="btn btn-fluid text-md" onClick={openModal}>
        Signup
      </a>
      <dialog id={dialogId} className="modal" ref={dialog}>
        <div className="modal-box">
          <h3 className="font-bold text-lg mb-2">Sign up form</h3>
          <form className="my-2 space-y-2">
            <label className="form-control w-full max-x-xs">
              <label className="input input-bordered flex items-center gap-2">
                <UsernameIcon />
                <input
                  type="text"
                  className="grow"
                  placeholder="Username"
                  tabIndex={1}
                  value={username}
                  onChange={(e) => setUsername(e.currentTarget.value)}
                />
              </label>
              {error && (
                <div className="label">
                  <span className="label-text-alt text-error">{error}</span>
                </div>
              )}
            </label>
            <label className="input input-bordered flex items-center gap-2">
              <PasswordIcon />
              <input
                type="password"
                placeholder="Type here"
                tabIndex={2}
                value={password}
                onChange={(e) => setPassword(e.currentTarget.value)}
              />
            </label>
          </form>
          <div className="modal-action">
            <button className="btn" onClick={closeModal} tabIndex={4}>
              Close
            </button>
            <button className="btn btn-primary" tabIndex={3} onClick={signup}>
              Sign up
            </button>
          </div>
        </div>
      </dialog>
    </>
  );
}
