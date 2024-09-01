"use client";

import { PropsWithChildren } from "react";
import { motion, AnimatePresence } from "framer-motion";

type Props = PropsWithChildren<{
  type: "info" | "error";
  show: boolean;
}>;

const ToastTypeClass = {
  info: "alert-info",
  error: "alert-error",
};

const Toast = ({ type, show, children }: Props) => {
  return (
    <AnimatePresence>
      {show && (
        <motion.div
          initial={{ opacity: 0, y: 10 }}
          animate={{ opacity: 1, y: 0 }}
          exit={{ opacity: 0, y: 10 }}
          className="toast animate-none"
        >
          <div className={`alert ${ToastTypeClass[type]}`}>
            <div>{children}</div>
          </div>
        </motion.div>
      )}
    </AnimatePresence>
  );
};

export default Toast;