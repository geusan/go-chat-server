"use client";

import { LoginButton } from "./login-button";
import { SignupButton } from "./signup-button";
import { RootContext } from "@/lib/contexts/root-context";

export function AuthButton() {
  return (
    <RootContext.Consumer>
      {({isLoggedIn}) => (
        <div>
          {isLoggedIn ? (
            <a className="btn btn-fluid text-md">Logout</a>
          ) : (
            <div className="space-x-2">
              <SignupButton />
              <LoginButton />
            </div>
          )}
        </div>
      )}
    </RootContext.Consumer>
  );
}
