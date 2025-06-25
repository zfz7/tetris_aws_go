import React from "react";
import ReactDOM from "react-dom/client";
import "./index.css";
import App from "./App";
import reportWebVitals from "./reportWebVitals";
import { getInfo } from "./utils/client";
import { Amplify } from "aws-amplify";
import { Authenticator } from "@aws-amplify/ui-react";

const config = await getInfo();

Amplify.configure({
  Auth: {
    Cognito: {
      identityPoolId: "",
      userPoolId: config.userPoolId ?? "",
      userPoolClientId: config.userPoolWebClientId ?? "",
    },
  },
});

const root = ReactDOM.createRoot(
  document.getElementById("root") as HTMLElement,
);
root.render(
  <React.StrictMode>
    <Authenticator.Provider>
      <App />
    </Authenticator.Provider>
  </React.StrictMode>,
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
